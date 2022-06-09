package amm

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	utils "github.com/cosmosquad-labs/squad/types"
)

func parseOrders(s string) []Order {
	orderRe := regexp.MustCompile(`(\d+)(?:\((\d+)\))?`)
	var orders []Order
	for _, line := range strings.Split(s, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		chunks := strings.Split(line, "|")
		if len(chunks) != 3 {
			panic(fmt.Errorf("wrong number of chunks in %q: %d", line, len(chunks)))
		}
		price := sdk.MustNewDecFromStr(strings.TrimSpace(chunks[1]))
		parseSide := func(dir OrderDirection, s string) []Order {
			var orders []Order
			for _, chunks := range orderRe.FindAllStringSubmatch(s, -1) {
				amt, ok := sdk.NewIntFromString(chunks[1])
				if !ok {
					panic(fmt.Errorf("invalid amount: %s", chunks[1]))
				}
				batchId, err := strconv.ParseUint(chunks[2], 10, 64)
				if err != nil {
					if chunks[2] == "" {
						batchId = 0
					} else {
						panic(fmt.Errorf("invalid batch id: %s", chunks[2]))
					}
				}
				orders = append(orders, NewUserOrder(0, batchId, dir, price, amt))
			}
			return orders
		}
		orders = append(orders, parseSide(Sell, chunks[0])...)
		orders = append(orders, parseSide(Buy, chunks[2])...)
	}
	return orders
}

func TestInstantMatch(t *testing.T) {
	orders := parseOrders(`
    | 1.2 | 5 7
5   | 0.9 | 3
6 3 | 0.8 |
4   | 0.7 |
`)
	ob := NewOrderBook(orders...)
	matched := ob.InstantMatch(utils.ParseDec("1.0"))
	require.True(t, matched)
	for _, order := range orders {
		fmt.Printf("%s %s at %s\n", order.GetDirection(), order.GetAmount(), order.GetPrice())
		if len(order.GetMatchRecords()) == 0 {
			continue
		}
		for _, record := range order.GetMatchRecords() {
			fmt.Printf("  match %s at %s\n", record.Amount, record.Price)
		}
	}
}
