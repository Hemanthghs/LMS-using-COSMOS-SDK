package keeper_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

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

///////////////////// Register Admin Tests ////////////////////////

func (s *TestSuite) TestRegisterAdmin() {
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
				Name:    "Sai",
				Address: sdk.AccAddress("sakjhfdd").String(),
			},
			expected: "Admin Registered Successfully",
		},
		{
			arg1: types.RegisterAdminRequest{
				Name:    "Vishal",
				Address: "",
			},
			expected: "Address cannot be empty",
		},
		{
			arg1: types.RegisterAdminRequest{
				Name:    "",
				Address: sdk.AccAddress("kgjdk").String(),
			},
			expected: "Name cannot be empty",
		},
	}

	require := s.Require()
	for _, test := range registerAdminTests {

		if output := s.stdntKeeper.RegisterAdmin(s.ctx, &test.arg1); output != test.expected {
			require.Equal(test.expected, output)
		}
		// s.stdntKeeper.GetAdmin(s.ctx, sdk.AccAddress("sakjhfdd").String())
	}

}

///////////////////// Add Student Tests ////////////////////////

func (s *TestSuite) TestAddStudent() {
	students := []*types.Student{
		{
			Address: sdk.AccAddress("lms1").String(),
			Name:    "hemanth1",
			Id:      "1",
		},
		{
			Address: sdk.AccAddress("lms2").String(),
			Name:    "hemanth2",
			Id:      "2",
		},
		{
			Address: sdk.AccAddress("lms3").String(),
			Name:    "hemanth3",
			Id:      "3",
		},
	}
	req := types.AddStudentRequest{
		Admin:    "Hemanthsai",
		Students: students,
	}
	s.stdntKeeper.AddStudent(s.ctx, &req)
}

//////////////////// Apply Leave Tests ////////////////////////

func (s *TestSuite) TestApplyLeave() {
	type applyLeaveTest struct {
		arg1     types.ApplyLeaveRequest
		expected string
	}
	dateString := "2006-Jan-02"
	fromDate, _ := time.Parse(dateString, "2023-Feb-22")
	toDate, _ := time.Parse(dateString, "2023-Feb-26")
	var applyLeaveTests = []applyLeaveTest{
		{
			arg1: types.ApplyLeaveRequest{
				Address: sdk.AccAddress("lms1").String(),
				Reason:  "I am feeling sick",
				From:    &fromDate,
				To:      &toDate,
			},
			expected: "Leave Applied Successfully",
		},
		{
			arg1: types.ApplyLeaveRequest{
				Address: sdk.AccAddress("lms2").String(),
				Reason:  "I have to attend a wedding",
				From:    &fromDate,
				To:      &toDate,
			},
			expected: "Leave Applied Successfully",
		},
	}
	require := s.Require()
	for _, test := range applyLeaveTests {
		if output := s.stdntKeeper.ApplyLeave(s.ctx, &test.arg1); output != test.expected {
			require.Equal(test.expected, output)
		}
	}

}

// ///////////////// Accept Leave Tests ////////////////////////
func (s *TestSuite) TestAcceptLeave() {
	req := types.AcceptLeaveRequest{
		Admin:   sdk.AccAddress("abcdef").String(),
		LeaveId: "1",
		Status:  types.LeaveStatus_STATUS_ACCEPTED,
	}
	res := s.stdntKeeper.AcceptLeave(s.ctx, &req)
	fmt.Println(res)

}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
