package models

import (
	"config"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"module/logger"
	"strconv"
	"strings"
)

var db map[string]*sql.DB

func init() {
	defaultConnection()
}

func defaultConnection() {
	dbconf := config.GetEnv()
	pgurl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbconf.DB_USERNAME, dbconf.DB_PASSWORD, dbconf.DB_IP, dbconf.DB_PORT, dbconf.DB_NAME)
	dbobj, err := sql.Open("postgres", pgurl)
	if err != nil {
		//dbobj.Close()
		panic(err.Error())
	}

	dbobj.SetMaxOpenConns(2000)
	dbobj.SetMaxIdleConns(1000) // Setting max idle Connections.

	err = dbobj.Ping()
	if err != nil {
		dbobj.Close()
		panic(err.Error())
	}

	db = map[string]*sql.DB{
		"default": dbobj,
	}
	fmt.Println("Successfully connected! :", dbconf.DB_NAME)
}

type sqlQuery struct {
	TableString   string
	FieldsString  string
	WhereSring    string
	OrderByString string
	LimitString   string
	Data          []interface{}
}

func (s *sqlQuery) Table(tableName string) *sqlQuery {
	s.TableString = tableName
	return s
}

func (s *sqlQuery) Field(columns []string) *sqlQuery {
	s.FieldsString = strings.Join(columns, ", ")
	return s
}

/*
**  where: " = ", "like"
**  "AND"
 */
func (s *sqlQuery) Where(columns []string, operator string, value []interface{}) *sqlQuery {
	var err error
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()
	if len(columns) != len(value) || len(columns) == 0 || len(value) == 0 {
		err = errors.New(" Args error.")
	}
	var condition = "WHERE "
	var val string
	for i := 0; i < len(columns); i++ {
		switch value[i].(type) {
		case int:
			val = strconv.Itoa(value[i].(int))
		case int64:
			val = strconv.FormatInt(value[i].(int64), 10)
		case float64:
			val = strconv.FormatFloat(value[i].(float64), 'e', 2, 64)
		case bool:
			val = strconv.FormatBool(value[i].(bool))
		case string:
			val = `"` + value[i].(string) + `"`
		default:
			err = errors.New("Unknow type of Value!")
		}
		if i == (len(columns) - 1) {
			condition += columns[i] + operator + val
			break
		}
		condition += columns[i] + operator + val + "AND"
	}
	s.WhereSring = condition
	return s
}

func (s *sqlQuery) WhereIn(column string, value []interface{}) *sqlQuery {
	var err error
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()
	var condition = `WHERE ` + column + ` IN (`
	var val string
	for i := 0; i < len(value); i++ {
		switch value[i].(type) {
		case int:
			val = strconv.Itoa(value[i].(int))
		case int64:
			val = strconv.FormatInt(value[i].(int64), 10)
		case float64:
			val = strconv.FormatFloat(value[i].(float64), 'e', 2, 64)
		case bool:
			val = strconv.FormatBool(value[i].(bool))
		case string:
			val = `"` + value[i].(string) + `"`
		default:
			err = errors.New("Unknow type of Value!")
		}
		if i == (len(value) - 1) {
			condition += val + `)`
			break
		}
		condition += val + `, `
	}
	s.WhereSring = condition
	return s
}

func (s *sqlQuery) WhereBetween(column string, value1 interface{}, value2 interface{}) *sqlQuery {
	s.WhereSring = fmt.Sprintf("WHERE %s BETWEEN %v AND %v", column, value1, value2)
	return s
}

func (s *sqlQuery) Order(column string, desc string) *sqlQuery {
	s.LimitString = fmt.Sprintf("ORDER BY %s %s", column, desc)
	return s
}

func (s *sqlQuery) Limit(page int, limit int) *sqlQuery {
	offset := limit * (page - 1)
	s.LimitString = fmt.Sprintf("LIMIT %d OFFSET %d", limit, offset)
	return s
}

func (s *sqlQuery) Select() (data []interface{}, err error) {
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()
	queryString := `SELECT ` + s.TableString + s.FieldsString + s.WhereSring + s.OrderByString + s.LimitString
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
		err = rows.Scan(s.Data...)
		data = append(data, s.Data)
	}
	return
}

func (s *sqlQuery) Create(data []interface{}) (result string, err error) {
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()
	var condition = ` VALUES(`
	for i := 0; i < len(s.Data); i++ {
		if i == (len(s.Data) - 1) {
			condition += `$` + strconv.Itoa(i) + `)`
			break
		}
		condition += `$` + strconv.Itoa(i) + ` ,`
	}
	queryString := `INSERT INTO` + s.TableString + `(` + s.FieldsString + `)` + condition
	stmt, err := db["default"].Prepare(queryString)
	defer stmt.Close()
	if err != nil {
		return
	}
	res, err := stmt.Exec(s.Data[:]...)
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

func (s *sqlQuery) Update() (result string, err error) {
	return
}

func (s *sqlQuery) Delete() (result string, err error) {
	return
}

func getCount(table string, column []string, value []interface{}, operator string) (num int, err error) {
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()
	if table == "" {
		err = errors.New("Table name is not exist.")
		return
	}
	if len(column) != len(value) {
		err = errors.New("Columns number unequal to Args.")
		return
	}
	if len(column) == 0 {
		err = db["default"].QueryRow("SELECT COUNT(*) FROM $1", table).Scan(&num)
		return
	}
	var condition string
	var val string
	for i := 0; i < len(column); i++ {
		switch value[i].(type) {
		case int:
			val = strconv.Itoa(value[i].(int))
		case float64:
			val = strconv.FormatFloat(value[i].(float64), 'e', 2, 64)
		case bool:
			val = strconv.FormatBool(value[i].(bool))
		case string:
			val = `"` + value[i].(string) + `"`
		default:
			err = errors.New("Unknow type of levelValue!")
			return
		}
		if i == (len(column) - 1) {
			condition += column[i] + operator + val
			break
		}
		condition += column[i] + operator + val + "AND"
	}
	queryString := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s", table, condition)
	err = db["default"].QueryRow(queryString).Scan(&num)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			return
		}
	}
	return
}

func updateRows(table string, column []string, value []interface{}, id int64) (result string, err error) {
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()
	if table == "" {
		err = errors.New("Table name is not exist.")
		return
	}
	if len(column) != len(value) {
		err = errors.New("Columns number unequal to Args.")
		return
	}
	if len(column) == 0 {
		err = errors.New("Can not find columns for update.")
		return
	}
	var condition string
	for i := 0; i < len(column); i++ {
		if i == (len(column) - 1) {
			condition += column[i] + ` = ` + `$` + strconv.Itoa(i+1)
			break
		}
		condition += column[i] + ` = ` + `$` + strconv.Itoa(i+1) + `, `
	}
	queryString := fmt.Sprintf("UPDATE %s SET %s WHERE id = %d", table, condition, id)
	stmt, err := db["default"].Prepare(queryString)
	defer stmt.Close()
	if err != nil {
		return
	}
	res, err := stmt.Exec(value...)
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

func deleteRows(table string, column []string, value []interface{}, levelName string, levelValue interface{}) (result string, err error) {
	defer func() {
		if err != nil {
			logger.Error(err)
		}
	}()
	if table == "" {
		err = errors.New("Table name is not exist.")
		return
	}
	if len(column) != len(value) || len(column) == 0 || len(value) == 0 {
		err = errors.New(" Args error.")
		return
	}
	if levelName != "" && levelValue != nil {
		var total int
		var leVal string
		switch levelValue.(type) {
		case int:
			leVal = strconv.Itoa(levelValue.(int))
		case int64:
			leVal = strconv.FormatInt(levelValue.(int64), 10)
		case float64:
			leVal = strconv.FormatFloat(levelValue.(float64), 'e', 2, 64)
		case bool:
			leVal = strconv.FormatBool(levelValue.(bool))
		case string:
			leVal = `"` + levelValue.(string) + `"`
		default:
			err = errors.New("Unknow type of levelValue!")
			return
		}
		queryString := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = %s", table, levelName, leVal)
		err = db["default"].QueryRow(queryString).Scan(&total)
		if err != nil {
			if err == sql.ErrNoRows {
				err = nil
			} else {
				return
			}
		}
		if total > 0 {
			err = errors.New("It is parent of many id.")
			return
		}
	}
	var condition string
	for i := 0; i < len(column); i++ {
		if i == (len(column) - 1) {
			condition += column[i] + ` = ` + `$` + strconv.Itoa(i+1)
			break
		}
		condition += column[i] + ` = ` + `$` + strconv.Itoa(i+1) + `, `
	}
	queryString := fmt.Sprintf("DELETE FROM %s WHERE %s", table, condition)
	stmt, err := db["default"].Prepare(queryString)
	defer stmt.Close()
	if err != nil {
		return
	}
	res, err := stmt.Exec(value...)
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
