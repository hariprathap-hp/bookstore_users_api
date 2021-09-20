package serviceUser

import (
	"github.com/hariprathap-hp/bookstore_users_api/src/domains/domain_user"
	"github.com/hariprathap-hp/bookstore_users_api/src/utils/cryptoUtil"
	"github.com/hariprathap-hp/bookstore_users_api/src/utils/dateUtil"
	"github.com/hariprathap-hp/bookstore_users_api/src/utils/errors"
)

func GetByID(id int64) (*domain_user.User, *errors.RestErr) {
	user := domain_user.User{Id: id}
	getErr := user.Get()
	if getErr != nil {
		return nil, getErr
	}
	return &user, nil
}

func SaveUser(user *domain_user.User) (*domain_user.User, *errors.RestErr) {
	valErr := user.Validate()
	if valErr != nil {
		return nil, valErr
	}
	user.DateCreated = dateUtil.GetTimeDBFormat()
	user.Status = user.GetStatus()
	user.Password = cryptoUtil.GetMD5(user.Password)
	saveErr := user.Save()
	if saveErr != nil {
		return nil, saveErr
	}
	return user, nil
}

func UpdateUser(isPatch bool, user *domain_user.User) (*domain_user.User, *errors.RestErr) {
	current := &domain_user.User{Id: user.Id}
	err := current.Get()
	if err != nil {
		return nil, err
	}
	if isPatch {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}

		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	updateErr := current.Update()
	if updateErr != nil {
		return nil, updateErr
	}
	return current, nil
}
func DeleteByID(id int64) (*domain_user.User, *errors.RestErr) {
	user := domain_user.User{Id: id}
	getErr := user.Delete()
	if getErr != nil {
		return nil, getErr
	}
	return &user, nil
}

func Search(status string) (domain_user.Users, *errors.RestErr) {
	user := domain_user.User{}
	return user.Search(status)
}

func LoginUser(request *domain_user.LoginUser) (*domain_user.User, *errors.RestErr) {
	user := &domain_user.User{
		Email:    request.Email,
		Password: cryptoUtil.GetMD5(request.Password),
	}
	if err := user.Login(); err != nil {
		return nil, err
	}
	return user, nil
}
