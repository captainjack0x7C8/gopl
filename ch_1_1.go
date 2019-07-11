package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func p_1_1() {
	fmt.Println(os.Args[:])
}

func p_1_2() {
	for i, s := range os.Args[:] {
		fmt.Printf("index = %d, value = %s\n", i, s)
	}
}

func p_1_3() {
	before := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	after := time.Now()
	d1 := after.Sub(before)
	fmt.Printf("Time of first way is %s\n", d1.String())
	before = time.Now()
	s, step := "", ""
	for _, arg := range os.Args[1:] {
		s += step + arg
		step = " "
	}
	fmt.Println(s)
	after = time.Now()
	d2 := after.Sub(before)
	fmt.Printf("Time of second way is %s\n", d2.String())
	if d2 > d1 {
		fmt.Printf("%s is faster than %s", "first way", "second way")
	} else {
		fmt.Printf("%s is faster than %s", "second way", "first way")
	}
}

func main() {
	p_1_1()
	p_1_2()
	p_1_3()
}
