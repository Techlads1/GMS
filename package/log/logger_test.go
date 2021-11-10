package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger_WithOptions(t *testing.T) {
	l := New(Development())
	n := l.WithOptions(AddCaller())

	l.Debug("debug original")
	n.Debug("debug new")
}

func TestLogger(t *testing.T) {
	l := New(Development(), AddCaller())

	l.Debug("debug", "test")
	l.Debugf("debugf %s", "test")
	l.Debugln("debugln", "test")
	l.Debugw("debugw", "aaa", 1)

	l.Info("info", "test")
	l.Infof("infof %s", "test")
	l.Infoln("infoln", "test")
	l.Infow("infow", "aaa", 1)

	l.Warn("warn", "test")
	l.Warnf("warnf %s", "test")
	l.Warnln("warnln", "test")
	l.Warnw("warnw", "aaa", 1)

	l.Error("error", "test")
	l.Errorf("errorf %s", "test")
	l.Errorln("errorln", "test")
	l.Errorw("errorw", "aaa", 1)

	assert.Panics(t, func() {
		l.DPanic("dpanic", "test")
	})
	assert.Panics(t, func() {
		l.DPanicf("dpanicf %s", "test")
	})
	assert.Panics(t, func() {
		l.DPanicln("dpanicln", "test")
	})
	assert.Panics(t, func() {
		l.DPanicw("dpanicw", "aaa", 1)
	})

	assert.Panics(t, func() {
		l.Panic("panic", "test")
	})
	assert.Panics(t, func() {
		l.Panicf("panicf %s", "test")
	})
	assert.Panics(t, func() {
		l.Panicln("panicln", "test")
	})
	assert.Panics(t, func() {
		l.Panicw("panic", "aaa", 1)
	})
}

func Test_LoggerPanic(t *testing.T) {
	l := New(Development())
	n := l.WithOptions(WithDevelopment(false))

	assert.Panics(t, func() {
		l.DPanicln("dpanic", "This should panic")
	})

	n.DPanicln("dpanic", "This should not panic")

	assert.Panics(t, func() {
		l.Panicln("panic", "This should panic")
	})

	assert.Panics(t, func() {
		n.Panicln("panic", "This should panic")
	})
}

func TestLogger_LogDir(t *testing.T) {
	l := New(WithLogDirs("log"))

	l.Print("info")
	l.Print("info")
	l.Infoln("info, fff")
	l.Warn("debug")
}
