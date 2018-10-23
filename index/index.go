package index

import (
	"sync"
)

//BuildIndex builds invert index for defined files
//files - это map, где ключи - именя файлов, а значения - массив слов из этих файлов
//dict - это map, где ключи это слова, значения - это maps где ключи - названия файла
//из которого взято слово, а значение - колличество повторений этого слова в файле
func BuildIndex(dict map[string]map[string]int, files map[string][]string) {
	var dictMutex = &sync.Mutex{}
	wg := &sync.WaitGroup{}
	for nameoffile, onefile := range files {
		go indexFile(nameoffile, onefile, dict, dictMutex, wg)
	}
	wg.Wait()
}

func indexFile(nameoffile string, onefile []string, dict map[string]map[string]int,
	dictMutex *sync.Mutex, wg *sync.WaitGroup) {
	dictMutex.Lock()
	defer dictMutex.Unlock()
	wg.Add(1)
	defer wg.Done()
	for _, word := range onefile {
		//если слово встретилось первый раз
		if dict[word] == nil {
			filemap := make(map[string]int)
			dict[word] = filemap
		}
		dict[word][nameoffile]++
	}
}

//FindPhrase finds phrase in invert index
//phrase - это массив слов из фразы
func FindPhrase(dict map[string]map[string]int, phrase []string) map[string]int {
	samewords := make(map[string]map[string]int)
	for item := range dict {
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
