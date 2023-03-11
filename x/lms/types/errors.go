package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrAdminNameNil       = sdkerrors.Register(ModuleName, 1, "Admin name cannot be empty")
	ErrAdminAddressNil    = sdkerrors.Register(ModuleName, 2, "Admin address cannot be empty")
	ErrAdminNotRegistered = sdkerrors.Register(ModuleName, 3, "Admin not yet registered")
	ErrLeaveNotFound      = sdkerrors.Register(ModuleName, 4, "Leave does not exist with the given id")
	ErrStudentNotFound    = sdkerrors.Register(ModuleName, 5, "Student not found with the given id")
)
