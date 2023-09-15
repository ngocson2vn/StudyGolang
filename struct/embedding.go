package main

import "fmt"

type Model struct {
	id   int
	name string
}

func (m *Model) GetName() string {
	fmt.Printf("m has address: %p\n", m)
	return m.name
}

type OnlineModel struct {
	*Model
	region string
}

func main() {
	om := OnlineModel{&Model{id: 0, name: "test_model_01"}, "maliva"}
	fmt.Printf("om has address: %p\n", &om)
	fmt.Printf("om.Model has address: %p\n\n", om.Model)
	fmt.Printf("Model name: %s\n", om.GetName())
}
