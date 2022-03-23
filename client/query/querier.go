package query

import (
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributionTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/strangelove-ventures/lens/client"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

type Query struct {
	Client  *client.ChainClient
	Options *QueryOptions
}

// Bank queries

// Balances returns the balance of all coins for a single account.
func (q *Query) Balances(address string) (*bankTypes.QueryAllBalancesResponse, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return balanceWithAddressRPC(q, address)
}

// TotalSupply returns the supply of all coins
func (q *Query) TotalSupply() (*bankTypes.QueryTotalSupplyResponse, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return totalSupplyRPC(q)
}

// DenomsMetadata returns the metadata for all denoms
func (q *Query) DenomsMetadata() (*bankTypes.QueryDenomsMetadataResponse, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return denomsMetadataRPC(q)
}

// Staking queries

// Delegation returns the delegations to a particular validator
func (q *Query) Delegation(delegator, validator string) (*stakingTypes.DelegationResponse, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return delegationRPC(q, delegator, validator)
}

// Delegations returns all the delegations
func (q *Query) Delegations(delegator string) (*stakingTypes.QueryDelegatorDelegationsResponse, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return delegationsRPC(q, delegator)
}

// ValidatorDelegations returns all the delegations for a validator
func (q *Query) ValidatorDelegations(validator string) (*stakingTypes.QueryValidatorDelegationsResponse, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return validatorDelegationssRPC(q, validator)
}

// Distribution queries

// DelegatorValidators returns the validators of a delegator
func (q *Query) DelegatorValidators(delegator string) (*distributionTypes.QueryDelegatorValidatorsResponse, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return delegatorValidatorsRPC(q, delegator)
}

// Tendermint queries

// Block returns information about a block
func (q *Query) Block() (*coretypes.ResultBlock, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return blockRPC(q)
}

// BlockByHash returns information about a block by hash
func (q *Query) BlockByHash(hash string) (*coretypes.ResultBlock, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return blockByHashRPC(q, hash)
}

// BlockResults returns information about a block by hash
func (q *Query) BlockResults() (*coretypes.ResultBlockResults, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return blockResultsRPC(q)
}

// Status returns information about a node status
func (q *Query) Status() (*coretypes.ResultStatus, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return statusRPC(q)
}

// ABCIInfo returns general information about the ABCI application
func (q *Query) ABCIInfo() (*coretypes.ResultABCIInfo, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return abciInfoRPC(q)
}

// ABCIQuery returns data from a particular path in the ABCI application
func (q *Query) ABCIQuery(path string, data string, prove bool) (*coretypes.ResultABCIQuery, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return abciQueryRPC(q, path, data, prove)
}
