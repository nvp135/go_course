package main

import (
	"fmt"
	"math"

	"github.com/nvp135/go_course/03_packages/strutil"
	//"./strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("hello"))
}
