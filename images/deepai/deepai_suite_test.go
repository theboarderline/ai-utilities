package deepai_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDeepai(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Deep AI Suite")
}
