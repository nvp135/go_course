package main

import "fmt"

func main() {
	peoplesNames1 := [3]string{"Firs", "Second", "Third"}
	peoplesNames2 := [2]byte{1, 2}

	fmt.Println(len(peoplesNames1))
	fmt.Println(peoplesNames2)

	slc := make([]string, 3)
	slc[0] = "[0]"
	slc[1] = "[1]"
	slc[2] = "[2]"

	slc = append(slc, "[3]")
	slc = append(slc, "[4]", "[5]")
	fmt.Println("slc ", slc)

	slc2 := make([]string, len(slc))
	copy(slc2, slc)
	fmt.Println("slc2 ", slc2)

	slc3 := slc[2:5]
	fmt.Println("slc3 ", slc3)

	slc3 = slc[:5]
	fmt.Println("slc3 [:5]", slc3)

	slc3 = slc[2:]
	fmt.Println("slc3 [2:]", slc3)
	// fmt.Println(slc2)

	twoD := make([][]int, 5)
	for i := 0; i < 5; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

}
