package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"release"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "test1", Year: 1942, Color: false, Actors: []string{"1111", "2222"}},
	{Title: "test2", Year: 1943, Color: true, Actors: []string{"1231", "1123"}},
	{Title: "test3", Year: 1944, Color: false, Actors: []string{"423", "443"}},
	{Title: "test4", Year: 1945, Color: true, Actors: []string{"54", "552"}},
	{Title: "test5", Year: 1946, Color: false, Actors: []string{"36345", "1231"}},
}

func main() {
	test1()
	test2()
	test3()
}

func test1() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}

func test2() {
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}

func test3() {
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}
