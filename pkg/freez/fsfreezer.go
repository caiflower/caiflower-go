package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("缺少父进程PID参数")
		return
	}
	parentPid, _ := strconv.Atoi(os.Args[1])

	// 1. 执行冻结操作（模拟sleep 2秒）
	fmt.Println("fsfreezer: 正在冻结文件系统...")
	time.Sleep(2 * time.Second)
	fmt.Println("fsfreezer: 冻结完成，通知父进程...")

	// 2. 给父进程发SIGUSR1
	syscall.Kill(parentPid, syscall.SIGUSR1)

	// 3. 等待父进程的SIGUSR1
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGUSR1)
	fmt.Println("fsfreezer: 等待父进程完成快照...")
	<-sigs
	fmt.Println("fsfreezer: 收到父进程SIGUSR1，解冻并退出")
}
