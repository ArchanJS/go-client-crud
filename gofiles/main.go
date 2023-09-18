package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name   string   `json:"coursename"`
	Author string   `json:"author"`
	Price  int      `json:"price"`
	Tags   []string `json:"tags,omitempty"`
}

func convertIntoJson() {
	courses := []course{
		{Name: "web dev", Author: "archan", Price: 900, Tags: []string{"web-dev", "javascript"}},
		{Name: "ml", Author: "sudin", Price: 1000, Tags: []string{"ml", "python"}},
		{Name: "app dev", Author: "jaydip", Price: 1200, Tags: nil},
	}
	coursesJson, _ := json.MarshalIndent(courses, "", " ")
	fmt.Printf("%s", coursesJson)
}

func consumeJson() {
	rawData := []byte(`
	{
		"coursename":"blockchain",
		"author":"harry",
		"price":500,
		"tags":["blockchain","web-dev","decentralization"]
	}
	`)
	isValid := json.Valid(rawData)
	if isValid == false {
		panic("JSON's format is wrong!")
	} else {
		// var testCourse course
		var testMapData map[string]interface{}
		// err := json.Unmarshal(rawData, &testCourse) //using course instance
		err := json.Unmarshal(rawData, &testMapData) // using map
		if err != nil {
			panic(err)
		}
		// fmt.Printf("%#v\n", &testMapData)

		//using for loop to print map's data
		for key, val := range testMapData {
			fmt.Printf("for key: %v, value is: %v\n", key, val)
		}

	}
}

func main() {
	// convertIntoJson()
	consumeJson()
}
