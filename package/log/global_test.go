package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGlobalSetOptions(t *testing.T) {
	SetOptions(Development(), WithCaller(true))

	Debug("debug")
}

func TestGlobalLogger(t *testing.T) {
	SetOptions(Development(), WithCaller(true))

	Debug("global: debug", "test")
	Debugf("global: debugf %s", "test")
	Debugln("global: debugln", "test")
	Debugw("global: debugw", "aaa", 1)

	Info("global: info", "test")
	Info("global: info", "test")
	Info("global: info", 1, 1)
	Info("global: info", 1, "test")
	Infof("global: infof %s", "test")
	Infoln("global: infoln", "test")
	Infoln("global: infoln", 1, "test")
	Infow("global: info", "aaa", 1)

	Warn("global: debug", "test")
	Warnf("global: debugf %s", "test")
	Warnln("global: debugln", "test")
	Warnw("global: debugw", "aaa", 1)

	Error("global: error", "test")
	Errorf("global: errorf %s", "test")
	Errorln("global: errorln", "test")
	Errorw("global: errorw", "aaa", 1)

	assert.Panics(t, func() {
		DPanic("global: dpanic", "test")
	})
	assert.Panics(t, func() {
		DPanicf("global: dpanicf %s", "test")
	})
	assert.Panics(t, func() {
		DPanicln("global: dpanicln", "test")
	})
	assert.Panics(t, func() {
		DPanicw("global: dpanicw", "aaa", 1)
	})
}

func TestPanic(t *testing.T) {
	SetOptions(Development(), WithCaller(true))

	assert.Panics(t, func() {
		DPanic("global: dpanic ", "this should panic")
	})
	assert.Panics(t, func() {
		Panic("global: panic ", "this should panic")
	})

	SetOptions(WithDevelopment(false))

	DPanic("global: dpanic ", "this should not panic")
	assert.Panics(t, func() {
		Panic("global: panic ", "this should panic")
	})
}
