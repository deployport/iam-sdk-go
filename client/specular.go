package iam

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	clientruntime "go.deployport.com/specular-runtime/client"

	godeployportcomapiservicescorelib "go.deployport.com/api-services-corelib"

	godeployportcomapiservicescorelibconfigurator "go.deployport.com/api-services-corelib/configurator"
)

// NewUserInformationSSOProvider creates a new UserInformationSSOProvider
func NewUserInformationSSOProvider() *UserInformationSSOProvider {
	s := &UserInformationSSOProvider{}
	s.InitializeDefaults()
	return s
}

// UserInformationSSOProvider struct
type UserInformationSSOProvider struct {
	// provider name
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetName returns the value for the field name
func (e *UserInformationSSOProvider) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *UserInformationSSOProvider) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *UserInformationSSOProvider) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserInformationSSOProvider.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserInformationSSOProvider) InitializeDefaults() {
}

// userInformationSSOProviderAlias is defined to help pre and post JSON marshaling without recursive loops
type userInformationSSOProviderAlias UserInformationSSOProvider

// UnmarshalJSON implements json.Unmarshaler
func (e *UserInformationSSOProvider) UnmarshalJSON(data []byte) error {
	var alias userInformationSSOProviderAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserInformationSSOProvider)(&alias)).InitializeDefaults()
	*e = UserInformationSSOProvider(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserInformationSSOProvider) MarshalJSON() ([]byte, error) {
	alias := userInformationSSOProviderAlias(e)
	return json.Marshal(alias)
}

// NewUserInformationSSOProfile creates a new UserInformationSSOProfile
func NewUserInformationSSOProfile() *UserInformationSSOProfile {
	s := &UserInformationSSOProfile{}
	s.InitializeDefaults()
	return s
}

// UserInformationSSOProfile struct
type UserInformationSSOProfile struct {
	Email      string `json:"email,omitempty" yaml:"email,omitempty"`
	Fullname   string `json:"fullname,omitempty" yaml:"fullname,omitempty"`
	PictureURL string `json:"pictureURL,omitempty" yaml:"pictureURL,omitempty"`
}

// GetEmail returns the value for the field email
func (e *UserInformationSSOProfile) GetEmail() string {
	return e.Email
}

// SetEmail sets the value for the field email
func (e *UserInformationSSOProfile) SetEmail(email string) {
	e.Email = email
}

// GetFullname returns the value for the field fullname
func (e *UserInformationSSOProfile) GetFullname() string {
	return e.Fullname
}

// SetFullname sets the value for the field fullname
func (e *UserInformationSSOProfile) SetFullname(fullname string) {
	e.Fullname = fullname
}

// GetPictureURL returns the value for the field pictureURL
func (e *UserInformationSSOProfile) GetPictureURL() string {
	return e.PictureURL
}

// SetPictureURL sets the value for the field pictureURL
func (e *UserInformationSSOProfile) SetPictureURL(pictureURL string) {
	e.PictureURL = pictureURL
}

// StructPath returns StructPath
func (e *UserInformationSSOProfile) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserInformationSSOProfile.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserInformationSSOProfile) InitializeDefaults() {
}

// userInformationSSOProfileAlias is defined to help pre and post JSON marshaling without recursive loops
type userInformationSSOProfileAlias UserInformationSSOProfile

// UnmarshalJSON implements json.Unmarshaler
func (e *UserInformationSSOProfile) UnmarshalJSON(data []byte) error {
	var alias userInformationSSOProfileAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserInformationSSOProfile)(&alias)).InitializeDefaults()
	*e = UserInformationSSOProfile(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserInformationSSOProfile) MarshalJSON() ([]byte, error) {
	alias := userInformationSSOProfileAlias(e)
	return json.Marshal(alias)
}

// NewUserInformationSSO creates a new UserInformationSSO
func NewUserInformationSSO() *UserInformationSSO {
	s := &UserInformationSSO{}
	s.InitializeDefaults()
	return s
}

// UserInformationSSO struct
type UserInformationSSO struct {
	Profile  *UserInformationSSOProfile  `json:"profile,omitempty" yaml:"profile,omitempty"`
	Provider *UserInformationSSOProvider `json:"provider,omitempty" yaml:"provider,omitempty"`
}

// GetProfile returns the value for the field profile
func (e *UserInformationSSO) GetProfile() *UserInformationSSOProfile {
	return e.Profile
}

// SetProfile sets the value for the field profile
func (e *UserInformationSSO) SetProfile(profile *UserInformationSSOProfile) {
	e.Profile = profile
}

// GetProvider returns the value for the field provider
func (e *UserInformationSSO) GetProvider() *UserInformationSSOProvider {
	return e.Provider
}

// SetProvider sets the value for the field provider
func (e *UserInformationSSO) SetProvider(provider *UserInformationSSOProvider) {
	e.Provider = provider
}

// StructPath returns StructPath
func (e *UserInformationSSO) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserInformationSSO.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserInformationSSO) InitializeDefaults() {
}

// userInformationSSOAlias is defined to help pre and post JSON marshaling without recursive loops
type userInformationSSOAlias UserInformationSSO

// UnmarshalJSON implements json.Unmarshaler
func (e *UserInformationSSO) UnmarshalJSON(data []byte) error {
	var alias userInformationSSOAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserInformationSSO)(&alias)).InitializeDefaults()
	*e = UserInformationSSO(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserInformationSSO) MarshalJSON() ([]byte, error) {
	alias := userInformationSSOAlias(e)
	return json.Marshal(alias)
}

// NewUserInformation creates a new UserInformation
func NewUserInformation() *UserInformation {
	s := &UserInformation{}
	s.InitializeDefaults()
	return s
}

// UserInformation struct
type UserInformation struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// when the user comes from SSO, this field is populated with the extra information
	Sso      *UserInformationSSO `json:"sso,omitempty" yaml:"sso,omitempty"`
	Username string              `json:"username,omitempty" yaml:"username,omitempty"`
}

// GetDescription returns the value for the field description
func (e *UserInformation) GetDescription() string {
	return e.Description
}

// SetDescription sets the value for the field description
func (e *UserInformation) SetDescription(description string) {
	e.Description = description
}

// GetSso returns the value for the field sso
func (e *UserInformation) GetSso() *UserInformationSSO {
	return e.Sso
}

// SetSso sets the value for the field sso
func (e *UserInformation) SetSso(sso *UserInformationSSO) {
	e.Sso = sso
}

// GetUsername returns the value for the field username
func (e *UserInformation) GetUsername() string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *UserInformation) SetUsername(username string) {
	e.Username = username
}

// StructPath returns StructPath
func (e *UserInformation) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserInformation.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserInformation) InitializeDefaults() {
}

// userInformationAlias is defined to help pre and post JSON marshaling without recursive loops
type userInformationAlias UserInformation

// UnmarshalJSON implements json.Unmarshaler
func (e *UserInformation) UnmarshalJSON(data []byte) error {
	var alias userInformationAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserInformation)(&alias)).InitializeDefaults()
	*e = UserInformation(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserInformation) MarshalJSON() ([]byte, error) {
	alias := userInformationAlias(e)
	return json.Marshal(alias)
}

// NewRoleInformation creates a new RoleInformation
func NewRoleInformation() *RoleInformation {
	s := &RoleInformation{}
	s.InitializeDefaults()
	return s
}

// RoleInformation struct
type RoleInformation struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetDescription returns the value for the field description
func (e *RoleInformation) GetDescription() string {
	return e.Description
}

// SetDescription sets the value for the field description
func (e *RoleInformation) SetDescription(description string) {
	e.Description = description
}

// GetName returns the value for the field name
func (e *RoleInformation) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *RoleInformation) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *RoleInformation) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleInformation.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleInformation) InitializeDefaults() {
}

// roleInformationAlias is defined to help pre and post JSON marshaling without recursive loops
type roleInformationAlias RoleInformation

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleInformation) UnmarshalJSON(data []byte) error {
	var alias roleInformationAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleInformation)(&alias)).InitializeDefaults()
	*e = RoleInformation(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleInformation) MarshalJSON() ([]byte, error) {
	alias := roleInformationAlias(e)
	return json.Marshal(alias)
}

// NewCredentials creates a new Credentials
func NewCredentials() *Credentials {
	s := &Credentials{}
	s.InitializeDefaults()
	return s
}

// Credentials struct
type Credentials struct {
	AccessKeyID     string `json:"accessKeyID,omitempty" yaml:"accessKeyID,omitempty"`
	SecretAccessKey string `json:"secretAccessKey,omitempty" yaml:"secretAccessKey,omitempty"`
}

// GetAccessKeyID returns the value for the field accessKeyID
func (e *Credentials) GetAccessKeyID() string {
	return e.AccessKeyID
}

// SetAccessKeyID sets the value for the field accessKeyID
func (e *Credentials) SetAccessKeyID(accessKeyID string) {
	e.AccessKeyID = accessKeyID
}

// GetSecretAccessKey returns the value for the field secretAccessKey
func (e *Credentials) GetSecretAccessKey() string {
	return e.SecretAccessKey
}

// SetSecretAccessKey sets the value for the field secretAccessKey
func (e *Credentials) SetSecretAccessKey(secretAccessKey string) {
	e.SecretAccessKey = secretAccessKey
}

// StructPath returns StructPath
func (e *Credentials) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathCredentials.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *Credentials) InitializeDefaults() {
}

// credentialsAlias is defined to help pre and post JSON marshaling without recursive loops
type credentialsAlias Credentials

// UnmarshalJSON implements json.Unmarshaler
func (e *Credentials) UnmarshalJSON(data []byte) error {
	var alias credentialsAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*Credentials)(&alias)).InitializeDefaults()
	*e = Credentials(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e Credentials) MarshalJSON() ([]byte, error) {
	alias := credentialsAlias(e)
	return json.Marshal(alias)
}

// NewSSOProviderUnavailableProblem creates a new SSOProviderUnavailableProblem
func NewSSOProviderUnavailableProblem() *SSOProviderUnavailableProblem {
	s := &SSOProviderUnavailableProblem{}
	s.InitializeDefaults()
	return s
}

// SSOProviderUnavailableProblem struct
type SSOProviderUnavailableProblem struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *SSOProviderUnavailableProblem) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [SSOProviderUnavailableProblem]
func (e *SSOProviderUnavailableProblem) Is(err error) bool {
	_, ok := err.(*SSOProviderUnavailableProblem)
	return ok
}

// IsSSOProviderUnavailableProblem indicates whether the given error chain contains an error of type [SSOProviderUnavailableProblem]
func IsSSOProviderUnavailableProblem(err error) bool {
	return errors.Is(err, &SSOProviderUnavailableProblem{})
}

// GetMessage returns the value for the field message
func (e *SSOProviderUnavailableProblem) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *SSOProviderUnavailableProblem) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *SSOProviderUnavailableProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathSSOProviderUnavailableProblem.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *SSOProviderUnavailableProblem) InitializeDefaults() {
}

// sSOProviderUnavailableProblemAlias is defined to help pre and post JSON marshaling without recursive loops
type sSOProviderUnavailableProblemAlias SSOProviderUnavailableProblem

// UnmarshalJSON implements json.Unmarshaler
func (e *SSOProviderUnavailableProblem) UnmarshalJSON(data []byte) error {
	var alias sSOProviderUnavailableProblemAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*SSOProviderUnavailableProblem)(&alias)).InitializeDefaults()
	*e = SSOProviderUnavailableProblem(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e SSOProviderUnavailableProblem) MarshalJSON() ([]byte, error) {
	alias := sSOProviderUnavailableProblemAlias(e)
	return json.Marshal(alias)
}

// NewSSOFlow creates a new SSOFlow
func NewSSOFlow() *SSOFlow {
	s := &SSOFlow{}
	s.InitializeDefaults()
	return s
}

// SSOFlow struct
type SSOFlow struct {
	BrowseURL                 string    `json:"browseURL,omitempty" yaml:"browseURL,omitempty"`
	CompletionIntervalSeconds int32     `json:"completionIntervalSeconds,omitempty" yaml:"completionIntervalSeconds,omitempty"`
	ExpiresAt                 time.Time `json:"expiresAt,omitempty" yaml:"expiresAt,omitempty"`
}

// GetBrowseURL returns the value for the field browseURL
func (e *SSOFlow) GetBrowseURL() string {
	return e.BrowseURL
}

// SetBrowseURL sets the value for the field browseURL
func (e *SSOFlow) SetBrowseURL(browseURL string) {
	e.BrowseURL = browseURL
}

// GetCompletionIntervalSeconds returns the value for the field completionIntervalSeconds
func (e *SSOFlow) GetCompletionIntervalSeconds() int32 {
	return e.CompletionIntervalSeconds
}

// SetCompletionIntervalSeconds sets the value for the field completionIntervalSeconds
func (e *SSOFlow) SetCompletionIntervalSeconds(completionIntervalSeconds int32) {
	e.CompletionIntervalSeconds = completionIntervalSeconds
}

// GetExpiresAt returns the value for the field expiresAt
func (e *SSOFlow) GetExpiresAt() time.Time {
	return e.ExpiresAt
}

// SetExpiresAt sets the value for the field expiresAt
func (e *SSOFlow) SetExpiresAt(expiresAt time.Time) {
	e.ExpiresAt = expiresAt
}

// StructPath returns StructPath
func (e *SSOFlow) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathSSOFlow.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *SSOFlow) InitializeDefaults() {
}

// sSOFlowAlias is defined to help pre and post JSON marshaling without recursive loops
type sSOFlowAlias SSOFlow

// UnmarshalJSON implements json.Unmarshaler
func (e *SSOFlow) UnmarshalJSON(data []byte) error {
	var alias sSOFlowAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*SSOFlow)(&alias)).InitializeDefaults()
	*e = SSOFlow(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e SSOFlow) MarshalJSON() ([]byte, error) {
	alias := sSOFlowAlias(e)
	return json.Marshal(alias)
}

// NewAccount creates a new Account
func NewAccount() *Account {
	s := &Account{}
	s.InitializeDefaults()
	return s
}

// Account struct
type Account struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetName returns the value for the field name
func (e *Account) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *Account) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *Account) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccount.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *Account) InitializeDefaults() {
}

// accountAlias is defined to help pre and post JSON marshaling without recursive loops
type accountAlias Account

// UnmarshalJSON implements json.Unmarshaler
func (e *Account) UnmarshalJSON(data []byte) error {
	var alias accountAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*Account)(&alias)).InitializeDefaults()
	*e = Account(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e Account) MarshalJSON() ([]byte, error) {
	alias := accountAlias(e)
	return json.Marshal(alias)
}

// NewAccountSSOProvider creates a new AccountSSOProvider
func NewAccountSSOProvider() *AccountSSOProvider {
	s := &AccountSSOProvider{}
	s.InitializeDefaults()
	return s
}

// AccountSSOProvider struct
type AccountSSOProvider struct {
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetDisplayName returns the value for the field displayName
func (e *AccountSSOProvider) GetDisplayName() string {
	return e.DisplayName
}

// SetDisplayName sets the value for the field displayName
func (e *AccountSSOProvider) SetDisplayName(displayName string) {
	e.DisplayName = displayName
}

// GetName returns the value for the field name
func (e *AccountSSOProvider) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *AccountSSOProvider) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *AccountSSOProvider) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOProvider.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountSSOProvider) InitializeDefaults() {
}

// accountSSOProviderAlias is defined to help pre and post JSON marshaling without recursive loops
type accountSSOProviderAlias AccountSSOProvider

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountSSOProvider) UnmarshalJSON(data []byte) error {
	var alias accountSSOProviderAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountSSOProvider)(&alias)).InitializeDefaults()
	*e = AccountSSOProvider(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountSSOProvider) MarshalJSON() ([]byte, error) {
	alias := accountSSOProviderAlias(e)
	return json.Marshal(alias)
}

// NewPolicyNotFoundProblem creates a new PolicyNotFoundProblem
func NewPolicyNotFoundProblem() *PolicyNotFoundProblem {
	s := &PolicyNotFoundProblem{}
	s.InitializeDefaults()
	return s
}

// PolicyNotFoundProblem struct
type PolicyNotFoundProblem struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *PolicyNotFoundProblem) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [PolicyNotFoundProblem]
func (e *PolicyNotFoundProblem) Is(err error) bool {
	_, ok := err.(*PolicyNotFoundProblem)
	return ok
}

// IsPolicyNotFoundProblem indicates whether the given error chain contains an error of type [PolicyNotFoundProblem]
func IsPolicyNotFoundProblem(err error) bool {
	return errors.Is(err, &PolicyNotFoundProblem{})
}

// GetMessage returns the value for the field message
func (e *PolicyNotFoundProblem) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *PolicyNotFoundProblem) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *PolicyNotFoundProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathPolicyNotFoundProblem.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *PolicyNotFoundProblem) InitializeDefaults() {
}

// policyNotFoundProblemAlias is defined to help pre and post JSON marshaling without recursive loops
type policyNotFoundProblemAlias PolicyNotFoundProblem

// UnmarshalJSON implements json.Unmarshaler
func (e *PolicyNotFoundProblem) UnmarshalJSON(data []byte) error {
	var alias policyNotFoundProblemAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*PolicyNotFoundProblem)(&alias)).InitializeDefaults()
	*e = PolicyNotFoundProblem(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e PolicyNotFoundProblem) MarshalJSON() ([]byte, error) {
	alias := policyNotFoundProblemAlias(e)
	return json.Marshal(alias)
}

// NewAccountSSOAutoJoinPolicy creates a new AccountSSOAutoJoinPolicy
func NewAccountSSOAutoJoinPolicy() *AccountSSOAutoJoinPolicy {
	s := &AccountSSOAutoJoinPolicy{}
	s.InitializeDefaults()
	return s
}

// AccountSSOAutoJoinPolicy struct
type AccountSSOAutoJoinPolicy struct {
	// email-domain users come from, example "deployport.com"
	Domain  string `json:"domain,omitempty" yaml:"domain,omitempty"`
	Enabled bool   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	// unique identifer of the auto-join policy
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// where the users are coming from
	OriginAccountName string `json:"originAccountName,omitempty" yaml:"originAccountName,omitempty"`
	// SSO provider in the source account
	OriginProviderName string `json:"originProviderName,omitempty" yaml:"originProviderName,omitempty"`
}

// GetDomain returns the value for the field domain
func (e *AccountSSOAutoJoinPolicy) GetDomain() string {
	return e.Domain
}

// SetDomain sets the value for the field domain
func (e *AccountSSOAutoJoinPolicy) SetDomain(domain string) {
	e.Domain = domain
}

// GetEnabled returns the value for the field enabled
func (e *AccountSSOAutoJoinPolicy) GetEnabled() bool {
	return e.Enabled
}

// SetEnabled sets the value for the field enabled
func (e *AccountSSOAutoJoinPolicy) SetEnabled(enabled bool) {
	e.Enabled = enabled
}

// GetName returns the value for the field name
func (e *AccountSSOAutoJoinPolicy) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *AccountSSOAutoJoinPolicy) SetName(name string) {
	e.Name = name
}

// GetOriginAccountName returns the value for the field originAccountName
func (e *AccountSSOAutoJoinPolicy) GetOriginAccountName() string {
	return e.OriginAccountName
}

// SetOriginAccountName sets the value for the field originAccountName
func (e *AccountSSOAutoJoinPolicy) SetOriginAccountName(originAccountName string) {
	e.OriginAccountName = originAccountName
}

// GetOriginProviderName returns the value for the field originProviderName
func (e *AccountSSOAutoJoinPolicy) GetOriginProviderName() string {
	return e.OriginProviderName
}

// SetOriginProviderName sets the value for the field originProviderName
func (e *AccountSSOAutoJoinPolicy) SetOriginProviderName(originProviderName string) {
	e.OriginProviderName = originProviderName
}

// StructPath returns StructPath
func (e *AccountSSOAutoJoinPolicy) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOAutoJoinPolicy.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountSSOAutoJoinPolicy) InitializeDefaults() {
}

// accountSSOAutoJoinPolicyAlias is defined to help pre and post JSON marshaling without recursive loops
type accountSSOAutoJoinPolicyAlias AccountSSOAutoJoinPolicy

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountSSOAutoJoinPolicy) UnmarshalJSON(data []byte) error {
	var alias accountSSOAutoJoinPolicyAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountSSOAutoJoinPolicy)(&alias)).InitializeDefaults()
	*e = AccountSSOAutoJoinPolicy(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountSSOAutoJoinPolicy) MarshalJSON() ([]byte, error) {
	alias := accountSSOAutoJoinPolicyAlias(e)
	return json.Marshal(alias)
}

// NewAccountCreateInput creates a new AccountCreateInput
func NewAccountCreateInput() *AccountCreateInput {
	s := &AccountCreateInput{}
	s.InitializeDefaults()
	return s
}

// AccountCreateInput struct
type AccountCreateInput struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetName returns the value for the field name
func (e *AccountCreateInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *AccountCreateInput) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *AccountCreateInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountCreateInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountCreateInput) InitializeDefaults() {
}

// accountCreateInputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountCreateInputAlias AccountCreateInput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountCreateInput) UnmarshalJSON(data []byte) error {
	var alias accountCreateInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountCreateInput)(&alias)).InitializeDefaults()
	*e = AccountCreateInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountCreateInput) MarshalJSON() ([]byte, error) {
	alias := accountCreateInputAlias(e)
	return json.Marshal(alias)
}

// NewAccountCreateOutput creates a new AccountCreateOutput
func NewAccountCreateOutput() *AccountCreateOutput {
	s := &AccountCreateOutput{}
	s.InitializeDefaults()
	return s
}

// AccountCreateOutput struct
type AccountCreateOutput struct {
	Account *Account `json:"account,omitempty" yaml:"account,omitempty"`
}

// GetAccount returns the value for the field account
func (e *AccountCreateOutput) GetAccount() *Account {
	return e.Account
}

// SetAccount sets the value for the field account
func (e *AccountCreateOutput) SetAccount(account *Account) {
	e.Account = account
}

// StructPath returns StructPath
func (e *AccountCreateOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountCreateOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountCreateOutput) InitializeDefaults() {
}

// accountCreateOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountCreateOutputAlias AccountCreateOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountCreateOutput) UnmarshalJSON(data []byte) error {
	var alias accountCreateOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountCreateOutput)(&alias)).InitializeDefaults()
	*e = AccountCreateOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountCreateOutput) MarshalJSON() ([]byte, error) {
	alias := accountCreateOutputAlias(e)
	return json.Marshal(alias)
}

// NewAccountCreateInvalidNameProblem creates a new AccountCreateInvalidNameProblem
func NewAccountCreateInvalidNameProblem() *AccountCreateInvalidNameProblem {
	s := &AccountCreateInvalidNameProblem{}
	s.InitializeDefaults()
	return s
}

// AccountCreateInvalidNameProblem struct
type AccountCreateInvalidNameProblem struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *AccountCreateInvalidNameProblem) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [AccountCreateInvalidNameProblem]
func (e *AccountCreateInvalidNameProblem) Is(err error) bool {
	_, ok := err.(*AccountCreateInvalidNameProblem)
	return ok
}

// IsAccountCreateInvalidNameProblem indicates whether the given error chain contains an error of type [AccountCreateInvalidNameProblem]
func IsAccountCreateInvalidNameProblem(err error) bool {
	return errors.Is(err, &AccountCreateInvalidNameProblem{})
}

// GetMessage returns the value for the field message
func (e *AccountCreateInvalidNameProblem) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *AccountCreateInvalidNameProblem) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *AccountCreateInvalidNameProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountCreateInvalidNameProblem.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountCreateInvalidNameProblem) InitializeDefaults() {
}

// accountCreateInvalidNameProblemAlias is defined to help pre and post JSON marshaling without recursive loops
type accountCreateInvalidNameProblemAlias AccountCreateInvalidNameProblem

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountCreateInvalidNameProblem) UnmarshalJSON(data []byte) error {
	var alias accountCreateInvalidNameProblemAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountCreateInvalidNameProblem)(&alias)).InitializeDefaults()
	*e = AccountCreateInvalidNameProblem(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountCreateInvalidNameProblem) MarshalJSON() ([]byte, error) {
	alias := accountCreateInvalidNameProblemAlias(e)
	return json.Marshal(alias)
}

// NewAccountAssumeIdentityInput creates a new AccountAssumeIdentityInput
func NewAccountAssumeIdentityInput() *AccountAssumeIdentityInput {
	s := &AccountAssumeIdentityInput{}
	s.InitializeDefaults()
	return s
}

// AccountAssumeIdentityInput struct
type AccountAssumeIdentityInput struct {
	AccountName string `json:"accountName,omitempty" yaml:"accountName,omitempty"`
}

// GetAccountName returns the value for the field accountName
func (e *AccountAssumeIdentityInput) GetAccountName() string {
	return e.AccountName
}

// SetAccountName sets the value for the field accountName
func (e *AccountAssumeIdentityInput) SetAccountName(accountName string) {
	e.AccountName = accountName
}

// StructPath returns StructPath
func (e *AccountAssumeIdentityInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountAssumeIdentityInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountAssumeIdentityInput) InitializeDefaults() {
}

// accountAssumeIdentityInputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountAssumeIdentityInputAlias AccountAssumeIdentityInput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountAssumeIdentityInput) UnmarshalJSON(data []byte) error {
	var alias accountAssumeIdentityInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountAssumeIdentityInput)(&alias)).InitializeDefaults()
	*e = AccountAssumeIdentityInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountAssumeIdentityInput) MarshalJSON() ([]byte, error) {
	alias := accountAssumeIdentityInputAlias(e)
	return json.Marshal(alias)
}

// NewAccountAssumeIdentityOutput creates a new AccountAssumeIdentityOutput
func NewAccountAssumeIdentityOutput() *AccountAssumeIdentityOutput {
	s := &AccountAssumeIdentityOutput{}
	s.InitializeDefaults()
	return s
}

// AccountAssumeIdentityOutput struct
type AccountAssumeIdentityOutput struct {
	Credentials *Credentials `json:"credentials,omitempty" yaml:"credentials,omitempty"`
}

// GetCredentials returns the value for the field credentials
func (e *AccountAssumeIdentityOutput) GetCredentials() *Credentials {
	return e.Credentials
}

// SetCredentials sets the value for the field credentials
func (e *AccountAssumeIdentityOutput) SetCredentials(credentials *Credentials) {
	e.Credentials = credentials
}

// StructPath returns StructPath
func (e *AccountAssumeIdentityOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountAssumeIdentityOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountAssumeIdentityOutput) InitializeDefaults() {
}

// accountAssumeIdentityOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountAssumeIdentityOutputAlias AccountAssumeIdentityOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountAssumeIdentityOutput) UnmarshalJSON(data []byte) error {
	var alias accountAssumeIdentityOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountAssumeIdentityOutput)(&alias)).InitializeDefaults()
	*e = AccountAssumeIdentityOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountAssumeIdentityOutput) MarshalJSON() ([]byte, error) {
	alias := accountAssumeIdentityOutputAlias(e)
	return json.Marshal(alias)
}

// NewAccountSSOBeginAuthenticationInput creates a new AccountSSOBeginAuthenticationInput
func NewAccountSSOBeginAuthenticationInput() *AccountSSOBeginAuthenticationInput {
	s := &AccountSSOBeginAuthenticationInput{}
	s.InitializeDefaults()
	return s
}

// AccountSSOBeginAuthenticationInput struct
type AccountSSOBeginAuthenticationInput struct {
	AccountName string `json:"accountName,omitempty" yaml:"accountName,omitempty"`
	// client generated code challenge that will be used to verify the completion code, SHA-256 of the codeVerifier
	CodeChallenge string `json:"codeChallenge,omitempty" yaml:"codeChallenge,omitempty"`
	ProviderName  string `json:"providerName,omitempty" yaml:"providerName,omitempty"`
}

// GetAccountName returns the value for the field accountName
func (e *AccountSSOBeginAuthenticationInput) GetAccountName() string {
	return e.AccountName
}

// SetAccountName sets the value for the field accountName
func (e *AccountSSOBeginAuthenticationInput) SetAccountName(accountName string) {
	e.AccountName = accountName
}

// GetCodeChallenge returns the value for the field codeChallenge
func (e *AccountSSOBeginAuthenticationInput) GetCodeChallenge() string {
	return e.CodeChallenge
}

// SetCodeChallenge sets the value for the field codeChallenge
func (e *AccountSSOBeginAuthenticationInput) SetCodeChallenge(codeChallenge string) {
	e.CodeChallenge = codeChallenge
}

// GetProviderName returns the value for the field providerName
func (e *AccountSSOBeginAuthenticationInput) GetProviderName() string {
	return e.ProviderName
}

// SetProviderName sets the value for the field providerName
func (e *AccountSSOBeginAuthenticationInput) SetProviderName(providerName string) {
	e.ProviderName = providerName
}

// StructPath returns StructPath
func (e *AccountSSOBeginAuthenticationInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOBeginAuthenticationInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountSSOBeginAuthenticationInput) InitializeDefaults() {
}

// accountSSOBeginAuthenticationInputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountSSOBeginAuthenticationInputAlias AccountSSOBeginAuthenticationInput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountSSOBeginAuthenticationInput) UnmarshalJSON(data []byte) error {
	var alias accountSSOBeginAuthenticationInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountSSOBeginAuthenticationInput)(&alias)).InitializeDefaults()
	*e = AccountSSOBeginAuthenticationInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountSSOBeginAuthenticationInput) MarshalJSON() ([]byte, error) {
	alias := accountSSOBeginAuthenticationInputAlias(e)
	return json.Marshal(alias)
}

// NewAccountSSOBeginAuthenticationOutput creates a new AccountSSOBeginAuthenticationOutput
func NewAccountSSOBeginAuthenticationOutput() *AccountSSOBeginAuthenticationOutput {
	s := &AccountSSOBeginAuthenticationOutput{}
	s.InitializeDefaults()
	return s
}

// AccountSSOBeginAuthenticationOutput struct
type AccountSSOBeginAuthenticationOutput struct {
	Flow *SSOFlow `json:"flow,omitempty" yaml:"flow,omitempty"`
}

// GetFlow returns the value for the field flow
func (e *AccountSSOBeginAuthenticationOutput) GetFlow() *SSOFlow {
	return e.Flow
}

// SetFlow sets the value for the field flow
func (e *AccountSSOBeginAuthenticationOutput) SetFlow(flow *SSOFlow) {
	e.Flow = flow
}

// StructPath returns StructPath
func (e *AccountSSOBeginAuthenticationOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOBeginAuthenticationOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountSSOBeginAuthenticationOutput) InitializeDefaults() {
}

// accountSSOBeginAuthenticationOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountSSOBeginAuthenticationOutputAlias AccountSSOBeginAuthenticationOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountSSOBeginAuthenticationOutput) UnmarshalJSON(data []byte) error {
	var alias accountSSOBeginAuthenticationOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountSSOBeginAuthenticationOutput)(&alias)).InitializeDefaults()
	*e = AccountSSOBeginAuthenticationOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountSSOBeginAuthenticationOutput) MarshalJSON() ([]byte, error) {
	alias := accountSSOBeginAuthenticationOutputAlias(e)
	return json.Marshal(alias)
}

// NewAccountSSOBeginAuthenticationParameterProblem creates a new AccountSSOBeginAuthenticationParameterProblem
func NewAccountSSOBeginAuthenticationParameterProblem() *AccountSSOBeginAuthenticationParameterProblem {
	s := &AccountSSOBeginAuthenticationParameterProblem{}
	s.InitializeDefaults()
	return s
}

// AccountSSOBeginAuthenticationParameterProblem struct
type AccountSSOBeginAuthenticationParameterProblem struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *AccountSSOBeginAuthenticationParameterProblem) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [AccountSSOBeginAuthenticationParameterProblem]
func (e *AccountSSOBeginAuthenticationParameterProblem) Is(err error) bool {
	_, ok := err.(*AccountSSOBeginAuthenticationParameterProblem)
	return ok
}

// IsAccountSSOBeginAuthenticationParameterProblem indicates whether the given error chain contains an error of type [AccountSSOBeginAuthenticationParameterProblem]
func IsAccountSSOBeginAuthenticationParameterProblem(err error) bool {
	return errors.Is(err, &AccountSSOBeginAuthenticationParameterProblem{})
}

// GetMessage returns the value for the field message
func (e *AccountSSOBeginAuthenticationParameterProblem) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *AccountSSOBeginAuthenticationParameterProblem) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *AccountSSOBeginAuthenticationParameterProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOBeginAuthenticationParameterProblem.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountSSOBeginAuthenticationParameterProblem) InitializeDefaults() {
}

// accountSSOBeginAuthenticationParameterProblemAlias is defined to help pre and post JSON marshaling without recursive loops
type accountSSOBeginAuthenticationParameterProblemAlias AccountSSOBeginAuthenticationParameterProblem

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountSSOBeginAuthenticationParameterProblem) UnmarshalJSON(data []byte) error {
	var alias accountSSOBeginAuthenticationParameterProblemAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountSSOBeginAuthenticationParameterProblem)(&alias)).InitializeDefaults()
	*e = AccountSSOBeginAuthenticationParameterProblem(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountSSOBeginAuthenticationParameterProblem) MarshalJSON() ([]byte, error) {
	alias := accountSSOBeginAuthenticationParameterProblemAlias(e)
	return json.Marshal(alias)
}

// NewAccountSSOCompleteAuthenticationInput creates a new AccountSSOCompleteAuthenticationInput
func NewAccountSSOCompleteAuthenticationInput() *AccountSSOCompleteAuthenticationInput {
	s := &AccountSSOCompleteAuthenticationInput{}
	s.InitializeDefaults()
	return s
}

// AccountSSOCompleteAuthenticationInput struct
type AccountSSOCompleteAuthenticationInput struct {
	AccountName string `json:"accountName,omitempty" yaml:"accountName,omitempty"`
	// base64-url-encoded client generated code challenge that will be used to verify the completion code
	CodeVerifierB64 string `json:"codeVerifierB64,omitempty" yaml:"codeVerifierB64,omitempty"`
	ProviderName    string `json:"providerName,omitempty" yaml:"providerName,omitempty"`
}

// GetAccountName returns the value for the field accountName
func (e *AccountSSOCompleteAuthenticationInput) GetAccountName() string {
	return e.AccountName
}

// SetAccountName sets the value for the field accountName
func (e *AccountSSOCompleteAuthenticationInput) SetAccountName(accountName string) {
	e.AccountName = accountName
}

// GetCodeVerifierB64 returns the value for the field codeVerifierB64
func (e *AccountSSOCompleteAuthenticationInput) GetCodeVerifierB64() string {
	return e.CodeVerifierB64
}

// SetCodeVerifierB64 sets the value for the field codeVerifierB64
func (e *AccountSSOCompleteAuthenticationInput) SetCodeVerifierB64(codeVerifierB64 string) {
	e.CodeVerifierB64 = codeVerifierB64
}

// GetProviderName returns the value for the field providerName
func (e *AccountSSOCompleteAuthenticationInput) GetProviderName() string {
	return e.ProviderName
}

// SetProviderName sets the value for the field providerName
func (e *AccountSSOCompleteAuthenticationInput) SetProviderName(providerName string) {
	e.ProviderName = providerName
}

// StructPath returns StructPath
func (e *AccountSSOCompleteAuthenticationInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOCompleteAuthenticationInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountSSOCompleteAuthenticationInput) InitializeDefaults() {
}

// accountSSOCompleteAuthenticationInputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountSSOCompleteAuthenticationInputAlias AccountSSOCompleteAuthenticationInput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountSSOCompleteAuthenticationInput) UnmarshalJSON(data []byte) error {
	var alias accountSSOCompleteAuthenticationInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountSSOCompleteAuthenticationInput)(&alias)).InitializeDefaults()
	*e = AccountSSOCompleteAuthenticationInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountSSOCompleteAuthenticationInput) MarshalJSON() ([]byte, error) {
	alias := accountSSOCompleteAuthenticationInputAlias(e)
	return json.Marshal(alias)
}

// NewAccountSSOCompleteAuthenticationOutput creates a new AccountSSOCompleteAuthenticationOutput
func NewAccountSSOCompleteAuthenticationOutput() *AccountSSOCompleteAuthenticationOutput {
	s := &AccountSSOCompleteAuthenticationOutput{}
	s.InitializeDefaults()
	return s
}

// AccountSSOCompleteAuthenticationOutput struct
type AccountSSOCompleteAuthenticationOutput struct {
	Credentials *Credentials `json:"credentials,omitempty" yaml:"credentials,omitempty"`
}

// GetCredentials returns the value for the field credentials
func (e *AccountSSOCompleteAuthenticationOutput) GetCredentials() *Credentials {
	return e.Credentials
}

// SetCredentials sets the value for the field credentials
func (e *AccountSSOCompleteAuthenticationOutput) SetCredentials(credentials *Credentials) {
	e.Credentials = credentials
}

// StructPath returns StructPath
func (e *AccountSSOCompleteAuthenticationOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOCompleteAuthenticationOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountSSOCompleteAuthenticationOutput) InitializeDefaults() {
}

// accountSSOCompleteAuthenticationOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountSSOCompleteAuthenticationOutputAlias AccountSSOCompleteAuthenticationOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountSSOCompleteAuthenticationOutput) UnmarshalJSON(data []byte) error {
	var alias accountSSOCompleteAuthenticationOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountSSOCompleteAuthenticationOutput)(&alias)).InitializeDefaults()
	*e = AccountSSOCompleteAuthenticationOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountSSOCompleteAuthenticationOutput) MarshalJSON() ([]byte, error) {
	alias := accountSSOCompleteAuthenticationOutputAlias(e)
	return json.Marshal(alias)
}

// NewAccountSSOCompleteAuthenticationInvalidFlowProblem creates a new AccountSSOCompleteAuthenticationInvalidFlowProblem
func NewAccountSSOCompleteAuthenticationInvalidFlowProblem() *AccountSSOCompleteAuthenticationInvalidFlowProblem {
	s := &AccountSSOCompleteAuthenticationInvalidFlowProblem{}
	s.InitializeDefaults()
	return s
}

// AccountSSOCompleteAuthenticationInvalidFlowProblem struct
type AccountSSOCompleteAuthenticationInvalidFlowProblem struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *AccountSSOCompleteAuthenticationInvalidFlowProblem) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [AccountSSOCompleteAuthenticationInvalidFlowProblem]
func (e *AccountSSOCompleteAuthenticationInvalidFlowProblem) Is(err error) bool {
	_, ok := err.(*AccountSSOCompleteAuthenticationInvalidFlowProblem)
	return ok
}

// IsAccountSSOCompleteAuthenticationInvalidFlowProblem indicates whether the given error chain contains an error of type [AccountSSOCompleteAuthenticationInvalidFlowProblem]
func IsAccountSSOCompleteAuthenticationInvalidFlowProblem(err error) bool {
	return errors.Is(err, &AccountSSOCompleteAuthenticationInvalidFlowProblem{})
}

// GetMessage returns the value for the field message
func (e *AccountSSOCompleteAuthenticationInvalidFlowProblem) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *AccountSSOCompleteAuthenticationInvalidFlowProblem) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *AccountSSOCompleteAuthenticationInvalidFlowProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOCompleteAuthenticationInvalidFlowProblem.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountSSOCompleteAuthenticationInvalidFlowProblem) InitializeDefaults() {
}

// accountSSOCompleteAuthenticationInvalidFlowProblemAlias is defined to help pre and post JSON marshaling without recursive loops
type accountSSOCompleteAuthenticationInvalidFlowProblemAlias AccountSSOCompleteAuthenticationInvalidFlowProblem

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountSSOCompleteAuthenticationInvalidFlowProblem) UnmarshalJSON(data []byte) error {
	var alias accountSSOCompleteAuthenticationInvalidFlowProblemAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountSSOCompleteAuthenticationInvalidFlowProblem)(&alias)).InitializeDefaults()
	*e = AccountSSOCompleteAuthenticationInvalidFlowProblem(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountSSOCompleteAuthenticationInvalidFlowProblem) MarshalJSON() ([]byte, error) {
	alias := accountSSOCompleteAuthenticationInvalidFlowProblemAlias(e)
	return json.Marshal(alias)
}

// NewAccountSSOGetProvidersInput creates a new AccountSSOGetProvidersInput
func NewAccountSSOGetProvidersInput() *AccountSSOGetProvidersInput {
	s := &AccountSSOGetProvidersInput{}
	s.InitializeDefaults()
	return s
}

// AccountSSOGetProvidersInput struct
type AccountSSOGetProvidersInput struct {
	AccountName string `json:"accountName,omitempty" yaml:"accountName,omitempty"`
}

// GetAccountName returns the value for the field accountName
func (e *AccountSSOGetProvidersInput) GetAccountName() string {
	return e.AccountName
}

// SetAccountName sets the value for the field accountName
func (e *AccountSSOGetProvidersInput) SetAccountName(accountName string) {
	e.AccountName = accountName
}

// StructPath returns StructPath
func (e *AccountSSOGetProvidersInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOGetProvidersInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountSSOGetProvidersInput) InitializeDefaults() {
}

// accountSSOGetProvidersInputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountSSOGetProvidersInputAlias AccountSSOGetProvidersInput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountSSOGetProvidersInput) UnmarshalJSON(data []byte) error {
	var alias accountSSOGetProvidersInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountSSOGetProvidersInput)(&alias)).InitializeDefaults()
	*e = AccountSSOGetProvidersInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountSSOGetProvidersInput) MarshalJSON() ([]byte, error) {
	alias := accountSSOGetProvidersInputAlias(e)
	return json.Marshal(alias)
}

// NewAccountSSOGetProvidersOutput creates a new AccountSSOGetProvidersOutput
func NewAccountSSOGetProvidersOutput() *AccountSSOGetProvidersOutput {
	s := &AccountSSOGetProvidersOutput{}
	s.InitializeDefaults()
	return s
}

// AccountSSOGetProvidersOutput struct
type AccountSSOGetProvidersOutput struct {
	Providers []*AccountSSOProvider `json:"providers,omitempty" yaml:"providers,omitempty"`
}

// GetProviders returns the value for the field providers
func (e *AccountSSOGetProvidersOutput) GetProviders() []*AccountSSOProvider {
	return e.Providers
}

// SetProviders sets the value for the field providers
func (e *AccountSSOGetProvidersOutput) SetProviders(providers []*AccountSSOProvider) {
	e.Providers = providers
}

// StructPath returns StructPath
func (e *AccountSSOGetProvidersOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOGetProvidersOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountSSOGetProvidersOutput) InitializeDefaults() {
}

// accountSSOGetProvidersOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountSSOGetProvidersOutputAlias AccountSSOGetProvidersOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountSSOGetProvidersOutput) UnmarshalJSON(data []byte) error {
	var alias accountSSOGetProvidersOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountSSOGetProvidersOutput)(&alias)).InitializeDefaults()
	*e = AccountSSOGetProvidersOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountSSOGetProvidersOutput) MarshalJSON() ([]byte, error) {
	alias := accountSSOGetProvidersOutputAlias(e)
	return json.Marshal(alias)
}

// NewAccountSSOAutoJoinPolicyCreateInput creates a new AccountSSOAutoJoinPolicyCreateInput
func NewAccountSSOAutoJoinPolicyCreateInput() *AccountSSOAutoJoinPolicyCreateInput {
	s := &AccountSSOAutoJoinPolicyCreateInput{}
	s.InitializeDefaults()
	return s
}

// AccountSSOAutoJoinPolicyCreateInput struct
type AccountSSOAutoJoinPolicyCreateInput struct {
	Domain             string `json:"domain,omitempty" yaml:"domain,omitempty"`
	OriginAccountName  string `json:"originAccountName,omitempty" yaml:"originAccountName,omitempty"`
	OriginProviderName string `json:"originProviderName,omitempty" yaml:"originProviderName,omitempty"`
}

// GetDomain returns the value for the field domain
func (e *AccountSSOAutoJoinPolicyCreateInput) GetDomain() string {
	return e.Domain
}

// SetDomain sets the value for the field domain
func (e *AccountSSOAutoJoinPolicyCreateInput) SetDomain(domain string) {
	e.Domain = domain
}

// GetOriginAccountName returns the value for the field originAccountName
func (e *AccountSSOAutoJoinPolicyCreateInput) GetOriginAccountName() string {
	return e.OriginAccountName
}

// SetOriginAccountName sets the value for the field originAccountName
func (e *AccountSSOAutoJoinPolicyCreateInput) SetOriginAccountName(originAccountName string) {
	e.OriginAccountName = originAccountName
}

// GetOriginProviderName returns the value for the field originProviderName
func (e *AccountSSOAutoJoinPolicyCreateInput) GetOriginProviderName() string {
	return e.OriginProviderName
}

// SetOriginProviderName sets the value for the field originProviderName
func (e *AccountSSOAutoJoinPolicyCreateInput) SetOriginProviderName(originProviderName string) {
	e.OriginProviderName = originProviderName
}

// StructPath returns StructPath
func (e *AccountSSOAutoJoinPolicyCreateInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOAutoJoinPolicyCreateInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountSSOAutoJoinPolicyCreateInput) InitializeDefaults() {
}

// accountSSOAutoJoinPolicyCreateInputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountSSOAutoJoinPolicyCreateInputAlias AccountSSOAutoJoinPolicyCreateInput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountSSOAutoJoinPolicyCreateInput) UnmarshalJSON(data []byte) error {
	var alias accountSSOAutoJoinPolicyCreateInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountSSOAutoJoinPolicyCreateInput)(&alias)).InitializeDefaults()
	*e = AccountSSOAutoJoinPolicyCreateInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountSSOAutoJoinPolicyCreateInput) MarshalJSON() ([]byte, error) {
	alias := accountSSOAutoJoinPolicyCreateInputAlias(e)
	return json.Marshal(alias)
}

// NewAccountSSOAutoJoinPolicyCreateOutput creates a new AccountSSOAutoJoinPolicyCreateOutput
func NewAccountSSOAutoJoinPolicyCreateOutput() *AccountSSOAutoJoinPolicyCreateOutput {
	s := &AccountSSOAutoJoinPolicyCreateOutput{}
	s.InitializeDefaults()
	return s
}

// AccountSSOAutoJoinPolicyCreateOutput struct
type AccountSSOAutoJoinPolicyCreateOutput struct {
	Policy *AccountSSOAutoJoinPolicy `json:"policy,omitempty" yaml:"policy,omitempty"`
}

// GetPolicy returns the value for the field policy
func (e *AccountSSOAutoJoinPolicyCreateOutput) GetPolicy() *AccountSSOAutoJoinPolicy {
	return e.Policy
}

// SetPolicy sets the value for the field policy
func (e *AccountSSOAutoJoinPolicyCreateOutput) SetPolicy(policy *AccountSSOAutoJoinPolicy) {
	e.Policy = policy
}

// StructPath returns StructPath
func (e *AccountSSOAutoJoinPolicyCreateOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOAutoJoinPolicyCreateOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountSSOAutoJoinPolicyCreateOutput) InitializeDefaults() {
}

// accountSSOAutoJoinPolicyCreateOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountSSOAutoJoinPolicyCreateOutputAlias AccountSSOAutoJoinPolicyCreateOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountSSOAutoJoinPolicyCreateOutput) UnmarshalJSON(data []byte) error {
	var alias accountSSOAutoJoinPolicyCreateOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountSSOAutoJoinPolicyCreateOutput)(&alias)).InitializeDefaults()
	*e = AccountSSOAutoJoinPolicyCreateOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountSSOAutoJoinPolicyCreateOutput) MarshalJSON() ([]byte, error) {
	alias := accountSSOAutoJoinPolicyCreateOutputAlias(e)
	return json.Marshal(alias)
}

// NewAccountSSOAutoJoinPolicyListInput creates a new AccountSSOAutoJoinPolicyListInput
func NewAccountSSOAutoJoinPolicyListInput() *AccountSSOAutoJoinPolicyListInput {
	s := &AccountSSOAutoJoinPolicyListInput{}
	s.InitializeDefaults()
	return s
}

// AccountSSOAutoJoinPolicyListInput struct
type AccountSSOAutoJoinPolicyListInput struct {
}

// StructPath returns StructPath
func (e *AccountSSOAutoJoinPolicyListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOAutoJoinPolicyListInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountSSOAutoJoinPolicyListInput) InitializeDefaults() {
}

// accountSSOAutoJoinPolicyListInputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountSSOAutoJoinPolicyListInputAlias AccountSSOAutoJoinPolicyListInput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountSSOAutoJoinPolicyListInput) UnmarshalJSON(data []byte) error {
	var alias accountSSOAutoJoinPolicyListInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountSSOAutoJoinPolicyListInput)(&alias)).InitializeDefaults()
	*e = AccountSSOAutoJoinPolicyListInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountSSOAutoJoinPolicyListInput) MarshalJSON() ([]byte, error) {
	alias := accountSSOAutoJoinPolicyListInputAlias(e)
	return json.Marshal(alias)
}

// NewAccountSSOAutoJoinPolicyListOutput creates a new AccountSSOAutoJoinPolicyListOutput
func NewAccountSSOAutoJoinPolicyListOutput() *AccountSSOAutoJoinPolicyListOutput {
	s := &AccountSSOAutoJoinPolicyListOutput{}
	s.InitializeDefaults()
	return s
}

// AccountSSOAutoJoinPolicyListOutput struct
type AccountSSOAutoJoinPolicyListOutput struct {
	Policies []*AccountSSOAutoJoinPolicy `json:"policies,omitempty" yaml:"policies,omitempty"`
}

// GetPolicies returns the value for the field policies
func (e *AccountSSOAutoJoinPolicyListOutput) GetPolicies() []*AccountSSOAutoJoinPolicy {
	return e.Policies
}

// SetPolicies sets the value for the field policies
func (e *AccountSSOAutoJoinPolicyListOutput) SetPolicies(policies []*AccountSSOAutoJoinPolicy) {
	e.Policies = policies
}

// StructPath returns StructPath
func (e *AccountSSOAutoJoinPolicyListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOAutoJoinPolicyListOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountSSOAutoJoinPolicyListOutput) InitializeDefaults() {
}

// accountSSOAutoJoinPolicyListOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountSSOAutoJoinPolicyListOutputAlias AccountSSOAutoJoinPolicyListOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountSSOAutoJoinPolicyListOutput) UnmarshalJSON(data []byte) error {
	var alias accountSSOAutoJoinPolicyListOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountSSOAutoJoinPolicyListOutput)(&alias)).InitializeDefaults()
	*e = AccountSSOAutoJoinPolicyListOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountSSOAutoJoinPolicyListOutput) MarshalJSON() ([]byte, error) {
	alias := accountSSOAutoJoinPolicyListOutputAlias(e)
	return json.Marshal(alias)
}

// NewAccountSSOAutoJoinPolicyEnableInput creates a new AccountSSOAutoJoinPolicyEnableInput
func NewAccountSSOAutoJoinPolicyEnableInput() *AccountSSOAutoJoinPolicyEnableInput {
	s := &AccountSSOAutoJoinPolicyEnableInput{}
	s.InitializeDefaults()
	return s
}

// AccountSSOAutoJoinPolicyEnableInput struct
type AccountSSOAutoJoinPolicyEnableInput struct {
	Enabled    bool   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	PolicyName string `json:"policyName,omitempty" yaml:"policyName,omitempty"`
}

// GetEnabled returns the value for the field enabled
func (e *AccountSSOAutoJoinPolicyEnableInput) GetEnabled() bool {
	return e.Enabled
}

// SetEnabled sets the value for the field enabled
func (e *AccountSSOAutoJoinPolicyEnableInput) SetEnabled(enabled bool) {
	e.Enabled = enabled
}

// GetPolicyName returns the value for the field policyName
func (e *AccountSSOAutoJoinPolicyEnableInput) GetPolicyName() string {
	return e.PolicyName
}

// SetPolicyName sets the value for the field policyName
func (e *AccountSSOAutoJoinPolicyEnableInput) SetPolicyName(policyName string) {
	e.PolicyName = policyName
}

// StructPath returns StructPath
func (e *AccountSSOAutoJoinPolicyEnableInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOAutoJoinPolicyEnableInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountSSOAutoJoinPolicyEnableInput) InitializeDefaults() {
}

// accountSSOAutoJoinPolicyEnableInputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountSSOAutoJoinPolicyEnableInputAlias AccountSSOAutoJoinPolicyEnableInput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountSSOAutoJoinPolicyEnableInput) UnmarshalJSON(data []byte) error {
	var alias accountSSOAutoJoinPolicyEnableInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountSSOAutoJoinPolicyEnableInput)(&alias)).InitializeDefaults()
	*e = AccountSSOAutoJoinPolicyEnableInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountSSOAutoJoinPolicyEnableInput) MarshalJSON() ([]byte, error) {
	alias := accountSSOAutoJoinPolicyEnableInputAlias(e)
	return json.Marshal(alias)
}

// NewAccountSSOAutoJoinPolicyEnableOutput creates a new AccountSSOAutoJoinPolicyEnableOutput
func NewAccountSSOAutoJoinPolicyEnableOutput() *AccountSSOAutoJoinPolicyEnableOutput {
	s := &AccountSSOAutoJoinPolicyEnableOutput{}
	s.InitializeDefaults()
	return s
}

// AccountSSOAutoJoinPolicyEnableOutput struct
type AccountSSOAutoJoinPolicyEnableOutput struct {
}

// StructPath returns StructPath
func (e *AccountSSOAutoJoinPolicyEnableOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOAutoJoinPolicyEnableOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountSSOAutoJoinPolicyEnableOutput) InitializeDefaults() {
}

// accountSSOAutoJoinPolicyEnableOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountSSOAutoJoinPolicyEnableOutputAlias AccountSSOAutoJoinPolicyEnableOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountSSOAutoJoinPolicyEnableOutput) UnmarshalJSON(data []byte) error {
	var alias accountSSOAutoJoinPolicyEnableOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountSSOAutoJoinPolicyEnableOutput)(&alias)).InitializeDefaults()
	*e = AccountSSOAutoJoinPolicyEnableOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountSSOAutoJoinPolicyEnableOutput) MarshalJSON() ([]byte, error) {
	alias := accountSSOAutoJoinPolicyEnableOutputAlias(e)
	return json.Marshal(alias)
}

// NewInvalidUsernameProblem creates a new InvalidUsernameProblem
func NewInvalidUsernameProblem() *InvalidUsernameProblem {
	s := &InvalidUsernameProblem{}
	s.InitializeDefaults()
	return s
}

// InvalidUsernameProblem struct
type InvalidUsernameProblem struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InvalidUsernameProblem) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InvalidUsernameProblem]
func (e *InvalidUsernameProblem) Is(err error) bool {
	_, ok := err.(*InvalidUsernameProblem)
	return ok
}

// IsInvalidUsernameProblem indicates whether the given error chain contains an error of type [InvalidUsernameProblem]
func IsInvalidUsernameProblem(err error) bool {
	return errors.Is(err, &InvalidUsernameProblem{})
}

// GetMessage returns the value for the field message
func (e *InvalidUsernameProblem) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InvalidUsernameProblem) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InvalidUsernameProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvalidUsernameProblem.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvalidUsernameProblem) InitializeDefaults() {
}

// invalidUsernameProblemAlias is defined to help pre and post JSON marshaling without recursive loops
type invalidUsernameProblemAlias InvalidUsernameProblem

// UnmarshalJSON implements json.Unmarshaler
func (e *InvalidUsernameProblem) UnmarshalJSON(data []byte) error {
	var alias invalidUsernameProblemAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvalidUsernameProblem)(&alias)).InitializeDefaults()
	*e = InvalidUsernameProblem(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvalidUsernameProblem) MarshalJSON() ([]byte, error) {
	alias := invalidUsernameProblemAlias(e)
	return json.Marshal(alias)
}

// NewInvalidRoleNameProblem creates a new InvalidRoleNameProblem
func NewInvalidRoleNameProblem() *InvalidRoleNameProblem {
	s := &InvalidRoleNameProblem{}
	s.InitializeDefaults()
	return s
}

// InvalidRoleNameProblem struct
type InvalidRoleNameProblem struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InvalidRoleNameProblem) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InvalidRoleNameProblem]
func (e *InvalidRoleNameProblem) Is(err error) bool {
	_, ok := err.(*InvalidRoleNameProblem)
	return ok
}

// IsInvalidRoleNameProblem indicates whether the given error chain contains an error of type [InvalidRoleNameProblem]
func IsInvalidRoleNameProblem(err error) bool {
	return errors.Is(err, &InvalidRoleNameProblem{})
}

// GetMessage returns the value for the field message
func (e *InvalidRoleNameProblem) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InvalidRoleNameProblem) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InvalidRoleNameProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvalidRoleNameProblem.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvalidRoleNameProblem) InitializeDefaults() {
}

// invalidRoleNameProblemAlias is defined to help pre and post JSON marshaling without recursive loops
type invalidRoleNameProblemAlias InvalidRoleNameProblem

// UnmarshalJSON implements json.Unmarshaler
func (e *InvalidRoleNameProblem) UnmarshalJSON(data []byte) error {
	var alias invalidRoleNameProblemAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvalidRoleNameProblem)(&alias)).InitializeDefaults()
	*e = InvalidRoleNameProblem(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvalidRoleNameProblem) MarshalJSON() ([]byte, error) {
	alias := invalidRoleNameProblemAlias(e)
	return json.Marshal(alias)
}

// NewMemberAccount creates a new MemberAccount
func NewMemberAccount() *MemberAccount {
	s := &MemberAccount{}
	s.InitializeDefaults()
	return s
}

// MemberAccount struct
type MemberAccount struct {
	AccountName string `json:"accountName,omitempty" yaml:"accountName,omitempty"`
}

// GetAccountName returns the value for the field accountName
func (e *MemberAccount) GetAccountName() string {
	return e.AccountName
}

// SetAccountName sets the value for the field accountName
func (e *MemberAccount) SetAccountName(accountName string) {
	e.AccountName = accountName
}

// StructPath returns StructPath
func (e *MemberAccount) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathMemberAccount.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *MemberAccount) InitializeDefaults() {
}

// memberAccountAlias is defined to help pre and post JSON marshaling without recursive loops
type memberAccountAlias MemberAccount

// UnmarshalJSON implements json.Unmarshaler
func (e *MemberAccount) UnmarshalJSON(data []byte) error {
	var alias memberAccountAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*MemberAccount)(&alias)).InitializeDefaults()
	*e = MemberAccount(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e MemberAccount) MarshalJSON() ([]byte, error) {
	alias := memberAccountAlias(e)
	return json.Marshal(alias)
}

// NewCredentialInfo creates a new CredentialInfo
func NewCredentialInfo() *CredentialInfo {
	s := &CredentialInfo{}
	s.InitializeDefaults()
	return s
}

// CredentialInfo struct
type CredentialInfo struct {
	AccessKeyID string `json:"accessKeyID,omitempty" yaml:"accessKeyID,omitempty"`
}

// GetAccessKeyID returns the value for the field accessKeyID
func (e *CredentialInfo) GetAccessKeyID() string {
	return e.AccessKeyID
}

// SetAccessKeyID sets the value for the field accessKeyID
func (e *CredentialInfo) SetAccessKeyID(accessKeyID string) {
	e.AccessKeyID = accessKeyID
}

// StructPath returns StructPath
func (e *CredentialInfo) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathCredentialInfo.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *CredentialInfo) InitializeDefaults() {
}

// credentialInfoAlias is defined to help pre and post JSON marshaling without recursive loops
type credentialInfoAlias CredentialInfo

// UnmarshalJSON implements json.Unmarshaler
func (e *CredentialInfo) UnmarshalJSON(data []byte) error {
	var alias credentialInfoAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*CredentialInfo)(&alias)).InitializeDefaults()
	*e = CredentialInfo(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e CredentialInfo) MarshalJSON() ([]byte, error) {
	alias := credentialInfoAlias(e)
	return json.Marshal(alias)
}

// NewIdentityPolicyStatement creates a new IdentityPolicyStatement
func NewIdentityPolicyStatement() *IdentityPolicyStatement {
	s := &IdentityPolicyStatement{}
	s.InitializeDefaults()
	return s
}

// IdentityPolicyStatement struct
type IdentityPolicyStatement struct {
	Actions   []string              `json:"actions,omitempty" yaml:"actions,omitempty"`
	Effect    PolicyStatementEffect `json:"effect,omitempty" yaml:"effect,omitempty"`
	Resources []string              `json:"resources,omitempty" yaml:"resources,omitempty"`
}

// GetActions returns the value for the field actions
func (e *IdentityPolicyStatement) GetActions() []string {
	return e.Actions
}

// SetActions sets the value for the field actions
func (e *IdentityPolicyStatement) SetActions(actions []string) {
	e.Actions = actions
}

// GetEffect returns the value for the field effect
func (e *IdentityPolicyStatement) GetEffect() PolicyStatementEffect {
	return e.Effect
}

// SetEffect sets the value for the field effect
func (e *IdentityPolicyStatement) SetEffect(effect PolicyStatementEffect) {
	e.Effect = effect
}

// GetResources returns the value for the field resources
func (e *IdentityPolicyStatement) GetResources() []string {
	return e.Resources
}

// SetResources sets the value for the field resources
func (e *IdentityPolicyStatement) SetResources(resources []string) {
	e.Resources = resources
}

// StructPath returns StructPath
func (e *IdentityPolicyStatement) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyStatement.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *IdentityPolicyStatement) InitializeDefaults() {
	if e.Effect == "" {
		e.Effect = PolicyStatementEffectDeny
	}
}

// identityPolicyStatementAlias is defined to help pre and post JSON marshaling without recursive loops
type identityPolicyStatementAlias IdentityPolicyStatement

// UnmarshalJSON implements json.Unmarshaler
func (e *IdentityPolicyStatement) UnmarshalJSON(data []byte) error {
	var alias identityPolicyStatementAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*IdentityPolicyStatement)(&alias)).InitializeDefaults()
	*e = IdentityPolicyStatement(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e IdentityPolicyStatement) MarshalJSON() ([]byte, error) {
	alias := identityPolicyStatementAlias(e)
	if alias.Effect == PolicyStatementEffectDeny {
		alias.Effect = ""
	}
	return json.Marshal(alias)
}

// NewIdentityPolicy creates a new IdentityPolicy
func NewIdentityPolicy() *IdentityPolicy {
	s := &IdentityPolicy{}
	s.InitializeDefaults()
	return s
}

// IdentityPolicy struct
type IdentityPolicy struct {
	Builtin    bool                       `json:"builtin,omitempty" yaml:"builtin,omitempty"`
	Name       string                     `json:"name,omitempty" yaml:"name,omitempty"`
	Statements []*IdentityPolicyStatement `json:"statements,omitempty" yaml:"statements,omitempty"`
}

// GetBuiltin returns the value for the field builtin
func (e *IdentityPolicy) GetBuiltin() bool {
	return e.Builtin
}

// SetBuiltin sets the value for the field builtin
func (e *IdentityPolicy) SetBuiltin(builtin bool) {
	e.Builtin = builtin
}

// GetName returns the value for the field name
func (e *IdentityPolicy) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *IdentityPolicy) SetName(name string) {
	e.Name = name
}

// GetStatements returns the value for the field statements
func (e *IdentityPolicy) GetStatements() []*IdentityPolicyStatement {
	return e.Statements
}

// SetStatements sets the value for the field statements
func (e *IdentityPolicy) SetStatements(statements []*IdentityPolicyStatement) {
	e.Statements = statements
}

// StructPath returns StructPath
func (e *IdentityPolicy) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicy.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *IdentityPolicy) InitializeDefaults() {
}

// identityPolicyAlias is defined to help pre and post JSON marshaling without recursive loops
type identityPolicyAlias IdentityPolicy

// UnmarshalJSON implements json.Unmarshaler
func (e *IdentityPolicy) UnmarshalJSON(data []byte) error {
	var alias identityPolicyAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*IdentityPolicy)(&alias)).InitializeDefaults()
	*e = IdentityPolicy(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e IdentityPolicy) MarshalJSON() ([]byte, error) {
	alias := identityPolicyAlias(e)
	return json.Marshal(alias)
}

// NewIdentityPolicyAttachment creates a new IdentityPolicyAttachment
func NewIdentityPolicyAttachment() *IdentityPolicyAttachment {
	s := &IdentityPolicyAttachment{}
	s.InitializeDefaults()
	return s
}

// IdentityPolicyAttachment struct
type IdentityPolicyAttachment struct {
	PolicyName string `json:"policyName,omitempty" yaml:"policyName,omitempty"`
}

// GetPolicyName returns the value for the field policyName
func (e *IdentityPolicyAttachment) GetPolicyName() string {
	return e.PolicyName
}

// SetPolicyName sets the value for the field policyName
func (e *IdentityPolicyAttachment) SetPolicyName(policyName string) {
	e.PolicyName = policyName
}

// StructPath returns StructPath
func (e *IdentityPolicyAttachment) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyAttachment.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *IdentityPolicyAttachment) InitializeDefaults() {
}

// identityPolicyAttachmentAlias is defined to help pre and post JSON marshaling without recursive loops
type identityPolicyAttachmentAlias IdentityPolicyAttachment

// UnmarshalJSON implements json.Unmarshaler
func (e *IdentityPolicyAttachment) UnmarshalJSON(data []byte) error {
	var alias identityPolicyAttachmentAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*IdentityPolicyAttachment)(&alias)).InitializeDefaults()
	*e = IdentityPolicyAttachment(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e IdentityPolicyAttachment) MarshalJSON() ([]byte, error) {
	alias := identityPolicyAttachmentAlias(e)
	return json.Marshal(alias)
}

// NewUserNotFoundProblem creates a new UserNotFoundProblem
func NewUserNotFoundProblem() *UserNotFoundProblem {
	s := &UserNotFoundProblem{}
	s.InitializeDefaults()
	return s
}

// UserNotFoundProblem struct
type UserNotFoundProblem struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *UserNotFoundProblem) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [UserNotFoundProblem]
func (e *UserNotFoundProblem) Is(err error) bool {
	_, ok := err.(*UserNotFoundProblem)
	return ok
}

// IsUserNotFoundProblem indicates whether the given error chain contains an error of type [UserNotFoundProblem]
func IsUserNotFoundProblem(err error) bool {
	return errors.Is(err, &UserNotFoundProblem{})
}

// GetMessage returns the value for the field message
func (e *UserNotFoundProblem) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *UserNotFoundProblem) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *UserNotFoundProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserNotFoundProblem.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserNotFoundProblem) InitializeDefaults() {
}

// userNotFoundProblemAlias is defined to help pre and post JSON marshaling without recursive loops
type userNotFoundProblemAlias UserNotFoundProblem

// UnmarshalJSON implements json.Unmarshaler
func (e *UserNotFoundProblem) UnmarshalJSON(data []byte) error {
	var alias userNotFoundProblemAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserNotFoundProblem)(&alias)).InitializeDefaults()
	*e = UserNotFoundProblem(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserNotFoundProblem) MarshalJSON() ([]byte, error) {
	alias := userNotFoundProblemAlias(e)
	return json.Marshal(alias)
}

// NewRoleNotFoundProblem creates a new RoleNotFoundProblem
func NewRoleNotFoundProblem() *RoleNotFoundProblem {
	s := &RoleNotFoundProblem{}
	s.InitializeDefaults()
	return s
}

// RoleNotFoundProblem struct
type RoleNotFoundProblem struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *RoleNotFoundProblem) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [RoleNotFoundProblem]
func (e *RoleNotFoundProblem) Is(err error) bool {
	_, ok := err.(*RoleNotFoundProblem)
	return ok
}

// IsRoleNotFoundProblem indicates whether the given error chain contains an error of type [RoleNotFoundProblem]
func IsRoleNotFoundProblem(err error) bool {
	return errors.Is(err, &RoleNotFoundProblem{})
}

// GetMessage returns the value for the field message
func (e *RoleNotFoundProblem) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *RoleNotFoundProblem) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *RoleNotFoundProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleNotFoundProblem.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleNotFoundProblem) InitializeDefaults() {
}

// roleNotFoundProblemAlias is defined to help pre and post JSON marshaling without recursive loops
type roleNotFoundProblemAlias RoleNotFoundProblem

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleNotFoundProblem) UnmarshalJSON(data []byte) error {
	var alias roleNotFoundProblemAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleNotFoundProblem)(&alias)).InitializeDefaults()
	*e = RoleNotFoundProblem(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleNotFoundProblem) MarshalJSON() ([]byte, error) {
	alias := roleNotFoundProblemAlias(e)
	return json.Marshal(alias)
}

// NewIdentityInUseProblem creates a new IdentityInUseProblem
func NewIdentityInUseProblem() *IdentityInUseProblem {
	s := &IdentityInUseProblem{}
	s.InitializeDefaults()
	return s
}

// IdentityInUseProblem - Occurs when an identity is in use and cannot be deleted because it is attached to a resource
type IdentityInUseProblem struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *IdentityInUseProblem) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [IdentityInUseProblem]
func (e *IdentityInUseProblem) Is(err error) bool {
	_, ok := err.(*IdentityInUseProblem)
	return ok
}

// IsIdentityInUseProblem indicates whether the given error chain contains an error of type [IdentityInUseProblem]
func IsIdentityInUseProblem(err error) bool {
	return errors.Is(err, &IdentityInUseProblem{})
}

// GetMessage returns the value for the field message
func (e *IdentityInUseProblem) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *IdentityInUseProblem) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *IdentityInUseProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityInUseProblem.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *IdentityInUseProblem) InitializeDefaults() {
}

// identityInUseProblemAlias is defined to help pre and post JSON marshaling without recursive loops
type identityInUseProblemAlias IdentityInUseProblem

// UnmarshalJSON implements json.Unmarshaler
func (e *IdentityInUseProblem) UnmarshalJSON(data []byte) error {
	var alias identityInUseProblemAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*IdentityInUseProblem)(&alias)).InitializeDefaults()
	*e = IdentityInUseProblem(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e IdentityInUseProblem) MarshalJSON() ([]byte, error) {
	alias := identityInUseProblemAlias(e)
	return json.Marshal(alias)
}

// NewUserCreateInput creates a new UserCreateInput
func NewUserCreateInput() *UserCreateInput {
	s := &UserCreateInput{}
	s.InitializeDefaults()
	return s
}

// UserCreateInput struct
type UserCreateInput struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Username    string `json:"username,omitempty" yaml:"username,omitempty"`
}

// GetDescription returns the value for the field description
func (e *UserCreateInput) GetDescription() string {
	return e.Description
}

// SetDescription sets the value for the field description
func (e *UserCreateInput) SetDescription(description string) {
	e.Description = description
}

// GetUsername returns the value for the field username
func (e *UserCreateInput) GetUsername() string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *UserCreateInput) SetUsername(username string) {
	e.Username = username
}

// StructPath returns StructPath
func (e *UserCreateInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserCreateInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserCreateInput) InitializeDefaults() {
}

// userCreateInputAlias is defined to help pre and post JSON marshaling without recursive loops
type userCreateInputAlias UserCreateInput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserCreateInput) UnmarshalJSON(data []byte) error {
	var alias userCreateInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserCreateInput)(&alias)).InitializeDefaults()
	*e = UserCreateInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserCreateInput) MarshalJSON() ([]byte, error) {
	alias := userCreateInputAlias(e)
	return json.Marshal(alias)
}

// NewUserCreateOutput creates a new UserCreateOutput
func NewUserCreateOutput() *UserCreateOutput {
	s := &UserCreateOutput{}
	s.InitializeDefaults()
	return s
}

// UserCreateOutput struct
type UserCreateOutput struct {
	User *UserInformation `json:"user,omitempty" yaml:"user,omitempty"`
}

// GetUser returns the value for the field user
func (e *UserCreateOutput) GetUser() *UserInformation {
	return e.User
}

// SetUser sets the value for the field user
func (e *UserCreateOutput) SetUser(user *UserInformation) {
	e.User = user
}

// StructPath returns StructPath
func (e *UserCreateOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserCreateOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserCreateOutput) InitializeDefaults() {
}

// userCreateOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type userCreateOutputAlias UserCreateOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserCreateOutput) UnmarshalJSON(data []byte) error {
	var alias userCreateOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserCreateOutput)(&alias)).InitializeDefaults()
	*e = UserCreateOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserCreateOutput) MarshalJSON() ([]byte, error) {
	alias := userCreateOutputAlias(e)
	return json.Marshal(alias)
}

// NewUserGetInput creates a new UserGetInput
func NewUserGetInput() *UserGetInput {
	s := &UserGetInput{}
	s.InitializeDefaults()
	return s
}

// UserGetInput struct
type UserGetInput struct {
	Username *string `json:"username,omitempty" yaml:"username,omitempty"`
}

// GetUsername returns the value for the field username
func (e *UserGetInput) GetUsername() *string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *UserGetInput) SetUsername(username *string) {
	e.Username = username
}

// StructPath returns StructPath
func (e *UserGetInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserGetInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserGetInput) InitializeDefaults() {
}

// userGetInputAlias is defined to help pre and post JSON marshaling without recursive loops
type userGetInputAlias UserGetInput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserGetInput) UnmarshalJSON(data []byte) error {
	var alias userGetInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserGetInput)(&alias)).InitializeDefaults()
	*e = UserGetInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserGetInput) MarshalJSON() ([]byte, error) {
	alias := userGetInputAlias(e)
	return json.Marshal(alias)
}

// NewUserGetOutput creates a new UserGetOutput
func NewUserGetOutput() *UserGetOutput {
	s := &UserGetOutput{}
	s.InitializeDefaults()
	return s
}

// UserGetOutput struct
type UserGetOutput struct {
	User *UserInformation `json:"user,omitempty" yaml:"user,omitempty"`
}

// GetUser returns the value for the field user
func (e *UserGetOutput) GetUser() *UserInformation {
	return e.User
}

// SetUser sets the value for the field user
func (e *UserGetOutput) SetUser(user *UserInformation) {
	e.User = user
}

// StructPath returns StructPath
func (e *UserGetOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserGetOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserGetOutput) InitializeDefaults() {
}

// userGetOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type userGetOutputAlias UserGetOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserGetOutput) UnmarshalJSON(data []byte) error {
	var alias userGetOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserGetOutput)(&alias)).InitializeDefaults()
	*e = UserGetOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserGetOutput) MarshalJSON() ([]byte, error) {
	alias := userGetOutputAlias(e)
	return json.Marshal(alias)
}

// NewUserGetUserNotAvailableProblem creates a new UserGetUserNotAvailableProblem
func NewUserGetUserNotAvailableProblem() *UserGetUserNotAvailableProblem {
	s := &UserGetUserNotAvailableProblem{}
	s.InitializeDefaults()
	return s
}

// UserGetUserNotAvailableProblem struct
type UserGetUserNotAvailableProblem struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *UserGetUserNotAvailableProblem) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [UserGetUserNotAvailableProblem]
func (e *UserGetUserNotAvailableProblem) Is(err error) bool {
	_, ok := err.(*UserGetUserNotAvailableProblem)
	return ok
}

// IsUserGetUserNotAvailableProblem indicates whether the given error chain contains an error of type [UserGetUserNotAvailableProblem]
func IsUserGetUserNotAvailableProblem(err error) bool {
	return errors.Is(err, &UserGetUserNotAvailableProblem{})
}

// GetMessage returns the value for the field message
func (e *UserGetUserNotAvailableProblem) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *UserGetUserNotAvailableProblem) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *UserGetUserNotAvailableProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserGetUserNotAvailableProblem.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserGetUserNotAvailableProblem) InitializeDefaults() {
}

// userGetUserNotAvailableProblemAlias is defined to help pre and post JSON marshaling without recursive loops
type userGetUserNotAvailableProblemAlias UserGetUserNotAvailableProblem

// UnmarshalJSON implements json.Unmarshaler
func (e *UserGetUserNotAvailableProblem) UnmarshalJSON(data []byte) error {
	var alias userGetUserNotAvailableProblemAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserGetUserNotAvailableProblem)(&alias)).InitializeDefaults()
	*e = UserGetUserNotAvailableProblem(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserGetUserNotAvailableProblem) MarshalJSON() ([]byte, error) {
	alias := userGetUserNotAvailableProblemAlias(e)
	return json.Marshal(alias)
}

// NewUserDestroyInput creates a new UserDestroyInput
func NewUserDestroyInput() *UserDestroyInput {
	s := &UserDestroyInput{}
	s.InitializeDefaults()
	return s
}

// UserDestroyInput struct
type UserDestroyInput struct {
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}

// GetUsername returns the value for the field username
func (e *UserDestroyInput) GetUsername() string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *UserDestroyInput) SetUsername(username string) {
	e.Username = username
}

// StructPath returns StructPath
func (e *UserDestroyInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserDestroyInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserDestroyInput) InitializeDefaults() {
}

// userDestroyInputAlias is defined to help pre and post JSON marshaling without recursive loops
type userDestroyInputAlias UserDestroyInput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserDestroyInput) UnmarshalJSON(data []byte) error {
	var alias userDestroyInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserDestroyInput)(&alias)).InitializeDefaults()
	*e = UserDestroyInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserDestroyInput) MarshalJSON() ([]byte, error) {
	alias := userDestroyInputAlias(e)
	return json.Marshal(alias)
}

// NewUserDestroyOutput creates a new UserDestroyOutput
func NewUserDestroyOutput() *UserDestroyOutput {
	s := &UserDestroyOutput{}
	s.InitializeDefaults()
	return s
}

// UserDestroyOutput struct
type UserDestroyOutput struct {
}

// StructPath returns StructPath
func (e *UserDestroyOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserDestroyOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserDestroyOutput) InitializeDefaults() {
}

// userDestroyOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type userDestroyOutputAlias UserDestroyOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserDestroyOutput) UnmarshalJSON(data []byte) error {
	var alias userDestroyOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserDestroyOutput)(&alias)).InitializeDefaults()
	*e = UserDestroyOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserDestroyOutput) MarshalJSON() ([]byte, error) {
	alias := userDestroyOutputAlias(e)
	return json.Marshal(alias)
}

// NewUserListInput creates a new UserListInput
func NewUserListInput() *UserListInput {
	s := &UserListInput{}
	s.InitializeDefaults()
	return s
}

// UserListInput struct
type UserListInput struct {
}

// StructPath returns StructPath
func (e *UserListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserListInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserListInput) InitializeDefaults() {
}

// userListInputAlias is defined to help pre and post JSON marshaling without recursive loops
type userListInputAlias UserListInput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserListInput) UnmarshalJSON(data []byte) error {
	var alias userListInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserListInput)(&alias)).InitializeDefaults()
	*e = UserListInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserListInput) MarshalJSON() ([]byte, error) {
	alias := userListInputAlias(e)
	return json.Marshal(alias)
}

// NewUserListOutput creates a new UserListOutput
func NewUserListOutput() *UserListOutput {
	s := &UserListOutput{}
	s.InitializeDefaults()
	return s
}

// UserListOutput struct
type UserListOutput struct {
	Users []*UserInformation `json:"users,omitempty" yaml:"users,omitempty"`
}

// GetUsers returns the value for the field users
func (e *UserListOutput) GetUsers() []*UserInformation {
	return e.Users
}

// SetUsers sets the value for the field users
func (e *UserListOutput) SetUsers(users []*UserInformation) {
	e.Users = users
}

// StructPath returns StructPath
func (e *UserListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserListOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserListOutput) InitializeDefaults() {
}

// userListOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type userListOutputAlias UserListOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserListOutput) UnmarshalJSON(data []byte) error {
	var alias userListOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserListOutput)(&alias)).InitializeDefaults()
	*e = UserListOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserListOutput) MarshalJSON() ([]byte, error) {
	alias := userListOutputAlias(e)
	return json.Marshal(alias)
}

// NewUserMemberAccountsInput creates a new UserMemberAccountsInput
func NewUserMemberAccountsInput() *UserMemberAccountsInput {
	s := &UserMemberAccountsInput{}
	s.InitializeDefaults()
	return s
}

// UserMemberAccountsInput struct
type UserMemberAccountsInput struct {
}

// StructPath returns StructPath
func (e *UserMemberAccountsInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserMemberAccountsInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserMemberAccountsInput) InitializeDefaults() {
}

// userMemberAccountsInputAlias is defined to help pre and post JSON marshaling without recursive loops
type userMemberAccountsInputAlias UserMemberAccountsInput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserMemberAccountsInput) UnmarshalJSON(data []byte) error {
	var alias userMemberAccountsInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserMemberAccountsInput)(&alias)).InitializeDefaults()
	*e = UserMemberAccountsInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserMemberAccountsInput) MarshalJSON() ([]byte, error) {
	alias := userMemberAccountsInputAlias(e)
	return json.Marshal(alias)
}

// NewUserMemberAccountsOutput creates a new UserMemberAccountsOutput
func NewUserMemberAccountsOutput() *UserMemberAccountsOutput {
	s := &UserMemberAccountsOutput{}
	s.InitializeDefaults()
	return s
}

// UserMemberAccountsOutput struct
type UserMemberAccountsOutput struct {
	Accounts []*MemberAccount `json:"accounts,omitempty" yaml:"accounts,omitempty"`
}

// GetAccounts returns the value for the field accounts
func (e *UserMemberAccountsOutput) GetAccounts() []*MemberAccount {
	return e.Accounts
}

// SetAccounts sets the value for the field accounts
func (e *UserMemberAccountsOutput) SetAccounts(accounts []*MemberAccount) {
	e.Accounts = accounts
}

// StructPath returns StructPath
func (e *UserMemberAccountsOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserMemberAccountsOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserMemberAccountsOutput) InitializeDefaults() {
}

// userMemberAccountsOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type userMemberAccountsOutputAlias UserMemberAccountsOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserMemberAccountsOutput) UnmarshalJSON(data []byte) error {
	var alias userMemberAccountsOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserMemberAccountsOutput)(&alias)).InitializeDefaults()
	*e = UserMemberAccountsOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserMemberAccountsOutput) MarshalJSON() ([]byte, error) {
	alias := userMemberAccountsOutputAlias(e)
	return json.Marshal(alias)
}

// NewUserAccessKeyCreateInput creates a new UserAccessKeyCreateInput
func NewUserAccessKeyCreateInput() *UserAccessKeyCreateInput {
	s := &UserAccessKeyCreateInput{}
	s.InitializeDefaults()
	return s
}

// UserAccessKeyCreateInput struct
type UserAccessKeyCreateInput struct {
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}

// GetUsername returns the value for the field username
func (e *UserAccessKeyCreateInput) GetUsername() string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *UserAccessKeyCreateInput) SetUsername(username string) {
	e.Username = username
}

// StructPath returns StructPath
func (e *UserAccessKeyCreateInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserAccessKeyCreateInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserAccessKeyCreateInput) InitializeDefaults() {
}

// userAccessKeyCreateInputAlias is defined to help pre and post JSON marshaling without recursive loops
type userAccessKeyCreateInputAlias UserAccessKeyCreateInput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserAccessKeyCreateInput) UnmarshalJSON(data []byte) error {
	var alias userAccessKeyCreateInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserAccessKeyCreateInput)(&alias)).InitializeDefaults()
	*e = UserAccessKeyCreateInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserAccessKeyCreateInput) MarshalJSON() ([]byte, error) {
	alias := userAccessKeyCreateInputAlias(e)
	return json.Marshal(alias)
}

// NewUserAccessKeyCreateOutput creates a new UserAccessKeyCreateOutput
func NewUserAccessKeyCreateOutput() *UserAccessKeyCreateOutput {
	s := &UserAccessKeyCreateOutput{}
	s.InitializeDefaults()
	return s
}

// UserAccessKeyCreateOutput struct
type UserAccessKeyCreateOutput struct {
	Credentials *Credentials `json:"credentials,omitempty" yaml:"credentials,omitempty"`
}

// GetCredentials returns the value for the field credentials
func (e *UserAccessKeyCreateOutput) GetCredentials() *Credentials {
	return e.Credentials
}

// SetCredentials sets the value for the field credentials
func (e *UserAccessKeyCreateOutput) SetCredentials(credentials *Credentials) {
	e.Credentials = credentials
}

// StructPath returns StructPath
func (e *UserAccessKeyCreateOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserAccessKeyCreateOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserAccessKeyCreateOutput) InitializeDefaults() {
}

// userAccessKeyCreateOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type userAccessKeyCreateOutputAlias UserAccessKeyCreateOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserAccessKeyCreateOutput) UnmarshalJSON(data []byte) error {
	var alias userAccessKeyCreateOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserAccessKeyCreateOutput)(&alias)).InitializeDefaults()
	*e = UserAccessKeyCreateOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserAccessKeyCreateOutput) MarshalJSON() ([]byte, error) {
	alias := userAccessKeyCreateOutputAlias(e)
	return json.Marshal(alias)
}

// NewUserAccessKeyListInput creates a new UserAccessKeyListInput
func NewUserAccessKeyListInput() *UserAccessKeyListInput {
	s := &UserAccessKeyListInput{}
	s.InitializeDefaults()
	return s
}

// UserAccessKeyListInput struct
type UserAccessKeyListInput struct {
	Username *string `json:"username,omitempty" yaml:"username,omitempty"`
}

// GetUsername returns the value for the field username
func (e *UserAccessKeyListInput) GetUsername() *string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *UserAccessKeyListInput) SetUsername(username *string) {
	e.Username = username
}

// StructPath returns StructPath
func (e *UserAccessKeyListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserAccessKeyListInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserAccessKeyListInput) InitializeDefaults() {
}

// userAccessKeyListInputAlias is defined to help pre and post JSON marshaling without recursive loops
type userAccessKeyListInputAlias UserAccessKeyListInput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserAccessKeyListInput) UnmarshalJSON(data []byte) error {
	var alias userAccessKeyListInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserAccessKeyListInput)(&alias)).InitializeDefaults()
	*e = UserAccessKeyListInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserAccessKeyListInput) MarshalJSON() ([]byte, error) {
	alias := userAccessKeyListInputAlias(e)
	return json.Marshal(alias)
}

// NewUserAccessKeyListOutput creates a new UserAccessKeyListOutput
func NewUserAccessKeyListOutput() *UserAccessKeyListOutput {
	s := &UserAccessKeyListOutput{}
	s.InitializeDefaults()
	return s
}

// UserAccessKeyListOutput struct
type UserAccessKeyListOutput struct {
	Credentials []*CredentialInfo `json:"credentials,omitempty" yaml:"credentials,omitempty"`
}

// GetCredentials returns the value for the field credentials
func (e *UserAccessKeyListOutput) GetCredentials() []*CredentialInfo {
	return e.Credentials
}

// SetCredentials sets the value for the field credentials
func (e *UserAccessKeyListOutput) SetCredentials(credentials []*CredentialInfo) {
	e.Credentials = credentials
}

// StructPath returns StructPath
func (e *UserAccessKeyListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserAccessKeyListOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserAccessKeyListOutput) InitializeDefaults() {
}

// userAccessKeyListOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type userAccessKeyListOutputAlias UserAccessKeyListOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserAccessKeyListOutput) UnmarshalJSON(data []byte) error {
	var alias userAccessKeyListOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserAccessKeyListOutput)(&alias)).InitializeDefaults()
	*e = UserAccessKeyListOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserAccessKeyListOutput) MarshalJSON() ([]byte, error) {
	alias := userAccessKeyListOutputAlias(e)
	return json.Marshal(alias)
}

// NewUserAccessKeyDestroyInput creates a new UserAccessKeyDestroyInput
func NewUserAccessKeyDestroyInput() *UserAccessKeyDestroyInput {
	s := &UserAccessKeyDestroyInput{}
	s.InitializeDefaults()
	return s
}

// UserAccessKeyDestroyInput struct
type UserAccessKeyDestroyInput struct {
	AccessKeyID string `json:"accessKeyID,omitempty" yaml:"accessKeyID,omitempty"`
	Username    string `json:"username,omitempty" yaml:"username,omitempty"`
}

// GetAccessKeyID returns the value for the field accessKeyID
func (e *UserAccessKeyDestroyInput) GetAccessKeyID() string {
	return e.AccessKeyID
}

// SetAccessKeyID sets the value for the field accessKeyID
func (e *UserAccessKeyDestroyInput) SetAccessKeyID(accessKeyID string) {
	e.AccessKeyID = accessKeyID
}

// GetUsername returns the value for the field username
func (e *UserAccessKeyDestroyInput) GetUsername() string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *UserAccessKeyDestroyInput) SetUsername(username string) {
	e.Username = username
}

// StructPath returns StructPath
func (e *UserAccessKeyDestroyInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserAccessKeyDestroyInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserAccessKeyDestroyInput) InitializeDefaults() {
}

// userAccessKeyDestroyInputAlias is defined to help pre and post JSON marshaling without recursive loops
type userAccessKeyDestroyInputAlias UserAccessKeyDestroyInput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserAccessKeyDestroyInput) UnmarshalJSON(data []byte) error {
	var alias userAccessKeyDestroyInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserAccessKeyDestroyInput)(&alias)).InitializeDefaults()
	*e = UserAccessKeyDestroyInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserAccessKeyDestroyInput) MarshalJSON() ([]byte, error) {
	alias := userAccessKeyDestroyInputAlias(e)
	return json.Marshal(alias)
}

// NewUserAccessKeyDestroyOutput creates a new UserAccessKeyDestroyOutput
func NewUserAccessKeyDestroyOutput() *UserAccessKeyDestroyOutput {
	s := &UserAccessKeyDestroyOutput{}
	s.InitializeDefaults()
	return s
}

// UserAccessKeyDestroyOutput struct
type UserAccessKeyDestroyOutput struct {
}

// StructPath returns StructPath
func (e *UserAccessKeyDestroyOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserAccessKeyDestroyOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserAccessKeyDestroyOutput) InitializeDefaults() {
}

// userAccessKeyDestroyOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type userAccessKeyDestroyOutputAlias UserAccessKeyDestroyOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserAccessKeyDestroyOutput) UnmarshalJSON(data []byte) error {
	var alias userAccessKeyDestroyOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserAccessKeyDestroyOutput)(&alias)).InitializeDefaults()
	*e = UserAccessKeyDestroyOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserAccessKeyDestroyOutput) MarshalJSON() ([]byte, error) {
	alias := userAccessKeyDestroyOutputAlias(e)
	return json.Marshal(alias)
}

// NewUserIdentityPolicyAttachInput creates a new UserIdentityPolicyAttachInput
func NewUserIdentityPolicyAttachInput() *UserIdentityPolicyAttachInput {
	s := &UserIdentityPolicyAttachInput{}
	s.InitializeDefaults()
	return s
}

// UserIdentityPolicyAttachInput struct
type UserIdentityPolicyAttachInput struct {
	PolicyName string `json:"policyName,omitempty" yaml:"policyName,omitempty"`
	Username   string `json:"username,omitempty" yaml:"username,omitempty"`
}

// GetPolicyName returns the value for the field policyName
func (e *UserIdentityPolicyAttachInput) GetPolicyName() string {
	return e.PolicyName
}

// SetPolicyName sets the value for the field policyName
func (e *UserIdentityPolicyAttachInput) SetPolicyName(policyName string) {
	e.PolicyName = policyName
}

// GetUsername returns the value for the field username
func (e *UserIdentityPolicyAttachInput) GetUsername() string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *UserIdentityPolicyAttachInput) SetUsername(username string) {
	e.Username = username
}

// StructPath returns StructPath
func (e *UserIdentityPolicyAttachInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserIdentityPolicyAttachInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserIdentityPolicyAttachInput) InitializeDefaults() {
}

// userIdentityPolicyAttachInputAlias is defined to help pre and post JSON marshaling without recursive loops
type userIdentityPolicyAttachInputAlias UserIdentityPolicyAttachInput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserIdentityPolicyAttachInput) UnmarshalJSON(data []byte) error {
	var alias userIdentityPolicyAttachInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserIdentityPolicyAttachInput)(&alias)).InitializeDefaults()
	*e = UserIdentityPolicyAttachInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserIdentityPolicyAttachInput) MarshalJSON() ([]byte, error) {
	alias := userIdentityPolicyAttachInputAlias(e)
	return json.Marshal(alias)
}

// NewUserIdentityPolicyAttachOutput creates a new UserIdentityPolicyAttachOutput
func NewUserIdentityPolicyAttachOutput() *UserIdentityPolicyAttachOutput {
	s := &UserIdentityPolicyAttachOutput{}
	s.InitializeDefaults()
	return s
}

// UserIdentityPolicyAttachOutput struct
type UserIdentityPolicyAttachOutput struct {
}

// StructPath returns StructPath
func (e *UserIdentityPolicyAttachOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserIdentityPolicyAttachOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserIdentityPolicyAttachOutput) InitializeDefaults() {
}

// userIdentityPolicyAttachOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type userIdentityPolicyAttachOutputAlias UserIdentityPolicyAttachOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserIdentityPolicyAttachOutput) UnmarshalJSON(data []byte) error {
	var alias userIdentityPolicyAttachOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserIdentityPolicyAttachOutput)(&alias)).InitializeDefaults()
	*e = UserIdentityPolicyAttachOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserIdentityPolicyAttachOutput) MarshalJSON() ([]byte, error) {
	alias := userIdentityPolicyAttachOutputAlias(e)
	return json.Marshal(alias)
}

// NewUserIdentityPolicyListInput creates a new UserIdentityPolicyListInput
func NewUserIdentityPolicyListInput() *UserIdentityPolicyListInput {
	s := &UserIdentityPolicyListInput{}
	s.InitializeDefaults()
	return s
}

// UserIdentityPolicyListInput struct
type UserIdentityPolicyListInput struct {
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}

// GetUsername returns the value for the field username
func (e *UserIdentityPolicyListInput) GetUsername() string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *UserIdentityPolicyListInput) SetUsername(username string) {
	e.Username = username
}

// StructPath returns StructPath
func (e *UserIdentityPolicyListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserIdentityPolicyListInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserIdentityPolicyListInput) InitializeDefaults() {
}

// userIdentityPolicyListInputAlias is defined to help pre and post JSON marshaling without recursive loops
type userIdentityPolicyListInputAlias UserIdentityPolicyListInput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserIdentityPolicyListInput) UnmarshalJSON(data []byte) error {
	var alias userIdentityPolicyListInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserIdentityPolicyListInput)(&alias)).InitializeDefaults()
	*e = UserIdentityPolicyListInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserIdentityPolicyListInput) MarshalJSON() ([]byte, error) {
	alias := userIdentityPolicyListInputAlias(e)
	return json.Marshal(alias)
}

// NewUserIdentityPolicyListOutput creates a new UserIdentityPolicyListOutput
func NewUserIdentityPolicyListOutput() *UserIdentityPolicyListOutput {
	s := &UserIdentityPolicyListOutput{}
	s.InitializeDefaults()
	return s
}

// UserIdentityPolicyListOutput struct
type UserIdentityPolicyListOutput struct {
	Attachments []*IdentityPolicyAttachment `json:"attachments,omitempty" yaml:"attachments,omitempty"`
}

// GetAttachments returns the value for the field attachments
func (e *UserIdentityPolicyListOutput) GetAttachments() []*IdentityPolicyAttachment {
	return e.Attachments
}

// SetAttachments sets the value for the field attachments
func (e *UserIdentityPolicyListOutput) SetAttachments(attachments []*IdentityPolicyAttachment) {
	e.Attachments = attachments
}

// StructPath returns StructPath
func (e *UserIdentityPolicyListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserIdentityPolicyListOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserIdentityPolicyListOutput) InitializeDefaults() {
}

// userIdentityPolicyListOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type userIdentityPolicyListOutputAlias UserIdentityPolicyListOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserIdentityPolicyListOutput) UnmarshalJSON(data []byte) error {
	var alias userIdentityPolicyListOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserIdentityPolicyListOutput)(&alias)).InitializeDefaults()
	*e = UserIdentityPolicyListOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserIdentityPolicyListOutput) MarshalJSON() ([]byte, error) {
	alias := userIdentityPolicyListOutputAlias(e)
	return json.Marshal(alias)
}

// NewUserIdentityPolicyDetachInput creates a new UserIdentityPolicyDetachInput
func NewUserIdentityPolicyDetachInput() *UserIdentityPolicyDetachInput {
	s := &UserIdentityPolicyDetachInput{}
	s.InitializeDefaults()
	return s
}

// UserIdentityPolicyDetachInput struct
type UserIdentityPolicyDetachInput struct {
	PolicyName string `json:"policyName,omitempty" yaml:"policyName,omitempty"`
	Username   string `json:"username,omitempty" yaml:"username,omitempty"`
}

// GetPolicyName returns the value for the field policyName
func (e *UserIdentityPolicyDetachInput) GetPolicyName() string {
	return e.PolicyName
}

// SetPolicyName sets the value for the field policyName
func (e *UserIdentityPolicyDetachInput) SetPolicyName(policyName string) {
	e.PolicyName = policyName
}

// GetUsername returns the value for the field username
func (e *UserIdentityPolicyDetachInput) GetUsername() string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *UserIdentityPolicyDetachInput) SetUsername(username string) {
	e.Username = username
}

// StructPath returns StructPath
func (e *UserIdentityPolicyDetachInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserIdentityPolicyDetachInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserIdentityPolicyDetachInput) InitializeDefaults() {
}

// userIdentityPolicyDetachInputAlias is defined to help pre and post JSON marshaling without recursive loops
type userIdentityPolicyDetachInputAlias UserIdentityPolicyDetachInput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserIdentityPolicyDetachInput) UnmarshalJSON(data []byte) error {
	var alias userIdentityPolicyDetachInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserIdentityPolicyDetachInput)(&alias)).InitializeDefaults()
	*e = UserIdentityPolicyDetachInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserIdentityPolicyDetachInput) MarshalJSON() ([]byte, error) {
	alias := userIdentityPolicyDetachInputAlias(e)
	return json.Marshal(alias)
}

// NewUserIdentityPolicyDetachOutput creates a new UserIdentityPolicyDetachOutput
func NewUserIdentityPolicyDetachOutput() *UserIdentityPolicyDetachOutput {
	s := &UserIdentityPolicyDetachOutput{}
	s.InitializeDefaults()
	return s
}

// UserIdentityPolicyDetachOutput struct
type UserIdentityPolicyDetachOutput struct {
}

// StructPath returns StructPath
func (e *UserIdentityPolicyDetachOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserIdentityPolicyDetachOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserIdentityPolicyDetachOutput) InitializeDefaults() {
}

// userIdentityPolicyDetachOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type userIdentityPolicyDetachOutputAlias UserIdentityPolicyDetachOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserIdentityPolicyDetachOutput) UnmarshalJSON(data []byte) error {
	var alias userIdentityPolicyDetachOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserIdentityPolicyDetachOutput)(&alias)).InitializeDefaults()
	*e = UserIdentityPolicyDetachOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserIdentityPolicyDetachOutput) MarshalJSON() ([]byte, error) {
	alias := userIdentityPolicyDetachOutputAlias(e)
	return json.Marshal(alias)
}

// NewInlinePolicy creates a new InlinePolicy
func NewInlinePolicy() *InlinePolicy {
	s := &InlinePolicy{}
	s.InitializeDefaults()
	return s
}

// InlinePolicy struct
type InlinePolicy struct {
	Statements []*IdentityPolicyStatement `json:"statements,omitempty" yaml:"statements,omitempty"`
}

// GetStatements returns the value for the field statements
func (e *InlinePolicy) GetStatements() []*IdentityPolicyStatement {
	return e.Statements
}

// SetStatements sets the value for the field statements
func (e *InlinePolicy) SetStatements(statements []*IdentityPolicyStatement) {
	e.Statements = statements
}

// StructPath returns StructPath
func (e *InlinePolicy) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInlinePolicy.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InlinePolicy) InitializeDefaults() {
}

// inlinePolicyAlias is defined to help pre and post JSON marshaling without recursive loops
type inlinePolicyAlias InlinePolicy

// UnmarshalJSON implements json.Unmarshaler
func (e *InlinePolicy) UnmarshalJSON(data []byte) error {
	var alias inlinePolicyAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InlinePolicy)(&alias)).InitializeDefaults()
	*e = InlinePolicy(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InlinePolicy) MarshalJSON() ([]byte, error) {
	alias := inlinePolicyAlias(e)
	return json.Marshal(alias)
}

// NewPolicyStructureProblem creates a new PolicyStructureProblem
func NewPolicyStructureProblem() *PolicyStructureProblem {
	s := &PolicyStructureProblem{}
	s.InitializeDefaults()
	return s
}

// PolicyStructureProblem struct
type PolicyStructureProblem struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *PolicyStructureProblem) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [PolicyStructureProblem]
func (e *PolicyStructureProblem) Is(err error) bool {
	_, ok := err.(*PolicyStructureProblem)
	return ok
}

// IsPolicyStructureProblem indicates whether the given error chain contains an error of type [PolicyStructureProblem]
func IsPolicyStructureProblem(err error) bool {
	return errors.Is(err, &PolicyStructureProblem{})
}

// GetMessage returns the value for the field message
func (e *PolicyStructureProblem) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *PolicyStructureProblem) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *PolicyStructureProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathPolicyStructureProblem.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *PolicyStructureProblem) InitializeDefaults() {
}

// policyStructureProblemAlias is defined to help pre and post JSON marshaling without recursive loops
type policyStructureProblemAlias PolicyStructureProblem

// UnmarshalJSON implements json.Unmarshaler
func (e *PolicyStructureProblem) UnmarshalJSON(data []byte) error {
	var alias policyStructureProblemAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*PolicyStructureProblem)(&alias)).InitializeDefaults()
	*e = PolicyStructureProblem(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e PolicyStructureProblem) MarshalJSON() ([]byte, error) {
	alias := policyStructureProblemAlias(e)
	return json.Marshal(alias)
}

// NewRoleCreateInput creates a new RoleCreateInput
func NewRoleCreateInput() *RoleCreateInput {
	s := &RoleCreateInput{}
	s.InitializeDefaults()
	return s
}

// RoleCreateInput struct
type RoleCreateInput struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetDescription returns the value for the field description
func (e *RoleCreateInput) GetDescription() string {
	return e.Description
}

// SetDescription sets the value for the field description
func (e *RoleCreateInput) SetDescription(description string) {
	e.Description = description
}

// GetName returns the value for the field name
func (e *RoleCreateInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *RoleCreateInput) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *RoleCreateInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleCreateInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleCreateInput) InitializeDefaults() {
}

// roleCreateInputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleCreateInputAlias RoleCreateInput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleCreateInput) UnmarshalJSON(data []byte) error {
	var alias roleCreateInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleCreateInput)(&alias)).InitializeDefaults()
	*e = RoleCreateInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleCreateInput) MarshalJSON() ([]byte, error) {
	alias := roleCreateInputAlias(e)
	return json.Marshal(alias)
}

// NewRoleCreateOutput creates a new RoleCreateOutput
func NewRoleCreateOutput() *RoleCreateOutput {
	s := &RoleCreateOutput{}
	s.InitializeDefaults()
	return s
}

// RoleCreateOutput struct
type RoleCreateOutput struct {
	Role *RoleInformation `json:"role,omitempty" yaml:"role,omitempty"`
}

// GetRole returns the value for the field role
func (e *RoleCreateOutput) GetRole() *RoleInformation {
	return e.Role
}

// SetRole sets the value for the field role
func (e *RoleCreateOutput) SetRole(role *RoleInformation) {
	e.Role = role
}

// StructPath returns StructPath
func (e *RoleCreateOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleCreateOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleCreateOutput) InitializeDefaults() {
}

// roleCreateOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleCreateOutputAlias RoleCreateOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleCreateOutput) UnmarshalJSON(data []byte) error {
	var alias roleCreateOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleCreateOutput)(&alias)).InitializeDefaults()
	*e = RoleCreateOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleCreateOutput) MarshalJSON() ([]byte, error) {
	alias := roleCreateOutputAlias(e)
	return json.Marshal(alias)
}

// NewRoleGetInput creates a new RoleGetInput
func NewRoleGetInput() *RoleGetInput {
	s := &RoleGetInput{}
	s.InitializeDefaults()
	return s
}

// RoleGetInput struct
type RoleGetInput struct {
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetName returns the value for the field name
func (e *RoleGetInput) GetName() *string {
	return e.Name
}

// SetName sets the value for the field name
func (e *RoleGetInput) SetName(name *string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *RoleGetInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleGetInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleGetInput) InitializeDefaults() {
}

// roleGetInputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleGetInputAlias RoleGetInput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleGetInput) UnmarshalJSON(data []byte) error {
	var alias roleGetInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleGetInput)(&alias)).InitializeDefaults()
	*e = RoleGetInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleGetInput) MarshalJSON() ([]byte, error) {
	alias := roleGetInputAlias(e)
	return json.Marshal(alias)
}

// NewRoleGetOutput creates a new RoleGetOutput
func NewRoleGetOutput() *RoleGetOutput {
	s := &RoleGetOutput{}
	s.InitializeDefaults()
	return s
}

// RoleGetOutput struct
type RoleGetOutput struct {
	Role *RoleInformation `json:"role,omitempty" yaml:"role,omitempty"`
}

// GetRole returns the value for the field role
func (e *RoleGetOutput) GetRole() *RoleInformation {
	return e.Role
}

// SetRole sets the value for the field role
func (e *RoleGetOutput) SetRole(role *RoleInformation) {
	e.Role = role
}

// StructPath returns StructPath
func (e *RoleGetOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleGetOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleGetOutput) InitializeDefaults() {
}

// roleGetOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleGetOutputAlias RoleGetOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleGetOutput) UnmarshalJSON(data []byte) error {
	var alias roleGetOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleGetOutput)(&alias)).InitializeDefaults()
	*e = RoleGetOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleGetOutput) MarshalJSON() ([]byte, error) {
	alias := roleGetOutputAlias(e)
	return json.Marshal(alias)
}

// NewRoleDestroyInput creates a new RoleDestroyInput
func NewRoleDestroyInput() *RoleDestroyInput {
	s := &RoleDestroyInput{}
	s.InitializeDefaults()
	return s
}

// RoleDestroyInput struct
type RoleDestroyInput struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetName returns the value for the field name
func (e *RoleDestroyInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *RoleDestroyInput) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *RoleDestroyInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleDestroyInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleDestroyInput) InitializeDefaults() {
}

// roleDestroyInputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleDestroyInputAlias RoleDestroyInput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleDestroyInput) UnmarshalJSON(data []byte) error {
	var alias roleDestroyInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleDestroyInput)(&alias)).InitializeDefaults()
	*e = RoleDestroyInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleDestroyInput) MarshalJSON() ([]byte, error) {
	alias := roleDestroyInputAlias(e)
	return json.Marshal(alias)
}

// NewRoleDestroyOutput creates a new RoleDestroyOutput
func NewRoleDestroyOutput() *RoleDestroyOutput {
	s := &RoleDestroyOutput{}
	s.InitializeDefaults()
	return s
}

// RoleDestroyOutput struct
type RoleDestroyOutput struct {
}

// StructPath returns StructPath
func (e *RoleDestroyOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleDestroyOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleDestroyOutput) InitializeDefaults() {
}

// roleDestroyOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleDestroyOutputAlias RoleDestroyOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleDestroyOutput) UnmarshalJSON(data []byte) error {
	var alias roleDestroyOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleDestroyOutput)(&alias)).InitializeDefaults()
	*e = RoleDestroyOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleDestroyOutput) MarshalJSON() ([]byte, error) {
	alias := roleDestroyOutputAlias(e)
	return json.Marshal(alias)
}

// NewRoleListInput creates a new RoleListInput
func NewRoleListInput() *RoleListInput {
	s := &RoleListInput{}
	s.InitializeDefaults()
	return s
}

// RoleListInput struct
type RoleListInput struct {
}

// StructPath returns StructPath
func (e *RoleListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleListInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleListInput) InitializeDefaults() {
}

// roleListInputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleListInputAlias RoleListInput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleListInput) UnmarshalJSON(data []byte) error {
	var alias roleListInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleListInput)(&alias)).InitializeDefaults()
	*e = RoleListInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleListInput) MarshalJSON() ([]byte, error) {
	alias := roleListInputAlias(e)
	return json.Marshal(alias)
}

// NewRoleListOutput creates a new RoleListOutput
func NewRoleListOutput() *RoleListOutput {
	s := &RoleListOutput{}
	s.InitializeDefaults()
	return s
}

// RoleListOutput struct
type RoleListOutput struct {
	Roles []*RoleInformation `json:"roles,omitempty" yaml:"roles,omitempty"`
}

// GetRoles returns the value for the field roles
func (e *RoleListOutput) GetRoles() []*RoleInformation {
	return e.Roles
}

// SetRoles sets the value for the field roles
func (e *RoleListOutput) SetRoles(roles []*RoleInformation) {
	e.Roles = roles
}

// StructPath returns StructPath
func (e *RoleListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleListOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleListOutput) InitializeDefaults() {
}

// roleListOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleListOutputAlias RoleListOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleListOutput) UnmarshalJSON(data []byte) error {
	var alias roleListOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleListOutput)(&alias)).InitializeDefaults()
	*e = RoleListOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleListOutput) MarshalJSON() ([]byte, error) {
	alias := roleListOutputAlias(e)
	return json.Marshal(alias)
}

// NewRoleAssumeInput creates a new RoleAssumeInput
func NewRoleAssumeInput() *RoleAssumeInput {
	s := &RoleAssumeInput{}
	s.InitializeDefaults()
	return s
}

// RoleAssumeInput struct
type RoleAssumeInput struct {
	// duration in seconds for the assumed role
	DurationSeconds int32 `json:"durationSeconds,omitempty" yaml:"durationSeconds,omitempty"`
	// optional inline policy to attach to the assumed role
	InlinePolicy *InlinePolicy `json:"inlinePolicy,omitempty" yaml:"inlinePolicy,omitempty"`
	// name of the role to assume
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetDurationSeconds returns the value for the field durationSeconds
func (e *RoleAssumeInput) GetDurationSeconds() int32 {
	return e.DurationSeconds
}

// SetDurationSeconds sets the value for the field durationSeconds
func (e *RoleAssumeInput) SetDurationSeconds(durationSeconds int32) {
	e.DurationSeconds = durationSeconds
}

// GetInlinePolicy returns the value for the field inlinePolicy
func (e *RoleAssumeInput) GetInlinePolicy() *InlinePolicy {
	return e.InlinePolicy
}

// SetInlinePolicy sets the value for the field inlinePolicy
func (e *RoleAssumeInput) SetInlinePolicy(inlinePolicy *InlinePolicy) {
	e.InlinePolicy = inlinePolicy
}

// GetName returns the value for the field name
func (e *RoleAssumeInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *RoleAssumeInput) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *RoleAssumeInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleAssumeInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleAssumeInput) InitializeDefaults() {
}

// roleAssumeInputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleAssumeInputAlias RoleAssumeInput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleAssumeInput) UnmarshalJSON(data []byte) error {
	var alias roleAssumeInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleAssumeInput)(&alias)).InitializeDefaults()
	*e = RoleAssumeInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleAssumeInput) MarshalJSON() ([]byte, error) {
	alias := roleAssumeInputAlias(e)
	return json.Marshal(alias)
}

// NewRoleAssumeOutput creates a new RoleAssumeOutput
func NewRoleAssumeOutput() *RoleAssumeOutput {
	s := &RoleAssumeOutput{}
	s.InitializeDefaults()
	return s
}

// RoleAssumeOutput struct
type RoleAssumeOutput struct {
	// credentials for the assumed role
	Credentials *Credentials `json:"credentials,omitempty" yaml:"credentials,omitempty"`
}

// GetCredentials returns the value for the field credentials
func (e *RoleAssumeOutput) GetCredentials() *Credentials {
	return e.Credentials
}

// SetCredentials sets the value for the field credentials
func (e *RoleAssumeOutput) SetCredentials(credentials *Credentials) {
	e.Credentials = credentials
}

// StructPath returns StructPath
func (e *RoleAssumeOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleAssumeOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleAssumeOutput) InitializeDefaults() {
}

// roleAssumeOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleAssumeOutputAlias RoleAssumeOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleAssumeOutput) UnmarshalJSON(data []byte) error {
	var alias roleAssumeOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleAssumeOutput)(&alias)).InitializeDefaults()
	*e = RoleAssumeOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleAssumeOutput) MarshalJSON() ([]byte, error) {
	alias := roleAssumeOutputAlias(e)
	return json.Marshal(alias)
}

// NewRoleIdentityPolicyAttachInput creates a new RoleIdentityPolicyAttachInput
func NewRoleIdentityPolicyAttachInput() *RoleIdentityPolicyAttachInput {
	s := &RoleIdentityPolicyAttachInput{}
	s.InitializeDefaults()
	return s
}

// RoleIdentityPolicyAttachInput struct
type RoleIdentityPolicyAttachInput struct {
	PolicyName string `json:"policyName,omitempty" yaml:"policyName,omitempty"`
	RoleName   string `json:"roleName,omitempty" yaml:"roleName,omitempty"`
}

// GetPolicyName returns the value for the field policyName
func (e *RoleIdentityPolicyAttachInput) GetPolicyName() string {
	return e.PolicyName
}

// SetPolicyName sets the value for the field policyName
func (e *RoleIdentityPolicyAttachInput) SetPolicyName(policyName string) {
	e.PolicyName = policyName
}

// GetRoleName returns the value for the field roleName
func (e *RoleIdentityPolicyAttachInput) GetRoleName() string {
	return e.RoleName
}

// SetRoleName sets the value for the field roleName
func (e *RoleIdentityPolicyAttachInput) SetRoleName(roleName string) {
	e.RoleName = roleName
}

// StructPath returns StructPath
func (e *RoleIdentityPolicyAttachInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleIdentityPolicyAttachInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleIdentityPolicyAttachInput) InitializeDefaults() {
}

// roleIdentityPolicyAttachInputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleIdentityPolicyAttachInputAlias RoleIdentityPolicyAttachInput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleIdentityPolicyAttachInput) UnmarshalJSON(data []byte) error {
	var alias roleIdentityPolicyAttachInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleIdentityPolicyAttachInput)(&alias)).InitializeDefaults()
	*e = RoleIdentityPolicyAttachInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleIdentityPolicyAttachInput) MarshalJSON() ([]byte, error) {
	alias := roleIdentityPolicyAttachInputAlias(e)
	return json.Marshal(alias)
}

// NewRoleIdentityPolicyAttachOutput creates a new RoleIdentityPolicyAttachOutput
func NewRoleIdentityPolicyAttachOutput() *RoleIdentityPolicyAttachOutput {
	s := &RoleIdentityPolicyAttachOutput{}
	s.InitializeDefaults()
	return s
}

// RoleIdentityPolicyAttachOutput struct
type RoleIdentityPolicyAttachOutput struct {
}

// StructPath returns StructPath
func (e *RoleIdentityPolicyAttachOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleIdentityPolicyAttachOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleIdentityPolicyAttachOutput) InitializeDefaults() {
}

// roleIdentityPolicyAttachOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleIdentityPolicyAttachOutputAlias RoleIdentityPolicyAttachOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleIdentityPolicyAttachOutput) UnmarshalJSON(data []byte) error {
	var alias roleIdentityPolicyAttachOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleIdentityPolicyAttachOutput)(&alias)).InitializeDefaults()
	*e = RoleIdentityPolicyAttachOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleIdentityPolicyAttachOutput) MarshalJSON() ([]byte, error) {
	alias := roleIdentityPolicyAttachOutputAlias(e)
	return json.Marshal(alias)
}

// NewRoleIdentityPolicyListInput creates a new RoleIdentityPolicyListInput
func NewRoleIdentityPolicyListInput() *RoleIdentityPolicyListInput {
	s := &RoleIdentityPolicyListInput{}
	s.InitializeDefaults()
	return s
}

// RoleIdentityPolicyListInput struct
type RoleIdentityPolicyListInput struct {
	RoleName string `json:"roleName,omitempty" yaml:"roleName,omitempty"`
}

// GetRoleName returns the value for the field roleName
func (e *RoleIdentityPolicyListInput) GetRoleName() string {
	return e.RoleName
}

// SetRoleName sets the value for the field roleName
func (e *RoleIdentityPolicyListInput) SetRoleName(roleName string) {
	e.RoleName = roleName
}

// StructPath returns StructPath
func (e *RoleIdentityPolicyListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleIdentityPolicyListInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleIdentityPolicyListInput) InitializeDefaults() {
}

// roleIdentityPolicyListInputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleIdentityPolicyListInputAlias RoleIdentityPolicyListInput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleIdentityPolicyListInput) UnmarshalJSON(data []byte) error {
	var alias roleIdentityPolicyListInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleIdentityPolicyListInput)(&alias)).InitializeDefaults()
	*e = RoleIdentityPolicyListInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleIdentityPolicyListInput) MarshalJSON() ([]byte, error) {
	alias := roleIdentityPolicyListInputAlias(e)
	return json.Marshal(alias)
}

// NewRoleIdentityPolicyListOutput creates a new RoleIdentityPolicyListOutput
func NewRoleIdentityPolicyListOutput() *RoleIdentityPolicyListOutput {
	s := &RoleIdentityPolicyListOutput{}
	s.InitializeDefaults()
	return s
}

// RoleIdentityPolicyListOutput struct
type RoleIdentityPolicyListOutput struct {
	Attachments []*IdentityPolicyAttachment `json:"attachments,omitempty" yaml:"attachments,omitempty"`
}

// GetAttachments returns the value for the field attachments
func (e *RoleIdentityPolicyListOutput) GetAttachments() []*IdentityPolicyAttachment {
	return e.Attachments
}

// SetAttachments sets the value for the field attachments
func (e *RoleIdentityPolicyListOutput) SetAttachments(attachments []*IdentityPolicyAttachment) {
	e.Attachments = attachments
}

// StructPath returns StructPath
func (e *RoleIdentityPolicyListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleIdentityPolicyListOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleIdentityPolicyListOutput) InitializeDefaults() {
}

// roleIdentityPolicyListOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleIdentityPolicyListOutputAlias RoleIdentityPolicyListOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleIdentityPolicyListOutput) UnmarshalJSON(data []byte) error {
	var alias roleIdentityPolicyListOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleIdentityPolicyListOutput)(&alias)).InitializeDefaults()
	*e = RoleIdentityPolicyListOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleIdentityPolicyListOutput) MarshalJSON() ([]byte, error) {
	alias := roleIdentityPolicyListOutputAlias(e)
	return json.Marshal(alias)
}

// NewRoleIdentityPolicyDetachInput creates a new RoleIdentityPolicyDetachInput
func NewRoleIdentityPolicyDetachInput() *RoleIdentityPolicyDetachInput {
	s := &RoleIdentityPolicyDetachInput{}
	s.InitializeDefaults()
	return s
}

// RoleIdentityPolicyDetachInput struct
type RoleIdentityPolicyDetachInput struct {
	PolicyName string `json:"policyName,omitempty" yaml:"policyName,omitempty"`
	RoleName   string `json:"roleName,omitempty" yaml:"roleName,omitempty"`
}

// GetPolicyName returns the value for the field policyName
func (e *RoleIdentityPolicyDetachInput) GetPolicyName() string {
	return e.PolicyName
}

// SetPolicyName sets the value for the field policyName
func (e *RoleIdentityPolicyDetachInput) SetPolicyName(policyName string) {
	e.PolicyName = policyName
}

// GetRoleName returns the value for the field roleName
func (e *RoleIdentityPolicyDetachInput) GetRoleName() string {
	return e.RoleName
}

// SetRoleName sets the value for the field roleName
func (e *RoleIdentityPolicyDetachInput) SetRoleName(roleName string) {
	e.RoleName = roleName
}

// StructPath returns StructPath
func (e *RoleIdentityPolicyDetachInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleIdentityPolicyDetachInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleIdentityPolicyDetachInput) InitializeDefaults() {
}

// roleIdentityPolicyDetachInputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleIdentityPolicyDetachInputAlias RoleIdentityPolicyDetachInput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleIdentityPolicyDetachInput) UnmarshalJSON(data []byte) error {
	var alias roleIdentityPolicyDetachInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleIdentityPolicyDetachInput)(&alias)).InitializeDefaults()
	*e = RoleIdentityPolicyDetachInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleIdentityPolicyDetachInput) MarshalJSON() ([]byte, error) {
	alias := roleIdentityPolicyDetachInputAlias(e)
	return json.Marshal(alias)
}

// NewRoleIdentityPolicyDetachOutput creates a new RoleIdentityPolicyDetachOutput
func NewRoleIdentityPolicyDetachOutput() *RoleIdentityPolicyDetachOutput {
	s := &RoleIdentityPolicyDetachOutput{}
	s.InitializeDefaults()
	return s
}

// RoleIdentityPolicyDetachOutput struct
type RoleIdentityPolicyDetachOutput struct {
}

// StructPath returns StructPath
func (e *RoleIdentityPolicyDetachOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleIdentityPolicyDetachOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleIdentityPolicyDetachOutput) InitializeDefaults() {
}

// roleIdentityPolicyDetachOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleIdentityPolicyDetachOutputAlias RoleIdentityPolicyDetachOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleIdentityPolicyDetachOutput) UnmarshalJSON(data []byte) error {
	var alias roleIdentityPolicyDetachOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleIdentityPolicyDetachOutput)(&alias)).InitializeDefaults()
	*e = RoleIdentityPolicyDetachOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleIdentityPolicyDetachOutput) MarshalJSON() ([]byte, error) {
	alias := roleIdentityPolicyDetachOutputAlias(e)
	return json.Marshal(alias)
}

// NewIdentityPolicyCreateInput creates a new IdentityPolicyCreateInput
func NewIdentityPolicyCreateInput() *IdentityPolicyCreateInput {
	s := &IdentityPolicyCreateInput{}
	s.InitializeDefaults()
	return s
}

// IdentityPolicyCreateInput struct
type IdentityPolicyCreateInput struct {
	Name       string                     `json:"name,omitempty" yaml:"name,omitempty"`
	Statements []*IdentityPolicyStatement `json:"statements,omitempty" yaml:"statements,omitempty"`
}

// GetName returns the value for the field name
func (e *IdentityPolicyCreateInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *IdentityPolicyCreateInput) SetName(name string) {
	e.Name = name
}

// GetStatements returns the value for the field statements
func (e *IdentityPolicyCreateInput) GetStatements() []*IdentityPolicyStatement {
	return e.Statements
}

// SetStatements sets the value for the field statements
func (e *IdentityPolicyCreateInput) SetStatements(statements []*IdentityPolicyStatement) {
	e.Statements = statements
}

// StructPath returns StructPath
func (e *IdentityPolicyCreateInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyCreateInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *IdentityPolicyCreateInput) InitializeDefaults() {
}

// identityPolicyCreateInputAlias is defined to help pre and post JSON marshaling without recursive loops
type identityPolicyCreateInputAlias IdentityPolicyCreateInput

// UnmarshalJSON implements json.Unmarshaler
func (e *IdentityPolicyCreateInput) UnmarshalJSON(data []byte) error {
	var alias identityPolicyCreateInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*IdentityPolicyCreateInput)(&alias)).InitializeDefaults()
	*e = IdentityPolicyCreateInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e IdentityPolicyCreateInput) MarshalJSON() ([]byte, error) {
	alias := identityPolicyCreateInputAlias(e)
	return json.Marshal(alias)
}

// NewIdentityPolicyCreateOutput creates a new IdentityPolicyCreateOutput
func NewIdentityPolicyCreateOutput() *IdentityPolicyCreateOutput {
	s := &IdentityPolicyCreateOutput{}
	s.InitializeDefaults()
	return s
}

// IdentityPolicyCreateOutput struct
type IdentityPolicyCreateOutput struct {
	Policy *IdentityPolicy `json:"policy,omitempty" yaml:"policy,omitempty"`
}

// GetPolicy returns the value for the field policy
func (e *IdentityPolicyCreateOutput) GetPolicy() *IdentityPolicy {
	return e.Policy
}

// SetPolicy sets the value for the field policy
func (e *IdentityPolicyCreateOutput) SetPolicy(policy *IdentityPolicy) {
	e.Policy = policy
}

// StructPath returns StructPath
func (e *IdentityPolicyCreateOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyCreateOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *IdentityPolicyCreateOutput) InitializeDefaults() {
}

// identityPolicyCreateOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type identityPolicyCreateOutputAlias IdentityPolicyCreateOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *IdentityPolicyCreateOutput) UnmarshalJSON(data []byte) error {
	var alias identityPolicyCreateOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*IdentityPolicyCreateOutput)(&alias)).InitializeDefaults()
	*e = IdentityPolicyCreateOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e IdentityPolicyCreateOutput) MarshalJSON() ([]byte, error) {
	alias := identityPolicyCreateOutputAlias(e)
	return json.Marshal(alias)
}

// NewIdentityPolicyListInput creates a new IdentityPolicyListInput
func NewIdentityPolicyListInput() *IdentityPolicyListInput {
	s := &IdentityPolicyListInput{}
	s.InitializeDefaults()
	return s
}

// IdentityPolicyListInput struct
type IdentityPolicyListInput struct {
}

// StructPath returns StructPath
func (e *IdentityPolicyListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyListInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *IdentityPolicyListInput) InitializeDefaults() {
}

// identityPolicyListInputAlias is defined to help pre and post JSON marshaling without recursive loops
type identityPolicyListInputAlias IdentityPolicyListInput

// UnmarshalJSON implements json.Unmarshaler
func (e *IdentityPolicyListInput) UnmarshalJSON(data []byte) error {
	var alias identityPolicyListInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*IdentityPolicyListInput)(&alias)).InitializeDefaults()
	*e = IdentityPolicyListInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e IdentityPolicyListInput) MarshalJSON() ([]byte, error) {
	alias := identityPolicyListInputAlias(e)
	return json.Marshal(alias)
}

// NewIdentityPolicyListOutput creates a new IdentityPolicyListOutput
func NewIdentityPolicyListOutput() *IdentityPolicyListOutput {
	s := &IdentityPolicyListOutput{}
	s.InitializeDefaults()
	return s
}

// IdentityPolicyListOutput struct
type IdentityPolicyListOutput struct {
	Policies []*IdentityPolicy `json:"policies,omitempty" yaml:"policies,omitempty"`
}

// GetPolicies returns the value for the field policies
func (e *IdentityPolicyListOutput) GetPolicies() []*IdentityPolicy {
	return e.Policies
}

// SetPolicies sets the value for the field policies
func (e *IdentityPolicyListOutput) SetPolicies(policies []*IdentityPolicy) {
	e.Policies = policies
}

// StructPath returns StructPath
func (e *IdentityPolicyListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyListOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *IdentityPolicyListOutput) InitializeDefaults() {
}

// identityPolicyListOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type identityPolicyListOutputAlias IdentityPolicyListOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *IdentityPolicyListOutput) UnmarshalJSON(data []byte) error {
	var alias identityPolicyListOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*IdentityPolicyListOutput)(&alias)).InitializeDefaults()
	*e = IdentityPolicyListOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e IdentityPolicyListOutput) MarshalJSON() ([]byte, error) {
	alias := identityPolicyListOutputAlias(e)
	return json.Marshal(alias)
}

// NewIdentityPolicyRetrieveInput creates a new IdentityPolicyRetrieveInput
func NewIdentityPolicyRetrieveInput() *IdentityPolicyRetrieveInput {
	s := &IdentityPolicyRetrieveInput{}
	s.InitializeDefaults()
	return s
}

// IdentityPolicyRetrieveInput struct
type IdentityPolicyRetrieveInput struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetName returns the value for the field name
func (e *IdentityPolicyRetrieveInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *IdentityPolicyRetrieveInput) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *IdentityPolicyRetrieveInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyRetrieveInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *IdentityPolicyRetrieveInput) InitializeDefaults() {
}

// identityPolicyRetrieveInputAlias is defined to help pre and post JSON marshaling without recursive loops
type identityPolicyRetrieveInputAlias IdentityPolicyRetrieveInput

// UnmarshalJSON implements json.Unmarshaler
func (e *IdentityPolicyRetrieveInput) UnmarshalJSON(data []byte) error {
	var alias identityPolicyRetrieveInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*IdentityPolicyRetrieveInput)(&alias)).InitializeDefaults()
	*e = IdentityPolicyRetrieveInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e IdentityPolicyRetrieveInput) MarshalJSON() ([]byte, error) {
	alias := identityPolicyRetrieveInputAlias(e)
	return json.Marshal(alias)
}

// NewIdentityPolicyRetrieveOutput creates a new IdentityPolicyRetrieveOutput
func NewIdentityPolicyRetrieveOutput() *IdentityPolicyRetrieveOutput {
	s := &IdentityPolicyRetrieveOutput{}
	s.InitializeDefaults()
	return s
}

// IdentityPolicyRetrieveOutput struct
type IdentityPolicyRetrieveOutput struct {
	Policy *IdentityPolicy `json:"policy,omitempty" yaml:"policy,omitempty"`
}

// GetPolicy returns the value for the field policy
func (e *IdentityPolicyRetrieveOutput) GetPolicy() *IdentityPolicy {
	return e.Policy
}

// SetPolicy sets the value for the field policy
func (e *IdentityPolicyRetrieveOutput) SetPolicy(policy *IdentityPolicy) {
	e.Policy = policy
}

// StructPath returns StructPath
func (e *IdentityPolicyRetrieveOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyRetrieveOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *IdentityPolicyRetrieveOutput) InitializeDefaults() {
}

// identityPolicyRetrieveOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type identityPolicyRetrieveOutputAlias IdentityPolicyRetrieveOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *IdentityPolicyRetrieveOutput) UnmarshalJSON(data []byte) error {
	var alias identityPolicyRetrieveOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*IdentityPolicyRetrieveOutput)(&alias)).InitializeDefaults()
	*e = IdentityPolicyRetrieveOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e IdentityPolicyRetrieveOutput) MarshalJSON() ([]byte, error) {
	alias := identityPolicyRetrieveOutputAlias(e)
	return json.Marshal(alias)
}

// NewIdentityPolicyDestroyInput creates a new IdentityPolicyDestroyInput
func NewIdentityPolicyDestroyInput() *IdentityPolicyDestroyInput {
	s := &IdentityPolicyDestroyInput{}
	s.InitializeDefaults()
	return s
}

// IdentityPolicyDestroyInput struct
type IdentityPolicyDestroyInput struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetName returns the value for the field name
func (e *IdentityPolicyDestroyInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *IdentityPolicyDestroyInput) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *IdentityPolicyDestroyInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyDestroyInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *IdentityPolicyDestroyInput) InitializeDefaults() {
}

// identityPolicyDestroyInputAlias is defined to help pre and post JSON marshaling without recursive loops
type identityPolicyDestroyInputAlias IdentityPolicyDestroyInput

// UnmarshalJSON implements json.Unmarshaler
func (e *IdentityPolicyDestroyInput) UnmarshalJSON(data []byte) error {
	var alias identityPolicyDestroyInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*IdentityPolicyDestroyInput)(&alias)).InitializeDefaults()
	*e = IdentityPolicyDestroyInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e IdentityPolicyDestroyInput) MarshalJSON() ([]byte, error) {
	alias := identityPolicyDestroyInputAlias(e)
	return json.Marshal(alias)
}

// NewIdentityPolicyDestroyOutput creates a new IdentityPolicyDestroyOutput
func NewIdentityPolicyDestroyOutput() *IdentityPolicyDestroyOutput {
	s := &IdentityPolicyDestroyOutput{}
	s.InitializeDefaults()
	return s
}

// IdentityPolicyDestroyOutput struct
type IdentityPolicyDestroyOutput struct {
}

// StructPath returns StructPath
func (e *IdentityPolicyDestroyOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyDestroyOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *IdentityPolicyDestroyOutput) InitializeDefaults() {
}

// identityPolicyDestroyOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type identityPolicyDestroyOutputAlias IdentityPolicyDestroyOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *IdentityPolicyDestroyOutput) UnmarshalJSON(data []byte) error {
	var alias identityPolicyDestroyOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*IdentityPolicyDestroyOutput)(&alias)).InitializeDefaults()
	*e = IdentityPolicyDestroyOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e IdentityPolicyDestroyOutput) MarshalJSON() ([]byte, error) {
	alias := identityPolicyDestroyOutputAlias(e)
	return json.Marshal(alias)
}

// NewIdentityPolicyUpdateInput creates a new IdentityPolicyUpdateInput
func NewIdentityPolicyUpdateInput() *IdentityPolicyUpdateInput {
	s := &IdentityPolicyUpdateInput{}
	s.InitializeDefaults()
	return s
}

// IdentityPolicyUpdateInput struct
type IdentityPolicyUpdateInput struct {
	Name       string                     `json:"name,omitempty" yaml:"name,omitempty"`
	Statements []*IdentityPolicyStatement `json:"statements,omitempty" yaml:"statements,omitempty"`
}

// GetName returns the value for the field name
func (e *IdentityPolicyUpdateInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *IdentityPolicyUpdateInput) SetName(name string) {
	e.Name = name
}

// GetStatements returns the value for the field statements
func (e *IdentityPolicyUpdateInput) GetStatements() []*IdentityPolicyStatement {
	return e.Statements
}

// SetStatements sets the value for the field statements
func (e *IdentityPolicyUpdateInput) SetStatements(statements []*IdentityPolicyStatement) {
	e.Statements = statements
}

// StructPath returns StructPath
func (e *IdentityPolicyUpdateInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyUpdateInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *IdentityPolicyUpdateInput) InitializeDefaults() {
}

// identityPolicyUpdateInputAlias is defined to help pre and post JSON marshaling without recursive loops
type identityPolicyUpdateInputAlias IdentityPolicyUpdateInput

// UnmarshalJSON implements json.Unmarshaler
func (e *IdentityPolicyUpdateInput) UnmarshalJSON(data []byte) error {
	var alias identityPolicyUpdateInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*IdentityPolicyUpdateInput)(&alias)).InitializeDefaults()
	*e = IdentityPolicyUpdateInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e IdentityPolicyUpdateInput) MarshalJSON() ([]byte, error) {
	alias := identityPolicyUpdateInputAlias(e)
	return json.Marshal(alias)
}

// NewIdentityPolicyUpdateOutput creates a new IdentityPolicyUpdateOutput
func NewIdentityPolicyUpdateOutput() *IdentityPolicyUpdateOutput {
	s := &IdentityPolicyUpdateOutput{}
	s.InitializeDefaults()
	return s
}

// IdentityPolicyUpdateOutput struct
type IdentityPolicyUpdateOutput struct {
	Policy *IdentityPolicy `json:"policy,omitempty" yaml:"policy,omitempty"`
}

// GetPolicy returns the value for the field policy
func (e *IdentityPolicyUpdateOutput) GetPolicy() *IdentityPolicy {
	return e.Policy
}

// SetPolicy sets the value for the field policy
func (e *IdentityPolicyUpdateOutput) SetPolicy(policy *IdentityPolicy) {
	e.Policy = policy
}

// StructPath returns StructPath
func (e *IdentityPolicyUpdateOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyUpdateOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *IdentityPolicyUpdateOutput) InitializeDefaults() {
}

// identityPolicyUpdateOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type identityPolicyUpdateOutputAlias IdentityPolicyUpdateOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *IdentityPolicyUpdateOutput) UnmarshalJSON(data []byte) error {
	var alias identityPolicyUpdateOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*IdentityPolicyUpdateOutput)(&alias)).InitializeDefaults()
	*e = IdentityPolicyUpdateOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e IdentityPolicyUpdateOutput) MarshalJSON() ([]byte, error) {
	alias := identityPolicyUpdateOutputAlias(e)
	return json.Marshal(alias)
}

// NewServiceBearerToken creates a new ServiceBearerToken
func NewServiceBearerToken() *ServiceBearerToken {
	s := &ServiceBearerToken{}
	s.InitializeDefaults()
	return s
}

// ServiceBearerToken - Token you can use to perform operations on your behalf
type ServiceBearerToken struct {
	ExpiresAt time.Time `json:"expiresAt,omitempty" yaml:"expiresAt,omitempty"`
	Token     string    `json:"token,omitempty" yaml:"token,omitempty"`
}

// GetExpiresAt returns the value for the field expiresAt
func (e *ServiceBearerToken) GetExpiresAt() time.Time {
	return e.ExpiresAt
}

// SetExpiresAt sets the value for the field expiresAt
func (e *ServiceBearerToken) SetExpiresAt(expiresAt time.Time) {
	e.ExpiresAt = expiresAt
}

// GetToken returns the value for the field token
func (e *ServiceBearerToken) GetToken() string {
	return e.Token
}

// SetToken sets the value for the field token
func (e *ServiceBearerToken) SetToken(token string) {
	e.Token = token
}

// StructPath returns StructPath
func (e *ServiceBearerToken) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathServiceBearerToken.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *ServiceBearerToken) InitializeDefaults() {
}

// serviceBearerTokenAlias is defined to help pre and post JSON marshaling without recursive loops
type serviceBearerTokenAlias ServiceBearerToken

// UnmarshalJSON implements json.Unmarshaler
func (e *ServiceBearerToken) UnmarshalJSON(data []byte) error {
	var alias serviceBearerTokenAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*ServiceBearerToken)(&alias)).InitializeDefaults()
	*e = ServiceBearerToken(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e ServiceBearerToken) MarshalJSON() ([]byte, error) {
	alias := serviceBearerTokenAlias(e)
	return json.Marshal(alias)
}

// NewInvalidServiceBearerTokenDurationProblem creates a new InvalidServiceBearerTokenDurationProblem
func NewInvalidServiceBearerTokenDurationProblem() *InvalidServiceBearerTokenDurationProblem {
	s := &InvalidServiceBearerTokenDurationProblem{}
	s.InitializeDefaults()
	return s
}

// InvalidServiceBearerTokenDurationProblem struct
type InvalidServiceBearerTokenDurationProblem struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InvalidServiceBearerTokenDurationProblem) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InvalidServiceBearerTokenDurationProblem]
func (e *InvalidServiceBearerTokenDurationProblem) Is(err error) bool {
	_, ok := err.(*InvalidServiceBearerTokenDurationProblem)
	return ok
}

// IsInvalidServiceBearerTokenDurationProblem indicates whether the given error chain contains an error of type [InvalidServiceBearerTokenDurationProblem]
func IsInvalidServiceBearerTokenDurationProblem(err error) bool {
	return errors.Is(err, &InvalidServiceBearerTokenDurationProblem{})
}

// GetMessage returns the value for the field message
func (e *InvalidServiceBearerTokenDurationProblem) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InvalidServiceBearerTokenDurationProblem) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InvalidServiceBearerTokenDurationProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvalidServiceBearerTokenDurationProblem.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvalidServiceBearerTokenDurationProblem) InitializeDefaults() {
}

// invalidServiceBearerTokenDurationProblemAlias is defined to help pre and post JSON marshaling without recursive loops
type invalidServiceBearerTokenDurationProblemAlias InvalidServiceBearerTokenDurationProblem

// UnmarshalJSON implements json.Unmarshaler
func (e *InvalidServiceBearerTokenDurationProblem) UnmarshalJSON(data []byte) error {
	var alias invalidServiceBearerTokenDurationProblemAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvalidServiceBearerTokenDurationProblem)(&alias)).InitializeDefaults()
	*e = InvalidServiceBearerTokenDurationProblem(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvalidServiceBearerTokenDurationProblem) MarshalJSON() ([]byte, error) {
	alias := invalidServiceBearerTokenDurationProblemAlias(e)
	return json.Marshal(alias)
}

// NewInvalidServiceNameProblem creates a new InvalidServiceNameProblem
func NewInvalidServiceNameProblem() *InvalidServiceNameProblem {
	s := &InvalidServiceNameProblem{}
	s.InitializeDefaults()
	return s
}

// InvalidServiceNameProblem struct
type InvalidServiceNameProblem struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InvalidServiceNameProblem) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InvalidServiceNameProblem]
func (e *InvalidServiceNameProblem) Is(err error) bool {
	_, ok := err.(*InvalidServiceNameProblem)
	return ok
}

// IsInvalidServiceNameProblem indicates whether the given error chain contains an error of type [InvalidServiceNameProblem]
func IsInvalidServiceNameProblem(err error) bool {
	return errors.Is(err, &InvalidServiceNameProblem{})
}

// GetMessage returns the value for the field message
func (e *InvalidServiceNameProblem) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InvalidServiceNameProblem) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InvalidServiceNameProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvalidServiceNameProblem.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvalidServiceNameProblem) InitializeDefaults() {
}

// invalidServiceNameProblemAlias is defined to help pre and post JSON marshaling without recursive loops
type invalidServiceNameProblemAlias InvalidServiceNameProblem

// UnmarshalJSON implements json.Unmarshaler
func (e *InvalidServiceNameProblem) UnmarshalJSON(data []byte) error {
	var alias invalidServiceNameProblemAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvalidServiceNameProblem)(&alias)).InitializeDefaults()
	*e = InvalidServiceNameProblem(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvalidServiceNameProblem) MarshalJSON() ([]byte, error) {
	alias := invalidServiceNameProblemAlias(e)
	return json.Marshal(alias)
}

// NewServiceBearerTokenGetInput creates a new ServiceBearerTokenGetInput
func NewServiceBearerTokenGetInput() *ServiceBearerTokenGetInput {
	s := &ServiceBearerTokenGetInput{}
	s.InitializeDefaults()
	return s
}

// ServiceBearerTokenGetInput struct
type ServiceBearerTokenGetInput struct {
	// duration in seconds for the token to be active
	// the minimum value is 900 seconds (15 minutes)
	// the maximum value is 43200 seconds (12 hours)
	DurationSeconds int32  `json:"durationSeconds,omitempty" yaml:"durationSeconds,omitempty"`
	ServiceName     string `json:"serviceName,omitempty" yaml:"serviceName,omitempty"`
}

// GetDurationSeconds returns the value for the field durationSeconds
func (e *ServiceBearerTokenGetInput) GetDurationSeconds() int32 {
	return e.DurationSeconds
}

// SetDurationSeconds sets the value for the field durationSeconds
func (e *ServiceBearerTokenGetInput) SetDurationSeconds(durationSeconds int32) {
	e.DurationSeconds = durationSeconds
}

// GetServiceName returns the value for the field serviceName
func (e *ServiceBearerTokenGetInput) GetServiceName() string {
	return e.ServiceName
}

// SetServiceName sets the value for the field serviceName
func (e *ServiceBearerTokenGetInput) SetServiceName(serviceName string) {
	e.ServiceName = serviceName
}

// StructPath returns StructPath
func (e *ServiceBearerTokenGetInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathServiceBearerTokenGetInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *ServiceBearerTokenGetInput) InitializeDefaults() {
}

// serviceBearerTokenGetInputAlias is defined to help pre and post JSON marshaling without recursive loops
type serviceBearerTokenGetInputAlias ServiceBearerTokenGetInput

// UnmarshalJSON implements json.Unmarshaler
func (e *ServiceBearerTokenGetInput) UnmarshalJSON(data []byte) error {
	var alias serviceBearerTokenGetInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*ServiceBearerTokenGetInput)(&alias)).InitializeDefaults()
	*e = ServiceBearerTokenGetInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e ServiceBearerTokenGetInput) MarshalJSON() ([]byte, error) {
	alias := serviceBearerTokenGetInputAlias(e)
	return json.Marshal(alias)
}

// NewServiceBearerTokenGetOutput creates a new ServiceBearerTokenGetOutput
func NewServiceBearerTokenGetOutput() *ServiceBearerTokenGetOutput {
	s := &ServiceBearerTokenGetOutput{}
	s.InitializeDefaults()
	return s
}

// ServiceBearerTokenGetOutput struct
type ServiceBearerTokenGetOutput struct {
	Token *ServiceBearerToken `json:"token,omitempty" yaml:"token,omitempty"`
}

// GetToken returns the value for the field token
func (e *ServiceBearerTokenGetOutput) GetToken() *ServiceBearerToken {
	return e.Token
}

// SetToken sets the value for the field token
func (e *ServiceBearerTokenGetOutput) SetToken(token *ServiceBearerToken) {
	e.Token = token
}

// StructPath returns StructPath
func (e *ServiceBearerTokenGetOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathServiceBearerTokenGetOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *ServiceBearerTokenGetOutput) InitializeDefaults() {
}

// serviceBearerTokenGetOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type serviceBearerTokenGetOutputAlias ServiceBearerTokenGetOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *ServiceBearerTokenGetOutput) UnmarshalJSON(data []byte) error {
	var alias serviceBearerTokenGetOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*ServiceBearerTokenGetOutput)(&alias)).InitializeDefaults()
	*e = ServiceBearerTokenGetOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e ServiceBearerTokenGetOutput) MarshalJSON() ([]byte, error) {
	alias := serviceBearerTokenGetOutputAlias(e)
	return json.Marshal(alias)
}

// NewInternalAccessKeyUser creates a new InternalAccessKeyUser
func NewInternalAccessKeyUser() *InternalAccessKeyUser {
	s := &InternalAccessKeyUser{}
	s.InitializeDefaults()
	return s
}

// InternalAccessKeyUser struct
type InternalAccessKeyUser struct {
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}

// GetUsername returns the value for the field username
func (e *InternalAccessKeyUser) GetUsername() string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *InternalAccessKeyUser) SetUsername(username string) {
	e.Username = username
}

// StructPath returns StructPath
func (e *InternalAccessKeyUser) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalAccessKeyUser.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InternalAccessKeyUser) InitializeDefaults() {
}

// internalAccessKeyUserAlias is defined to help pre and post JSON marshaling without recursive loops
type internalAccessKeyUserAlias InternalAccessKeyUser

// UnmarshalJSON implements json.Unmarshaler
func (e *InternalAccessKeyUser) UnmarshalJSON(data []byte) error {
	var alias internalAccessKeyUserAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InternalAccessKeyUser)(&alias)).InitializeDefaults()
	*e = InternalAccessKeyUser(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InternalAccessKeyUser) MarshalJSON() ([]byte, error) {
	alias := internalAccessKeyUserAlias(e)
	return json.Marshal(alias)
}

// NewInternalAccessKeyRole creates a new InternalAccessKeyRole
func NewInternalAccessKeyRole() *InternalAccessKeyRole {
	s := &InternalAccessKeyRole{}
	s.InitializeDefaults()
	return s
}

// InternalAccessKeyRole struct
type InternalAccessKeyRole struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetName returns the value for the field name
func (e *InternalAccessKeyRole) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *InternalAccessKeyRole) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *InternalAccessKeyRole) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalAccessKeyRole.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InternalAccessKeyRole) InitializeDefaults() {
}

// internalAccessKeyRoleAlias is defined to help pre and post JSON marshaling without recursive loops
type internalAccessKeyRoleAlias InternalAccessKeyRole

// UnmarshalJSON implements json.Unmarshaler
func (e *InternalAccessKeyRole) UnmarshalJSON(data []byte) error {
	var alias internalAccessKeyRoleAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InternalAccessKeyRole)(&alias)).InitializeDefaults()
	*e = InternalAccessKeyRole(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InternalAccessKeyRole) MarshalJSON() ([]byte, error) {
	alias := internalAccessKeyRoleAlias(e)
	return json.Marshal(alias)
}

// NewInternalAccessKeyAccount creates a new InternalAccessKeyAccount
func NewInternalAccessKeyAccount() *InternalAccessKeyAccount {
	s := &InternalAccessKeyAccount{}
	s.InitializeDefaults()
	return s
}

// InternalAccessKeyAccount struct
type InternalAccessKeyAccount struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetName returns the value for the field name
func (e *InternalAccessKeyAccount) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *InternalAccessKeyAccount) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *InternalAccessKeyAccount) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalAccessKeyAccount.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InternalAccessKeyAccount) InitializeDefaults() {
}

// internalAccessKeyAccountAlias is defined to help pre and post JSON marshaling without recursive loops
type internalAccessKeyAccountAlias InternalAccessKeyAccount

// UnmarshalJSON implements json.Unmarshaler
func (e *InternalAccessKeyAccount) UnmarshalJSON(data []byte) error {
	var alias internalAccessKeyAccountAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InternalAccessKeyAccount)(&alias)).InitializeDefaults()
	*e = InternalAccessKeyAccount(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InternalAccessKeyAccount) MarshalJSON() ([]byte, error) {
	alias := internalAccessKeyAccountAlias(e)
	return json.Marshal(alias)
}

// NewInternalAccessKey creates a new InternalAccessKey
func NewInternalAccessKey() *InternalAccessKey {
	s := &InternalAccessKey{}
	s.InitializeDefaults()
	return s
}

// InternalAccessKey struct
type InternalAccessKey struct {
	AccessKeyID     string                    `json:"accessKeyID,omitempty" yaml:"accessKeyID,omitempty"`
	Account         *InternalAccessKeyAccount `json:"account,omitempty" yaml:"account,omitempty"`
	Role            *InternalAccessKeyRole    `json:"role,omitempty" yaml:"role,omitempty"`
	SecretAccessKey string                    `json:"secretAccessKey,omitempty" yaml:"secretAccessKey,omitempty"`
	User            *InternalAccessKeyUser    `json:"user,omitempty" yaml:"user,omitempty"`
}

// GetAccessKeyID returns the value for the field accessKeyID
func (e *InternalAccessKey) GetAccessKeyID() string {
	return e.AccessKeyID
}

// SetAccessKeyID sets the value for the field accessKeyID
func (e *InternalAccessKey) SetAccessKeyID(accessKeyID string) {
	e.AccessKeyID = accessKeyID
}

// GetAccount returns the value for the field account
func (e *InternalAccessKey) GetAccount() *InternalAccessKeyAccount {
	return e.Account
}

// SetAccount sets the value for the field account
func (e *InternalAccessKey) SetAccount(account *InternalAccessKeyAccount) {
	e.Account = account
}

// GetRole returns the value for the field role
func (e *InternalAccessKey) GetRole() *InternalAccessKeyRole {
	return e.Role
}

// SetRole sets the value for the field role
func (e *InternalAccessKey) SetRole(role *InternalAccessKeyRole) {
	e.Role = role
}

// GetSecretAccessKey returns the value for the field secretAccessKey
func (e *InternalAccessKey) GetSecretAccessKey() string {
	return e.SecretAccessKey
}

// SetSecretAccessKey sets the value for the field secretAccessKey
func (e *InternalAccessKey) SetSecretAccessKey(secretAccessKey string) {
	e.SecretAccessKey = secretAccessKey
}

// GetUser returns the value for the field user
func (e *InternalAccessKey) GetUser() *InternalAccessKeyUser {
	return e.User
}

// SetUser sets the value for the field user
func (e *InternalAccessKey) SetUser(user *InternalAccessKeyUser) {
	e.User = user
}

// StructPath returns StructPath
func (e *InternalAccessKey) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalAccessKey.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InternalAccessKey) InitializeDefaults() {
}

// internalAccessKeyAlias is defined to help pre and post JSON marshaling without recursive loops
type internalAccessKeyAlias InternalAccessKey

// UnmarshalJSON implements json.Unmarshaler
func (e *InternalAccessKey) UnmarshalJSON(data []byte) error {
	var alias internalAccessKeyAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InternalAccessKey)(&alias)).InitializeDefaults()
	*e = InternalAccessKey(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InternalAccessKey) MarshalJSON() ([]byte, error) {
	alias := internalAccessKeyAlias(e)
	return json.Marshal(alias)
}

// NewInternalQueryAssertionEntryValue creates a new InternalQueryAssertionEntryValue
func NewInternalQueryAssertionEntryValue() *InternalQueryAssertionEntryValue {
	s := &InternalQueryAssertionEntryValue{}
	s.InitializeDefaults()
	return s
}

// InternalQueryAssertionEntryValue struct
type InternalQueryAssertionEntryValue struct {
	B *bool   `json:"b,omitempty" yaml:"b,omitempty"`
	S *string `json:"s,omitempty" yaml:"s,omitempty"`
}

// GetB returns the value for the field b
func (e *InternalQueryAssertionEntryValue) GetB() *bool {
	return e.B
}

// SetB sets the value for the field b
func (e *InternalQueryAssertionEntryValue) SetB(b *bool) {
	e.B = b
}

// GetS returns the value for the field s
func (e *InternalQueryAssertionEntryValue) GetS() *string {
	return e.S
}

// SetS sets the value for the field s
func (e *InternalQueryAssertionEntryValue) SetS(s *string) {
	e.S = s
}

// StructPath returns StructPath
func (e *InternalQueryAssertionEntryValue) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalQueryAssertionEntryValue.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InternalQueryAssertionEntryValue) InitializeDefaults() {
}

// internalQueryAssertionEntryValueAlias is defined to help pre and post JSON marshaling without recursive loops
type internalQueryAssertionEntryValueAlias InternalQueryAssertionEntryValue

// UnmarshalJSON implements json.Unmarshaler
func (e *InternalQueryAssertionEntryValue) UnmarshalJSON(data []byte) error {
	var alias internalQueryAssertionEntryValueAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InternalQueryAssertionEntryValue)(&alias)).InitializeDefaults()
	*e = InternalQueryAssertionEntryValue(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InternalQueryAssertionEntryValue) MarshalJSON() ([]byte, error) {
	alias := internalQueryAssertionEntryValueAlias(e)
	return json.Marshal(alias)
}

// NewInternalQueryAssertionEntry creates a new InternalQueryAssertionEntry
func NewInternalQueryAssertionEntry() *InternalQueryAssertionEntry {
	s := &InternalQueryAssertionEntry{}
	s.InitializeDefaults()
	return s
}

// InternalQueryAssertionEntry struct
type InternalQueryAssertionEntry struct {
	Snippet *string                           `json:"snippet,omitempty" yaml:"snippet,omitempty"`
	Value   *InternalQueryAssertionEntryValue `json:"value,omitempty" yaml:"value,omitempty"`
}

// GetSnippet returns the value for the field snippet
func (e *InternalQueryAssertionEntry) GetSnippet() *string {
	return e.Snippet
}

// SetSnippet sets the value for the field snippet
func (e *InternalQueryAssertionEntry) SetSnippet(snippet *string) {
	e.Snippet = snippet
}

// GetValue returns the value for the field value
func (e *InternalQueryAssertionEntry) GetValue() *InternalQueryAssertionEntryValue {
	return e.Value
}

// SetValue sets the value for the field value
func (e *InternalQueryAssertionEntry) SetValue(value *InternalQueryAssertionEntryValue) {
	e.Value = value
}

// StructPath returns StructPath
func (e *InternalQueryAssertionEntry) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalQueryAssertionEntry.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InternalQueryAssertionEntry) InitializeDefaults() {
}

// internalQueryAssertionEntryAlias is defined to help pre and post JSON marshaling without recursive loops
type internalQueryAssertionEntryAlias InternalQueryAssertionEntry

// UnmarshalJSON implements json.Unmarshaler
func (e *InternalQueryAssertionEntry) UnmarshalJSON(data []byte) error {
	var alias internalQueryAssertionEntryAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InternalQueryAssertionEntry)(&alias)).InitializeDefaults()
	*e = InternalQueryAssertionEntry(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InternalQueryAssertionEntry) MarshalJSON() ([]byte, error) {
	alias := internalQueryAssertionEntryAlias(e)
	return json.Marshal(alias)
}

// NewInternalServicesAssertActionCallerForbiddenProblem creates a new InternalServicesAssertActionCallerForbiddenProblem
func NewInternalServicesAssertActionCallerForbiddenProblem() *InternalServicesAssertActionCallerForbiddenProblem {
	s := &InternalServicesAssertActionCallerForbiddenProblem{}
	s.InitializeDefaults()
	return s
}

// InternalServicesAssertActionCallerForbiddenProblem struct
type InternalServicesAssertActionCallerForbiddenProblem struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InternalServicesAssertActionCallerForbiddenProblem) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InternalServicesAssertActionCallerForbiddenProblem]
func (e *InternalServicesAssertActionCallerForbiddenProblem) Is(err error) bool {
	_, ok := err.(*InternalServicesAssertActionCallerForbiddenProblem)
	return ok
}

// IsInternalServicesAssertActionCallerForbiddenProblem indicates whether the given error chain contains an error of type [InternalServicesAssertActionCallerForbiddenProblem]
func IsInternalServicesAssertActionCallerForbiddenProblem(err error) bool {
	return errors.Is(err, &InternalServicesAssertActionCallerForbiddenProblem{})
}

// GetMessage returns the value for the field message
func (e *InternalServicesAssertActionCallerForbiddenProblem) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InternalServicesAssertActionCallerForbiddenProblem) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InternalServicesAssertActionCallerForbiddenProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalServicesAssertActionCallerForbiddenProblem.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InternalServicesAssertActionCallerForbiddenProblem) InitializeDefaults() {
}

// internalServicesAssertActionCallerForbiddenProblemAlias is defined to help pre and post JSON marshaling without recursive loops
type internalServicesAssertActionCallerForbiddenProblemAlias InternalServicesAssertActionCallerForbiddenProblem

// UnmarshalJSON implements json.Unmarshaler
func (e *InternalServicesAssertActionCallerForbiddenProblem) UnmarshalJSON(data []byte) error {
	var alias internalServicesAssertActionCallerForbiddenProblemAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InternalServicesAssertActionCallerForbiddenProblem)(&alias)).InitializeDefaults()
	*e = InternalServicesAssertActionCallerForbiddenProblem(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InternalServicesAssertActionCallerForbiddenProblem) MarshalJSON() ([]byte, error) {
	alias := internalServicesAssertActionCallerForbiddenProblemAlias(e)
	return json.Marshal(alias)
}

// NewInternalAssertActionQualifier creates a new InternalAssertActionQualifier
func NewInternalAssertActionQualifier() *InternalAssertActionQualifier {
	s := &InternalAssertActionQualifier{}
	s.InitializeDefaults()
	return s
}

// InternalAssertActionQualifier struct
type InternalAssertActionQualifier struct {
	// optional child qualifier, example: if the resource to assert is "s3:Bucket(myimages).Object(myimage.jpg)", the child would be {"resourceName": "Object", "value": "myimage.jpg}
	Child *InternalAssertActionQualifier `json:"child,omitempty" yaml:"child,omitempty"`
	// name of the function to use, usually a resource, example: if the resource to assert is "s3:Bucket(myimages)", the name would be "Bucket"
	Function string `json:"function,omitempty" yaml:"function,omitempty"`
	// value to match the qualifier, example: if the resource to assert is "s3:Bucket(myimages)", the value would be "myimages"
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}

// GetChild returns the value for the field child
func (e *InternalAssertActionQualifier) GetChild() *InternalAssertActionQualifier {
	return e.Child
}

// SetChild sets the value for the field child
func (e *InternalAssertActionQualifier) SetChild(child *InternalAssertActionQualifier) {
	e.Child = child
}

// GetFunction returns the value for the field function
func (e *InternalAssertActionQualifier) GetFunction() string {
	return e.Function
}

// SetFunction sets the value for the field function
func (e *InternalAssertActionQualifier) SetFunction(function string) {
	e.Function = function
}

// GetValue returns the value for the field value
func (e *InternalAssertActionQualifier) GetValue() string {
	return e.Value
}

// SetValue sets the value for the field value
func (e *InternalAssertActionQualifier) SetValue(value string) {
	e.Value = value
}

// StructPath returns StructPath
func (e *InternalAssertActionQualifier) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalAssertActionQualifier.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InternalAssertActionQualifier) InitializeDefaults() {
}

// internalAssertActionQualifierAlias is defined to help pre and post JSON marshaling without recursive loops
type internalAssertActionQualifierAlias InternalAssertActionQualifier

// UnmarshalJSON implements json.Unmarshaler
func (e *InternalAssertActionQualifier) UnmarshalJSON(data []byte) error {
	var alias internalAssertActionQualifierAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InternalAssertActionQualifier)(&alias)).InitializeDefaults()
	*e = InternalAssertActionQualifier(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InternalAssertActionQualifier) MarshalJSON() ([]byte, error) {
	alias := internalAssertActionQualifierAlias(e)
	return json.Marshal(alias)
}

// NewInternalServicesValidateAccessKeyInput creates a new InternalServicesValidateAccessKeyInput
func NewInternalServicesValidateAccessKeyInput() *InternalServicesValidateAccessKeyInput {
	s := &InternalServicesValidateAccessKeyInput{}
	s.InitializeDefaults()
	return s
}

// InternalServicesValidateAccessKeyInput struct
type InternalServicesValidateAccessKeyInput struct {
	RequestAccessKeyID string `json:"requestAccessKeyID,omitempty" yaml:"requestAccessKeyID,omitempty"`
}

// GetRequestAccessKeyID returns the value for the field requestAccessKeyID
func (e *InternalServicesValidateAccessKeyInput) GetRequestAccessKeyID() string {
	return e.RequestAccessKeyID
}

// SetRequestAccessKeyID sets the value for the field requestAccessKeyID
func (e *InternalServicesValidateAccessKeyInput) SetRequestAccessKeyID(requestAccessKeyID string) {
	e.RequestAccessKeyID = requestAccessKeyID
}

// StructPath returns StructPath
func (e *InternalServicesValidateAccessKeyInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalServicesValidateAccessKeyInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InternalServicesValidateAccessKeyInput) InitializeDefaults() {
}

// internalServicesValidateAccessKeyInputAlias is defined to help pre and post JSON marshaling without recursive loops
type internalServicesValidateAccessKeyInputAlias InternalServicesValidateAccessKeyInput

// UnmarshalJSON implements json.Unmarshaler
func (e *InternalServicesValidateAccessKeyInput) UnmarshalJSON(data []byte) error {
	var alias internalServicesValidateAccessKeyInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InternalServicesValidateAccessKeyInput)(&alias)).InitializeDefaults()
	*e = InternalServicesValidateAccessKeyInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InternalServicesValidateAccessKeyInput) MarshalJSON() ([]byte, error) {
	alias := internalServicesValidateAccessKeyInputAlias(e)
	return json.Marshal(alias)
}

// NewInternalServicesValidateAccessKeyOutput creates a new InternalServicesValidateAccessKeyOutput
func NewInternalServicesValidateAccessKeyOutput() *InternalServicesValidateAccessKeyOutput {
	s := &InternalServicesValidateAccessKeyOutput{}
	s.InitializeDefaults()
	return s
}

// InternalServicesValidateAccessKeyOutput struct
type InternalServicesValidateAccessKeyOutput struct {
	Key *InternalAccessKey `json:"key,omitempty" yaml:"key,omitempty"`
}

// GetKey returns the value for the field key
func (e *InternalServicesValidateAccessKeyOutput) GetKey() *InternalAccessKey {
	return e.Key
}

// SetKey sets the value for the field key
func (e *InternalServicesValidateAccessKeyOutput) SetKey(key *InternalAccessKey) {
	e.Key = key
}

// StructPath returns StructPath
func (e *InternalServicesValidateAccessKeyOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalServicesValidateAccessKeyOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InternalServicesValidateAccessKeyOutput) InitializeDefaults() {
}

// internalServicesValidateAccessKeyOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type internalServicesValidateAccessKeyOutputAlias InternalServicesValidateAccessKeyOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *InternalServicesValidateAccessKeyOutput) UnmarshalJSON(data []byte) error {
	var alias internalServicesValidateAccessKeyOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InternalServicesValidateAccessKeyOutput)(&alias)).InitializeDefaults()
	*e = InternalServicesValidateAccessKeyOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InternalServicesValidateAccessKeyOutput) MarshalJSON() ([]byte, error) {
	alias := internalServicesValidateAccessKeyOutputAlias(e)
	return json.Marshal(alias)
}

// NewInternalServicesValidateAccessKeyInvalidAccessKeyProblem creates a new InternalServicesValidateAccessKeyInvalidAccessKeyProblem
func NewInternalServicesValidateAccessKeyInvalidAccessKeyProblem() *InternalServicesValidateAccessKeyInvalidAccessKeyProblem {
	s := &InternalServicesValidateAccessKeyInvalidAccessKeyProblem{}
	s.InitializeDefaults()
	return s
}

// InternalServicesValidateAccessKeyInvalidAccessKeyProblem struct
type InternalServicesValidateAccessKeyInvalidAccessKeyProblem struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InternalServicesValidateAccessKeyInvalidAccessKeyProblem) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InternalServicesValidateAccessKeyInvalidAccessKeyProblem]
func (e *InternalServicesValidateAccessKeyInvalidAccessKeyProblem) Is(err error) bool {
	_, ok := err.(*InternalServicesValidateAccessKeyInvalidAccessKeyProblem)
	return ok
}

// IsInternalServicesValidateAccessKeyInvalidAccessKeyProblem indicates whether the given error chain contains an error of type [InternalServicesValidateAccessKeyInvalidAccessKeyProblem]
func IsInternalServicesValidateAccessKeyInvalidAccessKeyProblem(err error) bool {
	return errors.Is(err, &InternalServicesValidateAccessKeyInvalidAccessKeyProblem{})
}

// GetMessage returns the value for the field message
func (e *InternalServicesValidateAccessKeyInvalidAccessKeyProblem) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InternalServicesValidateAccessKeyInvalidAccessKeyProblem) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InternalServicesValidateAccessKeyInvalidAccessKeyProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalServicesValidateAccessKeyInvalidAccessKeyProblem.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InternalServicesValidateAccessKeyInvalidAccessKeyProblem) InitializeDefaults() {
}

// internalServicesValidateAccessKeyInvalidAccessKeyProblemAlias is defined to help pre and post JSON marshaling without recursive loops
type internalServicesValidateAccessKeyInvalidAccessKeyProblemAlias InternalServicesValidateAccessKeyInvalidAccessKeyProblem

// UnmarshalJSON implements json.Unmarshaler
func (e *InternalServicesValidateAccessKeyInvalidAccessKeyProblem) UnmarshalJSON(data []byte) error {
	var alias internalServicesValidateAccessKeyInvalidAccessKeyProblemAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InternalServicesValidateAccessKeyInvalidAccessKeyProblem)(&alias)).InitializeDefaults()
	*e = InternalServicesValidateAccessKeyInvalidAccessKeyProblem(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InternalServicesValidateAccessKeyInvalidAccessKeyProblem) MarshalJSON() ([]byte, error) {
	alias := internalServicesValidateAccessKeyInvalidAccessKeyProblemAlias(e)
	return json.Marshal(alias)
}

// NewInternalServicesValidateServiceBearerTokenInput creates a new InternalServicesValidateServiceBearerTokenInput
func NewInternalServicesValidateServiceBearerTokenInput() *InternalServicesValidateServiceBearerTokenInput {
	s := &InternalServicesValidateServiceBearerTokenInput{}
	s.InitializeDefaults()
	return s
}

// InternalServicesValidateServiceBearerTokenInput struct
type InternalServicesValidateServiceBearerTokenInput struct {
	RequestServiceBearerToken string `json:"requestServiceBearerToken,omitempty" yaml:"requestServiceBearerToken,omitempty"`
}

// GetRequestServiceBearerToken returns the value for the field requestServiceBearerToken
func (e *InternalServicesValidateServiceBearerTokenInput) GetRequestServiceBearerToken() string {
	return e.RequestServiceBearerToken
}

// SetRequestServiceBearerToken sets the value for the field requestServiceBearerToken
func (e *InternalServicesValidateServiceBearerTokenInput) SetRequestServiceBearerToken(requestServiceBearerToken string) {
	e.RequestServiceBearerToken = requestServiceBearerToken
}

// StructPath returns StructPath
func (e *InternalServicesValidateServiceBearerTokenInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalServicesValidateServiceBearerTokenInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InternalServicesValidateServiceBearerTokenInput) InitializeDefaults() {
}

// internalServicesValidateServiceBearerTokenInputAlias is defined to help pre and post JSON marshaling without recursive loops
type internalServicesValidateServiceBearerTokenInputAlias InternalServicesValidateServiceBearerTokenInput

// UnmarshalJSON implements json.Unmarshaler
func (e *InternalServicesValidateServiceBearerTokenInput) UnmarshalJSON(data []byte) error {
	var alias internalServicesValidateServiceBearerTokenInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InternalServicesValidateServiceBearerTokenInput)(&alias)).InitializeDefaults()
	*e = InternalServicesValidateServiceBearerTokenInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InternalServicesValidateServiceBearerTokenInput) MarshalJSON() ([]byte, error) {
	alias := internalServicesValidateServiceBearerTokenInputAlias(e)
	return json.Marshal(alias)
}

// NewInternalServicesValidateServiceBearerTokenOutput creates a new InternalServicesValidateServiceBearerTokenOutput
func NewInternalServicesValidateServiceBearerTokenOutput() *InternalServicesValidateServiceBearerTokenOutput {
	s := &InternalServicesValidateServiceBearerTokenOutput{}
	s.InitializeDefaults()
	return s
}

// InternalServicesValidateServiceBearerTokenOutput struct
type InternalServicesValidateServiceBearerTokenOutput struct {
	Key *InternalAccessKey `json:"key,omitempty" yaml:"key,omitempty"`
}

// GetKey returns the value for the field key
func (e *InternalServicesValidateServiceBearerTokenOutput) GetKey() *InternalAccessKey {
	return e.Key
}

// SetKey sets the value for the field key
func (e *InternalServicesValidateServiceBearerTokenOutput) SetKey(key *InternalAccessKey) {
	e.Key = key
}

// StructPath returns StructPath
func (e *InternalServicesValidateServiceBearerTokenOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalServicesValidateServiceBearerTokenOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InternalServicesValidateServiceBearerTokenOutput) InitializeDefaults() {
}

// internalServicesValidateServiceBearerTokenOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type internalServicesValidateServiceBearerTokenOutputAlias InternalServicesValidateServiceBearerTokenOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *InternalServicesValidateServiceBearerTokenOutput) UnmarshalJSON(data []byte) error {
	var alias internalServicesValidateServiceBearerTokenOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InternalServicesValidateServiceBearerTokenOutput)(&alias)).InitializeDefaults()
	*e = InternalServicesValidateServiceBearerTokenOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InternalServicesValidateServiceBearerTokenOutput) MarshalJSON() ([]byte, error) {
	alias := internalServicesValidateServiceBearerTokenOutputAlias(e)
	return json.Marshal(alias)
}

// NewInternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem creates a new InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem
func NewInternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem() *InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem {
	s := &InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem{}
	s.InitializeDefaults()
	return s
}

// InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem struct
type InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem]
func (e *InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem) Is(err error) bool {
	_, ok := err.(*InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem)
	return ok
}

// IsInternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem indicates whether the given error chain contains an error of type [InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem]
func IsInternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem(err error) bool {
	return errors.Is(err, &InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem{})
}

// GetMessage returns the value for the field message
func (e *InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem) InitializeDefaults() {
}

// internalServicesValidateServiceBearerTokenInvalidAccessKeyProblemAlias is defined to help pre and post JSON marshaling without recursive loops
type internalServicesValidateServiceBearerTokenInvalidAccessKeyProblemAlias InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem

// UnmarshalJSON implements json.Unmarshaler
func (e *InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem) UnmarshalJSON(data []byte) error {
	var alias internalServicesValidateServiceBearerTokenInvalidAccessKeyProblemAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem)(&alias)).InitializeDefaults()
	*e = InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem) MarshalJSON() ([]byte, error) {
	alias := internalServicesValidateServiceBearerTokenInvalidAccessKeyProblemAlias(e)
	return json.Marshal(alias)
}

// NewInternalServicesAssertActionInput creates a new InternalServicesAssertActionInput
func NewInternalServicesAssertActionInput() *InternalServicesAssertActionInput {
	s := &InternalServicesAssertActionInput{}
	s.InitializeDefaults()
	return s
}

// InternalServicesAssertActionInput struct
type InternalServicesAssertActionInput struct {
	CallerAccessKeyID string `json:"callerAccessKeyID,omitempty" yaml:"callerAccessKeyID,omitempty"`
	CallerAction      string `json:"callerAction,omitempty" yaml:"callerAction,omitempty"`
	// namespace of the resource to assert, e.g: "iam", "s3", "specular", etc
	CallerNamespace string                         `json:"callerNamespace,omitempty" yaml:"callerNamespace,omitempty"`
	CallerRegion    string                         `json:"callerRegion,omitempty" yaml:"callerRegion,omitempty"`
	CallerResource  *InternalAssertActionQualifier `json:"callerResource,omitempty" yaml:"callerResource,omitempty"`
}

// GetCallerAccessKeyID returns the value for the field callerAccessKeyID
func (e *InternalServicesAssertActionInput) GetCallerAccessKeyID() string {
	return e.CallerAccessKeyID
}

// SetCallerAccessKeyID sets the value for the field callerAccessKeyID
func (e *InternalServicesAssertActionInput) SetCallerAccessKeyID(callerAccessKeyID string) {
	e.CallerAccessKeyID = callerAccessKeyID
}

// GetCallerAction returns the value for the field callerAction
func (e *InternalServicesAssertActionInput) GetCallerAction() string {
	return e.CallerAction
}

// SetCallerAction sets the value for the field callerAction
func (e *InternalServicesAssertActionInput) SetCallerAction(callerAction string) {
	e.CallerAction = callerAction
}

// GetCallerNamespace returns the value for the field callerNamespace
func (e *InternalServicesAssertActionInput) GetCallerNamespace() string {
	return e.CallerNamespace
}

// SetCallerNamespace sets the value for the field callerNamespace
func (e *InternalServicesAssertActionInput) SetCallerNamespace(callerNamespace string) {
	e.CallerNamespace = callerNamespace
}

// GetCallerRegion returns the value for the field callerRegion
func (e *InternalServicesAssertActionInput) GetCallerRegion() string {
	return e.CallerRegion
}

// SetCallerRegion sets the value for the field callerRegion
func (e *InternalServicesAssertActionInput) SetCallerRegion(callerRegion string) {
	e.CallerRegion = callerRegion
}

// GetCallerResource returns the value for the field callerResource
func (e *InternalServicesAssertActionInput) GetCallerResource() *InternalAssertActionQualifier {
	return e.CallerResource
}

// SetCallerResource sets the value for the field callerResource
func (e *InternalServicesAssertActionInput) SetCallerResource(callerResource *InternalAssertActionQualifier) {
	e.CallerResource = callerResource
}

// StructPath returns StructPath
func (e *InternalServicesAssertActionInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalServicesAssertActionInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InternalServicesAssertActionInput) InitializeDefaults() {
}

// internalServicesAssertActionInputAlias is defined to help pre and post JSON marshaling without recursive loops
type internalServicesAssertActionInputAlias InternalServicesAssertActionInput

// UnmarshalJSON implements json.Unmarshaler
func (e *InternalServicesAssertActionInput) UnmarshalJSON(data []byte) error {
	var alias internalServicesAssertActionInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InternalServicesAssertActionInput)(&alias)).InitializeDefaults()
	*e = InternalServicesAssertActionInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InternalServicesAssertActionInput) MarshalJSON() ([]byte, error) {
	alias := internalServicesAssertActionInputAlias(e)
	return json.Marshal(alias)
}

// NewInternalServicesAssertActionOutput creates a new InternalServicesAssertActionOutput
func NewInternalServicesAssertActionOutput() *InternalServicesAssertActionOutput {
	s := &InternalServicesAssertActionOutput{}
	s.InitializeDefaults()
	return s
}

// InternalServicesAssertActionOutput struct
type InternalServicesAssertActionOutput struct {
}

// StructPath returns StructPath
func (e *InternalServicesAssertActionOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalServicesAssertActionOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InternalServicesAssertActionOutput) InitializeDefaults() {
}

// internalServicesAssertActionOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type internalServicesAssertActionOutputAlias InternalServicesAssertActionOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *InternalServicesAssertActionOutput) UnmarshalJSON(data []byte) error {
	var alias internalServicesAssertActionOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InternalServicesAssertActionOutput)(&alias)).InitializeDefaults()
	*e = InternalServicesAssertActionOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InternalServicesAssertActionOutput) MarshalJSON() ([]byte, error) {
	alias := internalServicesAssertActionOutputAlias(e)
	return json.Marshal(alias)
}

// NewInternalServicesAssertQueryInput creates a new InternalServicesAssertQueryInput
func NewInternalServicesAssertQueryInput() *InternalServicesAssertQueryInput {
	s := &InternalServicesAssertQueryInput{}
	s.InitializeDefaults()
	return s
}

// InternalServicesAssertQueryInput struct
type InternalServicesAssertQueryInput struct {
	CallerAccessKeyID string `json:"callerAccessKeyID,omitempty" yaml:"callerAccessKeyID,omitempty"`
	CallerAction      string `json:"callerAction,omitempty" yaml:"callerAction,omitempty"`
	// namespace of the resource to assert, e.g: "iam", "s3", "specular", etc
	CallerNamespace string `json:"callerNamespace,omitempty" yaml:"callerNamespace,omitempty"`
	CallerRegion    string `json:"callerRegion,omitempty" yaml:"callerRegion,omitempty"`
	// resource replacements where each resource value is a replacement for a SQL expression in the query
	CallerResourceReplacements *InternalAssertActionQualifier `json:"callerResourceReplacements,omitempty" yaml:"callerResourceReplacements,omitempty"`
}

// GetCallerAccessKeyID returns the value for the field callerAccessKeyID
func (e *InternalServicesAssertQueryInput) GetCallerAccessKeyID() string {
	return e.CallerAccessKeyID
}

// SetCallerAccessKeyID sets the value for the field callerAccessKeyID
func (e *InternalServicesAssertQueryInput) SetCallerAccessKeyID(callerAccessKeyID string) {
	e.CallerAccessKeyID = callerAccessKeyID
}

// GetCallerAction returns the value for the field callerAction
func (e *InternalServicesAssertQueryInput) GetCallerAction() string {
	return e.CallerAction
}

// SetCallerAction sets the value for the field callerAction
func (e *InternalServicesAssertQueryInput) SetCallerAction(callerAction string) {
	e.CallerAction = callerAction
}

// GetCallerNamespace returns the value for the field callerNamespace
func (e *InternalServicesAssertQueryInput) GetCallerNamespace() string {
	return e.CallerNamespace
}

// SetCallerNamespace sets the value for the field callerNamespace
func (e *InternalServicesAssertQueryInput) SetCallerNamespace(callerNamespace string) {
	e.CallerNamespace = callerNamespace
}

// GetCallerRegion returns the value for the field callerRegion
func (e *InternalServicesAssertQueryInput) GetCallerRegion() string {
	return e.CallerRegion
}

// SetCallerRegion sets the value for the field callerRegion
func (e *InternalServicesAssertQueryInput) SetCallerRegion(callerRegion string) {
	e.CallerRegion = callerRegion
}

// GetCallerResourceReplacements returns the value for the field callerResourceReplacements
func (e *InternalServicesAssertQueryInput) GetCallerResourceReplacements() *InternalAssertActionQualifier {
	return e.CallerResourceReplacements
}

// SetCallerResourceReplacements sets the value for the field callerResourceReplacements
func (e *InternalServicesAssertQueryInput) SetCallerResourceReplacements(callerResourceReplacements *InternalAssertActionQualifier) {
	e.CallerResourceReplacements = callerResourceReplacements
}

// StructPath returns StructPath
func (e *InternalServicesAssertQueryInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalServicesAssertQueryInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InternalServicesAssertQueryInput) InitializeDefaults() {
}

// internalServicesAssertQueryInputAlias is defined to help pre and post JSON marshaling without recursive loops
type internalServicesAssertQueryInputAlias InternalServicesAssertQueryInput

// UnmarshalJSON implements json.Unmarshaler
func (e *InternalServicesAssertQueryInput) UnmarshalJSON(data []byte) error {
	var alias internalServicesAssertQueryInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InternalServicesAssertQueryInput)(&alias)).InitializeDefaults()
	*e = InternalServicesAssertQueryInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InternalServicesAssertQueryInput) MarshalJSON() ([]byte, error) {
	alias := internalServicesAssertQueryInputAlias(e)
	return json.Marshal(alias)
}

// NewInternalServicesAssertQueryOutput creates a new InternalServicesAssertQueryOutput
func NewInternalServicesAssertQueryOutput() *InternalServicesAssertQueryOutput {
	s := &InternalServicesAssertQueryOutput{}
	s.InitializeDefaults()
	return s
}

// InternalServicesAssertQueryOutput struct
type InternalServicesAssertQueryOutput struct {
	Entries []*InternalQueryAssertionEntry `json:"entries,omitempty" yaml:"entries,omitempty"`
}

// GetEntries returns the value for the field entries
func (e *InternalServicesAssertQueryOutput) GetEntries() []*InternalQueryAssertionEntry {
	return e.Entries
}

// SetEntries sets the value for the field entries
func (e *InternalServicesAssertQueryOutput) SetEntries(entries []*InternalQueryAssertionEntry) {
	e.Entries = entries
}

// StructPath returns StructPath
func (e *InternalServicesAssertQueryOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalServicesAssertQueryOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InternalServicesAssertQueryOutput) InitializeDefaults() {
}

// internalServicesAssertQueryOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type internalServicesAssertQueryOutputAlias InternalServicesAssertQueryOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *InternalServicesAssertQueryOutput) UnmarshalJSON(data []byte) error {
	var alias internalServicesAssertQueryOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InternalServicesAssertQueryOutput)(&alias)).InitializeDefaults()
	*e = InternalServicesAssertQueryOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InternalServicesAssertQueryOutput) MarshalJSON() ([]byte, error) {
	alias := internalServicesAssertQueryOutputAlias(e)
	return json.Marshal(alias)
}

// NewInternalServicesAssertQueryReplacementProblem creates a new InternalServicesAssertQueryReplacementProblem
func NewInternalServicesAssertQueryReplacementProblem() *InternalServicesAssertQueryReplacementProblem {
	s := &InternalServicesAssertQueryReplacementProblem{}
	s.InitializeDefaults()
	return s
}

// InternalServicesAssertQueryReplacementProblem struct
type InternalServicesAssertQueryReplacementProblem struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InternalServicesAssertQueryReplacementProblem) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InternalServicesAssertQueryReplacementProblem]
func (e *InternalServicesAssertQueryReplacementProblem) Is(err error) bool {
	_, ok := err.(*InternalServicesAssertQueryReplacementProblem)
	return ok
}

// IsInternalServicesAssertQueryReplacementProblem indicates whether the given error chain contains an error of type [InternalServicesAssertQueryReplacementProblem]
func IsInternalServicesAssertQueryReplacementProblem(err error) bool {
	return errors.Is(err, &InternalServicesAssertQueryReplacementProblem{})
}

// GetMessage returns the value for the field message
func (e *InternalServicesAssertQueryReplacementProblem) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InternalServicesAssertQueryReplacementProblem) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InternalServicesAssertQueryReplacementProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalServicesAssertQueryReplacementProblem.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InternalServicesAssertQueryReplacementProblem) InitializeDefaults() {
}

// internalServicesAssertQueryReplacementProblemAlias is defined to help pre and post JSON marshaling without recursive loops
type internalServicesAssertQueryReplacementProblemAlias InternalServicesAssertQueryReplacementProblem

// UnmarshalJSON implements json.Unmarshaler
func (e *InternalServicesAssertQueryReplacementProblem) UnmarshalJSON(data []byte) error {
	var alias internalServicesAssertQueryReplacementProblemAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InternalServicesAssertQueryReplacementProblem)(&alias)).InitializeDefaults()
	*e = InternalServicesAssertQueryReplacementProblem(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InternalServicesAssertQueryReplacementProblem) MarshalJSON() ([]byte, error) {
	alias := internalServicesAssertQueryReplacementProblemAlias(e)
	return json.Marshal(alias)
}

// PolicyStatementEffect entity
type PolicyStatementEffect string

// PolicyStatementEffectDeny = "DENY"
const PolicyStatementEffectDeny PolicyStatementEffect = "DENY"

// PolicyStatementEffectAllow = "ALLOW"
const PolicyStatementEffectAllow PolicyStatementEffect = "ALLOW"

// String returns the string representation of the enum
func (en PolicyStatementEffect) String() string {
	return string(en)
}

// Optional returns the optional value
func (en PolicyStatementEffect) Optional() *PolicyStatementEffect {
	c := en
	return &c
}

// PolicyStatementEffectAllConstants returns all constants in enum PolicyStatementEffect
func PolicyStatementEffectAllConstants() []PolicyStatementEffect {
	return []PolicyStatementEffect{
		PolicyStatementEffectDeny,
		PolicyStatementEffectAllow,
	}
}

var packagePath = clientruntime.ModulePathFromTrustedValues(
	"deployport",
	"iam",
)

func newSpecularPackage() (pk *clientruntime.Package, err error) {
	pk = clientruntime.NewPackage(packagePath)
	if err := pk.Import(godeployportcomapiservicescorelib.SpecularMeta().Module()); err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserInformationSSOProvider, err = pk.NewType(
		"UserInformationSSOProvider",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserInformationSSOProvider()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserInformationSSOProfile, err = pk.NewType(
		"UserInformationSSOProfile",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserInformationSSOProfile()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserInformationSSO, err = pk.NewType(
		"UserInformationSSO",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserInformationSSO()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserInformation, err = pk.NewType(
		"UserInformation",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserInformation()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleInformation, err = pk.NewType(
		"RoleInformation",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleInformation()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathCredentials, err = pk.NewType(
		"Credentials",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewCredentials()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathSSOProviderUnavailableProblem, err = pk.NewType(
		"SSOProviderUnavailableProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewSSOProviderUnavailableProblem()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathSSOFlow, err = pk.NewType(
		"SSOFlow",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewSSOFlow()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccount, err = pk.NewType(
		"Account",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccount()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountSSOProvider, err = pk.NewType(
		"AccountSSOProvider",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountSSOProvider()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathPolicyNotFoundProblem, err = pk.NewType(
		"PolicyNotFoundProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewPolicyNotFoundProblem()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountSSOAutoJoinPolicy, err = pk.NewType(
		"AccountSSOAutoJoinPolicy",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountSSOAutoJoinPolicy()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountCreateInput, err = pk.NewType(
		"AccountCreateInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountCreateInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountCreateOutput, err = pk.NewType(
		"AccountCreateOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountCreateOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountCreateInvalidNameProblem, err = pk.NewType(
		"AccountCreateInvalidNameProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountCreateInvalidNameProblem()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountAssumeIdentityInput, err = pk.NewType(
		"AccountAssumeIdentityInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountAssumeIdentityInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountAssumeIdentityOutput, err = pk.NewType(
		"AccountAssumeIdentityOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountAssumeIdentityOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountSSOBeginAuthenticationInput, err = pk.NewType(
		"AccountSSOBeginAuthenticationInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountSSOBeginAuthenticationInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountSSOBeginAuthenticationOutput, err = pk.NewType(
		"AccountSSOBeginAuthenticationOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountSSOBeginAuthenticationOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountSSOBeginAuthenticationParameterProblem, err = pk.NewType(
		"AccountSSOBeginAuthenticationParameterProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountSSOBeginAuthenticationParameterProblem()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountSSOCompleteAuthenticationInput, err = pk.NewType(
		"AccountSSOCompleteAuthenticationInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountSSOCompleteAuthenticationInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountSSOCompleteAuthenticationOutput, err = pk.NewType(
		"AccountSSOCompleteAuthenticationOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountSSOCompleteAuthenticationOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountSSOCompleteAuthenticationInvalidFlowProblem, err = pk.NewType(
		"AccountSSOCompleteAuthenticationInvalidFlowProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountSSOCompleteAuthenticationInvalidFlowProblem()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountSSOGetProvidersInput, err = pk.NewType(
		"AccountSSOGetProvidersInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountSSOGetProvidersInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountSSOGetProvidersOutput, err = pk.NewType(
		"AccountSSOGetProvidersOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountSSOGetProvidersOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountSSOAutoJoinPolicyCreateInput, err = pk.NewType(
		"AccountSSOAutoJoinPolicyCreateInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountSSOAutoJoinPolicyCreateInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountSSOAutoJoinPolicyCreateOutput, err = pk.NewType(
		"AccountSSOAutoJoinPolicyCreateOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountSSOAutoJoinPolicyCreateOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountSSOAutoJoinPolicyListInput, err = pk.NewType(
		"AccountSSOAutoJoinPolicyListInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountSSOAutoJoinPolicyListInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountSSOAutoJoinPolicyListOutput, err = pk.NewType(
		"AccountSSOAutoJoinPolicyListOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountSSOAutoJoinPolicyListOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountSSOAutoJoinPolicyEnableInput, err = pk.NewType(
		"AccountSSOAutoJoinPolicyEnableInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountSSOAutoJoinPolicyEnableInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountSSOAutoJoinPolicyEnableOutput, err = pk.NewType(
		"AccountSSOAutoJoinPolicyEnableOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountSSOAutoJoinPolicyEnableOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvalidUsernameProblem, err = pk.NewType(
		"InvalidUsernameProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvalidUsernameProblem()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvalidRoleNameProblem, err = pk.NewType(
		"InvalidRoleNameProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvalidRoleNameProblem()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathMemberAccount, err = pk.NewType(
		"MemberAccount",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewMemberAccount()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathCredentialInfo, err = pk.NewType(
		"CredentialInfo",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewCredentialInfo()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathIdentityPolicyStatement, err = pk.NewType(
		"IdentityPolicyStatement",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewIdentityPolicyStatement()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathIdentityPolicy, err = pk.NewType(
		"IdentityPolicy",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewIdentityPolicy()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathIdentityPolicyAttachment, err = pk.NewType(
		"IdentityPolicyAttachment",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewIdentityPolicyAttachment()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserNotFoundProblem, err = pk.NewType(
		"UserNotFoundProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserNotFoundProblem()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleNotFoundProblem, err = pk.NewType(
		"RoleNotFoundProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleNotFoundProblem()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathIdentityInUseProblem, err = pk.NewType(
		"IdentityInUseProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewIdentityInUseProblem()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserCreateInput, err = pk.NewType(
		"UserCreateInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserCreateInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserCreateOutput, err = pk.NewType(
		"UserCreateOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserCreateOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserGetInput, err = pk.NewType(
		"UserGetInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserGetInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserGetOutput, err = pk.NewType(
		"UserGetOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserGetOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserGetUserNotAvailableProblem, err = pk.NewType(
		"UserGetUserNotAvailableProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserGetUserNotAvailableProblem()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserDestroyInput, err = pk.NewType(
		"UserDestroyInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserDestroyInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserDestroyOutput, err = pk.NewType(
		"UserDestroyOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserDestroyOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserListInput, err = pk.NewType(
		"UserListInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserListInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserListOutput, err = pk.NewType(
		"UserListOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserListOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserMemberAccountsInput, err = pk.NewType(
		"UserMemberAccountsInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserMemberAccountsInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserMemberAccountsOutput, err = pk.NewType(
		"UserMemberAccountsOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserMemberAccountsOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserAccessKeyCreateInput, err = pk.NewType(
		"UserAccessKeyCreateInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserAccessKeyCreateInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserAccessKeyCreateOutput, err = pk.NewType(
		"UserAccessKeyCreateOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserAccessKeyCreateOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserAccessKeyListInput, err = pk.NewType(
		"UserAccessKeyListInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserAccessKeyListInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserAccessKeyListOutput, err = pk.NewType(
		"UserAccessKeyListOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserAccessKeyListOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserAccessKeyDestroyInput, err = pk.NewType(
		"UserAccessKeyDestroyInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserAccessKeyDestroyInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserAccessKeyDestroyOutput, err = pk.NewType(
		"UserAccessKeyDestroyOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserAccessKeyDestroyOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserIdentityPolicyAttachInput, err = pk.NewType(
		"UserIdentityPolicyAttachInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserIdentityPolicyAttachInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserIdentityPolicyAttachOutput, err = pk.NewType(
		"UserIdentityPolicyAttachOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserIdentityPolicyAttachOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserIdentityPolicyListInput, err = pk.NewType(
		"UserIdentityPolicyListInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserIdentityPolicyListInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserIdentityPolicyListOutput, err = pk.NewType(
		"UserIdentityPolicyListOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserIdentityPolicyListOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserIdentityPolicyDetachInput, err = pk.NewType(
		"UserIdentityPolicyDetachInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserIdentityPolicyDetachInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserIdentityPolicyDetachOutput, err = pk.NewType(
		"UserIdentityPolicyDetachOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserIdentityPolicyDetachOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInlinePolicy, err = pk.NewType(
		"InlinePolicy",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInlinePolicy()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathPolicyStructureProblem, err = pk.NewType(
		"PolicyStructureProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewPolicyStructureProblem()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleCreateInput, err = pk.NewType(
		"RoleCreateInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleCreateInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleCreateOutput, err = pk.NewType(
		"RoleCreateOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleCreateOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleGetInput, err = pk.NewType(
		"RoleGetInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleGetInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleGetOutput, err = pk.NewType(
		"RoleGetOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleGetOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleDestroyInput, err = pk.NewType(
		"RoleDestroyInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleDestroyInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleDestroyOutput, err = pk.NewType(
		"RoleDestroyOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleDestroyOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleListInput, err = pk.NewType(
		"RoleListInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleListInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleListOutput, err = pk.NewType(
		"RoleListOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleListOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleAssumeInput, err = pk.NewType(
		"RoleAssumeInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleAssumeInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleAssumeOutput, err = pk.NewType(
		"RoleAssumeOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleAssumeOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleIdentityPolicyAttachInput, err = pk.NewType(
		"RoleIdentityPolicyAttachInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleIdentityPolicyAttachInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleIdentityPolicyAttachOutput, err = pk.NewType(
		"RoleIdentityPolicyAttachOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleIdentityPolicyAttachOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleIdentityPolicyListInput, err = pk.NewType(
		"RoleIdentityPolicyListInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleIdentityPolicyListInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleIdentityPolicyListOutput, err = pk.NewType(
		"RoleIdentityPolicyListOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleIdentityPolicyListOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleIdentityPolicyDetachInput, err = pk.NewType(
		"RoleIdentityPolicyDetachInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleIdentityPolicyDetachInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleIdentityPolicyDetachOutput, err = pk.NewType(
		"RoleIdentityPolicyDetachOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleIdentityPolicyDetachOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathIdentityPolicyCreateInput, err = pk.NewType(
		"IdentityPolicyCreateInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewIdentityPolicyCreateInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathIdentityPolicyCreateOutput, err = pk.NewType(
		"IdentityPolicyCreateOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewIdentityPolicyCreateOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathIdentityPolicyListInput, err = pk.NewType(
		"IdentityPolicyListInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewIdentityPolicyListInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathIdentityPolicyListOutput, err = pk.NewType(
		"IdentityPolicyListOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewIdentityPolicyListOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathIdentityPolicyRetrieveInput, err = pk.NewType(
		"IdentityPolicyRetrieveInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewIdentityPolicyRetrieveInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathIdentityPolicyRetrieveOutput, err = pk.NewType(
		"IdentityPolicyRetrieveOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewIdentityPolicyRetrieveOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathIdentityPolicyDestroyInput, err = pk.NewType(
		"IdentityPolicyDestroyInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewIdentityPolicyDestroyInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathIdentityPolicyDestroyOutput, err = pk.NewType(
		"IdentityPolicyDestroyOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewIdentityPolicyDestroyOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathIdentityPolicyUpdateInput, err = pk.NewType(
		"IdentityPolicyUpdateInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewIdentityPolicyUpdateInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathIdentityPolicyUpdateOutput, err = pk.NewType(
		"IdentityPolicyUpdateOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewIdentityPolicyUpdateOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathServiceBearerToken, err = pk.NewType(
		"ServiceBearerToken",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewServiceBearerToken()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvalidServiceBearerTokenDurationProblem, err = pk.NewType(
		"InvalidServiceBearerTokenDurationProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvalidServiceBearerTokenDurationProblem()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvalidServiceNameProblem, err = pk.NewType(
		"InvalidServiceNameProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvalidServiceNameProblem()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathServiceBearerTokenGetInput, err = pk.NewType(
		"ServiceBearerTokenGetInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewServiceBearerTokenGetInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathServiceBearerTokenGetOutput, err = pk.NewType(
		"ServiceBearerTokenGetOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewServiceBearerTokenGetOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInternalAccessKeyUser, err = pk.NewType(
		"InternalAccessKeyUser",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInternalAccessKeyUser()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInternalAccessKeyRole, err = pk.NewType(
		"InternalAccessKeyRole",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInternalAccessKeyRole()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInternalAccessKeyAccount, err = pk.NewType(
		"InternalAccessKeyAccount",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInternalAccessKeyAccount()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInternalAccessKey, err = pk.NewType(
		"InternalAccessKey",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInternalAccessKey()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInternalQueryAssertionEntryValue, err = pk.NewType(
		"InternalQueryAssertionEntryValue",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInternalQueryAssertionEntryValue()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInternalQueryAssertionEntry, err = pk.NewType(
		"InternalQueryAssertionEntry",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInternalQueryAssertionEntry()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInternalServicesAssertActionCallerForbiddenProblem, err = pk.NewType(
		"InternalServicesAssertActionCallerForbiddenProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInternalServicesAssertActionCallerForbiddenProblem()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInternalAssertActionQualifier, err = pk.NewType(
		"InternalAssertActionQualifier",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInternalAssertActionQualifier()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInternalServicesValidateAccessKeyInput, err = pk.NewType(
		"InternalServicesValidateAccessKeyInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInternalServicesValidateAccessKeyInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInternalServicesValidateAccessKeyOutput, err = pk.NewType(
		"InternalServicesValidateAccessKeyOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInternalServicesValidateAccessKeyOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInternalServicesValidateAccessKeyInvalidAccessKeyProblem, err = pk.NewType(
		"InternalServicesValidateAccessKeyInvalidAccessKeyProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInternalServicesValidateAccessKeyInvalidAccessKeyProblem()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInternalServicesValidateServiceBearerTokenInput, err = pk.NewType(
		"InternalServicesValidateServiceBearerTokenInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInternalServicesValidateServiceBearerTokenInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInternalServicesValidateServiceBearerTokenOutput, err = pk.NewType(
		"InternalServicesValidateServiceBearerTokenOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInternalServicesValidateServiceBearerTokenOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem, err = pk.NewType(
		"InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInternalServicesAssertActionInput, err = pk.NewType(
		"InternalServicesAssertActionInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInternalServicesAssertActionInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInternalServicesAssertActionOutput, err = pk.NewType(
		"InternalServicesAssertActionOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInternalServicesAssertActionOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInternalServicesAssertQueryInput, err = pk.NewType(
		"InternalServicesAssertQueryInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInternalServicesAssertQueryInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInternalServicesAssertQueryOutput, err = pk.NewType(
		"InternalServicesAssertQueryOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInternalServicesAssertQueryOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInternalServicesAssertQueryReplacementProblem, err = pk.NewType(
		"InternalServicesAssertQueryReplacementProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInternalServicesAssertQueryReplacementProblem()
		}),
	)
	if err != nil {
		return nil, err
	}
	resAccount, err := pk.NewResource("Account")
	if err != nil {
		return nil, err
	}
	_ = resAccount

	var op *clientruntime.Operation
	op, err = resAccount.NewOperation("Create")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountCreateInputStruct())
	op.SetOutput(SpecularMeta().AccountCreateInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())
	op.RegisterProblemType(SpecularMeta().AccountCreateInvalidNameProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resAccount.NewOperation("AssumeIdentity")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountAssumeIdentityInputStruct())
	op.SetOutput(SpecularMeta().AccountAssumeIdentityInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})
	// subresource AccountSSO
	resAccountSSO, err := resAccount.NewSubResource("SSO")
	if err != nil {
		return nil, err
	}
	_ = resAccountSSO

	op, err = resAccountSSO.NewOperation("BeginAuthentication")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountSSOBeginAuthenticationInputStruct())
	op.SetOutput(SpecularMeta().AccountSSOBeginAuthenticationInputStruct())
	op.RegisterProblemType(SpecularMeta().SSOProviderUnavailableProblemStruct())
	op.RegisterProblemType(SpecularMeta().AccountSSOBeginAuthenticationParameterProblemStruct())

	op, err = resAccountSSO.NewOperation("CompleteAuthentication")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountSSOCompleteAuthenticationInputStruct())
	op.SetOutput(SpecularMeta().AccountSSOCompleteAuthenticationInputStruct())
	op.RegisterProblemType(SpecularMeta().AccountSSOCompleteAuthenticationInvalidFlowProblemStruct())

	op, err = resAccountSSO.NewOperation("GetProviders")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountSSOGetProvidersInputStruct())
	op.SetOutput(SpecularMeta().AccountSSOGetProvidersInputStruct())

	// subresource AccountSSOAutoJoinPolicy
	resAccountSSOAutoJoinPolicy, err := resAccountSSO.NewSubResource("AutoJoinPolicy")
	if err != nil {
		return nil, err
	}
	_ = resAccountSSOAutoJoinPolicy

	op, err = resAccountSSOAutoJoinPolicy.NewOperation("Create")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountSSOAutoJoinPolicyCreateInputStruct())
	op.SetOutput(SpecularMeta().AccountSSOAutoJoinPolicyCreateInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resAccountSSOAutoJoinPolicy.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountSSOAutoJoinPolicyListInputStruct())
	op.SetOutput(SpecularMeta().AccountSSOAutoJoinPolicyListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resAccountSSOAutoJoinPolicy.NewOperation("Enable")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountSSOAutoJoinPolicyEnableInputStruct())
	op.SetOutput(SpecularMeta().AccountSSOAutoJoinPolicyEnableInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	resUser, err := pk.NewResource("User")
	if err != nil {
		return nil, err
	}
	_ = resUser

	op, err = resUser.NewOperation("Create")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserCreateInputStruct())
	op.SetOutput(SpecularMeta().UserCreateInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())
	op.RegisterProblemType(SpecularMeta().InvalidUsernameProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUser.NewOperation("Get")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserGetInputStruct())
	op.SetOutput(SpecularMeta().UserGetInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(SpecularMeta().UserGetUserNotAvailableProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUser.NewOperation("Destroy")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserDestroyInputStruct())
	op.SetOutput(SpecularMeta().UserDestroyInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())
	op.RegisterProblemType(SpecularMeta().IdentityInUseProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUser.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserListInputStruct())
	op.SetOutput(SpecularMeta().UserListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUser.NewOperation("MemberAccounts")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserMemberAccountsInputStruct())
	op.SetOutput(SpecularMeta().UserMemberAccountsInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})
	// subresource UserAccessKey
	resUserAccessKey, err := resUser.NewSubResource("AccessKey")
	if err != nil {
		return nil, err
	}
	_ = resUserAccessKey

	op, err = resUserAccessKey.NewOperation("Create")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserAccessKeyCreateInputStruct())
	op.SetOutput(SpecularMeta().UserAccessKeyCreateInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUserAccessKey.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserAccessKeyListInputStruct())
	op.SetOutput(SpecularMeta().UserAccessKeyListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUserAccessKey.NewOperation("Destroy")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserAccessKeyDestroyInputStruct())
	op.SetOutput(SpecularMeta().UserAccessKeyDestroyInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	// subresource UserIdentityPolicy
	resUserIdentityPolicy, err := resUser.NewSubResource("IdentityPolicy")
	if err != nil {
		return nil, err
	}
	_ = resUserIdentityPolicy

	op, err = resUserIdentityPolicy.NewOperation("Attach")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserIdentityPolicyAttachInputStruct())
	op.SetOutput(SpecularMeta().UserIdentityPolicyAttachInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())
	op.RegisterProblemType(SpecularMeta().PolicyNotFoundProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUserIdentityPolicy.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserIdentityPolicyListInputStruct())
	op.SetOutput(SpecularMeta().UserIdentityPolicyListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUserIdentityPolicy.NewOperation("Detach")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserIdentityPolicyDetachInputStruct())
	op.SetOutput(SpecularMeta().UserIdentityPolicyDetachInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())
	op.RegisterProblemType(SpecularMeta().PolicyNotFoundProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	resRole, err := pk.NewResource("Role")
	if err != nil {
		return nil, err
	}
	_ = resRole

	op, err = resRole.NewOperation("Create")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().RoleCreateInputStruct())
	op.SetOutput(SpecularMeta().RoleCreateInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())
	op.RegisterProblemType(SpecularMeta().InvalidRoleNameProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resRole.NewOperation("Get")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().RoleGetInputStruct())
	op.SetOutput(SpecularMeta().RoleGetInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resRole.NewOperation("Destroy")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().RoleDestroyInputStruct())
	op.SetOutput(SpecularMeta().RoleDestroyInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())
	op.RegisterProblemType(SpecularMeta().IdentityInUseProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resRole.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().RoleListInputStruct())
	op.SetOutput(SpecularMeta().RoleListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resRole.NewOperation("Assume")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().RoleAssumeInputStruct())
	op.SetOutput(SpecularMeta().RoleAssumeInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())
	op.RegisterProblemType(SpecularMeta().PolicyStructureProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})
	// subresource RoleIdentityPolicy
	resRoleIdentityPolicy, err := resRole.NewSubResource("IdentityPolicy")
	if err != nil {
		return nil, err
	}
	_ = resRoleIdentityPolicy

	op, err = resRoleIdentityPolicy.NewOperation("Attach")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().RoleIdentityPolicyAttachInputStruct())
	op.SetOutput(SpecularMeta().RoleIdentityPolicyAttachInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())
	op.RegisterProblemType(SpecularMeta().PolicyNotFoundProblemStruct())
	op.RegisterProblemType(SpecularMeta().RoleNotFoundProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resRoleIdentityPolicy.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().RoleIdentityPolicyListInputStruct())
	op.SetOutput(SpecularMeta().RoleIdentityPolicyListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resRoleIdentityPolicy.NewOperation("Detach")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().RoleIdentityPolicyDetachInputStruct())
	op.SetOutput(SpecularMeta().RoleIdentityPolicyDetachInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())
	op.RegisterProblemType(SpecularMeta().PolicyNotFoundProblemStruct())
	op.RegisterProblemType(SpecularMeta().RoleNotFoundProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	resIdentityPolicy, err := pk.NewResource("IdentityPolicy")
	if err != nil {
		return nil, err
	}
	_ = resIdentityPolicy

	op, err = resIdentityPolicy.NewOperation("Create")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().IdentityPolicyCreateInputStruct())
	op.SetOutput(SpecularMeta().IdentityPolicyCreateInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())
	op.RegisterProblemType(SpecularMeta().PolicyStructureProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resIdentityPolicy.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().IdentityPolicyListInputStruct())
	op.SetOutput(SpecularMeta().IdentityPolicyListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resIdentityPolicy.NewOperation("Retrieve")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().IdentityPolicyRetrieveInputStruct())
	op.SetOutput(SpecularMeta().IdentityPolicyRetrieveInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())
	op.RegisterProblemType(SpecularMeta().PolicyNotFoundProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resIdentityPolicy.NewOperation("Destroy")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().IdentityPolicyDestroyInputStruct())
	op.SetOutput(SpecularMeta().IdentityPolicyDestroyInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resIdentityPolicy.NewOperation("Update")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().IdentityPolicyUpdateInputStruct())
	op.SetOutput(SpecularMeta().IdentityPolicyUpdateInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())
	op.RegisterProblemType(SpecularMeta().PolicyStructureProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	resServiceBearerToken, err := pk.NewResource("ServiceBearerToken")
	if err != nil {
		return nil, err
	}
	_ = resServiceBearerToken

	op, err = resServiceBearerToken.NewOperation("Get")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().ServiceBearerTokenGetInputStruct())
	op.SetOutput(SpecularMeta().ServiceBearerTokenGetInputStruct())
	op.RegisterProblemType(SpecularMeta().InvalidServiceBearerTokenDurationProblemStruct())
	op.RegisterProblemType(SpecularMeta().InvalidServiceNameProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	resInternal, err := pk.NewResource("Internal")
	if err != nil {
		return nil, err
	}
	_ = resInternal

	// subresource InternalServices
	resInternalServices, err := resInternal.NewSubResource("Services")
	if err != nil {
		return nil, err
	}
	_ = resInternalServices

	op, err = resInternalServices.NewOperation("ValidateAccessKey")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().InternalServicesValidateAccessKeyInputStruct())
	op.SetOutput(SpecularMeta().InternalServicesValidateAccessKeyInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())
	op.RegisterProblemType(SpecularMeta().InternalServicesValidateAccessKeyInvalidAccessKeyProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resInternalServices.NewOperation("ValidateServiceBearerToken")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().InternalServicesValidateServiceBearerTokenInputStruct())
	op.SetOutput(SpecularMeta().InternalServicesValidateServiceBearerTokenInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())
	op.RegisterProblemType(SpecularMeta().InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resInternalServices.NewOperation("AssertAction")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().InternalServicesAssertActionInputStruct())
	op.SetOutput(SpecularMeta().InternalServicesAssertActionInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())
	op.RegisterProblemType(SpecularMeta().InternalServicesAssertActionCallerForbiddenProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resInternalServices.NewOperation("AssertQuery")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().InternalServicesAssertQueryInputStruct())
	op.SetOutput(SpecularMeta().InternalServicesAssertQueryInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblemStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblemStruct())
	op.RegisterProblemType(SpecularMeta().InternalServicesAssertActionCallerForbiddenProblemStruct())
	op.RegisterProblemType(SpecularMeta().InternalServicesAssertQueryReplacementProblemStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	pk.AddAnnotation(&godeployportcomapiservicescorelib.ServiceSignatureV1{
		ServiceName: "iam",
	})
	return pk, nil
}

// AccountResource is the AccountResource resource client
type AccountResource struct {
	transport      clientruntime.Transport
	res            *clientruntime.Resource
	SSO            *AccountSSOResource
	create         *clientruntime.Operation
	assumeIdentity *clientruntime.Operation
}

func newAccountResource(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*AccountResource, error) {
	res := finder.FindResource("Account")
	r := &AccountResource{
		transport: transport,
		res:       res,
	}
	var err error
	r.SSO, err = newAccountSSOResource(transport, res)
	if err != nil {
		return nil, err
	}
	r.create = res.FindOperation("Create")
	r.assumeIdentity = res.FindOperation("AssumeIdentity")
	return r, nil
}

// Create - Creates a new Account with the given name
// Create account only works for users of the "global" account
func (res *AccountResource) Create(ctx context.Context, input *AccountCreateInput) (*AccountCreateOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.create,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*AccountCreateOutput)
	return output, nil
}

// AssumeIdentity - Assume identity in the given account
func (res *AccountResource) AssumeIdentity(ctx context.Context, input *AccountAssumeIdentityInput) (*AccountAssumeIdentityOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.assumeIdentity,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*AccountAssumeIdentityOutput)
	return output, nil
}

// AccountSSOResource is the AccountSSOResource resource client
type AccountSSOResource struct {
	transport              clientruntime.Transport
	res                    *clientruntime.Resource
	AutoJoinPolicy         *AccountSSOAutoJoinPolicyResource
	beginAuthentication    *clientruntime.Operation
	completeAuthentication *clientruntime.Operation
	getProviders           *clientruntime.Operation
}

func newAccountSSOResource(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*AccountSSOResource, error) {
	res := finder.FindResource("SSO")
	r := &AccountSSOResource{
		transport: transport,
		res:       res,
	}
	var err error
	r.AutoJoinPolicy, err = newAccountSSOAutoJoinPolicyResource(transport, res)
	if err != nil {
		return nil, err
	}
	r.beginAuthentication = res.FindOperation("BeginAuthentication")
	r.completeAuthentication = res.FindOperation("CompleteAuthentication")
	r.getProviders = res.FindOperation("GetProviders")
	return r, nil
}

// BeginAuthentication - Returns the public link to the login page for the given SSO provider
func (res *AccountSSOResource) BeginAuthentication(ctx context.Context, input *AccountSSOBeginAuthenticationInput) (*AccountSSOBeginAuthenticationOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.beginAuthentication,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*AccountSSOBeginAuthenticationOutput)
	return output, nil
}

// CompleteAuthentication - Authorizes the user with the given code and returns a token
func (res *AccountSSOResource) CompleteAuthentication(ctx context.Context, input *AccountSSOCompleteAuthenticationInput) (*AccountSSOCompleteAuthenticationOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.completeAuthentication,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*AccountSSOCompleteAuthenticationOutput)
	return output, nil
}

// GetProviders - Returns the providers available for the given account
// This is a public operation that does not require any authentication
func (res *AccountSSOResource) GetProviders(ctx context.Context, input *AccountSSOGetProvidersInput) (*AccountSSOGetProvidersOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.getProviders,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*AccountSSOGetProvidersOutput)
	return output, nil
}

// AccountSSOAutoJoinPolicyResource is the AccountSSOAutoJoinPolicyResource resource client
type AccountSSOAutoJoinPolicyResource struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	create    *clientruntime.Operation
	list      *clientruntime.Operation
	enable    *clientruntime.Operation
}

func newAccountSSOAutoJoinPolicyResource(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*AccountSSOAutoJoinPolicyResource, error) {
	res := finder.FindResource("AutoJoinPolicy")
	r := &AccountSSOAutoJoinPolicyResource{
		transport: transport,
		res:       res,
	}
	r.create = res.FindOperation("Create")
	r.list = res.FindOperation("List")
	r.enable = res.FindOperation("Enable")
	return r, nil
}

// Create - Creates an auto-join policy in the account
func (res *AccountSSOAutoJoinPolicyResource) Create(ctx context.Context, input *AccountSSOAutoJoinPolicyCreateInput) (*AccountSSOAutoJoinPolicyCreateOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.create,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*AccountSSOAutoJoinPolicyCreateOutput)
	return output, nil
}

// List - Lists policies in the account that allow users in other accounts to join automatically
// if they match certain criteria
func (res *AccountSSOAutoJoinPolicyResource) List(ctx context.Context, input *AccountSSOAutoJoinPolicyListInput) (*AccountSSOAutoJoinPolicyListOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.list,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*AccountSSOAutoJoinPolicyListOutput)
	return output, nil
}

// Enable - Enables the auto-join policy
func (res *AccountSSOAutoJoinPolicyResource) Enable(ctx context.Context, input *AccountSSOAutoJoinPolicyEnableInput) (*AccountSSOAutoJoinPolicyEnableOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.enable,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*AccountSSOAutoJoinPolicyEnableOutput)
	return output, nil
}

// UserResource is the UserResource resource client
type UserResource struct {
	transport      clientruntime.Transport
	res            *clientruntime.Resource
	AccessKey      *UserAccessKeyResource
	IdentityPolicy *UserIdentityPolicyResource
	create         *clientruntime.Operation
	get            *clientruntime.Operation
	destroy        *clientruntime.Operation
	list           *clientruntime.Operation
	memberAccounts *clientruntime.Operation
}

func newUserResource(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*UserResource, error) {
	res := finder.FindResource("User")
	r := &UserResource{
		transport: transport,
		res:       res,
	}
	var err error
	r.AccessKey, err = newUserAccessKeyResource(transport, res)
	if err != nil {
		return nil, err
	}
	r.IdentityPolicy, err = newUserIdentityPolicyResource(transport, res)
	if err != nil {
		return nil, err
	}
	r.create = res.FindOperation("Create")
	r.get = res.FindOperation("Get")
	r.destroy = res.FindOperation("Destroy")
	r.list = res.FindOperation("List")
	r.memberAccounts = res.FindOperation("MemberAccounts")
	return r, nil
}

// Create - Creates a new user with the given username in the current account
// Requires permission action iam:CreateUser over resource iam:User(<username>)
func (res *UserResource) Create(ctx context.Context, input *UserCreateInput) (*UserCreateOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.create,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*UserCreateOutput)
	return output, nil
}

// Get - Returns information about the current user or the user with the given username
func (res *UserResource) Get(ctx context.Context, input *UserGetInput) (*UserGetOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.get,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*UserGetOutput)
	return output, nil
}

// Destroy - Destroys the user with the given username
func (res *UserResource) Destroy(ctx context.Context, input *UserDestroyInput) (*UserDestroyOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.destroy,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*UserDestroyOutput)
	return output, nil
}

// List - Returns the list of users in the current account
func (res *UserResource) List(ctx context.Context, input *UserListInput) (*UserListOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.list,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*UserListOutput)
	return output, nil
}

// MemberAccounts - Returns the accounts the user is member of
// this operation is available only to User identities
func (res *UserResource) MemberAccounts(ctx context.Context, input *UserMemberAccountsInput) (*UserMemberAccountsOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.memberAccounts,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*UserMemberAccountsOutput)
	return output, nil
}

// UserAccessKeyResource is the UserAccessKeyResource resource client
type UserAccessKeyResource struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	create    *clientruntime.Operation
	list      *clientruntime.Operation
	destroy   *clientruntime.Operation
}

func newUserAccessKeyResource(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*UserAccessKeyResource, error) {
	res := finder.FindResource("AccessKey")
	r := &UserAccessKeyResource{
		transport: transport,
		res:       res,
	}
	r.create = res.FindOperation("Create")
	r.list = res.FindOperation("List")
	r.destroy = res.FindOperation("Destroy")
	return r, nil
}

// Create - Creates a new access key for the given user
// Requires permission action iam:CreateUserAccessKey over resource iam:User(<username>).AccessKey
func (res *UserAccessKeyResource) Create(ctx context.Context, input *UserAccessKeyCreateInput) (*UserAccessKeyCreateOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.create,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*UserAccessKeyCreateOutput)
	return output, nil
}

// List - Returns the list of access keys for the given user or the current user
// Requires permission action iam:ListUserAccessKeys over resource iam:User(<username>)
func (res *UserAccessKeyResource) List(ctx context.Context, input *UserAccessKeyListInput) (*UserAccessKeyListOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.list,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*UserAccessKeyListOutput)
	return output, nil
}

// Destroy - Destroy an access key for the given user
// Requires permission action iam:DeleteUserAccessKey over resource iam:User(<username>).AccessKey(<accessKeyID>)
func (res *UserAccessKeyResource) Destroy(ctx context.Context, input *UserAccessKeyDestroyInput) (*UserAccessKeyDestroyOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.destroy,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*UserAccessKeyDestroyOutput)
	return output, nil
}

// UserIdentityPolicyResource is the UserIdentityPolicyResource resource client
type UserIdentityPolicyResource struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	attach    *clientruntime.Operation
	list      *clientruntime.Operation
	detach    *clientruntime.Operation
}

func newUserIdentityPolicyResource(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*UserIdentityPolicyResource, error) {
	res := finder.FindResource("IdentityPolicy")
	r := &UserIdentityPolicyResource{
		transport: transport,
		res:       res,
	}
	r.attach = res.FindOperation("Attach")
	r.list = res.FindOperation("List")
	r.detach = res.FindOperation("Detach")
	return r, nil
}

// Attach - Attaches an identity policy to a user
// Requires permission action iam:AttachUserIdentityPolicy over resource iam:User(<username>).IdentityPolicy(<name>)
func (res *UserIdentityPolicyResource) Attach(ctx context.Context, input *UserIdentityPolicyAttachInput) (*UserIdentityPolicyAttachOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.attach,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*UserIdentityPolicyAttachOutput)
	return output, nil
}

// List - List returns the identity policies of a user
// Requires permission action iam:ListUserPolicies over resource iam:User(<username>).IdentityPolicy
func (res *UserIdentityPolicyResource) List(ctx context.Context, input *UserIdentityPolicyListInput) (*UserIdentityPolicyListOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.list,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*UserIdentityPolicyListOutput)
	return output, nil
}

// Detach - Deatach an identity policy from a user
// Requires permission action iam:DeattachUserIdentityPolicy over resource iam:User(<username>).IdentityPolicy(<name>)
func (res *UserIdentityPolicyResource) Detach(ctx context.Context, input *UserIdentityPolicyDetachInput) (*UserIdentityPolicyDetachOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.detach,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*UserIdentityPolicyDetachOutput)
	return output, nil
}

// RoleResource is the RoleResource resource client
type RoleResource struct {
	transport      clientruntime.Transport
	res            *clientruntime.Resource
	IdentityPolicy *RoleIdentityPolicyResource
	create         *clientruntime.Operation
	get            *clientruntime.Operation
	destroy        *clientruntime.Operation
	list           *clientruntime.Operation
	assume         *clientruntime.Operation
}

func newRoleResource(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*RoleResource, error) {
	res := finder.FindResource("Role")
	r := &RoleResource{
		transport: transport,
		res:       res,
	}
	var err error
	r.IdentityPolicy, err = newRoleIdentityPolicyResource(transport, res)
	if err != nil {
		return nil, err
	}
	r.create = res.FindOperation("Create")
	r.get = res.FindOperation("Get")
	r.destroy = res.FindOperation("Destroy")
	r.list = res.FindOperation("List")
	r.assume = res.FindOperation("Assume")
	return r, nil
}

// Create - Creates a new role with the given name in the current account
// Requires permission action iam:CreateRole over resource iam:Role(<name>)
func (res *RoleResource) Create(ctx context.Context, input *RoleCreateInput) (*RoleCreateOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.create,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*RoleCreateOutput)
	return output, nil
}

// Get - Returns information about a role
func (res *RoleResource) Get(ctx context.Context, input *RoleGetInput) (*RoleGetOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.get,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*RoleGetOutput)
	return output, nil
}

// Destroy - Destroys the role with the given name
func (res *RoleResource) Destroy(ctx context.Context, input *RoleDestroyInput) (*RoleDestroyOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.destroy,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*RoleDestroyOutput)
	return output, nil
}

// List - Returns the list of roles in the current account
func (res *RoleResource) List(ctx context.Context, input *RoleListInput) (*RoleListOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.list,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*RoleListOutput)
	return output, nil
}

// Assume - Returns credentials for the given assumed role
// the caller needs to have permission of action iam:AssumeRole over resource iam:Role(<roleName>)
func (res *RoleResource) Assume(ctx context.Context, input *RoleAssumeInput) (*RoleAssumeOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.assume,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*RoleAssumeOutput)
	return output, nil
}

// RoleIdentityPolicyResource is the RoleIdentityPolicyResource resource client
type RoleIdentityPolicyResource struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	attach    *clientruntime.Operation
	list      *clientruntime.Operation
	detach    *clientruntime.Operation
}

func newRoleIdentityPolicyResource(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*RoleIdentityPolicyResource, error) {
	res := finder.FindResource("IdentityPolicy")
	r := &RoleIdentityPolicyResource{
		transport: transport,
		res:       res,
	}
	r.attach = res.FindOperation("Attach")
	r.list = res.FindOperation("List")
	r.detach = res.FindOperation("Detach")
	return r, nil
}

// Attach - Attaches an identity policy to a role
// Requires permission action iam:AttachRoleIdentityPolicy over resource iam:Role(<name>).IdentityPolicy(<name>)
func (res *RoleIdentityPolicyResource) Attach(ctx context.Context, input *RoleIdentityPolicyAttachInput) (*RoleIdentityPolicyAttachOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.attach,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*RoleIdentityPolicyAttachOutput)
	return output, nil
}

// List - List returns the identity policies of a role
// Requires permission action iam:ListRolePolicies over resource iam:Role(<rolename>).IdentityPolicy
func (res *RoleIdentityPolicyResource) List(ctx context.Context, input *RoleIdentityPolicyListInput) (*RoleIdentityPolicyListOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.list,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*RoleIdentityPolicyListOutput)
	return output, nil
}

// Detach - Deatach an identity policy from a role
// Requires permission action iam:DeattachRoleIdentityPolicy over resource iam:Role(<rolename>).IdentityPolicy(<name>)
func (res *RoleIdentityPolicyResource) Detach(ctx context.Context, input *RoleIdentityPolicyDetachInput) (*RoleIdentityPolicyDetachOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.detach,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*RoleIdentityPolicyDetachOutput)
	return output, nil
}

// IdentityPolicyResource is the IdentityPolicyResource resource client
type IdentityPolicyResource struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	create    *clientruntime.Operation
	list      *clientruntime.Operation
	retrieve  *clientruntime.Operation
	destroy   *clientruntime.Operation
	update    *clientruntime.Operation
}

func newIdentityPolicyResource(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*IdentityPolicyResource, error) {
	res := finder.FindResource("IdentityPolicy")
	r := &IdentityPolicyResource{
		transport: transport,
		res:       res,
	}
	r.create = res.FindOperation("Create")
	r.list = res.FindOperation("List")
	r.retrieve = res.FindOperation("Retrieve")
	r.destroy = res.FindOperation("Destroy")
	r.update = res.FindOperation("Update")
	return r, nil
}

// Create - Creates a new identity policy
// Requires permission action iam:CreateIdentityPolicy over resource iam:IdentityPolicy(<name>)
func (res *IdentityPolicyResource) Create(ctx context.Context, input *IdentityPolicyCreateInput) (*IdentityPolicyCreateOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.create,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*IdentityPolicyCreateOutput)
	return output, nil
}

// List - Retrieves list of identity policies
func (res *IdentityPolicyResource) List(ctx context.Context, input *IdentityPolicyListInput) (*IdentityPolicyListOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.list,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*IdentityPolicyListOutput)
	return output, nil
}

// Retrieve - Retrieves identity policy by name
func (res *IdentityPolicyResource) Retrieve(ctx context.Context, input *IdentityPolicyRetrieveInput) (*IdentityPolicyRetrieveOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.retrieve,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*IdentityPolicyRetrieveOutput)
	return output, nil
}

// Destroy - Destroys an identity policy
// Requires permission action iam:DestroyIdentityPolicy over resource iam:IdentityPolicy(<name>)
func (res *IdentityPolicyResource) Destroy(ctx context.Context, input *IdentityPolicyDestroyInput) (*IdentityPolicyDestroyOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.destroy,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*IdentityPolicyDestroyOutput)
	return output, nil
}

// Update - Updates an identity policy
// Requires permission action iam:UpdateIdentityPolicy over resource iam:IdentityPolicy(<name>)
func (res *IdentityPolicyResource) Update(ctx context.Context, input *IdentityPolicyUpdateInput) (*IdentityPolicyUpdateOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.update,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*IdentityPolicyUpdateOutput)
	return output, nil
}

// ServiceBearerTokenResource is the ServiceBearerTokenResource resource client
type ServiceBearerTokenResource struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	get       *clientruntime.Operation
}

func newServiceBearerTokenResource(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*ServiceBearerTokenResource, error) {
	res := finder.FindResource("ServiceBearerToken")
	r := &ServiceBearerTokenResource{
		transport: transport,
		res:       res,
	}
	r.get = res.FindOperation("Get")
	return r, nil
}

// Get - Generates a temporary authorization token for accessing a a service or feature in a service that only supports bearer tokens
// The operations the token are limited to IAM permissions
// Requires permission to execute "iam:GetServiceBearerToken" over "*"
func (res *ServiceBearerTokenResource) Get(ctx context.Context, input *ServiceBearerTokenGetInput) (*ServiceBearerTokenGetOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.get,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*ServiceBearerTokenGetOutput)
	return output, nil
}

// InternalResource is the InternalResource resource client
type InternalResource struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	Services  *InternalServicesResource
}

func newInternalResource(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*InternalResource, error) {
	res := finder.FindResource("Internal")
	r := &InternalResource{
		transport: transport,
		res:       res,
	}
	var err error
	r.Services, err = newInternalServicesResource(transport, res)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// InternalServicesResource is the InternalServicesResource resource client
type InternalServicesResource struct {
	transport                  clientruntime.Transport
	res                        *clientruntime.Resource
	validateAccessKey          *clientruntime.Operation
	validateServiceBearerToken *clientruntime.Operation
	assertAction               *clientruntime.Operation
	assertQuery                *clientruntime.Operation
}

func newInternalServicesResource(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*InternalServicesResource, error) {
	res := finder.FindResource("Services")
	r := &InternalServicesResource{
		transport: transport,
		res:       res,
	}
	r.validateAccessKey = res.FindOperation("ValidateAccessKey")
	r.validateServiceBearerToken = res.FindOperation("ValidateServiceBearerToken")
	r.assertAction = res.FindOperation("AssertAction")
	r.assertQuery = res.FindOperation("AssertQuery")
	return r, nil
}

// ValidateAccessKey operation
func (res *InternalServicesResource) ValidateAccessKey(ctx context.Context, input *InternalServicesValidateAccessKeyInput) (*InternalServicesValidateAccessKeyOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.validateAccessKey,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*InternalServicesValidateAccessKeyOutput)
	return output, nil
}

// ValidateServiceBearerToken operation
func (res *InternalServicesResource) ValidateServiceBearerToken(ctx context.Context, input *InternalServicesValidateServiceBearerTokenInput) (*InternalServicesValidateServiceBearerTokenOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.validateServiceBearerToken,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*InternalServicesValidateServiceBearerTokenOutput)
	return output, nil
}

// AssertAction operation
func (res *InternalServicesResource) AssertAction(ctx context.Context, input *InternalServicesAssertActionInput) (*InternalServicesAssertActionOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.assertAction,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*InternalServicesAssertActionOutput)
	return output, nil
}

// AssertQuery operation
func (res *InternalServicesResource) AssertQuery(ctx context.Context, input *InternalServicesAssertQueryInput) (*InternalServicesAssertQueryOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.assertQuery,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*InternalServicesAssertQueryOutput)
	return output, nil
}

// Client is the main client of the API
type Client struct {
	transport clientruntime.Transport
	pk        *clientruntime.Package
	// Account operations
	Account *AccountResource
	// User operations
	User *UserResource
	// Role operations
	Role *RoleResource
	// IdentityPolicy operations
	IdentityPolicy *IdentityPolicyResource
	// ServiceBearerToken - Service Bearer Tokens
	ServiceBearerToken *ServiceBearerTokenResource
	// Internal - Internal operations exposed only to integrating services
	Internal *InternalResource
}

// WithTransport configures the transport in the client
func WithTransport(transport clientruntime.Transport) clientruntime.OptionFunc {
	return func(o any) error {
		c, ok := o.(*Client)
		if !ok {
			return nil
		}
		c.transport = transport
		return nil
	}
}

// NewEndpointTransport returns a transport with the package base endpoint
func NewEndpointTransport(options ...clientruntime.Option) (clientruntime.Transport, error) {
	o := options
	if mo, err := godeployportcomapiservicescorelibconfigurator.DefaultClientOptions(options...); err != nil {
		return nil, err
	} else {
		o = append(options, mo...)
	}
	return clientruntime.NewHTTPJSONTransport(
		"http://localhost:3000",
		o...,
	)
}

// NewClient returns a new instance of Client
func NewClient(options ...clientruntime.Option) (*Client, error) {
	pk := SpecularMeta().Module()
	c := &Client{
		pk: pk,
	}
	if err := clientruntime.ApplyOptions(c, options...); err != nil {
		return nil, err
	}
	if c.transport == nil {
		t, err := NewEndpointTransport(options...)
		if err != nil {
			return nil, err
		}
		c.transport = t
	}
	transport := c.transport
	var err error
	c.Account, err = newAccountResource(transport, pk)
	if err != nil {
		return nil, err
	}
	c.User, err = newUserResource(transport, pk)
	if err != nil {
		return nil, err
	}
	c.Role, err = newRoleResource(transport, pk)
	if err != nil {
		return nil, err
	}
	c.IdentityPolicy, err = newIdentityPolicyResource(transport, pk)
	if err != nil {
		return nil, err
	}
	c.ServiceBearerToken, err = newServiceBearerTokenResource(transport, pk)
	if err != nil {
		return nil, err
	}
	c.Internal, err = newInternalResource(transport, pk)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func init() {
	initSpecularMeta()
}

func initSpecularMeta() {
	pk, err := newSpecularPackage()
	if err != nil {
		panic(errors.New("failed to initialize shared allow package deployport/iam"))
	}
	localSpecularMeta.mod = pk
}

// SpecularMetaInfo defines metadata of the specular module
type SpecularMetaInfo struct {
	mod                                                                         *clientruntime.Package
	structPathUserInformationSSOProvider                                        *clientruntime.StructDefinition
	structPathUserInformationSSOProfile                                         *clientruntime.StructDefinition
	structPathUserInformationSSO                                                *clientruntime.StructDefinition
	structPathUserInformation                                                   *clientruntime.StructDefinition
	structPathRoleInformation                                                   *clientruntime.StructDefinition
	structPathCredentials                                                       *clientruntime.StructDefinition
	structPathSSOProviderUnavailableProblem                                     *clientruntime.StructDefinition
	structPathSSOFlow                                                           *clientruntime.StructDefinition
	structPathAccount                                                           *clientruntime.StructDefinition
	structPathAccountSSOProvider                                                *clientruntime.StructDefinition
	structPathPolicyNotFoundProblem                                             *clientruntime.StructDefinition
	structPathAccountSSOAutoJoinPolicy                                          *clientruntime.StructDefinition
	structPathAccountCreateInput                                                *clientruntime.StructDefinition
	structPathAccountCreateOutput                                               *clientruntime.StructDefinition
	structPathAccountCreateInvalidNameProblem                                   *clientruntime.StructDefinition
	structPathAccountAssumeIdentityInput                                        *clientruntime.StructDefinition
	structPathAccountAssumeIdentityOutput                                       *clientruntime.StructDefinition
	structPathAccountSSOBeginAuthenticationInput                                *clientruntime.StructDefinition
	structPathAccountSSOBeginAuthenticationOutput                               *clientruntime.StructDefinition
	structPathAccountSSOBeginAuthenticationParameterProblem                     *clientruntime.StructDefinition
	structPathAccountSSOCompleteAuthenticationInput                             *clientruntime.StructDefinition
	structPathAccountSSOCompleteAuthenticationOutput                            *clientruntime.StructDefinition
	structPathAccountSSOCompleteAuthenticationInvalidFlowProblem                *clientruntime.StructDefinition
	structPathAccountSSOGetProvidersInput                                       *clientruntime.StructDefinition
	structPathAccountSSOGetProvidersOutput                                      *clientruntime.StructDefinition
	structPathAccountSSOAutoJoinPolicyCreateInput                               *clientruntime.StructDefinition
	structPathAccountSSOAutoJoinPolicyCreateOutput                              *clientruntime.StructDefinition
	structPathAccountSSOAutoJoinPolicyListInput                                 *clientruntime.StructDefinition
	structPathAccountSSOAutoJoinPolicyListOutput                                *clientruntime.StructDefinition
	structPathAccountSSOAutoJoinPolicyEnableInput                               *clientruntime.StructDefinition
	structPathAccountSSOAutoJoinPolicyEnableOutput                              *clientruntime.StructDefinition
	structPathInvalidUsernameProblem                                            *clientruntime.StructDefinition
	structPathInvalidRoleNameProblem                                            *clientruntime.StructDefinition
	structPathMemberAccount                                                     *clientruntime.StructDefinition
	structPathCredentialInfo                                                    *clientruntime.StructDefinition
	structPathIdentityPolicyStatement                                           *clientruntime.StructDefinition
	structPathIdentityPolicy                                                    *clientruntime.StructDefinition
	structPathIdentityPolicyAttachment                                          *clientruntime.StructDefinition
	structPathUserNotFoundProblem                                               *clientruntime.StructDefinition
	structPathRoleNotFoundProblem                                               *clientruntime.StructDefinition
	structPathIdentityInUseProblem                                              *clientruntime.StructDefinition
	structPathUserCreateInput                                                   *clientruntime.StructDefinition
	structPathUserCreateOutput                                                  *clientruntime.StructDefinition
	structPathUserGetInput                                                      *clientruntime.StructDefinition
	structPathUserGetOutput                                                     *clientruntime.StructDefinition
	structPathUserGetUserNotAvailableProblem                                    *clientruntime.StructDefinition
	structPathUserDestroyInput                                                  *clientruntime.StructDefinition
	structPathUserDestroyOutput                                                 *clientruntime.StructDefinition
	structPathUserListInput                                                     *clientruntime.StructDefinition
	structPathUserListOutput                                                    *clientruntime.StructDefinition
	structPathUserMemberAccountsInput                                           *clientruntime.StructDefinition
	structPathUserMemberAccountsOutput                                          *clientruntime.StructDefinition
	structPathUserAccessKeyCreateInput                                          *clientruntime.StructDefinition
	structPathUserAccessKeyCreateOutput                                         *clientruntime.StructDefinition
	structPathUserAccessKeyListInput                                            *clientruntime.StructDefinition
	structPathUserAccessKeyListOutput                                           *clientruntime.StructDefinition
	structPathUserAccessKeyDestroyInput                                         *clientruntime.StructDefinition
	structPathUserAccessKeyDestroyOutput                                        *clientruntime.StructDefinition
	structPathUserIdentityPolicyAttachInput                                     *clientruntime.StructDefinition
	structPathUserIdentityPolicyAttachOutput                                    *clientruntime.StructDefinition
	structPathUserIdentityPolicyListInput                                       *clientruntime.StructDefinition
	structPathUserIdentityPolicyListOutput                                      *clientruntime.StructDefinition
	structPathUserIdentityPolicyDetachInput                                     *clientruntime.StructDefinition
	structPathUserIdentityPolicyDetachOutput                                    *clientruntime.StructDefinition
	structPathInlinePolicy                                                      *clientruntime.StructDefinition
	structPathPolicyStructureProblem                                            *clientruntime.StructDefinition
	structPathRoleCreateInput                                                   *clientruntime.StructDefinition
	structPathRoleCreateOutput                                                  *clientruntime.StructDefinition
	structPathRoleGetInput                                                      *clientruntime.StructDefinition
	structPathRoleGetOutput                                                     *clientruntime.StructDefinition
	structPathRoleDestroyInput                                                  *clientruntime.StructDefinition
	structPathRoleDestroyOutput                                                 *clientruntime.StructDefinition
	structPathRoleListInput                                                     *clientruntime.StructDefinition
	structPathRoleListOutput                                                    *clientruntime.StructDefinition
	structPathRoleAssumeInput                                                   *clientruntime.StructDefinition
	structPathRoleAssumeOutput                                                  *clientruntime.StructDefinition
	structPathRoleIdentityPolicyAttachInput                                     *clientruntime.StructDefinition
	structPathRoleIdentityPolicyAttachOutput                                    *clientruntime.StructDefinition
	structPathRoleIdentityPolicyListInput                                       *clientruntime.StructDefinition
	structPathRoleIdentityPolicyListOutput                                      *clientruntime.StructDefinition
	structPathRoleIdentityPolicyDetachInput                                     *clientruntime.StructDefinition
	structPathRoleIdentityPolicyDetachOutput                                    *clientruntime.StructDefinition
	structPathIdentityPolicyCreateInput                                         *clientruntime.StructDefinition
	structPathIdentityPolicyCreateOutput                                        *clientruntime.StructDefinition
	structPathIdentityPolicyListInput                                           *clientruntime.StructDefinition
	structPathIdentityPolicyListOutput                                          *clientruntime.StructDefinition
	structPathIdentityPolicyRetrieveInput                                       *clientruntime.StructDefinition
	structPathIdentityPolicyRetrieveOutput                                      *clientruntime.StructDefinition
	structPathIdentityPolicyDestroyInput                                        *clientruntime.StructDefinition
	structPathIdentityPolicyDestroyOutput                                       *clientruntime.StructDefinition
	structPathIdentityPolicyUpdateInput                                         *clientruntime.StructDefinition
	structPathIdentityPolicyUpdateOutput                                        *clientruntime.StructDefinition
	structPathServiceBearerToken                                                *clientruntime.StructDefinition
	structPathInvalidServiceBearerTokenDurationProblem                          *clientruntime.StructDefinition
	structPathInvalidServiceNameProblem                                         *clientruntime.StructDefinition
	structPathServiceBearerTokenGetInput                                        *clientruntime.StructDefinition
	structPathServiceBearerTokenGetOutput                                       *clientruntime.StructDefinition
	structPathInternalAccessKeyUser                                             *clientruntime.StructDefinition
	structPathInternalAccessKeyRole                                             *clientruntime.StructDefinition
	structPathInternalAccessKeyAccount                                          *clientruntime.StructDefinition
	structPathInternalAccessKey                                                 *clientruntime.StructDefinition
	structPathInternalQueryAssertionEntryValue                                  *clientruntime.StructDefinition
	structPathInternalQueryAssertionEntry                                       *clientruntime.StructDefinition
	structPathInternalServicesAssertActionCallerForbiddenProblem                *clientruntime.StructDefinition
	structPathInternalAssertActionQualifier                                     *clientruntime.StructDefinition
	structPathInternalServicesValidateAccessKeyInput                            *clientruntime.StructDefinition
	structPathInternalServicesValidateAccessKeyOutput                           *clientruntime.StructDefinition
	structPathInternalServicesValidateAccessKeyInvalidAccessKeyProblem          *clientruntime.StructDefinition
	structPathInternalServicesValidateServiceBearerTokenInput                   *clientruntime.StructDefinition
	structPathInternalServicesValidateServiceBearerTokenOutput                  *clientruntime.StructDefinition
	structPathInternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem *clientruntime.StructDefinition
	structPathInternalServicesAssertActionInput                                 *clientruntime.StructDefinition
	structPathInternalServicesAssertActionOutput                                *clientruntime.StructDefinition
	structPathInternalServicesAssertQueryInput                                  *clientruntime.StructDefinition
	structPathInternalServicesAssertQueryOutput                                 *clientruntime.StructDefinition
	structPathInternalServicesAssertQueryReplacementProblem                     *clientruntime.StructDefinition
}

// Module returns the module definition
func (m *SpecularMetaInfo) Module() *clientruntime.Package {
	return m.mod
}

// UserInformationSSOProviderStruct allows easy access to structure
func (m *SpecularMetaInfo) UserInformationSSOProviderStruct() *clientruntime.StructDefinition {
	return m.structPathUserInformationSSOProvider
}

// UserInformationSSOProfileStruct allows easy access to structure
func (m *SpecularMetaInfo) UserInformationSSOProfileStruct() *clientruntime.StructDefinition {
	return m.structPathUserInformationSSOProfile
}

// UserInformationSSOStruct allows easy access to structure
func (m *SpecularMetaInfo) UserInformationSSOStruct() *clientruntime.StructDefinition {
	return m.structPathUserInformationSSO
}

// UserInformationStruct allows easy access to structure
func (m *SpecularMetaInfo) UserInformationStruct() *clientruntime.StructDefinition {
	return m.structPathUserInformation
}

// RoleInformationStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleInformationStruct() *clientruntime.StructDefinition {
	return m.structPathRoleInformation
}

// CredentialsStruct allows easy access to structure
func (m *SpecularMetaInfo) CredentialsStruct() *clientruntime.StructDefinition {
	return m.structPathCredentials
}

// SSOProviderUnavailableProblemStruct allows easy access to structure
func (m *SpecularMetaInfo) SSOProviderUnavailableProblemStruct() *clientruntime.StructDefinition {
	return m.structPathSSOProviderUnavailableProblem
}

// SSOFlowStruct allows easy access to structure
func (m *SpecularMetaInfo) SSOFlowStruct() *clientruntime.StructDefinition {
	return m.structPathSSOFlow
}

// AccountStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountStruct() *clientruntime.StructDefinition {
	return m.structPathAccount
}

// AccountSSOProviderStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOProviderStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOProvider
}

// PolicyNotFoundProblemStruct allows easy access to structure
func (m *SpecularMetaInfo) PolicyNotFoundProblemStruct() *clientruntime.StructDefinition {
	return m.structPathPolicyNotFoundProblem
}

// AccountSSOAutoJoinPolicyStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOAutoJoinPolicyStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOAutoJoinPolicy
}

// AccountCreateInputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountCreateInputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountCreateInput
}

// AccountCreateOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountCreateOutputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountCreateOutput
}

// AccountCreateInvalidNameProblemStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountCreateInvalidNameProblemStruct() *clientruntime.StructDefinition {
	return m.structPathAccountCreateInvalidNameProblem
}

// AccountAssumeIdentityInputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountAssumeIdentityInputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountAssumeIdentityInput
}

// AccountAssumeIdentityOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountAssumeIdentityOutputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountAssumeIdentityOutput
}

// AccountSSOBeginAuthenticationInputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOBeginAuthenticationInputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOBeginAuthenticationInput
}

// AccountSSOBeginAuthenticationOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOBeginAuthenticationOutputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOBeginAuthenticationOutput
}

// AccountSSOBeginAuthenticationParameterProblemStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOBeginAuthenticationParameterProblemStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOBeginAuthenticationParameterProblem
}

// AccountSSOCompleteAuthenticationInputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOCompleteAuthenticationInputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOCompleteAuthenticationInput
}

// AccountSSOCompleteAuthenticationOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOCompleteAuthenticationOutputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOCompleteAuthenticationOutput
}

// AccountSSOCompleteAuthenticationInvalidFlowProblemStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOCompleteAuthenticationInvalidFlowProblemStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOCompleteAuthenticationInvalidFlowProblem
}

// AccountSSOGetProvidersInputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOGetProvidersInputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOGetProvidersInput
}

// AccountSSOGetProvidersOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOGetProvidersOutputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOGetProvidersOutput
}

// AccountSSOAutoJoinPolicyCreateInputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOAutoJoinPolicyCreateInputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOAutoJoinPolicyCreateInput
}

// AccountSSOAutoJoinPolicyCreateOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOAutoJoinPolicyCreateOutputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOAutoJoinPolicyCreateOutput
}

// AccountSSOAutoJoinPolicyListInputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOAutoJoinPolicyListInputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOAutoJoinPolicyListInput
}

// AccountSSOAutoJoinPolicyListOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOAutoJoinPolicyListOutputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOAutoJoinPolicyListOutput
}

// AccountSSOAutoJoinPolicyEnableInputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOAutoJoinPolicyEnableInputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOAutoJoinPolicyEnableInput
}

// AccountSSOAutoJoinPolicyEnableOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOAutoJoinPolicyEnableOutputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOAutoJoinPolicyEnableOutput
}

// InvalidUsernameProblemStruct allows easy access to structure
func (m *SpecularMetaInfo) InvalidUsernameProblemStruct() *clientruntime.StructDefinition {
	return m.structPathInvalidUsernameProblem
}

// InvalidRoleNameProblemStruct allows easy access to structure
func (m *SpecularMetaInfo) InvalidRoleNameProblemStruct() *clientruntime.StructDefinition {
	return m.structPathInvalidRoleNameProblem
}

// MemberAccountStruct allows easy access to structure
func (m *SpecularMetaInfo) MemberAccountStruct() *clientruntime.StructDefinition {
	return m.structPathMemberAccount
}

// CredentialInfoStruct allows easy access to structure
func (m *SpecularMetaInfo) CredentialInfoStruct() *clientruntime.StructDefinition {
	return m.structPathCredentialInfo
}

// IdentityPolicyStatementStruct allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyStatementStruct() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyStatement
}

// IdentityPolicyStruct allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyStruct() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicy
}

// IdentityPolicyAttachmentStruct allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyAttachmentStruct() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyAttachment
}

// UserNotFoundProblemStruct allows easy access to structure
func (m *SpecularMetaInfo) UserNotFoundProblemStruct() *clientruntime.StructDefinition {
	return m.structPathUserNotFoundProblem
}

// RoleNotFoundProblemStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleNotFoundProblemStruct() *clientruntime.StructDefinition {
	return m.structPathRoleNotFoundProblem
}

// IdentityInUseProblemStruct allows easy access to structure
func (m *SpecularMetaInfo) IdentityInUseProblemStruct() *clientruntime.StructDefinition {
	return m.structPathIdentityInUseProblem
}

// UserCreateInputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserCreateInputStruct() *clientruntime.StructDefinition {
	return m.structPathUserCreateInput
}

// UserCreateOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserCreateOutputStruct() *clientruntime.StructDefinition {
	return m.structPathUserCreateOutput
}

// UserGetInputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserGetInputStruct() *clientruntime.StructDefinition {
	return m.structPathUserGetInput
}

// UserGetOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserGetOutputStruct() *clientruntime.StructDefinition {
	return m.structPathUserGetOutput
}

// UserGetUserNotAvailableProblemStruct allows easy access to structure
func (m *SpecularMetaInfo) UserGetUserNotAvailableProblemStruct() *clientruntime.StructDefinition {
	return m.structPathUserGetUserNotAvailableProblem
}

// UserDestroyInputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserDestroyInputStruct() *clientruntime.StructDefinition {
	return m.structPathUserDestroyInput
}

// UserDestroyOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserDestroyOutputStruct() *clientruntime.StructDefinition {
	return m.structPathUserDestroyOutput
}

// UserListInputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserListInputStruct() *clientruntime.StructDefinition {
	return m.structPathUserListInput
}

// UserListOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserListOutputStruct() *clientruntime.StructDefinition {
	return m.structPathUserListOutput
}

// UserMemberAccountsInputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserMemberAccountsInputStruct() *clientruntime.StructDefinition {
	return m.structPathUserMemberAccountsInput
}

// UserMemberAccountsOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserMemberAccountsOutputStruct() *clientruntime.StructDefinition {
	return m.structPathUserMemberAccountsOutput
}

// UserAccessKeyCreateInputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserAccessKeyCreateInputStruct() *clientruntime.StructDefinition {
	return m.structPathUserAccessKeyCreateInput
}

// UserAccessKeyCreateOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserAccessKeyCreateOutputStruct() *clientruntime.StructDefinition {
	return m.structPathUserAccessKeyCreateOutput
}

// UserAccessKeyListInputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserAccessKeyListInputStruct() *clientruntime.StructDefinition {
	return m.structPathUserAccessKeyListInput
}

// UserAccessKeyListOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserAccessKeyListOutputStruct() *clientruntime.StructDefinition {
	return m.structPathUserAccessKeyListOutput
}

// UserAccessKeyDestroyInputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserAccessKeyDestroyInputStruct() *clientruntime.StructDefinition {
	return m.structPathUserAccessKeyDestroyInput
}

// UserAccessKeyDestroyOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserAccessKeyDestroyOutputStruct() *clientruntime.StructDefinition {
	return m.structPathUserAccessKeyDestroyOutput
}

// UserIdentityPolicyAttachInputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserIdentityPolicyAttachInputStruct() *clientruntime.StructDefinition {
	return m.structPathUserIdentityPolicyAttachInput
}

// UserIdentityPolicyAttachOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserIdentityPolicyAttachOutputStruct() *clientruntime.StructDefinition {
	return m.structPathUserIdentityPolicyAttachOutput
}

// UserIdentityPolicyListInputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserIdentityPolicyListInputStruct() *clientruntime.StructDefinition {
	return m.structPathUserIdentityPolicyListInput
}

// UserIdentityPolicyListOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserIdentityPolicyListOutputStruct() *clientruntime.StructDefinition {
	return m.structPathUserIdentityPolicyListOutput
}

// UserIdentityPolicyDetachInputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserIdentityPolicyDetachInputStruct() *clientruntime.StructDefinition {
	return m.structPathUserIdentityPolicyDetachInput
}

// UserIdentityPolicyDetachOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserIdentityPolicyDetachOutputStruct() *clientruntime.StructDefinition {
	return m.structPathUserIdentityPolicyDetachOutput
}

// InlinePolicyStruct allows easy access to structure
func (m *SpecularMetaInfo) InlinePolicyStruct() *clientruntime.StructDefinition {
	return m.structPathInlinePolicy
}

// PolicyStructureProblemStruct allows easy access to structure
func (m *SpecularMetaInfo) PolicyStructureProblemStruct() *clientruntime.StructDefinition {
	return m.structPathPolicyStructureProblem
}

// RoleCreateInputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleCreateInputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleCreateInput
}

// RoleCreateOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleCreateOutputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleCreateOutput
}

// RoleGetInputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleGetInputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleGetInput
}

// RoleGetOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleGetOutputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleGetOutput
}

// RoleDestroyInputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleDestroyInputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleDestroyInput
}

// RoleDestroyOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleDestroyOutputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleDestroyOutput
}

// RoleListInputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleListInputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleListInput
}

// RoleListOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleListOutputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleListOutput
}

// RoleAssumeInputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleAssumeInputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleAssumeInput
}

// RoleAssumeOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleAssumeOutputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleAssumeOutput
}

// RoleIdentityPolicyAttachInputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleIdentityPolicyAttachInputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleIdentityPolicyAttachInput
}

// RoleIdentityPolicyAttachOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleIdentityPolicyAttachOutputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleIdentityPolicyAttachOutput
}

// RoleIdentityPolicyListInputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleIdentityPolicyListInputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleIdentityPolicyListInput
}

// RoleIdentityPolicyListOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleIdentityPolicyListOutputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleIdentityPolicyListOutput
}

// RoleIdentityPolicyDetachInputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleIdentityPolicyDetachInputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleIdentityPolicyDetachInput
}

// RoleIdentityPolicyDetachOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleIdentityPolicyDetachOutputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleIdentityPolicyDetachOutput
}

// IdentityPolicyCreateInputStruct allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyCreateInputStruct() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyCreateInput
}

// IdentityPolicyCreateOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyCreateOutputStruct() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyCreateOutput
}

// IdentityPolicyListInputStruct allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyListInputStruct() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyListInput
}

// IdentityPolicyListOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyListOutputStruct() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyListOutput
}

// IdentityPolicyRetrieveInputStruct allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyRetrieveInputStruct() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyRetrieveInput
}

// IdentityPolicyRetrieveOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyRetrieveOutputStruct() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyRetrieveOutput
}

// IdentityPolicyDestroyInputStruct allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyDestroyInputStruct() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyDestroyInput
}

// IdentityPolicyDestroyOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyDestroyOutputStruct() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyDestroyOutput
}

// IdentityPolicyUpdateInputStruct allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyUpdateInputStruct() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyUpdateInput
}

// IdentityPolicyUpdateOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyUpdateOutputStruct() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyUpdateOutput
}

// ServiceBearerTokenStruct allows easy access to structure
func (m *SpecularMetaInfo) ServiceBearerTokenStruct() *clientruntime.StructDefinition {
	return m.structPathServiceBearerToken
}

// InvalidServiceBearerTokenDurationProblemStruct allows easy access to structure
func (m *SpecularMetaInfo) InvalidServiceBearerTokenDurationProblemStruct() *clientruntime.StructDefinition {
	return m.structPathInvalidServiceBearerTokenDurationProblem
}

// InvalidServiceNameProblemStruct allows easy access to structure
func (m *SpecularMetaInfo) InvalidServiceNameProblemStruct() *clientruntime.StructDefinition {
	return m.structPathInvalidServiceNameProblem
}

// ServiceBearerTokenGetInputStruct allows easy access to structure
func (m *SpecularMetaInfo) ServiceBearerTokenGetInputStruct() *clientruntime.StructDefinition {
	return m.structPathServiceBearerTokenGetInput
}

// ServiceBearerTokenGetOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) ServiceBearerTokenGetOutputStruct() *clientruntime.StructDefinition {
	return m.structPathServiceBearerTokenGetOutput
}

// InternalAccessKeyUserStruct allows easy access to structure
func (m *SpecularMetaInfo) InternalAccessKeyUserStruct() *clientruntime.StructDefinition {
	return m.structPathInternalAccessKeyUser
}

// InternalAccessKeyRoleStruct allows easy access to structure
func (m *SpecularMetaInfo) InternalAccessKeyRoleStruct() *clientruntime.StructDefinition {
	return m.structPathInternalAccessKeyRole
}

// InternalAccessKeyAccountStruct allows easy access to structure
func (m *SpecularMetaInfo) InternalAccessKeyAccountStruct() *clientruntime.StructDefinition {
	return m.structPathInternalAccessKeyAccount
}

// InternalAccessKeyStruct allows easy access to structure
func (m *SpecularMetaInfo) InternalAccessKeyStruct() *clientruntime.StructDefinition {
	return m.structPathInternalAccessKey
}

// InternalQueryAssertionEntryValueStruct allows easy access to structure
func (m *SpecularMetaInfo) InternalQueryAssertionEntryValueStruct() *clientruntime.StructDefinition {
	return m.structPathInternalQueryAssertionEntryValue
}

// InternalQueryAssertionEntryStruct allows easy access to structure
func (m *SpecularMetaInfo) InternalQueryAssertionEntryStruct() *clientruntime.StructDefinition {
	return m.structPathInternalQueryAssertionEntry
}

// InternalServicesAssertActionCallerForbiddenProblemStruct allows easy access to structure
func (m *SpecularMetaInfo) InternalServicesAssertActionCallerForbiddenProblemStruct() *clientruntime.StructDefinition {
	return m.structPathInternalServicesAssertActionCallerForbiddenProblem
}

// InternalAssertActionQualifierStruct allows easy access to structure
func (m *SpecularMetaInfo) InternalAssertActionQualifierStruct() *clientruntime.StructDefinition {
	return m.structPathInternalAssertActionQualifier
}

// InternalServicesValidateAccessKeyInputStruct allows easy access to structure
func (m *SpecularMetaInfo) InternalServicesValidateAccessKeyInputStruct() *clientruntime.StructDefinition {
	return m.structPathInternalServicesValidateAccessKeyInput
}

// InternalServicesValidateAccessKeyOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) InternalServicesValidateAccessKeyOutputStruct() *clientruntime.StructDefinition {
	return m.structPathInternalServicesValidateAccessKeyOutput
}

// InternalServicesValidateAccessKeyInvalidAccessKeyProblemStruct allows easy access to structure
func (m *SpecularMetaInfo) InternalServicesValidateAccessKeyInvalidAccessKeyProblemStruct() *clientruntime.StructDefinition {
	return m.structPathInternalServicesValidateAccessKeyInvalidAccessKeyProblem
}

// InternalServicesValidateServiceBearerTokenInputStruct allows easy access to structure
func (m *SpecularMetaInfo) InternalServicesValidateServiceBearerTokenInputStruct() *clientruntime.StructDefinition {
	return m.structPathInternalServicesValidateServiceBearerTokenInput
}

// InternalServicesValidateServiceBearerTokenOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) InternalServicesValidateServiceBearerTokenOutputStruct() *clientruntime.StructDefinition {
	return m.structPathInternalServicesValidateServiceBearerTokenOutput
}

// InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblemStruct allows easy access to structure
func (m *SpecularMetaInfo) InternalServicesValidateServiceBearerTokenInvalidAccessKeyProblemStruct() *clientruntime.StructDefinition {
	return m.structPathInternalServicesValidateServiceBearerTokenInvalidAccessKeyProblem
}

// InternalServicesAssertActionInputStruct allows easy access to structure
func (m *SpecularMetaInfo) InternalServicesAssertActionInputStruct() *clientruntime.StructDefinition {
	return m.structPathInternalServicesAssertActionInput
}

// InternalServicesAssertActionOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) InternalServicesAssertActionOutputStruct() *clientruntime.StructDefinition {
	return m.structPathInternalServicesAssertActionOutput
}

// InternalServicesAssertQueryInputStruct allows easy access to structure
func (m *SpecularMetaInfo) InternalServicesAssertQueryInputStruct() *clientruntime.StructDefinition {
	return m.structPathInternalServicesAssertQueryInput
}

// InternalServicesAssertQueryOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) InternalServicesAssertQueryOutputStruct() *clientruntime.StructDefinition {
	return m.structPathInternalServicesAssertQueryOutput
}

// InternalServicesAssertQueryReplacementProblemStruct allows easy access to structure
func (m *SpecularMetaInfo) InternalServicesAssertQueryReplacementProblemStruct() *clientruntime.StructDefinition {
	return m.structPathInternalServicesAssertQueryReplacementProblem
}

var localSpecularMeta *SpecularMetaInfo = &SpecularMetaInfo{}

// SpecularMeta returns metadata of the specular module
func SpecularMeta() *SpecularMetaInfo {
	return localSpecularMeta
}
