package tstruct

import (
	"strings"
)

// Object 结构
type Object struct {
	desc string
}

// Type 结构类型
func (me *Object) Type() ETStructType {
	return TStructObject
}

// Desc 结构描述
func (me *Object) Desc() string {
	return me.desc
}

// Hash 结构Hash值
func (me *Object) Hash() string {
	return "object"
}

// IsBasic 是否是基本结构
func (me *Object) IsBasic() bool {
	return false
}

// ModuleName 模块名称
func (me *Object) ModuleName() string {
	return strings.Title(strings.ToLower(me.Desc()))
}

// InterfaceName 接口名称
func (me *Object) InterfaceName() string {
	return "I" + me.ModuleName()
}

// NewObject Object结构构造函数
func NewObject(desc string) *Object {
	return &Object{desc}
}
