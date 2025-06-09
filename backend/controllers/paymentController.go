package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	merchantID = "3002607"
	hashKey    = "pwFHCqoQZGmho4w6"
	hashIV     = "EkRm7iFT261dpevs"
)

func Payment(c *gin.Context) {
	now := time.Now()
	tradeNo := "VIP" + now.Format("20060102150405")
	tradeDate := now.Format("2006/01/02 15:04:05")

	form := url.Values{}
	form.Add("MerchantID", merchantID)
	form.Add("MerchantTradeNo", tradeNo)
	form.Add("MerchantTradeDate", tradeDate)
	form.Add("PaymentType", "aio")
	form.Add("TotalAmount", "99")
	form.Add("TradeDesc", "VIP會員升級")
	form.Add("ItemName", "升級會員方案")
	form.Add("ReturnURL", "https://e312-2401-e180-8820-8783-ed6c-8b9f-c725-1ce2.ngrok-free.app/ecpay-return")
	form.Add("ChoosePayment", "Credit")
	form.Add("ClientBackURL", "http://localhost:3000/payment-success") // 可自訂完成頁
	form.Add("NotifyURL", "https://e312-2401-e180-xxxx-xxxx.ngrok-free.app/ecpay-notify")

	checkMac := generateCheckMacValue(form)
	form.Add("CheckMacValue", checkMac)

	html := `
		<form id="ecpay" method="post" action="https://payment-stage.ecpay.com.tw/Cashier/AioCheckOut/V5">
			` + formToInputs(form) + `
		</form>
		<script>document.getElementById("ecpay").submit();</script>
	`

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}

func generateCheckMacValue(params url.Values) string {
	// 1. 排序
	var keys []string
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 2. 組合為 query string
	var raw strings.Builder
	raw.WriteString("HashKey=" + hashKey)
	for _, key := range keys {
		raw.WriteString("&" + key + "=" + params.Get(key))
	}
	raw.WriteString("&HashIV=" + hashIV)

	// 3. URL encode 再轉大寫
	encoded := url.QueryEscape(raw.String())
	encoded = strings.ToLower(encoded)
	encoded = strings.ReplaceAll(encoded, "%2d", "-")
	encoded = strings.ReplaceAll(encoded, "%5f", "_")
	encoded = strings.ReplaceAll(encoded, "%2e", ".")
	encoded = strings.ReplaceAll(encoded, "%21", "!")
	encoded = strings.ReplaceAll(encoded, "%2a", "*")
	encoded = strings.ReplaceAll(encoded, "%28", "(")
	encoded = strings.ReplaceAll(encoded, "%29", ")")

	// 4. MD5 加密
	hash := md5.Sum([]byte(encoded))
	return strings.ToUpper(hex.EncodeToString(hash[:]))
}

func formToInputs(values url.Values) string {
	var b strings.Builder
	for key, vals := range values {
		for _, val := range vals {
			b.WriteString(`<input type="hidden" name="` + key + `" value="` + val + `">`)
		}
	}
	return b.String()
}
