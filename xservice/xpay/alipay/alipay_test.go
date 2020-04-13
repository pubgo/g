package alipay

import (
	"testing"

	"log"
)

func TestPreCreate(t *testing.T) {
	//bizContent := new(BizContent)
	//bizContent.Subject = "测试收款"
	//bizContent.OutTradeNo = strconv.Itoa(tools.GetSystemCurTime()) + "386871"
	//bizContent.TotalAmount = 0.01
	//bizContent.ProductCode = "FAST_INSTANT_TRADE_PAY"
	//e, err := json.Marshal(bizContent)
	//if err != nil {
	//	log.Println(err)
	//}
	//res, err := PagePay("2019050664325962", "https://www.baidu.com", "", string(e))
	////res, err := PagePay("2019050664325962", "", "", string(e))
	//log.Printf("%+v", res)

	res, _ := main.TradeQuery("15573900496704971295")
	log.Println(res)
}
