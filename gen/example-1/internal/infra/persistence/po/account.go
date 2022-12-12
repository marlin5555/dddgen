// Package po Code generated, DO NOT EDIT.
package po

import (
	"time"

	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/base"
)

// Account table name: t_account
type Account struct {
	CreatedAt   time.Time `gorm:"column:created_at;type:time;<-:create;autoCreateTime"`
	Creator     string    `gorm:"column:creator;type:varchar(255)"`
	DeletedAt   int64     `gorm:"column:deleted_at;default:0"`
	Description string    `gorm:"column:description;type:varchar(4096);default null"`
	ID          string    `gorm:"column:id;type:varchar(36);primaryKey;NOT NULL"`
	Name        string    `gorm:"column:name;type:varchar(255);NOT NULL;uniqueIndex:t_account_u1"`
	Nickname    string    `gorm:"column:nickname;type:varchar(255);"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:time;autoUpdateTime"`
	Updater     string    `gorm:"column:updater;type:varchar(255)"`
}

type Accounts []Account

func (p Accounts) GetAccountIDs() []string {
	return base.ToAccountIDs(p)
}

// TableName Account impl schema.Tabler
func (p Account) TableName() string {
	return "t_account"
}

// GetCreatedAt Account impl base.CreatedAt
func (p Account) GetCreatedAt() time.Time {
	return p.CreatedAt
}

// GetCreator Account impl base.Creator
func (p Account) GetCreator() string {
	return p.Creator
}

// GetDeletedAt Account impl base.DeletedAt
func (p Account) GetDeletedAt() int64 {
	return p.DeletedAt
}

// GetAccountDescription Account impl base.Description
func (p Account) GetAccountDescription() string {
	return p.Description
}

// GetAccountID Account impl base.ID
func (p Account) GetAccountID() string {
	return p.ID
}

// GetAccountIDs Account impl base.IDs
func (p Account) GetAccountIDs() []string {
	return []string{p.ID}
}

// GetAccountName Account impl base.Name
func (p Account) GetAccountName() string {
	return p.Name
}

// GetAccountNames Account impl base.Names
func (p Account) GetAccountNames() []string {
	return []string{p.Name}
}

// GetAccountNickname Account impl base.Nickname
func (p Account) GetAccountNickname() string {
	return p.Nickname
}

// GetUpdatedAt Account impl base.UpdatedAt
func (p Account) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

// GetUpdater Account impl base.Updater
func (p Account) GetUpdater() string {
	return p.Updater
}

// Passport table name: t_passport
type Passport struct {
	AccountID     string    `gorm:"column:account_id;type:varchar(255);"`
	CreatedAt     time.Time `gorm:"column:created_at;type:time;<-:create;autoCreateTime"`
	Creator       string    `gorm:"column:creator;type:varchar(255)"`
	DeletedAt     int64     `gorm:"column:deleted_at;default:0"`
	ID            string    `gorm:"column:id;type:varchar(36);primaryKey;NOT NULL"`
	Nationality   string    `gorm:"column:nationality;type:varchar(255)"`
	PostalAddress string    `gorm:"column:postal_address;type:varchar(1024)"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:time;autoUpdateTime"`
	Updater       string    `gorm:"column:updater;type:varchar(255)"`
}

type Passports []Passport

func (p Passports) GetPassportIDs() []string {
	return base.ToPassportIDs(p)
}
func (p Passports) GetAccountIDs() []string {
	return base.ToAccountIDs(p)
}

// TableName Passport impl schema.Tabler
func (p Passport) TableName() string {
	return "t_passport"
}

// GetAccountID Passport impl base.AccountID
func (p Passport) GetAccountID() string {
	return p.AccountID
}

// GetAccountIDs Passport impl base.AccountIDs
func (p Passport) GetAccountIDs() []string {
	return []string{p.AccountID}
}

// GetCreatedAt Passport impl base.CreatedAt
func (p Passport) GetCreatedAt() time.Time {
	return p.CreatedAt
}

// GetCreator Passport impl base.Creator
func (p Passport) GetCreator() string {
	return p.Creator
}

// GetDeletedAt Passport impl base.DeletedAt
func (p Passport) GetDeletedAt() int64 {
	return p.DeletedAt
}

// GetPassportID Passport impl base.ID
func (p Passport) GetPassportID() string {
	return p.ID
}

// GetPassportIDs Passport impl base.IDs
func (p Passport) GetPassportIDs() []string {
	return []string{p.ID}
}

// GetPassportNationality Passport impl base.Nationality
func (p Passport) GetPassportNationality() string {
	return p.Nationality
}

// GetPassportNationalitys Passport impl base.Nationalitys
func (p Passport) GetPassportNationalitys() []string {
	return []string{p.Nationality}
}

// GetPassportPostalAddress Passport impl base.PostalAddress
func (p Passport) GetPassportPostalAddress() string {
	return p.PostalAddress
}

// GetUpdatedAt Passport impl base.UpdatedAt
func (p Passport) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

// GetUpdater Passport impl base.Updater
func (p Passport) GetUpdater() string {
	return p.Updater
}

// Secret table name: t_secret
type Secret struct {
	AccountID string    `gorm:"column:account_id;type:varchar(255);"`
	CreatedAt time.Time `gorm:"column:created_at;type:time;<-:create;autoCreateTime"`
	Creator   string    `gorm:"column:creator;type:varchar(255)"`
	DeletedAt int64     `gorm:"column:deleted_at;default:0"`
	ID        string    `gorm:"column:id;type:varchar(36);primaryKey;NOT NULL"`
	Secret    string    `gorm:"column:secret;type:varchar(255);NOT NULL"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:time;autoUpdateTime"`
	Updater   string    `gorm:"column:updater;type:varchar(255)"`
}

type Secrets []Secret

func (p Secrets) GetSecretIDs() []string {
	return base.ToSecretIDs(p)
}
func (p Secrets) GetAccountIDs() []string {
	return base.ToAccountIDs(p)
}

// TableName Secret impl schema.Tabler
func (p Secret) TableName() string {
	return "t_secret"
}

// GetAccountID Secret impl base.AccountID
func (p Secret) GetAccountID() string {
	return p.AccountID
}

// GetCreatedAt Secret impl base.CreatedAt
func (p Secret) GetCreatedAt() time.Time {
	return p.CreatedAt
}

// GetCreator Secret impl base.Creator
func (p Secret) GetCreator() string {
	return p.Creator
}

// GetDeletedAt Secret impl base.DeletedAt
func (p Secret) GetDeletedAt() int64 {
	return p.DeletedAt
}

// GetSecretID Secret impl base.ID
func (p Secret) GetSecretID() string {
	return p.ID
}

// GetSecretIDs Secret impl base.IDs
func (p Secret) GetSecretIDs() []string {
	return []string{p.ID}
}

// GetSecretSecret Secret impl base.Secret
func (p Secret) GetSecretSecret() string {
	return p.Secret
}

// GetUpdatedAt Secret impl base.UpdatedAt
func (p Secret) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

// GetUpdater Secret impl base.Updater
func (p Secret) GetUpdater() string {
	return p.Updater
}

// TechRelation table name: t_tech_relation
type TechRelation struct {
	CreatedAt time.Time `gorm:"column:created_at;type:time;<-:create;autoCreateTime"`
	Creator   string    `gorm:"column:creator;type:varchar(255)"`
	DeletedAt int64     `gorm:"column:deleted_at;default:0"`
	ID        string    `gorm:"column:id;type:varchar(36)"`
	StudentID string    `gorm:"column:student_id;type:varchar(36)"`
	TeacherID string    `gorm:"column:teacher_id;type:varchar(36)"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:time;autoUpdateTime"`
	Updater   string    `gorm:"column:updater;type:varchar(255)"`
}

type TechRelations []TechRelation

func (p TechRelations) GetTechRelationIDs() []string {
	return base.ToTechRelationIDs(p)
}
func (p TechRelations) GetStudentIDs() []string {
	return base.ToStudentIDs(p)
}
func (p TechRelations) GetTeacherIDs() []string {
	return base.ToTeacherIDs(p)
}

// TableName TechRelation impl schema.Tabler
func (p TechRelation) TableName() string {
	return "t_tech_relation"
}

// GetCreatedAt TechRelation impl base.CreatedAt
func (p TechRelation) GetCreatedAt() time.Time {
	return p.CreatedAt
}

// GetCreator TechRelation impl base.Creator
func (p TechRelation) GetCreator() string {
	return p.Creator
}

// GetDeletedAt TechRelation impl base.DeletedAt
func (p TechRelation) GetDeletedAt() int64 {
	return p.DeletedAt
}

// GetTechRelationID TechRelation impl base.ID
func (p TechRelation) GetTechRelationID() string {
	return p.ID
}

// GetTechRelationIDs TechRelation impl base.IDs
func (p TechRelation) GetTechRelationIDs() []string {
	return []string{p.ID}
}

// GetStudentID TechRelation impl base.StudentID
func (p TechRelation) GetStudentID() string {
	return p.StudentID
}

// GetTeacherID TechRelation impl base.TeacherID
func (p TechRelation) GetTeacherID() string {
	return p.TeacherID
}

// GetUpdatedAt TechRelation impl base.UpdatedAt
func (p TechRelation) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

// GetUpdater TechRelation impl base.Updater
func (p TechRelation) GetUpdater() string {
	return p.Updater
}
