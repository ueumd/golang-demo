package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/myapiserver?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, _ := fetchRows(db, "SELECT * FROM test")
	for _, v := range *rows {
		fmt.Println(v["a"], v["b"])
	}

	//fmt.Println(insert(db, "INSERT INTO test( b ) VALUES( ? )", 1))
	//row, _ := fetchRow(db, "SELECT * FROM test where a = ?", 1)
	//fmt.Println(*row)
}
func fetchRows(db *sql.DB, sqlstr string, args ...interface{}) (*[]map[string]string, error) {
	stmtOut, err := db.Prepare(sqlstr)

	if err != nil {
		panic(err.Error())
	}

	defer stmtOut.Close()

	rows, err := stmtOut.Query(args...)
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	ret := make([]map[string]string, 0)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string
		vmap := make(map[string]string, len(scanArgs))
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			vmap[columns[i]] = value
		}
		ret = append(ret, vmap)
	}
	return &ret, nil

}