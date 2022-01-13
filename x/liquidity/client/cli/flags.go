package cli

// DONTCOVER

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagPairId        = "pair-id"
	FlagXDenom        = "x-denom"
	FlagYDenom        = "y-denom"
	FlagPoolCoinDenom = "pool-coin-denom"
	FlagReserveAcc    = "reserve-acc"
)

func flagSetPools() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagXDenom, "", "The X coin denomination")
	fs.String(FlagYDenom, "", "The Y coin denomination")
	fs.String(FlagPairId, "", "The pair id")

	return fs
}

func flagSetPool() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagPoolCoinDenom, "", "The denomination of the pool coin")
	fs.String(FlagReserveAcc, "", "The Bech32 address of the reserve account")

	return fs
}

func flagSetPairs() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagXDenom, "", "The X coin denomination")
	fs.String(FlagYDenom, "", "The Y coin denomination")

	return fs
}