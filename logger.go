package logger

import "os"
import "log"

var errlog = log.New(os.Stderr, "", log.LstdFlags | log.Lmicroseconds)
var stdlog = log.New(os.Stdout, "", log.LstdFlags | log.Lmicroseconds)
var DEBUG = false

// SetDebug will set the DEBUG variable
func SetDebug(debug bool) {
  DEBUG = debug
}

// IfDebug will return DEBUG
func IfDebug() (debug bool) {
  debug = DEBUG
  return
}

// Fatal will log a message to Stderr and exit with 1
func Fatal(format string, v ...interface{}) {
  errlog.Fatalf(format, v...)
}

// Error will log a message to Stderr
func Error(format string, v ...interface{}) {
  errlog.Printf(format, v...)
}

// Info will log a message to Stdout
func Info(v ...interface{}) {
  stdlog.Println(v...)
}

// Debug will log a message to Stdout if DEBUG is enabled
func Debug(v ...interface{}) {
  if IfDebug() { stdlog.Println(v...) }
}
