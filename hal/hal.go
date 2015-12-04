package hal

import "encoding/json"

// HAL is the base datastructure for creating HAL apis
type HAL struct {
	Embedded map[string]interface{}
	Links    map[string]interface{}
	Data     interface{}
}

// JSON parses HAL into correct JSON
func JSON(hal HAL) string {
	t := Translate(hal)
	parsed, _ := json.Marshal(t)
	return string(parsed)
}

func asMap(post interface{}) map[string]interface{} {
	byt, _ := json.Marshal(post)
	var dat map[string]interface{}
	json.Unmarshal(byt, &dat)
	return dat
}

func Translate(hal HAL) map[string]interface{} {
	m := map[string]interface{}{}

	if hal.Embedded != nil {
		m["_embedded"] = hal.Embedded
	}

	if hal.Links != nil {
		m["_links"] = hal.Links
	}

	for k, v := range asMap(hal.Data) {
		m[k] = v
	}
	return m
}
