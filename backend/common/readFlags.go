package common

import (
	"flag"
)

func ReadFlags() {
	flag.BoolVar(&UseJwtAsApiResponse, "token", false, "Use jwt token into api handlers")
	flag.Parse()

	if UseJwtAsApiResponse {
		LogInfo("Using JWT as response")
	} else {
		LogInfo("Using proto as response")
	}
}
