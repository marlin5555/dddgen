// Package req Code generated, DO NOT EDIT.
package req

import "github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/base"

// get create update delete req for Application
type (
	// GetApplicationsNoForeign get Applications no foreign key
	GetApplicationsNoForeign interface {
		base.FuzzyApplicationName
		base.FuzzyApplicationNickname
		base.ApplicationID
		base.ApplicationIDs
		base.ApplicationName
		base.ApplicationNames
	}
	// GetApplicationsPure get Applications pure, 仅仅作用在单一实体表上
	GetApplicationsPure interface {
		base.FuzzyApplicationName
		base.FuzzyApplicationNickname
		base.ApplicationID
		base.ApplicationIDs
		base.ApplicationName
		base.ApplicationNames
	}
	// GetApplicationsPO get Applications po, 用于 persistence 上的 get 操作
	GetApplicationsPO interface {
		GetApplicationsPure
		base.PageOrderOperator
	}

	// GetApplicationsReq get Applications req, 用于对外 service 中使用
	GetApplicationsReq interface {
		GetApplicationsPure
		base.Fields
		base.PageOrderOperator
		base.Validator
	}

	DeleteApplicationReq interface {
		base.ApplicationID
		base.Operator
		base.Validator
	}

	UpdateApplicationReq interface {
		base.ApplicationID
		base.ApplicationName
		base.ApplicationNickname
		base.Operator
		base.Validator
	}
	CreateApplicationReq interface {
		base.ApplicationID
		base.ApplicationName
		base.ApplicationNickname
		base.Operator
		base.Validator
	}
)
