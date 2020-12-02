package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/ritarock/ctj/lib/date"
)

func TestReader(t *testing.T) {
	filePath := "../../test_data/test.csv"
	result := Reader(filePath)
	expect := [][]string{
		{"col1", "col2", "col3", "col4", "col5", "col6"},
		{"val01", "val02", "val03", "val04", "val5", "val06"},
		{"val11", "val12", "val13", "val14", "val5", "val16"},
		{"val21", "val22", "val23", "val24", "val5", "val26"},
	}

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("result = %q, want %q", result, expect)
	}
}

func TestMakeJsonFile(t *testing.T) {
	filePath := "../../test_data/test.csv"
	outputFile := strings.Split(filepath.Base(filePath), ".")[0] + "_to_json_" + date.ToYYYYMMDD() + ".json"
	resultArray := []string{
		"{\"col1\":\"val01\",\"col2\":\"val02\",\"col3\":\"val03\",\"col4\":\"val04\",\"col5\":\"val5\",\"col6\":\"val06\"}",
		"{\"col1\":\"val11\",\"col2\":\"val12\",\"col3\":\"val13\",\"col4\":\"val14\",\"col5\":\"val5\",\"col6\":\"val16\"}",
		"{\"col1\":\"val21\",\"col2\":\"val22\",\"col3\":\"val23\",\"col4\":\"val24\",\"col5\":\"val5\",\"col6\":\"val26\"}",
	}
	expect := "[{\"col1\":\"val01\",\"col2\":\"val02\",\"col3\":\"val03\",\"col4\":\"val04\",\"col5\":\"val5\",\"col6\":\"val06\"}," +
		"{\"col1\":\"val11\",\"col2\":\"val12\",\"col3\":\"val13\",\"col4\":\"val14\",\"col5\":\"val5\",\"col6\":\"val16\"}," +
		"{\"col1\":\"val21\",\"col2\":\"val22\",\"col3\":\"val23\",\"col4\":\"val24\",\"col5\":\"val5\",\"col6\":\"val26\"}]"

	MakeJsonFile(outputFile, resultArray)

	f, err := os.Open(outputFile)
	if err != nil {
		t.Fatalf("err = %v", err)
	}

	defer func() {
		err := os.Remove(outputFile)
		if err != nil {
			t.Fatalf("err = %v", err)
		}
	}()

	defer f.Close()

	result, err := ioutil.ReadAll(f)

	if string(result) != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}
