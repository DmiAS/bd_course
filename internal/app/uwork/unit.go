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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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

func (u *Unit) WithTransaction() UnitOfWork {
	tx := u.db.Begin()
	return &Unit{
		//cr: u.cr,
		//wr: u.wr,
		//ar: u.ar,
		db: tx,
	}
}

func (u Unit) Commit() {
	u.db.Commit()
}

func (u Unit) Rollback() {
	u.db.Rollback()
}

//func (u Unit) GetClientRepository() repository.IClientRepository {
//	return u.cr
//}

func (u Unit) GetWorkerRepository() repository.IWorkerRepository {
	wr := orm.NewWorkerRepository(conn)
	return wr
}

func (u Unit) GetAuthRepository() repository.IAuthRepository {
	ar := orm.NewAuthRepository(conn)
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
