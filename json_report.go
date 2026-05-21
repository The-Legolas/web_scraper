package main

import (
	"encoding/json"
	"os"
	"sort"
)

func writeJSONReport(pages map[string]PageData, filename string) error {
	var keys []string

	for key := range pages {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var sorted []PageData
	for _, val := range keys {
		sorted = append(sorted, pages[val])
	}

	data, err := json.MarshalIndent(sorted, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
