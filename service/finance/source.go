package finance

import (
	"fmt"
	"net/http"
	"roybot/service/admin"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const future = "https://tw.screener.finance.yahoo.net/future/aa02"

//GetFuture is func
func GetFuture() string {
	res, err := http.Get(future)

	if err != nil {
		admin.CallAdmin("GetFuture", err)
		return ""
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		admin.CallAdmin("GetFuture:NewDocumentFromReader", err)
		return ""
	}

	stock := doc.Find("#ystkfutb .tbd tbody > tr").Slice(0, 1)
	futrue1 := doc.Find("#ystkfutb .tbd tbody > tr").Slice(1, 2)
	futrue2 := doc.Find("#ystkfutb .tbd tbody > tr").Slice(7, 8)
	smark, f1mark, f2mark := "▲", "▲", "▲"
	if stock.Find("td").Slice(1, 2).Text() == "跌" {
		smark = "▼"
	}
	if strings.TrimSpace(futrue1.Find("td").Slice(1, 2).Text()) == "跌" {
		f1mark = "▼"
	}
	if strings.TrimSpace(futrue2.Find("td").Slice(1, 2).Text()) == "跌" {
		f2mark = "▼"
	}

	result := fmt.Sprintf("上市大盤    %s\n%s %s%s %s\n"+
		"台指期近一    %s\n%s %s%s %s\n"+
		"台指期近二    %s\n%s %s%s %s\n",
		stock.Find("td").Slice(4, 5).Text(), stock.Find("td").Slice(0, 1).Text(),
		smark, stock.Find("td").Slice(2, 3).Text(), stock.Find("td").Slice(3, 4).Text(),
		futrue1.Find("td").Slice(4, 5).Text(), futrue1.Find("td").Slice(0, 1).Text(),
		f1mark, strings.Replace(futrue1.Find("td").Slice(2, 3).Text(), "-", "", 1), futrue1.Find("td").Slice(3, 4).Text(),
		futrue2.Find("td").Slice(4, 5).Text(), futrue2.Find("td").Slice(0, 1).Text(),
		f2mark, strings.Replace(futrue2.Find("td").Slice(2, 3).Text(), "-", "", 1), futrue2.Find("td").Slice(3, 4).Text())

	return result
}
