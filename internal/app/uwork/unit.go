package uwork

import (
	"context"
	"log"

	"github.com/DmiAS/bd_course/internal/app/config"
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/DmiAS/bd_course/internal/app/repository"
	"github.com/DmiAS/bd_course/internal/app/repository/orm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Unit struct {
	currentDB *gorm.DB
	conns     map[models.Role]*gorm.DB
}

func New(ctx context.Context, cfg config.DBConfig) (*Unit, error) {
	adminConn, err := gorm.Open(postgres.Open(cfg.AdminDSN))
	if err != nil {
		return nil, err
	}
	workerConn, err := gorm.Open(postgres.Open(cfg.WorkerDSN))
	if err != nil {
		return nil, err
	}
	clientConn, err := gorm.Open(postgres.Open(cfg.ClientDSN))
	if err != nil {
		return nil, err
	}
	dummyConn, err := gorm.Open(postgres.Open(cfg.DummyDSN))
	if err != nil {
		return nil, err
	}
	conns := createMap(adminConn, workerConn, clientConn, dummyConn)
	go func() {
		<-ctx.Done()
		for role, conn := range conns {
			closeConn(conn, role)
		}
	}()
	return &Unit{conns: conns, currentDB: dummyConn}, nil
}

func closeConn(conn *gorm.DB, connType models.Role) {
	db, err := conn.DB()
	if err != nil {
		log.Printf("can't get db connection for %s - %s", connType, err.Error())
	}
	if err := db.Close(); err != nil {
		log.Printf("can't close db connection for %s - %s", connType, err.Error())
	}
}

func (u *Unit) WithRole(role models.Role) UnitOfWork {
	db := u.getConnection(role)
	return &Unit{
		currentDB: db,
		conns:     u.conns,
	}
}

func (u *Unit) WithTransaction(f func(u UnitOfWork) error) error {
	return u.currentDB.Transaction(func(tx *gorm.DB) error {
		un := &Unit{currentDB: tx, conns: u.conns}
		if err := f(un); err != nil {
			return err
		}
		return nil
	})
}

func (u Unit) GetUserRepository() repository.IUserRepository {
	ur := orm.NewUserRepository(u.currentDB)
	return ur
}

func (u Unit) GetWorkerRepository() repository.IWorkerRepository {
	wr := orm.NewWorkerRepository(u.currentDB)
	return wr
}

func (u Unit) GetAuthRepository() repository.IAuthRepository {
	ar := orm.NewAuthRepository(u.currentDB)
	return ar
}

func (u Unit) GetProjectRepository() repository.IProjectRepository {
	pr := orm.NewProjectRepository(u.currentDB)
	return pr
}

func (u Unit) GetThreadsRepository() repository.IThreadRepository {
	tr := orm.NewThreadRepository(u.currentDB)
	return tr
}

func (u Unit) GetCampaignsRepository() repository.ICampaignRepository {
	cr := orm.NewCampaignRepository(u.currentDB)
	return cr
}

func (u *Unit) getConnection(role models.Role) *gorm.DB {
	if conn, ok := u.conns[role]; !ok {
		return u.conns[models.DummyRole]
	} else {
		return conn
	}
}

func createMap(adminConn, workerConn, clientConn, dummyConn *gorm.DB) map[models.Role]*gorm.DB {
	return map[models.Role]*gorm.DB{
		models.AdminRole:  adminConn,
		models.ClientRole: clientConn,
		models.WorkerRole: workerConn,
		models.DummyRole:  dummyConn,
	}
}
