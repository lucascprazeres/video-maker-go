package robots

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	fmt.Println("Type a wikipedia search term:")

	input, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(input)
}

func askAndReturnPrefix() string {
	prefixes := []string{"Who is", "What is", "The history of"}

	for i, option := range prefixes {
		fmt.Printf("[%v] - %v\n", i+1, option)
	}

	var option int
	fmt.Println("Selecione uma opcao")
	fmt.Scanln(&option)
	return prefixes[option-1]
}
