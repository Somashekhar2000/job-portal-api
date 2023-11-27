package service

import (
	"errors"
	"job-portal-api/internal/cache"
	"job-portal-api/internal/model"
	"job-portal-api/internal/repository"
	"reflect"
	"testing"

	gomock "go.uber.org/mock/gomock"
)

func TestService_CreateJobByCompanyId(t *testing.T) {
	type args struct {
		jobDetails model.NewJobs
		cID        uint
	}
	tests := []struct {
		name         string
		args         args
		want         model.Response
		wantErr      bool
		mockResponse func() (model.Response, error)
	}{
		{
			name:    "failure",
			args:    args{jobDetails: model.NewJobs{}, cID: 0},
			want:    model.Response{},
			wantErr: true,
			mockResponse: func() (model.Response, error) {
				return model.Response{}, errors.New("error")
			},
		},
		{
			name: "success",
			args: args{jobDetails: model.NewJobs{
				Jobname:         "asdfghj",
				MinNoticePeriod: 2,
				MaxNoticePeriod: 60,
				Location:        []uint{1, 2},
				TechnologyStack: []uint{1, 2},
				Description:     "ASDFGHJKL",
				MinExperience:   1,
				MaxExperience:   2,
				Qualifications:  []uint{1, 2},
				Shift:           []uint{1, 2},
				Jobtype:         []uint{1, 2},
			}, cID: 1},
			want:    model.Response{},
			wantErr: false,
			mockResponse: func() (model.Response, error) {
				return model.Response{}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mj := repository.NewMockJobRepository(mc)
			mca := cache.NewMockCaching(mc)
			s, _ := NewJobService(mj, mca)
			if tt.mockResponse != nil {
				mj.EXPECT().CreateJob(gomock.Any()).Return(tt.mockResponse()).AnyTimes()
			}
			got, err := s.CreateJobByCompanyId(tt.args.jobDetails, tt.args.cID)
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

func TestService_ViewJobByCompanyID(t *testing.T) {
	type args struct {
		cID uint
	}
	tests := []struct {
		name         string
		args         args
		want         []model.Job
		wantErr      bool
		mockResponse func() ([]model.Job, error)
	}{
		{
			name:    "failure - 1",
			args:    args{cID: 0},
			want:    nil,
			wantErr: true,
			mockResponse: func() ([]model.Job, error) {
				return nil, errors.New("error")
			},
		},
		{
			name:    "success",
			args:    args{cID: 0},
			want:    []model.Job{},
			wantErr: false,
			mockResponse: func() ([]model.Job, error) {
				return []model.Job{}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mj := repository.NewMockJobRepository(mc)
			mca := cache.NewMockCaching(mc)
			s, _ := NewJobService(mj, mca)
			if tt.mockResponse != nil {
				mj.EXPECT().GetJobByCompanyID(gomock.Any()).Return(tt.mockResponse()).AnyTimes()
			}
			got, err := s.ViewJobByCompanyID(tt.args.cID)
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

func TestService_ViewJobByJobID(t *testing.T) {
	type args struct {
		jID uint
	}
	tests := []struct {
		name         string
		args         args
		want         model.Job
		wantErr      bool
		mockResponse func() (model.Job, error)
	}{
		{
			name:    "failure",
			args:    args{jID: 0},
			want:    model.Job{},
			wantErr: true,
			mockResponse: func() (model.Job, error) {
				return model.Job{}, errors.New("error")
			},
		},
		{
			name:    "success",
			args:    args{jID: 0},
			want:    model.Job{},
			wantErr: false,
			mockResponse: func() (model.Job, error) {
				return model.Job{}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mj := repository.NewMockJobRepository(mc)
			mca := cache.NewMockCaching(mc)
			s, _ := NewJobService(mj, mca)
			if tt.mockResponse != nil {
				mj.EXPECT().GetJobByJobID(gomock.Any()).Return(tt.mockResponse()).AnyTimes()
			}
			got, err := s.ViewJobByJobID(tt.args.jID)
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

func TestService_ViewAllJobs(t *testing.T) {
	tests := []struct {
		name         string
		want         []model.Job
		wantErr      bool
		mockResponse func() ([]model.Job, error)
	}{
		{
			name:    "failure",
			want:    nil,
			wantErr: true,
			mockResponse: func() ([]model.Job, error) {
				return nil, errors.New("error")
			},
		},
		{
			name:    "success",
			want:    []model.Job{},
			wantErr: false,
			mockResponse: func() ([]model.Job, error) {
				return []model.Job{}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mj := repository.NewMockJobRepository(mc)
			mca := cache.NewMockCaching(mc)
			s, _ := NewJobService(mj, mca)
			if tt.mockResponse != nil {
				mj.EXPECT().GetAllJobs().Return(tt.mockResponse()).AnyTimes()
			}
			got, err := s.ViewAllJobs()
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
