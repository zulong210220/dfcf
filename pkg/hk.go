package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/buger/jsonparser"
	"github.com/sirupsen/logrus"
	. "github.com/zulong210220/dfcf/types"
)

func GetAllHkCodes() []*CodeMarket {
	pn := 1
	pz := 200
	ret := []*CodeMarket{}

	for {
		uu := "http://push2.eastmoney.com/api/qt/clist/get?pn=%d&pz=%d&po=1&np=1&fltt=2&invt=2&fid=f3&fs=m:128+t:3,m:128+t:4,m:128+t:1,m:128+t:2&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f19,f20,f21,f23,f24,f25,f26,f22,f33,f11,f62,f128,f136,f115,f152"

		uu = fmt.Sprintf(uu, pn, pz)
		items := getHkDiffs(uu)

		for _, it := range items {
			ret = append(ret, &CodeMarket{it.Code, it.Market})
		}

		pn++
		time.Sleep(100 * time.Millisecond)
		if atomic.LoadInt64(&hkInvalid)+atomic.LoadInt64(&hkValid) >= atomic.LoadInt64(&hkTotal) {
			//logrus.Info("GetAllHkCodes ", atomic.LoadInt64(&hkInvalid), atomic.LoadInt64(&hkValid), atomic.LoadInt64(&hkTotal))
			logrus.Infof("GetAllHkCodes hkInvalid:%d hkValid:%d hkTotal:%d", hkInvalid, hkValid, hkTotal)
			break
		}
	}

	logrus.Info("GetAllHkCodes", "len:", len(ret))
	atomic.StoreInt64(&hkInvalid, 0)
	atomic.StoreInt64(&hkValid, 0)
	atomic.StoreInt64(&hkTotal, 0)

	return ret
}

func getHkDiffs(uu string) []*DfcfListItem {
	req, err := http.NewRequest("GET", uu, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// log.Println("Reading body...")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// log.Printf("%s\n", body)
	db := &DfcfListBody{}
	err = json.Unmarshal(body, db)
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("%+v", db.Data.Diff)
	if db.Data != nil {
		atomic.StoreInt64(&hkTotal, db.Data.Total)
		return convertHkDiff(db.Data.Diff)
	}
	return nil
}

func convertHkDiff(values []byte) []*DfcfListItem {
	ret := []*DfcfListItem{}
	jsonparser.ArrayEach([]byte(values), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		dsi := &DfcfListItem{}
		val, dt, _, err := jsonparser.Get(value, "f2")
		if dt == jsonparser.String {
			atomic.AddInt64(&hkInvalid, 1)
			return
		}
		dsi.Name, err = jsonparser.GetString(value, "f14")
		dsi.Code, err = jsonparser.GetString(value, "f12")

		dsi.Current, err = strconv.ParseFloat(string(val), 64)
		dsi.Percent, err = jsonparser.GetFloat(value, "f3")
		dsi.Chg, err = jsonparser.GetFloat(value, "f4")
		dsi.Volume, err = jsonparser.GetInt(value, "f5")
		dsi.Amount, err = jsonparser.GetFloat(value, "f6")
		dsi.Amplitude, err = jsonparser.GetFloat(value, "f7")
		dsi.Turnoverrate, err = jsonparser.GetFloat(value, "f8")
		dsi.Market, err = jsonparser.GetInt(value, "f13")
		dsi.High, err = jsonparser.GetFloat(value, "f15")
		dsi.Low, err = jsonparser.GetFloat(value, "f16")
		dsi.Open, err = jsonparser.GetFloat(value, "f17")

		ret = append(ret, dsi)
		atomic.AddInt64(&hkValid, 1)
	})
	return ret
}

// "2021-03-26,3.67,3.80,3.95,3.64,544498,206862632.00,8.40,2.98,0.11,6.24",
// 日期 开盘 收盘 最高 最低 成交量 成交额 振幅 涨幅 涨跌额 还手
func GetHkHisItems(cm *CodeMarket) []*DfcfHisDay {
	code := cm.Code
	market := cm.Market
	uu := "http://push2his.eastmoney.com/api/qt/stock/kline/get?secid=%d.%s&fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61&klt=101&fqt=1&beg=0&end=20500101&smplmt=1194&lmt=1000000"

	uu = fmt.Sprintf(uu, market, code)
	var items []*DfcfHisDay
	var name string

	retry := 0
	for retry < 3 {
		items, name = getDiffs(uu)
		if len(items) > 0 {
			break
		}
		retry++
		time.Sleep(500 * time.Millisecond)
	}
	//log.Println(uu, len(items))
	if len(items) == 0 {
		log.Println("code", code, uu, name)
		return items
	}
	sort.Sort(DHDs(items))
	return items
}

// http://emweb.securities.eastmoney.com/PC_HKF10/CompanyProfile/PageAjax?code=00023
