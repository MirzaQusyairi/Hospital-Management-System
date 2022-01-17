package admins

import (
	middleware "Hospital-Management-System/app/middlewares"
	"Hospital-Management-System/business"
	"Hospital-Management-System/helpers/encrypt"
	"time"
)

type serviceAdmin struct {
	adminRepository Repository
	contextTimeout  time.Duration
	jwtAuth         *middleware.ConfigJWT
}

func NewServiceAdmin(repoAdmin Repository, timeout time.Duration, jwtauth *middleware.ConfigJWT) Service {
	return &serviceAdmin{
		adminRepository: repoAdmin,
		contextTimeout:  timeout,
		jwtAuth:         jwtauth,
	}
}

func (serv *serviceAdmin) Register(domain *Domain) (Domain, error) {

	hashedPassword, err := encrypt.HashingPassword(domain.Password)

	if err != nil {
		return Domain{}, business.ErrInternalServer
	}

	domain.Password = hashedPassword

	result, err := serv.adminRepository.Register(domain)

	if result == (Domain{}) {
		return Domain{}, business.ErrDuplicateData
	}

	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	return result, nil
}

func (serv *serviceAdmin) Login(username, password string) (Domain, error) {

	result, err := serv.adminRepository.Login(username, password)

	if err != nil {
		return Domain{}, business.ErrEmailorPass
	}

	checkPass := encrypt.CheckPasswordHash(password, result.Password)

	if !checkPass {
		return Domain{}, business.ErrEmailorPass
	}

	result.Token = serv.jwtAuth.GenerateToken(result.ID, "admin")

	return result, nil
}
