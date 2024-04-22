package api

import (
	"github.com/bcmills/unsafeslice"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/pstano1/go-cart/internal/pkg"
	"go.uber.org/zap"
)

func (a *InstanceAPI) CreateUser(request *pkg.UserCreate) (*string, error) {
	a.log.Debug("creating account",
		zap.String("username", request.Username),
	)
	var user pkg.User
	err := copier.Copy(&user, request)
	if err != nil {
		a.log.Error("error while copying request",
			zap.Error(err),
		)
		return nil, err
	}
	user.Password = request.PasswordCheck
	users, err := a.dbController.GetUsers(&pkg.UserFilter{
		Username: user.Username,
	})
	if err != nil {
		a.log.Error("error while retreving users",
			zap.Error(err),
		)
		return nil, pkg.ErrRetrievingUsers
	}
	if len(users) != 0 {
		return nil, pkg.ErrUserAlreadyExists
	}
	hashPassword, err := getHash(unsafeslice.OfString(user.Password))
	if err != nil {
		a.log.Error("error while hashing password",
			zap.Error(err),
		)
		return nil, err
	}
	user.Id = uuid.New().String()
	user.Password = hashPassword
	err = a.dbController.Create(&user)
	if err != nil {
		a.log.Error("error while creating account",
			zap.Error(err),
		)
		return nil, pkg.ErrCreatingUser
	}
	return &user.Id, nil
}
