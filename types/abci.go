package types

import abci "github.com/tendermint/tendermint/abci/types"

// InitChainer initializes application state at genesis
type InitChainer func(ctx Context, req abci.RequestInitChain) abci.ResponseInitChain

// BeginBlocker runs code before the transactions in a block
//
// Note: applications which set create_empty_blocks=false will not have regular block timing and should use
// e.g. BFT timestamps rather than block height for any periodic BeginBlock logic
type BeginBlocker func(ctx Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock

// EndBlocker runs code after the transactions in a block and return updates to the validator set
//
// Note: applications which set create_empty_blocks=false will not have regular block timing and should use
// e.g. BFT timestamps rather than block height for any periodic EndBlock logic
type EndBlocker func(ctx Context, req abci.RequestEndBlock) abci.ResponseEndBlock

// PeerFilter responds to p2p filtering queries from Tendermint
type PeerFilter func(info string) abci.ResponseQuery

//
// side channel
//

// BeginSideBlocker runs code before the side transactions in a block
type BeginSideBlocker func(ctx Context, req abci.RequestBeginSideBlock) abci.ResponseBeginSideBlock

// DeliverSideTxHandler runs during each side trasaction in a block
type DeliverSideTxHandler func(ctx Context, tx Tx, req abci.RequestDeliverSideTx) abci.ResponseDeliverSideTx

// PostDeliverTxHandler runs after deliver tx
type PostDeliverTxHandler func(ctx Context, tx Tx, result Result)
