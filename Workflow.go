// pkg/models/models.go
package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Workflow struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Steps []Step `json:"steps"`
}

// Implement a safe UnmarshalJSON for Workflow
func (w *Workflow) UnmarshalJSON(data []byte) error {

	type Alias Workflow       // Use an alias type to
	aux := &struct{ *Alias }{ // avoid infinite recursion
		Alias: (*Alias)(w),
	}

	return json.Unmarshal(data, &aux)
}

// Custom String method for Workflow
func (w Workflow) String() string {

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Id: %s\n", w.Id))
	sb.WriteString(fmt.Sprintf("Name: %s\n", w.Name))

	if len(w.Steps) > 0 {
		sb.WriteString("Steps:\n")
		for i, step := range w.Steps {
			sb.WriteString(fmt.Sprintf(" Step %d:\n%s", i+1, step.String()))
		}
	}

	return sb.String()
}
