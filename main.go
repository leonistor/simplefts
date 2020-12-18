package main

import (
	"flag"
	"log"
	"time"
)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "./data/enwiki-latest-abstract1.xml", "wiki abstract dump path")
	flag.StringVar(&query, "q", "cat", "search query")
	flag.Parse()

	log.Println("Starting simplefts")

	start := time.Now()
	docs, err := loadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	// matchedIDs := search(docs, query)
	// log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))

	// for _, doc := range matchedIDs {
	// log.Printf("%d\t%s\n", doc.ID, doc.Text)
	// }
}
