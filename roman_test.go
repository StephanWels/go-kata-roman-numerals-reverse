package roman

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRomanNumeralEncoder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Roman Numeral Suite")
}

var _ = Describe("test decimal to roman numeral converter", func() {
	It("should give roman numeral for positive integer", func() {
// 		Expect(Decode(...)).To(Equal(...))
	})
})
