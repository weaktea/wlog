package wlog

import (
	"testing"
)
//CMD: go test -v -count=1 -run TestAllLevel -logThreshold=WARNING

func TestDebug(t *testing.T) {
	t.Run("Info", func(t *testing.T) {
		Debug("this is a debug log")
		Debugln("debug ln", 123, "fsef")
		Flush()
	})
}

func TestInfo(t *testing.T) {
	t.Run("Info", func(t *testing.T) {
		Info("this is a Info log")
		Infoln("Infoln", 123, "fsef")
		Flush()
	})
}

func TestAllLevel(t *testing.T){
	t.Run("AllLevel", func(t *testing.T) {
		Debugln("debug ln", 123, "fsef")
		Infoln("Infoln", 123, "fsef")
		Warningln("Warningln", 123, "fsef")
		Errorln("Errorln", 123, "fsef")
		Flush()
	})
}