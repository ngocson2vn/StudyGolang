package main

import (
	"fmt"
	"time"
)

type TaskAPI interface {
	Start() error
	Stop() error
}

type LagrangeTask struct {
	Model string
	PSM   string
}

func (l *LagrangeTask) Start() error {
	fmt.Printf("[LagrangeTask] PSM %s starts serving model %s\n", l.PSM, l.Model)
	return nil
}

func (l *LagrangeTask) Stop() error {
	fmt.Printf("[LagrangeTask] PSM %s stops serving model %s\n", l.PSM, l.Model)
	return nil
}

func createTask(model string, psm string) (TaskAPI, error) {
	lt := &LagrangeTask{Model: model, PSM: psm}
	return lt, nil
}

func main() {
	task, _ := createTask("test_model_1", "test_psm_1")
	fmt.Printf("task has type %T\n\n", task)

	task.Start()
	for i := 0; i < 10; i++ {
		fmt.Printf(".")
		time.Sleep(1 * time.Second)
	}
	fmt.Println()
	task.Stop()
}
