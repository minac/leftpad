package main

import (
	"fmt"

	"github.com/minac/leftpad"
)

func main() {
	fmt.Println(leftpad.Format("hello", 15))
	fmt.Println(leftpad.Format("goodbye", 15))
	fmt.Println(leftpad.Format("Como estás?", 15))
	fmt.Println(leftpad.Format("Internationalization", 15))
	fmt.Println(leftpad.FormatRune("hello", 15, '😀'))
	fmt.Println(leftpad.FormatRune("goodbye", 15, '😀'))
	fmt.Println(leftpad.FormatRune("Como estás?", 15, '😀'))
}
