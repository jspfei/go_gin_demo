package tool

import (
	"CloudRestaurant/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DbEngine *Orm

type Orm struct {
	*xorm.Engine
}

func OrmEngine(cfg *Config) (*Orm, error) {
	databse := cfg.Database
	conn := databse.User + ":" + databse.Password + "@tcp(" + databse.Host + ":" + databse.Port + ")/" + databse.DbName + "?charset=" + databse.Charset
	engine, err := xorm.NewEngine(databse.Dirver, conn)
	if err != nil {
		return nil, err
	}

	engine.ShowSQL(databse.ShowSql)
	err = engine.Sync2(new(model.SmsCode))
	if err != nil {
		return nil, err
	}

	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm
	return orm, nil
}
