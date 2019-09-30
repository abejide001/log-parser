package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	total := 0

	var (
		domains []string // to store unique domain names
		sum     map[string]int
	)
	sum = make(map[string]int) // create a map to store the domain name

	in := bufio.NewScanner(os.Stdin) // get the input

	for in.Scan() { // loop through to get the text
		fields := strings.Fields(in.Text()) // split the line
		if len(fields) != 2 {
			log.Panic("fields must be two")
			return
		}

		domain := fields[0]                    // the domain name
		visits, err := strconv.Atoi(fields[1]) // the number of visits
		if visits < 0 || err != nil {
			log.Panic("wrong input")
			return
		}
		_, ok := sum[domain] // check if there is a domain

		if !ok {
			domains = append(domains, domain) // add domain to the slice
		}
		sum[domain] += visits
	}

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS") // headers
	fmt.Println(strings.Repeat("-", 45))           // print number of lines

	for _, domain := range domains {
		visits := sum[domain]
		total += visits
		fmt.Printf("%-30s %10d\n", domain, visits)
	}
	fmt.Println(strings.Repeat("-", 45))

	fmt.Printf("%-30s %10d\n", "TOTAL", total)
}
