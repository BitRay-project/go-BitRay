// Copyright 2021 The go-BitRay Authors
// This file is part of the go-BitRay library.
//
// The go-BitRay library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-BitRay library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-BitRay library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/BitRay-project/go-BitRay/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// BitRay Foundation Go Bootnodes
	
}


const dnsPrefix = " "

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. 
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".btrcdisco.net"
}
