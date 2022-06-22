package main

import (
	"fmt"
	"sort"

	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func main() {
	// all lower-case ASCII alnum runes
	var runes []rune
	for r := 'a'; r <= 'z'; r++ {
		runes = append(runes, r)
	}
	for r := '0'; r <= '9'; r++ {
		runes = append(runes, r)
	}
	fmt.Printf("Using the following runes: %c\n", runes)

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
			domain := string(second) + string(permutation)
			name2c = append(name2c, domain)
		}
	}
	names = append(names, name2c...)

	// Permutation 3 letters
	var name3c []string
	for _, permutation := range runes {
		for _, second := range runes {
			for _, third := range runes {
				domain := string(third) + string(second) + string(permutation)
				name3c = append(name3c, domain)
			}
		}
	}
	names = append(names, name3c...)

	sort.Strings(names)
	// Whois
	fmt.Printf("Using the following domain names: %s\n", names)
	tld := ".de"
	for _, domain := range names {
		domain := string(domain) + tld

		fmt.Printf("Testing %s", domain)
		result, err := whois.Whois(domain)
		if err == nil {
			_, err := whoisparser.Parse(result)
			if err != nil {
				fmt.Println(" - FREE")
			} else {
				fmt.Println(" - not available")
			}
		}
	}
}
