package background

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// MergeMap returns map concatenated two maps.
func MergeMap(a map[string][]string, b map[string][]string) map[string][]string {
	for k, v := range b {
		if _, ok := a[k]; ok {
			a[k] = append(a[k], v...)
		} else {
			a[k] = v
		}
	}

	return a
}

// Filter rerturn map[string][]string
// Read from local filter file.
// Search log with keyword based filter.
// Create Map key is filtering key, data is log.
func Filter(log string) map[string][]string {
	f, err := os.Open("filter.json")

	var key []string
	var idxMap map[string][]string

	if err != nil {
		key = make([]string, 1)
		key[0] = "err"
	} else {
		byteValue, _ := ioutil.ReadAll(f)
		json.Unmarshal(byteValue, &key)
		fmt.Println(key)
		fmt.Println(string(byteValue[1 : len(byteValue)-2]))
	}

	for _, v := range key {
		if strings.Contains(log, v) {
			idxMap[v] = make([]string, 1)
			idxMap[v][0] = log
		}
	}

	return idxMap
}
