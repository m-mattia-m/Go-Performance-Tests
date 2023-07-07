# Hashmap

Ist schneller als ein Binery-Search, vor allem wenn man die Values noch sortieren muss, ist aber langsamer als eine normale iteration.

```go
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
```