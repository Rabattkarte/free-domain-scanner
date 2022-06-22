package main

import (
	"fmt"

	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func main() {
	// TODO: Brute Force Algo, 1-3 chars. Permutation
	// TODO: buffer results
	// TODO: Threading
	// TODO: Threading: print no of domains every second
	// TODO: Cobra for own charset

	// all lower-case ASCII alnum runes
	var runes []rune
	for r := 'a'; r <= 'z'; r++ {
		runes = append(runes, r)
	}
	for r := '0'; r <= '9'; r++ {
		runes = append(runes, r)
	}

	fmt.Printf("Using the following runes: %c", runes)

	for _, domain := range runes {
		domain := string(domain) + ".de"
		result, err := whois.Whois(domain)
		if err == nil {
			result, err := whoisparser.Parse(result)
			if err == nil {
				// Print the domain status
				fmt.Println(result.Domain.Domain+": ", result.Domain.Status)
			}
		}
	}
}
