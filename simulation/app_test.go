// Copyright 2018 The dexon-consensus Authors
// This file is part of the dexon-consensus library.
//
// The dexon-consensus library is free software: you can redistribute it
// and/or modify it under the terms of the GNU Lesser General Public License as
// published by the Free Software Foundation, either version 3 of the License,
// or (at your option) any later version.
//
// The dexon-consensus library is distributed in the hope that it will be
// useful, but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser
// General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the dexon-consensus library. If not, see
// <http://www.gnu.org/licenses/>.

package simulation

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"gitlab.com/byzantine-lab/tangerine-consensus/core"
	"gitlab.com/byzantine-lab/tangerine-consensus/core/types"
)

type SimAppSuite struct {
	suite.Suite
}

func (s *SimAppSuite) TestAppInterface() {
	var app core.Application
	app = newSimApp(types.NodeID{}, nil, nil)
	s.NotPanics(func() {
		_ = app.(core.Debug)
	})
}

func TestSimApp(t *testing.T) {
	suite.Run(t, new(SimAppSuite))
}
