package usecase

import (
	"database/sql"
	mocks "miniProject2/mocks/modules/account/repository"
	"miniProject2/modules/account/model/domain"
	"miniProject2/modules/account/model/entity"
	"miniProject2/modules/account/repository"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/mock"
)

func TestNewAccountUseCase(t *testing.T) {
	type args struct {
		AccountRepository repository.AccountRepository
		DB                *sql.DB
	}
	tests := []struct {
		name string
		args args
		want AccountUseCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccountUseCase(tt.args.AccountRepository, tt.args.DB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccountUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountUseCaseImpl_GetAllAdmin(t *testing.T) {

	// mock
	mockRepo := mocks.NewAccountRepository(t)
	mockRepo.EXPECT().
		GetAllAdmin(mock.Anything, mock.AnythingOfType("entity.Actor"), mock.AnythingOfType("entity.Pagination")).
		Return([]entity.Actor{}, nil).
		Once()
	mockRepo.EXPECT().
		Pagination(mock.Anything, mock.AnythingOfType("entity.Pagination")).
		Return(entity.Pagination{}, nil).
		Once()

	type args struct {
		req  domain.Actor
		pagi domain.Pagination
	}
	tests := []struct {
		name    string
		uc      *AccountUseCaseImpl
		args    args
		want    domain.ListActorWithPaging
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			uc: &AccountUseCaseImpl{
				AccountRepository: mockRepo,
				DB:                nil,
			},
			args: args{
				req: domain.Actor{},
				pagi: domain.Pagination{
					Page: 1,
				},
			},
			want: domain.ListActorWithPaging{
				Pagination: domain.Pagination{
					Page:    1,
					PerPage: 6,
				},
				Admins: []domain.Actor{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetAllAdmin(tt.args.req, tt.args.pagi)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountUseCaseImpl.GetAllAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountUseCaseImpl.GetAllAdmin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountUseCaseImpl_DeleteAdminByID(t *testing.T) {
	// mock DB...
	db, sql, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sql.ExpectBegin()
	sql.ExpectCommit()

	//mockRepo
	mockRepo := mocks.NewAccountRepository(t)
	mockRepo.EXPECT().
		DeleteAdminByID(mock.Anything, mock.AnythingOfType("entity.Actor")).
		Return(1, nil).
		Once()

	type args struct {
		req domain.Actor
	}
	tests := []struct {
		name    string
		uc      *AccountUseCaseImpl
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			uc: &AccountUseCaseImpl{
				AccountRepository: mockRepo,
				DB:                db,
			},
			args: args{
				req: domain.Actor{},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.DeleteAdminByID(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountUseCaseImpl.DeleteAdminByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AccountUseCaseImpl.DeleteAdminByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountUseCaseImpl_UpdateAdminStatusByID(t *testing.T) {
	// mock DB...
	db, sql, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sql.ExpectBegin()
	sql.ExpectCommit()

	//mockRepo
	mockRepo := mocks.NewAccountRepository(t)
	mockRepo.EXPECT().
		UpdateAdminRegStatusByAdminID(mock.Anything, mock.AnythingOfType("entity.AdminReg")).
		Return(1, nil).
		Once()

	mockRepo.EXPECT().
		UpdateAdminStatusByAdminID(mock.Anything, mock.AnythingOfType("entity.Actor")).
		Return(1, nil).
		Once()

	type args struct {
		reqReg   domain.AdminReg
		reqActor domain.Actor
	}
	tests := []struct {
		name    string
		uc      *AccountUseCaseImpl
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			uc: &AccountUseCaseImpl{
				AccountRepository: mockRepo,
				DB:                db,
			},
			args: args{
				reqReg:   domain.AdminReg{},
				reqActor: domain.Actor{},
			},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.UpdateAdminStatusByID(tt.args.reqReg, tt.args.reqActor)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountUseCaseImpl.UpdateAdminStatusByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AccountUseCaseImpl.UpdateAdminStatusByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountUseCaseImpl_GetAllApprovalAdmin(t *testing.T) {
	// mock DB...
	db, sql, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sql.ExpectBegin()
	sql.ExpectCommit()

	//mockRepo
	mockRepo := mocks.NewAccountRepository(t)
	mockRepo.EXPECT().
		GetAllApprovalAdmin(mock.Anything).
		Return([]entity.AdminReg{}, nil).
		Once()

	tests := []struct {
		name    string
		uc      *AccountUseCaseImpl
		want    []domain.AdminReg
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			uc: &AccountUseCaseImpl{
				AccountRepository: mockRepo,
				DB:                db,
			},
			want:    []domain.AdminReg{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetAllApprovalAdmin()
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountUseCaseImpl.GetAllApprovalAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountUseCaseImpl.GetAllApprovalAdmin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountUseCaseImpl_VerifyActorCredential(t *testing.T) {
	// mock DB...
	db, sql, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sql.ExpectBegin()
	sql.ExpectCommit()

	//mockRepo
	mockRepo := mocks.NewAccountRepository(t)
	mockRepo.EXPECT().
		VerifyActorCredential(mock.Anything, mock.AnythingOfType("entity.Actor")).
		Return(entity.Actor{}, nil).
		Once()

	type args struct {
		req domain.Actor
	}
	tests := []struct {
		name    string
		uc      *AccountUseCaseImpl
		args    args
		want    domain.ResToken
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Fail",
			uc: &AccountUseCaseImpl{
				AccountRepository: mockRepo,
				DB:                db,
			},
			args: args{
				req: domain.Actor{},
			},
			want:    domain.ResToken{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.VerifyActorCredential(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountUseCaseImpl.VerifyActorCredential() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountUseCaseImpl.VerifyActorCredential() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountUseCaseImpl_AddActor(t *testing.T) {
	// mock DB...
	db, sql, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sql.ExpectBegin()
	sql.ExpectCommit()

	//mockRepo
	mockRepo := mocks.NewAccountRepository(t)
	mockRepo.EXPECT().
		AddActor(mock.Anything, mock.AnythingOfType("entity.Actor")).
		Return(1, nil).
		Once()
	mockRepo.EXPECT().
		RegisterAdmin(mock.Anything, mock.AnythingOfType("entity.AdminReg")).
		Return(1, nil).
		Once()

	type args struct {
		req domain.Actor
	}
	tests := []struct {
		name    string
		uc      *AccountUseCaseImpl
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			uc: &AccountUseCaseImpl{
				AccountRepository: mockRepo,
				DB:                db,
			},
			args: args{
				req: domain.Actor{},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.AddActor(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountUseCaseImpl.AddActor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AccountUseCaseImpl.AddActor() = %v, want %v", got, tt.want)
			}
		})
	}
}
