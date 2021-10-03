package business

import "errors"

var (
	ErrInternalServer = errors.New("something gone wrong, contact administrator")

	ErrNotFound = errors.New("data not found")

	ErrAPIFound = errors.New("movie from API not found")

	ErrNotProFound = errors.New("you are not premium")

	ErrSubsNotFound = errors.New("subscription id not found")

	ErrPayNotFound = errors.New("payment id not found")

	// ErrTransNotFound = errors.New("transaction not found")

	ErrMovieResource = errors.New("movie id not found")

	ErrUserResource = errors.New("user id not found")

	// ErrRatingResource = errors.New("rating not found")

	// ErrCategoryNotFound = errors.New("genre not found")

	ErrDuplicateData = errors.New("duplicate data")

	ErrFillData = errors.New("please fill in the required data")

	ErrUsernamePasswordNotFound = errors.New("check Your email and password")
)
