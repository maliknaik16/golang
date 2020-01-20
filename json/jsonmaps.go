
package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	jsonStr := `
		{
			"students": [
				{
					"name": "John Doe",
					"class": "5A",
					"rno": "10"
				},
				{
					"name": "Petey Cruiser",
					"class": "6A",
					"rno": "05"
				}
			]
		}
	`

	jsonMap := make(map[string]interface{})

	err := json.Unmarshal([]byte(jsonStr), &jsonMap)

	if err != nil {
		fmt.Println(err)
	}

	students := jsonMap["students"].([]interface{})

	for _, value := range students {
		student := value.(map[string]interface{})
		fmt.Println("Name: ", student["name"])
		fmt.Println("Class: ", student["class"])
		fmt.Println("Roll Number: ", student["rno"])

		fmt.Println()
	}
}
