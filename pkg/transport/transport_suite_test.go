package transport

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTransport(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Transport Suite")
}
