package client

import (
	"errors"
	"github.com/soloth/go-chainanalysis/client/oraclesanctions"
	"time"
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

func (c *Client) IsSanctionedConcurrent(walletAddress string) (bool, oraclesanctions.OracleSanctionListNetwork, error) {

	checkChan := make(chan oraclesanctions.IsSanctionedStruct)

	go func() {
		go c.OracleNetworks.Ethereum.IsSanctionedConcurrent(walletAddress, checkChan)
		go c.OracleNetworks.Celo.IsSanctionedConcurrent(walletAddress, checkChan)
		go c.OracleNetworks.Fantom.IsSanctionedConcurrent(walletAddress, checkChan)
		go c.OracleNetworks.Arbitrum.IsSanctionedConcurrent(walletAddress, checkChan)
		go c.OracleNetworks.Optimism.IsSanctionedConcurrent(walletAddress, checkChan)
		go c.OracleNetworks.Avalance.IsSanctionedConcurrent(walletAddress, checkChan)
		go c.OracleNetworks.Binance.IsSanctionedConcurrent(walletAddress, checkChan)
		go c.OracleNetworks.Polygon.IsSanctionedConcurrent(walletAddress, checkChan)
	}()


	for i := 0; i < 8; i++ {
		select {
		case data := <-checkChan:
			if data.Err != nil {
				return data.IsSanctioned, data.Network, data.Err
			}
		case <-time.After(30 * time.Second):
			return false, oraclesanctions.OracleSanctionListNetwork{}, errors.New("use new rpc urls, current are too slow")
		}
	}

	return false, oraclesanctions.OracleSanctionListNetwork{}, nil
}
