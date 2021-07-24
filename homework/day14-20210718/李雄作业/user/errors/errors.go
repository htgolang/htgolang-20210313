package errors

// 利用一个且切片来存储错误信息,并在模板中按需获取
type Errors map[string][]string

func NewErrors() Errors {
	return make(Errors)
}

// 添加错误信息
func (e Errors) AddError(key, err string) {
	e[key] = append(e[key], err)
}

// 删除错误信息
func (e Errors) Clear() {
	keys := make([]string, 0, len(e))
	for k := range e {
		keys = append(keys, k)
	}
	for _, key := range keys {
		delete(e, key)
	}

}
