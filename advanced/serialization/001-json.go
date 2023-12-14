package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
JSON 与 Go 类型对应如下：
- bool 对应 JSON 的 boolean
- float64 对应 JSON 的 number
- string 对应 JSON 的 string
- nil 对应 JSON 的 null

不是所有的数据都可以编码为 JSON 类型：只有验证通过的数据结构才能被编码：
1. JSON 对象只支持字符串类型的 key；要编码一个 Go map 类型，map 必须是 map[string]T（T是 json 包中支持的任何类型）
2. Channel，复杂类型和函数类型不能被编码
3. 不支持循环数据结构；它将引起序列化进入一个无限循环
4. 指针可以被编码，实际上是对指针指向的值进行编码（或者指针是 nil）
*/

type PersonJ struct {
	id   uint64 // 未导出字段，不会写入json
	Name string
	Age  int
}

func main() {
	shanliao := PersonJ{
		001,
		"shanliao",
		22,
	}
	king := PersonJ{
		002,
		"king",
		22,
	}
	ret, err := json.Marshal(shanliao)
	if err == nil {
		fmt.Println(string(ret))
	}

	file, _ := os.OpenFile("json.dat", os.O_CREATE|os.O_WRONLY, 0666)
	encoder := json.NewEncoder(file)
	_ = encoder.Encode(shanliao)
	_ = encoder.Encode(king)

	file, _ = os.Open("json.dat")
	decoder := json.NewDecoder(file)
	for decoder.More() {
		var receive PersonJ
		_ = decoder.Decode(&receive)
		fmt.Println(receive)
	}
	// {0 shanliao 22}
	// {0 king 22}

	// 可使用以下方式访问任意类型（或未知类型）的数据
	var anyType interface{}
	_ = json.Unmarshal(ret, &anyType)
	m := anyType.(map[string]interface{})
	for key, value := range m {
		fmt.Println(key, "->", value)
	}
}
