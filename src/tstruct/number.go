package tstruct

// Number 结构
type Number struct {
	desc string
}

// Type 结构类型
func (me *Number) Type() ETStructType {
	return TStructNumber
}

// Desc 结构描述
func (me *Number) Desc() string {
	return me.desc
}

// Hash 结构Hash值
func (me *Number) Hash() string {
	return "null"
}

// IsBasic 是否是基本结构
func (me *Number) IsBasic() bool {
	return true
}

// TsName TypeScript类型名称
func (me *Number) TsName() string {
	return "number"
}

// NewNumber Number结构构造函数
func NewNumber(desc string) *Number {
	return &Number{desc}
}
