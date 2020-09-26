package gotask

import (
	"fmt"
	"testing"
)

type printTask struct {
}

func (printTask) Name() string {
	return "printTask"
}

func (printTask) Scheduled() string {
	return "@every 1s"
}

func (printTask) Run() func() {
	return func() {
		fmt.Println("helloworld")
	}
}

func TestTask(t *testing.T) {
	Add(printTask{})
	Start()
	Wait()
}
