package first

import (
	"fmt"
	"pkg/second"
)

const variable = 20

func main() {
	var z = second.CreateInt()
	fmt.Printf("%g %g", variable, z)
}
