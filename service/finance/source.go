package finance

import (
	"github.com/djimenez/iconv-go"
	"fmt"
	"net/http"
	"roybot/service/admin"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	future        = "https://tw.screener.finance.yahoo.net/future/aa02"
	foreignFuture = "https://tw.stock.yahoo.com/"
)

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

func GetForeignFuture() string {
	res, err := http.Get(foreignFuture)

	if err != nil {
		admin.CallAdmin("GetForeignFuture", err)
		return ""
	}

	defer res.Body.Close()
	utfBody, err := iconv.NewReader(res.Body, "big5", "utf-8")

	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		admin.CallAdmin("GetForeignFuture", err)
		return ""
	}

	result := ""

	doc.Find("#ystkchatwtb div tbody tr").Each(func(i int, s *goquery.Selection) {
		// th := s.Find("th")
		t, _ := s.Html()
		fmt.Printf("%d\n%+v\n", i, t)
		if i < 4 {
			smark := "▲"
			if s.Find("td").Slice(1, 2).Text() == "跌" {
				smark = "▼"
			}
			result += fmt.Sprintf("%s %s %s %s\n", s.Find("a").Slice(0, 1).Text(), s.Find("td").Slice(0, 1).Text(), 
			smark + s.Find("td").Slice(2, 3).Text(), s.Find("td").Slice(3, 4).Text())
		} else if (i != 8 && i != 13 && i != 18) {
			smark := "▲"
			if s.Find("td").Slice(1, 2).Text() == "跌" {
				smark = "▼"
			}
			result += fmt.Sprintf("%s %s %s\n", s.Find("a").Slice(0, 1).Text(), s.Find("td").Slice(0, 1).Text(), smark + s.Find("td").Slice(2, 3).Text())
		}
		
	})
	return result
}
