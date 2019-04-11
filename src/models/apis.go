package models

import (
	"database/sql"
	"errors"
	"fmt"
	"module/logger"
	"strconv"
	"time"
)

type ApisModel struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	Status     bool   `json:"status"`
	CreateTime string `json:"create_time"`
}

func (a ApisModel) RetrieveLevelForTable(page int, limit int, parentId int) (apis []interface{}, total int, err error) {
	return
}

func (a ApisModel) RetrieveForTable(page int, limit int) (apis []interface{}, total int, err error) {
	// It dou't need to close,because it hava been closed inside Scan function
	err = db["default"].QueryRow("SELECT COUNT(*) FROM api_list").Scan(&total)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			return
		}
	}
	offset := limit * (page - 1)
	queryString := fmt.Sprintf("SELECT * FROM api_list ORDER BY id DESC LIMIT %d OFFSET %d", limit, offset)
	rows, err := db["default"].Query(queryString)
	defer rows.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			return
		}
	}
	var temp time.Time
	for rows.Next() {
		err = rows.Scan(&a.Id, &a.Name, &a.Url, &a.Status, &temp)
		a.CreateTime = temp.Format("2006-01-02 15:04:05")
		apis = append(apis, a)
	}
	return
}

func (a ApisModel) Retrieve() (apis []interface{}, err error) {
	return
}

func (m ApisModel) RetrieveIndividual() (user interface{}, err error) {
	return
}

// TODO: change interface args to slice
func (a ApisModel) RetrieveWhereIn(list []int64) (apis []interface{}, err error) {
	var condition string = ""
	for i := 0; i < len(list); i++ {
		// int64 to 10 bit string
		item := strconv.FormatInt(list[i], 10)
		if last := len(list) - 1; i == last {
			condition += item
		}
		condition += item + `,`
	}
	/*
	 columns := []string{"id", "name"}
	 id := 1
	 sql := fmt.Sprintf("select %s from user where id=%d", strings.Join(columns, ", "), id)
	*/
	queryString := fmt.Sprintf("SELECT * FROM api_list WHERE id IN (%s)", condition)
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
		err = rows.Scan(&a.Id, &a.Name, &a.Url, &a.Status, &a.CreateTime)
		if err != nil {
			return
		}
		apis = append(apis, a)
	}
	return
}

func (a ApisModel) Create() (result string, err error) {
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()
	stmt, err := db["default"].Prepare(`INSERT INTO api_list ("name", "url", "status") VALUES($1,$2,$3)`)
	defer stmt.Close()
	if err != nil {
		return
	}
	res, err := stmt.Exec(a.Name, a.Url, a.Status)
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

func (a ApisModel) Update() (result string, err error) {
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()
	stmt, err := db["default"].Prepare("UPDATE api_list SET name = $1, url = $2, status = $3 WHERE id = $4")
	defer stmt.Close()
	if err != nil {
		return
	}
	res, err := stmt.Exec(a.Name, a.Url, a.Status, a.Id)
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

func (a ApisModel) Delete() (result string, err error) {
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()
	stmt, err := db["default"].Prepare("DELETE FROM api_list WHERE id = $1")
	defer stmt.Close()
	if err != nil {
		return
	}
	res, err := stmt.Exec(a.Id)
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
