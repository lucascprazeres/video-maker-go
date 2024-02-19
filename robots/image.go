package robots

import (
	"math/rand"

	"github.com/lucascprazeres/video-maker/logging"
)

func Image() {
	logging.Info("[image-robot] Starting...\n")
	state := LoadState()

	fetchImagesOfAllSentences(state)

	SetState(state)
}

func fetchImagesOfAllSentences(state *State) {
	logging.Info("[image-robot] Fetching images from Google Images\n")

	for i := 0; i < len(state.Sentences); i++ {
		var query string

		if i == 0 {
			query = state.SearchTerm
		} else {
			randomIndex := rand.Intn(len(state.Sentences[i].Keywords))
			query = state.SearchTerm + " " + state.Sentences[i].Keywords[randomIndex]
		}

		state.Sentences[i].GoogleSearchQuery = query
	}
}
