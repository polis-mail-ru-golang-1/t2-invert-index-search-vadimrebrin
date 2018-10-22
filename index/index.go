package index

func BuildIndex(dict map[string]map[string]int, files map[string][]string) {
	//files - это map, где ключи - именя файлов, а значения - массив слов из этих файлов
	//dict - это map, где ключи это слова, значения - это maps где ключи - названия файла
	//из которого взято слово, а значение - колличество повторений этого слова в файле
	for nameoffile, onefile := range files {
		for _, word := range onefile {
			//если слово встретилось первый раз
			if dict[word] == nil {
				filemap := make(map[string]int)
				dict[word] = filemap
			}
			dict[word][nameoffile]++
		}
	}
}

func FindPhrase(dict map[string]map[string]int, phrase []string) map[string]int {
	//phrase - это массив слов из фразы
	samewords := make(map[string]map[string]int)
	for item, _ := range dict {
		for _, word := range phrase {
			if item == word {
				samewords[word] = dict[word]
			}
		}
	}
	res := make(map[string]int)
	for _, item := range samewords {
		for name, i := range item {
			res[name] = res[name] + i
		}
	}
	return res
}
