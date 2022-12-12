// Package base Code generated, DO NOT EDIT.
package base

import "time"

// CreatedAt interface has GetCreatedAt
type CreatedAt interface {
	GetCreatedAt() time.Time
}

// Creator interface has GetCreator
type Creator interface {
	GetCreator() string
}

// DeletedAt interface has GetDeletedAt
type DeletedAt interface {
	GetDeletedAt() int64
}

// UpdatedAt interface has GetUpdatedAt
type UpdatedAt interface {
	GetUpdatedAt() time.Time
}

// Updater interface has GetUpdater
type Updater interface {
	GetUpdater() string
}

type (
	// RspCode interface return int32
	RspCode interface {
		GetCode() int32
	}

	// RspMsg interface return string
	RspMsg interface {
		GetMsg() string
	}

	// RspErr interface return error
	RspErr interface {
		GetError() error
		IsOK() bool
		NotOK() bool
	}
)
type (
	// Rtx interface GetRtx return string
	Rtx interface {
		GetRtx() string
	}

	// Operator combination interface: Rtx
	Operator interface {
		Rtx
	}
)

type (
	// Offset interface GetOffset return int
	Offset interface {
		GetOffset() int
	}

	// Limit interface GetLimit return int
	Limit interface {
		GetLimit() int
	}

	// Order interface OrderStr return string
	Order interface {
		OrderStr() string
	}

	// PageQuery combination interface: Offset Limit
	PageQuery interface {
		Offset
		Limit
	}

	// PageOrder combination interface: PageQuery Order
	PageOrder interface {
		PageQuery
		Order
	}

	// PageOrderOperator combination interface: PageOrder Operator
	PageOrderOperator interface {
		PageOrder
		Operator
	}
)

// Fields interface return field names []string
type Fields interface {
	GetFieldNames() []string
}

// Validator interface Validate when ok return nil, others return error
type Validator interface {
	Validate() error
}
