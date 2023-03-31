package types

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type DfcfListItem struct {
	Current      float64 `json:"f2"`
	Percent      float64 `json:"f3"`
	Chg          float64 `json:"f4"`
	Volume       int64   `json:"f5"` // 手
	Amount       float64 `json:"f6"`
	Amplitude    float64 `json:"f7"`
	High         float64 `json:"f15"`
	Low          float64 `json:"f16"`
	Open         float64 `json:"f17"`
	Turnoverrate float64 `json:"f8"`
	Name         string  `json:"f14"`
	Code         string  `json:"f12"`
	Market       int64   `json:"f13"`
}

type DfcfListBody struct {
	Data *DfcfListData `json:"data"`
}

type DfcfListData struct {
	Total int64           `json:"total"`
	Diff  json.RawMessage `json:"diff"`
}

type DfcfHisData struct {
	Code      string   `json:"code"`
	Market    int      `json:"market"`
	Name      string   `json:"name"`
	Decimal   int      `json:"decimal"` //
	Dktotal   int      `json:"dktotal"`
	PreKPrice float64  `json:"preKPrice"`
	Klines    []string `json:"klines"`
}

type DfcfHis struct {
	Rc   int          `json:"rc"`
	Rt   int          `json:"rt"`
	Svr  int          `json:"svr"`
	Full int          `json:"full"`
	Data *DfcfHisData `json:"data"`
}

//

type CodeMarket struct {
	Code   string
	Market int64
}

func (cm *CodeMarket) String() string {
	if cm == nil {
		return ""
	}
	return strconv.FormatInt(cm.Market, 10) + cm.Code
}

func GetHsCMKey(cm *CodeMarket) string {
	return fmt.Sprintf("%d%s", cm.Market, cm.Code)
}

//

type DfcfHisDay struct {
	Market       int     `json:"market"`
	Code         string  `json:"code"`
	Volume       int64   `json:"volume"`
	Open         float64 `json:"open"`
	High         float64 `json:"high"`
	Low          float64 `json:"low"`
	Close        float64 `json:"close"`
	Chg          float64 `json:"chg"`
	Percent      float64 `json:"percent"`
	Turnoverrate float64 `json:"turnoverrate"`
	Amount       float64 `json:"amount"`
	Amplitude    float64 `json:"amplitude"`
	Date         string  `json:"date"`
}

func (dhd *DfcfHisDay) MarshalBinary() (data []byte, err error) {
	return json.Marshal(dhd)
}

func (dhd *DfcfHisDay) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, dhd)
}

//

type DfcfUsChinese struct {
	Rc     int    `json:"rc"`
	Rt     int    `json:"rt"`
	Svr    int    `json:"svr"`
	Lt     int    `json:"lt"`
	Full   int    `json:"full"`
	Dlmkts string `json:"dlmkts"`
	Data   struct {
		Total int `json:"total"`
		Diff  []struct {
			//			F1  int     `json:"f1"`
			//			F2  float64 `json:"f2"`
			//			F3  float64 `json:"f3"`
			//			F4  float64 `json:"f4"`
			//			F5  int     `json:"f5"`
			//			F6  float64 `json:"f6"`
			//			F7  float64 `json:"f7"`
			//			F8  float64 `json:"f8"`
			//			F9  float64 `json:"f9"`
			//			F10 float64 `json:"f10"`
			//			F11 float64 `json:"f11"`
			F12 string `json:"f12"` // code
			F13 int64  `json:"f13"` // market
			F14 string `json:"f14"` // 全称
			//			F15  float64 `json:"f15"`
			//			F16  float64 `json:"f16"`
			//			F17  float64 `json:"f17"`
			//			F18  float64 `json:"f18"`
			//			F20  int     `json:"f20"`
			//			F21  int     `json:"f21"`
			//			F22  float64 `json:"f22"`
			//			F23  float64 `json:"f23"`
			//			F24  float64 `json:"f24"`
			//			F25  float64 `json:"f25"`
			//			F26  int     `json:"f26"`
			//			F33  float64 `json:"f33"`
			//			F62  float64 `json:"f62"`
			//			F115 float64 `json:"f115"`
			//			F128 string  `json:"f128"`
			//			F140 string  `json:"f140"`
			//			F141 string  `json:"f141"`
			//			F136 string  `json:"f136"`
			//			F152 int     `json:"f152"`
		} `json:"diff"`
	} `json:"data"`
}
