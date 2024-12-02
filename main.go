package main

import (
	"aoc-2024/d1"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		panic("requried args: {day} {part}\n")
	}

	d := os.Args[1]
	p := os.Args[2]

	switch q := d + "-" + p; q {
	case "1-1":
		d1.P1()
	case "1-2":
		d1.P2()
	}

}
