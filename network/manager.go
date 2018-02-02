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
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	uuid "github.com/satori/go.uuid"
)

// Enumeration of different networks available. A node configured with one
// network will only successfully connect to nodes of the same network. To be
// used for testing & iteration.
var (
	Odin = []byte{79, 68, 73, 78}
	Thor = []byte{84, 72, 79, 82}
	Loki = []byte{76, 79, 75, 73}
)

// wrapper around the standard dial function
var dial = func(address string) (net.Conn, error) { return net.Dial("tcp", address) }

// Manager represents the manager of all network components.
type Manager struct {
	log      zerolog.Logger
	wg       *sync.WaitGroup
	cfg      *Config
	book     *Book
	slots    slotManager
	peers    peerManager
	rep      reputationManager
	handlers handlerManager
	stop     chan struct{}
	pending  uint
}

// NewManager will initialize the completely wired up networking dependencies.
func NewManager(log zerolog.Logger, codec Codec, options ...func(*Config)) *Manager {

	// add the package information to the top package level logger
	log = log.With().Str("package", "network").Logger()

	// initialize the package-wide waitgroup
	wg := &sync.WaitGroup{}

	// initialize the default configuration and apply custom options
	cfg := &Config{
		network:    Odin,
		listen:     false,
		address:    "0.0.0.0:31337",
		minPeers:   3,
		maxPeers:   10,
		maxPending: 16,
		nonce:      uuid.NewV4().Bytes(),
		interval:   time.Second * 1,
		codec:      codec,
		bufferSize: 16,
	}
	for _, option := range options {
		option(cfg)
	}

	// TODO: validate the configuration parameters

	// initialize the network component with all state
	mgr := &Manager{
		log:      log,
		wg:       wg,
		cfg:      cfg,
		slots:    newSimpleSlotManager(cfg.maxPending),
		peers:    newSimplePeerManager(cfg.minPeers, cfg.maxPeers),
		rep:      newSimpleReputationManager(),
		handlers: &simpleHandlerManager{},
		stop:     make(chan struct{}),
	}

	// TODO: separate book package and inject so we can add addresses in main
	mgr.book.Found("127.0.0.1:31330")

	// blacklist our own address
	mgr.book.Invalid(cfg.address)

	// initialize the connection dropper, the outgoing connection dialer and
	// the incoming connection server
	// TODO: restart the initial handlers

	return mgr
}

// Stop will shut down all routines and wait for them to end.
func (mgr *Manager) Stop() {
	close(mgr.stop)
	mgr.peers.DropAll()
	mgr.wg.Wait()
}

// PendingCount returns the number of pending peer connections.
func (mgr *Manager) PendingCount() uint {
	return mgr.pending
}

// ClaimSlot claims one pending connection slot.
func (mgr *Manager) ClaimSlot() error {
	mgr.pending++
	return nil
}

// ReleaseSlot releases one pending connection slot.
func (mgr *Manager) ReleaseSlot() {
	mgr.pending--
}

// GetAddress returns a random address for connection.
func (mgr *Manager) GetAddress() (string, error) {
	addresses := mgr.book.Sample(1, isActive(false), byRandom())
	if len(addresses) == 0 {
		return "", errors.New("no inactive addresses in book")
	}
	return addresses[0], nil
}

// AddressSample returns a random address sample.
func (mgr *Manager) AddressSample() ([]string, error) {
	addresses := mgr.book.Sample(16, isAny(), byRandom())
	if len(addresses) == 0 {
		return nil, errors.New("no addresses in book")
	}
	return addresses, nil
}

// StartConnector will try to launch a new connection attempt.
func (mgr *Manager) StartConnector() {
	addresses := mgr.book.Sample(1, isActive(false), byRandom())
	if len(addresses) == 0 {
		mgr.log.Error().Msg("could not get address for connector")
		return
	}
	mgr.wg.Add(1)
	// TODO: launch handler to create connection
}

// StartListener will start a listener on a given port.
func (mgr *Manager) StartListener(stop <-chan struct{}) {
	mgr.wg.Add(1)
	listener := &simpleListenManager{}
	go handleListening(mgr.log, mgr.wg, mgr.cfg, mgr.handlers, listener, stop)
}

// StartAcceptor will start accepting an incoming connection.
func (mgr *Manager) StartAcceptor(conn net.Conn) {
	mgr.wg.Add(1)
	go handleAccepting(mgr.log, mgr.wg, mgr.cfg, mgr.slots, mgr.peers, mgr.rep, conn)
}
