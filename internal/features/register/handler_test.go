package register_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"jennifer/dealls-tech-test/internal/domain/models"
	"jennifer/dealls-tech-test/internal/domain/others"
	"jennifer/dealls-tech-test/internal/features/register"
	"jennifer/dealls-tech-test/internal/server"
	"jennifer/dealls-tech-test/internal/server/middlewares"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("Handler", func() {
	var db *gorm.DB

	BeforeEach(func() {
		db, _ = server.Conn(context.Background())
		Expect(models.AutoMigrate(db)).To(Succeed())
	})

	Context("Register", func() {
		middlewares.SetupLogger()

		It("should be exception empty username", func() {
			handler := register.New(db)
			r := httptest.NewRequest(http.MethodPost, "/authentication/register", nil)
			requestBody := register.Request{
				Username: "",
				Password: "test234",
				Name:     "John Doe",
				Phone:    "023123124131",
				Email:    "john.doe@gmail.com",
			}
			requestJSON, _ := json.Marshal(requestBody)
			r.Body = io.NopCloser(bytes.NewReader(requestJSON))
			w := httptest.NewRecorder()

			handler(w, r, nil)
			Expect(w.Result().StatusCode).To(Equal(http.StatusUnprocessableEntity))

			var response others.BaseResponse
			Expect(json.NewDecoder(w.Result().Body).Decode(&response)).To(Succeed())

			expected := others.BaseResponse{
				Message:    "Request Payload Validation Not Satisfied",
				StatusCode: http.StatusUnprocessableEntity,
			}

			Expect(response.Message).To(Equal(expected.Message))
			Expect(response.StatusCode).To(Equal(expected.StatusCode))
		})

		It("should be exception username too long", func() {
			handler := register.New(db)
			r := httptest.NewRequest(http.MethodPost, "/authentication/register", nil)
			requestBody := register.Request{
				Username: "PwlQOkUAwfNBkIgZrEpJ QrGUmePUVgZANyTfLoxD nxVbmhaWhofdOqxHUncR kXRZIqHqAnBVIiuLDVIw GHfjTcyTiHzpLndoEoaN",
				Password: "test234",
				Name:     "John Doe",
				Phone:    "023123124131",
				Email:    "john.doe@gmail.com",
			}
			requestJSON, _ := json.Marshal(requestBody)
			r.Body = io.NopCloser(bytes.NewReader(requestJSON))
			w := httptest.NewRecorder()

			handler(w, r, nil)
			Expect(w.Result().StatusCode).To(Equal(http.StatusUnprocessableEntity))

			var response others.BaseResponse
			Expect(json.NewDecoder(w.Result().Body).Decode(&response)).To(Succeed())

			expected := others.BaseResponse{
				Message:    "Request Payload Validation Not Satisfied",
				StatusCode: http.StatusUnprocessableEntity,
			}

			Expect(response.Message).To(Equal(expected.Message))
			Expect(response.StatusCode).To(Equal(expected.StatusCode))
		})

		It("should be able to create user", func() {
			handler := register.New(db)
			r := httptest.NewRequest(http.MethodPost, "/authentication/register", nil)
			requestBody := register.Request{
				Username: "john.doe",
				Password: "test234",
				Name:     "John Doe",
				Phone:    "023123124131",
				Email:    "john.doe@gmail.com",
			}
			requestJSON, _ := json.Marshal(requestBody)
			r.Body = io.NopCloser(bytes.NewReader(requestJSON))
			w := httptest.NewRecorder()

			handler(w, r, nil)
			var resultData others.BaseResponse
			Expect(json.NewDecoder(w.Result().Body).Decode(&resultData)).To(Succeed())
			result := resultData.Data.(map[string]interface{})
			id, _ := uuid.Parse(result["id"].(string))
			expect := register.Response{
				ID:       id.String(),
				Username: "john.doe",
				Name:     "John Doe",
				Phone:    "023123124131",
				Email:    "john.doe@gmail.com",
			}
			Expect(w.Result().StatusCode).To(Equal(http.StatusCreated))
			Expect(result["id"]).To(Equal(expect.ID))
			Expect(result["username"]).To(Equal(expect.Username))
			Expect(result["name"]).To(Equal(expect.Name))
			Expect(result["phone"]).To(Equal(expect.Phone))
			Expect(result["email"]).To(Equal(expect.Email))
		})
	})
})
