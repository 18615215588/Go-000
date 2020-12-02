package dao

import (
	"database/sql"
	"errors"

	errs "github.com/pkg/errors"
)

//UserDetail 用户详细信息数据
type UserDetail struct {
	UID      string
	UserName string
	Age      uint8
	Address  string
	Email    string
	PhoneNum string
}

//GetUserFromDB 模拟从数据库获取用户信息，返回ErrNoRows错误，仅仅模拟:)
func GetUserFromDB(uid string) (*UserDetail, error) {
	//return nil, fmt.Errorf("other error:%s", "no matter")
	return nil, sql.ErrNoRows
	//panic("opoos panic!!")
}

//GetUserDetails 获取用户详细信息
func GetUserDetails(uid string) (*UserDetail, error) {
	//模拟从数据库获取数据,会返回ErrNoRows错误
	u, e := GetUserFromDB(uid)
	if e != nil && errors.Is(e, sql.ErrNoRows) {
		//没有记录，返回降级数据，吞掉错误信息
		return &UserDetail{
			"na",
			"na",
			0,
			"na",
			"na",
			"na",
		}, nil
	} else if e != nil { //其他错误信息封装一次返回上层
		return nil, errs.Wrap(e, "query db error")
	}
	return u, nil
}
