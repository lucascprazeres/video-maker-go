package robots

import (
	"regexp"
	"strings"

	"github.com/neurosnap/sentences/english"
	gowiki "github.com/trietmn/go-wiki"
)

func Text() {
	state := LoadState()

	fetchContentFromWikipedia(state)
	sanitizeContent(state)
	breakContentIntoSentences(state)
	limitMaximumSentences(state)
	fetchKeywordsOffAllSentences(state)

	SetState(state)
}

func fetchContentFromWikipedia(state *State) {
	search := (state.Prefix) + " " + state.SearchTerm

	page, err := gowiki.GetPage(search, -1, false, true)

	if err != nil {
		panic(err)
	}

	content, err := page.GetSummary()
	if err != nil {
		panic(err)
	}

	state.SourceContentOriginal = content
}

func sanitizeContent(state *State) {
	withoutDatesInParenthesis := removeDatesInParenthesis(state.SourceContentOriginal)
	state.SourceContentSanized = withoutDatesInParenthesis
}

func removeDatesInParenthesis(text string) string {
	regex := regexp.MustCompile(`\((?:\([^()]*\)|[^()])*\)`)
	result := regex.ReplaceAllString(text, "")
	return result
}

func breakContentIntoSentences(state *State) {
	tokenizer, err := english.NewSentenceTokenizer(nil)
	if err != nil {
		panic(err)
	}

	sentences := tokenizer.Tokenize(state.SourceContentSanized)

	sentencesContent := make([]*Sentence, len(sentences))

	for i, s := range sentences {
		sentence := NewSentence()
		sentence.Text = strings.TrimSpace(s.Text)
		sentence.Keywords = make([]string, 0)
		sentence.Images = make([]string, 0)
		sentencesContent[i] = sentence
	}

	state.Sentences = sentencesContent
}

func limitMaximumSentences(state *State) {
	if len(state.Sentences) > state.MaximumSentences {
		state.Sentences = state.Sentences[:state.MaximumSentences]
	}
}

func fetchKeywordsOffAllSentences(state *State) {
	tokenizer, err := english.NewSentenceTokenizer(nil)
	if err != nil {
		panic(err)
	}

	for sentenceIndex, s := range state.Sentences {
		result := tokenizer.WordTokenizer.Tokenize(s.Text, false)
		var keywords []string

		for _, ky := range result {
			// TODO: check for a better way to filter english common words
			if len(ky.Tok) <= 3 {
				continue
			}

			keywords = append(keywords, ky.Tok)
		}
		state.Sentences[sentenceIndex].Keywords = removeDuplicateKeywords(keywords)
	}

}

func removeDuplicateKeywords(keywords []string) []string {
	uniqueKeywords := make(map[string]bool)
	result := []string{}

	for _, ky := range keywords {
		if !uniqueKeywords[ky] {
			uniqueKeywords[ky] = true
			result = append(result, ky)
		}
	}

	return result
}
