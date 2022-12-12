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

// applicationGetter impl entity.ApplicationGetter
type applicationGetter struct {
	value entity.Application
}

// GetApplication return entity.Application
func (g applicationGetter) GetApplication() entity.Application {
	return g.value
}

// applicationEntity application entity impl entity.Application
type applicationEntity struct {
	base.CreatedAt
	base.Creator
	base.DeletedAt
	base.ApplicationDescription
	base.ApplicationID
	base.ApplicationName
	base.ApplicationNickname
	base.UpdatedAt
	base.Updater
}

func (e *applicationEntity) setCreatedAt(pred bool, value interface{}) {
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
func (e *applicationEntity) setCreator(pred bool, value interface{}) {
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
func (e *applicationEntity) setDeletedAt(pred bool, value interface{}) {
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
func (e *applicationEntity) setApplicationDescription(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.ApplicationDescription); ok {
			e.ApplicationDescription = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.ApplicationDescription", value)
	}
	// 零值处理
	e.ApplicationDescription = zero.ApplicationDescription{}
}
func (e *applicationEntity) setApplicationID(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.ApplicationID); ok {
			e.ApplicationID = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.ApplicationID", value)
	}
	// 零值处理
	e.ApplicationID = zero.ApplicationID{}
}
func (e *applicationEntity) setApplicationName(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.ApplicationName); ok {
			e.ApplicationName = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.ApplicationName", value)
	}
	// 零值处理
	e.ApplicationName = zero.ApplicationName{}
}
func (e *applicationEntity) setApplicationNickname(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.ApplicationNickname); ok {
			e.ApplicationNickname = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.ApplicationNickname", value)
	}
	// 零值处理
	e.ApplicationNickname = zero.ApplicationNickname{}
}
func (e *applicationEntity) setUpdatedAt(pred bool, value interface{}) {
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
func (e *applicationEntity) setUpdater(pred bool, value interface{}) {
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

// buildApplication build entity
func buildApplication(checker fieldChecker, po po.Application) *applicationEntity {
	e := &applicationEntity{}
	t := newTable(persistence.T_Application)
	for _, tup := range []tuple{
		{pred: t.col("created_at").check(checker.hasColumn), attr: po, do: e.setCreatedAt},
		{pred: t.col("creator").check(checker.hasColumn), attr: po, do: e.setCreator},
		{pred: t.col("deleted_at").check(checker.hasColumn), attr: po, do: e.setDeletedAt},
		{pred: t.col("description").check(checker.hasColumn), attr: po, do: e.setApplicationDescription},
		{pred: t.col("id").check(checker.hasColumn), attr: po, do: e.setApplicationID},
		{pred: t.col("name").check(checker.hasColumn), attr: po, do: e.setApplicationName},
		{pred: t.col("nickname").check(checker.hasColumn), attr: po, do: e.setApplicationNickname},
		{pred: t.col("updated_at").check(checker.hasColumn), attr: po, do: e.setUpdatedAt},
		{pred: t.col("updater").check(checker.hasColumn), attr: po, do: e.setUpdater},
	} {
		builderAppend(tup)
	}
	return e
}
