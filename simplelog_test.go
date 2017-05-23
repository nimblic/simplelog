package simplelog

import (
	"testing"
	"time"
)

var loc *time.Location

func TestLog(t *testing.T) {
	l, _ := ParseLevel("debug")
	Init(l)
	Debugf("testing debug log with param %s", "p1")
	Infof("testing info log with param %s", "p1")
	Noticef("testing notice log with param %s", "p1")
	Warningf("testing warning log with param %s", "p1")
	Errorf("testing error log with param %s", "p1")

	m := make(map[string]string)
	m["k1"] = "v1"
	m["k2"] = "v2"
	Debugm("Some values", m)

	l, _ = ParseLevel("info")
	Init(l)
	Debugf("testing debug log with param %s", "p1")
	Infof("testing info log with param %s", "p1")
	Noticef("testing notice log with param %s", "p1")
	Warningf("testing warning log with param %s", "p1")
	Errorf("testing error log with param %s", "p1")
	l, _ = ParseLevel("error")
	Init(l)
	Debugf("testing debug log with param %s", "p1")
	Infof("testing info log with param %s", "p1")
	Noticef("testing notice log with param %s", "p1")
	Warningf("testing warning log with param %s", "p1")
	Errorf("testing error log with param %s", "p1")

}
