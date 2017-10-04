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
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestSetBalance(t *testing.T) {
	config := Config{}
	balance := 5 * time.Second

	SetBalance(balance)(&config)

	assert.Equal(t, balance, config.balance)
}

func TestSetBook(t *testing.T) {
	config := Config{}
	book := NewSimpleBook()

	SetBook(book)(&config)

	assert.Equal(t, book, config.book)
}

func TestSetCodec(t *testing.T) {
	config := Config{}
	codec := DummyCodec{}

	SetCodec(codec)(&config)

	assert.Equal(t, codec, config.codec)
}

func TestSetDiscovery(t *testing.T) {
	config := Config{}
	discovery := 5 * time.Second

	SetDiscovery(discovery)(&config)

	assert.Equal(t, discovery, config.discovery)
}

func TestSetHeartbeat(t *testing.T) {
	config := Config{}
	heartbeat := 5 * time.Second

	SetHeartbeat(heartbeat)(&config)

	assert.Equal(t, heartbeat, config.heartbeat)
}

func TestSetLog(t *testing.T) {
	config := Config{}
	log, _ := zap.NewDevelopment()

	SetLog(log)(&config)

	assert.Equal(t, log, config.log)
}

func TestSetSubscriber(t *testing.T) {
	config := Config{}
	subscriber := make(chan interface{})

	SetSubscriber(subscriber)(&config)

	assert.ObjectsAreEqual(subscriber, config.subscriber)
}

func TestSetTimeout(t *testing.T) {
	config := Config{}
	timeout := 5 * time.Second

	SetTimeout(timeout)(&config)

	assert.Equal(t, timeout, config.timeout)
}

type DummyCodec struct{}

func (s DummyCodec) Encode(w io.Writer, i interface{}) error {
	return nil
}

func (s DummyCodec) Decode(r io.Reader) (interface{}, error) {
	return 1, nil
}
