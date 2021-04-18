package popcount

//import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

//PopCount returns pop count
func PopCount(x uint64) int {

	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

//PopCount2 returns pop count
/*func PopCount2(x uint64) int {
	/*for i := range pc {
		fmt.Printf("%v ", pc[i])
	}
	fmt.Printf("Value = %v", pc[byte(x>>(0*8))])
	var res int = 0
	for i :=  {
		res += pc[byte(x>>(i*8))]
	}
	return res
}
*/
