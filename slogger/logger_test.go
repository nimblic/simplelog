package slogger

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"testing"
)

// decode parses the single JSON log line in buf.
func decode(t *testing.T, buf *bytes.Buffer) map[string]any {
	t.Helper()
	if buf.Len() == 0 {
		t.Fatal("no log output produced")
	}
	var rec map[string]any
	if err := json.Unmarshal(buf.Bytes(), &rec); err != nil {
		t.Fatalf("output is not valid JSON: %v\noutput: %s", err, buf.String())
	}
	return rec
}

// TestAuditEmittedAtServerLevel covers the placement contract: AUDIT sits
// between INFO (0) and NOTICE (2), so it is emitted at the server's debug
// runtime level but filtered once the configured level rises to NOTICE or above.
func TestAuditEmittedAtServerLevel(t *testing.T) {
	// Emitted at debug (production runtime level).
	var emitted bytes.Buffer
	h, _, err := newHandler(Config{Level: "debug", Format: "json", SupportCustomLevels: true}, &emitted)
	if err != nil {
		t.Fatalf("newHandler: %v", err)
	}
	slog.New(h).Log(context.Background(), LevelAudit, "audit line")
	if rec := decode(t, &emitted); rec["msg"] != "audit line" {
		t.Fatalf("audit record not emitted at debug; got msg=%v", rec["msg"])
	}

	// Filtered at notice and above (LevelAudit=1 < LevelNotice=2).
	for _, lvl := range []string{"notice", "warn", "error"} {
		var dropped bytes.Buffer
		h, _, err := newHandler(Config{Level: lvl, Format: "json", SupportCustomLevels: true}, &dropped)
		if err != nil {
			t.Fatalf("newHandler(%s): %v", lvl, err)
		}
		slog.New(h).Log(context.Background(), LevelAudit, "audit line")
		if dropped.Len() != 0 {
			t.Fatalf("audit unexpectedly emitted at level %q: %s", lvl, dropped.String())
		}
	}
}

// TestAuditRendering covers the level-name rendering in both SupportCustomLevels
// modes: "AUDIT" when on, slog's default "INFO+1" passthrough when off (audit=1).
func TestAuditRendering(t *testing.T) {
	tests := []struct {
		name        string
		customLevel bool
		wantLevel   string
	}{
		{name: "custom levels on", customLevel: true, wantLevel: "AUDIT"},
		{name: "custom levels off", customLevel: false, wantLevel: "INFO+1"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			h, _, err := newHandler(Config{Level: "debug", Format: "json", SupportCustomLevels: tc.customLevel}, &buf)
			if err != nil {
				t.Fatalf("newHandler: %v", err)
			}
			slog.New(h).Log(context.Background(), LevelAudit, "audit line")

			rec := decode(t, &buf)
			if rec["level"] != tc.wantLevel {
				t.Fatalf("level = %q, want %q", rec["level"], tc.wantLevel)
			}
		})
	}
}

// TestCustomLevelRenderingUnchanged guards against regressions in the existing
// NOTICE/VERBOSE mapping when the AUDIT case was added.
func TestCustomLevelRenderingUnchanged(t *testing.T) {
	tests := []struct {
		level     slog.Level
		wantLevel string
	}{
		{LevelNotice, "NOTICE"},
		{LevelVerbose, "VERBOSE"},
		{LevelAudit, "AUDIT"},
		{slog.LevelInfo, "INFO"},
		{slog.LevelError, "ERROR"},
	}
	for _, tc := range tests {
		var buf bytes.Buffer
		h, _, err := newHandler(Config{Level: "verbose", Format: "json", SupportCustomLevels: true}, &buf)
		if err != nil {
			t.Fatalf("newHandler: %v", err)
		}
		slog.New(h).Log(context.Background(), tc.level, "line")
		rec := decode(t, &buf)
		if rec["level"] != tc.wantLevel {
			t.Fatalf("level %d rendered as %q, want %q", tc.level, rec["level"], tc.wantLevel)
		}
	}
}

// TestSetLevelValidation confirms SetLevel rejects unknown level names (so the
// runtime SetLogLevel endpoint can return 400) and accepts the valid ones.
func TestSetLevelValidation(t *testing.T) {
	for _, bad := range []string{"", "bogus", "trace", "audit"} {
		if err := SetLevel(bad); err == nil {
			t.Fatalf("SetLevel(%q) = nil, want error", bad)
		}
	}
	for _, good := range []string{"debug", "info", "notice", "warn", "warning", "error", "verbose"} {
		if err := SetLevel(good); err != nil {
			t.Fatalf("SetLevel(%q) = %v, want nil", good, err)
		}
	}
}

// TestLevelFromString covers the strict name→level mapping.
func TestLevelFromString(t *testing.T) {
	cases := map[string]slog.Level{
		"debug": slog.LevelDebug, "info": slog.LevelInfo, "warn": slog.LevelWarn,
		"warning": slog.LevelWarn, "error": slog.LevelError, "notice": LevelNotice, "verbose": LevelVerbose,
	}
	for name, want := range cases {
		if got, ok := levelFromString(name); !ok || got != want {
			t.Fatalf("levelFromString(%q) = (%v,%v), want (%v,true)", name, got, ok, want)
		}
	}
	if _, ok := levelFromString("audit"); ok {
		t.Fatal("levelFromString(\"audit\") ok=true, want false (not selectable)")
	}
}

// TestAuditAttrsSurvive confirms structured attributes (including a nested
// raw-JSON payload such as clientAudit emits) pass through intact.
func TestAuditAttrsSurvive(t *testing.T) {
	var buf bytes.Buffer
	h, _, err := newHandler(Config{Level: "debug", Format: "json", SupportCustomLevels: true}, &buf)
	if err != nil {
		t.Fatalf("newHandler: %v", err)
	}
	slog.New(h).Log(context.Background(), LevelAudit, "task created",
		"taskId", "abc-123",
		"clientAudit", json.RawMessage(`{"action":"create","ok":true}`),
	)

	rec := decode(t, &buf)
	if rec["taskId"] != "abc-123" {
		t.Fatalf("taskId attr lost; got %v", rec["taskId"])
	}
	ca, ok := rec["clientAudit"].(map[string]any)
	if !ok {
		t.Fatalf("clientAudit not an object; got %T (%v)", rec["clientAudit"], rec["clientAudit"])
	}
	if ca["action"] != "create" || ca["ok"] != true {
		t.Fatalf("clientAudit payload altered; got %v", ca)
	}
}
