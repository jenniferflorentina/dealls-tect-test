package server_test

import (
	"context"
	"net/http"

	"jennifer/dealls-tech-test/internal/server"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server", func() {
	Context("Server", func() {
		var httpServer *http.Server

		ctx := context.Background()
		It("should be able to start", func() {
			httpServer = server.InitServer(ctx, nil)
			Expect(httpServer).NotTo(BeNil())
		})
		It("should be able to be stopped", func() {
			Expect(httpServer.Shutdown(ctx)).To(Succeed())
		})
		It("should not be able to listen if the server is already closed", func() {
			Expect(httpServer.ListenAndServe()).NotTo(Succeed())
		})
	})
})
