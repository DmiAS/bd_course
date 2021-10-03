package uwork

import (
	"github.com/DmiAS/bd_course/internal/app/repository"
	"github.com/DmiAS/bd_course/internal/app/repository/orm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func init() {
	dsn := "host=localhost user=postgres password=password dbname=agency sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: false})
	if err != nil {
		log.Fatalln(err)
	}
	conn = db
}

var conn *gorm.DB

type Unit struct {
	db *gorm.DB
}

func New() *Unit {
	return &Unit{}
}

func (u *Unit) WithRole(role Role) UnitOfWork {
	db := getConnection(role)
	return &Unit{
		//wr: u.wr,
		//ar: u.ar,
		db: db,
	}
}

//func (u *Unit) WithTransaction() UnitOfWork {
//	//db := u.db.Session(&gorm.Session{SkipDefaultTransaction: true})
//	//u.db.Transaction()
//	tx := u.db.Begin()
//	fmt.Println(tx.Error)
//	return &Unit{
//		//cr: u.cr,
//		//wr: u.wr,
//		//ar: u.ar,
//		db: tx,
//	}
//}

func (u *Unit) WithTransaction(f func(u UnitOfWork) error) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		un := &Unit{db: tx}
		if err := f(un); err != nil {
			return err
		}
		return nil
	})
}

//func (u Unit) GetClientRepository() repository.IClientRepository {
//	return u.cr
//}

func (u Unit) GetWorkerRepository() repository.IWorkerRepository {
	wr := orm.NewWorkerRepository(u.db)
	return wr
}

func (u Unit) GetAuthRepository() repository.IAuthRepository {
	ar := orm.NewAuthRepository(u.db)
	return ar
}

func getConnection(role Role) *gorm.DB {
	//conn, ok := u.conns[role]
	//if !ok{
	//	return nil, errors.New("invalid access role")
	//}
	//return conn, nil
	return conn
}
