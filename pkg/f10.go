package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
	. "github.com/zulong210220/dfcf/types"
)

func GetF10(cm *CodeMarket) {
	if cm == nil {
		return
	}
	code := ""
	if cm.Market == 0 {
		code = "SZ" + cm.Code
	} else if cm.Market == 1 {
		code = "SH" + cm.Code
	}
	uu := "http://f10.eastmoney.com/CompanySurvey/CompanySurveyAjax?code=" + code

	req, err := http.NewRequest("GET", uu, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	// log.Println("Reading body...")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	f10 := &F10Data{}
	err = json.Unmarshal(body, f10)
	if err != nil {
		logrus.Fatal("F10", err)
		return
	}
	fmt.Println(uu)
	fmt.Println(f10.Jbzl)
}

// http://f10.eastmoney.com/CompanySurvey/CompanySurveyAjax?code=SZ002023

//

// us
// http://emweb.eastmoney.com/pc_usf10/CompanyInfo/PageAjax?fullCode=TNXP.O

// MITQ A 纽交所 107
// NVFY O 纳斯达克 105
// MATX N 纽交所 106

var (
	codeMarketMap = map[int64]string{
		105: "O",
		106: "N",
		107: "A",
	}
)

func getUsMarket(c int64) string {
	r := codeMarketMap[c]
	if r != "" {
		return r
	}
	return "O"
}

func GetUsF10(cm *CodeMarket) []byte {
	o := getUsMarket(cm.Market)
	uu := fmt.Sprintf("http://emweb.eastmoney.com/pc_usf10/CompanyInfo/PageAjax?fullCode=%s.%s", cm.Code, o)
	req, err := http.NewRequest("GET", uu, nil)
	if err != nil {
		logrus.Fatal(err)
	}
	resp, err := cli.Do(req)
	if err != nil {
		logrus.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatal(err)
	}
	return body
}
