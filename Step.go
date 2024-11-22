// pkg/models/models.go
package models

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Step struct {
	Endpoint     string      `json:"endpoint"`
	Method       string      `json:"method"`
	Parameters   interface{} `json:"parameters,omitempty"`
	Dependencies []string    `json:"dependencies,omitempty"`
}

// Implement a safe UnmarshalJSON for Step
func (s *Step) UnmarshalJSON(data []byte) error {

	type Alias Step // Use an alias type to avoid infinite recursion
	aux := &struct {
		Parameters json.RawMessage `json:"parameters,omitempty"` // Keep raw for parsing later
		*Alias
	}{
		Alias: (*Alias)(s),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Dynamically parse Parameters
	if aux.Parameters != nil {
		var params interface{}
		if err := json.Unmarshal(aux.Parameters, &params); err == nil {
			s.Parameters = params // Assign parsed value
		} else {
			s.Parameters = aux.Parameters // Fallback to raw message
		}
	}

	return nil
}

// Custom String method for Step
func (s Step) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("  Endpoint: %s\n", s.Endpoint))
	sb.WriteString(fmt.Sprintf("  Method: %s\n", s.Method))

	if s.Parameters != nil {
		sb.WriteString("  Parameters: ")
		prettyParameters := prettyPrintParameters(s.Parameters)
		sb.WriteString(prettyParameters + "\n")
	}

	if len(s.Dependencies) > 0 {
		sb.WriteString(fmt.Sprintf("  Dependencies: %v\n", s.Dependencies))
	}
	return sb.String()
}

// Pretty-print parameters
func prettyPrintParameters(params interface{}) string {
	switch v := params.(type) {
	case map[string]interface{}:
		var sb strings.Builder
		sb.WriteString("{\n")
		for key, value := range v {
			sb.WriteString(fmt.Sprintf("    %s: %v\n", key, value))
		}
		sb.WriteString("  }")
		return sb.String()
	case string:
		return v
	case []byte:
		return string(v)
	default:
		// Use reflection for unknown types
		if reflect.ValueOf(params).Kind() == reflect.Map {
			return fmt.Sprintf("%v", params)
		}
		return fmt.Sprintf("%v", params)
	}
}
