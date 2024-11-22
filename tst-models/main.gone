package main

import (
	"fmt"

	"github.com/Peter-Bird/models"
)

func main() {
	// Create a Workflow object
	workflow := models.Workflow{
		Id:   "1",
		Name: "Sample Workflow",
		Steps: []models.Step{
			{
				Endpoint:   "/example",
				Method:     "POST",
				Parameters: map[string]interface{}{"key": "value"},
				Dependencies: []string{
					"Step1",
				},
			},
		},
	}

	// Convert Workflow to JSON
	jsonStr, err := models.WorkflowToJSON(workflow)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	js, err := models.PrettyPrintJSON(jsonStr)
	fmt.Println("Workflow as JSON:", js)

	// Convert JSON back to Workflow
	parsedWorkflow, err := models.JSONToWorkflow(jsonStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Pretty print the workflow
	fmt.Println("Workflow as Struct:")
	fmt.Println(parsedWorkflow.String())
}
