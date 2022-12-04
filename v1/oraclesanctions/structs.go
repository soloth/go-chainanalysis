package oraclesanctions

import "math/big"

type OracleSanctionListNetwork struct {
	Name                        string
	ContractAddress string
	RPCURL                      string
	ChainID                     *big.Int
}
