package logger

import (
	"testing"
)

func TestLog(t *testing.T) {
	Tracef("aaaa%s222", "111")
	Debugf("aaaa%s222", "111")
	Infof("aaaa%s222", "111")
	Warnf("aaaa%s222", "111")
	Errorf("aaaa%s222", "111")
	Criticalf("aaaa%s222", "111")
}
