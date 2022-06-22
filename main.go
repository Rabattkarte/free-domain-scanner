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
	fmt.Printf("Using two of the following runes: %c\n", runes)

	tld := ".de"
	for _, d1 := range runes {
		for _, d2 := range runes {
			domain := fmt.Sprintf("%c%c%s", d1, d2, tld)

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
}
