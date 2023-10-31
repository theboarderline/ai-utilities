package dalle_test

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sashabaranov/go-openai"
	"github.com/theboarderline/ai-utilities/images/dalle"
	"os"
)

var _ = Describe("Dalle", func() {

	var (
		dalleClient *dalle.Client
	)

	BeforeEach(func() {
		config := openai.DefaultConfig(os.Getenv("OPENAI_API_KEY"))
		config.OrgID = os.Getenv("OPENAI_ORG")
		dalleClient = dalle.NewClient(config)
	})

	It("should generate an image", func() {
		res, err := dalleClient.GenerateImage(context.Background(), "Hello World")
		Expect(err).NotTo(HaveOccurred())
		Expect(res).NotTo(BeNil())
	})

})
