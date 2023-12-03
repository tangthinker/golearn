package main

/*
	常量类型只可以为：布尔型、数字型（整数型、浮点型和复数）和字符串型
					bool   number(int\float\complex) string
*/
func main() {
	// go语言的类型推导是在编译阶段完成
	// 不同于Python或Ruby在运行时推断
	const Pi float32 = 3.1415926
	const (
		a, b, c = "a", "b", "c"
		e, f, g = "e", "f", "g"
	)
	println(a, b, c)

	// 作为枚举值
	const (
		ENABLE  = 0
		DISABLE = 1
	)
	println(ENABLE, DISABLE)

	// iota整形自动递增
	const (
		RED = iota
		GREEN
		_
		WHITE
	)
	println(RED, GREEN, WHITE)

	type CustomType int
	const (
		custom1 CustomType = iota
		custom2
		custom3
	)
	println(custom1, custom2, custom3)

	// funny usage
	const (
		_  = iota             // 使用 _ 忽略不需要的 iota
		KB = 1 << (10 * iota) // 1 << (10*1)
		MB                    // 1 << (10*2)
		GB                    // 1 << (10*3)
		TB                    // 1 << (10*4)
		PB                    // 1 << (10*5)
		EB                    // 1 << (10*6)
		//ZB                    // 1 << (10*7)	越界
		//YB                    // 1 << (10*8) 越界
	)

	println(PB, EB)

}
