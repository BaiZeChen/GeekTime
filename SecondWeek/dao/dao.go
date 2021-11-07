package dao

import (
	"database/sql"
	pkgerr "github.com/pkg/errors"
)

func GetUserNameById(id int) (string, error) {

	DB, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return "", pkgerr.Wrap(err, "数据库建立链接失败！")
	}
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		return "", pkgerr.Wrap(err, "数据库链接失败！")
	}

	var name string
	err = DB.QueryRow("select name from users where id = ?", id).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			// 如果没有获取到对应的数据，直接返回就好，不应该当做一个错误
			return "", nil
		} else {
			return "", pkgerr.Wrap(err, "查询数据失败！")
		}
	}

	return name, nil

}
