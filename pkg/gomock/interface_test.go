package gomock

import (
	"fmt"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestPeople(t *testing.T) {
	//创建gomock控制器，用来记录后续操作信息
	mockCtl := gomock.NewController(t)
	//调用mock文件中的NewMockMyInter方法，创建一个MyInter接口的mock实例
	mockMyInter := NewMockPeople(mockCtl)

	//EXPECT()接口设置预期返回值
	mockMyInter.EXPECT().DescribeAge().Return("11")
	mockMyInter.EXPECT().DescribeName().Return("caiflower")
	mockMyInter.EXPECT().SaySame("hello world").Return("hello world")

	//将mock的MyInterface传入GetUser函数
	fmt.Println(mockMyInter.DescribeAge())
	fmt.Println(mockMyInter.DescribeName())
	fmt.Println(mockMyInter.SaySame("hello world"))
}
