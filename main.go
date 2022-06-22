package main

import (
	"fmt"

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

	tld := ".de"
	for _, domain := range runes {
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
