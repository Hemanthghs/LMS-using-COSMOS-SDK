package keeper_test

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/Leave-Management-System/lms-cosmos/x/lms/types"
// 	"github.com/cosmos/cosmos-sdk/codec"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
// 	"github.com/stretchr/testify/suite"
// )

// type KeeperTestSuite struct {
// 	suite.Suite

// 	cdc           codec.Codec
// 	ctx           sdk.Context
// 	legacyMsgSrvr v1beta1.MsgServer
// 	lmsKeeper     Keeper
// }

// func (suite *KeeperTestSuite) SetupSuite() {
// 	suite.reset()
// }

// func (suite *KeeperTestSuite) reset() {
// }

// func (suite *KeeperTestSuite) TestSetStudent(t *testing.T) {
// 	fmt.Println("Tests stared...")
// 	student := types.Student{
// 		Address: "123",
// 		Name:    "Hemanth",
// 		Id:      "1",
// 	}
// 	Keeper.SetStudent(Keeper{}, suite.ctx, student)
// }

// func TestKeeperTestSuite(t *testing.T) {
// 	suite.Run(t, new(KeeperTestSuite))
// }

import (
	"fmt"
	"testing"

	"github.com/Leave-Management-System/lms-cosmos/x/lms/keeper"
	"github.com/Leave-Management-System/lms-cosmos/x/lms/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	dbm "github.com/tendermint/tm-db"
)

type TestSuite struct {
	suite.Suite
	ctx         sdk.Context
	stdntKeeper keeper.Keeper
}

func (s *TestSuite) SetupTest() {
	fmt.Println("I am in setup")
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	encCfg := simapp.MakeTestEncodingConfig()
	lmsKey := sdk.NewKVStoreKey(types.StoreKey)
	ctx := testutil.DefaultContext(lmsKey, sdk.NewTransientStoreKey("transient_test"))
	keeper := keeper.NewKeeper(lmsKey, encCfg.Codec)
	cms.MountStoreWithDB(lmsKey, storetypes.StoreTypeIAVL, db)
	s.Require().NoError(cms.LoadLatestVersion())
	s.stdntKeeper = keeper
	s.ctx = ctx
}
func (s *TestSuite) TestRegisterAdmin() {
	addr := sdk.AccAddress("abcdef")
	req := types.RegisterAdminRequest{
		Name:    "Hemanthsai",
		Address: addr.String(),
	}
	s.stdntKeeper.RegisterAdmin(s.ctx, &req)
}
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
