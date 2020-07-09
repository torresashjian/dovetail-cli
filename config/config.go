/*
* Copyright © 2018. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package config

import (
	"github.com/spf13/viper"
)

// Blockchains Returns a list of the supported blockchains
func Blockchains() []string {
	return []string{HYPERLEDGER_FABRIC, CORDA}
}

// IsNodeVerbose returns whether the node log should be verbose
func IsNodeVerbose() bool {
	return viper.GetBool(NODE_VERBOSE_KEY)
}
