package hal

// HAL is the base datastructure for creating HAL apis
type HAL struct {
	Embedded map[string]interface{} `json:"_embedded"`
	Links    map[string]interface{} `json:"_links"`
}
