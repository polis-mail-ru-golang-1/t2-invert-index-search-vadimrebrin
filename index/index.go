package index

import (
	"sync"
)

//BuildIndex builds invert index for defined files
//files - это map, где ключи - именя файлов, а значения - массив слов из этих файлов
//dict - это map, где ключи это слова, значения - это maps где ключи - названия файла
//из которого взято слово, а значение - колличество повторений этого слова в файле
func BuildIndex(dict map[string]map[string]int, files map[string][]string) {
	chanel := make(chan map[string]map[string]int)
	var wgindex sync.WaitGroup
	var wgput sync.WaitGroup

	for nameoffile, onefile := range files {
		wgindex.Add(1)
		go indexFile(nameoffile, onefile, &wgindex, chanel)
	}
	wgput.Add(1)
	go putToDict(chanel, dict, &wgput)

	wgindex.Wait()
	close(chanel)

	wgput.Wait()
}

func putToDict(chanel <-chan map[string]map[string]int, dict map[string]map[string]int, wgindex *sync.WaitGroup) {
	defer wgindex.Done()
	for data := range chanel {
		for word, value := range data {
			if len(dict[word]) == 0 {
				dict[word] = value
			} else {
				for key, val := range value {
					//если первый раз
					//если не первый раз
					dict[word][key] = val
				}
			}
		}
	}
}

//nameoffile - имя файла
//onefile - slice со всеми словами из этого файла
//dict - это map, где ключи это слова, значения - это maps где ключи - названия файла
func indexFile(nameoffile string, onefile []string, wgindex *sync.WaitGroup,
	chanel chan<- map[string]map[string]int) {
	dict := make(map[string]map[string]int)

	defer wgindex.Done()
	for _, word := range onefile {
		//если слово встретилось первый раз
		if dict[word] == nil {
			filemap := make(map[string]int)
			dict[word] = filemap
		}
		dict[word][nameoffile]++
	}
	chanel <- dict
}

//FindPhrase finds phrase in invert index
//phrase - это массив слов из фразы
func FindPhrase(dict map[string]map[string]int, phrase []string) map[string]int {
	samewords := make(map[string]map[string]int)
	res := make(map[string]int)

	for item := range dict {
		for _, word := range phrase {
			if item == word {
				samewords[word] = dict[word]
			}
		}
	}
	//Checks if all words from phrase are found
	for _, word := range phrase {
		isInDict := false
		for item := range samewords {
			if word == item {
				isInDict = true
				break
			}
		}
		if !isInDict {
			return res
		}
	}

	for _, item := range samewords {
		for name, i := range item {
			res[name] = res[name] + i
		}
	}
	return res
}
