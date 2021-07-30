package background

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"regexp"
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
// Search log with filterword based filter.
// Create Map filter is filtering key, data is log.
func Filter(log string) error {
	// Open은 한 Time만 호출하고
	// 뒤에서부터는 seek로 핸들링
	f, err := os.Open("filter.json")
	defer f.Close()

	var filter []string

	if err != nil {
		filter = make([]string, 1)
		filter[0] = "kijunking"
	} else {
		byteValue, _ := ioutil.ReadAll(f)
		json.Unmarshal(byteValue, &filter)
	}
	for _, v := range filter {
		// regex로 체크
		if ok, _ := regexp.MatchString(v, log); ok == true {
			return nil
		}
	}
	return errors.New("no filtered messages")
}
