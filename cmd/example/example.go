package main

import (
	"fmt"

	"github.com/pav5000/algoviz-go/collector"
)

func main() {
	c := collector.New(800, 600)

	step1 := c.NewStep()

	fmt.Println(string(c.Serialize()))
}
