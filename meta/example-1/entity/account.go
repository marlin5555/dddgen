package entity

// Account Table: t_account; Group: account; 账户定义
type Account struct {
	// use guid 36 char
	ID string `gorm:"column:id;type:varchar(36);primaryKey;NOT NULL" ddd:"exact;batch;id:auto;must_a:Secret.AccountID;may_a:Passport.AccountID;may_s:TechRelation.StudentID,TechRelation.TeacherID;simple"`
	// name unique
	Name string `gorm:"column:name;type:varchar(255);NOT NULL;uniqueIndex:t_account_u1" ddd:"exact;batch;fuzzy;upsert;simple;detail" json:"name"`
	// nickname
	Nickname string `gorm:"column:nickname;type:varchar(255);" ddd:"exact;fuzzy;upsert;simple;detail" json:"nickname"`
	// 账户描述
	Description string `gorm:"column:description;type:varchar(4096);default null" json:"description"`
	Trailer
}

// Secret Table: t_secret; Group: account; 密钥定义
type Secret struct {
	// use guid 36 char
	ID string `gorm:"column:id;type:varchar(36);primaryKey;NOT NULL" ddd:"exact;batch;id:auto;simple" json:"id"`
	// account id
	AccountID string `gorm:"column:account_id;type:varchar(255);" ddd:"exact;tsum_a:Account.ID" json:"account_id"`
	// secret
	Secret string `gorm:"column:secret;type:varchar(255);NOT NULL" json:"secret"`
	Trailer
}

// Passport Table: t_passport; Group: account; 护照定义
type Passport struct {
	// use guid 36 char
	ID string `gorm:"column:id;type:varchar(36);primaryKey;NOT NULL" ddd:"exact;batch;id:auto;simple" json:"id"`
	// account id
	AccountID string `gorm:"column:account_id;type:varchar(255);" ddd:"exact;batch;yam_a:Account.ID" json:"account_id"`
	// nationality 国籍
	Nationality string `gorm:"column:nationality;type:varchar(255)" ddd:"exact;batch" json:"nationality"`
	// postal address 通讯地址
	PostalAddress string `gorm:"column:postal_address;type:varchar(1024)" ddd:"fuzzy;" json:"postal_address"`
	Trailer
}

// TechRelation Table:t_tech_relation; Group: account;
type TechRelation struct {
	ID        string `gorm:"column:id;type:varchar(36)" ddd:"exact;batch;id:auto;simple" `
	StudentID string `gorm:"column:student_id;type:varchar(36)" ddd:"yam_s:Account.ID;role:student"`
	TeacherID string `gorm:"column:teacher_id;type:varchar(36)" ddd:"yam_s:Account.ID;role:teacher"`
	Trailer
}
