package uwork

import (
	"github.com/DmiAS/bd_course/internal/app/repository"
	"gorm.io/gorm"
)

type Unit struct {
	cr repository.IClientRepository
	wr repository.IWorkerRepository
	ar repository.IAuthRepository

	db *gorm.DB
}

func New(
	cr repository.IClientRepository,
	wr repository.IWorkerRepository,
	ar repository.IAuthRepository,
) *Unit {
	return &Unit{cr: cr, wr: wr, ar: ar}
}

func (u *Unit) WithRole(role Role) (*Unit, error) {
	db, err := getConnection(role)
	if err != nil {
		return nil, err
	}
	return &Unit{
		cr: u.cr,
		wr: u.wr,
		ar: u.ar,
		db: db,
	}, nil
}

func (u *Unit) WithTransaction() *Unit {
	tx := u.db.Begin()
	return &Unit{
		cr: u.cr,
		wr: u.wr,
		ar: u.ar,
		db: tx,
	}
}

func (u Unit) Commit() {
	u.db.Commit()
}

func (u Unit) Rollback() {
	u.db.Rollback()
}

func (u Unit) GetClientRepository() repository.IClientRepository {
	return u.cr
}

func (u Unit) GetWorkerRepository() repository.IWorkerRepository {
	return u.wr
}

func (u Unit) GetAuthRepository() repository.IAuthRepository {
	return u.ar
}

func getConnection(role Role) (*gorm.DB, error) {
	//conn, ok := u.conns[role]
	//if !ok{
	//	return nil, errors.New("invalid access role")
	//}
	//return conn, nil
	return nil, nil
}
