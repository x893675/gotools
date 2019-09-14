package main

import (
	"fmt"
	"time"
)

func main() {
	Test()
}
func Run(task_id, sleeptime, timeout int, ch chan string) {
	ch_run := make(chan string)
	go run(task_id, sleeptime, ch_run)
	select {
	case re := <-ch_run:
		ch <- re
	case <-time.After(time.Duration(timeout) * time.Second):
		re := fmt.Sprintf("task id %d, timeout ", task_id)
		ch <- re
	}
}
func run(task_id, sleeptime int, ch chan string) {
	time.Sleep(time.Duration(sleeptime) * time.Second)
	ch <- fmt.Sprintf("task id %d,sleep %d second", task_id, sleeptime)
	return
}
func Test() {
	input := []int{5, 3, 8}
	timeout := 8
	// 创建N个任务管道，用来接收各个并发任务的完成结果
	chs := make([]chan string, len(input))
	defer func() {
		for _, c := range chs {
			if c != nil {
				close(c)
			}
		}
	}()
	sTime := time.Now()
	fmt.Println("start")
	for i, sleeptime := range input {
		chs[i] = make(chan string)
		go Run(i, sleeptime, timeout, chs[i])
	}
	// 获取结果
	for _, ch := range chs {
		fmt.Println(<-ch)
	}
	eTime := time.Now()
	fmt.Printf("finished,Process time %s. Number of task is %d \n", eTime.Sub(sTime), len(input))
}
