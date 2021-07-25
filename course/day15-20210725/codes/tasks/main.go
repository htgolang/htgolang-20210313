package main

import (
	"context"
	"fmt"
	"time"

	"github.com/beego/beego/v2/task"
)

func main() {
	t := task.NewTask("t1", "*/5 * * * * *", func(ctx context.Context) error {
		fmt.Println("t1", time.Now())
		return nil
	})

	// t.Run(context.Background())
	task.AddTask("t1", t)

	task.StartTask()
	defer task.StopTask()

	time.Sleep(3 * time.Minute)
}
