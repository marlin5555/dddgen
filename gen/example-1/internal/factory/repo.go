// Package factory Code generated, DO NOT EDIT.
package factory

import (
	"sync"

	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/ports/repo"
	implrepo "github.com/marlin5555/dddgen/gen/example-1/internal/domain/repo"
)

// RepoFactory 存储层数据访问对象工厂
type RepoFactory struct {
	accountRepo     repo.AccountRepository
	applicationRepo repo.ApplicationRepository
	eventRepo       repo.EventRepository
}

var (
	repoOnce    sync.Once
	repoFactory RepoFactory
)

// RepositoryFactory 构造 PersistenceFactory
func RepositoryFactory(pF *PersistenceFactory) *RepoFactory {
	repoOnce.Do(func() {
		repoFactory = RepoFactory{
			accountRepo:     implrepo.NewAccountRepo(pF.AccountDAO(), pF.PassportDAO(), pF.SecretDAO(), pF.TechRelationDAO()),
			applicationRepo: implrepo.NewApplicationRepo(pF.ApplicationDAO()),
			eventRepo:       implrepo.NewEventRepo(pF.ApplicationDAO(), pF.EventBusDAO(), pF.EventTypeDAO(), pF.PublicationDAO(), pF.SubscriptionDAO()),
		}
	})
	return &repoFactory
}
func (f *RepoFactory) AccountRepo() repo.AccountRepository {
	if f == nil {
		return nil
	}
	return f.accountRepo
}
func (f *RepoFactory) ApplicationRepo() repo.ApplicationRepository {
	if f == nil {
		return nil
	}
	return f.applicationRepo
}
func (f *RepoFactory) EventRepo() repo.EventRepository {
	if f == nil {
		return nil
	}
	return f.eventRepo
}
