package tstruct

// Date 结构
type Date struct {
	desc string
}

// Type 结构类型
func (me *Date) Type() ETStructType {
	return TStructDate
}

// Desc 结构描述
func (me *Date) Desc() string {
	return me.desc
}

// Hash 结构Hash值
func (me *Date) Hash() string {
	return "null"
}

// IsBasic 是否是基本结构
func (me *Date) IsBasic() bool {
	return true
}

// TsName TypeScript类型名称
func (me *Date) TsName() string {
	return "Date"
}

// NewDate Date结构构造函数
func NewDate(desc string) *Date {
	return &Date{desc}
}
