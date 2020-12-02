package file

import (
	"encoding/csv"
	"log"
	"os"
)

func Reader(filePath string) [][]string {
	readerFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(readerFile)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}

func MakeJsonFile(outPutfile string, resultArray []string) {
	file, err := os.Create(outPutfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString("[")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(resultArray); i++ {
		if i+1 == len(resultArray) {
			_, err := file.WriteString(resultArray[i])
			if err != nil {
				log.Fatal(err)
			}
		} else {
			_, err = file.WriteString(resultArray[i] + ",")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	_, err = file.WriteString("]")
	if err != nil {
		log.Fatal(err)
	}
}
