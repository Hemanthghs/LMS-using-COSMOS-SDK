package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&AddStudentRequest{}, "cosmos-lms/AddStudent", nil)
	cdc.RegisterConcrete(&RegisterAdminRequest{}, "cosmos-lms/RegisterAdmin", nil)
	cdc.RegisterConcrete(&ApplyLeaveRequest{}, "cosmos-lms/ApplyLeave", nil)
	cdc.RegisterConcrete(&AcceptLeaveRequest{}, "cosmos-lsm/AcceptLeave", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&RegisterAdminRequest{},
		&AddStudentRequest{},
		&ApplyLeaveRequest{},
		&AcceptLeaveRequest{},
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
}
