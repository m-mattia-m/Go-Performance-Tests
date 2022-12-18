package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

type Evaluation struct {
	Version         string
	LoopDuration    string
	BinaryDuration  BinaryDuration
	HashMapduration HashMapduration
}

type BinaryDuration struct {
	Duration      string
	SortDuration  string
	FoundDuration string
}

type HashMapduration struct {
	Duration      string
	MapDuration   string
	FoundDuration string
}

var stringResults []Evaluation
var byteResults []Evaluation

func main() {
	// The file with 20 million records was too big, so I could not upload it. 'data-20000000.csv'
	files := []string{"data-10.csv", "data-200.csv", "data-1000.csv", "data-15000.csv", "data-50000.csv", "data-500000.csv", "data-1000000.csv", "data-20000000.csv"}

	for i, _ := range files {
		data := getData(files, i)

		var values []string
		for y, _ := range data {
			values = append(values, data[y]...)
		}

		key := values[len(values)-2]
		StringBenchmarkt(files, key, values, i)
		ByteBenchmakt(files, key, values, i)

	}
	fmt.Println("-------------------------------------------")
	fmt.Println("String-Results")
	jsonString, _ := json.Marshal(stringResults)
	fmt.Println(string(jsonString))
	fmt.Println("-------------------------------------------")
	fmt.Println("Byte-Results")
	jsonByte, _ := json.Marshal(byteResults)
	fmt.Println(string(jsonByte))
	fmt.Println("-------------------------------------------")
}

func getData(files []string, i int) [][]string {
	file, err := os.Open("./source/" + files[i])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'
	csvReader.LazyQuotes = true
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err, "\n"+files[i])
	}

	return data
}

func StringBenchmarkt(files []string, key string, values []string, i int) {
	loopDuration := ""
	binaryDuration := ""

	var startTime = time.Now().UnixNano()
	status := searchWithLoop(key, values)
	loopDuration = strconv.FormatUint(uint64(time.Now().UnixNano())-uint64(startTime), 10) + "ns"
	if !status {
		return
	}
	startTime = time.Now().UnixNano()
	status, sortValue, foundValue := searchWithBinary(key, values)
	binaryDuration = strconv.FormatUint(uint64(time.Now().UnixNano())-uint64(startTime), 10) + "ns"
	if !status {
		return
	}
	startTime = time.Now().UnixNano()
	status, mapDuration, hashMapFoundValue := searchWithHashMap(key, values)
	hashMapDuration := strconv.FormatUint(uint64(time.Now().UnixNano())-uint64(startTime), 10) + "ns"
	if !status {
		return
	}

	stringResults = append(stringResults, Evaluation{
		Version:      files[i],
		LoopDuration: loopDuration,
		BinaryDuration: BinaryDuration{
			Duration:      binaryDuration,
			SortDuration:  sortValue,
			FoundDuration: foundValue,
		},
		HashMapduration: HashMapduration{
			Duration:      hashMapDuration,
			MapDuration:   mapDuration,
			FoundDuration: hashMapFoundValue,
		},
	})
	fmt.Println("Finished String Round " + strconv.Itoa(i) + " - " + files[i])
}

func ByteBenchmakt(files []string, key string, values []string, i int) {

}

func searchWithLoop(key string, values []string) bool {
	for _, value := range values {
		if value == key {
			return true
		}
	}
	return false
}

func searchWithBinary(key string, values []string) (bool, string, string) {
	var startTime = time.Now().UnixNano()
	sort.Strings(values)
	Intermediate := time.Now().UnixNano()
	sortValue := strconv.FormatUint(uint64(Intermediate)-uint64(startTime), 10) + "ns"
	i := sort.Search(len(values), func(i int) bool { return key <= values[i] })
	if i < len(values) && values[i] == key {
		foundValue := strconv.FormatUint(uint64(time.Now().UnixNano())-uint64(Intermediate), 10) + "ns"
		return true, sortValue, foundValue
	}
	return false, "", ""
}

func searchWithHashMap(key string, values []string) (bool, string, string) {
	var startTime = time.Now().UnixNano()
	valuesMap := make(map[string]string)
	for _, word := range values {
		valuesMap[word] = word
	}
	Intermediate := time.Now().UnixNano()
	hashMapDuration := strconv.FormatUint(uint64(Intermediate)-uint64(startTime), 10) + "ns"

	if _, foundState := valuesMap[key]; foundState {
		foundDuration := strconv.FormatUint(uint64(time.Now().UnixNano())-uint64(Intermediate), 10) + "ns"
		return true, hashMapDuration, foundDuration
	}
	return false, "", ""
}
