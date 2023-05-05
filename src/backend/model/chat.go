package model

type Chat struct {
	Chat_ID			int64 	`gorm:"primaryKey" json:"id"`
	Question		string	`gorm:"type:text" json:"question"`
	Answer			string	`gorm:"type:text" json:"answer"`
}