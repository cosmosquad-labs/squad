package types

import (
	"fmt"
	"regexp"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	reBearingNameString = fmt.Sprintf(`[a-zA-Z][a-zA-Z0-9-]{0,%d}`, MaxBearingNameLength-1)
	reBearingName       = regexp.MustCompile(fmt.Sprintf(`^%s$`, reBearingNameString))
)

// String returns a human-readable string representation of the bearing.
func (bearing Bearing) String() string {
	out, _ := bearing.MarshalYAML()
	return out.(string)
}

// MarshalYAML returns the YAML representation of a Bearing.
func (bearing Bearing) MarshalYAML() (interface{}, error) {
	bz, err := codec.MarshalYAML(codec.NewProtoCodec(codectypes.NewInterfaceRegistry()), &bearing)
	if err != nil {
		return nil, err
	}
	return string(bz), err
}

// Validate validates the bearing.
func (bearing Bearing) Validate() error {
	if err := ValidateName(bearing.Name); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(bearing.DestinationAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid destination address %s: %v", bearing.DestinationAddress, err)
	}

	if _, err := sdk.AccAddressFromBech32(bearing.SourceAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid source address %s: %v", bearing.SourceAddress, err)
	}

	if !bearing.EndTime.After(bearing.StartTime) {
		return ErrInvalidStartEndTime
	}

	if !bearing.Rate.IsPositive() {
		return sdkerrors.Wrapf(ErrInvalidBearingRate, "bearing rate must not be positive: %s", bearing.Rate)
	} else if bearing.Rate.GT(sdk.OneDec()) {
		return sdkerrors.Wrapf(ErrInvalidBearingRate, "bearing rate must not exceed 1: %s", bearing.Rate)
	}

	return nil
}

// Collectible validates the bearing has reached its start time and that the end time has not elapsed.
func (bearing Bearing) Collectible(blockTime time.Time) bool {
	return !bearing.StartTime.After(blockTime) && bearing.EndTime.After(blockTime)
}

// CollectibleBearings returns only the valid and started and not expired bearings based on the given block time.
func CollectibleBearings(bearings []Bearing, blockTime time.Time) (collectibleBearings []Bearing) {
	for _, bearing := range bearings {
		if bearing.Collectible(blockTime) {
			collectibleBearings = append(collectibleBearings, bearing)
		}
	}
	return
}

// ValidateName is the default validation function for Bearing.Name.
// A bearing name only allows alphabet letters(`A-Z, a-z`), digit numbers(`0-9`), and `-`.
// It doesn't allow spaces and the maximum length is 50 characters.
func ValidateName(name string) error {
	if !reBearingName.MatchString(name) {
		return sdkerrors.Wrap(ErrInvalidBearingName, name)
	}
	return nil
}

// BearingsBySource defines the total rate of bearing lists.
type BearingsBySource struct {
	Bearings        []Bearing
	CollectionCoins []sdk.Coins
	TotalRate       sdk.Dec
}

type BearingsBySourceMap map[string]BearingsBySource

// GetBearingsBySourceMap returns BearingsBySourceMap that has a list of bearings and their total rate
// which contain the same SourceAddress. It can be used to track of what bearings are available with SourceAddress
// and validate their total rate.
func GetBearingsBySourceMap(bearings []Bearing) BearingsBySourceMap {
	bearingsMap := make(BearingsBySourceMap)
	for _, bearing := range bearings {
		if bearingsBySource, ok := bearingsMap[bearing.SourceAddress]; ok {
			bearingsBySource.TotalRate = bearingsBySource.TotalRate.Add(bearing.Rate)
			bearingsBySource.Bearings = append(bearingsBySource.Bearings, bearing)
			bearingsMap[bearing.SourceAddress] = bearingsBySource
		} else {
			bearingsMap[bearing.SourceAddress] = BearingsBySource{
				Bearings:  []Bearing{bearing},
				TotalRate: bearing.Rate,
			}
		}
	}
	return bearingsMap
}
