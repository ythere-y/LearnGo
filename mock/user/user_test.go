package user

import (
	"code.byted.org/hc_test/mock/mock"
	"fmt"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUser_GetUserInfo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var id int = 1
	mockMale := mock.NewMockMale(ctl)
	gomock.InOrder(
		mockMale.EXPECT().Get(id).Return(id))

	user := NewUser(mockMale)
	reGet := user.GetUserInfo(id)
	fmt.Printf("get result = %v\n", reGet)
	reGet = user.GetUserInfo(-100)
	fmt.Printf("get result = %v\n", reGet)
}
