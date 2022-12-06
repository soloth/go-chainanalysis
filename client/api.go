package client

import "github.com/soloth/go-chainanalysis/client/api"

func (c *Client) GetSanctions(walletAddress string) (int, error) {
	var a api.API
	return a.IsSanctioned(walletAddress, c.APIKey)
}
