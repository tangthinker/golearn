package main

var (
	int1       int
	float1     float32
	bool1      bool
	string1    string
	pointer    *int
	interface1 interface{}
)

func main() {
	println(int1, float1, bool1, string1, pointer, interface1)
	// 0 +0.000000e+000 false  0x0 (0x0,0x0)
	println(pointer == nil)
	// true
}
