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

// These are the multipliers for BTRC denominations.
// Example: To get the Rb value of an amount in 'grb', use
//
//    new(big.Int).Mul(value, big.NewInt(params.GRb))
//
const (
	Rb   = 1
	GRb  = 1e9
	BTRC = 1e18
)
