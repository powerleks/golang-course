package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readSentences() []string {
	fmt.Println("Введите количество предложений: ")
	var amount int
	fmt.Scan(&amount)

	var sentences = make([]string, amount);
	in := bufio.NewReader(os.Stdin)


	for i := 0; i < amount; i++ {
		fmt.Printf("Введите предложение %v:\n", i + 1)
		var sentence string
		sentence, _ = in.ReadString('\n')
		words := strings.Fields(sentence)
		sentences[i] = words[len(words) - 1]
	}
	fmt.Println(sentences)

	return sentences
}

func readChars() []rune {
	in := bufio.NewReader(os.Stdin)
	fmt.Println("Введите символы через пробел: ")
	var sentence string
	sentence, _ = in.ReadString('\n')

	splits := strings.Split(sentence, " ")
	fmt.Println(splits)
	var chars = make([]rune, len(splits));

	for i, c := range(splits) {
		chars[i] = []rune(c)[0]
	}

	return chars
}

func parseTest(sentences []string, chars []rune) ([][]int, error) {
	// if len(sentences) == 0 || len(chars) == 0 {
	// 	empty := [][]int{{},{}}
	// 	return empty, fmt.Errorf("Отсутствуют предложения или символы")
	// }
	positions := make([][]int, len(sentences))
	for i := 0; i < len(sentences); i++ {
		positions[i] = make([]int, len(chars))
	}
	for i := 0; i < len(sentences); i++ {
		for j := 0; j < len(chars); j++ {
			positions[i][j] = strings.IndexRune(sentences[i], chars[j])
		}
	} 
	return positions, nil
}

func main() {
	sentences := readSentences()
	chars := readChars()

	positions, err := parseTest(sentences, chars)

	if err != nil {	
		fmt.Println("Позиции вхождения:", positions)
	}
}