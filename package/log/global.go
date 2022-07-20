package log

var global = New()

func SetOptions(opts ...Option) {
	global = global.WithOptions(opts...)
}

func Debug(args ...interface{}) {
	global.debug(global.base, args...)
}

func Debugf(format string, args ...interface{}) {
	global.debugf(global.base, format, args...)
}

func Debugln(args ...interface{}) {
	global.debug(global.base, sprintln(args...))
}

func Debugw(msg string, keysAndValues ...interface{}) {
	global.debugw(global.base, msg, keysAndValues...)
}

func Info(args ...interface{}) {
	global.info(global.base, args...)
}

func Infof(format string, args ...interface{}) {
	global.infof(global.base, format, args...)
}

func Infoln(args ...interface{}) {
	global.info(global.base, sprintln(args...))
}

func Infow(msg string, keysAndValues ...interface{}) {
	global.infow(global.base, msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	global.warn(global.base, args...)
}

func Warnf(template string, args ...interface{}) {
	global.warnf(global.base, template, args...)
}

func Warnln(args ...interface{}) {
	global.warn(global.base, sprintln(args...))
}

func Warnw(msg string, keysAndValues ...interface{}) {
	global.warnw(global.base, msg, keysAndValues...)
}

func Error(args ...interface{}) {
	global.error(global.base, args...)
}

func Errorf(template string, args ...interface{}) {
	global.errorf(global.base, template, args...)
}

func Errorln(args ...interface{}) {
	global.error(global.base, sprintln(args...))
}

func Errorw(msg string, keysAndValues ...interface{}) {
	global.errorw(global.base, msg, keysAndValues...)
}

func DPanic(args ...interface{}) {
	global.dpanic(global.base, args...)
}

func DPanicf(template string, args ...interface{}) {
	global.dpanicf(global.base, template, args...)
}

func DPanicln(args ...interface{}) {
	global.dpanic(global.base, sprintln(args...))
}

func DPanicw(msg string, keysAndValues ...interface{}) {
	global.dpanicw(global.base, msg, keysAndValues...)
}

func Panic(args ...interface{}) {
	global.panic(global.base, args...)
}

func Panicf(template string, args ...interface{}) {
	global.panicf(global.base, template, args...)
}

func Panicln(args ...interface{}) {
	global.panic(global.base, sprintln(args...))
}

func Panicw(msg string, keysAndValues ...interface{}) {
	global.panicw(global.base, msg, keysAndValues...)
}

func Fatal(args ...interface{}) {
	global.fatal(global.base, args...)
}

func Fatalf(template string, args ...interface{}) {
	global.fatalf(global.base, template, args...)
}

func Fatalln(args ...interface{}) {
	global.fatal(global.base, sprintln(args...))
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	global.fatalw(global.base, msg, keysAndValues...)
}

func Rotate() error {
	return global.Rotate()
}
