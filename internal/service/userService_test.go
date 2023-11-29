package service

import (
	"errors"
	"job-portal-api/internal/authentication"
	"job-portal-api/internal/cache"
	"job-portal-api/internal/model"
	"job-portal-api/internal/repository"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestService_UserSignup(t *testing.T) {
	type args struct {
		userData model.UserSignup
	}
	tests := []struct {
		name             string
		args             args
		want             model.User
		wantErr          bool
		mockUserResponse func() (model.User, error)
	}{
		{
			name:    "failure",
			args:    args{userData: model.UserSignup{UserName: "qwertyui", EmailID: "sdfgh@gmail.com", Password: ""}},
			want:    model.User{},
			wantErr: true,
			mockUserResponse: func() (model.User, error) {
				return model.User{}, errors.New("invalid user")
			},
		},
		{
			name:    "failure",
			args:    args{userData: model.UserSignup{UserName: "qwertyu", EmailID: "wertyui@gmail.com", Password: "12345678"}},
			want:    model.User{},
			wantErr: true,
			mockUserResponse: func() (model.User, error) {
				return model.User{}, errors.New("error")
			},
		},
		{
			name:    "success",
			args:    args{userData: model.UserSignup{UserName: "qwertyu", EmailID: "wertyui@gmail.com", Password: "12345678"}},
			want:    model.User{UserName: "qwertyu", EmailID: "wertyui@gmail.com"},
			wantErr: false,
			mockUserResponse: func() (model.User, error) {
				return model.User{UserName: "qwertyu", EmailID: "wertyui@gmail.com"}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			ms := repository.NewMockUserRepository(mc)
			ma := authentication.NewMockAuthenticaton(mc)
			mca := cache.NewMockCaching(mc)
			s, _ := NewUserService(ms, ma, mca)
			if tt.mockUserResponse != nil {
				ms.EXPECT().CreateUser(gomock.Any()).Return(tt.mockUserResponse()).AnyTimes()
			}
			got, err := s.UserSignup(tt.args.userData)
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
		userSignin model.UserLogin
	}
	tests := []struct {
		name             string
		args             args
		want             string
		wantErr          bool
		mockUserResponse func() (model.User, error)
		mockAuth         func() (string, error)
	}{

		{
			name:    "faile",
			args:    args{userSignin: model.UserLogin{}},
			want:    "",
			wantErr: true,
			mockUserResponse: func() (model.User, error) {
				return model.User{}, errors.New("error")
			},
			mockAuth: func() (string, error) {
				return "", errors.New("error")
			},
		},
		{
			name:    "invalid paasword",
			args:    args{userSignin: model.UserLogin{EmailID: "abc@gmail.com", Password: "12345678"}},
			want:    "",
			wantErr: true,
			mockUserResponse: func() (model.User, error) {
				return model.User{
					EmailID:  "abc@gmail.com",
					Password: "$2a$10$hNkswO/Wr.gDQJPnaYqvoOh0oQSnw8PkNm6Ipj6890CVEYTpNetUPabC",
				}, nil
			},
			mockAuth: func() (string, error) {
				return "", errors.New("error")
			},
		},
		{
			name:    "success",
			args:    args{userSignin: model.UserLogin{EmailID: "abc@gmail.com", Password: "12345678"}},
			want:    "",
			wantErr: false,
			mockUserResponse: func() (model.User, error) {
				return model.User{
					EmailID:  "abc@gmail.com",
					Password: "$2a$10$hNkswO/Wr.gDQJPnaYqvoOh0oQSnw8PkNm6Ipj6CVEYTpNetUPabC",
				}, nil
			},
			mockAuth: func() (string, error) {
				return "", nil
			},
		},
		{
			name:    "success",
			args:    args{userSignin: model.UserLogin{EmailID: "abc@gmail.com", Password: "12345678"}},
			want:    "",
			wantErr: true,
			mockUserResponse: func() (model.User, error) {
				return model.User{
					EmailID:  "abc@gmail.com",
					Password: "$2a$10$hNkswO/Wr.gDQJPnaYqvoOh0oQSnw8PkNm6Ipj6CVEYTpNetUPabC",
				}, nil
			},
			mockAuth: func() (string, error) {
				return "", errors.New("error")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			ms := repository.NewMockUserRepository(mc)
			ma := authentication.NewMockAuthenticaton(mc)
			mca := cache.NewMockCaching(mc)
			s, _ := NewUserService(ms, ma, mca)
			if tt.mockUserResponse != nil {
				ms.EXPECT().CheckUser(gomock.Any()).Return(tt.mockUserResponse()).AnyTimes()
				ma.EXPECT().GenerateToken(gomock.Any()).Return(tt.mockAuth()).AnyTimes()
			}
			got, err := s.Userlogin(tt.args.userSignin)
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

func TestService_OTPGeneration(t *testing.T) {
	type args struct {
		userdetails model.ChangePassword
	}
	tests := []struct {
		name              string
		args              args
		want              string
		wantErr           bool
		mockUserResponse  func() (model.User, error)
		mockCacheResponse func() error
	}{
		{
			name:    "failure - 1",
			args:    args{userdetails: model.ChangePassword{EmailID: "asdfghjk", DOB: "sdfghj"}},
			want:    "",
			wantErr: true,
			mockUserResponse: func() (model.User, error) {
				return model.User{}, errors.New("error")
			},
			mockCacheResponse: func() error {
				return errors.New("error")
			},
		},
		{
			name:    "failure - 2",
			args:    args{userdetails: model.ChangePassword{EmailID: "asdfghjk", DOB: "sdfghj"}},
			want:    "",
			wantErr: true,
			mockUserResponse: func() (model.User, error) {
				return model.User{}, nil
			},
			mockCacheResponse: func() error {
				return errors.New("error")
			},
		},
		{
			name:    "success",
			args:    args{userdetails: model.ChangePassword{EmailID: "asdfghjk", DOB: "sdfghj"}},
			want:    "",
			wantErr: false,
			mockUserResponse: func() (model.User, error) {
				return model.User{}, nil
			},
			mockCacheResponse: func() error {
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			ma := authentication.NewMockAuthenticaton(mc)
			ms := repository.NewMockUserRepository(mc)
			mca := cache.NewMockCaching(mc)
			s, _ := NewUserService(ms, ma, mca)
			if tt.mockUserResponse != nil {
				ms.EXPECT().CheckUser(gomock.Any()).Return(tt.mockUserResponse()).AnyTimes()
				mca.EXPECT().AddOTP(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.mockCacheResponse()).AnyTimes()
			}
			got, err := s.OTPGeneration(tt.args.userdetails)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.OTPGeneration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.OTPGeneration() = %v, want %v", got, tt.want)
			}
		})
	}
}
