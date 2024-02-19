package robots

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lucascprazeres/video-maker/logging"
)

func Input() {
	content := map[string]interface{}{
		"maximumSentences": 7,
	}

	content["searchTerm"] = askAndReturnSearchTerm()
	content["prefix"] = askAndReturnPrefix()

	SetState(content)
}

func askAndReturnSearchTerm() string {
	reader := bufio.NewReader(os.Stdin)

	logging.Prompt("> Type a wikipedia search term: ")

	input, _ := reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func askAndReturnPrefix() string {
	prefixes := []string{"Who is", "What is", "The history of"}

	logging.Prompt("\nSelect a search prefix:\n")

	for i, option := range prefixes {
		logging.Prompt("[%v] - %v\n", i+1, option)
	}

	logging.Prompt("\n> ")

	var option int
	fmt.Scanln(&option)

	if option > len(prefixes) {
		logging.Error("> Selected option is outside range!\n")
		return ""
	}

	return prefixes[option-1]
}
