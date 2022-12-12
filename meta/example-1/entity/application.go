package entity

// Application Table: t_application; Group: application; 应用定义
type Application struct {
	// use guid 36 char
	ID string `gorm:"column:id;type:varchar(36);primaryKey;NOT NULL" ddd:"exact;batch;upsert;may_s:EventType.MainAppID;s_may_s:Subscription.SubscriberID,Publication.PublisherID;simple" json:"id"`
	// name unique
	Name string `gorm:"column:name;type:varchar(255);NOT NULL;uniqueIndex:t_application_u1" ddd:"exact;batch;fuzzy;upsert;simple;detail" json:"name"`
	// nickname
	Nickname string `gorm:"column:nickname;type:varchar(255);" ddd:"fuzzy;upsert;simple;detail" json:"nickname"`
	// 应用描述
	Description string `gorm:"column:description;type:varchar(4096);default null" json:"description"`
	Trailer
}

// EventType Table: t_event_type; Group: event; 事件类型定义
type EventType struct {
	// use guid 36 char
	ID string `gorm:"column:id;type:varchar(36);primaryKey;NOT NULL" ddd:"exact;batch;upsert;simple;s_may_s:Subscription.EventTypeID"`
	// name unique
	Name string `gorm:"column:name;type:varchar(255);NOT NULL;uniqueIndex:t_event_type_u1" ddd:"exact;batch;fuzzy;upsert;simple;detail" json:"name"`
	// main app ref
	MainAppID string `gorm:"column:main_app_id;type:varchar(36);NOT NULL" ddd:"exact;batch;upsert;yam_s:Application.ID;role:ownership" json:"main_app_id"`
	// event bus ref
	EventBusID string `gorm:"column:event_bus_id;type:varchar(36);NOT NULL" ddd:"exact;batch;upsert;yam_s:EventBus.ID" json:"event_bus_id"`
	// 关系状态
	Status int `gorm:"column:status;type:int;NOT NULL;default:0" ddd:"exact;batch;upsert;simple;detail;status" json:"status"`
	// 事件类型描述
	Description string `gorm:"column:description;type:varchar(4096);default null" json:"description"`
	Trailer
}

// Subscription Table: t_subscription; Group: event; 订阅关系定义
type Subscription struct {
	// use guid 36 char
	ID string `gorm:"column:id;type:varchar(36);primaryKey;NOT NULL" ddd:"exact;batch;upsert;simple"`
	// main app ref
	SubscriberID string `gorm:"column:main_app_id;type:varchar(36);NOT NULL" ddd:"exact;batch;upsert;yam_s:Application.ID;role:subscriber" json:"subscriber_id"`
	// name unique
	EventTypeID string `gorm:"column:name;type:varchar(255);NOT NULL" ddd:"exact;batch;fuzzy;upsert;simple;detail;yam_s:EventType.ID" json:"event_type_id"`
	// 关系状态
	Status int `gorm:"column:status;type:int;NOT NULL;default:0" ddd:"exact;batch;upsert;simple;detail;status" json:"status"`
	// 订阅描述
	Description string `gorm:"column:description;type:varchar(4096);default null" json:"description"`
	Trailer
}

// Publication Table: t_publication; Group: event; 发布关系定义
type Publication struct {
	// use guid 36 char
	ID string `gorm:"column:id;type:varchar(36);primaryKey;NOT NULL" ddd:"exact;batch;upsert;simple"`
	// publisher app ref
	PublisherID string `gorm:"column:main_app_id;type:varchar(36);NOT NULL" ddd:"exact;batch;upsert;yam_s:Application.ID;role:publisher" json:"publisher_id"`
	// event type ref
	EventTypeID string `gorm:"column:name;type:varchar(255);NOT NULL" ddd:"exact;batch;fuzzy;upsert;simple;detail;yam_s:EventType.ID" json:"event_type_id"`
	// 关系状态
	Status int `gorm:"column:status;type:int;NOT NULL;default:0" ddd:"exact;batch;upsert;simple;detail;status" json:"status"`
	// 发布描述
	Description string `gorm:"column:description;type:varchar(4096);default null" json:"description"`
	Trailer
}

// EventBus Table: t_event_bus; Group: event; 事件集定义
type EventBus struct {
	// use guid 36 char
	ID string `gorm:"column:id;type:varchar(36);primaryKey;NOT NULL" ddd:"exact;batch;upsert;simple;may_s:EventType.EventBusID" json:"id"`
	// name unique
	Name string `gorm:"column:name;type:varchar(255);NOT NULL;uniqueIndex:t_event_bus_u1" ddd:"exact;batch;fuzzy;upsert;simple;detail" json:"name"`
	// 事件集参数
	Params string `gorm:"column:params;type:varchar(255)" ddd:"upsert;simple;detail" json:"params"`
	// 事件集描述
	Description string `gorm:"column:description;type:varchar(4096);default null" json:"description"`
	Trailer
}
