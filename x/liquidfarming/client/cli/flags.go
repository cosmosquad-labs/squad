package cli

// DONTCOVER

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagFarmer = "farmer"
)

func flagSetQueuedFarmings() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagFarmer, "", "farmer address")

	return fs
}
