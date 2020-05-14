//  Copyright 2018 The go-ethereum Authors
//  Copyright 2019 The go-aigar Authors
//  This file is part of the go-aigar library.
//
//  The go-aigar library is free software: you can redistribute it and/or modify
//  it under the terms of the GNU Lesser General Public License as published by
//  the Free Software Foundation, either version 3 of the License, or
//  (at your option) any later version.
//
//  The go-aigar library is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
//  GNU Lesser General Public License for more details.
//
//  You should have received a copy of the GNU Lesser General Public License
//  along with the go-aigar library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Aigar network.
var MainnetBootnodes = []string{
	"enode://b92d12d807bfbb058369640ea6a17dc74adc79cec02ff50831cde012a48d7bd80ba91df08e23728e2c175dfa2accf4ea2a6e2b8a7bd53957fc34d85973e3581a@178.57.222.100:30301",  //Aigar
	"enode://8b8586095253483153c78c20670fa91eb26aa0c09e041234d3451079873470ad27c2aa5ff3ebdf2894be7b1edd307162083479925caea6619f4fde1cb3d8586d@185.22.232.218:30301",  //Guzelka
	"enode://aaba57ebf68850b0ed2aeb839fcbf844cfc2fcdd4eae1674e843a6ed0a8153abff4419bd1633856007d7e4ef1e2da28b52033ba9f447818a6c9d4cf0d8416539@185.22.233.129:30301",  //Bob Slay
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes []string

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes []string

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes []string

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes []string
