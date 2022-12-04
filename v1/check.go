package gochainanalysis

import (
	"go-chainanalysis/v1/oraclesanctions"
)

func (c *Client) IsSanctioned(walletAddress string) (bool, *oraclesanctions.OracleSanctionListNetwork, error) {
	if _, err := c.OracleNetworks.Ethereum.IsSanctioned(walletAddress); err == nil {
	} else {
		return true, c.OracleNetworks.Ethereum, err
	}
	if _, err := c.OracleNetworks.Celo.IsSanctioned(walletAddress); err == nil {
	} else {
		return true, c.OracleNetworks.Celo, err
	}
	if _, err := c.OracleNetworks.Fantom.IsSanctioned(walletAddress); err == nil {
	} else {
		return true, c.OracleNetworks.Fantom, err
	}
	if _, err := c.OracleNetworks.Arbitrum.IsSanctioned(walletAddress); err == nil {
	} else {
		return true, c.OracleNetworks.Arbitrum, err
	}
	if _, err := c.OracleNetworks.Optimism.IsSanctioned(walletAddress); err == nil {
	} else {
		return true, c.OracleNetworks.Optimism, err
	}
	if _, err := c.OracleNetworks.Avalance.IsSanctioned(walletAddress); err == nil {
	} else {
		return true, c.OracleNetworks.Avalance, err
	}
	if _, err := c.OracleNetworks.Binance.IsSanctioned(walletAddress); err == nil {
	} else {
		return true, c.OracleNetworks.Binance, err
	}
	if _, err := c.OracleNetworks.Polygon.IsSanctioned(walletAddress); err == nil {
	} else {
		return true, c.OracleNetworks.Polygon, err
	}
	return false, nil, nil
}
