package keeper

import (
	//"cosmossdk.io/store/prefix"

	"log"

	"github.com/Leave-Management-System/lms-cosmos/x/lms/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) RegisterAdmin(ctx sdk.Context, registerAdmin *types.RegisterAdminRequest) string {
	if registerAdmin.Name == "" {
		return "Name cannot be empty"
	} else if registerAdmin.Address == "" {
		return "Address cannot be empty"
	} else {
		store := ctx.KVStore(k.storeKey)
		k.cdc.MustMarshal(registerAdmin)
		marshalRegisterAdmin, err := k.cdc.Marshal(registerAdmin)
		if err != nil {
			log.Fatal(err)
		}
		store.Set(types.AdminStoreKey(registerAdmin.Address), marshalRegisterAdmin)
		return "Admin Registered Successfully"
	}
}
