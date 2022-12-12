// Package factory Code generated, DO NOT EDIT.
package factory

import (
	"sync"

	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/ports/infra/persistence"
	"github.com/marlin5555/dddgen/gen/example-1/internal/infra/persistence/db"
	"github.com/marlin5555/dddgen/gen/example-1/internal/infra/persistence/po"
	config "github.com/marlin5555/dddgen/gen/example-1/pkg/conf"
	"github.com/marlin5555/dddgen/gen/example-1/pkg/log"
)

// PersistenceFactory 存储层数据访问对象工厂
type PersistenceFactory struct {
	accountDao      persistence.AccountDAO
	applicationDao  persistence.ApplicationDAO
	eventBusDao     persistence.EventBusDAO
	eventTypeDao    persistence.EventTypeDAO
	passportDao     persistence.PassportDAO
	publicationDao  persistence.PublicationDAO
	secretDao       persistence.SecretDAO
	subscriptionDao persistence.SubscriptionDAO
	techRelationDao persistence.TechRelationDAO
}

var (
	dbOnce         sync.Once
	persistFactory PersistenceFactory
)

// DBPersistenceFactory 构造 PersistenceFactory
func DBPersistenceFactory(conf config.Config) *PersistenceFactory {
	dbOnce.Do(func() {
		// 如果没有传入 db config 则使用 SQLite
		mydb := db.NewSQLiteDB()
		if conf.DBConfig().Host != "" {
			mydb = db.NewMySQLDB(conf.DBConfig())
		}
		persistFactory = PersistenceFactory{
			accountDao:      db.NewAccountDAO(mydb),
			applicationDao:  db.NewApplicationDAO(mydb),
			eventBusDao:     db.NewEventBusDAO(mydb),
			eventTypeDao:    db.NewEventTypeDAO(mydb),
			passportDao:     db.NewPassportDAO(mydb),
			publicationDao:  db.NewPublicationDAO(mydb),
			secretDao:       db.NewSecretDAO(mydb),
			subscriptionDao: db.NewSubscriptionDAO(mydb),
			techRelationDao: db.NewTechRelationDAO(mydb),
		}
		var err error
		err = mydb.AutoMigrate(&po.Account{})
		log.InfofWithFuncName("Auto Migrate Account got err = %+v", err)
		err = mydb.AutoMigrate(&po.Application{})
		log.InfofWithFuncName("Auto Migrate Application got err = %+v", err)
		err = mydb.AutoMigrate(&po.EventBus{})
		log.InfofWithFuncName("Auto Migrate EventBus got err = %+v", err)
		err = mydb.AutoMigrate(&po.EventType{})
		log.InfofWithFuncName("Auto Migrate EventType got err = %+v", err)
		err = mydb.AutoMigrate(&po.Passport{})
		log.InfofWithFuncName("Auto Migrate Passport got err = %+v", err)
		err = mydb.AutoMigrate(&po.Publication{})
		log.InfofWithFuncName("Auto Migrate Publication got err = %+v", err)
		err = mydb.AutoMigrate(&po.Secret{})
		log.InfofWithFuncName("Auto Migrate Secret got err = %+v", err)
		err = mydb.AutoMigrate(&po.Subscription{})
		log.InfofWithFuncName("Auto Migrate Subscription got err = %+v", err)
		err = mydb.AutoMigrate(&po.TechRelation{})
		log.InfofWithFuncName("Auto Migrate TechRelation got err = %+v", err)
	})
	return &persistFactory
}

// AccountDAO 获取 persistence.AccountDAO
func (f *PersistenceFactory) AccountDAO() persistence.AccountDAO {
	if f == nil {
		return nil
	}
	return f.accountDao
}

// ApplicationDAO 获取 persistence.ApplicationDAO
func (f *PersistenceFactory) ApplicationDAO() persistence.ApplicationDAO {
	if f == nil {
		return nil
	}
	return f.applicationDao
}

// EventBusDAO 获取 persistence.EventBusDAO
func (f *PersistenceFactory) EventBusDAO() persistence.EventBusDAO {
	if f == nil {
		return nil
	}
	return f.eventBusDao
}

// EventTypeDAO 获取 persistence.EventTypeDAO
func (f *PersistenceFactory) EventTypeDAO() persistence.EventTypeDAO {
	if f == nil {
		return nil
	}
	return f.eventTypeDao
}

// PassportDAO 获取 persistence.PassportDAO
func (f *PersistenceFactory) PassportDAO() persistence.PassportDAO {
	if f == nil {
		return nil
	}
	return f.passportDao
}

// PublicationDAO 获取 persistence.PublicationDAO
func (f *PersistenceFactory) PublicationDAO() persistence.PublicationDAO {
	if f == nil {
		return nil
	}
	return f.publicationDao
}

// SecretDAO 获取 persistence.SecretDAO
func (f *PersistenceFactory) SecretDAO() persistence.SecretDAO {
	if f == nil {
		return nil
	}
	return f.secretDao
}

// SubscriptionDAO 获取 persistence.SubscriptionDAO
func (f *PersistenceFactory) SubscriptionDAO() persistence.SubscriptionDAO {
	if f == nil {
		return nil
	}
	return f.subscriptionDao
}

// TechRelationDAO 获取 persistence.TechRelationDAO
func (f *PersistenceFactory) TechRelationDAO() persistence.TechRelationDAO {
	if f == nil {
		return nil
	}
	return f.techRelationDao
}
