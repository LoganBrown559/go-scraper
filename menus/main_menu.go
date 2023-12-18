// This package is for printing menus found in my program
package menus

import "fmt"

func MainMenu() {
	fmt.Println("\n\nHello, and welcome to Go Scraper!\n")
	fmt.Println("Please choose what you'd like from the menu:")

	fmt.Println("\t1. Quick Scrape...")
	fmt.Println("\t2. Filtered Scrape...")
	fmt.Println("\t3. Custom  Scrape...")
	fmt.Println("\t4. Quit...\n\n")
}
