package models

type Balance struct {
	Payer		string		`json:"payer" xml:"payer" form:"payer"`
	Points 		int			`json:"points" xml:"points" form:"points"`
}