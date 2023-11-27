package service

import (
	"errors"
	"job-portal-api/internal/model"
	"job-portal-api/internal/repository"
	"reflect"
	"testing"

	gomock "go.uber.org/mock/gomock"
)

func TestService_AddingCompany(t *testing.T) {
	type args struct {
		company model.AddCompany
	}
	tests := []struct {
		name             string
		args             args
		want             model.Company
		wantErr          bool
		mockUserResponse func() (model.Company, error)
	}{
		{
			name:    "failure",
			args:    args{company: model.AddCompany{}},
			want:    model.Company{},
			wantErr: true,
			mockUserResponse: func() (model.Company, error) {
				return model.Company{}, errors.New("error")
			},
		},
		{
			name:    "success",
			args:    args{company: model.AddCompany{CompanyName: "wertyu", Address: "qwertyui", Domain: "wertyui"}},
			want:    model.Company{CompanyName: "wertyu", Address: "qwertyui", Domain: "wertyui"},
			wantErr: false,
			mockUserResponse: func() (model.Company, error) {
				return model.Company{CompanyName: "wertyu", Address: "qwertyui", Domain: "wertyui"}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			ms := repository.NewMockComapnyRepo(mc)
			s, _ := NewCompanyService(ms)
			if tt.mockUserResponse != nil {
				ms.EXPECT().CreateComapny(gomock.Any()).Return(tt.mockUserResponse()).AnyTimes()
			}
			got, err := s.AddingCompany(tt.args.company)
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

func TestService_ViewAllCompanies(t *testing.T) {
	tests := []struct {
		name             string
		want             []model.Company
		wantErr          bool
		mockUserResponse func() ([]model.Company, error)
	}{
		{
			name:    "failure",
			want:    nil,
			wantErr: true,
			mockUserResponse: func() ([]model.Company, error) {
				return nil, errors.New("error")
			},
		},
		{
			name:    "success",
			want:    []model.Company{},
			wantErr: false,
			mockUserResponse: func() ([]model.Company, error) {
				return []model.Company{}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			ms := repository.NewMockComapnyRepo(mc)
			s, _ := NewCompanyService(ms)
			if tt.mockUserResponse != nil {
				ms.EXPECT().GetAllCompanies().Return(tt.mockUserResponse()).AnyTimes()
			}
			got, err := s.ViewAllCompanies()
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

func TestService_ViewCompanyById(t *testing.T) {
	type args struct {
		cId uint64
	}
	tests := []struct {
		name             string
		args             args
		want             model.Company
		wantErr          bool
		mockUserResponse func() (model.Company, error)
	}{
		{
			name:    "failure",
			args:    args{cId: 0},
			want:    model.Company{},
			wantErr: true,
			mockUserResponse: func() (model.Company, error) {
				return model.Company{}, errors.New("error")
			},
		},
		{
			name:    "success",
			args:    args{cId: 1},
			want:    model.Company{},
			wantErr: false,
			mockUserResponse: func() (model.Company, error) {
				return model.Company{}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			ms := repository.NewMockComapnyRepo(mc)
			s, _ := NewCompanyService(ms)
			if tt.mockUserResponse != nil {
				ms.EXPECT().GetCompanyByID(gomock.Any()).Return(tt.mockUserResponse()).AnyTimes()
			}
			got, err := s.ViewCompanyById(tt.args.cId)
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
