package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
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

type Object struct {
	Id int
}

var stringResults []Evaluation
var byteResults []Evaluation
var intResults []Evaluation

func main() {
	// The file with 20 million records was too big, so I could not upload it. 'data-20000000.csv'
	//files := []string{"data-10.csv", "data-200.csv", "data-1000.csv", "data-15000.csv", "data-50000.csv", "data-500000.csv", "data-1000000.csv", "data-20000000.csv"}
	files := []string{}

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

	arrayLength := []int{10, 200, 1000, 15000, 50000, 500000, 1000000, 20000000}
	for _, length := range arrayLength {
		rand.Seed(time.Now().Unix())
		//numbers := rand.Perm(length)

		numbers := make([]int, length)
		for i := 0; i < len(numbers); i++ {
			numbers[i] = rand.Intn(200000000)
		}
		index := rand.Intn(len(numbers)-0) + 0

		_, _, binaryDuration := searchWithBinaryInt(numbers[index], numbers)
		_, _, loopDuration := searchWithLoopInt(numbers[index], numbers)

		intResults = append(intResults, Evaluation{
			Version:        fmt.Sprintf("Int - %d", length),
			LoopDuration:   loopDuration,
			BinaryDuration: binaryDuration,
		})
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
	fmt.Println("Int-Results")
	jsonInt, _ := json.Marshal(intResults)
	fmt.Println(string(jsonInt))
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
	status := searchWithLoopString(key, values)
	loopDuration = strconv.FormatUint(uint64(time.Now().UnixNano())-uint64(startTime), 10) + "ns"
	if !status {
		return
	}
	startTime = time.Now().UnixNano()
	status, sortValue, foundValue := searchWithBinaryString(key, values)
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

func searchWithLoopString(key string, values []string) bool {
	for _, value := range values {
		if value == key {
			return true
		}
	}
	return false
}

func searchWithBinaryString(key string, values []string) (bool, string, string) {
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

func searchWithLoopInt(number int, numbers []int) (bool, int, string) {
	Intermediate := time.Now().UnixNano()
	for i, _ := range numbers {
		if number == numbers[i] {
			return true, number, strconv.FormatUint(uint64(time.Now().UnixNano())-uint64(Intermediate), 10) + "ns"
		}
	}

	return false, number, strconv.FormatUint(uint64(time.Now().UnixNano())-uint64(Intermediate), 10) + "ns"
}

func searchWithBinaryInt(number int, numbers []int) (bool, int, BinaryDuration) {
	IntermediateSort := time.Now().UnixNano()
	sort.Slice(numbers, func(i, j int) bool { return numbers[i] <= numbers[j] })
	sortDuration := strconv.FormatUint(uint64(time.Now().UnixNano())-uint64(IntermediateSort), 10) + "ns"
	IntermediateSearch := time.Now().UnixNano()
	index := sort.Search(len(numbers), func(i int) bool { return numbers[i] <= number })
	searchDuration := strconv.FormatUint(uint64(time.Now().UnixNano())-uint64(IntermediateSearch), 10) + "ns"
	wholeDuration := strconv.FormatUint(uint64(time.Now().UnixNano())-uint64(IntermediateSort), 10) + "ns"

	if index >= len(numbers) {
		return false, -1, BinaryDuration{
			Duration:      wholeDuration,
			SortDuration:  sortDuration,
			FoundDuration: searchDuration,
		}
	}

	if numbers[index] == number {
		return true, numbers[index], BinaryDuration{
			Duration:      wholeDuration,
			SortDuration:  sortDuration,
			FoundDuration: searchDuration,
		}
	}
	return false, -1, BinaryDuration{
		Duration:      wholeDuration,
		SortDuration:  sortDuration,
		FoundDuration: searchDuration,
	}
}

//func searchWithBinaryObjectId(object Object, objects []Object) (bool, int, BinaryDuration) {
//	sort.Slice(objects, func(i, j int) bool { return objects[i].Id <= objects[j].Id })
//	sort.Search(len(objects), func(i int) bool { return objects[i].Id <= object.Id })
//}
