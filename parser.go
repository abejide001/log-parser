package main

import (
	"log"
	"strconv"
	"strings"
)

type result struct {
	visits int
	domain string
}
type parser struct {
	sum     map[string]result
	domains []string // to store unique domain names
	total   int
}

func newParser() parser {
	return parser{sum: make(map[string]result)}

}

func parse(line string) result {
	var parsed result
	var err error
	fields := strings.Fields(line) // split the line
	if len(fields) != 2 {
		log.Panic("fields must be two")
		return parsed
	}

	parsed.domain = fields[0]                    // the domain name
	parsed.visits, err = strconv.Atoi(fields[1]) // the number of visits
	if parsed.visits < 0 || err != nil {
		log.Panic("wrong input")
		return parsed
	}
	return parsed
}

func update(ps *parser, parsed result) {
	domain, visits := parsed.domain, parsed.visits
	_, ok := ps.sum[domain] // check if there is a domain
	if !ok {
		ps.domains = append(ps.domains, domain) // add domain to the slice
	}
	ps.sum[domain] = result{
		domain: domain,
		visits: visits + ps.sum[domain].visits,
	}
}
