package uwork

import (
	"log"

	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/app/repository"
	"github.com/DmiAS/bd_course/internal/app/repository/orm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func (u *Unit) WithRole(role models.Role) UnitOfWork {
	db := getConnection(role)
	return &Unit{
		db: db,
	}
}

func (u *Unit) WithTransaction(f func(u UnitOfWork) error) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		un := &Unit{db: tx}
		if err := f(un); err != nil {
			return err
		}
		return nil
	})
}

func (u Unit) GetUserRepository() repository.IUserRepository {
	ur := orm.NewUserRepository(u.db)
	return ur
}

func (u Unit) GetWorkerRepository() repository.IWorkerRepository {
	wr := orm.NewWorkerRepository(u.db)
	return wr
}

func (u Unit) GetAuthRepository() repository.IAuthRepository {
	ar := orm.NewAuthRepository(u.db)
	return ar
}

func (u Unit) GetProjectRepository() repository.IProjectRepository {
	pr := orm.NewProjectRepository(u.db)
	return pr
}

func (u Unit) GetThreadsRepository() repository.IThreadRepository {
	tr := orm.NewThreadRepository(u.db)
	return tr
}

func (u Unit) GetCampaignsRepository() repository.ICampaignRepository {
	cr := orm.NewCampaignRepository(u.db)
	return cr
}

func getConnection(role models.Role) *gorm.DB {
	//conn, ok := u.conns[role]
	//if !ok{
	//	return nil, errors.New("invalid access role")
	//}
	//return conn, nil
	return conn
}
