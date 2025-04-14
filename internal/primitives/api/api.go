// Copyright 2025 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package api

// ProvideRuntimeAPI is the interface that provides a runtime api.
type ProvideRuntimeAPI[API any] interface {
	// Returns the runtime api.
	// The returned instance will keep track of modifications to the storage. Any successful call to an api function,
	// will commit its changes to an internal buffer. Otherwise, the modifications will be discarded. The modifications
	// will not be applied to the storage, even on a commit.
	RuntimeAPI() API
}
