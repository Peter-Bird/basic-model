package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Converts a Workflow struct to a JSON string
func WorkflowToJSON(workflow Workflow) (string, error) {
	jsonData, err := json.Marshal(workflow)
	if err != nil {
		return "", fmt.Errorf("failed to convert Workflow to JSON: %w", err)
	}
	return string(jsonData), nil
}

// Converts a JSON string to a Workflow struct
func JSONToWorkflow(jsonStr string) (*Workflow, error) {
	var workflow Workflow
	err := json.Unmarshal([]byte(jsonStr), &workflow)
	if err != nil {
		return nil, fmt.Errorf("failed to convert JSON to Workflow: %w", err)
	}
	return &workflow, nil
}

func PrettyPrintJSON(jsonStr string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(jsonStr), "", "  "); err != nil {
		return "", fmt.Errorf("failed to pretty print JSON: %w", err)
	}
	return prettyJSON.String(), nil
}
