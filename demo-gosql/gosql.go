package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv" //这个是为了把int转换为string
)

func main()  {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/myapiserver?charset=utf8")

	if err != nil {
		panic(err.Error()) 		  //抛出异常
		fmt.Println(err.Error())  //仅仅显示异常
	}

	// 只有在前面用了 panic 这时defer才能起作用，如果链接数据的时候出问题，他会往err写数据
	defer db.Close()

	rows, err := db.Query("select * from users")

	if err != nil {
		panic(err.Error())
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	var (
		id			int
		username	string
		password	string
	)

	for rows.Next() {
		//数据指针，会把得到的数据，往刚才id 和 lvs引入
		rerr := rows.Scan(&id, &username, &password)
		if rerr == nil {
			// fmt.Println("id:",strconv.Itoa(id) + "        username:" + username +"        password:" + password)
			fmt.Printf("id: %s \t name: %s \t\t\tpwd: %s \n", strconv.Itoa(id), username, password)
		}
	}

}
