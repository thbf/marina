package main

import (
	"fmt"
	"links"
	"log"
	"os"
)

func main() {
	breadthFirst(crawl, os.Args[1:])
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	//TODO: find out if url is job posting
	//TODO: if url is job posting, extract info
	// interested in: company, location, salary, programming language
	list, err := links.Extract(url)

	if err != nil {
		log.Print(err)
	}
	return list
}
