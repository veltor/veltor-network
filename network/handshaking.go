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
	"bytes"
	"net"
	"sync"

	"github.com/rs/zerolog"
)

func handleIncoming(log zerolog.Logger, wg *sync.WaitGroup, network []byte, nonce []byte, input <-chan net.Conn, output chan<- net.Conn) {
	defer wg.Done()
	log = log.With().Str("component", "incoming").Logger()
	log.Info().Msg("incoming handshake routine started")
	defer log.Info().Msg("incoming handshake routine stopped")
	ack := append(network, nonce...)
	syn := make([]byte, len(ack))
	for conn := range input {
		address := conn.RemoteAddr().String()
		_, err := conn.Read(syn)
		if err != nil {
			log.Error().Str("address", address).Err(err).Msg("could not read syn packet")
			conn.Close()
			continue
		}
		networkIn := syn[:len(network)]
		if !bytes.Equal(networkIn, network) {
			log.Error().Str("address", address).Bytes("network", network).Bytes("network_in", networkIn).Msg("network mismatch")
			conn.Close()
			continue
		}
		nonceIn := syn[len(network):]
		if bytes.Equal(nonceIn, nonce) {
			log.Error().Str("address", address).Bytes("nonce", nonce).Msg("identical nonce")
			conn.Close()
			continue
		}
		_, err = conn.Write(ack)
		if err != nil {
			log.Error().Str("address", address).Err(err).Msg("could not write ack packet")
			conn.Close()
			continue
		}
		output <- conn
	}
}

func handleOutgoing(log zerolog.Logger, wg *sync.WaitGroup, network []byte, nonce []byte, input <-chan net.Conn, output chan<- net.Conn) {
	defer wg.Done()
	log = log.With().Str("component", "outgoing").Logger()
	log.Info().Msg("outgoing handshake routine started")
	defer log.Info().Msg("outgoing handshake routine stopped")
	syn := append(network, nonce...)
	ack := make([]byte, len(syn))
	for conn := range input {
		address := conn.RemoteAddr().String()
		_, err := conn.Write(syn)
		if err != nil {
			log.Error().Str("address", address).Err(err).Msg("could not write syn packet")
			conn.Close()
			continue
		}
		_, err = conn.Read(ack)
		if err != nil {
			log.Error().Str("address", address).Err(err).Msg("could not read ack packet")
			conn.Close()
			continue
		}
		networkIn := syn[:len(network)]
		if !bytes.Equal(networkIn, network) {
			log.Error().Str("address", address).Bytes("network", network).Bytes("network_in", networkIn).Msg("network mismatch")
			conn.Close()
			continue
		}
		nonceIn := syn[len(network):]
		if bytes.Equal(nonceIn, nonce) {
			log.Error().Str("address", address).Bytes("nonce", nonce).Msg("identical nonce")
			conn.Close()
			continue
		}
		output <- conn
	}
}
