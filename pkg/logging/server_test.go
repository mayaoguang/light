package logging

import "testing"

func TestNewLogger(t *testing.T) {
	Init(Conf{"../../logs/light.log", ConsoleEncoder})
	defer Sync()
	Debug("debug msg")
	Debugf("debugf %s", "light")
	Info("info msg")
	Infof("infof %d", 10)
	Warn("warn msg")
	Warnf("warnf %v", true)
	Error("err msg")
	Errorf("errorf %v", []int{1, 2, 3})
	Fatal("fatal msg")
}

func TestFatalf(t *testing.T) {
	Init(Conf{"../../logs/light.log", JsonEncoder})
	defer Sync()
	Error("err msg")
	Errorf("errorf %v", []int{1, 2, 3})
	Fatalf("fatalf %v", map[string]interface{}{"name": "master"})
}
