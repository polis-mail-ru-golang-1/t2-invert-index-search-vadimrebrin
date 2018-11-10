package index

import (
	"sync"
)

type Index map[string]map[string]int
type Result struct {
	File  string
	Count int
}

//BuildIndex builds invert index for defined files
//files - это map, где ключи - именя файлов, а значения - массив слов из этих файлов
//dict - это map, где ключи это слова, значения - это maps где ключи - названия файла
//из которого взято слово, а значение - колличество повторений этого слова в файле
func BuildIndex(dict Index, files map[string][]string) {
	fileCounter := len(files)
	chanel := make(chan Index)
	var wgindex sync.WaitGroup

	for nameoffile, onefile := range files {
		wgindex.Add(1)
		go indexFile(nameoffile, onefile, &wgindex, chanel)
	}
	putToDict(chanel, dict, fileCounter)
	wgindex.Wait()
	close(chanel)
}

func putToDict(chanel <-chan Index, dict Index, fileCounter int) {
	for i := 0; i < fileCounter; i++ {
		data := <-chanel
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
	chanel chan<- Index) {
	dict := make(Index)

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
func FindPhrase(dict Index, phrase []string) []Result {
	samewords := make(Index)
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
			return printInfo(res)
		}
	}

	for _, item := range samewords {
		for name, i := range item {
			res[name] = res[name] + i
		}
	}
	return printInfo(res)
}

//printInfo prints statistics of search
func printInfo(dict map[string]int) []Result {
	if len(dict) == 0 {
		return nil
	}
	var result []Result
	for name, count := range dict {
		var tmp Result
		tmp.File = name
		tmp.Count = count
		result = append(result, tmp)
	}
	for i := 0; i < len(result); i++ {
		for j := i; j < len(result); j++ {
			if result[i].Count < result[j].Count {
				tempcount := result[i]
				result[i] = result[j]
				result[j] = tempcount
			}
		}
	}

	return result
}
