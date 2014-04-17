package main

import (
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		go A()
		if i%2 == 0 {
			go B()
		}
		if i%3 == 0 {
			go C()
		}
		if i%4 == 0 {
			go D()
		}
		if i%5 == 0 {
			go Sleep()
		}
	}
	time.Sleep(100 * time.Millisecond)
	panic("oh no")
}

func A() {
	B()
}

func B() {
	C()
}

func C() {
	D()
}

func D() {
	Sleep()
}

func Sleep() {
	time.Sleep(time.Second)
}
