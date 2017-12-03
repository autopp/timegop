package timegop_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTimegop(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Timegop Suite")
}
