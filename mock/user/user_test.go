package user

import (
	"code.byted.org/gopkg/gomonkey"
	"code.byted.org/hc_test/mock/person"
	"testing"

	. "code.byted.org/gopkg/mockito"
	"code.byted.org/hc_test/mock/mock"
	"github.com/golang/mock/gomock"
	"github.com/smartystreets/goconvey/convey"
)

func TestUser_GetUserInfo(t *testing.T) {
	type args struct {
		id int
	}
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	PatchConvey("test get user info", t, func() {
		mockMale := mock.NewMockMale(ctl)
		PatchConvey("test id = 1", func() {
			var (
				mockReq    = 1
				mockResp   = 1
				testReq    = 1
				expectResp = 1
			)

			gomock.InOrder(
				mockMale.EXPECT().Get(mockReq).Return(mockResp).AnyTimes())

			user := NewUser(mockMale)
			reGet := user.GetUserInfo(testReq)
			convey.So(reGet, convey.ShouldEqual, expectResp)
		})
		PatchConvey("test id = -1", func() {
			var (
				testReq    = -1
				expectResp = -1
			)

			user := NewUser(mockMale)
			reGet := user.GetUserInfo(testReq)
			convey.So(reGet, convey.ShouldEqual, expectResp)
		})
	})
}

func TestUser_GetUserName(t *testing.T) {
	type args struct {
		id int
	}
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	PatchConvey("test get user name", t, func() {
		male := &person.RealMale{}
		user := NewUser(male)
		PatchConvey("test idx = 1", func() {
			gomonkey.ApplyFunc(person.GetName, func(idx int) string {
				return "Queen"
			})
			idx := 1
			expectName := "Queen"
			name := user.GetUserName(idx)
			convey.So(name, convey.ShouldEqual, expectName)
		})
	})
}
