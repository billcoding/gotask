package gotask

import (
	"fmt"
	"testing"
)

//Define printTask struct
type printTask struct {
}

//Implements Task Name method
func (printTask) Name() string {
	return "printTask"
}

//Implements Task Scheduled method
func (printTask) Scheduled() string {
	return "@every 1s"
}

//Implements Task Run method
func (printTask) Run() func() {
	return func() {
		fmt.Println("helloworld")
	}
}

//Test Task
func TestTask(t *testing.T) {
	Add(printTask{})
	Start()
	Wait()
}
