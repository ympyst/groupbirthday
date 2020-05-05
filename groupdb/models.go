package groupdb

import "time"

type Group struct {
	ID      		int			`gorm:"AUTO_INCREMENT"`
	Name    		string
	TelegramChatId	int64		`gorm:"unique"`
	Members 		[]Member 	`gorm:"many2many:group_members;unique"`
	Description 	string
}

type Member struct {
	ID               int			`gorm:"AUTO_INCREMENT"`
	FirstName        string
	LastName         string
	Birthday         *time.Time		`gorm:"default:NULL"`
	TelegramUsername string			`gorm:"unique"`
	TelegramUserId   int32			`gorm:"unique"`
	Groups           []Group		`gorm:"many2many:group_members;"`
}

type Fundraiser struct {
	Group       Group
	GroupId     int
	Member      Member
	MemberId    int
	PhoneNumber string
	BanksList   string
}

type BirthdayCelebration struct {
	Group          Group
	GroupId        int
	Member         Member
	MemberId       int
	Fundraiser     Fundraiser
	FundraiserId   int
	TelegramChatId int
	Year           int
}
