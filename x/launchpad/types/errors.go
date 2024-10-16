package types

import (
	errorsmod "cosmossdk.io/errors"
	"fmt"
)

var (
	ErrorPoolNotExist = errorsmod.Register(ModuleName, 1, "Pool is not exist")
)

type PoolDoesNotExistError struct {
	PoolId uint64
}

func (e PoolDoesNotExistError) Error() string {
	return fmt.Sprintf("pool with ID %d does not exist", e.PoolId)
}
