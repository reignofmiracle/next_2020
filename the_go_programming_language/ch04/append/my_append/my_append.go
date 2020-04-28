package my_append

import "fmt"

// AppendInt ...
func AppendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
		fmt.Printf("expand %v, %d\n", x, y)
	}
	z[len(x)] = y
	return z
}

// func AppendInt2(x []int, y ...int) []int {
// 	var z []int
// 	zlen := len(x) + len(y)
// 	copy(z[len(x):], y)
// 	return z
// }
