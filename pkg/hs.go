package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/buger/jsonparser"
	"github.com/sirupsen/logrus"

	. "github.com/zulong210220/dfcf/types"
)

func GetAllHsCodes() []*CodeMarket {
	pn := 1
	//pz := 10
	pz := 200
	ret := []*CodeMarket{}

	for {
		uu := "http://push2.eastmoney.com/api/qt/clist/get?pn=%d&pz=%d&po=1&np=1&fltt=2&invt=2&fid=f3&fs=m:0+t:6,m:0+t:13,m:0+t:80,m:1+t:2,m:1+t:23&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f11,f62,f128,f136,f115,f152"

		uu = fmt.Sprintf(uu, pn, pz)
		items := getHsDiffs(uu)

		for _, it := range items {
			ret = append(ret, &CodeMarket{it.Code, it.Market})
		}

		pn++
		time.Sleep(100 * time.Millisecond)
		if atomic.LoadInt64(&hsInvalid)+atomic.LoadInt64(&hsValid) >= atomic.LoadInt64(&hsTotal) {
			logrus.Infof("GetAllHsCodes page:%d hsInvalid:%d hsValid:%d hsTotal:%d", pn, hsInvalid, hsValid, hsTotal)
			break
		}
	}

	logrus.Info("GetAllHsCodes ", len(ret), " ......")

	atomic.StoreInt64(&hsInvalid, 0)
	atomic.StoreInt64(&hsValid, 0)
	atomic.StoreInt64(&hsTotal, 0)
	return ret
}

func getHsDiffs(uu string) []*DfcfListItem {

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

	db := &DfcfListBody{}
	err = json.Unmarshal(body, db)
	if err != nil {
		log.Printf("%s\n", body)
		log.Fatal(err)
	}
	//log.Printf("%+v", db.Data.Diff)
	if db.Data != nil {
		atomic.StoreInt64(&hsTotal, db.Data.Total)
		return convertDiff(db.Data.Diff)
	}
	return nil
}

func convertDiff(values []byte) []*DfcfListItem {
	ret := []*DfcfListItem{}
	jsonparser.ArrayEach([]byte(values), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		dsi := &DfcfListItem{}
		val, dt, _, err := jsonparser.Get(value, "f2")
		if dt == jsonparser.String {
			atomic.AddInt64(&hsInvalid, 1)
			return
		}
		dsi.Name, err = jsonparser.GetString(value, "f14")
		dsi.Code, err = jsonparser.GetString(value, "f12")

		if strings.Index(dsi.Name, "*") >= 0 ||
			strings.Index(dsi.Name, "退") == 0 ||
			strings.Index(dsi.Name, "S") >= 0 {
			atomic.AddInt64(&hsInvalid, 1)
			return
		}

		if strings.Index(dsi.Code, "688") == 0 ||
			strings.Index(dsi.Code, "689") == 0 {
			atomic.AddInt64(&hsInvalid, 1)
			return
		}

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
		atomic.AddInt64(&hsValid, 1)
	})
	return ret
}

// get all kline history
func GetHisItems(cm *CodeMarket) []*DfcfHisDay {
	code := cm.Code
	market := cm.Market
	uu := "http://push2his.eastmoney.com/api/qt/stock/kline/get?fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61&klt=101&fqt=1&secid=%d.%s&beg=0&end=20500000"

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

func getDiffs(uu string) ([]*DfcfHisDay, string) {

	// log.Println(uu)
	req, err := http.NewRequest("GET", uu, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := cli.Do(req)
	if err != nil {
		log.Print(err)
		return nil, ""
	}
	defer resp.Body.Close()
	// log.Println("Reading body...")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return nil, ""
	}

	//log.Printf("%s\n", body)
	db := &DfcfHis{}
	err = json.Unmarshal(body, db)
	if err != nil {
		log.Print(err)
		return nil, ""
	}

	if db.Data == nil {
		return nil, ""
	}

	dd := db.Data
	ret := []*DfcfHisDay{}
	for _, line := range dd.Klines {
		fields := strings.Split(line, ",")
		// "2021-03-26,3.67,3.80,3.95,3.64,544498,206862632.00,8.40,2.98,0.11,6.24",
		// 日期 开盘 收盘 最高 最低 成交量 成交额 振幅 涨幅 涨跌额 还手
		it := &DfcfHisDay{}
		it.Code = dd.Code
		it.Market = dd.Market

		if len(fields) > 0 {
			it.Date = fields[0]
		}
		if len(fields) > 1 {
			f, err := strconv.ParseFloat(fields[1], 64)
			if err != nil {
				log.Fatal(err)
			}
			it.Open = f
		}
		if len(fields) > 2 {
			f, err := strconv.ParseFloat(fields[2], 64)
			if err != nil {
				log.Fatal(err)
			}
			it.Close = f
		}
		if len(fields) > 3 {
			f, err := strconv.ParseFloat(fields[3], 64)
			if err != nil {
				log.Fatal(err)
			}
			it.High = f
		}
		if len(fields) > 4 {
			f, err := strconv.ParseFloat(fields[4], 64)
			if err != nil {
				log.Fatal(err)
			}
			it.Low = f
		}
		// "2021-03-26,3.67,3.80,3.95,3.64,544498,206862632.00,8.40,2.98,0.11,6.24",
		// 日期 开盘 收盘 最高 最低 成交量 成交额 振幅 涨幅 涨跌额 还手
		if len(fields) > 5 {
			f, err := strconv.ParseInt(fields[5], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			it.Volume = f
		}
		if len(fields) > 6 {
			f, err := strconv.ParseFloat(fields[6], 64)
			if err != nil {
				log.Fatal(err)
			}
			it.Amount = f
		}
		if len(fields) > 7 {
			f, err := strconv.ParseFloat(fields[7], 64)
			if err != nil {
				log.Fatal(err)
			}
			if f < 99999999.99 && f > -99999999.99 {
				it.Amplitude = f
			} else {
				it.Amplitude = 0
			}
		}
		if len(fields) > 8 {
			f, err := strconv.ParseFloat(fields[8], 64)
			if err != nil {
				log.Fatal(err)
			}
			if f < 99999999.99 && f > -99999999.99 {
				it.Percent = f
			} else {
				it.Percent = 0
			}
		}
		if len(fields) > 9 {
			f, err := strconv.ParseFloat(fields[9], 64)
			if err != nil {
				log.Fatal(err)
			}
			//if f < 99999999.99 && f > -99999999.99 {
			it.Chg = f
			//} else {
			//	it.Chg = 0
			//}
		}
		if len(fields) > 10 {
			f, err := strconv.ParseFloat(fields[10], 64)
			if err != nil {
				log.Fatal(err)
			}
			it.Turnoverrate = f
		}
		ret = append(ret, it)
	}

	return ret, dd.Name
}
