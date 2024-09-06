package gorm

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlOption struct {
	User     string
	Password string
	Host     string
	Name     string
	Debug    bool
}

func NewOption(conf *viper.Viper) MysqlOption {
	return MysqlOption{
		User:     conf.GetString("mysql.user"),
		Password: conf.GetString("mysql.password"),
		Host:     conf.GetString("mysql.host"),
		Name:     conf.GetString("mysql.name"),
		Debug:    conf.GetBool("mysql.debug"),
	}
}

func NewMysql(mysqlOption MysqlOption) *gorm.DB {
	orm, err := gorm.Open(mysql.Open(
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			mysqlOption.User,
			mysqlOption.Password,
			mysqlOption.Host,
			mysqlOption.Name)))
	if err != nil {
		log.Fatalf("mysql connect error: %v", err)
	}

	if orm.Error != nil {
		log.Fatalf("mysql error: %v", orm.Error)
	}

	if mysqlOption.Debug {
		orm = orm.Debug()
	}

	return orm
}
