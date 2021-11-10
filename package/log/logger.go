package log

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type logFunc func(logger *zap.SugaredLogger, args ...interface{})
type logfFunc func(logger *zap.SugaredLogger, format string, args ...interface{})
type logwFunc func(logger *zap.SugaredLogger, msg string, keysAndValues ...interface{})

type Logger struct {
	base    *zap.SugaredLogger
	writers []*lumberjack.Logger
	options options

	print  logFunc
	printf logfFunc
	printw logwFunc

	debug  logFunc
	debugf logfFunc
	debugw logwFunc

	info  logFunc
	infof logfFunc
	infow logwFunc

	warn  logFunc
	warnf logfFunc
	warnw logwFunc

	error  logFunc
	errorf logfFunc
	errorw logwFunc

	dpanic  logFunc
	dpanicf logfFunc
	dpanicw logwFunc

	panic  logFunc
	panicf logfFunc
	panicw logwFunc

	fatal  logFunc
	fatalf logfFunc
	fatalw logwFunc
}

var defaultOptions = options{
	RotationConfig: RotationConfig{
		MaxAge:     28,
		MaxBackups: 7,
		MaxSize:    500,
		LocalTime:  true,
		Compress:   false,
	},

	Development: false,
	Format:      FormatJSON,
	Level:       InfoLevel,
	LogToStdout: true,
	AddCaller:   false,
	CallerSkip:  1,
}

func New(opt ...Option) *Logger {
	opts := defaultOptions

	for _, o := range opt {
		o.apply(&opts)
	}

	return newLogger(opts)
}

func newLogger(opts options) *Logger {
	var encoder zapcore.Encoder

	if opts.Encoder != nil {
		encoder = opts.Encoder
	} else {
		var encoderCfg zapcore.EncoderConfig

		if opts.Development {
			encoderCfg = zap.NewDevelopmentEncoderConfig()
			encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder
			encoderCfg.EncodeCaller = zapcore.FullCallerEncoder
		} else {
			encoderCfg = zap.NewProductionEncoderConfig()
			encoderCfg.TimeKey = "time"
			encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder
		}

		switch opts.Format {
		case FormatJSON:
			encoder = zapcore.NewJSONEncoder(encoderCfg)
		default:
			encoder = zapcore.NewConsoleEncoder(encoderCfg)
		}
	}

	cores := make([]zapcore.Core, 0)
	writers := make([]*lumberjack.Logger, 0)

	// add stdout log
	if opts.LogToStdout {
		stdoutCore := zapcore.NewCore(
			encoder,
			zapcore.Lock(os.Stdout),
			zap.LevelEnablerFunc(opts.ZapLevelEnabled),
		)
		cores = append(cores, stdoutCore)
	}

	// add output core
	if opts.Output != nil {
		outputCore := zapcore.NewCore(
			encoder,
			zapcore.Lock(zapcore.AddSync(opts.Output)),
			zap.LevelEnablerFunc(opts.ZapLevelEnabled),
		)
		cores = append(cores, outputCore)
	}

	// parse log dirs
	for _, dir := range opts.LogDirs {
		if dir == "" {
			continue
		}

		for _, level := range []zapcore.Level{
			zapcore.DebugLevel,
			zapcore.InfoLevel,
			zapcore.WarnLevel,
			zapcore.ErrorLevel,
			zapcore.DPanicLevel,
			zapcore.PanicLevel,
			zapcore.FatalLevel,
		} {
			if opts.ZapLevelEnabled(level) {
				lvl := level

				lvlWriter := newRotateWriter(
					filepath.Join(dir, fmt.Sprint(lvl.String(), ".log")),
					opts.RotationConfig,
				)

				lvlCore := zapcore.NewCore(
					encoder,
					zapcore.AddSync(lvlWriter),
					zap.LevelEnablerFunc(func(l zapcore.Level) bool {
						return l == lvl
					}),
				)

				cores = append(cores, lvlCore)
				writers = append(writers, lvlWriter)
			}
		}
	}

	// parse log files
	for _, file := range opts.LogFiles {
		if file == "" {
			continue
		}

		writer := newRotateWriter(file, opts.RotationConfig)

		fileCore := zapcore.NewCore(
			encoder,
			zapcore.AddSync(writer),
			zap.LevelEnablerFunc(opts.ZapLevelEnabled),
		)

		cores = append(cores, fileCore)
		writers = append(writers, writer)
	}

	zapOptions := []zap.Option{
		zap.WithCaller(opts.AddCaller),
		zap.AddCallerSkip(opts.CallerSkip),
	}

	if opts.Development {
		zapOptions = append(zapOptions, zap.Development())
	}

	l := &Logger{
		base:    zap.New(zapcore.NewTee(cores...), zapOptions...).Sugar(),
		writers: writers,
		options: opts,

		debug:   (*zap.SugaredLogger).Debug,
		debugf:  (*zap.SugaredLogger).Debugf,
		debugw:  (*zap.SugaredLogger).Debugw,
		info:    (*zap.SugaredLogger).Info,
		infof:   (*zap.SugaredLogger).Infof,
		infow:   (*zap.SugaredLogger).Infow,
		warn:    (*zap.SugaredLogger).Warn,
		warnf:   (*zap.SugaredLogger).Warnf,
		warnw:   (*zap.SugaredLogger).Warnw,
		error:   (*zap.SugaredLogger).Error,
		errorf:  (*zap.SugaredLogger).Errorf,
		errorw:  (*zap.SugaredLogger).Errorw,
		dpanic:  (*zap.SugaredLogger).DPanic,
		dpanicf: (*zap.SugaredLogger).DPanicf,
		dpanicw: (*zap.SugaredLogger).DPanicw,
		panic:   (*zap.SugaredLogger).Panic,
		panicf:  (*zap.SugaredLogger).Panicf,
		panicw:  (*zap.SugaredLogger).Panicw,
		fatal:   (*zap.SugaredLogger).Fatal,
		fatalf:  (*zap.SugaredLogger).Fatalf,
		fatalw:  (*zap.SugaredLogger).Fatalw,
	}

	if opts.Development {
		l.print = (*zap.SugaredLogger).Debug
		l.printf = (*zap.SugaredLogger).Debugf
		l.printw = (*zap.SugaredLogger).Debugw
	} else {
		l.print = (*zap.SugaredLogger).Info
		l.printf = (*zap.SugaredLogger).Infof
		l.printw = (*zap.SugaredLogger).Infow
	}

	return l
}

func newRotateWriter(filename string, config RotationConfig) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    config.MaxSize,
		MaxAge:     config.MaxAge,
		MaxBackups: config.MaxBackups,
		LocalTime:  config.LocalTime,
		Compress:   config.Compress,
	}
}

func (l *Logger) WithOptions(opt ...Option) *Logger {
	opts := l.options.Clone()

	for _, o := range opt {
		o.apply(&opts)
	}

	return newLogger(opts)
}

func (l *Logger) Print(args ...interface{}) {
	l.print(l.base, args...)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.printf(l.base, format, args...)
}

func (l *Logger) Println(args ...interface{}) {
	l.print(l.base, sprintln(args...))
}

func (l *Logger) Printw(msg string, keysAndValues ...interface{}) {
	l.printw(l.base, msg, keysAndValues...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.debug(l.base, args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.debugf(l.base, format, args...)
}

func (l *Logger) Debugln(args ...interface{}) {
	l.debug(l.base, sprintln(args...))
}

func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.debugw(l.base, msg, keysAndValues...)
}

func (l *Logger) Info(args ...interface{}) {
	l.info(l.base, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.infof(l.base, format, args...)
}

func (l *Logger) Infoln(args ...interface{}) {
	l.info(l.base, sprintln(args...))
}

func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.infow(l.base, msg, keysAndValues...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.warn(l.base, args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.warnf(l.base, template, args...)
}

func (l *Logger) Warnln(args ...interface{}) {
	l.warn(l.base, sprintln(args...))
}

func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.warnw(l.base, msg, keysAndValues...)
}

func (l *Logger) Error(args ...interface{}) {
	l.error(l.base, args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.errorf(l.base, template, args...)
}

func (l *Logger) Errorln(args ...interface{}) {
	l.error(l.base, sprintln(args...))
}

func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.errorw(l.base, msg, keysAndValues...)
}

func (l *Logger) DPanic(args ...interface{}) {
	l.dpanic(l.base, args...)
}

func (l *Logger) DPanicf(template string, args ...interface{}) {
	l.dpanicf(l.base, template, args...)
}

func (l *Logger) DPanicln(args ...interface{}) {
	l.dpanic(l.base, sprintln(args...))
}

func (l *Logger) DPanicw(msg string, keysAndValues ...interface{}) {
	l.dpanicw(l.base, msg, keysAndValues...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.panic(l.base, args...)
}

func (l *Logger) Panicf(template string, args ...interface{}) {
	l.panicf(l.base, template, args...)
}

func (l *Logger) Panicln(args ...interface{}) {
	l.panic(l.base, sprintln(args...))
}

func (l *Logger) Panicw(msg string, keysAndValues ...interface{}) {
	l.panicw(l.base, msg, keysAndValues...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.fatal(l.base, args...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.fatalf(l.base, template, args...)
}

func (l *Logger) Fatalln(args ...interface{}) {
	l.fatal(l.base, sprintln(args...))
}

func (l *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.fatalw(l.base, msg, keysAndValues...)
}

func (l *Logger) Rotate() error {
	for _, w := range l.writers {
		if err := w.Rotate(); err != nil {
			return err
		}
	}

	return nil
}

func sprintln(args ...interface{}) string {
	return strings.TrimSuffix(fmt.Sprintln(args...), "\n")
}
