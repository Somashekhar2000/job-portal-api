// package services

// import (
// 	"errors"
// 	"project/internal/model"
// 	"project/internal/repository"
// 	"reflect"
// 	"testing"

// 	"go.uber.org/mock/gomock"
// )

// func TestService_CompanyCreate(t *testing.T) {
// 	type args struct {
// 		nc model.CreateCompany
// 	}
// 	tests := []struct {
// 		name             string
// 		args             args
// 		want             model.Company
// 		wantErr          bool
// 		mockRepoResponse func() (model.Company, error)
// 	}{
// 		{
// 			name: "=======success case========",
// 			want: model.Company{CompanyName: "TEK", Adress: "BENGALURU", Domain: "IT"},
// 			args: args{
// 				nc: model.CreateCompany{CompanyName: "TEK", Adress: "BENGALURU", Domain: "IT"},
// 			},
// 			wantErr: false,
// 			mockRepoResponse: func() (model.Company, error) {
// 				return model.Company{CompanyName: "TEK", Adress: "BENGALURU", Domain: "IT"}, nil
// 			},
// 		},
// 		{
// 			name: "===========invalid input : failure case==========",
// 			want: model.Company{},
// 			args: args{
// 				nc: model.CreateCompany{CompanyName: "", Adress: "asdfgh", Domain: "engineer"},
// 			},
// 			wantErr: true,
// 			mockRepoResponse: func() (model.Company, error) {
// 				return model.Company{}, errors.New("error in test case")
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockCompanyRepo := repository.NewMockCompany(mc)
// 			mockUserRepo := repository.NewMockUsers(mc)

// 			companyModel := model.Company{CompanyName: tt.args.nc.CompanyName, Adress: tt.args.nc.Adress, Domain: tt.args.nc.Domain}
// 			if tt.mockRepoResponse != nil {
// 				mockCompanyRepo.EXPECT().CreateCompany(companyModel).Return(tt.mockRepoResponse()).AnyTimes()
// 			}
// 			s, _ := NewService(mockUserRepo, mockCompanyRepo)
// 			got, err := s.CompanyCreate(tt.args.nc)

// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.CompanyCreate() error = %v, wantErr = %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.CompanyCreate() = %v, want = %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestService_GetCompany(t *testing.T) {
// 	type args struct {
// 		id int
// 	}
// 	tests := []struct {
// 		name             string
// 		args             args
// 		want             model.Company
// 		wantErr          bool
// 		mockRepoResponse func() (model.Company, error)
// 	}{
// 		{
// 			name: "success",
// 			want: model.Company{CompanyName: "google", Adress: "mysore", Domain: "development"},
// 			args: args{
// 				id: 1,
// 			},
// 			wantErr: false,
// 			mockRepoResponse: func() (model.Company, error) {
// 				return model.Company{CompanyName: "google", Adress: "mysore", Domain: "development"}, nil
// 			},
// 		},
// 		{
// 			name: "============failure case===========",
// 			want: model.Company{},
// 			args: args{
// 				id: 12,
// 			},
// 			wantErr: true,
// 			mockRepoResponse: func() (model.Company, error) {
// 				return model.Company{}, errors.New("ID is not been greater")
// 			},
// 		},
// 		{
// 			name: "============failure case===========",
// 			want: model.Company{},
// 			args: args{
// 				id: 1,
// 			},
// 			wantErr: true,
// 			mockRepoResponse: func() (model.Company, error) {
// 				return model.Company{}, errors.New("ID is not been greater")
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockCompanyRepo := repository.NewMockCompany(mc)
// 			mockUserRepo := repository.NewMockUsers(mc)

// 			if tt.mockRepoResponse != nil {
// 				mockCompanyRepo.EXPECT().GetCompany(tt.args.id).Return(tt.mockRepoResponse()).AnyTimes()
// 			}
// 			s, _ := NewService(mockUserRepo, mockCompanyRepo)
// 			got, err := s.GetCompanyById(tt.args.id)

// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.GetCompany() error = %v, wantErr = %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.GetCompany() = %v, want = %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestService_GetAllCompanies(t *testing.T) {
// 	tests := []struct {
// 		name             string
// 		want             []model.Company
// 		wantErr          bool
// 		mockRepoResponse func() ([]model.Company, error)
// 	}{
// 		{
// 			name:    "=======success case==========",
// 			want:    []model.Company{{CompanyName: "wipro", Adress: "sdfgh", Domain: "sde"}, {CompanyName: "micro", Adress: "dfgh", Domain: "wedfgh"}},
// 			wantErr: false,
// 			mockRepoResponse: func() ([]model.Company, error) {
// 				return []model.Company{{CompanyName: "wipro", Adress: "sdfgh", Domain: "sde"}, {CompanyName: "micro", Adress: "dfgh", Domain: "wedfgh"}}, nil
// 			},
// 		},
// 		{
// 			name:    "============failure case ==========",
// 			want:    nil,
// 			wantErr: true,
// 			mockRepoResponse: func() ([]model.Company, error) {
// 				return []model.Company{}, errors.New("cannot get companies no companies are created")
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			mc := gomock.NewController(t)
// 			mockCompanyRepo := repository.NewMockCompany(mc)
// 			mockUserRepo := repository.NewMockUsers(mc)

// 			if tt.mockRepoResponse != nil {
// 				mockCompanyRepo.EXPECT().GetAllCompany().Return(tt.mockRepoResponse()).AnyTimes()
// 			}
// 			s, _ := NewService(mockUserRepo, mockCompanyRepo)
// 			got, err := s.GetAllCompanies()

// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.GetAllCompanies() error = %v, wantErr = %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.GetAllCompanies() = %v, want = %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestService_GetJobs(t *testing.T) {
// 	type args struct {
// 		id int
// 	}
// 	tests := []struct {
// 		name             string
// 		args             args
// 		want             []model.Job
// 		wantErr          bool
// 		mockRepoResponse func() ([]model.Job, error)
// 	}{
// 		{
// 			name: "========success case======",
// 			want: []model.Job{{JobTitle: "sdfgh", JobSalary: "wedfg", Uid: 1}},
// 			args: args{
// 				id: 1,
// 			},
// 			wantErr: false,
// 			mockRepoResponse: func() ([]model.Job, error) {
// 				return []model.Job{{JobTitle: "sdfgh", JobSalary: "wedfg", Uid: 1}}, nil
// 			},
// 		},
// 		{
// 			name: "===========failure case ==========",
// 			want: nil,
// 			args: args{
// 				id: 12,
// 			},
// 			wantErr: true,
// 			mockRepoResponse: func() ([]model.Job, error) {
// 				return nil, errors.New("id doest not exists")
// 			},
// 		},
// 		{
// 			name: "===========failure case ==========",
// 			want: nil,
// 			args: args{
// 				id: 1,
// 			},
// 			wantErr: true,
// 			mockRepoResponse: func() ([]model.Job, error) {
// 				return nil, errors.New("id doest not exists")
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockCompanyRepo := repository.NewMockCompany(mc)
// 			mockUserRepo := repository.NewMockUsers(mc)

// 			if tt.mockRepoResponse != nil {
// 				mockCompanyRepo.EXPECT().GetJobs(tt.args.id).Return(tt.mockRepoResponse()).AnyTimes()
// 			}
// 			s, _ := NewService(mockUserRepo, mockCompanyRepo)
// 			got, err := s.GetJobsByCompanyId(tt.args.id)

// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.GetJob() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.GetJob() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestService_GetAllJobs(t *testing.T) {
// 	tests := []struct {
// 		name             string
// 		want             []model.Job
// 		wantErr          bool
// 		mockRepoResponse func() ([]model.Job, error)
// 	}{
// 		{
// 			name:    "==========success==========",
// 			want:    []model.Job{{JobTitle: "asdfg", JobSalary: "12345", Uid: 1}, {JobTitle: "swdf", JobSalary: "12345", Uid: 2}},
// 			wantErr: false,
// 			mockRepoResponse: func() ([]model.Job, error) {
// 				return []model.Job{{JobTitle: "asdfg", JobSalary: "12345", Uid: 1}, {JobTitle: "swdf", JobSalary: "12345", Uid: 2}}, nil
// 			},
// 		},
// 		{
// 			name:    "==========failure case=======",
// 			want:    nil,
// 			wantErr: true,
// 			mockRepoResponse: func() ([]model.Job, error) {
// 				return []model.Job{}, errors.New("jobs doesnot exists")
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockCompanyRepo := repository.NewMockCompany(mc)
// 			mockUserRepo := repository.NewMockUsers(mc)

// 			if tt.mockRepoResponse != nil {
// 				mockCompanyRepo.EXPECT().GetAllJobs().Return(tt.mockRepoResponse()).AnyTimes()
// 			}
// 			s, _ := NewService(mockUserRepo, mockCompanyRepo)
// 			got, err := s.GetAllJobs()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.GetAllJobs() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.GetAllJobs() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestService_GetJobByJobId(t *testing.T) {
// 	type args struct {
// 		id int
// 	}
// 	tests := []struct {
// 		name             string
// 		args             args
// 		want             model.Job
// 		wantErr          bool
// 		mockRepoResponse func() (model.Job, error)
// 	}{
// 		{
// 			name:    "invalid id",
// 			args:    args{id: 15},
// 			want:    model.Job{},
// 			wantErr: true,
// 			mockRepoResponse: func() (model.Job, error) {
// 				return model.Job{}, errors.New("ID cannot be greater than 10")
// 			},
// 		},
// 		{
// 			name:    "id does not exists",
// 			args:    args{id: 2},
// 			want:    model.Job{},
// 			wantErr: true,
// 			mockRepoResponse: func() (model.Job, error) {
// 				return model.Job{}, errors.New("Failed to retrieve job from the repository")
// 			},
// 		},
// 		{
// 			name:    "success",
// 			args:    args{id: 1},
// 			want:    model.Job{},
// 			wantErr: false,
// 			mockRepoResponse: func() (model.Job, error) {
// 				return model.Job{}, nil
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockCompanyRepo := repository.NewMockCompany(mc)
// 			mockUserRepo := repository.NewMockUsers(mc)

// 			if tt.mockRepoResponse != nil {
// 				mockCompanyRepo.EXPECT().GetJobsByJobId(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
// 			}
// 			s, _ := NewService(mockUserRepo, mockCompanyRepo)
// 			got, err := s.GetJobByJobId(tt.args.id)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.GetAllJobs() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.GetAllJobs() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestService_JobCreate(t *testing.T) {
// 	type args struct {
// 		newJob model.NewJobRequest
// 		id     uint64
// 	}
// 	tests := []struct {
// 		name             string
// 		args             args
// 		want             model.Response
// 		wantErr          bool
// 		mockRepoResponse func() (model.Response, error)
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockCompanyRepo := repository.NewMockCompany(mc)
// 			mockUserRepo := repository.NewMockUsers(mc)

//				if tt.mockRepoResponse != nil {
//					mockCompanyRepo.EXPECT().PostJob(gomock.Any())(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
//				}
//				s, _ := NewService(mockUserRepo, mockCompanyRepo)
//				got, err := s.GetJobByJobId(tt.args)
//				if (err != nil) != tt.wantErr {
//					t.Errorf("Service.GetAllJobs() error = %v, wantErr %v", err, tt.wantErr)
//					return
//				}
//				if !reflect.DeepEqual(got, tt.want) {
//					t.Errorf("Service.GetAllJobs() = %v, want %v", got, tt.want)
//				}
//			})
//		}
//	}
package services
