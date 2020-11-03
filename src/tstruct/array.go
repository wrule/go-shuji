package tstruct

// Array 结构
type Array struct {
	desc    string
	element ITStruct
}

// Type 结构类型
func (me *Array) Type() ETStructType {
	return TStructArray
}

// Desc 结构描述
func (me *Array) Desc() string {
	return me.desc
}

// Hash 结构Hash值
func (me *Array) Hash() string {
	return "array"
}

// IsBasic 是否是基本结构
func (me *Array) IsBasic() bool {
	return false
}

// TsName TypeScript类型名称
func (me *Array) TsName() string {
	return me.Element().TsName() + "[]"
}

// Element 数组元素结构
func (me *Array) Element() ITStruct {
	return me.element
}

// NewArray Array结构构造函数
func NewArray(desc string, element ITStruct) *Array {
	return &Array{desc, element}
}
