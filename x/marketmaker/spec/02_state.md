<!-- order: 2 -->

# State

The `marketmaker`
module keeps track of the market maker states.

## MarketMaker

```go
type MarketMaker struct {
    Address    string
    PairId     uint64
    Intent     bool
}
```

## Incentive

```go
type Incentive struct {
    Address   string
    Claimable sdk.Coins
}
```

# Parameter

- ModuleName: `marketmaker`
- RouterKey: `marketmaker`
- StoreKey: `marketmaker`
- QuerierRoute: `marketmaker`

# Store

Stores are KVStores in the `multistore`. The key to find the store is the first parameter in the list.

`Address -> nil`

### **The index key to get the market maker object by pair id and address**

- MarketMakerIndexByPairIdKey: `[]byte{0xc0} | PairId | Address -> nil`

### **The index key to get the market maker object by address and pair id**

- MarketMakerIndexByAddrKey: `[]byte{0xc1} | AddressLen (1 byte) | Address | PairId -> nil`

### **The key to get the incentive object**

- IncentiveKey: `[]byte{0xc5} | Address -> ProtocalBuffer(Incentive)`