package timegop_test

import (
	"time"

	. "github.com/autopp/timegop"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Timegop", func() {
	Describe("Now()", func() {
		freezedTime := time.Date(2017, time.December, 6, 9, 29, 0, 0, time.Local)

		Context("Without Freeze()", func() {
			It("Returns real time", func() {
				Expect(Now()).To(BeTemporally("~", time.Now(), time.Millisecond))
			})
		})

		Context("With Freeze()", func() {
			JustBeforeEach(func() {
				Freeze(freezedTime)
			})

			It("Returns freezed time", func() {
				Expect(Now()).To(BeTemporally("==", freezedTime))
			})
		})

		Context("With Freeze() and Return()", func() {
			JustBeforeEach(func() {
				Freeze(freezedTime)
				Return()
			})

			It("Returns real time", func() {
				Expect(Now()).To(BeTemporally("~", time.Now(), time.Millisecond))
			})
		})
	})
})
