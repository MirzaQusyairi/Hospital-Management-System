package doctors_test

import (
	"Hospital-Management-System/business"
	"Hospital-Management-System/business/doctors"
	_doctorMock "Hospital-Management-System/business/doctors/mocks"
	"Hospital-Management-System/helpers/encrypt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockDoctorRepository _doctorMock.Repository
	doctorService        doctors.Service
	doctorDomain         doctors.Domain
)

func TestMain(m *testing.M) {
	doctorDomain = doctors.Domain{
		ID:           13,
		Username:     "rudi",
		Password:     "rudi123",
		Name:         "Dr Rudi",
		Fullname:     "Dr Rudi Mahruzar",
		Gender:       "Laki Laki",
		Specialist:   "Penyakit Dalam",
		Address:      "Jl. Antariksa No. 11",
		Phone_Number: "0812545445454",
		DOB:          "12-12-1980",
	}

	m.Run()
}

func TestRegister(t *testing.T) {
	t.Run("test case 1, valid test for register", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("doctor")
		outputDomain := doctors.Domain{
			Username: "doctor",
			Password: password,
		}
		inputService := doctors.Domain{
			Username: "doctor",
			Password: "doctor",
		}
		mockDoctorRepository.On("Register", mock.Anything).Return(outputDomain, nil).Once()

		resp, err := doctorService.Register(&inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.Username, resp.Username)
	})

	t.Run("test case 2, invalid test for register wrong password", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("doctor")
		outputDomain := doctors.Domain{
			Username: "doctor",
			Password: password,
		}
		inputService := doctors.Domain{
			Username: "doctor",
			Password: "temon",
		}
		mockDoctorRepository.On("Register", mock.Anything).Return(outputDomain, business.ErrInternalServer).Once()

		resp, err := doctorService.Register(&inputService)
		assert.Equal(t, err, business.ErrInternalServer)
		assert.Empty(t, resp)
	})
}

func TestLogin(t *testing.T) {
	t.Run("test case 1, valid test for login", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("doctor")
		outputDomain := doctors.Domain{
			Username: "doctor",
			Password: password,
		}
		inputService := doctors.Domain{
			Username: "doctor",
			Password: "doctor",
		}

		mockDoctorRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, nil).Once()

		resp, err := doctorService.Login(inputService.Username, inputService.Password)

		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})

	t.Run("test case 2, invalid test for login wrong password", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("doctor")
		outputDomain := doctors.Domain{
			Username: "doctor",
			Password: password,
		}
		inputService := doctors.Domain{
			Username: "doctor",
			Password: "temon",
		}
		mockDoctorRepository.On("Register", mock.Anything).Return(outputDomain, business.ErrInternalServer).Once()

		mockDoctorRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, business.ErrEmailorPass).Once()

		resp, err := doctorService.Login(inputService.Username, inputService.Password)
		assert.Equal(t, err, business.ErrEmailorPass)
		assert.Empty(t, resp)
	})
}

func TestAllDoctor(t *testing.T) {
	t.Run("test case 1, valid all doctor", func(t *testing.T) {
		expectedReturn := []doctors.Domain{
			{
				ID:           13,
				Username:     "rudi",
				Password:     "rudi123",
				Name:         "Dr Rudi",
				Fullname:     "Dr Rudi Mahruzar",
				Gender:       "Laki Laki",
				Specialist:   "Penyakit Dalam",
				Address:      "Jl. Antariksa No. 11",
				Phone_Number: "0812545445454",
				DOB:          "12-12-1980",
			},
			{
				ID:           14,
				Username:     "rudi",
				Password:     "rudi123",
				Name:         "Dr Rudi",
				Fullname:     "Dr Rudi Mahruzar",
				Gender:       "Laki Laki",
				Specialist:   "Penyakit Dalam",
				Address:      "Jl. Antariksa No. 11",
				Phone_Number: "0812545445454",
				DOB:          "12-12-1980",
			},
		}
		mockDoctorRepository.On("AllDoctor").Return(expectedReturn, nil).Once()
		_, err := doctorService.AllDoctor()
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all doctor", func(t *testing.T) {
		expectedReturn := []doctors.Domain{}
		mockDoctorRepository.On("AllDoctor").Return(expectedReturn, assert.AnError).Once()
		_, err := doctorService.AllDoctor()
		assert.Equal(t, err, assert.AnError)

	})
}
