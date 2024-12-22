package schedulers

import (
	"time"
)

type Task interface {
	Execute()
}

type Scheduler interface {
	Start()
	Stop()
	StartInf()
}

type TaskScheduler struct {
	done          chan bool
	ticker        *time.Ticker
	RepeatedTasks []Task
}

func NewTaskScheduler(interval time.Duration, tasks []Task) *TaskScheduler {
	return &TaskScheduler{
		done:          make(chan bool),
		ticker:        time.NewTicker(interval),
		RepeatedTasks: tasks,
	}
}

func (s *TaskScheduler) Start() {
	go func() {
		for {
			select {
			case <-s.done:
				return
			case <-s.ticker.C:
				for _, task := range s.RepeatedTasks {
					go task.Execute()
				}
			}
		}
	}()
}

func (s *TaskScheduler) Stop() {
	s.ticker.Stop()
	s.done <- true
}

func (s *TaskScheduler) StartInf() {
	go func() {
		for range s.ticker.C {
			for _, task := range s.RepeatedTasks {
				go task.Execute()
			}
		}
	}()
	select {}
}
