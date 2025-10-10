package diskfree

import (
	"encoding/json"
	"fmt"
)

type DiskFree struct {
	Discarray []Discarray `json:"discarray"`
}
type Discarray struct {
	Mount string `json:"mount"`
	Size  string `json:"size"`
	Used  string `json:"used"`
	Avail string `json:"avail"`
	Use   string `json:"use%"`
}

func UnmarshalJSON(output []byte) (*DiskFree, error) {
	var df DiskFree

	err := json.Unmarshal(output, &df)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshal stdout: %w", err)
	}

	return &df, nil
}
