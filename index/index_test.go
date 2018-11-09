package index

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestBuildIndex(t *testing.T) {
	file, _ := os.Create("file1.txt")
	defer os.Remove("file1.txt")
	file.WriteString("lol kek lol golang")
	file2, _ := os.Create("file2.txt")
	defer os.Remove("file2.txt")
	file2.WriteString("lol testing degree vadim")
	file.Close()
	file2.Close()

	names := []string{"file1.txt", "file2.txt"}
	expected := map[string]map[string]int{
		"testing": {
			"file2.txt": 1,
		},
		"degree": {
			"file2.txt": 1,
		},
		"vadim": {
			"file2.txt": 1,
		},
		"kek": {
			"file1.txt": 1,
		},
		"golang": {
			"file1.txt": 1,
		},
		"lol": {
			"file1.txt": 2,
			"file2.txt": 1,
		},
	}

	actual := make(map[string]map[string]int)
	BuildIndex(actual, names)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v is not equal to expected %v", actual, expected)
	}
}

func TestFindPhrase(t *testing.T) {
	dict := make(map[string]map[string]int)
	file, _ := os.Create("file1.txt")
	defer os.Remove("file1.txt")
	file.WriteString("lol kek lol golang lol")
	file2, _ := os.Create("file2.txt")
	defer os.Remove("file2.txt")
	file2.WriteString("lol testing degree vadim")
	file.Close()
	file2.Close()

	names := []string{"file1.txt", "file2.txt"}
	expected := "File file1.txt contains 3 words of requested phrase\n\rFile file2.txt contains 2 words of requested phrase\n\r"
	phrase := []string{"lol", "testing"}
	BuildIndex(dict, names)
	actual := FindPhrase(dict, phrase)
	fmt.Println(actual)
	if actual != expected {
		t.Errorf("%v is not equal to expected %v", actual, expected)
	}

	expected = "Phrase not found\n\r"
	phrase = []string{"lol", "testing", "hello"}
	actual = FindPhrase(dict, phrase)
	if actual != expected {
		t.Errorf("%v is not equal to expected %v", actual, expected)
	}
}
