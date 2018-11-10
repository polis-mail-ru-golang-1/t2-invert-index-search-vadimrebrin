package index

import (
	"reflect"
	"testing"
)

func TestBuildIndex(t *testing.T) {
	in := make(map[string][]string)
	in["file1.txt"] = []string{"lol", "kek", "lol", "golang"}
	in["file2.txt"] = []string{"lol", "testing", "degree", "vadim"}

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
	BuildIndex(actual, in)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v is not equal to expected %v", actual, expected)
	}
}

func TestFindPhrase(t *testing.T) {
	dict := make(map[string]map[string]int)
	in := make(map[string][]string)
	in["file1.txt"] = []string{"lol", "kek", "lol", "golang", "testing"}
	in["file2.txt"] = []string{"lol", "testing", "degree", "vadim"}

	expected := "File file1.txt contains 3 words of requested phrase\n\rFile file2.txt contains 2 words of requested phrase\n\r"
	phrase := []string{"lol", "testing"}
	BuildIndex(dict, in)
	actual := FindPhrase(dict, phrase)
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
