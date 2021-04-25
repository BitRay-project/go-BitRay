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

/*
Package vm implements the BitRay Virtual Machine.

The vm package implements one BVM, a byte code VM. The BC (Byte Code) VM loops
over a set of bytes and executes them according to the set of rules defined
in the Ethereum yellow paper. //yellow paper
*/
package vm
