package main

import (
	"fmt"
	"sort"

	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

const DEBUG bool = false

func main() {
	// all lower-case ASCII alnum runes
	var runes []rune
	for r := 'a'; r <= 'z'; r++ {
		runes = append(runes, r)
	}
	for r := '0'; r <= '9'; r++ {
		runes = append(runes, r)
	}
	if DEBUG {
		fmt.Printf("Using the following runes: %c\n", runes)
	}

	// Create domain names to check
	// Append all single runes to names
	var names []string
	for _, v := range runes {
		names = append(names, string(v))
	}

	// Permutation 2 letters
	var name2c []string
	for _, permutation := range runes {
		for _, second := range runes {
			d2 := string(second) + string(permutation)
			name2c = append(name2c, d2)
		}
	}
	names = append(names, name2c...)

	// Permutation 3 letters
	var name3c []string
	for _, permutation := range runes {
		for _, second := range runes {
			for _, third := range runes {
				d3 := string(third) + string(second) + string(permutation)
				name3c = append(name3c, d3)
			}
		}
	}
	names = append(names, name3c...)

	sort.Strings(names)

	// Whois
	if DEBUG {
		fmt.Printf("Using the following domain names: %s\n", names)
	}
	tld := ".de"
	fmt.Printf("Scanning %d domains", len(names))
	var free []string
	for _, domain := range names {
		fmt.Printf(".")
		domain := string(domain) + tld

		if DEBUG {
			fmt.Printf("Testing\t%s", domain)
		}
		result, err := whois.Whois(domain)
		if err == nil {
			_, err := whoisparser.Parse(result)
			if err != nil {
				if DEBUG {
					fmt.Println("\t-> FREE")
				} else {
					free = append(free, domain)
				}
			} else {
				if DEBUG {
					fmt.Println("\t-> not available")
				}
			}
		}
	}

	fmt.Printf("\nFree domains: %s", free)
}
