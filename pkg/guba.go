package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	. "github.com/zulong210220/dfcf/types"
)

/*
curl 'http://gbapi.eastmoney.com/webexparticlelist/api/Article/ArticleListForWebExplore?code=600379&plat=web&version=300&product=guba&type=0&sorttype=0&ps=40&p=3' | pbcopy

http://192.168.1.4:9200/_cat/nodes?pretty
*/

var (
	indexName = "guba"
	servers   = []string{"http://192.168.1.4:9200/"}
)

func getCode(code string) {
	ps := 100
	p := 1
	count := 0
	uu := "http://gbapi.eastmoney.com/webexparticlelist/api/Article/ArticleListForWebExplore?code=%s&plat=web&version=300&product=guba&type=0&sorttype=0&ps=%d&p=%d"
	for {
		nn := getList(fmt.Sprintf(uu, code, ps, p))
		n := len(nn)
		log.Println("getList", code, n)
		count += n
		if n < 1 {
			break
		}
		p++
		for _, v := range nn {
			egr := &EsGubaReply{}
			egr.QCode = code
			egr.GubaRe = v
		}
		//time.Sleep(50 * time.Millisecond)
		log.Println("getCode", code, n)
	}
	log.Println("getCode-----", code, count)
}

func getList(uu string) []*GubaRe {
	//uu := "http://gbapi.eastmoney.com/webexparticlelist/api/Article/ArticleListForWebExplore?code=600379&plat=web&version=300&product=guba&type=0&sorttype=0&ps=100&p=3"
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
	db := &DfcfGubaBody{}
	err = json.Unmarshal(body, db)
	if err != nil {
		log.Fatal(err)
	}
	return db.Re
}
