package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/mochi8k/domain-name-generator/thesaurus"
)

func main() {

	apiKey := os.Getenv("BHT_APIKEY")
	if apiKey == "" {
		log.Fatalln("please export BHT_APIKEY={your api key}")
	}

	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		word := s.Text()

		if word == "" {
			continue
		}

		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalf("bighuge: %qの類語検索に失敗しました: %v", word, err)
		}

		if len(syns) == 0 {
			log.Fatalf("%qに類語はありませんでした\n", word)
		}

		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
