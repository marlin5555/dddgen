package test

import (
	"context"
	"sync"
	"testing"

	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/req"
	"github.com/marlin5555/dddgen/gen/example-1/internal/factory"
	config "github.com/marlin5555/dddgen/gen/example-1/pkg/conf"
	"github.com/marlin5555/dddgen/gen/example-1/pkg/log"
	"github.com/marlin5555/dddgen/gen/example-1/pkg/log/zap"

	"github.com/stretchr/testify/suite"
)

func init() {
	log.SetLogger(zap.InitLog("", "info"))
}

// TestEventTestSuite ...
func TestEventTestSuite(t *testing.T) {
	suite.Run(t, new(accountTestSuite))
}

type accountTestSuite struct {
	suite.Suite
}

// clearDB 保证完成 数据清理
func (s *accountTestSuite) clearDB() {
	wg := sync.WaitGroup{}
	//wg.Add(1)
	//factory.SQLiteDAOFactory().EventBusDAO().Transaction(func(tx *gorm.DB) error {
	//	tx.Model(&po.EventBus{}).Where("1=1").Delete("")
	//	tx.Model(&po.EventType{}).Where("1=1").Delete("")
	//	tx.Model(&po.TopicInfo{}).Where("1=1").Delete("")
	//	tx.Model(&po.AppInfo{}).Where("1=1").Delete("")
	//	tx.Model(&po.Secret{}).Where("1=1").Delete("")
	//	tx.Model(&po.Source{}).Where("1=1").Delete("")
	//	tx.Model(&po.WorkspacePO{}).Where("1=1").Delete("")
	//	wg.Done()
	//	return nil
	//})
	wg.Wait()
}

func (s *accountTestSuite) TearDownTest() {
	s.clearDB()
}

type caccount struct {
	n  string
	nn string
}

func (c caccount) GetAccountName() string {
	return c.n
}

func (c caccount) GetAccountNickname() string {
	return c.nn
}

func (c caccount) GetRtx() string {
	return "xubian"
}

func (c caccount) Validate() error {
	return nil
}

type raccount struct {
	fn  string
	n   string
	nn  string
	fnn string
}

func (r raccount) GetFuzzyAccountName() string {
	return r.fn
}

func (r raccount) GetFuzzyAccountNickname() string {
	return r.fnn
}

func (r raccount) GetAccountID() string {
	return ""
}

func (r raccount) GetAccountIDs() []string {
	return nil
}

func (r raccount) GetAccountName() string {
	return r.n
}

func (r raccount) GetAccountNames() []string {
	return nil
}

func (r raccount) GetAccountNickname() string {
	return r.nn
}

func (r raccount) GetFieldNames() []string {
	return nil
}

func (r raccount) GetOffset() int {
	return 0
}

func (r raccount) GetLimit() int {
	return 0
}

func (r raccount) OrderStr() string {
	return ""
}

func (r raccount) GetRtx() string {
	return "xubian"
}

func (r raccount) Validate() error {
	return nil
}

// account 添加/查询场景
func (s *accountTestSuite) Test_Account_Add_Get() {
	persistenceFactory := factory.DBPersistenceFactory(config.EmptyConfig{})
	repositoryFactory := factory.RepositoryFactory(persistenceFactory)
	accountRepo := repositoryFactory.AccountRepo()

	tests := []struct {
		name     string
		ctx      context.Context
		cAccount req.CreateAccountReq
		cWantErr bool
		rWantErr bool
	}{{
		ctx:      context.Background(),
		name:     "添加-查询",
		cAccount: caccount{n: "nlike", nn: "nnlike"},
	}}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			_, err := accountRepo.CreateAccount(tt.ctx, tt.cAccount)
			if (err != nil) != tt.cWantErr {
				s.Failf(tt.name, "when create account, gotErr[%+v], wantErr[%+v]", err, tt.cWantErr)
			}
			// accountName 等值测试
			result, _, err := accountRepo.GetAccounts(tt.ctx, raccount{n: tt.cAccount.GetAccountName()})
			if (err != nil) != tt.rWantErr {
				s.Failf(tt.name, "when get account, gotErr[%+v], wantErr[%+v]", err, tt.rWantErr)
			}
			if len(result) != 1 {
				s.Failf(tt.name, "when get account, got[%+v], wantGot[%+v]", len(result), 1)
			}
			if tt.cAccount.GetAccountName() != result[0].GetAccountName() {
				s.Failf(tt.name, "when get account, got.AccountName[%+v], wantGot[%+v]", result[0].GetAccountName(), tt.cAccount.GetAccountName())
			}
			// accountName like测试
			result, _, err = accountRepo.GetAccounts(tt.ctx, raccount{fn: tt.cAccount.GetAccountName()[2:]})
			if (err != nil) != tt.rWantErr {
				s.Failf(tt.name, "when get account, gotErr[%+v], wantErr[%+v]", err, tt.rWantErr)
			}
			if len(result) != 1 {
				s.Failf(tt.name, "when get account, got[%+v], wantGot[%+v]", len(result), 1)
			}
			if tt.cAccount.GetAccountName() != result[0].GetAccountName() {
				s.Failf(tt.name, "when get account, got.AccountName[%+v], wantGot[%+v]", result[0].GetAccountName(), tt.cAccount.GetAccountName())
			}
			// accountNickName eq测试
			result, _, err = accountRepo.GetAccounts(tt.ctx, raccount{nn: tt.cAccount.GetAccountNickname()})
			if (err != nil) != tt.rWantErr {
				s.Failf(tt.name, "when get account, gotErr[%+v], wantErr[%+v]", err, tt.rWantErr)
			}
			if len(result) != 1 {
				s.Failf(tt.name, "when get account, got[%+v], wantGot[%+v]", len(result), 1)
			}
			if tt.cAccount.GetAccountName() != result[0].GetAccountName() {
				s.Failf(tt.name, "when get account, got.AccountName[%+v], wantGot[%+v]", result[0].GetAccountName(), tt.cAccount.GetAccountName())
			}
			// accountNickName like测试
			result, _, err = accountRepo.GetAccounts(tt.ctx, raccount{fnn: tt.cAccount.GetAccountNickname()[2:]})
			if (err != nil) != tt.rWantErr {
				s.Failf(tt.name, "when get account, gotErr[%+v], wantErr[%+v]", err, tt.rWantErr)
			}
			if len(result) != 1 {
				s.Failf(tt.name, "when get account, got[%+v], wantGot[%+v]", len(result), 1)
			}
			if tt.cAccount.GetAccountName() != result[0].GetAccountName() {
				s.Failf(tt.name, "when get account, got.AccountName[%+v], wantGot[%+v]", result[0].GetAccountName(), tt.cAccount.GetAccountName())
			}
		})
	}
}
