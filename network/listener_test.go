// // Copyright (c) 2017 The Alvalor Authors
// //
// // This file is part of Alvalor.
// //
// // Alvalor is free software: you can redistribute it and/or modify
// // it under the terms of the GNU Affero General Public License as published by
// // the Free Software Foundation, either version 3 of the License, or
// // (at your option) any later version.
// //
// // Alvalor is distributed in the hope that it will be useful,
// // but WITHOUT ANY WARRANTY; without even the implied warranty of
// // MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// // GNU Affero General Public License for more detailb.
// //
// // You should have received a copy of the GNU Affero General Public License
// // along with Alvalor.  If not, see <http://www.gnu.org/licenses/>.
//
package network

import (
	"errors"
	"io/ioutil"
	"sync"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ListenerTestSuite struct {
	suite.Suite
	log zerolog.Logger
	wg  sync.WaitGroup
	cfg Config
}

func (suite *ListenerTestSuite) SetupTest() {
	suite.log = zerolog.New(ioutil.Discard)
	suite.wg = sync.WaitGroup{}
	suite.wg.Add(1)
	suite.cfg = Config{
		address: "66.37.13.55:5643",
	}
}

func (suite *ListenerTestSuite) TestHandleListeningReturnsIfListeningFails() {

	// arrange
	ln := &ListenerMock{}

	listener := &ListenManagerMock{}
	listener.On("Listen", suite.cfg.address).Return(ln, errors.New("could not listen"))

	handlers := &HandlerManagerMock{}

	// act
	go handleListening(suite.log, &suite.wg, &suite.cfg, handlers, listener, nil)
	suite.wg.Wait()

	// assert
	listener.AssertCalled(suite.T(), "Listen", suite.cfg.address)
	ln.AssertNotCalled(suite.T(), "Accept")
	ln.AssertNotCalled(suite.T(), "Close")
	handlers.AssertNotCalled(suite.T(), "Accept", mock.Anything)
}

func (suite *ListenerTestSuite) TestHandleListeningDoesNotStartAcceptorIfCantAcceptConnection() {

	// arrange
	ln := &ListenerMock{}
	ln.On("SetDeadline", mock.Anything).Return(nil)
	ln.On("Accept").Return(nil, errors.New("could not accept connection"))
	ln.On("Close").Return(nil)

	listener := &ListenManagerMock{}
	listener.On("Listen", suite.cfg.address).Return(ln, nil)

	handlers := &HandlerManagerMock{}

	// act
	go handleListening(suite.log, &suite.wg, &suite.cfg, handlers, listener, nil)
	suite.wg.Wait()

	// assert
	handlers.AssertNotCalled(suite.T(), "Accept", mock.Anything)
}

func (suite *ListenerTestSuite) TestHandleListeningStartsAcceptor() {

	// arrange
	conn := &ConnMock{}

	ln := &ListenerMock{}
	ln.On("SetDeadline", mock.Anything).Return(nil)
	ln.On("Accept").Return(conn, nil).Once()
	ln.On("Accept").Return(nil, errors.New("could not accept connection"))
	ln.On("Close").Return(nil)

	listener := &ListenManagerMock{}
	listener.On("Listen", suite.cfg.address).Return(ln, nil)

	handlers := &HandlerManagerMock{}
	handlers.On("Accept", conn)

	// act
	go handleListening(suite.log, &suite.wg, &suite.cfg, handlers, listener, nil)
	suite.wg.Wait()

	// assert
	handlers.AssertCalled(suite.T(), "Accept", conn)
}

func (suite *ListenerTestSuite) TestHandleListenerCloseError() {

	// arrange
	conn := &ConnMock{}

	ln := &ListenerMock{}
	ln.On("SetDeadline", mock.Anything).Return(nil)
	ln.On("Accept").Return(conn, nil).Once()
	ln.On("Accept").Return(nil, errors.New("could not accept connection"))
	ln.On("Close").Return(errors.New("could not close listener"))

	listener := &ListenManagerMock{}
	listener.On("Listen", suite.cfg.address).Return(ln, nil)

	handlers := &HandlerManagerMock{}
	handlers.On("Accept", conn)

	// act
	go handleListening(suite.log, &suite.wg, &suite.cfg, handlers, listener, nil)
	suite.wg.Wait()

	// assert
	ln.AssertCalled(suite.T(), "Close")
}

func (suite *ListenerTestSuite) TestListenerClosedChannelBreaksLoop() {

	// arrange
	stop := make(chan struct{})
	close(stop)

	ln := &ListenerMock{}
	ln.On("Close").Return(errors.New("could not close listener"))

	listener := &ListenManagerMock{}
	listener.On("Listen", suite.cfg.address).Return(ln, nil)

	handlers := &HandlerManagerMock{}

	// act
	go handleListening(suite.log, &suite.wg, &suite.cfg, handlers, listener, stop)
	suite.wg.Wait()

	// assert
	ln.AssertCalled(suite.T(), "Close")
}

func (suite *ListenerTestSuite) TestListenerTimeoutContinuesLoop() {

	// arrange
	conn := &ConnMock{}

	err := &ErrorMock{}
	err.On("Error").Return("error")
	err.On("Timeout").Return(true)

	ln := &ListenerMock{}
	ln.On("SetDeadline", mock.Anything).Return(nil)
	ln.On("Accept").Return(nil, err).Once()
	ln.On("Accept").Return(conn, nil).Once()
	ln.On("Accept").Return(nil, errors.New("could not accept connection"))
	ln.On("Close").Return(nil)

	listener := &ListenManagerMock{}
	listener.On("Listen", suite.cfg.address).Return(ln, nil)

	handlers := &HandlerManagerMock{}
	handlers.On("Accept", conn)

	// act
	go handleListening(suite.log, &suite.wg, &suite.cfg, handlers, listener, nil)
	suite.wg.Wait()

	// assert
	ln.AssertCalled(suite.T(), "Accept")
	handlers.AssertCalled(suite.T(), "Accept", conn)
}

func TestListenerTestSuite(t *testing.T) {
	suite.Run(t, new(ListenerTestSuite))
}
