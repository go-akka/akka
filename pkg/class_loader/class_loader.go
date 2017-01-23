package class_loader

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

var (
	ErrBadClassLoaderParentChain = errors.New("bad class loader parent chain")
)

var (
	Default = NewClassicClassLoader(nil)
)

type ClassLoader interface {
	Register(v interface{}, name string)
	ClassOf(v interface{}) (typ reflect.Type, exist bool)
	ClassNameOf(name string) (typ reflect.Type, exist bool)
	Parent() (loader ClassLoader)
}

type ClassicClassLoader struct {
	parent    ClassLoader
	nameTypes map[string]reflect.Type
	pathTypes map[string]reflect.Type

	locker sync.Mutex
}

func NewClassicClassLoader(parent ClassLoader) ClassLoader {
	loader := &ClassicClassLoader{
		nameTypes: make(map[string]reflect.Type),
		pathTypes: make(map[string]reflect.Type),
	}

	loader.setParent(parent)

	return loader
}

func (p *ClassicClassLoader) Register(v interface{}, name string) {
	p.locker.Lock()
	defer p.locker.Unlock()

	var vType reflect.Type

	switch tVal := v.(type) {
	case reflect.Type:
		{
			vType = tVal
		}
	default:
		vType = reflect.TypeOf(v).Elem()
	}

	path := fmt.Sprintf("%s.%s", vType.PkgPath(), vType.Name())

	if len(name) == 0 {
		name = path
	}

	if _, exist := p.ClassOf(v); !exist {
		p.pathTypes[path] = vType
	}

	if _, exist := p.ClassNameOf(name); !exist {
		p.nameTypes[name] = vType
	}

	return
}

func (p *ClassicClassLoader) ClassOf(v interface{}) (typ reflect.Type, exist bool) {

	vType := reflect.TypeOf(v).Elem()
	path := fmt.Sprintf("%s.%s", vType.PkgPath(), vType.Name())

	if t, exist := p.pathTypes[path]; !exist {
		if p.parent != nil {
			return p.parent.ClassOf(v)
		}
	} else {
		typ = t
		exist = exist
	}

	return
}

func (p *ClassicClassLoader) ClassNameOf(name string) (typ reflect.Type, exist bool) {
	if t, ex := p.nameTypes[name]; !ex {
		if p.parent != nil {
			typ, exist = p.parent.ClassNameOf(name)
			return
		}
	} else {
		typ = t
		exist = ex
	}

	return
}

func (p *ClassicClassLoader) Parent() (loader ClassLoader) {
	loader = p.parent
	return
}

func (p *ClassicClassLoader) setParent(parent ClassLoader) (err error) {
	if parent == nil {
		p.parent = nil
		return
	}

	if parent == p {
		return
	}

	if p.parent == nil {
		p.parent = parent
		return
	}

	parentCount := make(map[ClassLoader]int)

	pa := p.parent
	for pa != nil {
		parentCount[pa] = parentCount[pa] + 1
		pa = p.parent
	}

	for _, count := range parentCount {
		if count > 1 {
			err = ErrBadClassLoaderParentChain
			return
		}
	}

	return
}
