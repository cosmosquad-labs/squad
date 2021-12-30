package types

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

// PoolName returns unique name of the pool consists of given reserve coin denoms and type id.
func PoolName(reserveCoinDenoms []string, poolTypeID uint32) string {
	return strings.Join(append(SortDenoms(reserveCoinDenoms), strconv.FormatUint(uint64(poolTypeID), 10)), "/")
}

// AlphabeticalDenomPair returns denom pairs that are alphabetically sorted.
func AlphabeticalDenomPair(denom1, denom2 string) (resDenom1, resDenom2 string) {
	if denom1 > denom2 {
		return denom2, denom1
	}
	return denom1, denom2
}

// SortDenoms sorts denoms in alphabetical order.
func SortDenoms(denoms []string) []string {
	sort.Strings(denoms)
	return denoms
}

// GetPoolReserveAcc returns the address of the pool's reserve account.
func GetPoolReserveAcc(poolName string) sdk.AccAddress {
	return sdk.AccAddress(crypto.AddressHash([]byte(poolName)))
}

// GetPoolCoinDenom returns the denomination of the pool coin.
// PoolCoinDenom originally had prefix with / splitter, but it is removed to pass ibc-transfer validation.
func GetPoolCoinDenom(poolName string) string {
	return fmt.Sprintf("%s%X", PoolCoinDenomPrefix, sha256.Sum256([]byte(poolName)))
}

// GetReserveAcc extracts and returns reserve account from pool coin denom.
func GetReserveAcc(poolCoinDenom string) (sdk.AccAddress, error) {
	if err := sdk.ValidateDenom(poolCoinDenom); err != nil {
		return nil, err
	}

	if !strings.HasPrefix(poolCoinDenom, PoolCoinDenomPrefix) {
		return nil, ErrInvalidDenom
	}

	poolCoinDenom = strings.TrimPrefix(poolCoinDenom, PoolCoinDenomPrefix)
	if len(poolCoinDenom) != 64 {
		return nil, ErrInvalidDenom
	}

	return sdk.AccAddressFromHex(poolCoinDenom[:40])
}
