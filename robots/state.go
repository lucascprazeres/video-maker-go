package robots

import (
	"encoding/json"
	"os"
)

type Sentence struct {
	Text              string   `json:"text"`
	Keywords          []string `json:"keywords"`
	Images            []string `json:"images"`
	GoogleSearchQuery string   `json:"googleSearchQuery"`
}

type State struct {
	MaximumSentences      int    `json:"maximumSentences"`
	Prefix                string `json:"prefix"`
	SearchTerm            string `json:"searchTerm"`
	SourceContentOriginal string `json:"sourceContentOriginal"`
	SourceContentSanized  string `json:"sourceContentSanized"`
	Sentences             []*Sentence
}

func NewState() *State {
	return &State{}
}

func NewSentence() *Sentence {
	return &Sentence{}
}

func SetState(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "	")

	if err != nil {
		panic(err)
	}

	file, err := os.Create("content.json")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)

	if err != nil {
		panic(err)
	}
}

func LoadState() *State {
	jsonData, err := os.ReadFile("content.json")

	if err != nil {
		panic(err)
	}

	data := NewState()
	err = json.Unmarshal(jsonData, &data)

	if err != nil {
		panic(err)
	}

	return data
}
