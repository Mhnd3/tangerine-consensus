// Copyright 2018 The dexon-consensus-core Authors
// This file is part of the dexon-consensus-core library.
//
// The dexon-consensus-core library is free software: you can redistribute it
// and/or modify it under the terms of the GNU Lesser General Public License as
// published by the Free Software Foundation, either version 3 of the License,
// or (at your option) any later version.
//
// The dexon-consensus-core library is distributed in the hope that it will be
// useful, but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser
// General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the dexon-consensus-core library. If not, see
// <http://www.gnu.org/licenses/>.

package types

import (
	"github.com/dexon-foundation/dexon-consensus-core/common"
)

// AgreementResult describes an agremeent result.
type AgreementResult struct {
	BlockHash common.Hash `json:"block_hash"`
	Round     uint64      `json:"round"`
	Position  Position    `json:"position"`
	Votes     []Vote      `json:"votes"`
}

// BlockRandomnessResult describes a block randomness result
type BlockRandomnessResult struct {
	BlockHash  common.Hash `json:"block_hash"`
	Round      uint64      `json:"round"`
	Randomness []byte      `json:"randomness"`
}