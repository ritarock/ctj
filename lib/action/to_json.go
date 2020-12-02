package action

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/ritarock/ctj/lib/date"
	"github.com/ritarock/ctj/lib/file"
)

func ToJson(filePath string) {
	outputFile := strings.Split(filepath.Base(filePath), ".")[0] + "_to_json_" + date.ToYYYYMMDD() + ".json"
	fmt.Println(outputFile)
	records := file.Reader(filePath)

	m := make(map[string]string)
	var resultArray []string

	for i := 1; i < len(records); i++ {
		for j := 0; j < len(records[0]); j++ {
			m[records[0][j]] = records[i][j]
		}
		bytes, err := json.Marshal(m)
		if err != nil {
			log.Fatal(err)
			return
		}
		resultArray = append(resultArray, string(bytes))
	}
	file.MakeJsonFile(outputFile, resultArray)
}
