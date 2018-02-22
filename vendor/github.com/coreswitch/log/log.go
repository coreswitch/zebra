// Copyright 2017 CoreSwitch
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"
)

// Output source file information "filename:line number" to the log.
var SourceField bool = true

// Output function name to the log.
var FuncField bool = true

//
func SetLevel(level string) error {
	l, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	origLogger.SetLevel(l)
	return nil
}

//
func SetOutput(output string) error {
	switch output {
	case "stdout":
		origLogger.Out = os.Stdout
	case "stderr":
		origLogger.Out = os.Stderr
	default:
		return fmt.Errorf("output format error")
	}
	return nil
}

//
func SetJSONFormatter() {
	origLogger.Formatter = &logrus.JSONFormatter{}
}

//
func SetTextFormatter() {
	origLogger.Formatter = &logrus.TextFormatter{
		DisableTimestamp: true,
	}
}

type Logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Warn(...interface{})
	Warnf(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
	With(key string, value interface{}) Logger
}

type logger struct {
	entry *logrus.Entry
}

var origLogger = logrus.New()
var baseLogger = logger{entry: logrus.NewEntry(origLogger)}

func (l logger) sourced() *logrus.Entry {
	e := l.entry

	if !SourceField && !FuncField {
		return e
	}

	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return e
	}

	if SourceField {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
		e = e.WithField("source", fmt.Sprintf("%s:%d", file, line))
	}

	if FuncField {
		if f := runtime.FuncForPC(pc); f != nil {
			e = e.WithField("func", f.Name())
		}
	}

	return e
}

func (l logger) Debug(args ...interface{}) {
	l.sourced().Debug(args...)
}

func (l logger) Debugf(format string, args ...interface{}) {
	l.sourced().Debugf(format, args...)
}

func (l logger) Info(args ...interface{}) {
	l.sourced().Info(args...)
}

func (l logger) Infof(format string, args ...interface{}) {
	l.sourced().Infof(format, args...)
}

func (l logger) Warn(args ...interface{}) {
	l.sourced().Warn(args...)
}

func (l logger) Warnf(format string, args ...interface{}) {
	l.sourced().Warnf(format, args...)
}

func (l logger) Error(args ...interface{}) {
	l.sourced().Error(args...)
}

func (l logger) Errorf(format string, args ...interface{}) {
	l.sourced().Errorf(format, args...)
}

func (l logger) With(key string, value interface{}) Logger {
	return logger{l.entry.WithField(key, value)}
}

func With(key string, value interface{}) Logger {
	return baseLogger.With(key, value)
}

func Debug(args ...interface{}) {
	baseLogger.sourced().Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	baseLogger.sourced().Debugf(format, args...)
}

func Info(args ...interface{}) {
	baseLogger.sourced().Info(args...)
}

func Infof(format string, args ...interface{}) {
	baseLogger.sourced().Infof(format, args...)
}

func Warn(args ...interface{}) {
	baseLogger.sourced().Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	baseLogger.sourced().Warnf(format, args...)
}

func Error(args ...interface{}) {
	baseLogger.sourced().Error(args...)
}

func Errorf(format string, args ...interface{}) {
	baseLogger.sourced().Errorf(format, args...)
}
