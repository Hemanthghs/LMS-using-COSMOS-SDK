package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrAdminNameNil       = sdkerrors.Register(ModuleName, 1, "Admin name cannot be empty")
	ErrAdminAddressNil    = sdkerrors.Register(ModuleName, 2, "Admin address cannot be empty")
	ErrAdminNotRegistered = sdkerrors.Register(ModuleName, 3, "Admin not yet registered")
)
