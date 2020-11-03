package tstruct

// Unknow 结构
type Unknow struct {
	desc string
}

// Type 结构类型
func (me *Unknow) Type() ETStructType {
	return TStructUnknow
}

// Desc 结构描述
func (me *Unknow) Desc() string {
	return me.desc
}

// Hash 结构Hash值
func (me *Unknow) Hash() string {
	return "unknow"
}

// IsBasic 是否是基本结构
func (me *Unknow) IsBasic() bool {
	return true
}

// TsName TypeScript类型名称
func (me *Unknow) TsName() string {
	return "any"
}

// NewUnknow Unknow结构构造函数
func NewUnknow(desc string) *Unknow {
	return &Unknow{desc}
}
