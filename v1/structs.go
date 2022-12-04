package gochainanalysis

import (
	"go-chainanalysis/v1/oraclesanctions"
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