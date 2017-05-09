package simplelog

import (
	"testing"
	"time"
)

var loc *time.Location

func TestLog(t *testing.T) {

	Warnf("testing log with param %s", "p1")
}
