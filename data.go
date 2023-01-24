package godo

import (
	"encoding/json"
	"os"

	"github.com/graph-gophers/graphql-go"
)

type Task struct {
	ID          graphql.ID   `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	StartDate   graphql.Time `json:"startDate"`
	EndDate     graphql.Time `json:"endDate"`
	Done        bool         `json:"done"`
}

func importJSONDataFromFile(filename string, result any) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, result)
	if err != nil {
		return err
	}

	return nil
}

var tasks []Task
var _ = importJSONDataFromFile("./tasks.json", &tasks)
