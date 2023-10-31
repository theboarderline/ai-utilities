package deepai_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/ai-utilities/images/deepai"
)

var _ = Describe("Deepai", func() {

	It("should generate an image", func() {
		res, err := deepai.GenerateImageFromText("Hello World")
		Expect(err).NotTo(HaveOccurred())
		Expect(res).NotTo(BeNil())
	})

})
