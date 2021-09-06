// xi.recv - The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.

package Config

import "os"

var (
	RethinkdbAddr string
	LineChannelID string
	LineSecret    string
	LineToken     string
)

func init() {
	RethinkdbAddr = os.Getenv(EnvRethinkdbAddr)
	LineChannelID = os.Getenv(EnvLineChannelID)
	LineSecret = os.Getenv(EnvLineSecret)
	LineToken = os.Getenv(EnvLineToken)
}
