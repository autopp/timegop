package timegop_test

import (
	"time"

	. "github.com/autopp/timegop"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

var _ = Describe("Timegop", func() {
	t := time.Date(2017, time.December, 6, 9, 29, 0, 0, time.Local)

	BeRealTime := func() types.GomegaMatcher {
		return BeTemporally("~", time.Now(), time.Millisecond)
	}

	Describe("Now()", func() {
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

	Describe("Since()", func() {
		u := time.Date(2017, time.December, 6, 8, 0, 0, 0, time.Local)

		Context("Without Freeze()", func() {
			It("Returns duration from argument to the real time", func() {
				Expect(Since(u)).To(BeNumerically("~", time.Since(u), time.Millisecond))
			})
		})

		Context("With Freeze()", func() {
			JustBeforeEach(func() {
				Freeze(t)
			})

			AfterEach(func() {
				Return()
			})

			It("Returns duration from argument to the freezed time", func() {
				expected, _ := time.ParseDuration("+1h29m")
				Expect(Since(u)).To(Equal(expected))
			})
		})
	})

	Describe("Until()", func() {
		u := time.Date(2017, time.December, 6, 8, 0, 0, 0, time.Local)

		Context("Without Freeze()", func() {
			It("Returns duration from the freezed time to argument", func() {
				Expect(Until(u)).To(BeNumerically("~", time.Until(u), time.Millisecond))
			})
		})

		Context("With Freeze()", func() {
			JustBeforeEach(func() {
				Freeze(t)
			})

			AfterEach(func() {
				Return()
			})

			It("Returns duration from the freezed time to argument", func() {
				expected, _ := time.ParseDuration("-1h29m")
				Expect(Until(u)).To(Equal(expected))
			})
		})
	})
})
