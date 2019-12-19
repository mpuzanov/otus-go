package topwords

import (
	"sort"
	"strings"
	"unicode"
)

type pair struct {
	key   string
	value int
}

// TopWords10 возвращает 10 самых частовстречающихся в тексте слов
func TopWords10(in string) []string {

	var out []string
	const topCountWord = 10

	words := make(map[string]int)

	// подсчитываем слова
	countWords(in, words)

	pairs := make([]pair, 0, len(words))

	// заносим словарь в структуру для сортировки
	for k, v := range words {
		pairs = append(pairs, pair{k, v})
	}

	// сортируем по убыванию
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value > pairs[j].value || (pairs[i].value == pairs[j].value && pairs[i].key < pairs[j].key)
	})

	// выводим первые topCountWord
	for i := 0; i < len(pairs) && i < topCountWord; i++ {
		out = append(out, pairs[i].key)
	}
	return out
}

func countWords(in string, words map[string]int) {

	f := func(c rune) bool {
		return unicode.IsSpace(c) || unicode.IsPunct(c)
	}
	sa := strings.FieldsFunc(in, f)
	for _, word := range sa {
		words[strings.ToLower(word)]++
	}
}
