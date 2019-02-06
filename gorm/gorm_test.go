package gorm


import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)



/**
cd myapiserver
go test -v
 */

func TestDoView(t *testing.T) {
	Convey("Given 用户在阅读", t, func() {
		ip := "127.0.0.1"
		ua := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"
		title := "golang 网络框架之 grpc"

		Convey("When 浏览了一次", func() {
			err := DoView(ip, ua, title)
			So(err, ShouldBeNil)
			Convey("Then 数据库里面应该有一条记录", func() {
				var count int
				db.Model(&View{}).Where(&View{Ip: ip, Ua: ua, Title: title}).Count(&count)
				So(count, ShouldEqual, 1)
			})
		})

		Convey("When 又浏览了一次", func() {
			err := DoView(ip, ua, title)
			So(err, ShouldBeNil)
			Convey("Then 数据库里面应该还是一条记录", func() {
				var count int
				db.Model(&View{}).Where(&View{Ip: ip, Ua: ua, Title: title}).Count(&count)
				So(count, ShouldEqual, 2)
			})
		})

		//Convey("Finally 删除记录", func() {
		//	db.Where(&View{Ip: ip, Ua: ua, Title: title}).Delete(View{})
		//})
	})
}

func TestAdd(t *testing.T) {
	Convey("将两数相加", t, func() {
		So(Add(1, 2), ShouldEqual, 3)
	})
}
