1. 复习
    if conditon {

    } else if condition_1 {

    }
    ...
    {

    } else if condition_n {

    } else {

    }

    switch var {
        case value:
        case value1, value2, value3:
        ...
        case valuen:
        default:
    }
    switch {
        case condition:
        case condition_1:
        ...
        case condition_n:
        default:
    }


    for {

    }
    for condition {

    }

    for init; cond; end {

    }

    switch
2. 作业
    => 助教
3. 新内容
    var t T
    var pointer *T = &t

    *pointer
    *pointer = V

    序列： 由相同类型元素组成的一个有序的集
     0123456
    "abcdefg"

    数组: 由相同类型（任意类型）元素组成的一个固定长度的有序集
    声明: [length]T
         [3]int
         [5]int

    切片: 由相同类型(任意类型)元素组成的长度可变的有序集
    声明: []T
    赋值:
        字面量
        make
        数组切片操作
        切片的切片操作
    操作:
        不能进行==和!= 运算
        函数
            len,cap,append,copy
        元素访问和修改
            slice[i]
            slice[i] = value
        遍历
            for
            for range
        切片操作:
            slice[start:end:cap_end]
            start<=end<=cap_en<=cap(slice)

            array[start:end:cap_end]
            start<=end<=cap_end<=len(array)

            start=0 slice[:end:cap_end]/array[:end:cap_end]
            end = len(slice)/len(array)
                slice[start::cap_end]/array[start::cap_end]
    // 练习
    // 队列: 先进先出
        queue := []int{}
        1 2 3=> 1 2 3
        进切片(添加到后面)
            append(1)
            append(2)
            append(3)
        出切片
            获取索引为0的元素，删除索引为0的元素

            queue = queue[1:]
    // 堆栈: 先进后出
        stack := []int{}
        1, 2, 3 => 3, 2, 1
        进切片
            append(1)
            append(2)
            append(3)
        出切片
            获取索引为len()-1 删除len-1的元素
            stack = stack[:len-1]

        []T

        // T => [3]int
        [][3]int
        // T => []int
        [][]int

        映射：
            含义: 数据 => ID
                定义key=>value对的集(可能是有序/无序的)
                go中无序(放入key的顺序与再内存中的顺序无关)
                        与遍历无关
                映射实现方式：hashtable(go的实现方式), treemap
                key对value进行增，删，改，获取
                key唯一的
                key类型的要求：== !=
                value类型任务
            定义: map[keyType]valueType
            赋值:
                make:
                    make(map[keyType]valueType)
                字面量
                    map[keyType]valueType{} //空map
                    map[keyType]valueType{k:v, k:v, k:v}
                零值: nil

            操作：
                元素数量: len
                元素操作:
                    获取: mapName[key]
                        key存在
                            mapName[key] => 对应value值
                        key不存在
                            mapName[key] => 不报错, valueType类型的0值
                            如何判断key是否存在
                            v, ok := mapName[key]
                            v := mapName[key]
                    增:
                        mapName[key] = value
                        key不存在
                    删:
                        delete(mapName, key)
                    改:
                        mapName[key] = value
                        key存在