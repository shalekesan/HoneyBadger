/*
 *    events.go - HoneyBadger core library for detecting TCP attacks.
 *
 *    Copyright (C) 2014  David Stainton
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU General Public License as published by
 *    the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *
 *    You should have received a copy of the GNU General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package types

import (
	"sync"
	"time"
)

type Event struct {
	Type          string
	Flow          *TcpIpFlow
	Time          time.Time
	HijackSeq     uint32
	HijackAck     uint32
	Payload       []byte
	Overlap       []byte
	StartSequence Sequence
	EndSequence   Sequence
	OverlapStart  int
	OverlapEnd    int
}

type EventWithMutex struct {
	Event *Event
	Mutex sync.Mutex
}

type Logger interface {
	Log(r *Event, l sync.Mutex)
}