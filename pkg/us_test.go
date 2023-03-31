package pkg

import (
	"fmt"
	"strings"
	"testing"
)

func TestCM(t *testing.T) {
	cms := GetAllUsCodes()
	cmm := map[int64]string{}

	for _, cm := range cms {
		cmm[cm.Market] = cm.Code
	}
	fmt.Println(cmm)
}

// MITQ A 纽交所 107
// NVFY O 纳斯达克 105
// MATX N 纽交所 106

func TestCCM(t *testing.T) {
	cms := GetUsChinese()
	fmt.Println(len(cms))
	fmt.Println(cms)
}

func TestFilter(t *testing.T) {
	codes := `74#CMND
74#ASNS
74#OKYO
74#VWEWW
74#DRMA
74#LNSR
74#TNON
74#LIDR
74#TOI
74#BEAT
`

	allCodes := strings.Split(codes, "#")
	fmt.Println(allCodes[0])
	fmt.Println(allCodes[1])
	return
	tmp := []string{}

	for _, code := range allCodes {
		if code == "74" {
			continue
		}
		tmp = append(tmp, code)
	}

	allCodes = tmp

	cms := GetUsChinese()

	ret := []string{}

	ucs := []string{}
	for _, code := range allCodes {
		isUC := false
		for _, cm := range cms {
			if strings.Contains(code, cm.Code) {
				isUC = true
				break
			}
		}

		if !isUC {
			ret = append(ret, code)
		}

		if isUC {
			ucs = append(ucs, code)
		}
	}

	fmt.Println(ret)
	fmt.Println()
	fmt.Println(ucs)
}
