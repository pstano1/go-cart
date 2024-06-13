// Package pkg provides models & provider implementations for the application
// this file contains errors definitions
package pkg

import "errors"

var (
	ErrUnsupportedLang              = errors.New("ERROR-UNSUPPORTED-LANG")
	ErrRetrievingUsers              = errors.New("ERROR-RETRIEVING-USERS")
	ErrUserNotFound                 = errors.New("ERROR-USER-NOT-FOUND")
	ErrUserAlreadyExists            = errors.New("ERROR-USER-EXISTS")
	ErrCreatingUser                 = errors.New("ERROR-CREATING-USER")
	ErrUnableToReadPayload          = errors.New("ERROR-UNABLE-TO-READ-PAYLOAD")
	ErrUserUnauthorized             = errors.New("ERROR-USER-UNAUTHORIZED")
	ErrUserForbidden                = errors.New("ERROR-FORBIDDEN")
	ErrIncorrectImplementation      = errors.New("ERROR-INCORRECT-IMPLEMENTATION")
	ErrInvalidToken                 = errors.New("ERROR-INVALID-TOKEN")
	ErrUpdatingUser                 = errors.New("ERROR-UPDATING-USER")
	ErrCreatingProduct              = errors.New("ERROR-CREATING-PRODUCT")
	ErrInvalidDescriptionKeyOrValue = errors.New("ERROR-INVALID-PRODUCT-DESCRIPTION")
	ErrInvalidPriceKeyOrValue       = errors.New("ERROR-INVALID-PRODUCT-PRICE")
	ErrProductNotFound              = errors.New("ERROR-PRODUCT-NOT-FOUND")
	ErrUpdatingProduct              = errors.New("ERROR-UPDATING-PRODUCT")
	ErrCreatingCategory             = errors.New("ERROR-CREATING-CATEGORY")
	ErrCategoryNotFound             = errors.New("ERROR-CATEGORY-NOT-FOUND")
	ErrUpdatingCategory             = errors.New("ERROR-UPDATING-CATEGORY")
	ErrCouponNotFound               = errors.New("ERROR-COUPON-NOT-FOUND")
	ErrUpdatingCoupon               = errors.New("ERROR-UPDATING-COUPON")
	ErrOrderNotFound                = errors.New("ERROR-ORDER-NOT-FOUND")
	ErrUpdatingOrder                = errors.New("ERROR-UPDATING-ORDER")
	ErrCreatingOrder                = errors.New("ERROR-CREATING-ORDER")
	ErrInvalidBasketValue           = errors.New("ERROR-INVALID-BASKET-VALUE")
	ErrInvalidNameKeyOrValue        = errors.New("ERROR-INVALID-NAME")
	ErrPassowordsDontMatch          = errors.New("ERROR-PASSWORDS-DONT-MATCH")

	// errors not send to the end user
	ErrStatusNotOK              = errors.New("received status different than 200")
	ErrBaseCurrencyNotAvailable = errors.New("currency requested to be exchanged from is not available in respository")
	ErrCurrencyNotAvailable     = errors.New("currency requested to be exchanged to is not available in respository")
)
