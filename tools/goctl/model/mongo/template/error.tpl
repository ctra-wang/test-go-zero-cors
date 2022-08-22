package model

import (
	"errors"

	"github.com/Holyson/test-go-zero-cors/core/stores/mon"
)

var (
	ErrNotFound        = mon.ErrNotFound
	ErrInvalidObjectId = errors.New("invalid objectId")
)
