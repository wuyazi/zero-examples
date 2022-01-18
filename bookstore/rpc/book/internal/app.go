package internal

import (
	"context"
	"database/sql"
	"fmt"
	"gitee.com/wuyazi2018/gddd"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tal-tech/go-zero/core/logx"
)

var _repository *gddd.Repository

func Repository() *gddd.Repository {
	if _repository == nil {
		panic(fmt.Errorf("get users repository failed, it is not exist"))
	}
	return _repository
}

func init() {
	logx.DisableStat()
	err := newDomainEventBus(context.TODO())
	if err != nil {
	}
	err = newDomainEventStore()
	if err != nil {
	}
	err = NewRepository()
	if err != nil {
	}
}

func NewRepository() (err error) {
	repository, repositoryErr := gddd.NewRepository(&gddd.RepositoryConfig{
		EventStore:   usersEventStore(),
		EventBus:     usersEventBus(),
		SaveListener: &userRepositorySaveListener{},
	})
	if repositoryErr != nil {
		err = fmt.Errorf("new domain repository failed, create repository failed")
		return
	}
	_repository = repository
	return
}

type userRepositorySaveListener struct{}

func (l *userRepositorySaveListener) Handle(ctx context.Context, aggregate gddd.Aggregate) {
	//fmt.Println("\n------- userRepositorySaveListener --------")
}

// EventStore

var _eventStore gddd.EventStore

func usersEventStore() gddd.EventStore {
	if _eventStore == nil {
		panic(fmt.Errorf("get users event store failed, it is not exist"))
	}
	return _eventStore
}

func newDomainEventStore() (err error) {
	var db *sql.DB
	db, err = sql.Open("mysql", "root:123123@tcp(localhost:3306)/test1?charset=utf8mb4&parseTime=true")
	if err != nil {
		return fmt.Errorf("openDB error: %w", err)
	}
	_eventStore, err = gddd.NewMysqlEventStore(context.TODO(), "users", &gddd.MysqlEventStoreConfig{
		DomainEventTableName:          "test1.users_domain_events",
		DomainSnapshotTableName:       "test1.users_domain_snapshot",
		DB:                            db,
		LoadAggregateFromDatabaseFunc: nil,
	})
	return
}

// EventBus

var _usersEventBus gddd.EventBus

func usersEventBus() gddd.EventBus {
	if _usersEventBus == nil {
		panic(fmt.Errorf("get users event bus failed, it is not exist"))
	}
	return _usersEventBus
}

func newDomainEventBus(ctx context.Context) (err error) {
	_usersEventBus = gddd.NewLocalEventBus("users_aggregate")
	err = _usersEventBus.Start(ctx)
	if err != nil {
		panic(fmt.Errorf("start users event bus failed, err: %w", err))
	}
	return
}
