package curl

import _ "embed"

//go:generate curl -LO https://cosmo.zip/pub/cosmos/bin/curl
//go:embed curl
var Curl []byte
