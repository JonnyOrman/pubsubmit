package main

import "github.com/jonnyorman/pubsubmit"

type TestModel struct {
	Prop1 string `json:"prop1"`
	Prop2 int    `json:"prop2"`
}

func main() {
	pubsubmit.RunTyped[TestModel]()
}
