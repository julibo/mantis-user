package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

type (
	Members struct {
		ID			int64
		Token		string		`json:"token", xorm: "varchar(255) notnull 'token'"`
		Account		string		`json:"account", xorm: "varchar(40) notnull 'account'"`
		Password 	string		`json:"password", xorm: "varchar(40) notnull 'password'"`
		Email 		string		`json:"email", xorm: "varchar(100) 'email'"`
		Mobile 		string		`json:"mobile", xorm: "varchar(20) 'mobile'"`
		CreateTime	time.Timer 	`json:"create_time", xorm: "timestamp notnull 'update_time'"`
		UpdateTime	time.Time	`json:"update_time", xorm: "timestamp notnull 'update_time'"`
	}
	MembersModel struct {
		mysql *xorm.Engine
	}
)

func NewMembersModel(mysql *xorm.Engine) *MembersModel {
	return &MembersModel{
		mysql: mysql,
	}
}

func (m *MembersModel) FindByToken(token string) (*Members, error) {
	members := new(Members)
	if _, err := m.mysql.Where("token=?", token).Get(members); err != nil {
		return nil, err
	}
	return members, nil
}

func (m *MembersModel) FindByID(id int64) (*Members, error) {
	fmt.Println("risini ", id)
	members := new(Members)
	members.ID = 1
	members.Password = "123456"
	members.Token = "1020300400606060"
	//if _, err := m.mysql.Get(members); err != nil {
	//	return nil, err
	//}
	return members, nil
}
