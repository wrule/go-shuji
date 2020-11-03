package tstruct

// Tuple 结构
type Tuple struct {
	desc     string
	elements []ITStruct
}

// Type 结构类型
func (me *Tuple) Type() ETStructType {
	return TStructTuple
}

// Desc 结构描述
func (me *Tuple) Desc() string {
	return me.desc
}

// Hash 结构Hash值
func (me *Tuple) Hash() string {
	return "tuple"
}

// IsBasic 是否是基本结构
func (me *Tuple) IsBasic() bool {
	return false
}

// TsName TypeScript类型名称
func (me *Tuple) TsName() string {
	var result = "["
	for index, num := range me.Elements() {
		if index > 0 {
			result += ", "
		}
		result += num.TsName()
	}
	result += "]"
	return result
}

// Elements 元组结构元素列表
func (me *Tuple) Elements() []ITStruct {
	return me.elements
}

// NewTuple Tuple结构构造函数
func NewTuple(desc string, elements []ITStruct) *Tuple {
	return &Tuple{desc, elements}
}
