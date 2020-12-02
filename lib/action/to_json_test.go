package action

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/ritarock/ctj/lib/date"
)

func TestToJson(t *testing.T) {
	filePath := "../../test_data/test.csv"
	resultFilePath := "./test_to_json_" + date.ToYYYYMMDD() + ".json"
	expectJson := "[{\"col1\":\"val01\",\"col2\":\"val02\",\"col3\":\"val03\",\"col4\":\"val04\",\"col5\":\"val5\",\"col6\":\"val06\"}," +
		"{\"col1\":\"val11\",\"col2\":\"val12\",\"col3\":\"val13\",\"col4\":\"val14\",\"col5\":\"val5\",\"col6\":\"val16\"}," +
		"{\"col1\":\"val21\",\"col2\":\"val22\",\"col3\":\"val23\",\"col4\":\"val24\",\"col5\":\"val5\",\"col6\":\"val26\"}]"

	ToJson(filePath)

	expectResultFilePath := func() bool {
		_, err := os.Stat(resultFilePath)
		return err == nil
	}()

	if !expectResultFilePath {
		t.Errorf("want = %v", resultFilePath)
	}

	f, err := os.Open(resultFilePath)
	if err != nil {
		t.Fatalf("err = %v", err)
	}

	defer func() {
		err := os.Remove(resultFilePath)
		if err != nil {
			t.Fatalf("err = %v", err)
		}
	}()
	defer f.Close()
	resultJson, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("err = %v", err)
	}

	if string(resultJson) != expectJson {
		t.Errorf("result = %v, want = %v", string(resultJson), expectJson)
	}
}
