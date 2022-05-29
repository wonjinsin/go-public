package model

import (
	"testing"

	. "github.com/onsi/ginkgo"
	gm "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	gm.RegisterFailHandler(Fail)
	RunSpecs(t, "Test Model Suite")
}
