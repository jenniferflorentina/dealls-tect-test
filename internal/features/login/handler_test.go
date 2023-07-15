package login_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"jennifer/dealls-tech-test/internal/features/utils"
	"jennifer/dealls-tech-test/internal/server/middlewares"

	"jennifer/dealls-tech-test/internal/domain/models"
	"jennifer/dealls-tech-test/internal/domain/others"
	"jennifer/dealls-tech-test/internal/features/login"
	"jennifer/dealls-tech-test/internal/server"

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

	Context("Login", func() {
		middlewares.SetupLogger()

		userID, _ := uuid.Parse("994175cc-0dbd-434f-9989-f7cced32d466")
		user := models.Users{
			ID:       userID,
			Username: "john.doe",
			Password: utils.HashAndSalt([]byte("test234")),
			Name:     "John Doe",
			Phone:    "023123124131",
			Email:    "john.doe@gmail.com",
		}

		It("Should be exception empty Username", func() {
			handler := login.New(db)
			r := httptest.NewRequest(http.MethodGet, "/authentication/login", nil)
			requestBody := login.Request{
				Username: "",
				Password: "test1234",
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

		It("should be exception username not found", func() {
			db.Create(&user)

			handler := login.New(db)
			r := httptest.NewRequest(http.MethodGet, "/authentication/login", nil)
			requestBody := login.Request{
				Username: "john",
				Password: "test1234",
			}
			requestJSON, _ := json.Marshal(requestBody)
			r.Body = io.NopCloser(bytes.NewReader(requestJSON))
			w := httptest.NewRecorder()

			handler(w, r, nil)
			Expect(w.Result().StatusCode).To(Equal(http.StatusNotFound))

			var response others.BaseResponse
			Expect(json.NewDecoder(w.Result().Body).Decode(&response)).To(Succeed())

			expected := others.BaseResponse{
				Message:    "Username or Password Wrong",
				StatusCode: http.StatusNotFound,
			}

			Expect(response.Message).To(Equal(expected.Message))
			Expect(response.StatusCode).To(Equal(expected.StatusCode))
		})

		It("should be exception wrong password", func() {
			db.Create(&user)

			handler := login.New(db)
			r := httptest.NewRequest(http.MethodGet, "/authentication/login", nil)
			requestBody := login.Request{
				Username: "john.doe",
				Password: "test1234",
			}
			requestJSON, _ := json.Marshal(requestBody)
			r.Body = io.NopCloser(bytes.NewReader(requestJSON))
			w := httptest.NewRecorder()

			handler(w, r, nil)
			Expect(w.Result().StatusCode).To(Equal(http.StatusNotFound))

			var response others.BaseResponse
			Expect(json.NewDecoder(w.Result().Body).Decode(&response)).To(Succeed())

			expected := others.BaseResponse{
				Message:    "Username or Password Wrong",
				StatusCode: http.StatusNotFound,
			}

			Expect(response.Message).To(Equal(expected.Message))
			Expect(response.StatusCode).To(Equal(expected.StatusCode))
		})

		It("Should be able to login", func() {

			jwtTokenReq := others.Claims{
				UserID:   userID,
				Username: "john.doe",
			}
			token, _ := utils.GenerateJwtToken(jwtTokenReq)

			db.Create(&user)

			handler := login.New(db)
			r := httptest.NewRequest(http.MethodGet, "/authentication/login", nil)
			requestBody := login.Request{
				Username: "john.doe",
				Password: "test234",
			}
			requestJSON, _ := json.Marshal(requestBody)
			r.Body = io.NopCloser(bytes.NewReader(requestJSON))
			w := httptest.NewRecorder()

			handler(w, r, nil)

			var resultData others.BaseResponse
			Expect(json.NewDecoder(w.Result().Body).Decode(&resultData)).To(Succeed())

			expect := login.Response{
				UserID:   userID.String(),
				Username: "john.doe",
				Name:     "John Doe",
				Token:    token,
			}
			result := resultData.Data.(map[string]interface{})

			Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
			Expect(result["token"]).To(Equal(expect.Token))
			Expect(result["name"]).To(Equal(expect.Name))
			Expect(result["userId"]).To(Equal(expect.UserID))
		})

	})
})
