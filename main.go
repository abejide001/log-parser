package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var domains []string
	var sum map[string]int
	sum = make(map[string]int)
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		fields := strings.Fields(in.Text())
		fmt.Printf("domain: %s - visits: %s\n", fields[0], fields[1])
		domain := fields[0]
		visits, _ := strconv.Atoi(fields[1])
		sum[domain] += visits
		_, ok := sum[domain]
		if !ok {
			domains = append(domains, domain)
		}
	}
	fmt.Printf("%-30s, %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))
	for domain, visits := range sum {
		fmt.Printf("%-30s, %10d\n", domain, visits)
	}
}
