//author ljf
//date:2021-03-24
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	//读取dream.log
	content, err := ioutil.ReadFile("dream.log")
	//判断文件是否为空
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Printf("File contents: %s", content)
	num := make(map[string]int, 0)
	//通过ASCII码表判断并统计
	//大写字母 A~Z 65-90，小写字母a~z 97-122
	for _, v := range content {
		if v >= 65 && v <= 90 || v >= 97 && v <= 122 {
			//将符合条件的ASCII码转换成string并统计
			num[string(v)]++
		}
	}
	fmt.Println("打印结果如下：")
	for k, b := range num {
		fmt.Printf("字母:%s\t 数量:%d\n", k, b)

	}

}

/*打印结果如下：
字母:t   数量:649
字母:n   数量:449
字母:o   数量:587
字母:s   数量:407
字母:F   数量:4
字母:c   数量:172
字母:z   数量:6
字母:b   数量:105
字母:k   数量:51
字母:R   数量:2
字母:M   数量:8
字母:u   数量:176
字母:K   数量:1
字母:J   数量:2
字母:T   数量:10
字母:O   数量:3
字母:q   数量:7
字母:Y   数量:3
字母:I   数量:24
字母:B   数量:4
字母:N   数量:24
字母:y   数量:122
字母:L   数量:16
字母:h   数量:377
字母:g   数量:159
字母:d   数量:260
字母:A   数量:27
字母:x   数量:5
字母:H   数量:3
字母:i   数量:528
字母:p   数量:90
字母:w   数量:147
字母:f   数量:213
字母:E   数量:1
字母:P   数量:4
字母:W   数量:15
字母:a   数量:527
字母:G   数量:9
字母:D   数量:2
字母:e   数量:870
字母:r   数量:402
字母:m   数量:179
字母:j   数量:20
字母:l   数量:318
字母:C   数量:6
字母:S   数量:4
字母:v   数量:81
*/
