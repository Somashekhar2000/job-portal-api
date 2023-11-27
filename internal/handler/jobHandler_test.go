package handler

import (
	"context"
	"errors"
	"job-portal-api/internal/authentication"
	"job-portal-api/internal/middleware"
	"job-portal-api/internal/model"
	"job-portal-api/internal/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/mock/gomock"
	"gopkg.in/go-playground/assert.v1"
)

func TestHandler_CreateJobByCompanyID(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.JobService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				hr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(hr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, hr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error ":"Internal Server Error"}`,
		},
		{
			name: "missing jwt claims",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")

				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   `{"error ":"Unauthorized"}`,
		},
		{
			name: "invalid id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(
					`{
						"jobName": "asdfghj",
						"minNoticePeriod": 1,
						"maxNoticePeriod": 50,
						"location": [1,2],
						"technologyStack": [1, 2],
						"description": "Exciting job opportunity for a software Developer...",
						"minExperience": 1,
						"maxExperience": 6,
						"qualifications": [1, 2],
						"shifts": [1,2],
						"jobtype": [1,2]
					  }`,
				))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
				ctx = context.WithValue(ctx, authentication.AuthKey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "abc"})
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error ":"Bad Request"}`,
		},
		{
			name: "error in decoding",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(
					`{
						jobName": "Software Developer",
						"minNoticePeriod": 1,
						"maxNoticePeriod": 50,
						"location": [1,2],
						"technologyStack": [1, 2],
						"description": "Exciting job opportunity for a software Developer...",
						"minExperience": 1,
						"maxExperience": 6,
						"qualifications": [1, 2],
						"shifts": [1,2],
						"jobtype": [1,2]
					  }`,
				))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
				ctx = context.WithValue(ctx, authentication.AuthKey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error ":"Bad Request"}`,
		},
		{
			name: "error in validating",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(
					`{
						"jobName": "",
						"minNoticePeriod": 1,
						"maxNoticePeriod": 50,
						"location": [1,2],
						"technologyStack": [1, 2],
						"description": "Exciting job opportunity for a software Developer...",
						"minExperience": 1,
						"maxExperience": 6,
						"qualifications": [1, 2],
						"shifts": [1,2],
						"jobtype": [1,2]
					  }`,
				))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
				ctx = context.WithValue(ctx, authentication.AuthKey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error ":"Bad Request"}`,
		},
		{
			name: "invalid id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(
					`{
						"jobName": "asdfghj",
						"minNoticePeriod": 1,
						"maxNoticePeriod": 50,
						"location": [1,2],
						"technologyStack": [1, 2],
						"description": "Exciting job opportunity for a software Developer...",
						"minExperience": 1,
						"maxExperience": 6,
						"qualifications": [1, 2],
						"shifts": [1,2],
						"jobtype": [1,2]
					  }`,
				))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
				ctx = context.WithValue(ctx, authentication.AuthKey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "abc"})
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error ":"Bad Request"}`,
		},
		{
			name: "failure",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(
					`{
						"jobName": "software testing",
						"minNoticePeriod": 1,
						"maxNoticePeriod": 50,
						"location": [1,2],
						"technologyStack": [1, 2],
						"description": "Exciting job opportunity for a software Developer...",
						"minExperience": 1,
						"maxExperience": 6,
						"qualifications": [1, 2],
						"shifts": [1,2],
						"jobtype": [1,2]
					  }`,
				))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
				ctx = context.WithValue(ctx, authentication.AuthKey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpRequest

				mc := gomock.NewController(t)
				mj := service.NewMockJobService(mc)

				mj.EXPECT().CreateJobByCompanyId(gomock.Any(), gomock.Any()).Return(model.Response{}, errors.New("error")).AnyTimes()

				return c, rr, mj
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error ":"Internal Server Error"}`,
		},
		{
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(
					`{
						"jobName": "software testing",
						"minNoticePeriod": 1,
						"maxNoticePeriod": 50,
						"location": [1,2],
						"technologyStack": [1, 2],
						"description": "Exciting job opportunity for a software Developer...",
						"minExperience": 1,
						"maxExperience": 6,
						"qualifications": [1, 2],
						"shifts": [1,2],
						"jobtype": [1,2]
					  }`,
				))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "1")
				ctx = context.WithValue(ctx, authentication.AuthKey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpRequest

				mc := gomock.NewController(t)
				mj := service.NewMockJobService(mc)

				mj.EXPECT().CreateJobByCompanyId(gomock.Any(), gomock.Any()).Return(model.Response{}, nil).AnyTimes()

				return c, rr, mj
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"id":0}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, mj := tt.setup()
			h := Handler{
				serviceJob: mj,
			}
			h.CreateJobByCompanyID(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func TestHandler_ViewJobByCompanyId(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.JobService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				hr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(hr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, hr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "missing jwt claims",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")

				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   `{"error":"Unauthorized"}`,
		},
		{
			name: "invalid company id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
				ctx = context.WithValue(ctx, authentication.AuthKey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "abc"})
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "failure",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
				ctx = context.WithValue(ctx, authentication.AuthKey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpRequest

				mc := gomock.NewController(t)
				mj := service.NewMockJobService(mc)

				mj.EXPECT().ViewJobByCompanyID(gomock.Any()).Return([]model.Job{}, errors.New("error"))

				return c, rr, mj
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
				ctx = context.WithValue(ctx, authentication.AuthKey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpRequest

				mc := gomock.NewController(t)
				mj := service.NewMockJobService(mc)

				mj.EXPECT().ViewJobByCompanyID(gomock.Any()).Return([]model.Job{}, nil)

				return c, rr, mj
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `[]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, mj := tt.setup()
			h := Handler{
				serviceJob: mj,
			}
			h.ViewJobByCompanyId(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func TestHandler_ViewJobByJobID(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.JobService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				hr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(hr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, hr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "missing jwt claims",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")

				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   `{"error":"Unauthorized"}`,
		},
		{
			name: "invalid company id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
				ctx = context.WithValue(ctx, authentication.AuthKey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "abc"})
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "failure",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
				ctx = context.WithValue(ctx, authentication.AuthKey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpRequest

				mc := gomock.NewController(t)
				mj := service.NewMockJobService(mc)

				mj.EXPECT().ViewJobByJobID(gomock.Any()).Return(model.Job{}, errors.New("error"))

				return c, rr, mj
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
				ctx = context.WithValue(ctx, authentication.AuthKey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpRequest

				mc := gomock.NewController(t)
				mj := service.NewMockJobService(mc)

				mj.EXPECT().ViewJobByJobID(gomock.Any()).Return(model.Job{}, nil)

				return c, rr, mj
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `[]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, mj := tt.setup()
			h := Handler{
				serviceJob: mj,
			}
			h.ViewJobByJobID(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func TestHandler_ViewAllJobs(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.JobService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				hr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(hr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, hr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "missing jwt claims",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")

				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   `{"error":"Unauthorized"}`,
		},
		{
			name: "failure",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
				ctx = context.WithValue(ctx, authentication.AuthKey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				mj := service.NewMockJobService(mc)

				mj.EXPECT().ViewAllJobs().Return(nil, errors.New("error"))

				return c, rr, mj
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
				ctx = context.WithValue(ctx, authentication.AuthKey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				mj := service.NewMockJobService(mc)

				mj.EXPECT().ViewAllJobs().Return([]model.Job{}, nil)

				return c, rr, mj
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `[]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, mj := tt.setup()
			h := Handler{
				serviceJob: mj,
			}
			h.ViewAllJobs(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func TestHandler_ProcessJobApplication(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.JobService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				hr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(hr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, hr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "missing jwt claims",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")

				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   `{"error":"Unauthorized"}`,
		},
		{
			name: "error in decoding",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(
					`
										[
						{
							"name": "John Doe 1",
							"age": "30",
							"jid": 7,
							"job_application": {
								"noticePeriod": 100,
								"location": [
									1,
									3
								],
								"technologyStack": [
									1,
									3
								],
								"experience": 90,
								"qualifications": [
									1,
									2
								],
								"shifts": [
									1,
									2
								],
								"jobtype": [
									1,
									2
								]
							}
						},
						{
							"name": "John Doe 2",
							"age": "30",
							"jid": 8,
							"job_application": {
								"noticePeriod": 30,
								"location": [
									1,
									2
								],
								"technologyStack": [
									1,
									2
								],
								"experience": 5,
								"qualifications": [
									1,
									2
								],
								"shifts": [
									1,
									2
								],
								jobtype": [
									1,
									2
								]
							}
						}
					]
					`,
				))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
				ctx = context.WithValue(ctx, authentication.AuthKey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "error in validating",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(
					`
					[
    {
        "name": "",
        "age": "",
        "jid": 7,
        "job_application": {
            "noticePeriod": 100,
            "location": [
                1,
                3
            ],
            "technologyStack": [
                1,
                3
            ],
            "experience": 90,
            "qualifications": [
                1,
                2
            ],
            "shifts": [
                1,
                2
            ],
            "jobtype": [
                1,
                2
            ]
        }
    },
    {
        "name": "",
        "age": "",
        "jid": 8,
        "job_application": {
            "noticePeriod": 30,
            "location": [
                1,
                2
            ],
            "technologyStack": [
                1,
                2
            ],
            "experience": 5,
            "qualifications": [
                1,
                2
            ],
            "shifts": [
                1,
                2
            ],
            "jobtype": [
                1,
                2
            ]
        }
    }
]
					`,
				))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
				ctx = context.WithValue(ctx, authentication.AuthKey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				mj := service.NewMockJobService(mc)

				mj.EXPECT().ProcessApplication(gomock.Any()).Return(nil).AnyTimes()

				return c, rr, mj

			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error all applications rejected ":"Bad Request"}`,
		},
		{
			name: "error in decoding",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(
					`
					[
    {
        "name": "John Doe 1",
        "age": "30",
        "jid": 7,
        "job_application": {
            "noticePeriod": 100,
            "location": [
                1,
                3
            ],
            "technologyStack": [
                1,
                3
            ],
            "experience": 90,
            "qualifications": [
                1,
                2
            ],
            "shifts": [
                1,
                2
            ],
            "jobtype": [
                1,
                2
            ]
        }
    },
    {
        "name": "John Doe 2",
        "age": "30",
        "jid": 8,
        "job_application": {
            "noticePeriod": 30,
            "location": [
                1,
                2
            ],
            "technologyStack": [
                1,
                2
            ],
            "experience": 5,
            "qualifications": [
                1,
                2
            ],
            "shifts": [
                1,
                2
            ],
            "jobtype": [
                1,
                2
            ]
        }
    }
]
					`,
				))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
				ctx = context.WithValue(ctx, authentication.AuthKey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				mj := service.NewMockJobService(mc)

				mj.EXPECT().ProcessApplication(gomock.Any()).Return(nil).AnyTimes()

				return c, rr, mj
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error all applications rejected ":"Bad Request"}`,
		},
		{
			name: "error in decoding",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.JobService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(
					`
					[
    {
        "name": "John Doe 1",
        "age": "30",
        "jid": 7,
        "job_application": {
            "noticePeriod": 100,
            "location": [
                1,
                3
            ],
            "technologyStack": [
                1,
                3
            ],
            "experience": 90,
            "qualifications": [
                1,
                2
            ],
            "shifts": [
                1,
                2
            ],
            "jobtype": [
                1,
                2
            ]
        }
    },
    {
        "name": "John Doe 2",
        "age": "30",
        "jid": 8,
        "job_application": {
            "noticePeriod": 30,
            "location": [
                1,
                2
            ],
            "technologyStack": [
                1,
                2
            ],
            "experience": 5,
            "qualifications": [
                1,
                2
            ],
            "shifts": [
                1,
                2
            ],
            "jobtype": [
                1,
                2
            ]
        }
    }
]
					`,
				))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
				ctx = context.WithValue(ctx, authentication.AuthKey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				mj := service.NewMockJobService(mc)

				mj.EXPECT().ProcessApplication(gomock.Any()).Return([]model.NewUserApplication{}).AnyTimes()

				return c, rr, mj
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `[]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, mj := tt.setup()
			h := Handler{
				serviceJob: mj,
			}
			h.ProcessJobApplication(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}
