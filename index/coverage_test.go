package index

import (
	"reflect"
	"testing"
)

func TestBuildIndex(t *testing.T) {
	in := make(map[string][]string)
	in["file1.txt"] = []string{"lol", "kek", "lol", "golang"}
	in["file2.txt"] = []string{"lol", "testing", "degree", "vadim"}
	expected := make(map[string]map[string]int)
	{
		tmp := make(map[string]int)
		tmp["file1.txt"] = 2
		expected["lol"] = tmp
	}
	expected["lol"]["file2.txt"] = 1
	{
		tmp := make(map[string]int)
		tmp["file1.txt"] = 1
		expected["kek"] = tmp
		expected["golang"] = tmp
	}
	{
		tmp := make(map[string]int)
		tmp["file2.txt"] = 1
		expected["testing"] = tmp
		expected["degree"] = tmp
		expected["vadim"] = tmp
	}
	actual := make(map[string]map[string]int)
	BuildIndex(actual, in)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v is not equal to expected %v", actual, expected)
	}
}
