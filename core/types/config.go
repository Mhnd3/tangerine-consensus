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
	"encoding/binary"
	"math"
	"time"
)

// Config stands for Current Configuration Parameters.
type Config struct {
	// Network related.
	NumChains uint32

	// Lambda related.
	LambdaBA  time.Duration
	LambdaDKG time.Duration

	// Total ordering related.
	K        int
	PhiRatio float32

	// NodeSet related.
	NumNotarySet  int
	NumWitnessSet int
	NumDKGSet     int

	// Time related.
	RoundInterval    time.Duration
	MinBlockInterval time.Duration
	MaxBlockInterval time.Duration
}

// Bytes returns []byte representation of Config.
func (c *Config) Bytes() []byte {
	binaryNumChains := make([]byte, 4)
	binary.LittleEndian.PutUint32(binaryNumChains, c.NumChains)

	binaryLambdaBA := make([]byte, 8)
	binary.LittleEndian.PutUint64(
		binaryLambdaBA, uint64(c.LambdaBA.Nanoseconds()))
	binaryLambdaDKG := make([]byte, 8)
	binary.LittleEndian.PutUint64(
		binaryLambdaDKG, uint64(c.LambdaDKG.Nanoseconds()))

	binaryK := make([]byte, 4)
	binary.LittleEndian.PutUint32(binaryK, uint32(c.K))
	binaryPhiRatio := make([]byte, 4)
	binary.LittleEndian.PutUint32(binaryPhiRatio, math.Float32bits(c.PhiRatio))

	binaryNumNotarySet := make([]byte, 4)
	binary.LittleEndian.PutUint32(binaryNumNotarySet, uint32(c.NumNotarySet))
	binaryNumWitnessSet := make([]byte, 4)
	binary.LittleEndian.PutUint32(binaryNumWitnessSet, uint32(c.NumWitnessSet))
	binaryNumDKGSet := make([]byte, 4)
	binary.LittleEndian.PutUint32(binaryNumDKGSet, uint32(c.NumDKGSet))

	binaryRoundInterval := make([]byte, 8)
	binary.LittleEndian.PutUint64(binaryRoundInterval,
		uint64(c.RoundInterval.Nanoseconds()))
	binaryMinBlockInterval := make([]byte, 8)
	binary.LittleEndian.PutUint64(binaryMinBlockInterval,
		uint64(c.MinBlockInterval.Nanoseconds()))
	binaryMaxBlockInterval := make([]byte, 8)
	binary.LittleEndian.PutUint64(binaryMaxBlockInterval,
		uint64(c.MaxBlockInterval.Nanoseconds()))

	enc := make([]byte, 0, 40)
	enc = append(enc, binaryNumChains...)
	enc = append(enc, binaryLambdaBA...)
	enc = append(enc, binaryLambdaDKG...)
	enc = append(enc, binaryK...)
	enc = append(enc, binaryPhiRatio...)
	enc = append(enc, binaryNumNotarySet...)
	enc = append(enc, binaryNumWitnessSet...)
	enc = append(enc, binaryNumDKGSet...)
	enc = append(enc, binaryRoundInterval...)
	enc = append(enc, binaryMinBlockInterval...)
	enc = append(enc, binaryMaxBlockInterval...)
	return enc
}
