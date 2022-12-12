// Package potoentity Code generated, DO NOT EDIT.
package potoentity

import (
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/base"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/zero"
	"github.com/marlin5555/dddgen/gen/example-1/internal/infra/persistence"
	"github.com/marlin5555/dddgen/gen/example-1/internal/infra/persistence/po"
	"github.com/marlin5555/dddgen/gen/example-1/pkg/log"
)

// accountGetter impl entity.AccountGetter
type accountGetter struct {
	value entity.Account
}

// GetAccount return entity.Account
func (g accountGetter) GetAccount() entity.Account {
	return g.value
}

// passportGetter impl entity.PassportGetter
type passportGetter struct {
	value entity.Passport
}

// GetPassport return entity.Passport
func (g passportGetter) GetPassport() entity.Passport {
	return g.value
}

// secretGetter impl entity.SecretGetter
type secretGetter struct {
	value entity.Secret
}

// GetSecret return entity.Secret
func (g secretGetter) GetSecret() entity.Secret {
	return g.value
}

// techRelationGetter impl entity.TechRelationGetter
type techRelationGetter struct {
	value entity.TechRelation
}

// GetTechRelation return entity.TechRelation
func (g techRelationGetter) GetTechRelation() entity.TechRelation {
	return g.value
}

// studentGetter impl entity.StudentGetter
type studentGetter struct {
	value entity.Account
}

// GetStudent return entity.Account
func (g studentGetter) GetStudent() entity.Account {
	return g.value
}

// teacherGetter impl entity.TeacherGetter
type teacherGetter struct {
	value entity.Account
}

// GetTeacher return entity.Account
func (g teacherGetter) GetTeacher() entity.Account {
	return g.value
}

// accountEntity account entity impl entity.Account
type accountEntity struct {
	base.CreatedAt
	base.Creator
	base.DeletedAt
	base.AccountDescription
	base.AccountID
	secretGetter
	passportGetter
	base.AccountName
	base.AccountNickname
	base.UpdatedAt
	base.Updater
}

func (e *accountEntity) setCreatedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.CreatedAt); ok {
			e.CreatedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.CreatedAt", value)
	}
	// 零值处理
	e.CreatedAt = zero.CreatedAt{}
}
func (e *accountEntity) setCreator(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.Creator); ok {
			e.Creator = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.Creator", value)
	}
	// 零值处理
	e.Creator = zero.Creator{}
}
func (e *accountEntity) setDeletedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.DeletedAt); ok {
			e.DeletedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.DeletedAt", value)
	}
	// 零值处理
	e.DeletedAt = zero.DeletedAt{}
}
func (e *accountEntity) setAccountDescription(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.AccountDescription); ok {
			e.AccountDescription = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.AccountDescription", value)
	}
	// 零值处理
	e.AccountDescription = zero.AccountDescription{}
}
func (e *accountEntity) setAccountID(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.AccountID); ok {
			e.AccountID = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.AccountID", value)
	}
	// 零值处理
	e.AccountID = zero.AccountID{}
}
func (e *accountEntity) setAccountName(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.AccountName); ok {
			e.AccountName = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.AccountName", value)
	}
	// 零值处理
	e.AccountName = zero.AccountName{}
}
func (e *accountEntity) setAccountNickname(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.AccountNickname); ok {
			e.AccountNickname = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.AccountNickname", value)
	}
	// 零值处理
	e.AccountNickname = zero.AccountNickname{}
}
func (e *accountEntity) setUpdatedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.UpdatedAt); ok {
			e.UpdatedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.UpdatedAt", value)
	}
	// 零值处理
	e.UpdatedAt = zero.UpdatedAt{}
}
func (e *accountEntity) setUpdater(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.Updater); ok {
			e.Updater = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.Updater", value)
	}
	// 零值处理
	e.Updater = zero.Updater{}
}

// buildAccount build entity
func buildAccount(checker fieldChecker, po po.Account) *accountEntity {
	e := &accountEntity{}
	t := newTable(persistence.T_Account)
	for _, tup := range []tuple{
		{pred: t.col("created_at").check(checker.hasColumn), attr: po, do: e.setCreatedAt},
		{pred: t.col("creator").check(checker.hasColumn), attr: po, do: e.setCreator},
		{pred: t.col("deleted_at").check(checker.hasColumn), attr: po, do: e.setDeletedAt},
		{pred: t.col("description").check(checker.hasColumn), attr: po, do: e.setAccountDescription},
		{pred: t.col("id").check(checker.hasColumn), attr: po, do: e.setAccountID},
		{pred: t.col("name").check(checker.hasColumn), attr: po, do: e.setAccountName},
		{pred: t.col("nickname").check(checker.hasColumn), attr: po, do: e.setAccountNickname},
		{pred: t.col("updated_at").check(checker.hasColumn), attr: po, do: e.setUpdatedAt},
		{pred: t.col("updater").check(checker.hasColumn), attr: po, do: e.setUpdater},
	} {
		builderAppend(tup)
	}
	return e
}

// passportEntity passport entity impl entity.Passport
type passportEntity struct {
	accountGetter
	base.CreatedAt
	base.Creator
	base.DeletedAt
	base.PassportID
	base.PassportNationality
	base.PassportPostalAddress
	base.UpdatedAt
	base.Updater
}

func (e *passportEntity) setCreatedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.CreatedAt); ok {
			e.CreatedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.CreatedAt", value)
	}
	// 零值处理
	e.CreatedAt = zero.CreatedAt{}
}
func (e *passportEntity) setCreator(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.Creator); ok {
			e.Creator = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.Creator", value)
	}
	// 零值处理
	e.Creator = zero.Creator{}
}
func (e *passportEntity) setDeletedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.DeletedAt); ok {
			e.DeletedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.DeletedAt", value)
	}
	// 零值处理
	e.DeletedAt = zero.DeletedAt{}
}
func (e *passportEntity) setPassportID(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.PassportID); ok {
			e.PassportID = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.PassportID", value)
	}
	// 零值处理
	e.PassportID = zero.PassportID{}
}
func (e *passportEntity) setPassportNationality(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.PassportNationality); ok {
			e.PassportNationality = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.PassportNationality", value)
	}
	// 零值处理
	e.PassportNationality = zero.PassportNationality{}
}
func (e *passportEntity) setPassportPostalAddress(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.PassportPostalAddress); ok {
			e.PassportPostalAddress = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.PassportPostalAddress", value)
	}
	// 零值处理
	e.PassportPostalAddress = zero.PassportPostalAddress{}
}
func (e *passportEntity) setUpdatedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.UpdatedAt); ok {
			e.UpdatedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.UpdatedAt", value)
	}
	// 零值处理
	e.UpdatedAt = zero.UpdatedAt{}
}
func (e *passportEntity) setUpdater(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.Updater); ok {
			e.Updater = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.Updater", value)
	}
	// 零值处理
	e.Updater = zero.Updater{}
}

// buildPassport build entity
func buildPassport(checker fieldChecker, po po.Passport) *passportEntity {
	e := &passportEntity{}
	t := newTable(persistence.T_Passport)
	for _, tup := range []tuple{
		{pred: t.col("created_at").check(checker.hasColumn), attr: po, do: e.setCreatedAt},
		{pred: t.col("creator").check(checker.hasColumn), attr: po, do: e.setCreator},
		{pred: t.col("deleted_at").check(checker.hasColumn), attr: po, do: e.setDeletedAt},
		{pred: t.col("id").check(checker.hasColumn), attr: po, do: e.setPassportID},
		{pred: t.col("nationality").check(checker.hasColumn), attr: po, do: e.setPassportNationality},
		{pred: t.col("postal_address").check(checker.hasColumn), attr: po, do: e.setPassportPostalAddress},
		{pred: t.col("updated_at").check(checker.hasColumn), attr: po, do: e.setUpdatedAt},
		{pred: t.col("updater").check(checker.hasColumn), attr: po, do: e.setUpdater},
	} {
		builderAppend(tup)
	}
	return e
}

// secretEntity secret entity impl entity.Secret
type secretEntity struct {
	accountGetter
	base.CreatedAt
	base.Creator
	base.DeletedAt
	base.SecretID
	base.SecretSecret
	base.UpdatedAt
	base.Updater
}

func (e *secretEntity) setCreatedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.CreatedAt); ok {
			e.CreatedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.CreatedAt", value)
	}
	// 零值处理
	e.CreatedAt = zero.CreatedAt{}
}
func (e *secretEntity) setCreator(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.Creator); ok {
			e.Creator = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.Creator", value)
	}
	// 零值处理
	e.Creator = zero.Creator{}
}
func (e *secretEntity) setDeletedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.DeletedAt); ok {
			e.DeletedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.DeletedAt", value)
	}
	// 零值处理
	e.DeletedAt = zero.DeletedAt{}
}
func (e *secretEntity) setSecretID(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.SecretID); ok {
			e.SecretID = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.SecretID", value)
	}
	// 零值处理
	e.SecretID = zero.SecretID{}
}
func (e *secretEntity) setSecretSecret(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.SecretSecret); ok {
			e.SecretSecret = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.SecretSecret", value)
	}
	// 零值处理
	e.SecretSecret = zero.SecretSecret{}
}
func (e *secretEntity) setUpdatedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.UpdatedAt); ok {
			e.UpdatedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.UpdatedAt", value)
	}
	// 零值处理
	e.UpdatedAt = zero.UpdatedAt{}
}
func (e *secretEntity) setUpdater(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.Updater); ok {
			e.Updater = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.Updater", value)
	}
	// 零值处理
	e.Updater = zero.Updater{}
}

// buildSecret build entity
func buildSecret(checker fieldChecker, po po.Secret) *secretEntity {
	e := &secretEntity{}
	t := newTable(persistence.T_Secret)
	for _, tup := range []tuple{
		{pred: t.col("created_at").check(checker.hasColumn), attr: po, do: e.setCreatedAt},
		{pred: t.col("creator").check(checker.hasColumn), attr: po, do: e.setCreator},
		{pred: t.col("deleted_at").check(checker.hasColumn), attr: po, do: e.setDeletedAt},
		{pred: t.col("id").check(checker.hasColumn), attr: po, do: e.setSecretID},
		{pred: t.col("secret").check(checker.hasColumn), attr: po, do: e.setSecretSecret},
		{pred: t.col("updated_at").check(checker.hasColumn), attr: po, do: e.setUpdatedAt},
		{pred: t.col("updater").check(checker.hasColumn), attr: po, do: e.setUpdater},
	} {
		builderAppend(tup)
	}
	return e
}

// techRelationEntity techRelation entity impl entity.TechRelation
type techRelationEntity struct {
	base.CreatedAt
	base.Creator
	base.DeletedAt
	base.TechRelationID
	studentGetter
	teacherGetter
	base.UpdatedAt
	base.Updater
}

func (e *techRelationEntity) setCreatedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.CreatedAt); ok {
			e.CreatedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.CreatedAt", value)
	}
	// 零值处理
	e.CreatedAt = zero.CreatedAt{}
}
func (e *techRelationEntity) setCreator(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.Creator); ok {
			e.Creator = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.Creator", value)
	}
	// 零值处理
	e.Creator = zero.Creator{}
}
func (e *techRelationEntity) setDeletedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.DeletedAt); ok {
			e.DeletedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.DeletedAt", value)
	}
	// 零值处理
	e.DeletedAt = zero.DeletedAt{}
}
func (e *techRelationEntity) setTechRelationID(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.TechRelationID); ok {
			e.TechRelationID = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.TechRelationID", value)
	}
	// 零值处理
	e.TechRelationID = zero.TechRelationID{}
}
func (e *techRelationEntity) setUpdatedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.UpdatedAt); ok {
			e.UpdatedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.UpdatedAt", value)
	}
	// 零值处理
	e.UpdatedAt = zero.UpdatedAt{}
}
func (e *techRelationEntity) setUpdater(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.Updater); ok {
			e.Updater = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.Updater", value)
	}
	// 零值处理
	e.Updater = zero.Updater{}
}

// buildTechRelation build entity
func buildTechRelation(checker fieldChecker, po po.TechRelation) *techRelationEntity {
	e := &techRelationEntity{}
	t := newTable(persistence.T_TechRelation)
	for _, tup := range []tuple{
		{pred: t.col("created_at").check(checker.hasColumn), attr: po, do: e.setCreatedAt},
		{pred: t.col("creator").check(checker.hasColumn), attr: po, do: e.setCreator},
		{pred: t.col("deleted_at").check(checker.hasColumn), attr: po, do: e.setDeletedAt},
		{pred: t.col("id").check(checker.hasColumn), attr: po, do: e.setTechRelationID},
		{pred: t.col("updated_at").check(checker.hasColumn), attr: po, do: e.setUpdatedAt},
		{pred: t.col("updater").check(checker.hasColumn), attr: po, do: e.setUpdater},
	} {
		builderAppend(tup)
	}
	return e
}
