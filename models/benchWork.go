package models

type BenchWork struct {
	Id          int64  `xorm:"BIGINT(20)"`
	Benchid     string `xorm:"not null pk VARCHAR(16)"`
	Benchcateid string `xorm:"VARCHAR(16)"`
	Upcode      string `xorm:"VARCHAR(255)"`
	Fullcode    string `xorm:"VARCHAR(128)"`
	Benchname   string `xorm:"VARCHAR(128)"`
	Inputstr    string `xorm:"VARCHAR(128)"`
	Createdate  string `xorm:"VARCHAR(32)"`
	Createempid string `xorm:"VARCHAR(32)"`
}
