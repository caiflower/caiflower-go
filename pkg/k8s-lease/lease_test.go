package k8s_lease

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"testing"
	"time"
)

var kubeconfig *string

var k8sclientset *kubernetes.Clientset

func init() {
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	k8sclientset = clientSet
}

// k8s version is v1.18.10
func TestLease(t *testing.T) {

	ctx, cancelFunc := context.WithCancel(context.Background())
	ctx1, cancelFunc1 := context.WithCancel(context.Background())
	ctx2, cancelFunc2 := context.WithCancel(context.Background())
	go leaderelection.RunOrDie(ctx, *buildLeaderElectionConfig("test1"))
	go leaderelection.RunOrDie(ctx1, *buildLeaderElectionConfig("test2"))
	go leaderelection.RunOrDie(ctx2, *buildLeaderElectionConfig("test3"))

	time.Sleep(10 * time.Second)
	cancelFunc()

	time.Sleep(10 * time.Second)
	cancelFunc1()

	time.Sleep(10 * time.Second)
	cancelFunc2()

	time.Sleep(10 * time.Second)
}

func buildLeaderElectionConfig(identity string) *leaderelection.LeaderElectionConfig {
	leaseLock := &resourcelock.LeaseLock{
		Client: k8sclientset.CoordinationV1(),
		LeaseMeta: metav1.ObjectMeta{
			Name:      "test-lease",
			Namespace: "mlops-system",
		},
		LockConfig: resourcelock.ResourceLockConfig{
			Identity: identity,
		},
	}

	return &leaderelection.LeaderElectionConfig{
		Lock:          leaseLock,
		LeaseDuration: 15 * time.Second, // 获取lease的时长
		RenewDeadline: 10 * time.Second, // 获取到lease后，续租lease的超时时间
		RetryPeriod:   2 * time.Second,  // 执行获取lease的间隔
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: func(ctx context.Context) {
				// 当前实例成为 Leader
				// 在这里执行 Leader 专属的逻辑
				fmt.Printf("i am leader \n")
			},
			OnStoppedLeading: func() {
				// 当前实例结束
				// 可以在这里执行清理工作
				fmt.Printf("clear work \n")
			},
			OnNewLeader: func(identity string) {
				// 有新的 Leader 产生
				// 可以在这实现Slaver的工作
				fmt.Printf("leader is %v \n", identity)
			},
		},
	}
}
