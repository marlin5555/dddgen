// Package zero Code generated, DO NOT EDIT.
package zero

// zero struct for Application entity start

// ApplicationDescription struct has GetApplicationDescription
type ApplicationDescription struct {
	Value string
}

// GetApplicationDescription return ApplicationDescription.Description impl base.ApplicationDescription
func (e ApplicationDescription) GetApplicationDescription() string {
	return e.Value
}

// FuzzyApplicationName struct has GetFuzzyApplicationName
type FuzzyApplicationName struct {
	Value string
}

// GetFuzzyApplicationName return FuzzyApplicationName.FuzzyApplicationName impl base.FuzzyApplicationName
func (e FuzzyApplicationName) GetFuzzyApplicationName() string {
	return e.Value
}

// FuzzyApplicationNickname struct has GetFuzzyApplicationNickname
type FuzzyApplicationNickname struct {
	Value string
}

// GetFuzzyApplicationNickname return FuzzyApplicationNickname.FuzzyApplicationNickname impl base.FuzzyApplicationNickname
func (e FuzzyApplicationNickname) GetFuzzyApplicationNickname() string {
	return e.Value
}

// ApplicationID struct has GetApplicationID
type ApplicationID struct {
	Value string
}

// GetApplicationID return ApplicationID.ID impl base.ApplicationID
func (e ApplicationID) GetApplicationID() string {
	return e.Value
}

// ApplicationIDs struct has GetApplicationIDs
type ApplicationIDs struct {
	Value []string
}

// GetApplicationIDs return ApplicationIDs.IDs impl base.ApplicationIDs
func (e ApplicationIDs) GetApplicationIDs() []string {
	return e.Value
}

// ApplicationName struct has GetApplicationName
type ApplicationName struct {
	Value string
}

// GetApplicationName return ApplicationName.Name impl base.ApplicationName
func (e ApplicationName) GetApplicationName() string {
	return e.Value
}

// ApplicationNames struct has GetApplicationNames
type ApplicationNames struct {
	Value []string
}

// GetApplicationNames return ApplicationNames.Names impl base.ApplicationNames
func (e ApplicationNames) GetApplicationNames() []string {
	return e.Value
}

// ApplicationNickname struct has GetApplicationNickname
type ApplicationNickname struct {
	Value string
}

// GetApplicationNickname return ApplicationNickname.Nickname impl base.ApplicationNickname
func (e ApplicationNickname) GetApplicationNickname() string {
	return e.Value
}
