package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

type User struct {
	ID   string
	Name string
}

func main() {
	// 定义模板字符串
	// 语法
	// 	显示数据 	{{ . }}   . 传递的数据

	// txt := `{{ . }}`
	// data := "kk"
	// data := 1

	// 条件
	// data true 显示 我叫kk false 显示 无内容
	// txt := `{{ if . }}
	// 我叫kk
	// {{ else }}
	// 无内容
	// {{end}}`
	// data := false
	// data := false
	// txt := `{{ if . }}
	// 我叫kk
	// {{ end }}`

	// 模板函数 index 切片 indexNum
	// txt := `{{ index . 2 }}`

	// 遍历切片
	// range 中的. 表示遍历的每个元素
	// txt := `{{ range . }}
	// 	{{ . }}
	// {{ end }}`
	// 	txt := `{{ range $i := . }}
	// 	{{ $i }}
	// {{ end }}`
	// txt := `{{ range $i, $v := . }}
	// 	{{ $i }}: {{ $v }}
	// {{ end }}`
	// data := []int{}

	// txt := `{{ len .}}`
	// 遍历
	// txt := `{{ range . }}
	// {{ . }}
	// {{ end }}`
	// txt := `{{- range . -}}
	// {{ . }}
	// {{- else -}}
	// 无数据
	// {{- end -}}`

	// txt := `{{ range $v := . }}
	// {{ $v }}
	// {{ end }}`
	// txt := `{{ range $k, $v := . }}
	// {{ $k }} : {{ $v }}
	// {{ end }}`
	// data := map[string]string{"kk": "西安", "lujianfeng": "北京"}
	// 想获取Name属性并显示
	// txt := `登陆：{{ .Name }}`
	// data := User{"1", "kk"}

	// data := []User{{"1", "kk"}, {"2", "lujianfeng"}}

	// 条件判断 > < == != >= <= gt lt eq ne ge le
	// 显示成绩 级别 >= 90 A >= 80 B >= 70 C 其他 D

	// txt := `
	// {{ if ge . 90 }}
	// A
	// {{ else if ge . 80 }}
	// B
	// {{ else if ge . 70}}
	// C
	// {{ else }}
	// D
	// {{ end }}
	// `
	// data := 60

	// 显示表格
	// txt := `
	// <table>
	// 	<thead>
	// 		<tr>
	// 			<th>ID</th>
	// 			<th>Name</th>
	// 		</tr>
	// 	</thead>
	// 	<tbody>
	// 	{{ range . }}
	// 		<tr>
	// 			<td>{{ .ID }}</td>
	// 			<td>{{ .Name }}</td>
	// 		</tr>
	// 	{{ end }}
	// 	</tbody>
	// </table>
	// `

	// 内置函数
	// index len

	// 自定义函数 upper lower
	txt := `{{ . }} {{ upper . }} {{ lower . }} {{ tl . true }} {{ tl . false }}`
	data := "Abc"

	funcs := template.FuncMap{
		"upper": func(txt string) string {
			return strings.ToUpper(txt)
		},
		"lower": strings.ToLower,
		"tl": func(txt string, isLower bool) string {
			if isLower {
				return strings.ToLower(txt)
			}
			return strings.ToUpper(txt)
		},
	}

	// 创建模板
	tpl := template.New("txtTemplate")
	tpl = tpl.Funcs(funcs)

	// 解析模板字符串
	tpl, err := tpl.Parse(txt)
	if err != nil {
		log.Fatal(err)
	}
	// 执行 输出到控制台
	err = tpl.Execute(os.Stdout, data)
	fmt.Println()
	fmt.Println(err)
}
