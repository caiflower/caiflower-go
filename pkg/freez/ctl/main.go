package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 1. 启动子进程（传递父进程PID给子进程）
	parentPID := os.Getpid()
	cmd := exec.Command("./pkg/freez/ctl/fsfreezer", fmt.Sprintf("%d", parentPID))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("启动子进程失败:", err)
		return
	}
	fmt.Println("子进程PID:", cmd.Process.Pid)

	// 2. 等待SIGUSR1信号
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGUSR1)

	fmt.Println("等待子进程SIGUSR1（冻结完成通知）...")
	<-sigs // 阻塞直到收到SIGUSR1
	fmt.Println("收到SIGUSR1，开始快照流程...")

	// 3. 模拟快照操作
	time.Sleep(2 * time.Second)
	fmt.Println("快照完成，通知子进程继续...")

	// 4. 给子进程发SIGUSR1
	err := cmd.Process.Signal(syscall.SIGUSR1)
	if err != nil {
		fmt.Println("发送SIGUSR1给子进程失败:", err)
	}

	// 5. 等待子进程退出
	cmd.Wait()
	fmt.Println("子进程退出，流程结束")
}
