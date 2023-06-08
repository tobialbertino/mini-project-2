package usecase

import (
	"database/sql"
	mocks "miniProject2/mocks/modules/customer/repository"
	"miniProject2/modules/customer/model/domain"
	"miniProject2/modules/customer/model/entity"
	"miniProject2/modules/customer/repository"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

// var mockRepo = mocks.CustomerRepository{}

func TestNewCustomerUseCase(t *testing.T) {
	type args struct {
		CustomerRepo repository.CustomerRepository
		DB           *sql.DB
	}
	tests := []struct {
		name string
		args args
		want CustomertUseCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomerUseCase(tt.args.CustomerRepo, tt.args.DB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomerUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO: Mock sql.DB, because start transaction
func TestCustomerUseCaseImpl_CreateCustomer(t *testing.T) {
	// var DB sql.DB // TODO mock DB...
	// var mockRepo = mocks.NewCustomerRepository(t)

	type args struct {
		dt domain.Customer
	}
	tests := []struct {
		name    string
		uc      *CustomerUseCaseImpl
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.CreateCustomer(tt.args.dt)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerUseCaseImpl.CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CustomerUseCaseImpl.CreateCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerUseCaseImpl_DeleteCustomerByID(t *testing.T) {
	type args struct {
		dt domain.Customer
	}
	tests := []struct {
		name    string
		uc      *CustomerUseCaseImpl
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.DeleteCustomerByID(tt.args.dt)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerUseCaseImpl.DeleteCustomerByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CustomerUseCaseImpl.DeleteCustomerByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerUseCaseImpl_GetAllCustomer(t *testing.T) {
	var mockRepo = mocks.NewCustomerRepository(t)
	mockRepo.EXPECT().GetAllCustomer(mock.Anything, mock.Anything, mock.Anything).
		Return([]entity.Customer{}, nil).
		Once()
	mockRepo.EXPECT().Pagination(mock.Anything, mock.Anything).
		Return(entity.Pagination{}, nil).
		Once()

	type args struct {
		dt   domain.Customer
		pagi domain.Pagination
	}
	tests := []struct {
		name    string
		uc      *CustomerUseCaseImpl
		args    args
		want    domain.ListActorWithPaging
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			uc: &CustomerUseCaseImpl{
				CustomerRepository: mockRepo,
				DB:                 nil,
			},
			args: args{
				dt: domain.Customer{},
				pagi: domain.Pagination{
					Page: 1,
				},
			},
			want: domain.ListActorWithPaging{
				Pagination: domain.Pagination{
					Page:    1,
					PerPage: 6,
				},
				Customers: []domain.Customer{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetAllCustomer(tt.args.dt, tt.args.pagi)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerUseCaseImpl.GetAllCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerUseCaseImpl.GetAllCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerUseCaseImpl_GetCustomerByID(t *testing.T) {
	type args struct {
		dt domain.Customer
	}
	tests := []struct {
		name    string
		uc      *CustomerUseCaseImpl
		args    args
		want    domain.Customer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetCustomerByID(tt.args.dt)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerUseCaseImpl.GetCustomerByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerUseCaseImpl.GetCustomerByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerUseCaseImpl_UpdateCustomerByID(t *testing.T) {
	type args struct {
		dt domain.Customer
	}
	tests := []struct {
		name    string
		uc      *CustomerUseCaseImpl
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.UpdateCustomerByID(tt.args.dt)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerUseCaseImpl.UpdateCustomerByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CustomerUseCaseImpl.UpdateCustomerByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
