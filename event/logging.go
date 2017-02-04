package event

import (
	"github.com/go-akka/akka"
	"reflect"
	"strconv"
)

func LogClassFor(level akka.LogLevel) reflect.Type {

	switch level {
	case akka.DebugLevel:
		return reflect.TypeOf((*Debug)(nil)).Elem()
	case akka.InfoLevel:
		return reflect.TypeOf((*Info)(nil)).Elem()
	case akka.WarningLevel:
		return reflect.TypeOf((*Warning)(nil)).Elem()
	case akka.ErrorLevel:
		return reflect.TypeOf((*Error)(nil)).Elem()
	}

	panic("Unknown LogLevel: " + strconv.Itoa(int(level)))
}
