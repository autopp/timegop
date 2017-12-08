package timegop_test

import (
	"time"

	. "github.com/autopp/timegop"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

var _ = Describe("Timegop", func() {
	BeRealTime := func() types.GomegaMatcher {
		return BeTemporally("~", time.Now(), time.Millisecond)
	}

	Describe("Now()", func() {
		t := time.Date(2017, time.December, 6, 9, 29, 0, 0, time.Local)

		Context("Without Freeze()", func() {
			It("Returns real time", func() {
				Expect(Now()).To(BeRealTime())
			})
		})

		Context("With Freeze()", func() {
			JustBeforeEach(func() {
				Freeze(t)
			})

			AfterEach(func() {
				Return()
			})

			It("Returns freezed time", func() {
				Expect(Now()).To(BeTemporally("==", t))
			})
		})

		Context("With Freeze() and Return()", func() {
			JustBeforeEach(func() {
				Freeze(t)
				Return()
			})

			It("Returns real time", func() {
				Expect(Now()).To(BeRealTime())
			})
		})

		Context("With Freeze() and defered calling returned value", func() {
			JustBeforeEach(func() {
				defer Freeze(t)()
			})

			It("Returns real time", func() {
				Expect(Now()).To(BeRealTime())
			})
		})
	})
})
