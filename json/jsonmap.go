
package main

import (
  "fmt"
  "encoding/json"
)

func main() {
  jsonStr1 := `
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

  jsonStr2 := `
    {
      "students": {
          "name": "John Doe",
          "class": "5A",
          "rno": "10"
        }
    }
  `

  jsonMap := make(map[string]interface{})

  err := json.Unmarshal([]byte(jsonStr1), &jsonMap)

  if err != nil {
    fmt.Println(err)
  }

  if _, ok := jsonMap["students"].([]interface{}); ok {

    students := jsonMap["students"].([]interface{})

    for _, value := range students {
      student := value.(map[string]interface{})
      fmt.Println("Name: ", student["name"])
      fmt.Println("Class: ", student["class"])
      fmt.Println("Roll Number: ", student["rno"])

      fmt.Println()
    }
  } else {
    student := jsonMap["students"].(map[string]interface{})
    fmt.Println("Name: ", student["name"])
    fmt.Println("Class: ", student["class"])
    fmt.Println("Roll Number: ", student["rno"])
  }

  if t := fmt.Sprintf("%T", jsonStr2); t == "string" {
    // DO NOTHING
  }
}
