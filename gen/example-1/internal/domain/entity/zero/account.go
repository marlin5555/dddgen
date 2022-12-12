// Package zero Code generated, DO NOT EDIT.
package zero

// zero struct for Account entity start

// AccountDescription struct has GetAccountDescription
type AccountDescription struct {
	Value string
}

// GetAccountDescription return AccountDescription.Description impl base.AccountDescription
func (e AccountDescription) GetAccountDescription() string {
	return e.Value
}

// FuzzyAccountName struct has GetFuzzyAccountName
type FuzzyAccountName struct {
	Value string
}

// GetFuzzyAccountName return FuzzyAccountName.FuzzyAccountName impl base.FuzzyAccountName
func (e FuzzyAccountName) GetFuzzyAccountName() string {
	return e.Value
}

// FuzzyAccountNickname struct has GetFuzzyAccountNickname
type FuzzyAccountNickname struct {
	Value string
}

// GetFuzzyAccountNickname return FuzzyAccountNickname.FuzzyAccountNickname impl base.FuzzyAccountNickname
func (e FuzzyAccountNickname) GetFuzzyAccountNickname() string {
	return e.Value
}

// AccountID struct has GetAccountID
type AccountID struct {
	Value string
}

// GetAccountID return AccountID.ID impl base.AccountID
func (e AccountID) GetAccountID() string {
	return e.Value
}

// AccountIDs struct has GetAccountIDs
type AccountIDs struct {
	Value []string
}

// GetAccountIDs return AccountIDs.IDs impl base.AccountIDs
func (e AccountIDs) GetAccountIDs() []string {
	return e.Value
}

// AccountName struct has GetAccountName
type AccountName struct {
	Value string
}

// GetAccountName return AccountName.Name impl base.AccountName
func (e AccountName) GetAccountName() string {
	return e.Value
}

// AccountNames struct has GetAccountNames
type AccountNames struct {
	Value []string
}

// GetAccountNames return AccountNames.Names impl base.AccountNames
func (e AccountNames) GetAccountNames() []string {
	return e.Value
}

// AccountNickname struct has GetAccountNickname
type AccountNickname struct {
	Value string
}

// GetAccountNickname return AccountNickname.Nickname impl base.AccountNickname
func (e AccountNickname) GetAccountNickname() string {
	return e.Value
}

// zero struct for Passport entity start

// FuzzyPassportPostalAddress struct has GetFuzzyPassportPostalAddress
type FuzzyPassportPostalAddress struct {
	Value string
}

// GetFuzzyPassportPostalAddress return FuzzyPassportPostalAddress.FuzzyPassportPostalAddress impl base.FuzzyPassportPostalAddress
func (e FuzzyPassportPostalAddress) GetFuzzyPassportPostalAddress() string {
	return e.Value
}

// PassportID struct has GetPassportID
type PassportID struct {
	Value string
}

// GetPassportID return PassportID.ID impl base.PassportID
func (e PassportID) GetPassportID() string {
	return e.Value
}

// PassportIDs struct has GetPassportIDs
type PassportIDs struct {
	Value []string
}

// GetPassportIDs return PassportIDs.IDs impl base.PassportIDs
func (e PassportIDs) GetPassportIDs() []string {
	return e.Value
}

// PassportNationality struct has GetPassportNationality
type PassportNationality struct {
	Value string
}

// GetPassportNationality return PassportNationality.Nationality impl base.PassportNationality
func (e PassportNationality) GetPassportNationality() string {
	return e.Value
}

// PassportNationalitys struct has GetPassportNationalitys
type PassportNationalitys struct {
	Value []string
}

// GetPassportNationalitys return PassportNationalitys.Nationalitys impl base.PassportNationalitys
func (e PassportNationalitys) GetPassportNationalitys() []string {
	return e.Value
}

// PassportPostalAddress struct has GetPassportPostalAddress
type PassportPostalAddress struct {
	Value string
}

// GetPassportPostalAddress return PassportPostalAddress.PostalAddress impl base.PassportPostalAddress
func (e PassportPostalAddress) GetPassportPostalAddress() string {
	return e.Value
}

// zero struct for Secret entity start

// SecretID struct has GetSecretID
type SecretID struct {
	Value string
}

// GetSecretID return SecretID.ID impl base.SecretID
func (e SecretID) GetSecretID() string {
	return e.Value
}

// SecretIDs struct has GetSecretIDs
type SecretIDs struct {
	Value []string
}

// GetSecretIDs return SecretIDs.IDs impl base.SecretIDs
func (e SecretIDs) GetSecretIDs() []string {
	return e.Value
}

// SecretSecret struct has GetSecretSecret
type SecretSecret struct {
	Value string
}

// GetSecretSecret return SecretSecret.Secret impl base.SecretSecret
func (e SecretSecret) GetSecretSecret() string {
	return e.Value
}

// zero struct for TechRelation entity start

// TechRelationID struct has GetTechRelationID
type TechRelationID struct {
	Value string
}

// GetTechRelationID return TechRelationID.ID impl base.TechRelationID
func (e TechRelationID) GetTechRelationID() string {
	return e.Value
}

// TechRelationIDs struct has GetTechRelationIDs
type TechRelationIDs struct {
	Value []string
}

// GetTechRelationIDs return TechRelationIDs.IDs impl base.TechRelationIDs
func (e TechRelationIDs) GetTechRelationIDs() []string {
	return e.Value
}

// StudentID struct has GetStudentID
type StudentID struct {
	Value string
}

func (e StudentID) GetStudentID() string {
	return e.Value
}

// StudentIDs struct has GetStudentIDs
type StudentIDs struct {
	Value []string
}

func (e StudentIDs) GetStudentIDs() []string {
	return e.Value
}

// TeacherID struct has GetTeacherID
type TeacherID struct {
	Value string
}

func (e TeacherID) GetTeacherID() string {
	return e.Value
}

// TeacherIDs struct has GetTeacherIDs
type TeacherIDs struct {
	Value []string
}

func (e TeacherIDs) GetTeacherIDs() []string {
	return e.Value
}
