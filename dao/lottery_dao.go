package dao

import (
	"database/sql"
	"fmt"
	"roybot/config"
	"roybot/model"
	"roybot/service/admin"
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
