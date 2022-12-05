### go-chainanalysis

A Golang SDK for [chainanalysis](https://www.chainalysis.com/free-cryptocurrency-sanctions-screening-tools/) API.

For best compatibility, please use Go >= 1.7.

A list of currently compatible networks can be found [here](https://go.chainalysis.com/chainalysis-oracle-docs.html)

### Dependencies
Name | Description | 
------------ | ------------ |
[go-ethereum](https://geth.ethereum.org/) | Used for interacting with EVM |


### API List

Name | Description | Status
------------ | ------------ | ------------
[chainanalysis-oracle-for-sanctions-screening](https://go.chainalysis.com/chainalysis-oracle-docs.html) | Details on the on-chain API | <input type="checkbox" checked> Implemented

### Installation

```shell
go get github.com/soloth/go-chainanalysis
```

### Importing

```golang
import (
    "github.com/soloth/go-chainanalysis/v1"
)
```

### Generating a client

```golang
import (
	"github.com/soloth/go-chainanalysis/v1"
)

func main() {
	client := gochainanalysis.NewClient()
}
```

### Using default config
The default config uses public RPC endpoints. It is recommends you change them for a production environment.
```golang
import (
	"github.com/soloth/go-chainanalysis/v1"
)

func main() {
	client := gochainanalysis.NewClient().UseDefault()
}
```

### An example of changing an rpc endpoint
```golang
import (
	"go-chainanalysis/v1"
	"math/big"
)

func main() {
	client := gochainanalysis.NewClient()
	
	client.OracleSanctionListNetworks.Ethereum = gochainanalysis.OracleSanctionListNetwork{
		Name: "Ethereum",
		ContractAddress: "0x40C57923924B5c5c5455c48D93317139ADDaC8fb",
		RPCURL: "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161",
		ChainID: big.NewInt(1),
	}
}

```

### An example of updating all rpc URLS
Please note, any empty values in `OracleSanctionListNetwork` other than the name will be ignored during any checks.
```golang
import (
	"go-chainanalysis/v1"
	"math/big"
)

func main() {
	client := gochainanalysis.NewClient()
	
	client.OracleSanctionListNetworks = gochainanalysis.OracleSanctionListNetworks{
		Ethereum: gochainanalysis.OracleSanctionListNetwork{
			Name: "Ethereum",
			ContractAddress: "0x40C57923924B5c5c5455c48D93317139ADDaC8fb",
			RPCURL: "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161",
			ChainID: big.NewInt(1),
		},
		Polygon:  gochainanalysis.OracleSanctionListNetwork{
			Name: "Polygon",
			ContractAddress: "0x40C57923924B5c5c5455c48D93317139ADDaC8fb",
			RPCURL: "https://polygon-rpc.com",
			ChainID: big.NewInt(137),
		},
		...
	},
}

```

### Checking for sanctions
```golang
import (
	"go-chainanalysis/v1"
	"log"
)

func main() {
	// Check each RPC node one at a time
	client := gochainanalysis.NewClient().UseDefault()
	isSanctioned, network, err := client.IsSanctioned("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")
	if err != nil {
		log.Fatal(network, err)
	}
	log.Println(isSanctioned)

	// Check all RPC nodes at the same time
	client := gochainanalysis.NewClient().UseDefault()
	isSanctioned, network, err := client.IsSanctionedConcurrent("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")
	if err != nil {
		log.Fatal(network, err)
	}
	log.Println(isSanctioned)
}
```