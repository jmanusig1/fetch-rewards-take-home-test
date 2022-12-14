package models

import (
	"time"
)

type Transaction struct {
	Payer			string		`json:"payer" xml:"payer" form:"payer"`
	Points 			int			`json:"points" xml:"points" form:"points"`
	TimeStamp 		time.Time 	`json:"timestamp" xml:"timestamp" form:"timestamp"`
}