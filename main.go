package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"index/suffixarray"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"unicode"
)

func main() {
	searcher := Searcher{}
	err := searcher.Load("completeworks.txt")
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/search", handleSearch(searcher))
	http.HandleFunc("/recommendations", handleRecommendations(searcher))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	fmt.Printf("Listening on port %s...", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Searcher struct {
	CompleteWorks string
	SuffixArray   *suffixarray.Index
}

func handleSearch(searcher Searcher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query, ok := r.URL.Query()["q"]
		if !ok || len(query[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing search query in URL params"))
			return
		}
		results := searcher.Search(query[0])
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		err := enc.Encode(results)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("encoding failure"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf.Bytes())
	}
}

func (s *Searcher) Load(filename string) error {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Load: %w", err)
	}
	s.CompleteWorks = string(dat)
	s.SuffixArray = suffixarray.New(dat)
	return nil
}

func IsUpper(s string) bool {
    for _, r := range s {
        if !unicode.IsUpper(r) && unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

func (s *Searcher) Search(query string) []string {
	isQueryMayus := IsUpper(query)
	_rxp := `([A-Z](.+)[\s\.-]){4}`+query+`((.+)[\s\.-]){4}`
	if isQueryMayus {
		_rxp = query+`((.+)[\s\.-]){5}`
	}
	fmt.Println(_rxp)
	rxp, err := regexp.Compile(_rxp)
	fmt.Println(rxp)
	results := []string{}
	if err != nil {
		panic("something went wrong with our search")
	}
	idxs := s.SuffixArray.FindAllIndex(rxp, -1)
	fmt.Println(idxs)
	for _, idx := range idxs {
		results = append(results, s.CompleteWorks[idx[0]:idx[1]])
	}
	fmt.Println(results)
	return results
}
