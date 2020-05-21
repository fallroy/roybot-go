package dao

import (
	"database/sql"
	"fmt"
	"roybot/config"
	"roybot/model"
	"roybot/service/admin"
	"strconv"
	"time"
)

//GetSummary is a func
func GetSummary() string {
	rows, err := config.DB.Query("SELECT   " +
		"(SELECT COUNT(*) AS total FROM d649) AS total,      " +
		"ROUND(sum(n_avg) / (SELECT COUNT(*) AS total FROM d649),2) as average ,      " +
		"	ROUND((select sum(n_sum) / 10 from ( " +
		"	select n_sum, count(n_sum) as cc from d649 group by n_sum order by cc desc limit 10 " +
		"	) tmp),2) AS sum_avg_top10 ," +
		"ROUND((select count(*)  from d649 where n_seq != '') / (SELECT COUNT(*) AS total FROM d649) * 100,2)  AS continuous_per,   " +
		"ROUND(sum(n_same) / (SELECT COUNT(*) AS total FROM d649) * 100,2) AS same_pri_per " +
		"FROM d649 order by issue asc ")
	defer rows.Close()

	if err != nil {
		admin.CallAdmin("GetSummary", err)
		return ""
	}

	s, total, avg, sumtop10avg, consecutive, samerate := "", "", "", "", "", ""

	if rows.Next() {
		if err := rows.Scan(&total, &avg, &sumtop10avg, &consecutive, &samerate); err != nil {
			admin.CallAdmin("GetSummary", err)
			return ""
		}
		s = fmt.Sprintf("Total : %s\nAvg : %s\nTop10Avg : %s\nConsecutive : %s %%\nSame : %s %%", total, avg, sumtop10avg, consecutive, samerate)
	}
	return s
}

//GetLatestData is func
func GetLatestData() model.DailyData {
	// fmt.Printf("issue :" + dailyDate.Issue)
	rows, err := config.DB.Query("SELECT id, open_date, issue, n1, n2, n3, n4, n5, n6, sp FROM d649 ORDER BY open_date DESC LIMIT 1 ")
	defer rows.Close()
	if err != nil {
		admin.CallAdmin("GetLatestData", err)
	}
	m := model.DailyData{}
	if rows.Next() {
		if err := rows.Scan(&m.ID, &m.OpenDate, &m.Issue, &m.N1, &m.N2, &m.N3, &m.N4, &m.N5, &m.N6, &m.SP); err != nil {
			admin.CallAdmin("GetLatestData", err)
		}
		return m
	}
	return m
}

//GetLatestRecord is func
func GetLatestRecord() model.DailyData {
	rows, err := config.DB.Query("SELECT id, open_date, issue, n1, n2, n3, n4, n5, n6 FROM d649_record ORDER BY issue DESC LIMIT 1 ")
	defer rows.Close()
	if err != nil {
		admin.CallAdmin("GetLatestRecord", err)
	}
	m := model.DailyData{}
	if rows.Next() {
		if err := rows.Scan(&m.ID, &m.OpenDate, &m.Issue, &m.N1, &m.N2, &m.N3, &m.N4, &m.N5, &m.N6); err != nil {
			admin.CallAdmin("GetLatestRecord", err)
		}
		return m
	}
	return m
}

//GetRecordByIssue is func
func GetRecordByIssue(issue string) model.DailyData {
	rows, err := config.DB.Query("SELECT id, open_date, issue, n1, n2, n3, n4, n5, n6 FROM d649_record WHERE issue = ? LIMIT 1 ", issue)
	defer rows.Close()
	if err != nil {
		admin.CallAdmin("GetRecordByIssue", err)
	}
	m := model.DailyData{}
	if rows.Next() {
		if err := rows.Scan(&m.ID, &m.OpenDate, &m.Issue, &m.N1, &m.N2, &m.N3, &m.N4, &m.N5, &m.N6); err != nil {
			admin.CallAdmin("GetRecordByIssue", err)
		}
		return m
	}
	return m
}

/*	Insert Functions*/

//Add649New is func
func Add649New(issue string, n1, n2, n3, n4, n5, n6 int, source, open_date string) {
	openDate, _ := time.Parse("2006-01-02", open_date)
	record := model.DailyData{
		Issue:    issue,
		OpenDate: openDate,
		N1:       n1,
		N2:       n2,
		N3:       n3,
		N4:       n4,
		N5:       n5,
		N6:       n6,
		FLDiff:   n6 - n1,
		SDRate:   getSDRate(n1, n2, n3, n4, n5, n6),
	}
	// fmt.Printf("Insert %+v", record)
	InsertRecordDate(&record, source)
}

//InsertRecordDate is func
func InsertRecordDate(dailyDate *model.DailyData, source string) {
	dailyDate.NSum = dailyDate.N1 + dailyDate.N2 + dailyDate.N3 + dailyDate.N4 + dailyDate.N5 + dailyDate.N6
	dailyDate.NAvg = dailyDate.NSum / 6
	// fmt.Printf("Insert %+v \n", dailyDate)

	_, err := config.DB.Exec(
		"Insert INTO d649_record(id, issue, open_date, n1, n2, n3, n4, n5, n6, n_sum, n_avg, fl_diff, sd_rate, source, created_date, updated_date) VALUES(UUID(),?,?,?,?,?,?,?,?,?,?,?,?,?,SYSDATE(),SYSDATE())",
		dailyDate.Issue, dailyDate.OpenDate,
		dailyDate.N1, dailyDate.N2, dailyDate.N3, dailyDate.N4, dailyDate.N5, dailyDate.N6,
		dailyDate.NSum, dailyDate.NAvg,
		dailyDate.FLDiff, dailyDate.SDRate, source,
	)

	if err != nil {
		admin.CallAdmin("InsertRecordDate", err)
	}

}

/*	Private Functions	*/
func convertToModel(rows *sql.Rows) []model.DailyData {
	result := make([]model.DailyData, 0)
	for rows.Next() {
		m := model.DailyData{}
		if err := rows.Scan(&m.ID, &m.OpenDate, &m.Issue, &m.N1, &m.N2, &m.N3, &m.N4, &m.N5, &m.N6, &m.SP); err != nil {
			admin.CallAdmin("convertToModel", err)
		}
		result = append(result, m)
	}
	return result
}

func getSDRate(nums ...int) string {
	s, d := 0, 0
	for _, num := range nums {
		if num%2 == 0 {
			d++
		} else {
			s++
		}
	}
	return strconv.Itoa(s) + ":" + strconv.Itoa(d)
}
