package client

import (
	"github.com/soloth/go-chainanalysis/client/oraclesanctions"
)

type Client struct {
	OracleNetworks *OracleSanctionListNetworks
}

type OracleSanctionListNetworks struct {
	Ethereum *oraclesanctions.OracleSanctionListNetwork
	Polygon  *oraclesanctions.OracleSanctionListNetwork
	Binance  *oraclesanctions.OracleSanctionListNetwork
	Avalance *oraclesanctions.OracleSanctionListNetwork
	Optimism *oraclesanctions.OracleSanctionListNetwork
	Arbitrum *oraclesanctions.OracleSanctionListNetwork
	Fantom   *oraclesanctions.OracleSanctionListNetwork
	Celo     *oraclesanctions.OracleSanctionListNetwork
}