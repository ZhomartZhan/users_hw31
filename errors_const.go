package users_hw31

import "errors"

var (
	ErrNoUser = errors.New("user not found by this username and password")
)
