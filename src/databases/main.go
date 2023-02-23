package databases

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

// SQLConnect DB接続
func GormConnect() (database *gorm.DB) {
    DBMS := "mysql"                   // MySQL
    PROTOCOL := "tcp(container_db)"   // MySQLコンテナ名
    DBNAME := "golang"                // テーブル名
    USER := "akidon"                  // MySQLユーザー名
    PASS := "12345"                   // パスワード

    CONNECT := USER + ":" + PASS + "@" +PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"  // 修正!!
    db, err := gorm.Open(DBMS, CONNECT)
    if err != nil {
        panic(err.Error())
    } else {
        fmt.Println("DB接続成功")
    }
    return db
}