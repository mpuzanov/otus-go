package topwords

import (
	"sort"
	"strings"
)

type pair struct {
	key   string
	value int
}

// TopWords10 возвращает 10 самых частовстречающихся в тексте слов
func TopWords10(in string) (out []string) {
	
	const topCountWord = 10
	var pairs []pair
	
	words := make(map[string]int)

	// подсчитываем слова
	countWords(in, words)

	// заносим словарь в структуру для сортировки
	for k, v := range words {
		pairs = append(pairs, pair{k, v})
	}

	// сортируем по убыванию
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value > pairs[j].value
	})

	// выводим первые topCountWord
	for i := 0; i < len(pairs) && i < topCountWord; i++ {
		out = append(out, pairs[i].key)
	}
	return
}

func countWords(in string, words map[string]int) {
	sa := strings.Split(in, " ")
	for _, word := range sa {
		if word != "-" {
			words[strings.ToLower(word)]++
		}
	}
}
