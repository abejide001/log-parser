package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	ps := newParser()
	in := bufio.NewScanner(os.Stdin) // get the input

	for in.Scan() { // loop through to get the text
		parsed := parse(in.Text())
		update(&ps, parsed)
	}

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS") // headers
	fmt.Println(strings.Repeat("-", 45))           // print number of lines

	for _, domain := range ps.domains {
		parsed := ps.sum[domain]
		ps.total += parsed.visits
		fmt.Printf("%-30s %10d\n", domain, parsed.visits) // print  the domain and visits
	}
	fmt.Println(strings.Repeat("-", 45)) // print lines

	fmt.Printf("%-30s %10d\n", "TOTAL", ps.total) // print the total
}
