package models

import (
	"database/sql"
	"errors"
	"fmt"
	"module/logger"
	"time"
)

type UsersModel struct {
	Id            int64     `json:"id"`
	UserName      string    `json:"username"`        //当前平台注册用户名(手机号)
	Password      string    `json:"password"`        //登录密码 SHA256 len = 64
	Name          string    `json:"name"`            //姓名
	LoginPlatform string    `json:"login_platform"`  //第三方登录平台
	UnionId       string    `json:"union_id"`        //第三方登录union_id
	PayPassword   string    `json:"pay_password"`    //支付密码 sha256
	VipId         int64     `json:"vip_id"`          //VIP表信息编号
	VipExpireTime time.Time `json:"vip_expire_time"` //VIP到期时间
	Credit        float64   `json:"credit"`          //积分
	Commission    float64   `json:"commission"`      //佣金
	ReferrerId    int64     `json:"referrer_id"`     //推荐人
	Age           int64     `json:"age"`             //年龄
	Sex           string    `json:"sex"`             //性别
	Address       string    `json:"address"`         //住址
	IdCard        string    `json:"id_card"`         //身份证号
	Img           string    `json:"img"`             //图像
	LastLoginTime string    `json:"last_login_time"` //最后登录
	CreateTime    string    `json:"create_time"`
	Status        bool      `json:"status"` //状态
	Email         string    `json:"email"`
}

func (m UsersModel) RetrieveLevelForTable(page int, limit int, referrerId int) (menus []interface{}, total int, err error) {
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()
	queryString := fmt.Sprintf("SELECT COUNT(*) FROM users WHERE referrer_id = %d", referrerId)
	err = db["default"].QueryRow(queryString).Scan(&total)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			return
		}
	}
	offset := limit * (page - 1)
	queryString = fmt.Sprintf("SELECT * FROM users WHERE referrer_id = %d ORDER BY id DESC LIMIT %d OFFSET %d", referrerId, limit, offset)
	rows, err := db["default"].Query(queryString)
	defer rows.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			return
		}
	}
	var temp1, temp2 time.Time
	for rows.Next() {
		err = rows.Scan(&m.Id, &m.UserName, &m.Password, &m.Name, &m.LoginPlatform, &m.UnionId,
			&m.PayPassword, &m.VipId, &m.VipExpireTime, &m.Credit, &m.Commission, &m.ReferrerId,
			&m.Age, &m.Sex, &m.Address, &m.IdCard, &m.Img, &temp1, &temp2, &m.Status, &m.Email)
		m.LastLoginTime = temp1.Format("2006-01-02 15:04:05")
		m.CreateTime = temp2.Format("2006-01-02 15:04:05")
		menus = append(menus, m)
	}
	return
}

func (m UsersModel) RetrieveForTable(page int, limit int) (menus []interface{}, total int, err error) {
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()
	err = db["default"].QueryRow("SELECT COUNT(*) FROM users").Scan(&total)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			return
		}
	}
	offset := limit * (page - 1)
	queryString := fmt.Sprintf("SELECT * FROM users ORDER BY id DESC LIMIT %d OFFSET %d", limit, offset)
	rows, err := db["default"].Query(queryString)
	defer rows.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			return
		}
	}
	var temp1, temp2 time.Time
	for rows.Next() {
		err = rows.Scan(&m.Id, &m.UserName, &m.Password, &m.Name, &m.LoginPlatform, &m.UnionId,
			&m.PayPassword, &m.VipId, &m.VipExpireTime, &m.Credit, &m.Commission, &m.ReferrerId,
			&m.Age, &m.Sex, &m.Address, &m.IdCard, &m.Img, &temp1, &temp2, &m.Status, &m.Email)
		m.LastLoginTime = temp1.Format("2006-01-02 15:04:05")
		m.CreateTime = temp2.Format("2006-01-02 15:04:05")
		menus = append(menus, m)
	}
	return
}

func (m UsersModel) Retrieve() (menus []interface{}, err error) {
	return
}

func (m UsersModel) RetrieveIndividual() (user interface{}, err error) {
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()

	/*
		err = db["default"].QueryRow("SELECT * FROM users WHERE username = $1",m.UserName).Scan(
			&m.Id, &m.UserName, &m.Password, &m.Name, &m.LoginPlatform, &m.UnionId,
			&m.PayPassword, &m.VipId, &m.VipExpireTime, &m.Credit, &m.Commission, &m.ReferrerId,
			&m.Age, &m.Sex, &m.Address, &m.IdCard, &m.Img, &m.LastLoginTime, &m.CreateTime, &m.Status, &m.Email,
			)
	*/
	if err != nil {
		return
	}
	user = m
	return
}

func (m UsersModel) RetrieveWhereIn(list []int64) (users []interface{}, err error) {
	return
}

func (m UsersModel) Create() (result string, err error) {
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()
	stmt, err := db["default"].Prepare(`INSERT INTO public.users(username, password, name, login_platform, union_id, pay_password, vip_id, vip_expire_time, credit, commission, referrer_id, age, sex, address, id_card, img, status, email)VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)`)
	defer stmt.Close()
	if err != nil {
		return
	}
	res, err := stmt.Exec(m.UserName, m.Password, m.Name,
		m.LoginPlatform, m.UnionId, m.PayPassword, m.VipId,
		m.VipExpireTime, m.Credit, m.Commission, m.ReferrerId,
		m.Age, m.Sex, m.Address, m.IdCard, m.Img, m.Status, m.Email)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect < 1 {
		err = errors.New("Insert error.")
		return
	}
	result = fmt.Sprintf("The affected numbers is %d", affect)
	return

	var (
		table  = "users"
		column = []string{`username`, `password`, `name`,
			`login_platform`, `union_id`, `pay_password`, `vip_id`,
			`vip_expire_time`, `credit`, `commission`, `referrer_id`,
			`age`, `sex`, `address`, `id_card`, `img`, `status`, `email`,
		}
		value = []interface{}{m.UserName, m.Password, m.Name,
			m.LoginPlatform, m.UnionId, m.PayPassword, m.VipId,
			m.VipExpireTime, m.Credit, m.Commission, m.ReferrerId,
			m.Age, m.Sex, m.Address, m.IdCard, m.Img, m.Status, m.Email,
		}
		id = m.Id
	)
	result, err = updateRows(table, column, value, id)
	return
}

func (m UsersModel) Update() (result string, err error) {
	var (
		table  = "users"
		column = []string{`username`, `password`, `name`,
			`login_platform`, `union_id`, `pay_password`, `vip_id`,
			`vip_expire_time`, `credit`, `commission`, `referrer_id`,
			`age`, `sex`, `address`, `id_card`, `img`, `status`, `email`,
		}
		value = []interface{}{m.UserName, m.Password, m.Name,
			m.LoginPlatform, m.UnionId, m.PayPassword, m.VipId,
			m.VipExpireTime, m.Credit, m.Commission, m.ReferrerId,
			m.Age, m.Sex, m.Address, m.IdCard, m.Img, m.Status, m.Email,
		}
		id = m.Id
	)
	result, err = updateRows(table, column, value, id)
	return
}

func (m UsersModel) Delete() (result string, err error) {
	var (
		table      = "users"
		column     = []string{`id`}
		value      = []interface{}{m.Id}
		levelName  = `referrer_id`
		levelValue = m.Id
	)
	result, err = deleteRows(table, column, value, levelName, levelValue)
	return
}
