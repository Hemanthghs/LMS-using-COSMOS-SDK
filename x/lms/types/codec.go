package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&AddStudentRequest{}, "cosmos-lms/AddStudent", nil)
	cdc.RegisterConcrete(&RegisterAdminRequest{}, "cosmos-lms/RegisterAdmin", nil)
	cdc.RegisterConcrete(&ApplyLeaveRequest{}, "cosmos-lms/ApplyLeave", nil)
	cdc.RegisterConcrete(&AcceptLeaveRequest{}, "cosmos-lms/AcceptLeave", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&RegisterAdminRequest{},
		&AddStudentRequest{},
		&ApplyLeaveRequest{},
		&AcceptLeaveRequest{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)
}
