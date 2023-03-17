package service

import (
	"airbnb/feature/homestay"
	"airbnb/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAdd(t *testing.T) {
	repo := mocks.NewRepositoryInterface(t)
	srv := NewRoom(repo)
	inputData := homestay.HomestayEntity{
		Homestay_Id: 0,
		Name:        "villa 1",
		Price:       500000,
		Deskripsi:   "indah",
	}
	resData := homestay.HomestayEntity{
		Homestay_Id: 1,
		Name:        "villa 1",
		Price:       500000,
		Deskripsi:   "indah",
	}

	t.Run("success post homestay", func(t *testing.T) {
		repo.On("Add", uint(1), inputData).Return(resData, nil).Once()

		err := srv.Create(inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.Homestay_Id, inputData.Homestay_Id)
		repo.AssertExpectations(t)
	})

	t.Run("eror", func(t *testing.T) {
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		repo.AssertExpectations(t)
	})

	t.Run("cannot post homestay", func(t *testing.T) {
		repo.On("Add", uint(1), mock.Anything).Return(homestay.HomestayEntity{}, errors.New("server error"))

		err := srv.Create(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), resData.Homestay_Id)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

}

func TestUpdate(t *testing.T) {
	repo := mocks.NewRepositoryInterface(t)
	srv := NewRoom(repo)
	inputData := homestay.HomestayEntity{
		Homestay_Id: 0,
		Name:        "villa 1",
		Price:       500000,
		Deskripsi:   "indah",
	}
	resData := homestay.HomestayEntity{
		Homestay_Id: 1,
		Name:        "villa 1",
		Price:       500000,
		Deskripsi:   "indah",
	}

	t.Run("success update homestay", func(t *testing.T) {
		repo.On("Update", uint(1), uint(1), inputData).Return(resData, nil).Once()
		err := srv.Update(1, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.Homestay_Id, resData.Homestay_Id)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("Update", uint(1), uint(1), inputData).Return(homestay.HomestayEntity{}, errors.New("data not found")).Once()
		err := srv.Update(0, inputData)
		assert.Equal(t, uint(0), resData.Homestay_Id)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("Update", uint(1), uint(1), inputData).Return(homestay.HomestayEntity{}, errors.New("server problem")).Once()

		err := srv.Update(1, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.NotEqual(t, 0, resData.Homestay_Id)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewRepositoryInterface(t)
	srv := NewRoom(repo)

	resData := homestay.HomestayEntity{
		Homestay_Id: 1,
		Name:        "villa 1",
		Price:       500000,
		Deskripsi:   "indah",
	}
	t.Run("success delete homestay", func(t *testing.T) {
		repo.On("Delete", uint(1), resData).Return(nil).Once()
		err := srv.Delete(1)
		assert.Nil(t, err)
		repo.AssertExpectations(t)

	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("Delete", uint(2), uint(2)).Return(errors.New("data not found")).Once()

		err := srv.Delete(2)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		repo.AssertExpectations(t)
	})
}

func TestGet(t *testing.T) {
	repo := mocks.NewRepositoryInterface(t)
	srv := NewRoom(repo)
	resData := homestay.HomestayEntity{
		Homestay_Id: 1,
		Name:        "villa 1",
		Price:       500000,
		Deskripsi:   "indah",
	}

	t.Run("success get homestay", func(t *testing.T) {
		repo.On("ShowDetail", uint(1)).Return(resData, nil).Once()
		res, err := srv.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, resData.Homestay_Id, res)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("ShowDetail", uint(1)).Return(homestay.HomestayEntity{}, errors.New("data not found")).Once()
		res, err := srv.GetAll()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.NotEqual(t, 0, res)
		repo.AssertExpectations(t)
	})

}

func TestShow(t *testing.T) {
	repo := mocks.NewRepositoryInterface(t)
	srv := NewRoom(repo)
	resData := homestay.HomestayEntity{
		Homestay_Id: 1,
		Name:        "villa 1",
		Price:       500000,
		Deskripsi:   "indah",
	}

	t.Run("success get homestay", func(t *testing.T) {
		repo.On("Show").Return(resData, nil).Once()
		res, err := srv.GetUserHome(1)
		assert.Nil(t, err)
		assert.Equal(t, resData, res)
		repo.AssertExpectations(t)
	})

	t.Run("homestay not found", func(t *testing.T) {
		repo.On("Show").Return([]homestay.HomestayEntity{}, errors.New("homestay not found")).Once()
		res, err := srv.GetUserHome(1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("ShowAll").Return([]homestay.HomestayEntity{}, errors.New("server problem")).Once()
		res, err := srv.GetUserHome(1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})
}
