/*
   Race condition problem https://medium.com/german-gorelkin/race-8936927dba20
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go func() {
		fmt.Printf("A->")
	}()

	go func() {
		fmt.Printf("B")
	}()

	time.Sleep(time.Second)

	//wrongForFunc() // rightForFunc fixs it
	rightForFunc()
}

func wrongForFunc() {
	x := 0
	for {
		go func() {
			x++
		}()

		go func() {
			if x%2 == 0 {
				time.Sleep(1 * time.Millisecond)
				fmt.Println(x)
			}
		}()
	}

}

func rightForFunc() {
	var mu sync.Mutex
	x := 0
	for {
		go func() {
			mu.Lock()
			x++
			mu.Unlock()
		}()
		// 1st
		go func() {
			mu.Lock()
			if x%2 == 0 {
				time.Sleep(1 * time.Millisecond)
				fmt.Println(x)
			}
			mu.Unlock()
		}()
		//2nd
		go func() {
			y := x
			if y%2 == 0 {
				time.Sleep(1 * time.Millisecond)
				fmt.Println(x)
			}
		}()
	}

}

/*
        func f(from string) {
            for i := 0; i < 3; i++ {
                fmt.Println(from, ":", i)
    }
}

func main() {

    f("direct")

    go f("goroutine")

    go func(msg string) {
        fmt.Println(msg)
    }("going")

    time.Sleep(time.Second)
    fmt.Println("done")
}*/
