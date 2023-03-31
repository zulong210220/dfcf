package pkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	. "github.com/zulong210220/dfcf/types"
)

func TestHL(t *testing.T) {
	items := getAllHs()
	ret := []string{}
	for _, item := range items {
		if item.High == item.Open && item.Open > item.Low {
			ret = append(ret, fmt.Sprintf("%d%s", item.Market, item.Code))
		}
	}

	fmt.Println(strings.Join(ret, "\n"))
}

func getAllHs() []*DfcfListItem {
	pn := 1
	//pz := 10
	pz := 200
	ret := []*DfcfListItem{}

	for {
		uu := "http://push2.eastmoney.com/api/qt/clist/get?pn=%d&pz=%d&po=1&np=1&fltt=2&invt=2&fid=f3&fs=m:0+t:6,m:0+t:13,m:0+t:80,m:1+t:2,m:1+t:23&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f11,f62,f128,f136,f115,f152"

		uu = fmt.Sprintf(uu, pn, pz)
		items := getHsDiffs(uu)

		ret = append(ret, items...)

		pn++
		if len(items) == 0 {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	logrus.Info("GetAllHs ", len(ret), " ......")

	return ret
}

// curl 'http://push2his.eastmoney.com/api/qt/stock/trends2/get?fields1=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f11,f12,f13&fields2=f51,f52,f53,f54,f55,f56,f57,f58&secid=0.300561&ndays=1&iscr=1&iscca=0'   -H 'Accept: */*' \
//   -H 'Accept-Language: zh-TW,zh;q=0.9,en-US;q=0.8,en;q=0.7' \
//   -H 'Cookie: qgqp_b_id=7a38c4828b3464cffcfccf9711dbada8; st_si=78289256805040; st_asi=delete; HAList=ty-0-300561-%u6C47%u91D1%u79D1%u6280; st_pvi=25309417537189; st_sp=2023-03-13%2021%3A28%3A22; st_inirUrl=http%3A%2F%2Fquote.eastmoney.com%2Fchanges%2Fstocks%2Fsz300561.html; st_sn=4; st_psi=20230313213105730-113200301201-6621244810' \
//   -H 'Proxy-Connection: keep-alive' \
//   -H 'Referer: http://quote.eastmoney.com/' \
//   -H 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36' \
//   --compressed \
//   --insecure \

func doCurl(cm string) []byte {
	requestURL := fmt.Sprintf("http://push2his.eastmoney.com/api/qt/stock/trends2/get?fields1=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f11,f12,f13&fields2=f51,f52,f53,f54,f55,f56,f57,f58&secid=%s&ndays=1&iscr=1&iscca=0", cm)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("'Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-TW,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Cookie", "qgqp_b_id=7a38c4828b3464cffcfccf9711dbada8; st_si=78289256805040; st_asi=delete; HAList=ty-0-300561-%u6C47%u91D1%u79D1%u6280; st_pvi=25309417537189; st_sp=2023-03-13%2021%3A28%3A22; st_inirUrl=http%3A%2F%2Fquote.eastmoney.com%2Fchanges%2Fstocks%2Fsz300561.html; st_sn=4; st_psi=20230313213105730-113200301201-6621244810")
	req.Header.Set("Proxy-Connection", "keep-alive")
	req.Header.Set("Referer", "http://quote.eastmoney.com/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")

	resp, err := cli.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return []byte{}
	}
	defer resp.Body.Close()
	//log.Println("Reading body...", requestURL)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("client: body error making http request: %s\n", err)
		return []byte{}
	}
	return body
}

func TestDay(t *testing.T) {
	day := time.Now().Format("2006-01-02")
	f, err := os.Create(day + ".txt")

	if err != nil {
		os.Exit(1)
	}

	defer f.Close()

	items := getAllHs()
	for _, item := range items {
		cm := fmt.Sprintf("%d.%s", item.Market, item.Code)
		body := doCurl(cm)
		k := 0
		for len(body) == 0 {
			body = doCurl(cm)
			k++
			if k > 3 {
				break
			}
		}
		if k > 3 {
			continue
		}
		f.Write(body)
		f.Write([]byte{'\n'})
	}
}
