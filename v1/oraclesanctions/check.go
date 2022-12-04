package oraclesanctions

import (
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (c *OracleSanctionListNetwork) IsSanctioned(walletAddress string) (bool, error) {
	conn, err := ethclient.Dial(c.RPCURL)
	if err != nil {
		return true, err
	}

	verify, err := NewConsumerContract(common.HexToAddress(c.ContractAddress), conn)
	if err != nil {
		return true, err
	}

	isSanctioned, err := verify.IsSanctioned(nil, common.HexToAddress(walletAddress))
	if err != nil {
		return true, err
	}

	return isSanctioned, err
}

func (c *OracleSanctionListNetwork) IsSanctionedConcurrent(walletAddress string, _isSanctioned chan IsSanctionedStruct) {
	conn, err := ethclient.Dial(c.RPCURL)
	if err != nil {
		_isSanctioned <- IsSanctionedStruct{
			Network: *c,
			IsSanctioned: true,
			Err: err,
		}
		return
	}

	verify, err := NewConsumerContract(common.HexToAddress(c.ContractAddress), conn)
	if err != nil {
		_isSanctioned <- IsSanctionedStruct{
			Network: *c,
			IsSanctioned: true,
			Err: err,
		}
		return
	}

	isSanctioned, err := verify.IsSanctioned(nil, common.HexToAddress(walletAddress))
	if err != nil {
		_isSanctioned <- IsSanctionedStruct{
			Network: *c,
			IsSanctioned: true,
			Err: err,
		}
		return
	}
	_isSanctioned <- IsSanctionedStruct{
		Network: *c,
		IsSanctioned: isSanctioned,
		Err: err,
	}
	return
	return
}
