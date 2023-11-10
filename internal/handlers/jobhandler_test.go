package handlers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"project/internal/middlewear"
	"project/internal/model"
	"project/internal/repository"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
	"gopkg.in/go-playground/assert.v1"
)

func Test_handler_companyCreation(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		// {
		// 	name: "input validation",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`{"company_name":"names",
		// 		"company_adress":    "name@gmail.com",
		// 		"domain": "hfhhfhfh"}`))
		// 		ctx := httpRequest.Context()
		// 		ctx = context.WithValue(ctx, middlewear.TraceIdKey, "1")
		// 		httpRequest = httpRequest.WithContext(ctx)
		// 		c.Request = httpRequest

		// 		return c, rr, nil
		// 	},
		// 	expectedStatusCode: http.StatusUnauthorized,
		// 	expectedResponse:   `{"error":"Unauthorized"}`,
		// },
		// {name: "input validation failure",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`{internal server error}`))
		// 		ctx := httpRequest.Context()
		// 		ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
		// 		httpRequest = httpRequest.WithContext(ctx)
		// 		c.Request = httpRequest

		// 		return c, rr, nil
		// 	},
		// 	expectedStatusCode: http.StatusInternalServerError,
		// 	expectedResponse:   `{"msg":"Internal Server Error"}`,
		// },
		{
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{"company_name":"names",
				"company_adress":    "name@gmail.com",
				"domain": "hfhhfhfh"}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "1")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := repository.NewMockCompanyService(mc)

				ms.EXPECT().CompanyCreate(gomock.Any()).Return(model.Company{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"company_name":"","company_adress":"","domain":""}`,
		},
		{
			name: "failure case",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{"company_name":"names",
				"company_adress":    "name@gmail.com",
				"domain": "hfhhfhfh"}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := repository.NewMockCompanyService(mc)

				ms.EXPECT().CompanyCreate(gomock.Any()).Return(model.Company{}, errors.New("errors")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"compoany creation failed"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				cs: ms,
			}
			h.companyCreation(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_getAllCompany(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		}, {
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := repository.NewMockCompanyService(mc)

				ms.EXPECT().GetAllCompanies().Return([]model.Company{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   "[]",
		},
		{
			name: "fail",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := repository.NewMockCompanyService(mc)

				ms.EXPECT().GetAllCompanies().Return(nil, errors.New("error")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"msg":" could not get all companies"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				cs: ms,
			}
			h.getAllCompany(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_getCompanyById(t *testing.T) {

	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "error while fetching companies from service",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "abc"})
				// mc := gomock.NewController(t)
				// ms := repository.NewMockCompanyService(mc)

				// ms.EXPECT().GetCompanyById(gomock.Any()).Return(model.Company{}, errors.New("test service error")).AnyTimes()

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "error while fetching companies from service==================",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "1"})
				mc := gomock.NewController(t)
				ms := repository.NewMockCompanyService(mc)

				ms.EXPECT().GetCompanyById(gomock.Any()).Return(model.Company{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		// {
		// 	name: "fail",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
		// 		ctx := httpRequest.Context()
		// 		ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")

		// 		httpRequest = httpRequest.WithContext(ctx)
		// 		c.Request = httpRequest
		// 		c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "2"})
		// 		mc := gomock.NewController(t)
		// 		ms := repository.NewMockCompanyService(mc)

		// 		ms.EXPECT().GetCompanyById(gomock.Any()).Return(model.Company{}, nil).AnyTimes()

		// 		return c, rr, ms
		// 	},
		// 	expectedStatusCode: http.StatusBadRequest,
		// 	expectedResponse:   `{"error":"Bad Request"}`,
		// },
		{
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "1"})
				mc := gomock.NewController(t)
				ms := repository.NewMockCompanyService(mc)

				ms.EXPECT().GetCompanyById(gomock.Any()).Return(model.Company{}, errors.New("test service error")).AnyTimes()

				return c, rr, nil
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				cs: ms,
			}
			h.getCompanyById(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_postJobByCompany(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{name: "traceid missing from context",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://tests.com", nil)
				c.Request = httpRequest
				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		// // {
		// // 	name: "missing jwt claims",
		// // 	setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {

		// // 		rr := httptest.NewRecorder()
		// // 		c, _ := gin.CreateTestContext(rr)
		// // 		httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
		// // 		ctx := httpRequest.Context()
		// // 		ctx = context.WithValue(ctx, middlewear.TraceIdKey, "1")
		// // 		httpRequest = httpRequest.WithContext(ctx)
		// // 		c.Request = httpRequest

		// // 		return c, rr, nil
		// // 	},
		// // 	expectedStatusCode: http.StatusUnauthorized,
		// // 	expectedResponse:   `{"error":"Unauthorized"}`,
		// // },
		// {
		// 	name: "id invalid",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
		// 		ctx := httpRequest.Context()
		// 		ctx = context.WithValue(ctx, middlewear.TraceIdKey, "1")
		// 		httpRequest = httpRequest.WithContext(ctx)
		// 		c.Params = append(c.Params, gin.Param{Key: "cid", Value: "abc"})
		// 		c.Request = httpRequest
		// 		mc := gomock.NewController(t)
		// 		ms := repository.NewMockCompanyService(mc)
		// 		return c, rr, ms
		// 	},
		// 	expectedStatusCode: http.StatusBadRequest,
		// 	expectedResponse:   `{"error":"Bad Request"}`,
		// },
		{name: "request validation failure",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://tests.com", strings.NewReader(`{
				"sal": "85000,
				"minNp": 3,
				"maxNp": 60,
				"budget": 85000.00,
				"jobDesc": "We are hiring a software engineer...",
				"minExp": 2.5,
				"maxExp": 5.5,
				"locationIDs": [1,2],
				"skillIDs": [1],
				"workModeIDs": [1],
				"qualificationIDs": [1],
				"shiftIDs": [1],
				"jobTypeIDs": [1]}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "1")
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "123"})
				c.Request = httpRequest
				mc := gomock.NewController(t)
				ms := repository.NewMockCompanyService(mc)
				return c, rr, ms
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{name: "failure success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://tests.com", strings.NewReader(`{
					            "jobTitle": "asdfghj",
								"sal": "85000",
				 				"minNp": 3,
				 				"maxNp": 60,
								"budget": 85000,
				 				"jobDesc": "We are hiring a software engineer...",
				 				"minExp": 2.5,
								"maxExp": 5.5,
				 				"locationIDs": [1,2],
								"skillIDs": [1],
				 				"workModeIDs": [1],
				 				"qualificationIDs": [1],
				 				"shiftIDs": [1],
				 				"jobTypeIDs": [1]}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "1")
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "cid", Value: "1"})
				c.Request = httpRequest
				mc := gomock.NewController(t)
				ms := repository.NewMockCompanyService(mc)
				// ms.EXPECT().JobCreate(gomock.Any()).Return(model.Response{}, errors.New("error in adding job")).AnyTimes()
				return c, rr, ms
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"error in adding job"}`,
		},
		// {name: "success success",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UserService) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpRequest, _ := http.NewRequest(http.MethodGet, "http://tests.com", strings.NewReader(`{
		// 			            "jobTitle": "asdfghj",
		// 						"sal": "85000",
		// 		 				"minNp": 3,
		// 		 				"maxNp": 60,
		// 						"budget": 85000.0,
		// 		 				"jobDesc": "We are hiring a software engineer...",
		// 		 				"minExp": 2.5,
		// 						"maxExp": 5.5,
		// 		 				"locationIDs": [1,2],
		// 						"skillIDs": [1],
		// 		 				"workModeIDs": [1],
		// 		 				"qualificationIDs": [1],
		// 		 				"shiftIDs": [1],
		// 		 				"jobTypeIDs": [1]}`))
		// 		ctx := httpRequest.Context()
		// 		ctx = context.WithValue(ctx, middlewares.TraceIdKey, "1")
		// 		ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
		// 		httpRequest = httpRequest.WithContext(ctx)
		// 		c.Params = append(c.Params, gin.Param{Key: "cid", Value: "1"})
		// 		c.Request = httpRequest
		// 		mc := gomock.NewController(t)
		// 		ms := services.NewMockUserService(mc)
		// 		ms.EXPECT().AddJobDetails(gomock.Any(), gomock.Any(), gomock.Any()).Return(models.Response{}, nil).AnyTimes()
		// 		return c, rr, ms
		// 	},
		// 	expectedStatusCode: http.StatusOK,
		// 	expectedResponse:   `{"ID":0}`,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{cs: ms}
			h.postJobByCompany(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())

		})
	}
}

func Test_handler_getJobByCompany(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		//  {
		// 	name: "error while fetching jobs from service",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
		// 		ctx := httpRequest.Context()
		// 		ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
		// 		httpRequest = httpRequest.WithContext(ctx)
		// 		c.Request = httpRequest
		// 		c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "abc"})
		// 		mc := gomock.NewController(t)
		// 		ms := repository.NewMockCompanyService(mc)

		// 		ms.EXPECT().GetCompanyById(gomock.Any()).Return(model.Job{}, errors.New("test service error")).AnyTimes()

		// 		return c, rr, ms
		// 	},
		// 	expectedStatusCode: http.StatusBadRequest,
		// 	expectedResponse:   `{"error":"Bad Request"}`,
		// },
		{
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")

				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "2"})
				mc := gomock.NewController(t)
				ms := repository.NewMockCompanyService(mc)

				ms.EXPECT().GetJobsByCompanyId(gomock.Any()).Return([]model.Job{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `[]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				cs: ms,
			}
			h.getJobByCompany(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_getAllJob(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		}, {
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := repository.NewMockCompanyService(mc)

				ms.EXPECT().GetAllJobs().Return([]model.Job{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   "[]",
		},
		{
			name: "fail",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, repository.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := repository.NewMockCompanyService(mc)

				ms.EXPECT().GetAllJobs().Return(nil, errors.New("error")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				cs: ms,
			}
			h.getAllJob(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}
