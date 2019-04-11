package models

import (
	"database/sql"
	"errors"
	"fmt"
	"module/logger"
	"time"
)

type MenuModel struct {
	Id         int64  `json:"id"`
	ParentId   int64  `json:"parent_id"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	Ico        string `json:"ico"`
	Sort       int8   `json:"sort"`
	Status     bool   `json:"status"`
	CreateTime string `json:"create_time"`
}

//type MenuModels []MenuModel

func (m MenuModel) RetrieveLevelForTable(page int, limit int, parentId int) (menus []interface{}, total int, err error) {
	queryString := fmt.Sprintf("SELECT COUNT(*) FROM menu WHERE parent_id = %d", parentId)
	// It dou't need to close,because it hava been closed inside Scan function
	err = db["default"].QueryRow(queryString).Scan(&total)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			return
		}
	}
	offset := limit * (page - 1)
	queryString = fmt.Sprintf("SELECT * FROM menu WHERE parent_id = %d ORDER BY id DESC LIMIT %d OFFSET %d", parentId, limit, offset)
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
		err = rows.Scan(&m.Id, &m.ParentId, &m.Name, &m.Url, &m.Ico, &m.Sort, &m.Status, &temp)
		m.CreateTime = temp.Format("2006-01-02 15:04:05")
		menus = append(menus, m)
	}
	return
}

func (m MenuModel) RetrieveForTable(page int, limit int) (menus []interface{}, total int, err error) {
	// It dou't need to close,because it hava been closed inside Scan function
	err = db["default"].QueryRow("SELECT COUNT(*) FROM menu").Scan(&total)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			return
		}
	}
	offset := limit * (page - 1)
	queryString := fmt.Sprintf("SELECT * FROM menu ORDER BY id DESC LIMIT %d OFFSET %d", limit, offset)
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
		err = rows.Scan(&m.Id, &m.ParentId, &m.Name, &m.Url, &m.Ico, &m.Sort, &m.Status, &temp)
		m.CreateTime = temp.Format("2006-01-02 15:04:05")
		menus = append(menus, m)
	}
	return
}

func (m MenuModel) Retrieve() (menus []interface{}, err error) {
	rows, err := db["default"].Query("SELECT * FROM menu WHERE status = true")
	defer rows.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			return
		}
	}
	for rows.Next() {
		err = rows.Scan(&m.Id, &m.ParentId, &m.Name, &m.Url, &m.Ico, &m.Sort, &m.Status, &m.CreateTime)
		if err != nil {
			return
		}
		menus = append(menus, m)
	}
	return
}

func (m MenuModel) Create() (result string, err error) {
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()
	stmt, err := db["default"].Prepare(`INSERT INTO menu ("parent_id", "name", "url", "ico", "sort", "status") VALUES($1, $2, $3, $4, $5, $6)`)
	defer stmt.Close()
	if err != nil {
		return
	}
	res, err := stmt.Exec(m.ParentId, m.Name, m.Url, m.Ico, m.Sort, m.Status)
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

func (m MenuModel) Update() (result string, err error) {
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()
	stmt, err := db["default"].Prepare("UPDATE menu SET parent_id = $1, name = $2, url = $3, ico = $4, sort = $5, status = $6 WHERE id = $7")
	defer stmt.Close()
	if err != nil {
		return
	}
	res, err := stmt.Exec(m.ParentId, m.Name, m.Url, m.Ico, m.Sort, m.Status, m.Id)
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

func (m MenuModel) Delete() (result string, err error) {
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()
	var total int
	queryString := fmt.Sprintf("SELECT COUNT(*) FROM menu WHERE parent_id = %d", m.Id)
	err = db["default"].QueryRow(queryString).Scan(&total)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			return
		}
	}
	if total > 0 {
		err = errors.New("It is parent_id of many menus.")
		return
	}
	stmt, err := db["default"].Prepare("DELETE FROM menu WHERE id=$1")
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

/*
// 批量插入
func (ms MenuModels) BulkWay() {
	tx, err := db["default"].Begin()
	if err != nil {
		fmt.Println("Beginx error:", err)
		panic(err)
	}
	stmt, err := tx.Prepare("insert into test_b (name, alia, city_id) values ($1, $2, $3)")
	if err != nil {
		fmt.Println("Prepare error:", err)
		panic(err)
	}
	for _, value := range values {
		_, err = stmt.Exec(value...)
		if err != nil {
			fmt.Println("Exec error:", err)
			panic(err)
		}
	}

	err = stmt.Close()
	if err != nil {
		fmt.Println("stmt close error:", err)
		panic(err)
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("commit error:", err)
		panic(err)
	}
}
*/
