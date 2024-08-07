package main



require (
filippo.io/edwards25519 v1.1.0 // indirect
github.com/go-sql-driver/mysql v1.8.1 // indirect
github.com/jinzhu/inflection v1.0.0 // indirect
github.com/jinzhu/now v1.1.5 // indirect
golang.org/x/text v0.16.0 // indirect
gorm.io/driver/mysql v1.5.7 // indirect
gorm.io/gorm v1.25.11 // indirect
)
import (
"gorm.io/driver/mysql"
"gorm.io/gorm"
)
func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
