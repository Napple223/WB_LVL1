package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.

//Способ с помощью time.After.
//Под капотом time.NewTimer.

func sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	start := time.Now()
	sleep(3 * time.Second)
	fmt.Println(time.Since(start))
}
