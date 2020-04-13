package alipay

import (
	"github.com/pubgo/g/pkg/ethutil"
	"github.com/pubgo/g/pkg/timeutil"
	"github.com/pubgo/g/xerror"
	"net/http"
	"net/url"
	"time"

	"fmt"
	"io/ioutil"
	"log"

	"encoding/json"
)

func PagePay(returnUrl, notifyUrl, bizContent string) (results *url.URL, err error) {
	var data = url.Values{}
	data.Add("app_id", main.app_id)
	data.Add("method", "alipay.trade.page.pay")
	if returnUrl != "" {
		data.Add("return_url", returnUrl)
	}
	if notifyUrl != "" {
		data.Add("notify_url", notifyUrl)
	}
	data.Add("format", "json")
	data.Add("charset", "utf-8")
	data.Add("sign_type", "RSA2")
	data.Add("timestamp", time.Now().Format("2006-01-02 03:04:05"))
	data.Add("version", "1.0")
	data.Add("biz_content", bizContent)
	data.Add("sign", main.sign(data))
	results, err = url.Parse("https://openapi.alipay.com/gateway.do" + "?" + data.Encode())
	if err != nil {
		return nil, err
	}
	return results, err
}

type QueryRep struct {
	Data AlipayTradeQueryResponse `json:"alipay_trade_query_response"`
	Sign string                   `json:"sign"`
}

type AlipayTradeQueryResponse struct {
	Code           string `json:"code"`
	Msg            string `json:"msg"`
	SubCode        string `json:"sub_code"`
	SubMsg         string `json:"sub_msg"`
	OutTradeNo     string `json:"out_trade_no"`
	BuyerPayAmount string `json:"buyer_pay_amount"`
	InvoiceAmount  string `json:"invoice_amount"`
	PointAmount    string `json:"point_amount"`
	ReceiptAmount  string `json:"receipt_amount"`
}

func TradeQuery(OutTradeNo string) (string, error) {

	res := new(main.BizContent)
	res.OutTradeNo = OutTradeNo
	e, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	var data = url.Values{}
	data.Add("app_id", main.app_id)
	data.Add("method", "alipay.trade.query")
	data.Add("format", "json")
	data.Add("charset", "utf-8")
	data.Add("sign_type", "RSA2")
	data.Add("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	data.Add("version", "1.0")
	data.Add("biz_content", string(e))
	data.Add("sign", main.sign(data))
	resp, err := http.PostForm("https://openapi.alipay.com/gateway.do", data)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	rep := &QueryRep{}
	err = json.Unmarshal(body, rep)
	if err != nil {
		log.Println("response struct err :", string(body))
	}
	return rep.Data.Msg, nil
}

func GetTradeNotification(req *http.Request) (noti *main.TradeNotification, err error) {
	if req == nil {
		return nil, fmt.Errorf("request 参数不能为空")
	}

	noti = &main.TradeNotification{}
	noti.AppId = req.FormValue("app_id")
	noti.AuthAppId = req.FormValue("auth_app_id")
	noti.NotifyId = req.FormValue("notify_id")
	noti.NotifyType = req.FormValue("notify_type")
	noti.NotifyTime = req.FormValue("notify_time")
	noti.TradeNo = req.FormValue("trade_no")
	noti.TradeStatus = req.FormValue("trade_status")
	noti.TotalAmount = req.FormValue("total_amount")
	noti.ReceiptAmount = req.FormValue("receipt_amount")
	noti.InvoiceAmount = req.FormValue("invoice_amount")
	noti.BuyerPayAmount = req.FormValue("buyer_pay_amount")
	noti.SellerId = req.FormValue("seller_id")
	noti.SellerEmail = req.FormValue("seller_email")
	noti.BuyerId = req.FormValue("buyer_id")
	noti.BuyerLogonId = req.FormValue("buyer_logon_id")
	noti.FundBillList = req.FormValue("fund_bill_list")
	noti.Charset = req.FormValue("charset")
	noti.PointAmount = req.FormValue("point_amount")
	noti.OutTradeNo = req.FormValue("out_trade_no")
	noti.OutBizNo = req.FormValue("out_biz_no")
	noti.GmtCreate = req.FormValue("gmt_create")
	noti.GmtPayment = req.FormValue("gmt_payment")
	noti.GmtRefund = req.FormValue("gmt_refund")
	noti.GmtClose = req.FormValue("gmt_close")
	noti.Subject = req.FormValue("subject")
	noti.Body = req.FormValue("body")
	noti.RefundFee = req.FormValue("refund_fee")
	noti.Version = req.FormValue("version")
	noti.SignType = req.FormValue("sign_type")
	noti.Sign = req.FormValue("sign")
	noti.PassbackParams = req.FormValue("passback_params")
	noti.VoucherDetailList = req.FormValue("voucher_detail_list")

	ok, err := main.verifySign(req.Form, []byte(main.alipay_public_key))
	if ok == false {
		return nil, err
	}
	return noti, err
}

func AckNotification(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write(main.kSuccess)
}

func TradeFinish(OutTradeNo string, noti *main.TradeNotification) (err error) {
	defer xerror.RespErr(&err)

	wheres := make(map[string]interface{}, 0)
	wheres["bill_id"] = OutTradeNo
	updates := make(map[string]interface{}, 0)
	updates["pay_id"] = noti.TradeNo
	updates["updated_at"] = timeutil.GetSystemCurTime()
	updates["bill_info"] = xerror.PanicErr(ethutil.ToString(noti))
	updates["status"] = 1
	return
}
