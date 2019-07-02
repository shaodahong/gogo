package log

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func Info(msg ...interface{}) {
	log.Info(msg...)
}

func Trace(msg ...interface{}) {
	log.Trace(msg...)
}

func Debug(msg ...interface{}) {
	log.Debug(msg...)
}

func Warn(msg ...interface{}) {
	log.Warn(msg...)
}

func Fatal(msg ...interface{}) {
	log.Fatal(msg...)
}

func Panic(msg ...interface{}) {
	log.Panic(msg...)
}

func Error(msg ...interface{}) {
	log.Error(msg...)
}
