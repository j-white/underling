package underlinglib

import (
	"github.com/tatsushid/go-fastping"
	"net"
	"errors"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const DefaultPingRetries = 2
const DefaultPingTimeout = 800
const DefaultPacketSize = 64

type Pinger interface {
	Ping(host string, retries int, timeout int) (rtt time.Duration, err error)
}

type DefaultPinger struct {
}

type response struct {
	addr *net.IPAddr
	rtt  time.Duration
}

func (pinger DefaultPinger) Ping(host string, retries int, timeout int) (rtt time.Duration, err error) {
	p := fastping.NewPinger()
	p.Network("udp")

	netProto := "ip4:icmp"
	if strings.Index(host, ":") != -1 {
		netProto = "ip6:ipv6-icmp"
	}
	ra, err := net.ResolveIPAddr(netProto, host)
	if err != nil {
		return 0, err
	}
	p.AddIPAddr(ra)

	onRecv, onIdle := make(chan *response), make(chan bool)
	p.OnRecv = func(addr *net.IPAddr, t time.Duration) {
		onRecv <- &response{addr: addr, rtt: t}
	}
	p.OnIdle = func() {
		onIdle <- true
	}

	p.MaxRTT = time.Second
	p.RunLoop()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	rtt, err = 0, errors.New("Unknown error")
loop:
	for {
		select {
		case <-c:
			rtt, err = 0, errors.New("Interrupted.")
			break loop
		case <-onRecv:
			rtt, err = 1, nil
			break loop
		case <-onIdle:
			rtt, err = 0, nil
			break loop
		case <-p.Done():
			if err != nil {
				rtt, err = 0, err
			} else {
				rtt, err = 0, errors.New("Unknown error")
			}
		}
	}
	signal.Stop(c)
	p.Stop()
	return rtt, err
}

func NewPinger() Pinger {
	return DefaultPinger{}
}
