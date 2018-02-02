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

package network

import (
	"net"
	"time"

	"github.com/stretchr/testify/mock"
)

type AddrMock struct {
	mock.Mock
}

func (am *AddrMock) Network() string {
	args := am.Called()
	return args.String(0)
}

func (am *AddrMock) String() string {
	args := am.Called()
	return args.String(0)
}

type ConnMock struct {
	mock.Mock
}

func (cm *ConnMock) Read(b []byte) (int, error) {
	args := cm.Called(b)
	return args.Int(0), args.Error(1)
}

func (cm *ConnMock) Write(b []byte) (int, error) {
	args := cm.Called(b)
	return args.Int(0), args.Error(1)
}

func (cm *ConnMock) Close() error {
	args := cm.Called()
	return args.Error(0)
}

func (cm *ConnMock) LocalAddr() net.Addr {
	args := cm.Called()
	return args.Get(0).(*AddrMock)
}

func (cm *ConnMock) RemoteAddr() net.Addr {
	args := cm.Called()
	return args.Get(0).(*AddrMock)
}

func (cm *ConnMock) SetDeadline(t time.Time) error {
	args := cm.Called(t)
	return args.Error(0)
}

func (cm *ConnMock) SetReadDeadline(t time.Time) error {
	args := cm.Called(t)
	return args.Error(0)
}

func (cm *ConnMock) SetWriteDeadline(t time.Time) error {
	args := cm.Called(t)
	return args.Error(0)
}

type SlotManagerMock struct {
	mock.Mock
}

func (sm *SlotManagerMock) Claim() error {
	args := sm.Called()
	return args.Error(0)
}

func (sm *SlotManagerMock) Release() error {
	args := sm.Called()
	return args.Error(0)
}

type PeerManagerMock struct {
	mock.Mock
}

func (pm *PeerManagerMock) Add(conn net.Conn, nonce []byte) error {
	args := pm.Called(conn, nonce)
	return args.Error(0)
}

func (pm *PeerManagerMock) DropAll() {
	_ = pm.Called()
}

func (pm *PeerManagerMock) Count() uint {
	args := pm.Called()
	return uint(args.Int(0))
}

func (pm *PeerManagerMock) Addresses() []string {
	args := pm.Called()
	return args.Get(0).([]string)
}

type ReputationManagerMock struct {
	mock.Mock
}

func (rm *ReputationManagerMock) Error(address string) {
	_ = rm.Called(address)
}

func (rm *ReputationManagerMock) Failure(address string) {
	_ = rm.Called(address)
}

func (rm *ReputationManagerMock) Invalid(address string) {
	_ = rm.Called(address)
}

func (rm *ReputationManagerMock) Success(address string) {
	_ = rm.Called(address)
}

func (rm *ReputationManagerMock) Score(address string) float32 {
	args := rm.Called(address)
	return float32(args.Get(0).(float64))
}
