package main

import (
	"bufio"
	"fmt"
	"os"
	"polytonicgreek/polytonicstring"
	"strconv"
)

func main() {
	str := ""
	width := 0
	fmt.Println("Width:")
	for width == 0 {
		fmt.Scanln(&str)
		width, _ = strconv.Atoi(str)
	}
	fmt.Println("Beta:")
	scanner := bufio.NewScanner(os.Stdin)
	str = ""
	for scanner.Scan() {
		str += scanner.Text() + "\n"
	}
	fmt.Println("\n\nGreek:")
	greek := polytonicstring.New(str)
	polytonicstring.WrapString(greek, width)
	fmt.Println(greek)
}
