package gochainanalysis

import "math/big"

type Client struct {
	OracleSanctionListNetworks OracleSanctionListNetworks
}

type OracleSanctionListNetworks struct {
	Ethereum OracleSanctionListNetwork
	Polygon  OracleSanctionListNetwork
	Binance  OracleSanctionListNetwork
	Avalance OracleSanctionListNetwork
	Optimism OracleSanctionListNetwork
	Arbitrum OracleSanctionListNetwork
	Fantom   OracleSanctionListNetwork
	Celo     OracleSanctionListNetwork
}

type OracleSanctionListNetwork struct {
	Name                        string
	ContractAddress string
	RPCURL                      string
	ChainID                     *big.Int
}
