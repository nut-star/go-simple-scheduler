# go-simple-scheduler
Scheduler for run functions. Go.

# Exmaples

```
type PrintTask struct {
	Message string
}

func (p PrintTask) Execute() {
	log.Info().Msg(p.Message)
}

t1 := PrintTask{Message: "Hello, Golang!"}
t2 := PrintTask{Message: "Hello, World!"}
t3 := PrintTask{Message: "Hello, Somebody!"}
taskScheduler := NewTaskScheduler(2*time.Second, []Task{t1, t2, t3})

//for period of time
taskScheduler.Start()
time.Sleep(10 * time.Second)
taskScheduler.Stop()

//inf execution, stop, when base flow stop
taskScheduler.StartInf()
```