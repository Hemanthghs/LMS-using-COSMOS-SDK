package keeper_test

import (
	"sync"
	"testing"

	"github.com/Leave-Management-System/lms-cosmos/x/lms/keeper"
	"github.com/Leave-Management-System/lms-cosmos/x/lms/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	dbm "github.com/tendermint/tm-db"
)

type TestSuite struct {
	suite.Suite
	ctx         sdk.Context
	stdntKeeper keeper.Keeper
	*assert.Assertions
	mu      sync.RWMutex
	require *require.Assertions
	t       *testing.T
}

func (s *TestSuite) SetupTest() {
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

// T retrieves the current *testing.T context.
func (suite *TestSuite) T() *testing.T {
	suite.mu.RLock()
	defer suite.mu.RUnlock()
	return suite.t
}

// SetT sets the current *testing.T context.
func (suite *TestSuite) SetT(t *testing.T) {
	suite.mu.Lock()
	defer suite.mu.Unlock()
	suite.t = t
	suite.Assertions = assert.New(t)
	suite.require = require.New(t)
}

// Require returns a require context for suite.
func (suite *TestSuite) Require() *require.Assertions {
	suite.mu.Lock()
	defer suite.mu.Unlock()
	if suite.require == nil {
		suite.require = require.New(suite.T())
	}
	return suite.require
}

type registerAdminTest struct {
	arg1     types.RegisterAdminRequest
	expected string
}

var registerAdminTests = []registerAdminTest{
	{
		arg1: types.RegisterAdminRequest{
			Name:    "Hemanthsai",
			Address: sdk.AccAddress("abcdef").String(),
		},
		expected: "Admin Registered Successfully",
	},
	{
		arg1: types.RegisterAdminRequest{
			Name:    "Hemanthsai",
			Address: sdk.AccAddress("abcdef").String(),
		},
		expected: "Admin Registered Successfully",
	},
}

func (s *TestSuite) TestRegisterAdmin() {
	// addr := sdk.AccAddress("abcdef")
	// req := types.RegisterAdminRequest{
	// 	Name:    "",
	// 	Address: addr.String(),
	// }
	// res := s.stdntKeeper.RegisterAdmin(s.ctx, &req)
	// fmt.Println(res)
	require := s.Require()
	for _, test := range registerAdminTests {
		if output := s.stdntKeeper.RegisterAdmin(s.ctx, &test.arg1); output != test.expected {
			require.Equal(test.expected, output)
		}
	}

}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
