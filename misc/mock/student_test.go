package mock

import (
	"fmt"
	"github.com/golang/mock/gomock"
	mock_mockDemo "golang-example/misc/mock/mock"
	"testing"
)

func TestGetPeopleName(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	// 构造 mock 类
	mockPeople := mock_mockDemo.NewMockIPeople(mockCtl)

	// 注入期望的返回值
	mockPeople.EXPECT().GetName().Return("height")

	mockedName := GetPeopleName(mockPeople)
	if "height" != mockedName {
		t.Error("Get wrong name: ", mockedName)
	}
}

func TestSetPeopleName(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	// 构造 mock 类
	mockPeople := mock_mockDemo.NewMockIPeople(mockCtl)

	//指定输入参数，返回指定结果
	mockPeople.EXPECT().SetName(gomock.Eq("he")).Return("ok")
	//输出参不做指定，但是指定返回结果
	mockPeople.EXPECT().SetName(gomock.Any()).Do(func(format string) {
		fmt.Println("recv param2 :", format)
	}).Return("ok1")

	mockedSetName := SetPeopleName(mockPeople, "he")
	fmt.Println("mockedSetName: ", mockedSetName)
	if "ok" != mockedSetName {
		t.Error("Set wrong name2: ", mockedSetName)
	}

	mockedSetName = SetPeopleName(mockPeople, "al222")
	fmt.Println("mockedSetName: ", mockedSetName)
	if "ok1" != mockedSetName {
		t.Error("Set wrong name2: ", mockedSetName)
	}
}
