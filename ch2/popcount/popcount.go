package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountOld возвращает степень заполнения значения X.
// Benchmark test:	0.2555 ns/op
func PopCountOld(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// Task 2.3 (Замена на цикл for)
// Benchmark test:	3.852 ns/op
func PopCountNew(x uint64) int {
	var result int
	for i := 0; i < 7; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}
