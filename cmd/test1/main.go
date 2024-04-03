package main

import (
	"fmt"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/search"
)

func main() {
	messages := []struct {
		Id   string
		From string
		Body string
	}{
		{
			Id:   "index-example1",
			From: "forfd8960@github.com",
			Body: "test with bleve indexing",
		},
		{
			Id:   "index-example2",
			From: "forfd8960@github.com",
			Body: "another document",
		},
	}

	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("example.bleve", mapping)
	if err != nil {
		panic(err)
	}

	for _, msg := range messages {
		if err := index.Index(msg.Id, msg); err != nil {
			fmt.Printf("index err: %v\n", err)
			return
		}
	}

	fmt.Printf("index: %v has been successful build\n", index)

	// index, _ = bleve.Open("example.bleve")
	query := bleve.NewQueryStringQuery("another")
	fmt.Printf("query: %v\n", *query)

	searchRequest := bleve.NewSearchRequest(query)
	fmt.Printf("search request: %v\n", *searchRequest)

	searchResult, _ := index.Search(searchRequest)
	fmt.Printf("search result: %v\n", *searchResult)
	fmt.Printf("search resul hits: %v\n", searchResult.Hits)
	fmt.Printf("search result took: %v\n", searchResult.Took)
	fmt.Printf("search result status: %v\n", searchResult.Status)

	for _, hit := range searchResult.Hits {
		printResult(hit)

		query := bleve.NewDocIDQuery([]string{hit.ID})
		fmt.Printf("docID query: %v\n", *query)
		sReq := bleve.NewSearchRequest(query)
		sRes, _ := index.Search(sReq)
		fmt.Printf("docID query result: %v\n", *sRes)

		for _, ht := range sRes.Hits {
			printResult(ht)
		}
	}
}

func printResult(hit *search.DocumentMatch) {
	fmt.Printf("search hit: %v\n", *hit)
	fmt.Printf("search hit fields: %v\n", hit.Fields)
	fmt.Printf("search hit inedx: %v\n", hit.Index)
	fmt.Printf("search hit ID: %v\n", hit.ID)
	fmt.Printf("search hit internal ID: %s\n", string(hit.IndexInternalID))
	fmt.Printf("search hit Expl: %v\n", hit.Expl)
}
