package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	. "github.com/zulong210220/dfcf/types"
)

func GetUsChinese() []*CodeMarket {
	ret := []*CodeMarket{}
	total := 0
	ut := "http://push2.eastmoney.com/api/qt/clist/get?pn=%d&pz=100&po=1&np=1&fltt=2&invt=2&wbp2u=|0|0|0|web&fid=f3&fs=b:MK0201&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f26,f22,f33,f11,f62,f128,f136,f115,f152"
	pn := 1
	for {
		uu := fmt.Sprintf(ut, pn)
		duc := getUsChinese(uu)

		if len(duc.Data.Diff) == 0 {
			break
		}

		if total < duc.Data.Total {
			total = duc.Data.Total
		}

		for _, d := range duc.Data.Diff {
			ret = append(ret, &CodeMarket{
				Code:   d.F12,
				Market: d.F13,
			})
		}

		pn++
	}

	return ret
}

//  curl 'http://push2.eastmoney.com/api/qt/clist/get?pn=1&pz=20&po=1&np=1&fltt=2&invt=2&wbp2u=|0|0|0|web&fid=f3&fs=b:MK0201&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f26,f22,f33,f11,f62,f128,f136,f115,f152&_=1679120261944' \
//  -H 'Accept: */*' \
//  -H 'Accept-Language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6' \
//  -H 'Cookie: qgqp_b_id=97b97dc94e713c4f9fe105f8159b7bf5; st_si=87928398163309; st_asi=delete; HAList=ty-105-BIDU-%u767E%u5EA6%2Cty-106-BABA-%u963F%u91CC%u5DF4%u5DF4; st_pvi=18699588424718; st_sp=2023-03-18%2014%3A00%3A00; st_inirUrl=http%3A%2F%2Fquote.eastmoney.com%2Fcenter%2Fmgsc.html; st_sn=12; st_psi=2023031814174214-113200301321-2077806659' \
//  -H 'Proxy-Connection: keep-alive' \
//  -H 'Referer: http://quote.eastmoney.com/' \
//  -H 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.69' \

func getUsChinese(uu string) *DfcfUsChinese {
	db := &DfcfUsChinese{}
	req, err := http.NewRequest("GET", uu, nil)
	if err != nil {
		log.Fatal(err)
		return db
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Proxy-Connection", "keep-alive")
	req.Header.Set("Referer", "http://quote.eastmoney.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.69")

	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
		return db
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return db
	}

	err = json.Unmarshal(body, db)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
