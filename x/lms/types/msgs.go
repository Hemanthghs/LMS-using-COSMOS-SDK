package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeRegisterAdmin = "register-admin"
	TypeAddStudent    = "add-student"
	TypeApplyLeave    = "apply-leave"
	TypeAcceptLeave   = "accept-leave"
)

var (
	_ sdk.Msg = &RegisterAdminRequest{}
	_ sdk.Msg = &AddStudentRequest{}
	_ sdk.Msg = &ApplyLeaveRequest{}
	_ sdk.Msg = &AcceptLeaveRequest{}
)

// Register Admin
func NewRegisterAdminReq(accAddr sdk.AccAddress, name string) *RegisterAdminRequest {
	return &RegisterAdminRequest{
		Address: accAddr.String(),
	}
}

func (msg RegisterAdminRequest) GetSignBytes() []byte {
	b := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(b)
}

func (msg RegisterAdminRequest) GetSigners() []sdk.AccAddress {
	valAddr, _ := sdk.AccAddressFromBech32(msg.Address) //ValAddressFromBech32(msg.AdminAddress)
	return []sdk.AccAddress{sdk.AccAddress(valAddr)}
}

func (msg RegisterAdminRequest) Route() string {
	return RouterKey
}
func (msg RegisterAdminRequest) Type() string {
	return TypeRegisterAdmin
}

func (msg RegisterAdminRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("account input address: %s", err)
	}
	return nil
}

// Add Student
func NewAddStudentReq(accountAddr sdk.AccAddress, students []*Student) *AddStudentRequest {
	return &AddStudentRequest{
		Admin: accountAddr.String(),
	}
}

func (msg AddStudentRequest) GetSignBytes() []byte {
	b := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(b)
}

func (msg AddStudentRequest) Route() string {
	return RouterKey
}

func (msg AddStudentRequest) Type() string {
	return TypeAddStudent
}

func (msg AddStudentRequest) GetSigners() []sdk.AccAddress {
	valAddr, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{sdk.AccAddress(valAddr)}
}

func (msg AddStudentRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("account input address: %s", err)
	}
	return nil
}

//Apply Leave

func NewApplyLeaveReq(accountAddr sdk.AccAddress, reason string, from *time.Time, to *time.Time) *ApplyLeaveRequest {
	return &ApplyLeaveRequest{
		Address: accountAddr.String(),
	}
}

func (msg ApplyLeaveRequest) Route() string {
	return RouterKey
}

func (msg ApplyLeaveRequest) Type() string {
	return TypeApplyLeave
}

func (msg ApplyLeaveRequest) GetSignBytes() []byte {
	b := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(b)
}

func (msg ApplyLeaveRequest) GetSigners() []sdk.AccAddress {
	valAddr, _ := sdk.AccAddressFromBech32(msg.Address)
	return []sdk.AccAddress{sdk.AccAddress(valAddr)}
}

func (msg ApplyLeaveRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("account input address: %s", err)
	}
	return nil
}

// Accept Leave

func NewAcceptLeaveReq(accountAddr sdk.AccAddress, LeaveId string, Status string) *AcceptLeaveRequest {
	return &AcceptLeaveRequest{
		Admin: accountAddr.String(),
	}
}

func (msg AcceptLeaveRequest) Route() string {
	return RouterKey
}

func (msg AcceptLeaveRequest) Type() string {
	return TypeAcceptLeave
}

func (msg AcceptLeaveRequest) GetSignBytes() []byte {
	b := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(b)
}

func (msg AcceptLeaveRequest) GetSigners() []sdk.AccAddress {
	valAddr, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{sdk.AccAddress(valAddr)}
}

func (msg AcceptLeaveRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("account input address: %s", err)
	}
	return nil
}
