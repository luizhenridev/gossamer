// Copyright 2025 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package common

// BlockOrigin is the block data origin
type BlockOrigin uint

const (
	// NetworkInitialSync means the block is part of the initial sync with the network.
	NetworkInitialSync BlockOrigin = iota
	// NetworkBroadcast means the block was broadcasted on the network.
	NetworkBroadcast
)
