package dalle_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDalle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dalle AI Suite")
}
