package campaign

import "github.com/stretchr/testify/mock"

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Create(campaign *Campaign) (*Campaign, error) {
	args := r.Called(campaign)
	return args.Get(0).(*Campaign), args.Error(1)
}

func (r *RepositoryMock) Get() []Campaign {
	args := r.Called()
	return args.Get(0).([]Campaign)
}

func (r *RepositoryMock) GetById(id string) (*Campaign, error) {
	args := r.Called(id)
	return args.Get(0).(*Campaign), args.Error(1)
}

func (r *RepositoryMock) Update(campaign *Campaign) (*Campaign, error) {
	args := r.Called(campaign)
	return args.Get(0).(*Campaign), args.Error(1)
}

func (r *RepositoryMock) Delete(id string) error {
	args := r.Called(id)
	return args.Error(0)
}