package pkg

import (
	"net"
	"net/http"
	"time"

	. "github.com/zulong210220/dfcf/types"
)

var (
	cli = &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
)

var (
	hsValid   int64
	hsTotal   int64
	hsInvalid int64
)

var (
	mapCMFunc = map[string]func() []*CodeMarket{
		MarketHs: GetAllHsCodes,
		MarketHk: GetAllHkCodes,
		MarketUs: GetAllUsCodes,
	}
)

func GetCMFunc(market string) func() []*CodeMarket {
	f, ok := mapCMFunc[market]
	if ok {
		return f
	}
	return GetAllHsCodes
}

type DHDs []*DfcfHisDay

func (s DHDs) Len() int {
	return len(s)
}

func (s DHDs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s DHDs) Less(i, j int) bool {
	return s[i].Date > s[j].Date
}

// hk

var (
	hkValid   int64
	hkTotal   int64
	hkInvalid int64
)

// us

var (
	usValid   int64
	usTotal   int64
	usInvalid int64
)
