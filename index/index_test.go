package index

import (
	"reflect"
	"testing"
)

func TestBuildIndex(t *testing.T) {
	in := make(map[string][]string)
	in["file1.txt"] = []string{"lol", "kek", "lol", "golang"}
	in["file2.txt"] = []string{"lol", "testing", "degree", "vadim"}

	expected := Index{
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

	actual := make(Index)
	BuildIndex(actual, in)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v is not equal to expected %v", actual, expected)
	}
}

func TestFindPhrase(t *testing.T) {
	dict := make(Index)
	in := make(map[string][]string)
	in["file1.txt"] = []string{"lol", "kek", "lol", "golang", "testing"}
	in["file2.txt"] = []string{"lol", "testing", "degree", "vadim"}

	expected := []Result{{"file1.txt", 3},
		{"file2.txt", 2}}
	phrase := []string{"lol", "testing"}
	BuildIndex(dict, in)
	actual := FindPhrase(dict, phrase)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v is not equal to expected %v", actual, expected)
	}

	expected = nil
	phrase = []string{"lol", "testing", "hello"}
	actual = FindPhrase(dict, phrase)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v is not equal to expected %v", actual, expected)
	}
}
