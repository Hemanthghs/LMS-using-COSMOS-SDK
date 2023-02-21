package keeper

import (
	"fmt"
	"testing"

	"github.com/Leave-Management-System/lms-cosmos/x/lms/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc           codec.Codec
	ctx           sdk.Context
	legacyMsgSrvr v1beta1.MsgServer
}

func (suite *KeeperTestSuite) SetupSuite() {
	suite.reset()
}

func (suite *KeeperTestSuite) reset() {
}

func (suite *KeeperTestSuite) TestSetStudent(t *testing.T) {
	fmt.Println("Tests stared...")
	student := types.Student{
		Address: "123",
		Name:    "Hemanth",
		Id:      "1",
	}
	Keeper.SetStudent(Keeper{}, suite.ctx, student)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
