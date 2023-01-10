package clickhouse

import (
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/mailru/go-clickhouse"
	"github.com/shopspring/decimal"
)

type (
	User struct {
		UserId     string
		UserType   int64
		Name       string
		Email      string
		Balance    decimal.Decimal
		CreateTime int64
	}
)

var (
	err error
	c   = Config{
		Address:  "http://127.0.0.1:8123",
		Database: "house",
	}
)

func (slf *User) TableName() string {
	return "house.user"
}

func (slf *User) CreateTable() string {
	s := NewSession()
	_, err = s.Exec(fmt.Sprintf("CREATE TABLE If Not Exists %s  "+
		"(`user_id` String, "+
		"`user_type` UInt16, "+
		"`name` String, "+
		"`email` String, "+
		"`balance` Decimal128(18), "+
		"`create_time` UInt64 )"+
		"ENGINE = MergeTree()  PARTITION BY toYYYYMMDD(toDateTime(create_time/1000000000)) "+
		"ORDER BY (create_time)  SETTINGS index_granularity = 8192;", slf.TableName()))
	if err != nil {
		log.Fatal("create table err: " + err.Error())
	}
	return "house.user"
}

func (slf *User) Insert(records []*User) (err error) {
	s := NewSession().InsertInto(slf.TableName()).Columns("user_id", "user_type", "name",
		"email", "balance", "create_time")
	for _, v := range records {
		s.Record(v)
	}
	res, err := s.Exec()
	log.Printf("res: %+v, err: %v", res, err)
	return
}

func TestInit(t *testing.T) {
	if err = Init(c); err != nil {
		log.Fatal("init clickhouse err: " + err.Error())
	}
	balance, _ := decimal.NewFromString("0.123456789123456789")
	new(User).CreateTable()
	users := make([]*User, 0)
	for i := 0; i < 10000; i++ {
		u := User{
			UserId:     fmt.Sprintf("%d", i),
			UserType:   1,
			Name:       fmt.Sprintf("%10d", 1),
			Email:      "",
			Balance:    balance.Add(decimal.NewFromInt(int64(i))),
			CreateTime: time.Now().UnixNano(),
		}
		users = append(users, &u)
	}
	err = new(User).Insert(users)

	s := NewSession()
	countQ := s.SelectBySql("select count(*) as total from user")
	var total int
	_, err = countQ.Load(&total)
	log.Printf("err: %v, total: %d", err, total)
}
