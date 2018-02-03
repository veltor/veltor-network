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
	"sync"
	"time"

	"github.com/rs/zerolog"
)

func handleDialing(log zerolog.Logger, wg *sync.WaitGroup, cfg *Config, peers peerManager, slots slotManager, handlers handlerManager, stop <-chan struct{}) {
	defer wg.Done()

	// extract needed configuration parameters
	var (
		interval = cfg.interval
		minPeers = cfg.minPeers
		maxPeers = cfg.maxPeers
	)

	// configure logger and add start/stop messages
	log = log.With().Str("component", "dialer").Logger()
	log.Info().Msg("dialing routine started")
	defer log.Info().Msg("dialing routine stopped")

	// on each tick, check if we are below minimum peers and should have free
	// connection slots, then start a new dialer
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-stop:
			ticker.Stop()
			return
		case <-ticker.C:
		}
		peerCount := peers.Count()
		pendingCount := slots.Pending()
		if peerCount >= minPeers {
			continue
		}
		if peerCount+pendingCount >= maxPeers {
			continue
		}
		handlers.Connect()
	}
}
