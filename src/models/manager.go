package models

import (
	"database/sql"
	"errors"
	"fmt"
	"module/logger"
	"time"
)

type ManagerModel struct {
	Id            int64  `json:"id"`
	User          string `json:"user"`
	Password      string `json:"password"`
	Name          string `json:"name"`
	LastLoginTime string `json:"last_login_time"`
	CreateTime    string `json:"create_time"`
}

func (m ManagerModel) RetrieveLevelForTable(page int, limit int, parentId int) (managers []interface{}, total int, err error) {
	return
}

func (m ManagerModel) RetrieveForTable(page int, limit int) (managers []interface{}, total int, err error) {
	// It dou't need to close,because it hava been closed inside Scan function
	err = db["default"].QueryRow("SELECT COUNT(*) FROM manager").Scan(&total)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			return
		}
	}
	offset := limit * (page - 1)
	queryString := fmt.Sprintf("SELECT * FROM manager ORDER BY id DESC LIMIT %d OFFSET %d", limit, offset)
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
		err = rows.Scan(&m.Id, &m.User, &m.Password, &m.Name, &temp1, &temp2)
		m.LastLoginTime = temp1.Format("2006-01-02 15:04:05")
		m.CreateTime = temp2.Format("2006-01-02 15:04:05")
		managers = append(managers, m)
	}
	return
}

func (m ManagerModel) Retrieve() (managers []interface{}, err error) {
	queryString := fmt.Sprintf("SELECT * FROM manager WHERE user = %s", m.User)
	rows, err := db["default"].Query(queryString)
	defer rows.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			return
		}
	}

	for rows.Next() {
		err = rows.Scan(&m.Id, &m.User, &m.Password, &m.Name, &m.LastLoginTime, &m.CreateTime)
		if err != nil {
			return
		}
		managers = append(managers, m)
	}
	return
}

func (m ManagerModel) RetrieveIndividual() (user interface{}, err error) {
	return
}

func (m ManagerModel) RetrieveWhereIn(list []int64) (apis []interface{}, err error) {
	return
}

func (m ManagerModel) Create() (result string, err error) {
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()
	stmt, err := db["default"].Prepare(`INSERT INTO manager ("user","password","name") VALUES($1,$2,$3)`)
	defer stmt.Close()
	if err != nil {
		return
	}
	res, err := stmt.Exec(m.User, m.Password, m.Name)
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
}

func (m ManagerModel) Update() (result string, err error) {
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()
	stmt, err := db["default"].Prepare("UPDATE manager SET password = $1, name = $2 WHERE id = $3")
	defer stmt.Close()
	if err != nil {
		return
	}
	res, err := stmt.Exec(m.Password, m.Name, m.Id)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect < 1 {
		err = errors.New("Update error.")
		return
	}
	result = fmt.Sprintf("The affected numbers is %d", affect)
	return
}

func (m ManagerModel) Delete() (result string, err error) {
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()
	stmt, err := db["default"].Prepare("DELETE FROM manager WHERE id=$1")
	defer stmt.Close()
	if err != nil {
		return
	}
	res, err := stmt.Exec(m.Id)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}

	if affect < 1 {
		err = errors.New("The affected numbers is less than 1.")
		return
	}
	result = fmt.Sprintf("The affected numbers is %d", affect)
	return
}
