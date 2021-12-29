package models

type StatsOutput struct {
	Userid    int     `valid:"" json:"userid"`
	Dbid      int     `valid:"" json:"dbid"`
	Query     string  `valid:"" json:"query"`
	Calls     int     `valid:"" json:"calls"`
	TotalTime float32 `valid:"" json:"total_time"`
	MeanTime  float32 `valid:"" json:"mean_time"`
}
