package task

import "fmt"

type Task interface {
	Run()
}

func Loop(done <-chan any, task func()) {
	for {
		select {
		case <-done:
			fmt.Println("loop is done!")
			return
		default:
			task()
		}
	}
}
