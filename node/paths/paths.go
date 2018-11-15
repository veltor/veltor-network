// Copyright (c) 2017 The Alvalor Authors
//
// This file is part of Alvalor.
//
// Alvalor is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Alvalor is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with Alvalor.  If not, see <http://www.gnu.org/licenses/>.

package paths

import (
	"github.com/alvalor/alvalor-go/types"
)

// Paths is responsible for tracking a certain path by downloading the entities
// required to complete it.
type Paths struct {
	current      []types.Hash
	inventories  Inventories
	transactions Transactions
	downloads    Downloads
}

// Follow sets a new path through the header tree to follow and complete.
func (tr *Paths) Follow(path []types.Hash) error {
	return nil
}

// Signal notifies the tracker than a new inventory has become available and
// related transaction downloads should be started, if pending.
func (tr *Paths) Signal(hash types.Hash) error {

	return nil
}
