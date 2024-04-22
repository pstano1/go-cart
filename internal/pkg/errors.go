package pkg

import "errors"

var (
	ErrRetrievingUsers         = errors.New("error while retrieving users")
	ErrUserNotFound            = errors.New("user not found in the repository")
	ErrUserAlreadyExists       = errors.New("user already exists in the repository")
	ErrCreatingUser            = errors.New("error while creating user")
	ErrUnableToReadPayload     = errors.New("error while unmarshaling payload")
	ErrUserUnauthorized        = errors.New("user unauthorized")
	ErrUserForbidden           = errors.New("insufficient permissions")
	ErrIncorrectImplementation = errors.New("internal error, incorrect implementation of populate method")
	ErrInvalidToken            = errors.New("invalid access token")
	ErrUpdatingUser            = errors.New("error while updating user")
)
