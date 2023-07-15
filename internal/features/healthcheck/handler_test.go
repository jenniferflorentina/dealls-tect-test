package healthcheck_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"jennifer/dealls-tech-test/internal/features/healthcheck"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
)

var _ = Describe("Handler", func() {
	It("should return build timestamp set by configuration", func() {
		viper.Set("buildTime", "config time")
		handler := healthcheck.New()

		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()

		handler(w, r, nil)

		Expect(w.Result().StatusCode).To(Equal(http.StatusOK))

		var response healthcheck.Response
		Expect(json.NewDecoder(w.Result().Body).Decode(&response)).To(Succeed())

		expectation := healthcheck.Response{
			BuildTime: "config time",
		}

		Expect(response).To(Equal(expectation))
	})
})
