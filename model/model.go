package model

import "time"

//DailyData is model
type DailyData struct {
	ID       string
	OpenDate time.Time
	Issue    string
	N1       int
	N2       int
	N3       int
	N4       int
	N5       int
	N6       int
	N7       int
	N8       int
	N9       int
	SP       int
	NSum     int
	NAvg     int
	NSeq     string
	FLDiff   int
	SDRate   string
}

//CompareDuring is model
type CompareDuring struct {
	Start string
	End   string
}
