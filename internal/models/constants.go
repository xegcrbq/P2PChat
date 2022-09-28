package models

import "gopkg.in/errgo.v2/errors"

var ErrDataNotFound = errors.New("data not found")
var ErrComandNotFound = errors.New("command not found")
var ErrInvalidSC = errors.New("invalid secret key")
var ErrAlreadyExists = errors.New("data already exists")
