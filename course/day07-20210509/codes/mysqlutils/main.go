package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

// 生成结构体对应的创建sql语句
// sql:"type(varchar);size(32);pk;null;default()"
type User struct {
	Id   int    `sql:"pk"`
	Name string `sql:"type(text)"`
	Addr string `sql:"type(varchar);size(1024);null;default(西安)"`
	Desc string
}

func (u User) TableName() string {
	return "account"
}

type Department struct {
	Id   int `sql:"pk"`
	Name string
	Addr string
}

func snake(name string) string {
	return strings.ToLower(name)
}

func tableName(value reflect.Value) string {
	tableName := value.Elem().Type().Name()
	methodValue := value.MethodByName("TableName")
	if !methodValue.IsValid() {
		return tableName
	}
	methodType := methodValue.Type()
	if 0 != methodType.NumIn() || 1 != methodType.NumOut() {
		return tableName
	}

	if outType := methodType.Out(0); outType.Kind() != reflect.String {
		return tableName
	}

	params := make([]reflect.Value, 0)
	rts := methodValue.Call(params)
	if rts[0].String() != "" {
		return rts[0].String()
	}
	return tableName
}

func typeMapper(field reflect.StructField) string {
	kind := field.Type.Kind()
	tag := field.Tag.Get("sql")

	typeValue := ""
	defaultValue := ""
	sizeValue := "32"

	switch kind {
	case reflect.Int:
		typeValue = "int"
	case reflect.String:
		typeValue = "varchar"
		sizeValue = "32"
	default:
		typeValue = "text"
	}

	pkValue := false
	nullValue := false

	for _, tagValue := range strings.Split(tag, ";") {
		if strings.HasPrefix(tagValue, "type(") {
			typeValue = tagValue[5 : len(tagValue)-1]
		} else if strings.HasPrefix(tagValue, "size(") {
			sizeValue = tagValue[5 : len(tagValue)-1]
		} else if strings.HasPrefix(tagValue, "default(") {
			defaultValue = tagValue[8 : len(tagValue)-1]
		} else if tagValue == "pk" {
			pkValue = true

		} else if tagValue == "null" {
			nullValue = true
		}
	}

	var builder strings.Builder
	builder.WriteString(typeValue)
	if typeValue == "varchar" {
		builder.WriteString(fmt.Sprintf("(%s)", sizeValue))
	}

	if pkValue {
		builder.WriteString(" primary key")
	}
	if !nullValue {
		builder.WriteString(" not null")
	}
	if defaultValue != "" {
		builder.WriteString(" default")
		if typeValue == "varchar" {
			builder.WriteString(fmt.Sprintf(` "%s"`, defaultValue))
		} else {
			builder.WriteString(fmt.Sprintf(` %s`, defaultValue))
		}
	}
	// varchar(32) primary key not null default "";
	return builder.String()
}

func sqldump(models ...interface{}) string {
	var builder strings.Builder
	for _, model := range models {
		t := reflect.TypeOf(model)
		v := reflect.ValueOf(model)

		if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct {
			st := t.Elem()

			builder.WriteString(fmt.Sprintf("CREATE TABLE `%s`(\n", snake(tableName(v))))
			fieldNum := st.NumField()
			if fieldNum == 0 {
				log.Panicf("%s field num is zero", st.Name())
			}

			for i := 0; i < fieldNum; i++ {
				field := st.Field(i)
				builder.WriteString(fmt.Sprintf("\t`%s` %s", snake(field.Name), typeMapper(field)))
				if i != fieldNum-1 {
					builder.WriteString(",")
				}
				builder.WriteString("\n")
			}

			builder.WriteString(fmt.Sprintf(")ENGINE=INNODB DEFAULT CHARSET utf8mb4;\n"))

		} else {
			panic("error")
		}
	}

	return builder.String()
}

// create table user (id int, name varchar(32)) engine=innodb default charset utf8mb4;
func main() {
	fmt.Println(sqldump(new(User), new(Department)))

}
