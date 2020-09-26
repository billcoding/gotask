package gotask

import (
	"github.com/robfig/cron"
	"log"
)

type Task interface {
	Name() string
	Scheduled() string
	Run() func()
}

var c = cron.New()
var close = make(chan bool, 1)

func Start() {
	defer func() {
		if re := recover(); re != nil {
			log.Fatal(re)
			Start()
		}
	}()
	c.Start()
	log.Println("gotask service started")
}

func Wait() {
	<-close
}

func Add(task Task) {
	c.AddFunc(task.Scheduled(), task.Run())
	log.Printf("gotask add task[name:%s, scheduled:%s]\n", task.Name(), task.Scheduled())
}
