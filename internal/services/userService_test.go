package services

import (
	"errors"
	"project/internal/model"
	"project/internal/repository"
	"reflect"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/mock/gomock"
)

func TestService_UserSignup(t *testing.T) {
	type args struct {
		nu model.UserSignup
	}
	tests := []struct {
		name             string
		args             args
		want             model.User
		wantErr          bool
		mockRepoResponse func() (model.User, error)
	}{
		{
			name: "======success case=========",
			want: model.User{UserName: "soma", Email: "soma@gmail.com"},
			args: args{
				nu: model.UserSignup{UserName: "soma", Email: "soma@gmail.com", Password: "12345678"},
			},
			wantErr: false,
			mockRepoResponse: func() (model.User, error) {
				return model.User{UserName: "soma", Email: "soma@gmail.com"}, nil
			},
		},
		{
			name: "=====failure case=====",
			want: model.User{},
			args: args{
				nu: model.UserSignup{UserName: "", Email: "abc@gmail.com", Password: "87654321"},
			},
			wantErr: true,
			mockRepoResponse: func() (model.User, error) {
				return model.User{}, errors.New("error in testing")
			},
		},
		{
			name: "=====failure case=====",
			want: model.User{},
			args: args{
				nu: model.UserSignup{UserName: "", Email: "abc@gmail.com", Password: "11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"},
			},
			wantErr: true,
			mockRepoResponse: func() (model.User, error) {
				return model.User{}, errors.New("error in testing")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockUserRepo := repository.NewMockUsers(mc)
			mockCompanyRepo := repository.NewMockCompany(mc)

			if tt.mockRepoResponse != nil {
				mockUserRepo.EXPECT().CreateUser(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s, _ := NewService(mockUserRepo, mockCompanyRepo)
			got, err := s.UserSignup(tt.args.nu)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.UserSignup() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.UserSignup() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestService_Userlogin(t *testing.T) {
	type args struct {
		l model.UserLogin
	}
	tests := []struct {
		name             string
		args             args
		want             jwt.RegisteredClaims
		wantErr          bool
		mockRepoResponse func() (model.User, error)
	}{
		{name: "======success case=========",
			args:    args{l: model.UserLogin{Email: "soma@gmail.com", Password: "12345678"}},
			want:    jwt.RegisteredClaims{Issuer: "service project", Subject: "0", Audience: jwt.ClaimStrings{"users"}, ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)), IssuedAt: jwt.NewNumericDate(time.Now())},
			wantErr: false,
			mockRepoResponse: func() (model.User, error) {
				return model.User{UserName: "sridath", Email: "soma@gmail.com", PasswordHash: "$2a$10$5yU8yp9SkhrO67PJaTIXNepoIoYopj9lI3aJxkxXoF.VLrT4ARBxa"}, nil
			},
		},
		{name: "=====failure case=====",
			args:    args{l: model.UserLogin{Email: "", Password: "12345678"}},
			want:    jwt.RegisteredClaims{},
			wantErr: true,
			mockRepoResponse: func() (model.User, error) {
				return model.User{}, errors.New("error in testing")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockUserRepo := repository.NewMockUsers(mc)
			mockCompanyRepo := repository.NewMockCompany(mc)

			if tt.mockRepoResponse != nil {
				mockUserRepo.EXPECT().FetchUserByEmail(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s, _ := NewService(mockUserRepo, mockCompanyRepo)
			got, err := s.Userlogin(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Userlogin() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Userlogin() = %v, want = %v", got, tt.want)
			}
		})
	}
}
