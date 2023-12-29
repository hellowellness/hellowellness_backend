package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//------------------------------------------------------------
// Web Server settings start
//------------------------------------------------------------

func WebServerPort() (string, error) {
	port := os.Getenv("HELLOWELLNESS_DB_PORT")
	if port == "" {
		port = "8081"
	}
	return ":" + port, nil
}

//------------------------------------------------------------
// Web Server settings end
//------------------------------------------------------------

// ------------------------------------------------------------
// DB settings start
// ------------------------------------------------------------
func DBPort() string {
	port := os.Getenv("HELLOWELLNESS_DB_PORT")
	if port == "" {
		port = "33306"
	}
	return port
}

func DBuser() string {
	user := os.Getenv("HELLOWELLNESS_DB_USER")
	if user == "" {
		user = "root"
	}
	return user
}

func DBpassword() string {
	password := os.Getenv("HELLOWELLNESS_DB_PASSWORD")
	if password == "" {
		password = "123"
	}
	return password
}

func DBip() string {
	password := os.Getenv("HELLOWELLNESS_DB_IP")
	if password == "" {
		password = "127.0.0.1" // localhost
	}
	return password
}

func DBName() string {
	password := os.Getenv("HELLOWELLNESS_DB_DBNAME")
	if password == "" {
		password = "hellowellness"
	}
	return password
}

func ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBuser(), DBpassword(), DBip(), DBPort(), DBName())

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

//------------------------------------------------------------
// DB settings end
//------------------------------------------------------------
