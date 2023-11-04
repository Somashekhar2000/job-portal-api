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
		name string
		//s       *Service
		args             args
		want             model.User
		wantErr          bool
		mockRepoResponse func() (model.User, error)
	}{
		{
			name: "success",
			want: model.User{UserName: "sridath", Email: "dath@gmail.com"},
			args: args{
				nu: model.UserSignup{UserName: "sridath", Email: "dath@gmail.com", Password: "bangalore"},
			},
			wantErr: false,
			mockRepoResponse: func() (model.User, error) {
				return model.User{UserName: "sridath", Email: "dath@gmail.com"}, nil
			},
		},
		{
			name: "failure",
			want: model.User{},
			args: args{
				nu: model.UserSignup{UserName: "", Email: "dath@gmail.com", Password: "bangalore"},
			},
			wantErr: true,
			mockRepoResponse: func() (model.User, error) {
				return model.User{}, errors.New("test error")
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
				t.Errorf("Service.UserSignup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.UserSignup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Userlogin(t *testing.T) {
	type args struct {
		l model.UserLogin
	}
	tests := []struct {
		name string
		//s       *Service
		args             args
		want             jwt.RegisteredClaims
		wantErr          bool
		mockRepoResponse func() (model.User, error)
	}{
		{name: "checking  sucess case",
			args:    args{l: model.UserLogin{Email: "name@gmail.com", Password: "hfhbhfrbfrbfrwbfbfbrfbrhfhfh"}},
			want:    jwt.RegisteredClaims{Issuer: "service project", Subject: "0", Audience: jwt.ClaimStrings{"users"}, ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)), IssuedAt: jwt.NewNumericDate(time.Now())},
			wantErr: false,
			mockRepoResponse: func() (model.User, error) {
				return model.User{UserName: "sridath", Email: "dath@gmail.com", PasswordHash: "$2a$10$dy5br0fE1KHYarImvJZhcu1VkGy2s/OGjL9cwQzPAPflCvgNeE8VG"}, nil
			},
		},
		{name: "checking  failure case",
			args:    args{l: model.UserLogin{Email: "", Password: "hfhhfhfh"}},
			want:    jwt.RegisteredClaims{},
			wantErr: true,
			mockRepoResponse: func() (model.User, error) {
				return model.User{}, errors.New("test error")
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
				t.Errorf("Service.Userlogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Userlogin() = %v, want %v", got, tt.want)
			}
		})
	}
}
