package logger

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	DebugLevel int = 0
	InfoLevel  int = 1
	WarnLevel  int = 2
	ErrorLevel int = 3
	FatalLevel int = 4
	PanicLevel int = 5

	TextFormat string = "text"
	JSONFormat string = "json"

	TimestampFieldName string = "timestamp"
	MessageFieldName   string = "message"
)

type Logger struct {
	Logger *logrus.Logger
	out    io.Writer
	level  int
	fmt    string
}

type Fields map[string]interface{}

var (
	NoFields = make(map[string]interface{})
)

func NewDefault() (*Logger, error) {
	return New(os.Stdout, DebugLevel, TextFormat)
}

func New(out io.Writer, level int, fmt string) (*Logger, error) {
	log := &Logger{
		Logger: logrus.New(),
		out:    out,
		level:  level,
		fmt:    fmt,
	}
	log.Logger.Out = out
	if err := log.SetFormatter(fmt); err != nil {
		return nil, err
	}
	if err := log.SetLogLevel(level); err != nil {
		return nil, err
	}

	return log, nil
}

func NewLogfile(filename string, level int, fmt string) (*Logger, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	return New(file, level, fmt)
}

func GetLogLevel(level string) int {
	switch strings.ToLower(level) {
	case "panic":
		return PanicLevel
	case "fatal":
		return FatalLevel
	case "error":
		return ErrorLevel
	case "warn":
		return WarnLevel
	case "info":
		return InfoLevel
	default:
		return DebugLevel
	}
}

func (this *Logger) SetFormatter(fmt string) error {
	var lf logrus.Formatter
	switch fmt {
	case TextFormat:
		lf = &logrus.TextFormatter{}
	case JSONFormat:
		lf = &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime: TimestampFieldName,
				logrus.FieldKeyMsg:  MessageFieldName,
			},
		}
	default:
		return errors.New("Unknown format!")
	}
	this.Logger.Formatter = lf
	return nil
}

func (this *Logger) SetLogLevel(level int) error {
	var ll logrus.Level
	switch level {
	case DebugLevel:
		ll = logrus.DebugLevel
	case InfoLevel:
		ll = logrus.InfoLevel
	case WarnLevel:
		ll = logrus.WarnLevel
	case ErrorLevel:
		ll = logrus.ErrorLevel
	case FatalLevel:
		ll = logrus.FatalLevel
	case PanicLevel:
		ll = logrus.PanicLevel
	default:
		return errors.New("Unknown log level!")
	}

	this.Logger.Level = ll

	return nil
}

func (this *Logger) GetLogrus() *logrus.Logger {
	return this.Logger
}

func (this *Logger) Panic(args ...interface{}) {
	this.Logger.Panic(args...)
}

func (this *Logger) Fatal(args ...interface{}) {
	this.Logger.Fatal(args...)
}

func (this *Logger) Error(args ...interface{}) {
	this.Logger.Error(args...)
}

func (this *Logger) Warn(args ...interface{}) {
	this.Logger.Warn(args...)
}

func (this *Logger) Info(args ...interface{}) {
	this.Logger.Info(args...)
}

func (this *Logger) Debug(args ...interface{}) {
	this.Logger.Debug(args...)
}

func (this *Logger) PanicWithFields(fields Fields, args ...interface{}) {
	this.Logger.WithFields(logrus.Fields(fields)).Panic(args...)
}

func (this *Logger) FatalWithFields(fields Fields, args ...interface{}) {
	this.Logger.WithFields(logrus.Fields(fields)).Fatal(args...)
}

func (this *Logger) ErrorWithFields(fields Fields, args ...interface{}) {
	this.Logger.WithFields(logrus.Fields(fields)).Error(args...)
}

func (this *Logger) WarnWithFields(fields Fields, args ...interface{}) {
	this.Logger.WithFields(logrus.Fields(fields)).Warn(args...)
}

func (this *Logger) InfoWithFields(fields Fields, args ...interface{}) {
	this.Logger.WithFields(logrus.Fields(fields)).Info(args...)
}

func (this *Logger) DebugWithFields(fields Fields, args ...interface{}) {
	this.Logger.WithFields(logrus.Fields(fields)).Debug(args...)
}

func (this *Logger) Panicf(fmt string, args ...interface{}) {
	this.Logger.Panicf(fmt, args...)
}

func (this *Logger) Fatalf(fmt string, args ...interface{}) {
	this.Logger.Fatalf(fmt, args...)
}

func (this *Logger) Errorf(fmt string, args ...interface{}) {
	this.Logger.Errorf(fmt, args...)
}

func (this *Logger) Warnf(fmt string, args ...interface{}) {
	this.Logger.Warnf(fmt, args...)
}

func (this *Logger) Infof(fmt string, args ...interface{}) {
	this.Logger.Infof(fmt, args...)
}

func (this *Logger) Debugf(fmt string, args ...interface{}) {
	this.Logger.Debugf(fmt, args...)
}
