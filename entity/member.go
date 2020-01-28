package entity

type Member struct {
	 Id  int      `xorm:"pk autoincr" `
     Name string  `xorm:"VARCHAR(20) notnull"`
	 Age  int   `xorm:"notnull INT(11)"`
}
