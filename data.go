package main

import (
	_ "embed"
	"encoding/csv"
	"io"
	"strconv"
	"strings"
)

//go:embed data/articles.csv
var csvArticles string

func readArticles() ([]Article, error) {
	r := csv.NewReader(strings.NewReader(csvArticles))
	results := make([]Article, 0)
	// skip the title row
	_, err := r.Read()
	if err == io.EOF {
		return results, nil
	}
	if err != nil {
		return nil, err
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		id, err := strconv.ParseInt(record[0], 10, 32)
		if err != nil {
			return nil, err
		}
		authorID, err := strconv.ParseInt(record[2], 10, 32)
		if err != nil {
			return nil, err
		}
		results = append(results, Article{
			ID:       int(id),
			Title:    record[1],
			AuthorID: int(authorID),
		})
	}

	return results, nil
}
