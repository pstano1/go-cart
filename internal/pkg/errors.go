package pkg

import "errors"

var (
	ErrRetrievingUsers              = errors.New("error while retrieving users")
	ErrUserNotFound                 = errors.New("user not found in the repository")
	ErrUserAlreadyExists            = errors.New("user already exists in the repository")
	ErrCreatingUser                 = errors.New("error while creating user")
	ErrUnableToReadPayload          = errors.New("error while unmarshaling payload")
	ErrUserUnauthorized             = errors.New("user unauthorized")
	ErrUserForbidden                = errors.New("insufficient permissions")
	ErrIncorrectImplementation      = errors.New("internal error, incorrect implementation of populate method")
	ErrInvalidToken                 = errors.New("invalid access token")
	ErrUpdatingUser                 = errors.New("error while updating user")
	ErrCreatingProduct              = errors.New("error while creating product")
	ErrInvalidDescriptionKeyOrValue = errors.New("invalid key or value provided for description")
	ErrInvalidPriceKeyOrValue       = errors.New("invalid key or value provided for price")
	ErrProductNotFound              = errors.New("product not found in the repository")
	ErrUpdatingProduct              = errors.New("could not update product")
	ErrCreatingCategory             = errors.New("could not create category")
	ErrCategoryNotFound             = errors.New("could not find category")
	ErrUpdatingCategory             = errors.New("could not update category")
	ErrCouponNotFound               = errors.New("could not find coupon")
	ErrUpdatingCoupon               = errors.New("could not update coupon")
)
