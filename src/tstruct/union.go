package tstruct

// Union 结构
type Union struct {
	desc    string
	members []ITStruct
}

// Type 结构类型
func (me *Union) Type() ETStructType {
	return TStructUnion
}

// Desc 结构描述
func (me *Union) Desc() string {
	return me.desc
}

// Hash 结构Hash值
func (me *Union) Hash() string {
	return "union"
}

// IsBasic 是否是基本结构
func (me *Union) IsBasic() bool {
	return false
}

// TsName TypeScript类型名称
func (me *Union) TsName() string {
	var result = "("
	for index, num := range me.Members() {
		if index > 0 {
			result += " | "
		}
		result += num.TsName()
	}
	result += ")"
	return result
}

// Members 联合结构成员列表
func (me *Union) Members() []ITStruct {
	return me.members
}

// NewUnion Union结构构造函数
func NewUnion(desc string, members []ITStruct) *Union {
	return &Union{desc, members}
}
