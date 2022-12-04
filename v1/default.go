package go_chainanalysis

import "math/big"

var Default = Client{
	OracleSanctionListNetworks: OracleSanctionListNetworks{
		Ethereum: OracleSanctionListNetwork{
			Name: "Ethereum",
			ContractAddress: "0x40C57923924B5c5c5455c48D93317139ADDaC8fb",
			RPCURL: "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161",
			ChainID: big.NewInt(1),
		},
		Polygon:  OracleSanctionListNetwork{
			Name: "Polygon",
			ContractAddress: "0x40C57923924B5c5c5455c48D93317139ADDaC8fb",
			RPCURL: "https://polygon-rpc.com",
			ChainID: big.NewInt(137),
		},
		Binance:  OracleSanctionListNetwork{
			Name: "Binance",
			ContractAddress: "0x40C57923924B5c5c5455c48D93317139ADDaC8fb",
			RPCURL: "https://bsc-dataseed.binance.org",
			ChainID: big.NewInt(56),
		},
		Avalance: OracleSanctionListNetwork{
			Name: "Avalance",
			ContractAddress: "0x40C57923924B5c5c5455c48D93317139ADDaC8fb",
			RPCURL: "https://api.avax.network/ext/bc/C/rpc",
			ChainID: big.NewInt(43114),
		},
		Optimism: OracleSanctionListNetwork{
			Name: "Optimism",
			ContractAddress: "0x40C57923924B5c5c5455c48D93317139ADDaC8fb",
			RPCURL: "https://mainnet.optimism.io",
			ChainID: big.NewInt(10),
		},
		Arbitrum: OracleSanctionListNetwork{
			Name: "Arbitrum",
			ContractAddress: "0x40C57923924B5c5c5455c48D93317139ADDaC8fb",
			RPCURL: "https://rpc.ankr.com/arbitrum",
			ChainID: big.NewInt(42161),
		},
		Fantom:   OracleSanctionListNetwork{
			Name: "Fantom",
			ContractAddress: "0x40c57923924b5c5c5455c48d93317139addac8fb",
			RPCURL: "https://rpc.ftm.tools/",
			ChainID: big.NewInt(250),
		},
		Celo:     OracleSanctionListNetwork{
			Name: "Celo",
			ContractAddress: "0x40C57923924B5c5c5455c48D93317139ADDaC8fb",
			RPCURL: "https://rpc.ankr.com/celo",
			ChainID: big.NewInt(42220),
		},
	},
}

func (c *Client) UseDefault() *Client {
	return &Default
}
