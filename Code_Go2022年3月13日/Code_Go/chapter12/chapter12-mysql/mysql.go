package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 连接数据库
	db, err := sql.Open("mysql", "root:123456@/go_test")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	fmt.Println("连接成功！！")

	// 创建test表
	_, err = db.Exec("CREATE TABLE `test` (" +
		"`id` bigint(20) NOT NULL AUTO_INCREMENT," +
		"`name` varchar(45) DEFAULT ''," +
		"`age` int(11) NOT NULL DEFAULT '0'," +
		"`gender` tinyint(3) NOT NULL DEFAULT '0'," +
		"PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;")
	if err != nil {
		panic(err)
	}
	fmt.Println("表创建成功")

	//插入数据
	stmtInsert, err := db.Prepare("INSERT INTO test SET name=?,age=?,gender=?")
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := stmtInsert.Exec("三酷猫", "18", "0")
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println(id)

	// 查找数据
	type User struct {
		Name string `db:"name"`
		Id   int    `db:"id"`
		Age  int    `db:"age"`
		Sex  int    `db:"sex"`
	}
	var user User
	//单行查询
	err = db.QueryRow("SELECT * FROM test WHERE id=5").Scan(&user.Id, &user.Name, &user.Age, &user.Sex)
	fmt.Println(user)

	//结果集查询
	rows, e := db.Query("SELECT * FROM test WHERE age = 18")
	if e != nil {
		panic(e)
	}
	for rows.Next() {
		e := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Sex)
		if e != nil {
			panic(e)
		}
		fmt.Println(user)
	}
	err = rows.Close()
	if err != nil {
		panic(err)
	}

	// 删除数据
	stmtDelete, err := db.Prepare("DELETE FROM test WHERE id=?")
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err = stmtDelete.Exec("1")
	id, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println(id)

	//修改数据
	stmtUpdate, err := db.Prepare("UPDATE test SET age=? WHERE id=2")
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err = stmtUpdate.Exec("24")
	id, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println(id)

	//事务添加5000条数据
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	//事务中要执行的具体操作
	i := 0
	for i < 5000 {
		stmtTransaction, err := db.Prepare("INSERT INTO test SET name=?,age=?,gender=?")
		if err != nil {
			panic(err)
		}
		//设置参数以及执行sql语句
		_, err = stmtTransaction.Exec("三酷猫", "18", "0")
		if err != nil {
			panic(err)
		}
		i++
	}
	//提交事务
	err = tx.Commit()
	if err != nil {
		panic(err)
	}

}
