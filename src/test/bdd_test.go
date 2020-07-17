package test

import (
	// 相当于引入的外部依赖可以直接在当前包下引用
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBdd(t *testing.T) {
	Convey("Given 2 even numbers", t, func() {
		a := 2
		b := 4

		Convey("When add the two numbers", func() {
			c := a + b

			Convey("Then the result is still even", func() {
				So(c%2, ShouldEqual, 0)
			})
		})
	})
}