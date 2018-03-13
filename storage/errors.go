package storage

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/errors"
	)


var (
	errIpfsServiceError = fmt.Errorf("ipfs not avaliable")
	errFileNotExists    = fmt.Errorf("file is not stored in IPFS")

	invalidInput = fmt.Errorf("you need to add more parameters")
)

func ErrIPFSOffline() error {
	return errors.WithCode(errIpfsServiceError, errors.CodeTypeInternalErr)
}

func ErrFileNotExists() error {
	return errors.WithCode(errFileNotExists, errors.CodeTypeBaseInvalidInput)
}
func ErrMissingInput() error {
	return errors.WithCode(invalidInput,errors.CodeTypeBaseInvalidInput)
}