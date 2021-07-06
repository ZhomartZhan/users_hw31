package users_hw31

import "errors"

var (
	ErrNoUser = errors.New("User not found by this username and password")
)
