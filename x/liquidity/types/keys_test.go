package types_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/crescent-network/crescent/x/liquidity/types"
)

type keysTestSuite struct {
	suite.Suite
}

func TestKeysTestSuite(t *testing.T) {
	suite.Run(t, new(keysTestSuite))
}

func (s *keysTestSuite) TestGetPairKey() {
	s.Require().Equal([]byte{0xa5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetPairKey(0))
	s.Require().Equal([]byte{0xa5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x9}, types.GetPairKey(9))
	s.Require().Equal([]byte{0xa5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa}, types.GetPairKey(10))
}

func (s *keysTestSuite) TestGetPoolKey() {
	s.Require().Equal([]byte{0xab, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1}, types.GetPoolKey(1))
	s.Require().Equal([]byte{0xab, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5}, types.GetPoolKey(5))
	s.Require().Equal([]byte{0xab, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa}, types.GetPoolKey(10))
}

func (s *keysTestSuite) TestPairIndexKey() {
	testCases := []struct {
		denomA   string
		denomB   string
		expected []byte
	}{
		{
			"denomA",
			"denomB",
			[]byte{0xa6, 0x6, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x41, 0x6, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x42},
		},
		{
			"denomC",
			"denomD",
			[]byte{0xa6, 0x6, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x43, 0x6, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x44},
		},
		{
			"denomE",
			"denomF",
			[]byte{0xa6, 0x6, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x45, 0x6, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x46},
		},
	}

	for _, tc := range testCases {
		key := types.GetPairIndexKey(tc.denomA, tc.denomB)
		s.Require().Equal(tc.expected, key)

		// TODO: parse pair index key?
	}
}

func (suite *keysTestSuite) TestReversePairIndexKey() {
	// TODO: not implemented yet
}

func (suite *keysTestSuite) TestPoolByReserveAccIndexKey() {
	// TODO: not implemented yet
}

func (suite *keysTestSuite) TestPoolsByPairIndexKey() {
	// TODO: not implemented yet
}

func (suite *keysTestSuite) TestDepositRequestKey() {
	// TODO: not implemented yet
}

func (suite *keysTestSuite) TestWithdrawRequestKey() {
	// TODO: not implemented yet
}

func (suite *keysTestSuite) TestSwapRequestKey() {
	// TODO: not implemented yet
}
