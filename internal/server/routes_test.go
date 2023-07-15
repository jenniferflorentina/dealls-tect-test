package server_test

import (
	"jennifer/dealls-tech-test/internal/server"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Routes", func() {
	Context("creating routes", func() {
		It("should be able to create new router handler", func() {
			Expect(server.Routes(nil)).NotTo(BeNil())
		})
	})
})
