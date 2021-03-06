package tstruct

// Undefined 结构
type Undefined struct {
	desc string
}

// Type 结构类型
func (me *Undefined) Type() ETStructType {
	return TStructUndefined
}

// Desc 结构描述
func (me *Undefined) Desc() string {
	return me.desc
}

// Hash 结构Hash值
func (me *Undefined) Hash() string {
	return "undefined"
}

// IsBasic 是否是基本结构
func (me *Undefined) IsBasic() bool {
	return true
}

// TsName TypeScript类型名称
func (me *Undefined) TsName() string {
	return "undefined"
}

// NewUndefined Undefined结构构造函数
func NewUndefined(desc string) *Undefined {
	return &Undefined{desc}
}
