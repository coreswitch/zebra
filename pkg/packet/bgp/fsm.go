// Copyright 2017 zebra project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bgp

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	BGP_FSM_IDLE int = iota
	BGP_FSM_CONNECT
	BGP_FSM_ACTIVE
	BGP_FSM_OPENSENT
	BGP_FSM_OPENCONFIRM
	BGP_FSM_ESTABLISHED
)

var BgpState2String = map[int]string{
	BGP_FSM_IDLE:        "Idle",
	BGP_FSM_CONNECT:     "Connect",
	BGP_FSM_ACTIVE:      "Active",
	BGP_FSM_OPENSENT:    "OpenSent",
	BGP_FSM_OPENCONFIRM: "OpenConfirm",
	BGP_FSM_ESTABLISHED: "Established",
}

func BgpStateString(s int) string {
	if str, ok := BgpState2String[s]; ok {
		return str
	} else {
		return "Unknown"
	}
}

const (
	_                                    = iota
	ManualStart                          // 1
	ManualStop                           // 2
	AutomaticStart                       // 3
	ManualStart_with_Passive             // 4
	AutomaticStart_with_Passive          // 5
	AutomaticStart_with_Damp             // 6
	AutomaticStart_with_Damp_and_Passive // 7
	AutomaticStop                        // 8
	ConnectRetryTimer_Expires            // 9
	HoldTimer_Expires                    // 10
	KeepaliveTimer_Expires               // 11
	DelayOpenTimer_Expires               // 12
	IdleHoldTimer_Expires                // 13
	TcpConnection_Valid                  // 14
	Tcp_CR_Invalid                       // 15
	Tcp_CR_Acked                         // 16
	TcpConnectionConfirmed               // 17
	TcpConnectionFails                   // 18
	BGPOpen                              // 19
	BGPOpen_with_DelayOpenTimer_running  // 20
	BGPHeaderErr                         // 21
	BGPOpenMsgErr                        // 22
	OpenCollisionDump                    // 23
	NotifMsgVerErr                       // 24
	NotifMsg                             // 25
	KeepAliveMsg                         // 26
	UpdateMsg                            // 27
	UpdateMsgErr                         // 28
	RouteRefreshMsg                      // 29 - local extention
	CapabilityMsg                        // 30 - local extension
)

var BgpEvent2String = map[int]string{
	ManualStart:                          "ManualStart",
	ManualStop:                           "ManualStop",
	AutomaticStart:                       "AutomaticStart",
	ManualStart_with_Passive:             "ManualStart_with_Passive",
	AutomaticStart_with_Passive:          "AutomaticStart_with_Passive",
	AutomaticStart_with_Damp:             "AutomaticStart_with_Damp",
	AutomaticStart_with_Damp_and_Passive: "AutomaticStart_with_Damp_and_Passive",
	AutomaticStop:                        "AutomaticStop",
	ConnectRetryTimer_Expires:            "ConnectRetryTimer_Expires",
	HoldTimer_Expires:                    "HoldTimer_Expires",
	KeepaliveTimer_Expires:               "KeepaliveTimer_Expires",
	DelayOpenTimer_Expires:               "DelayOpenTimer_Expires",
	IdleHoldTimer_Expires:                "IdleHoldTimer_Expires",
	TcpConnection_Valid:                  "TcpConnection_Valid",
	Tcp_CR_Invalid:                       "Tcp_CR_Invalid",
	Tcp_CR_Acked:                         "Tcp_CR_Acked",
	TcpConnectionConfirmed:               "TcpConnectionConfirmed",
	TcpConnectionFails:                   "TcpConnectionFails",
	BGPOpen:                              "BGPOpen",
	BGPOpen_with_DelayOpenTimer_running: "BGPOpen_with_DelayOpenTimer_running",
	BGPHeaderErr:                        "BGPHeaderErr",
	BGPOpenMsgErr:                       "BGPOpenMsgErr",
	OpenCollisionDump:                   "OpenCollisionDump",
	NotifMsgVerErr:                      "NotifMsgVerErr",
	NotifMsg:                            "NotifMsg",
	KeepAliveMsg:                        "KeepAliveMsg",
	UpdateMsg:                           "UpdateMsg",
	UpdateMsgErr:                        "UpdateMsgErr",
	RouteRefreshMsg:                     "RouteRefreshMsg",
	CapabilityMsg:                       "CapabilityMsg",
}

func BgpEventString(s int) string {
	if str, ok := BgpEvent2String[s]; ok {
		return str
	} else {
		return "Unknown"
	}
}

type Event struct {
	id  int
	msg *BgpMessage
	//error   *BgpError
	ioError error
}

type Fsm struct {
	neighbor         *Neighbor
	state            int
	conn             net.Conn
	idleTimer        *time.Timer
	idleTime         int
	connRetryCounter int
	connRetryTimer   *time.Timer
	connRetryTime    int
	holdTimer        *time.Timer
	holdTime         uint16
	keepaliveTimer   *time.Timer
	keepaliveTime    int
	sendCh           chan *BgpMessage
	eventCh          chan *Event
	doneCh           chan interface{}
	connExit         func()
	readFunc         func()
	writeFunc        func()
	wg               sync.WaitGroup
}

func NewFsm(n *Neighbor) *Fsm {
	fsm := &Fsm{
		neighbor: n,
		state:    BGP_FSM_IDLE,
		sendCh:   make(chan *BgpMessage, 1024),
		eventCh:  make(chan *Event, 1024),
		doneCh:   make(chan interface{}),
	}
	return fsm
}

func ReadFull(conn net.Conn, length int) ([]byte, error) {
	buf := make([]byte, length)
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func NewEvent(e int) *Event {
	return &Event{
		id: e,
	}
}

func NewIoErrorEvent(e int, err error) *Event {
	return &Event{
		id:      e,
		ioError: err,
	}
}

func NewMsgEvent(e int, msg *BgpMessage) *Event {
	return &Event{
		id:  e,
		msg: msg,
	}
}

func (f *Fsm) NotifyAndSendEvent(e int, err error) *Event {
	if notify := err.(*BgpNotification); notify != nil {
		f.SendNotification(notify)
		fmt.Println("NotifyAndSendEvent: f.conn.Close()")
		f.conn.Close()
	}
	return &Event{
		id: e,
	}
}

func (f *Fsm) TcpRead(conn net.Conn) {
	f.readFunc = func() {
		conn.Close()
	}

	f.wg.Add(1)
	go func() {
		defer func() {
			f.readFunc = nil
			f.wg.Done()
		}()

		for {
			// Header read.
			buf, err := ReadFull(conn, BGP_HEADER_LEN)
			if err != nil {
				f.eventCh <- NewIoErrorEvent(TcpConnectionFails, err)
				return
			}

			// Header parse.
			header := &BgpHeader{}
			err = header.DecodeFromBytes(buf)
			if err != nil {
				f.eventCh <- f.NotifyAndSendEvent(BGPHeaderErr, err)
				return
			}

			// Body read.
			buf, err = ReadFull(conn, int(header.Len)-BGP_HEADER_LEN)
			if err != nil {
				f.eventCh <- NewIoErrorEvent(TcpConnectionFails, err)
				return
			}

			// Body parse.
			msg, err := ParseBgpBody(header, buf)
			if err != nil {
				switch header.Type {
				case BGP_MSG_OPEN:
					f.eventCh <- f.NotifyAndSendEvent(BGPOpenMsgErr, err)
				case BGP_MSG_UPDATE:
					f.eventCh <- f.NotifyAndSendEvent(UpdateMsgErr, err)
				case BGP_MSG_NOTIFICATION:
					f.eventCh <- f.NotifyAndSendEvent(AutomaticStop, err)
				case BGP_MSG_KEEPALIVE:
					// Won't generate error, checked at BgpHeader DecodeFromBytes().
				case BGP_MSG_ROUTE_REFRESH, BGP_MSG_ROUTE_REFRESH_OLD:
				case BGP_MSG_CAPABILITY:
				}
				return
			}

			// Dispatch message.
			switch header.Type {
			case BGP_MSG_OPEN:
				f.neighbor.in.open++
				f.eventCh <- NewMsgEvent(BGPOpen, msg)
			case BGP_MSG_KEEPALIVE:
				f.neighbor.in.keepalive++
				f.eventCh <- NewMsgEvent(KeepAliveMsg, msg)
			case BGP_MSG_NOTIFICATION:
				f.neighbor.in.notification++
				f.eventCh <- NewMsgEvent(NotifMsg, msg)
			case BGP_MSG_UPDATE:
				f.neighbor.in.update++
				f.eventCh <- NewMsgEvent(UpdateMsg, msg)
			case BGP_MSG_ROUTE_REFRESH, BGP_MSG_ROUTE_REFRESH_OLD:
				f.neighbor.in.refresh++
				f.eventCh <- NewMsgEvent(RouteRefreshMsg, msg)
			case BGP_MSG_CAPABILITY:
				f.neighbor.in.capability++
				f.eventCh <- NewMsgEvent(CapabilityMsg, msg)
			}
		}
	}()
}

func (f *Fsm) TcpWrite(conn net.Conn) {
	done := make(chan interface{})
	f.writeFunc = func() {
		close(done)
	}

	f.wg.Add(1)
	go func() {
		defer func() {
			f.writeFunc = nil
			f.wg.Done()
		}()

		for {
			select {
			case <-done:
				return
			case msg := <-f.sendCh:
				pkt, err := msg.Serialize()
				if err != nil {
					// XXX other type of error event.
					f.SendEvent(TcpConnectionFails)
					return
				}
				_, err = f.conn.Write(pkt)
				if err != nil {
					f.SendEvent(TcpConnectionFails)
					return
				}
				switch msg.Type() {
				case BGP_MSG_OPEN:
					fmt.Println("[Open:Sent]")
					f.neighbor.out.open++
				case BGP_MSG_KEEPALIVE:
					fmt.Println("[Keepalive:Sent]")
					f.neighbor.out.keepalive++
				case BGP_MSG_NOTIFICATION:
					fmt.Println("[Notification:Sent]")
					f.neighbor.out.notification++
				case BGP_MSG_UPDATE:
					fmt.Println("[Update:Sent]")
					f.neighbor.out.update++
				case BGP_MSG_ROUTE_REFRESH, BGP_MSG_ROUTE_REFRESH_OLD:
					fmt.Println("[Route Refresh:Sent]")
					f.neighbor.out.refresh++
				case BGP_MSG_CAPABILITY:
					fmt.Println("[Capability:Sent]")
					f.neighbor.out.capability++
				}
			}
		}
	}()
}

func (f *Fsm) TcpConnect() func() {
	if f.connExit != nil {
		fmt.Println("TcpConnect() already running, cancel existing one")
		f.connExit()
	}

	ctx, cancel := context.WithCancel(context.Background())
	f.connExit = cancel
	//fmt.Println("Add TcpConnect")

	go func() {
		defer func() {
			//fmt.Println("Del TcpConnect")
		}()
		d := net.Dialer{}
		conn, err := d.DialContext(ctx, "tcp", fmt.Sprintf("%s:%d", f.neighbor.Address(), BGP_PORT))
		if err != nil {
			fmt.Println("TcpConnect err", err)
			f.eventCh <- NewIoErrorEvent(TcpConnectionFails, err)
		} else {
			f.conn = conn
			f.eventCh <- NewEvent(TcpConnectionConfirmed)
		}
	}()
	return nil
}

func (f *Fsm) ConnectRetryTimerStart() {
	f.connRetryTimer = time.AfterFunc(time.Second*time.Duration(f.neighbor.connRetryTime),
		func() {
			f.SendEvent(ConnectRetryTimer_Expires)
			f.connRetryTimer = nil
		},
	)
}

func (f *Fsm) IdleTimerStart() {
	f.idleTimer = time.AfterFunc(time.Second*time.Duration(1),
		func() {
			if f.neighbor.passive {
				f.SendEvent(AutomaticStart_with_Passive)
			} else {
				f.SendEvent(AutomaticStart)
			}
			f.idleTimer = nil
		},
	)
}

func (f *Fsm) KeepaliveTimerStart() {
	f.keepaliveTimer = time.AfterFunc(time.Second*time.Duration(f.keepaliveTime),
		func() {
			f.SendEvent(KeepaliveTimer_Expires)
		},
	)
}

func TimerStop(t *time.Timer) {
	if t != nil {
		t.Stop()
	}
}

func (f *Fsm) ChangeStateTo(s int) {
	fmt.Printf("[%s] -> [%s]\n", BgpStateString(f.state), BgpStateString(s))
	f.state = s
	switch s {
	case BGP_FSM_IDLE:
		f.Stop()
		f.IdleTimerStart()
	case BGP_FSM_CONNECT:
		TimerStop(f.idleTimer)
		TimerStop(f.keepaliveTimer)
	case BGP_FSM_ACTIVE:
		f.TcpDrop()
		TimerStop(f.idleTimer)
		TimerStop(f.keepaliveTimer)
	case BGP_FSM_OPENSENT:
		TimerStop(f.idleTimer)
	case BGP_FSM_OPENCONFIRM:
		TimerStop(f.idleTimer)
	case BGP_FSM_ESTABLISHED:
		if f.keepaliveTime != 0 {
			f.KeepaliveTimerStart()
		} else {
			TimerStop(f.keepaliveTimer)
		}
		TimerStop(f.idleTimer)
		TimerStop(f.connRetryTimer)
	}
}

func (f *Fsm) SendOpen() {
	f.SendPacket(NewBgpOpenMsg(f.neighbor))
}

func (f *Fsm) SendKeepAlive() {
	f.SendPacket(NewBgpKeepAliveMsg())
}

func (f *Fsm) SendNotification(notify *BgpNotification) {
	msg := NewBgpNotificationMsg(notify)
	fmt.Println("[Notification:Send]")
	pkt, _ := msg.Serialize()
	nbytes, err := f.conn.Write(pkt)
	if err != nil {
		fmt.Println("Write err", err)
	}
	fmt.Println("Notification sent", nbytes)
}

func (f *Fsm) DelayOpenTimer() bool {
	return false
}

func (f *Fsm) EventIdle(e *Event) {
	switch e.id {
	case ManualStart, AutomaticStart:
		f.neighbor.connRetryCounter = 0
		f.TcpConnect()
		f.ChangeStateTo(BGP_FSM_CONNECT)
	case ManualStart_with_Passive, AutomaticStart_with_Passive:
		f.neighbor.connRetryCounter = 0
		f.ChangeStateTo(BGP_FSM_ACTIVE)
	}
}

func (f *Fsm) BgpOpenProcess(msg *BgpMessage) {
	m := msg.Body.(*BgpOpen)
	if m == nil {
		log.WithFields(log.Fields{
			"error": "Assignment failed",
		}).Error("BgpOpenProcess:asignment to *BgpOpen failed")
		return
	}
	if f.neighbor.HoldTime() < m.HoldTime {
		f.holdTime = f.neighbor.HoldTime()
	} else {
		f.holdTime = m.HoldTime
	}
	f.keepaliveTime = int(f.holdTime / 3)
}

func (f *Fsm) SendPacket(msg *BgpMessage) {
	f.sendCh <- msg
}

func (f *Fsm) TcpEstablish() {
	fmt.Printf("TcpEstablish(): Start ")
	f.TcpRead(f.conn)
	f.TcpWrite(f.conn)
	fmt.Println(" -> End")
}

func (f *Fsm) TcpDrop() {
	fmt.Printf("TcpDrop(): Start")
	if f.readFunc != nil {
		f.readFunc()
	}
	if f.writeFunc != nil {
		f.writeFunc()
	}
	f.wg.Wait()
	fmt.Println(" -> End")
}

func (f *Fsm) EventConnect(e *Event) {
	switch e.id {
	case ManualStart, AutomaticStart:
		// Ignore.
	case TcpConnectionConfirmed:
		// XXX DelayOpenTimer check.
		f.TcpEstablish()
	case BGPOpen:
		// XXX DelayOpenTimer is running,
		// Stop ConnectRetryTimer, Stop DelayeOpenTimer.
		f.BgpOpenProcess(e.msg)
		f.SendOpen()
		f.SendKeepAlive()
		f.ChangeStateTo(BGP_FSM_OPENCONFIRM)
	case TcpConnectionFails:
		if f.DelayOpenTimer() {
			// XXX Stop DelayOpen
			f.ChangeStateTo(BGP_FSM_ACTIVE)
			f.ConnectRetryTimerStart()
		} else {
			// XXX Stop DelayOpen
			f.ChangeStateTo(BGP_FSM_ACTIVE)
			f.ConnectRetryTimerStart()
		}
	case NotifMsg, BGPHeaderErr:
		f.ChangeStateTo(BGP_FSM_IDLE)
	}
}

func (f *Fsm) EventActive(e *Event) {
	switch e.id {
	case ConnectRetryTimer_Expires:
		fmt.Println("ConnectRetryTimer_Expires occured")
		f.TcpConnect()
		f.ConnectRetryTimerStart()
		f.ChangeStateTo(BGP_FSM_CONNECT)
	case TcpConnectionFails:
		f.ChangeStateTo(BGP_FSM_IDLE)
	}
}

func (f *Fsm) EventOpenSent(e *Event) {
	switch e.id {
	case NotifMsg, TcpConnectionFails, BGPHeaderErr:
		f.ChangeStateTo(BGP_FSM_IDLE)
	}
}

func (f *Fsm) EventOpenConfirm(e *Event) {
	switch e.id {
	case KeepAliveMsg:
		// XXX f.HoldTimerStart()
		f.neighbor.uptime = time.Now()
		f.ChangeStateTo(BGP_FSM_ESTABLISHED)
	case NotifMsg, TcpConnectionFails, BGPHeaderErr:
		f.ChangeStateTo(BGP_FSM_IDLE)
	}
}

func (f *Fsm) EventEstablished(e *Event) {
	switch e.id {
	case NotifMsg, TcpConnectionFails, BGPHeaderErr:
		f.ChangeStateTo(BGP_FSM_IDLE)
	case KeepaliveTimer_Expires:
		f.SendKeepAlive()
		f.KeepaliveTimerStart()
	}
}

func (f *Fsm) Event(e *Event) {
	fmt.Printf("[%s] <-%s\n", BgpStateString(f.state), BgpEventString(e.id))

	switch f.state {
	case BGP_FSM_IDLE:
		f.EventIdle(e)
	case BGP_FSM_CONNECT:
		f.EventConnect(e)
	case BGP_FSM_ACTIVE:
		f.EventActive(e)
	case BGP_FSM_OPENSENT:
		f.EventOpenSent(e)
	case BGP_FSM_OPENCONFIRM:
		f.EventOpenConfirm(e)
	case BGP_FSM_ESTABLISHED:
		f.EventEstablished(e)
	}
}

func (f *Fsm) EventLoop(done chan interface{}) {
	f.neighbor.server.wg.Add(1)

	go func() {
		defer f.neighbor.server.wg.Done()
		for {
			select {
			case e := <-f.eventCh:
				f.Event(e)
			case <-f.doneCh:
				fmt.Println("stop Fsm.EventLoop()")
				f.Stop()
				return
			case <-done:
				fmt.Println("stop Fsm.EventLoop()")
				f.Stop()
				return
			}
		}
	}()
}

func (f *Fsm) SendEvent(e int) {
	f.eventCh <- NewEvent(e)
}

func (f *Fsm) Stop() {
	fmt.Println("Stop(): is called")
	TimerStop(f.idleTimer)
	TimerStop(f.connRetryTimer)
	TimerStop(f.keepaliveTimer)
	if f.connExit != nil {
		f.connExit()
		f.connExit = nil
	}
	f.TcpDrop()
}
