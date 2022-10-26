package model

import (
	"github.com/shopspring/decimal"
)

type AccountHistory struct {
	Id      uint           `gorm:"primary_key" json:"-"`
   Block   int            `gorm:"column:block; default:null" json:"block"`
   Time    int            `gorm:"column:time; default:null" json:"time"`
	From    string         `sql:"default: null;size:100" json:"from"`
	To      string         `sql:"default: null;size:100" json:"to"`
	Value decimal.Decimal  `json:"value" sql:"type:decimal(30,0);"`
   Result       string    `gorm:"type:varchar(50)" json:"result"`
}

// type AccountData struct {
// 	Nonce    int `json:"nonce"`
// 	RefCount int `json:"ref_count"`
// 	Data     struct {
// 		Free       decimal.Decimal `json:"free"`
// 		Reserved   decimal.Decimal `json:"reserved"`
// 		MiscFrozen decimal.Decimal `json:"miscFrozen"`
// 		FeeFrozen  decimal.Decimal `json:"feeFrozen"`
// 	} `json:"data"`
// }
