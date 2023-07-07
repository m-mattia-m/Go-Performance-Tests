# Generic binary search

Um ein Slice schneller zu durchsuchen, anstatt durch das ganze zu loopen, halbiert man es sich immer.

Wichtig, es kann aufsteigend oder absteigend sortiert werden, darum falls `<=` nicht funktioniert, verwende `>=` → wenn es nicht funktioniert merkt man daran, dass `i = len(users)` zurückgegeben wird

## Integer

Wenn die Zahlen bereits sortiert herkommen, ist ein Binery-Search massiv (teilweise bis zu 700x schneller) schneller als durch zu loopen, dafür muss aber zu 100% gewährleistet sein, dass diese auch Sortiert sind. Ansonsten ist der Loop schneller.

### Benchmarkt

Es werden zufällige Zahlen zwischen 0-200’000’000 in Go generiert und in verschieden Grossen Arrays wird danach gesucht. 

[Weitere Daten](docs/Weitere%20Daten%2045993e3d7b2546999ed5ae8b04e5b1ce.md)

| Anzahl Values | Sort I Binary-Search  | Loop |
| --- |-----------------------| --- |
| 10 | 4000ns I 0ns          | 0ns |
| 200 | 36000ns I 0ns         | 1000ns |
| 1’000 | 113000ns I 0ns        | 1000ns |
| 15’000 | 2224000ns I 1000ns    | 30000ns |
| 50’000 | 8120000ns I 1000ns    | 36000ns |
| 500’000 | 93436000ns I 1000ns   | 571000ns |
| 1’000’000 | 192711000ns I 1000ns  | 720000ns |
| 20’000’000 | 4645890000ns I 1000ns | 31653000ns |

## String

In vielen Fällen ist diese Methode nicht wirklich schnell, da das sortieren der Strings deutlich länger braucht als mit einer Loop durch alle Values zu loopen. → evtl. gibt es einen performanteren Sortier-Algorithmus → **im Normalfall macht es nur sinn eine bereits sortiertes Array mit einem Binary-Search zu durchsuchen.** 

### Benchmarkt

*Zufällige Zahlen zwischen 0-100’000’000 als String nicht der Reihe nach. → es wurde immer die Letze Zahl im Array genommen um dieses zu suchen.*

[Weitere Daten](docs/Weitere%20Daten%20455b5c39b9f34b0fb27f713ec7293657.md)

| Anzahl Values | Sort I Binary-Search | Loop |
| --- |----------------------| --- |
| 10 | 6000ns               I 1000ns | 0ns |
| 200 | 35000ns              I 1000ns | 1000ns |
| 1’000 | 131000ns             I 0ns | 2000ns |
| 15’000 | 2613000ns            I 1000ns | 25000ns |
| 50’000 | 11425000ns           I 4000ns | 258000ns |
| 500’000 | 106582000ns          I 5000ns | 1010000ns |
| 1’000’000 | 308133000ns          I 5000ns | 2589000ns |
| 20’000’000 | 11242654000ns        I 5000ns | 46937000ns |

[Hashmap](docs/Hashmap%20f055e970ce5d4a438ccdff6bb204d51c.md)

[Performat Sort](docs/Performat%20Sort%20d8910a594ecf4ffba50dc46f84e769d0.md)

```go
func searchWithBinary(key string, values []string) bool {
	sort.Strings(values)
	i := sort.Search(len(values), func(i int) bool { return key <= values[i] })
	if i < len(values) && values[i] == key {
		return true
	}
	return false
}
```

# Quellen

- [yourbasic.org](https://yourbasic.org/golang/find-search-contains-slice/)
- [pkg.go.dev](https://pkg.go.dev/sort#Search)
- [stackoverflow.com](https://stackoverflow.com/questions/28344757/golang-sort-search-cant-find-first-element-in-a-slice)


# outdated code
- [outdated exampe](docs/Old%20a099efe2d53f4bec99bb0c8f9de1d472.md)
