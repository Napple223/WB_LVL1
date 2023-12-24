package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.

//Вариант с использованием Timer.

func sleep(d time.Duration) {
	t := time.NewTimer(d)
	<-t.C
	t.Stop()
}

func main() {
	start := time.Now()
	sleep(3 * time.Second)
	fmt.Println(time.Since(start))
}
