package codegen

import (
	"github.com/Numenorean/thrift-iterator/spi"
	"reflect"
)

type Extension struct {
	spi.Extension
	ExtTypes []reflect.Type
}

func (ext *Extension) MangledName() string {
	// TODO: hash extension to represent different config
	return "default"
}