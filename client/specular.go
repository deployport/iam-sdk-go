package iam

import (
	"context"
	"errors"
	"time"

	clientruntime "go.deployport.com/specular-runtime/client"

	godeployportcomapiservicescorelib "go.deployport.com/api-services-corelib"

	godeployportcomapiservicescorelibconfigurator "go.deployport.com/api-services-corelib/configurator"
)

// NewUserInformationSSOProvider creates a new UserInformationSSOProvider
func NewUserInformationSSOProvider() *UserInformationSSOProvider {
	s := &UserInformationSSOProvider{}
	return s
}

// UserInformationSSOProvider entity
type UserInformationSSOProvider struct {
	// provider name
	Name string
}

// GetName returns the value for the field name
func (e *UserInformationSSOProvider) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *UserInformationSSOProvider) SetName(name string) {
	e.Name = name
}

// Hydrate implements struct hydrate
func (e *UserInformationSSOProvider) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "name", &e.Name); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserInformationSSOProvider) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("name", e.Name)
	return nil
}

// StructPath returns StructPath
func (e *UserInformationSSOProvider) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserInformationSSOProvider.Path()
}

// NewUserInformationSSOProfile creates a new UserInformationSSOProfile
func NewUserInformationSSOProfile() *UserInformationSSOProfile {
	s := &UserInformationSSOProfile{}
	return s
}

// UserInformationSSOProfile entity
type UserInformationSSOProfile struct {
	Email      string
	Fullname   string
	PictureURL string
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

// Hydrate implements struct hydrate
func (e *UserInformationSSOProfile) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "email", &e.Email); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "fullname", &e.Fullname); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "pictureURL", &e.PictureURL); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserInformationSSOProfile) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("email", e.Email)
	ctx.Content().SetProperty("fullname", e.Fullname)
	ctx.Content().SetProperty("pictureURL", e.PictureURL)
	return nil
}

// StructPath returns StructPath
func (e *UserInformationSSOProfile) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserInformationSSOProfile.Path()
}

// NewUserInformationSSO creates a new UserInformationSSO
func NewUserInformationSSO() *UserInformationSSO {
	s := &UserInformationSSO{}
	return s
}

// UserInformationSSO entity
type UserInformationSSO struct {
	Profile  *UserInformationSSOProfile
	Provider *UserInformationSSOProvider
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

// Hydrate implements struct hydrate
func (e *UserInformationSSO) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentObjectProperty(ctx.Content(), SpecularMeta().UserInformationSSOProfile().TypeBuilder(), ctx.Package(), "profile", &e.Profile); err != nil {
		return err
	}
	if err := clientruntime.ContentObjectProperty(ctx.Content(), SpecularMeta().UserInformationSSOProvider().TypeBuilder(), ctx.Package(), "provider", &e.Provider); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserInformationSSO) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {

	var fieldContentProfile clientruntime.Content
	if v := e.Profile; v != nil {
		fieldContentProfile = clientruntime.NewContent()
		if err := v.Dehydrate(ctx.CopyWithContent(fieldContentProfile)); err != nil {
			return clientruntime.NewPropertyDehydrationError("profile", err)
		}
	}
	if fieldContentProfile == nil {
		return clientruntime.NewPropertyRequiredError("profile")
	}
	ctx.Content().SetProperty("profile", fieldContentProfile.Map())

	var fieldContentProvider clientruntime.Content
	if v := e.Provider; v != nil {
		fieldContentProvider = clientruntime.NewContent()
		if err := v.Dehydrate(ctx.CopyWithContent(fieldContentProvider)); err != nil {
			return clientruntime.NewPropertyDehydrationError("provider", err)
		}
	}
	if fieldContentProvider == nil {
		return clientruntime.NewPropertyRequiredError("provider")
	}
	ctx.Content().SetProperty("provider", fieldContentProvider.Map())
	return nil
}

// StructPath returns StructPath
func (e *UserInformationSSO) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserInformationSSO.Path()
}

// NewUserInformation creates a new UserInformation
func NewUserInformation() *UserInformation {
	s := &UserInformation{}
	return s
}

// UserInformation entity
type UserInformation struct {
	Description string
	// when the user comes from SSO, this field is populated with the extra information
	Sso      *UserInformationSSO
	Username string
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

// Hydrate implements struct hydrate
func (e *UserInformation) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "description", &e.Description); err != nil {
		return err
	}
	if err := clientruntime.ContentObjectProperty(ctx.Content(), SpecularMeta().UserInformationSSO().TypeBuilder(), ctx.Package(), "sso", &e.Sso); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "username", &e.Username); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserInformation) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("description", e.Description)

	var fieldContentSso clientruntime.Content
	if v := e.Sso; v != nil {
		fieldContentSso = clientruntime.NewContent()
		if err := v.Dehydrate(ctx.CopyWithContent(fieldContentSso)); err != nil {
			return clientruntime.NewPropertyDehydrationError("sso", err)
		}
	}
	ctx.Content().SetProperty("sso", fieldContentSso.Map())
	ctx.Content().SetProperty("username", e.Username)
	return nil
}

// StructPath returns StructPath
func (e *UserInformation) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserInformation.Path()
}

// NewCredentials creates a new Credentials
func NewCredentials() *Credentials {
	s := &Credentials{}
	return s
}

// Credentials entity
type Credentials struct {
	AccessKeyID     string
	SecretAccessKey string
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

// Hydrate implements struct hydrate
func (e *Credentials) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "accessKeyID", &e.AccessKeyID); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "secretAccessKey", &e.SecretAccessKey); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *Credentials) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("accessKeyID", e.AccessKeyID)
	ctx.Content().SetProperty("secretAccessKey", e.SecretAccessKey)
	return nil
}

// StructPath returns StructPath
func (e *Credentials) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathCredentials.Path()
}

// NewSSOProviderUnavailableProblem creates a new SSOProviderUnavailableProblem
func NewSSOProviderUnavailableProblem() *SSOProviderUnavailableProblem {
	s := &SSOProviderUnavailableProblem{}
	return s
}

// SSOProviderUnavailableProblem entity
type SSOProviderUnavailableProblem struct {
	Message string
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

// Hydrate implements struct hydrate
func (e *SSOProviderUnavailableProblem) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "message", &e.Message); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *SSOProviderUnavailableProblem) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("message", e.Message)
	return nil
}

// StructPath returns StructPath
func (e *SSOProviderUnavailableProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathSSOProviderUnavailableProblem.Path()
}

// NewSSOFlow creates a new SSOFlow
func NewSSOFlow() *SSOFlow {
	s := &SSOFlow{}
	return s
}

// SSOFlow entity
type SSOFlow struct {
	BrowseURL                 string
	CompletionIntervalSeconds int32
	ExpiresAt                 time.Time
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

// Hydrate implements struct hydrate
func (e *SSOFlow) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "browseURL", &e.BrowseURL); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireNumericProperty(ctx.Content(), "completionIntervalSeconds", &e.CompletionIntervalSeconds); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireTimeProperty(ctx.Content(), "expiresAt", &e.ExpiresAt); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *SSOFlow) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("browseURL", e.BrowseURL)
	ctx.Content().SetProperty("completionIntervalSeconds", e.CompletionIntervalSeconds)
	ctx.Content().SetProperty("expiresAt", e.ExpiresAt)
	return nil
}

// StructPath returns StructPath
func (e *SSOFlow) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathSSOFlow.Path()
}

// NewAccount creates a new Account
func NewAccount() *Account {
	s := &Account{}
	return s
}

// Account entity
type Account struct {
	Name string
}

// GetName returns the value for the field name
func (e *Account) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *Account) SetName(name string) {
	e.Name = name
}

// Hydrate implements struct hydrate
func (e *Account) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "name", &e.Name); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *Account) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("name", e.Name)
	return nil
}

// StructPath returns StructPath
func (e *Account) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccount.Path()
}

// NewAccountSSOProvider creates a new AccountSSOProvider
func NewAccountSSOProvider() *AccountSSOProvider {
	s := &AccountSSOProvider{}
	return s
}

// AccountSSOProvider entity
type AccountSSOProvider struct {
	DisplayName string
	Name        string
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

// Hydrate implements struct hydrate
func (e *AccountSSOProvider) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "displayName", &e.DisplayName); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "name", &e.Name); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *AccountSSOProvider) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("displayName", e.DisplayName)
	ctx.Content().SetProperty("name", e.Name)
	return nil
}

// StructPath returns StructPath
func (e *AccountSSOProvider) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOProvider.Path()
}

// NewPolicyNotFoundProblem creates a new PolicyNotFoundProblem
func NewPolicyNotFoundProblem() *PolicyNotFoundProblem {
	s := &PolicyNotFoundProblem{}
	return s
}

// PolicyNotFoundProblem entity
type PolicyNotFoundProblem struct {
	Message string
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

// Hydrate implements struct hydrate
func (e *PolicyNotFoundProblem) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "message", &e.Message); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *PolicyNotFoundProblem) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("message", e.Message)
	return nil
}

// StructPath returns StructPath
func (e *PolicyNotFoundProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathPolicyNotFoundProblem.Path()
}

// NewAccountSSOAutoJoinPolicy creates a new AccountSSOAutoJoinPolicy
func NewAccountSSOAutoJoinPolicy() *AccountSSOAutoJoinPolicy {
	s := &AccountSSOAutoJoinPolicy{}
	return s
}

// AccountSSOAutoJoinPolicy entity
type AccountSSOAutoJoinPolicy struct {
	// email-domain users come from, example "deployport.com"
	Domain  string
	Enabled bool
	// unique identifer of the auto-join policy
	Name string
	// where the users are coming from
	OriginAccountName string
	// SSO provider in the source account
	OriginProviderName string
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

// Hydrate implements struct hydrate
func (e *AccountSSOAutoJoinPolicy) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "domain", &e.Domain); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireBoolProperty(ctx.Content(), "enabled", &e.Enabled); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "name", &e.Name); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "originAccountName", &e.OriginAccountName); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "originProviderName", &e.OriginProviderName); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *AccountSSOAutoJoinPolicy) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("domain", e.Domain)
	ctx.Content().SetProperty("enabled", e.Enabled)
	ctx.Content().SetProperty("name", e.Name)
	ctx.Content().SetProperty("originAccountName", e.OriginAccountName)
	ctx.Content().SetProperty("originProviderName", e.OriginProviderName)
	return nil
}

// StructPath returns StructPath
func (e *AccountSSOAutoJoinPolicy) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOAutoJoinPolicy.Path()
}

// NewAccountCreateInput creates a new AccountCreateInput
func NewAccountCreateInput() *AccountCreateInput {
	s := &AccountCreateInput{}
	return s
}

// AccountCreateInput entity
type AccountCreateInput struct {
	Name string
}

// GetName returns the value for the field name
func (e *AccountCreateInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *AccountCreateInput) SetName(name string) {
	e.Name = name
}

// Hydrate implements struct hydrate
func (e *AccountCreateInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "name", &e.Name); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *AccountCreateInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("name", e.Name)
	return nil
}

// StructPath returns StructPath
func (e *AccountCreateInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountCreateInput.Path()
}

// NewAccountCreateOutput creates a new AccountCreateOutput
func NewAccountCreateOutput() *AccountCreateOutput {
	s := &AccountCreateOutput{}
	return s
}

// AccountCreateOutput entity
type AccountCreateOutput struct {
	Account *Account
}

// GetAccount returns the value for the field account
func (e *AccountCreateOutput) GetAccount() *Account {
	return e.Account
}

// SetAccount sets the value for the field account
func (e *AccountCreateOutput) SetAccount(account *Account) {
	e.Account = account
}

// Hydrate implements struct hydrate
func (e *AccountCreateOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentObjectProperty(ctx.Content(), SpecularMeta().Account().TypeBuilder(), ctx.Package(), "account", &e.Account); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *AccountCreateOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {

	var fieldContentAccount clientruntime.Content
	if v := e.Account; v != nil {
		fieldContentAccount = clientruntime.NewContent()
		if err := v.Dehydrate(ctx.CopyWithContent(fieldContentAccount)); err != nil {
			return clientruntime.NewPropertyDehydrationError("account", err)
		}
	}
	if fieldContentAccount == nil {
		return clientruntime.NewPropertyRequiredError("account")
	}
	ctx.Content().SetProperty("account", fieldContentAccount.Map())
	return nil
}

// StructPath returns StructPath
func (e *AccountCreateOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountCreateOutput.Path()
}

// NewAccountCreateInvalidNameProblem creates a new AccountCreateInvalidNameProblem
func NewAccountCreateInvalidNameProblem() *AccountCreateInvalidNameProblem {
	s := &AccountCreateInvalidNameProblem{}
	return s
}

// AccountCreateInvalidNameProblem entity
type AccountCreateInvalidNameProblem struct {
	Message string
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

// Hydrate implements struct hydrate
func (e *AccountCreateInvalidNameProblem) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "message", &e.Message); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *AccountCreateInvalidNameProblem) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("message", e.Message)
	return nil
}

// StructPath returns StructPath
func (e *AccountCreateInvalidNameProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountCreateInvalidNameProblem.Path()
}

// NewAccountAssumeIdentityInput creates a new AccountAssumeIdentityInput
func NewAccountAssumeIdentityInput() *AccountAssumeIdentityInput {
	s := &AccountAssumeIdentityInput{}
	return s
}

// AccountAssumeIdentityInput entity
type AccountAssumeIdentityInput struct {
	AccountName string
}

// GetAccountName returns the value for the field accountName
func (e *AccountAssumeIdentityInput) GetAccountName() string {
	return e.AccountName
}

// SetAccountName sets the value for the field accountName
func (e *AccountAssumeIdentityInput) SetAccountName(accountName string) {
	e.AccountName = accountName
}

// Hydrate implements struct hydrate
func (e *AccountAssumeIdentityInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "accountName", &e.AccountName); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *AccountAssumeIdentityInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("accountName", e.AccountName)
	return nil
}

// StructPath returns StructPath
func (e *AccountAssumeIdentityInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountAssumeIdentityInput.Path()
}

// NewAccountAssumeIdentityOutput creates a new AccountAssumeIdentityOutput
func NewAccountAssumeIdentityOutput() *AccountAssumeIdentityOutput {
	s := &AccountAssumeIdentityOutput{}
	return s
}

// AccountAssumeIdentityOutput entity
type AccountAssumeIdentityOutput struct {
	Credentials *Credentials
}

// GetCredentials returns the value for the field credentials
func (e *AccountAssumeIdentityOutput) GetCredentials() *Credentials {
	return e.Credentials
}

// SetCredentials sets the value for the field credentials
func (e *AccountAssumeIdentityOutput) SetCredentials(credentials *Credentials) {
	e.Credentials = credentials
}

// Hydrate implements struct hydrate
func (e *AccountAssumeIdentityOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentObjectProperty(ctx.Content(), SpecularMeta().Credentials().TypeBuilder(), ctx.Package(), "credentials", &e.Credentials); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *AccountAssumeIdentityOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {

	var fieldContentCredentials clientruntime.Content
	if v := e.Credentials; v != nil {
		fieldContentCredentials = clientruntime.NewContent()
		if err := v.Dehydrate(ctx.CopyWithContent(fieldContentCredentials)); err != nil {
			return clientruntime.NewPropertyDehydrationError("credentials", err)
		}
	}
	if fieldContentCredentials == nil {
		return clientruntime.NewPropertyRequiredError("credentials")
	}
	ctx.Content().SetProperty("credentials", fieldContentCredentials.Map())
	return nil
}

// StructPath returns StructPath
func (e *AccountAssumeIdentityOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountAssumeIdentityOutput.Path()
}

// NewAccountSSOBeginAuthenticationInput creates a new AccountSSOBeginAuthenticationInput
func NewAccountSSOBeginAuthenticationInput() *AccountSSOBeginAuthenticationInput {
	s := &AccountSSOBeginAuthenticationInput{}
	return s
}

// AccountSSOBeginAuthenticationInput entity
type AccountSSOBeginAuthenticationInput struct {
	AccountName string
	// client generated code challenge that will be used to verify the completion code, SHA-256 of the codeVerifier
	CodeChallenge string
	ProviderName  string
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

// Hydrate implements struct hydrate
func (e *AccountSSOBeginAuthenticationInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "accountName", &e.AccountName); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "codeChallenge", &e.CodeChallenge); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "providerName", &e.ProviderName); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *AccountSSOBeginAuthenticationInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("accountName", e.AccountName)
	ctx.Content().SetProperty("codeChallenge", e.CodeChallenge)
	ctx.Content().SetProperty("providerName", e.ProviderName)
	return nil
}

// StructPath returns StructPath
func (e *AccountSSOBeginAuthenticationInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOBeginAuthenticationInput.Path()
}

// NewAccountSSOBeginAuthenticationOutput creates a new AccountSSOBeginAuthenticationOutput
func NewAccountSSOBeginAuthenticationOutput() *AccountSSOBeginAuthenticationOutput {
	s := &AccountSSOBeginAuthenticationOutput{}
	return s
}

// AccountSSOBeginAuthenticationOutput entity
type AccountSSOBeginAuthenticationOutput struct {
	Flow *SSOFlow
}

// GetFlow returns the value for the field flow
func (e *AccountSSOBeginAuthenticationOutput) GetFlow() *SSOFlow {
	return e.Flow
}

// SetFlow sets the value for the field flow
func (e *AccountSSOBeginAuthenticationOutput) SetFlow(flow *SSOFlow) {
	e.Flow = flow
}

// Hydrate implements struct hydrate
func (e *AccountSSOBeginAuthenticationOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentObjectProperty(ctx.Content(), SpecularMeta().SSOFlow().TypeBuilder(), ctx.Package(), "flow", &e.Flow); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *AccountSSOBeginAuthenticationOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {

	var fieldContentFlow clientruntime.Content
	if v := e.Flow; v != nil {
		fieldContentFlow = clientruntime.NewContent()
		if err := v.Dehydrate(ctx.CopyWithContent(fieldContentFlow)); err != nil {
			return clientruntime.NewPropertyDehydrationError("flow", err)
		}
	}
	if fieldContentFlow == nil {
		return clientruntime.NewPropertyRequiredError("flow")
	}
	ctx.Content().SetProperty("flow", fieldContentFlow.Map())
	return nil
}

// StructPath returns StructPath
func (e *AccountSSOBeginAuthenticationOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOBeginAuthenticationOutput.Path()
}

// NewAccountSSOBeginAuthenticationParameterProblem creates a new AccountSSOBeginAuthenticationParameterProblem
func NewAccountSSOBeginAuthenticationParameterProblem() *AccountSSOBeginAuthenticationParameterProblem {
	s := &AccountSSOBeginAuthenticationParameterProblem{}
	return s
}

// AccountSSOBeginAuthenticationParameterProblem entity
type AccountSSOBeginAuthenticationParameterProblem struct {
	Message string
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

// Hydrate implements struct hydrate
func (e *AccountSSOBeginAuthenticationParameterProblem) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "message", &e.Message); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *AccountSSOBeginAuthenticationParameterProblem) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("message", e.Message)
	return nil
}

// StructPath returns StructPath
func (e *AccountSSOBeginAuthenticationParameterProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOBeginAuthenticationParameterProblem.Path()
}

// NewAccountSSOCompleteAuthenticationInput creates a new AccountSSOCompleteAuthenticationInput
func NewAccountSSOCompleteAuthenticationInput() *AccountSSOCompleteAuthenticationInput {
	s := &AccountSSOCompleteAuthenticationInput{}
	return s
}

// AccountSSOCompleteAuthenticationInput entity
type AccountSSOCompleteAuthenticationInput struct {
	AccountName string
	// base64-url-encoded client generated code challenge that will be used to verify the completion code
	CodeVerifierB64 string
	ProviderName    string
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

// Hydrate implements struct hydrate
func (e *AccountSSOCompleteAuthenticationInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "accountName", &e.AccountName); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "codeVerifierB64", &e.CodeVerifierB64); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "providerName", &e.ProviderName); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *AccountSSOCompleteAuthenticationInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("accountName", e.AccountName)
	ctx.Content().SetProperty("codeVerifierB64", e.CodeVerifierB64)
	ctx.Content().SetProperty("providerName", e.ProviderName)
	return nil
}

// StructPath returns StructPath
func (e *AccountSSOCompleteAuthenticationInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOCompleteAuthenticationInput.Path()
}

// NewAccountSSOCompleteAuthenticationOutput creates a new AccountSSOCompleteAuthenticationOutput
func NewAccountSSOCompleteAuthenticationOutput() *AccountSSOCompleteAuthenticationOutput {
	s := &AccountSSOCompleteAuthenticationOutput{}
	return s
}

// AccountSSOCompleteAuthenticationOutput entity
type AccountSSOCompleteAuthenticationOutput struct {
	Credentials *Credentials
}

// GetCredentials returns the value for the field credentials
func (e *AccountSSOCompleteAuthenticationOutput) GetCredentials() *Credentials {
	return e.Credentials
}

// SetCredentials sets the value for the field credentials
func (e *AccountSSOCompleteAuthenticationOutput) SetCredentials(credentials *Credentials) {
	e.Credentials = credentials
}

// Hydrate implements struct hydrate
func (e *AccountSSOCompleteAuthenticationOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentObjectProperty(ctx.Content(), SpecularMeta().Credentials().TypeBuilder(), ctx.Package(), "credentials", &e.Credentials); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *AccountSSOCompleteAuthenticationOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {

	var fieldContentCredentials clientruntime.Content
	if v := e.Credentials; v != nil {
		fieldContentCredentials = clientruntime.NewContent()
		if err := v.Dehydrate(ctx.CopyWithContent(fieldContentCredentials)); err != nil {
			return clientruntime.NewPropertyDehydrationError("credentials", err)
		}
	}
	if fieldContentCredentials == nil {
		return clientruntime.NewPropertyRequiredError("credentials")
	}
	ctx.Content().SetProperty("credentials", fieldContentCredentials.Map())
	return nil
}

// StructPath returns StructPath
func (e *AccountSSOCompleteAuthenticationOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOCompleteAuthenticationOutput.Path()
}

// NewAccountSSOCompleteAuthenticationInvalidFlowProblem creates a new AccountSSOCompleteAuthenticationInvalidFlowProblem
func NewAccountSSOCompleteAuthenticationInvalidFlowProblem() *AccountSSOCompleteAuthenticationInvalidFlowProblem {
	s := &AccountSSOCompleteAuthenticationInvalidFlowProblem{}
	return s
}

// AccountSSOCompleteAuthenticationInvalidFlowProblem entity
type AccountSSOCompleteAuthenticationInvalidFlowProblem struct {
	Message string
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

// Hydrate implements struct hydrate
func (e *AccountSSOCompleteAuthenticationInvalidFlowProblem) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "message", &e.Message); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *AccountSSOCompleteAuthenticationInvalidFlowProblem) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("message", e.Message)
	return nil
}

// StructPath returns StructPath
func (e *AccountSSOCompleteAuthenticationInvalidFlowProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOCompleteAuthenticationInvalidFlowProblem.Path()
}

// NewAccountSSOGetProvidersInput creates a new AccountSSOGetProvidersInput
func NewAccountSSOGetProvidersInput() *AccountSSOGetProvidersInput {
	s := &AccountSSOGetProvidersInput{}
	return s
}

// AccountSSOGetProvidersInput entity
type AccountSSOGetProvidersInput struct {
	AccountName string
}

// GetAccountName returns the value for the field accountName
func (e *AccountSSOGetProvidersInput) GetAccountName() string {
	return e.AccountName
}

// SetAccountName sets the value for the field accountName
func (e *AccountSSOGetProvidersInput) SetAccountName(accountName string) {
	e.AccountName = accountName
}

// Hydrate implements struct hydrate
func (e *AccountSSOGetProvidersInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "accountName", &e.AccountName); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *AccountSSOGetProvidersInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("accountName", e.AccountName)
	return nil
}

// StructPath returns StructPath
func (e *AccountSSOGetProvidersInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOGetProvidersInput.Path()
}

// NewAccountSSOGetProvidersOutput creates a new AccountSSOGetProvidersOutput
func NewAccountSSOGetProvidersOutput() *AccountSSOGetProvidersOutput {
	s := &AccountSSOGetProvidersOutput{}
	return s
}

// AccountSSOGetProvidersOutput entity
type AccountSSOGetProvidersOutput struct {
	Providers []*AccountSSOProvider
}

// GetProviders returns the value for the field providers
func (e *AccountSSOGetProvidersOutput) GetProviders() []*AccountSSOProvider {
	return e.Providers
}

// SetProviders sets the value for the field providers
func (e *AccountSSOGetProvidersOutput) SetProviders(providers []*AccountSSOProvider) {
	e.Providers = providers
}

// Hydrate implements struct hydrate
func (e *AccountSSOGetProvidersOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentObjectArrayProperty(ctx.Content(), SpecularMeta().AccountSSOProvider().TypeBuilder(), ctx.Package(), "providers", &e.Providers); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *AccountSSOGetProvidersOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	var fieldValuesProviders []map[string]any

	if fsv := e.Providers; fsv != nil {
		fieldValuesProviders = make([]map[string]any, 0, len(fsv))
		for ix, fsvitem := range fsv {
			var fsvm map[string]any
			if fsvitem != nil {
				ct := clientruntime.NewContent()
				if err := fsvitem.Dehydrate(ctx.CopyWithContent(ct)); err != nil {
					return clientruntime.NewPropertyItemDehydrationError("providers", ix, err)
				}
				fsvm = ct.Map()
			}
			if fsvm == nil {
				return clientruntime.NewArrayItemMissingError("providers", ix)
			}
			fieldValuesProviders = append(fieldValuesProviders, fsvm)
		}
	}

	ctx.Content().SetProperty("providers", fieldValuesProviders)
	return nil
}

// StructPath returns StructPath
func (e *AccountSSOGetProvidersOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOGetProvidersOutput.Path()
}

// NewAccountSSOAutoJoinPolicyCreateInput creates a new AccountSSOAutoJoinPolicyCreateInput
func NewAccountSSOAutoJoinPolicyCreateInput() *AccountSSOAutoJoinPolicyCreateInput {
	s := &AccountSSOAutoJoinPolicyCreateInput{}
	return s
}

// AccountSSOAutoJoinPolicyCreateInput entity
type AccountSSOAutoJoinPolicyCreateInput struct {
	Domain             string
	OriginAccountName  string
	OriginProviderName string
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

// Hydrate implements struct hydrate
func (e *AccountSSOAutoJoinPolicyCreateInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "domain", &e.Domain); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "originAccountName", &e.OriginAccountName); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "originProviderName", &e.OriginProviderName); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *AccountSSOAutoJoinPolicyCreateInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("domain", e.Domain)
	ctx.Content().SetProperty("originAccountName", e.OriginAccountName)
	ctx.Content().SetProperty("originProviderName", e.OriginProviderName)
	return nil
}

// StructPath returns StructPath
func (e *AccountSSOAutoJoinPolicyCreateInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOAutoJoinPolicyCreateInput.Path()
}

// NewAccountSSOAutoJoinPolicyCreateOutput creates a new AccountSSOAutoJoinPolicyCreateOutput
func NewAccountSSOAutoJoinPolicyCreateOutput() *AccountSSOAutoJoinPolicyCreateOutput {
	s := &AccountSSOAutoJoinPolicyCreateOutput{}
	return s
}

// AccountSSOAutoJoinPolicyCreateOutput entity
type AccountSSOAutoJoinPolicyCreateOutput struct {
	Policy *AccountSSOAutoJoinPolicy
}

// GetPolicy returns the value for the field policy
func (e *AccountSSOAutoJoinPolicyCreateOutput) GetPolicy() *AccountSSOAutoJoinPolicy {
	return e.Policy
}

// SetPolicy sets the value for the field policy
func (e *AccountSSOAutoJoinPolicyCreateOutput) SetPolicy(policy *AccountSSOAutoJoinPolicy) {
	e.Policy = policy
}

// Hydrate implements struct hydrate
func (e *AccountSSOAutoJoinPolicyCreateOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentObjectProperty(ctx.Content(), SpecularMeta().AccountSSOAutoJoinPolicy().TypeBuilder(), ctx.Package(), "policy", &e.Policy); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *AccountSSOAutoJoinPolicyCreateOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {

	var fieldContentPolicy clientruntime.Content
	if v := e.Policy; v != nil {
		fieldContentPolicy = clientruntime.NewContent()
		if err := v.Dehydrate(ctx.CopyWithContent(fieldContentPolicy)); err != nil {
			return clientruntime.NewPropertyDehydrationError("policy", err)
		}
	}
	if fieldContentPolicy == nil {
		return clientruntime.NewPropertyRequiredError("policy")
	}
	ctx.Content().SetProperty("policy", fieldContentPolicy.Map())
	return nil
}

// StructPath returns StructPath
func (e *AccountSSOAutoJoinPolicyCreateOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOAutoJoinPolicyCreateOutput.Path()
}

// NewAccountSSOAutoJoinPolicyListInput creates a new AccountSSOAutoJoinPolicyListInput
func NewAccountSSOAutoJoinPolicyListInput() *AccountSSOAutoJoinPolicyListInput {
	s := &AccountSSOAutoJoinPolicyListInput{}
	return s
}

// AccountSSOAutoJoinPolicyListInput entity
type AccountSSOAutoJoinPolicyListInput struct {
}

// Hydrate implements struct hydrate
func (e *AccountSSOAutoJoinPolicyListInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	return nil
}

// Dehydrate implements struct dehydrate
func (e *AccountSSOAutoJoinPolicyListInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	return nil
}

// StructPath returns StructPath
func (e *AccountSSOAutoJoinPolicyListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOAutoJoinPolicyListInput.Path()
}

// NewAccountSSOAutoJoinPolicyListOutput creates a new AccountSSOAutoJoinPolicyListOutput
func NewAccountSSOAutoJoinPolicyListOutput() *AccountSSOAutoJoinPolicyListOutput {
	s := &AccountSSOAutoJoinPolicyListOutput{}
	return s
}

// AccountSSOAutoJoinPolicyListOutput entity
type AccountSSOAutoJoinPolicyListOutput struct {
	Policies []*AccountSSOAutoJoinPolicy
}

// GetPolicies returns the value for the field policies
func (e *AccountSSOAutoJoinPolicyListOutput) GetPolicies() []*AccountSSOAutoJoinPolicy {
	return e.Policies
}

// SetPolicies sets the value for the field policies
func (e *AccountSSOAutoJoinPolicyListOutput) SetPolicies(policies []*AccountSSOAutoJoinPolicy) {
	e.Policies = policies
}

// Hydrate implements struct hydrate
func (e *AccountSSOAutoJoinPolicyListOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentObjectArrayProperty(ctx.Content(), SpecularMeta().AccountSSOAutoJoinPolicy().TypeBuilder(), ctx.Package(), "policies", &e.Policies); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *AccountSSOAutoJoinPolicyListOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	var fieldValuesPolicies []map[string]any

	if fsv := e.Policies; fsv != nil {
		fieldValuesPolicies = make([]map[string]any, 0, len(fsv))
		for ix, fsvitem := range fsv {
			var fsvm map[string]any
			if fsvitem != nil {
				ct := clientruntime.NewContent()
				if err := fsvitem.Dehydrate(ctx.CopyWithContent(ct)); err != nil {
					return clientruntime.NewPropertyItemDehydrationError("policies", ix, err)
				}
				fsvm = ct.Map()
			}
			if fsvm == nil {
				return clientruntime.NewArrayItemMissingError("policies", ix)
			}
			fieldValuesPolicies = append(fieldValuesPolicies, fsvm)
		}
	}

	ctx.Content().SetProperty("policies", fieldValuesPolicies)
	return nil
}

// StructPath returns StructPath
func (e *AccountSSOAutoJoinPolicyListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOAutoJoinPolicyListOutput.Path()
}

// NewAccountSSOAutoJoinPolicyEnableInput creates a new AccountSSOAutoJoinPolicyEnableInput
func NewAccountSSOAutoJoinPolicyEnableInput() *AccountSSOAutoJoinPolicyEnableInput {
	s := &AccountSSOAutoJoinPolicyEnableInput{}
	return s
}

// AccountSSOAutoJoinPolicyEnableInput entity
type AccountSSOAutoJoinPolicyEnableInput struct {
	Enabled    bool
	PolicyName string
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

// Hydrate implements struct hydrate
func (e *AccountSSOAutoJoinPolicyEnableInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireBoolProperty(ctx.Content(), "enabled", &e.Enabled); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "policyName", &e.PolicyName); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *AccountSSOAutoJoinPolicyEnableInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("enabled", e.Enabled)
	ctx.Content().SetProperty("policyName", e.PolicyName)
	return nil
}

// StructPath returns StructPath
func (e *AccountSSOAutoJoinPolicyEnableInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOAutoJoinPolicyEnableInput.Path()
}

// NewAccountSSOAutoJoinPolicyEnableOutput creates a new AccountSSOAutoJoinPolicyEnableOutput
func NewAccountSSOAutoJoinPolicyEnableOutput() *AccountSSOAutoJoinPolicyEnableOutput {
	s := &AccountSSOAutoJoinPolicyEnableOutput{}
	return s
}

// AccountSSOAutoJoinPolicyEnableOutput entity
type AccountSSOAutoJoinPolicyEnableOutput struct {
}

// Hydrate implements struct hydrate
func (e *AccountSSOAutoJoinPolicyEnableOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	return nil
}

// Dehydrate implements struct dehydrate
func (e *AccountSSOAutoJoinPolicyEnableOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	return nil
}

// StructPath returns StructPath
func (e *AccountSSOAutoJoinPolicyEnableOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOAutoJoinPolicyEnableOutput.Path()
}

// NewInvalidUsernameProblem creates a new InvalidUsernameProblem
func NewInvalidUsernameProblem() *InvalidUsernameProblem {
	s := &InvalidUsernameProblem{}
	return s
}

// InvalidUsernameProblem entity
type InvalidUsernameProblem struct {
	Message string
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

// Hydrate implements struct hydrate
func (e *InvalidUsernameProblem) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "message", &e.Message); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *InvalidUsernameProblem) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("message", e.Message)
	return nil
}

// StructPath returns StructPath
func (e *InvalidUsernameProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvalidUsernameProblem.Path()
}

// NewMemberAccount creates a new MemberAccount
func NewMemberAccount() *MemberAccount {
	s := &MemberAccount{}
	return s
}

// MemberAccount entity
type MemberAccount struct {
	AccountName string
}

// GetAccountName returns the value for the field accountName
func (e *MemberAccount) GetAccountName() string {
	return e.AccountName
}

// SetAccountName sets the value for the field accountName
func (e *MemberAccount) SetAccountName(accountName string) {
	e.AccountName = accountName
}

// Hydrate implements struct hydrate
func (e *MemberAccount) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "accountName", &e.AccountName); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *MemberAccount) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("accountName", e.AccountName)
	return nil
}

// StructPath returns StructPath
func (e *MemberAccount) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathMemberAccount.Path()
}

// NewCredentialInfo creates a new CredentialInfo
func NewCredentialInfo() *CredentialInfo {
	s := &CredentialInfo{}
	return s
}

// CredentialInfo entity
type CredentialInfo struct {
	AccessKeyID string
}

// GetAccessKeyID returns the value for the field accessKeyID
func (e *CredentialInfo) GetAccessKeyID() string {
	return e.AccessKeyID
}

// SetAccessKeyID sets the value for the field accessKeyID
func (e *CredentialInfo) SetAccessKeyID(accessKeyID string) {
	e.AccessKeyID = accessKeyID
}

// Hydrate implements struct hydrate
func (e *CredentialInfo) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "accessKeyID", &e.AccessKeyID); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *CredentialInfo) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("accessKeyID", e.AccessKeyID)
	return nil
}

// StructPath returns StructPath
func (e *CredentialInfo) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathCredentialInfo.Path()
}

// NewIdentityPolicyStatement creates a new IdentityPolicyStatement
func NewIdentityPolicyStatement() *IdentityPolicyStatement {
	s := &IdentityPolicyStatement{}
	return s
}

// IdentityPolicyStatement entity
type IdentityPolicyStatement struct {
	Actions   []string
	Resources []string
}

// GetActions returns the value for the field actions
func (e *IdentityPolicyStatement) GetActions() []string {
	return e.Actions
}

// SetActions sets the value for the field actions
func (e *IdentityPolicyStatement) SetActions(actions []string) {
	e.Actions = actions
}

// GetResources returns the value for the field resources
func (e *IdentityPolicyStatement) GetResources() []string {
	return e.Resources
}

// SetResources sets the value for the field resources
func (e *IdentityPolicyStatement) SetResources(resources []string) {
	e.Resources = resources
}

// Hydrate implements struct hydrate
func (e *IdentityPolicyStatement) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireBuiltinArrayProperty(ctx.Content(), "actions", &e.Actions); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireBuiltinArrayProperty(ctx.Content(), "resources", &e.Resources); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *IdentityPolicyStatement) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("actions", e.Actions)
	ctx.Content().SetProperty("resources", e.Resources)
	return nil
}

// StructPath returns StructPath
func (e *IdentityPolicyStatement) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyStatement.Path()
}

// NewIdentityPolicy creates a new IdentityPolicy
func NewIdentityPolicy() *IdentityPolicy {
	s := &IdentityPolicy{}
	return s
}

// IdentityPolicy entity
type IdentityPolicy struct {
	Builtin    bool
	Name       string
	Statements []*IdentityPolicyStatement
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

// Hydrate implements struct hydrate
func (e *IdentityPolicy) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireBoolProperty(ctx.Content(), "builtin", &e.Builtin); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "name", &e.Name); err != nil {
		return err
	}
	if err := clientruntime.ContentObjectArrayProperty(ctx.Content(), SpecularMeta().IdentityPolicyStatement().TypeBuilder(), ctx.Package(), "statements", &e.Statements); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *IdentityPolicy) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("builtin", e.Builtin)
	ctx.Content().SetProperty("name", e.Name)
	var fieldValuesStatements []map[string]any

	if fsv := e.Statements; fsv != nil {
		fieldValuesStatements = make([]map[string]any, 0, len(fsv))
		for ix, fsvitem := range fsv {
			var fsvm map[string]any
			if fsvitem != nil {
				ct := clientruntime.NewContent()
				if err := fsvitem.Dehydrate(ctx.CopyWithContent(ct)); err != nil {
					return clientruntime.NewPropertyItemDehydrationError("statements", ix, err)
				}
				fsvm = ct.Map()
			}
			if fsvm == nil {
				return clientruntime.NewArrayItemMissingError("statements", ix)
			}
			fieldValuesStatements = append(fieldValuesStatements, fsvm)
		}
	}

	ctx.Content().SetProperty("statements", fieldValuesStatements)
	return nil
}

// StructPath returns StructPath
func (e *IdentityPolicy) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicy.Path()
}

// NewIdentityPolicyAttachment creates a new IdentityPolicyAttachment
func NewIdentityPolicyAttachment() *IdentityPolicyAttachment {
	s := &IdentityPolicyAttachment{}
	return s
}

// IdentityPolicyAttachment entity
type IdentityPolicyAttachment struct {
	Policy *IdentityPolicy
}

// GetPolicy returns the value for the field policy
func (e *IdentityPolicyAttachment) GetPolicy() *IdentityPolicy {
	return e.Policy
}

// SetPolicy sets the value for the field policy
func (e *IdentityPolicyAttachment) SetPolicy(policy *IdentityPolicy) {
	e.Policy = policy
}

// Hydrate implements struct hydrate
func (e *IdentityPolicyAttachment) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentObjectProperty(ctx.Content(), SpecularMeta().IdentityPolicy().TypeBuilder(), ctx.Package(), "policy", &e.Policy); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *IdentityPolicyAttachment) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {

	var fieldContentPolicy clientruntime.Content
	if v := e.Policy; v != nil {
		fieldContentPolicy = clientruntime.NewContent()
		if err := v.Dehydrate(ctx.CopyWithContent(fieldContentPolicy)); err != nil {
			return clientruntime.NewPropertyDehydrationError("policy", err)
		}
	}
	if fieldContentPolicy == nil {
		return clientruntime.NewPropertyRequiredError("policy")
	}
	ctx.Content().SetProperty("policy", fieldContentPolicy.Map())
	return nil
}

// StructPath returns StructPath
func (e *IdentityPolicyAttachment) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyAttachment.Path()
}

// NewUserCreateInput creates a new UserCreateInput
func NewUserCreateInput() *UserCreateInput {
	s := &UserCreateInput{}
	return s
}

// UserCreateInput entity
type UserCreateInput struct {
	Description string
	Username    string
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

// Hydrate implements struct hydrate
func (e *UserCreateInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "description", &e.Description); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "username", &e.Username); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserCreateInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("description", e.Description)
	ctx.Content().SetProperty("username", e.Username)
	return nil
}

// StructPath returns StructPath
func (e *UserCreateInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserCreateInput.Path()
}

// NewUserCreateOutput creates a new UserCreateOutput
func NewUserCreateOutput() *UserCreateOutput {
	s := &UserCreateOutput{}
	return s
}

// UserCreateOutput entity
type UserCreateOutput struct {
	User *UserInformation
}

// GetUser returns the value for the field user
func (e *UserCreateOutput) GetUser() *UserInformation {
	return e.User
}

// SetUser sets the value for the field user
func (e *UserCreateOutput) SetUser(user *UserInformation) {
	e.User = user
}

// Hydrate implements struct hydrate
func (e *UserCreateOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentObjectProperty(ctx.Content(), SpecularMeta().UserInformation().TypeBuilder(), ctx.Package(), "user", &e.User); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserCreateOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {

	var fieldContentUser clientruntime.Content
	if v := e.User; v != nil {
		fieldContentUser = clientruntime.NewContent()
		if err := v.Dehydrate(ctx.CopyWithContent(fieldContentUser)); err != nil {
			return clientruntime.NewPropertyDehydrationError("user", err)
		}
	}
	if fieldContentUser == nil {
		return clientruntime.NewPropertyRequiredError("user")
	}
	ctx.Content().SetProperty("user", fieldContentUser.Map())
	return nil
}

// StructPath returns StructPath
func (e *UserCreateOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserCreateOutput.Path()
}

// NewUserGetInput creates a new UserGetInput
func NewUserGetInput() *UserGetInput {
	s := &UserGetInput{}
	return s
}

// UserGetInput entity
type UserGetInput struct {
	Username *string
}

// GetUsername returns the value for the field username
func (e *UserGetInput) GetUsername() *string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *UserGetInput) SetUsername(username *string) {
	e.Username = username
}

// Hydrate implements struct hydrate
func (e *UserGetInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentOptionalStringProperty(ctx.Content(), "username", &e.Username); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserGetInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("username", e.Username)
	return nil
}

// StructPath returns StructPath
func (e *UserGetInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserGetInput.Path()
}

// NewUserGetOutput creates a new UserGetOutput
func NewUserGetOutput() *UserGetOutput {
	s := &UserGetOutput{}
	return s
}

// UserGetOutput entity
type UserGetOutput struct {
	User *UserInformation
}

// GetUser returns the value for the field user
func (e *UserGetOutput) GetUser() *UserInformation {
	return e.User
}

// SetUser sets the value for the field user
func (e *UserGetOutput) SetUser(user *UserInformation) {
	e.User = user
}

// Hydrate implements struct hydrate
func (e *UserGetOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentObjectProperty(ctx.Content(), SpecularMeta().UserInformation().TypeBuilder(), ctx.Package(), "user", &e.User); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserGetOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {

	var fieldContentUser clientruntime.Content
	if v := e.User; v != nil {
		fieldContentUser = clientruntime.NewContent()
		if err := v.Dehydrate(ctx.CopyWithContent(fieldContentUser)); err != nil {
			return clientruntime.NewPropertyDehydrationError("user", err)
		}
	}
	if fieldContentUser == nil {
		return clientruntime.NewPropertyRequiredError("user")
	}
	ctx.Content().SetProperty("user", fieldContentUser.Map())
	return nil
}

// StructPath returns StructPath
func (e *UserGetOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserGetOutput.Path()
}

// NewUserGetUserNotAvailableProblem creates a new UserGetUserNotAvailableProblem
func NewUserGetUserNotAvailableProblem() *UserGetUserNotAvailableProblem {
	s := &UserGetUserNotAvailableProblem{}
	return s
}

// UserGetUserNotAvailableProblem entity
type UserGetUserNotAvailableProblem struct {
	Message string
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

// Hydrate implements struct hydrate
func (e *UserGetUserNotAvailableProblem) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "message", &e.Message); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserGetUserNotAvailableProblem) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("message", e.Message)
	return nil
}

// StructPath returns StructPath
func (e *UserGetUserNotAvailableProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserGetUserNotAvailableProblem.Path()
}

// NewUserDestroyInput creates a new UserDestroyInput
func NewUserDestroyInput() *UserDestroyInput {
	s := &UserDestroyInput{}
	return s
}

// UserDestroyInput entity
type UserDestroyInput struct {
	Username *string
}

// GetUsername returns the value for the field username
func (e *UserDestroyInput) GetUsername() *string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *UserDestroyInput) SetUsername(username *string) {
	e.Username = username
}

// Hydrate implements struct hydrate
func (e *UserDestroyInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentOptionalStringProperty(ctx.Content(), "username", &e.Username); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserDestroyInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("username", e.Username)
	return nil
}

// StructPath returns StructPath
func (e *UserDestroyInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserDestroyInput.Path()
}

// NewUserDestroyOutput creates a new UserDestroyOutput
func NewUserDestroyOutput() *UserDestroyOutput {
	s := &UserDestroyOutput{}
	return s
}

// UserDestroyOutput entity
type UserDestroyOutput struct {
}

// Hydrate implements struct hydrate
func (e *UserDestroyOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserDestroyOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	return nil
}

// StructPath returns StructPath
func (e *UserDestroyOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserDestroyOutput.Path()
}

// NewUserListInput creates a new UserListInput
func NewUserListInput() *UserListInput {
	s := &UserListInput{}
	return s
}

// UserListInput entity
type UserListInput struct {
}

// Hydrate implements struct hydrate
func (e *UserListInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserListInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	return nil
}

// StructPath returns StructPath
func (e *UserListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserListInput.Path()
}

// NewUserListOutput creates a new UserListOutput
func NewUserListOutput() *UserListOutput {
	s := &UserListOutput{}
	return s
}

// UserListOutput entity
type UserListOutput struct {
	Users []*UserInformation
}

// GetUsers returns the value for the field users
func (e *UserListOutput) GetUsers() []*UserInformation {
	return e.Users
}

// SetUsers sets the value for the field users
func (e *UserListOutput) SetUsers(users []*UserInformation) {
	e.Users = users
}

// Hydrate implements struct hydrate
func (e *UserListOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentObjectArrayProperty(ctx.Content(), SpecularMeta().UserInformation().TypeBuilder(), ctx.Package(), "users", &e.Users); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserListOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	var fieldValuesUsers []map[string]any

	if fsv := e.Users; fsv != nil {
		fieldValuesUsers = make([]map[string]any, 0, len(fsv))
		for ix, fsvitem := range fsv {
			var fsvm map[string]any
			if fsvitem != nil {
				ct := clientruntime.NewContent()
				if err := fsvitem.Dehydrate(ctx.CopyWithContent(ct)); err != nil {
					return clientruntime.NewPropertyItemDehydrationError("users", ix, err)
				}
				fsvm = ct.Map()
			}
			if fsvm == nil {
				return clientruntime.NewArrayItemMissingError("users", ix)
			}
			fieldValuesUsers = append(fieldValuesUsers, fsvm)
		}
	}

	ctx.Content().SetProperty("users", fieldValuesUsers)
	return nil
}

// StructPath returns StructPath
func (e *UserListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserListOutput.Path()
}

// NewUserMemberAccountsInput creates a new UserMemberAccountsInput
func NewUserMemberAccountsInput() *UserMemberAccountsInput {
	s := &UserMemberAccountsInput{}
	return s
}

// UserMemberAccountsInput entity
type UserMemberAccountsInput struct {
}

// Hydrate implements struct hydrate
func (e *UserMemberAccountsInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserMemberAccountsInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	return nil
}

// StructPath returns StructPath
func (e *UserMemberAccountsInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserMemberAccountsInput.Path()
}

// NewUserMemberAccountsOutput creates a new UserMemberAccountsOutput
func NewUserMemberAccountsOutput() *UserMemberAccountsOutput {
	s := &UserMemberAccountsOutput{}
	return s
}

// UserMemberAccountsOutput entity
type UserMemberAccountsOutput struct {
	Accounts []*MemberAccount
}

// GetAccounts returns the value for the field accounts
func (e *UserMemberAccountsOutput) GetAccounts() []*MemberAccount {
	return e.Accounts
}

// SetAccounts sets the value for the field accounts
func (e *UserMemberAccountsOutput) SetAccounts(accounts []*MemberAccount) {
	e.Accounts = accounts
}

// Hydrate implements struct hydrate
func (e *UserMemberAccountsOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentObjectArrayProperty(ctx.Content(), SpecularMeta().MemberAccount().TypeBuilder(), ctx.Package(), "accounts", &e.Accounts); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserMemberAccountsOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	var fieldValuesAccounts []map[string]any

	if fsv := e.Accounts; fsv != nil {
		fieldValuesAccounts = make([]map[string]any, 0, len(fsv))
		for ix, fsvitem := range fsv {
			var fsvm map[string]any
			if fsvitem != nil {
				ct := clientruntime.NewContent()
				if err := fsvitem.Dehydrate(ctx.CopyWithContent(ct)); err != nil {
					return clientruntime.NewPropertyItemDehydrationError("accounts", ix, err)
				}
				fsvm = ct.Map()
			}
			if fsvm == nil {
				return clientruntime.NewArrayItemMissingError("accounts", ix)
			}
			fieldValuesAccounts = append(fieldValuesAccounts, fsvm)
		}
	}

	ctx.Content().SetProperty("accounts", fieldValuesAccounts)
	return nil
}

// StructPath returns StructPath
func (e *UserMemberAccountsOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserMemberAccountsOutput.Path()
}

// NewUserAccessKeyCreateInput creates a new UserAccessKeyCreateInput
func NewUserAccessKeyCreateInput() *UserAccessKeyCreateInput {
	s := &UserAccessKeyCreateInput{}
	return s
}

// UserAccessKeyCreateInput entity
type UserAccessKeyCreateInput struct {
	Username string
}

// GetUsername returns the value for the field username
func (e *UserAccessKeyCreateInput) GetUsername() string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *UserAccessKeyCreateInput) SetUsername(username string) {
	e.Username = username
}

// Hydrate implements struct hydrate
func (e *UserAccessKeyCreateInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "username", &e.Username); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserAccessKeyCreateInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("username", e.Username)
	return nil
}

// StructPath returns StructPath
func (e *UserAccessKeyCreateInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserAccessKeyCreateInput.Path()
}

// NewUserAccessKeyCreateOutput creates a new UserAccessKeyCreateOutput
func NewUserAccessKeyCreateOutput() *UserAccessKeyCreateOutput {
	s := &UserAccessKeyCreateOutput{}
	return s
}

// UserAccessKeyCreateOutput entity
type UserAccessKeyCreateOutput struct {
	Credentials *Credentials
}

// GetCredentials returns the value for the field credentials
func (e *UserAccessKeyCreateOutput) GetCredentials() *Credentials {
	return e.Credentials
}

// SetCredentials sets the value for the field credentials
func (e *UserAccessKeyCreateOutput) SetCredentials(credentials *Credentials) {
	e.Credentials = credentials
}

// Hydrate implements struct hydrate
func (e *UserAccessKeyCreateOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentObjectProperty(ctx.Content(), SpecularMeta().Credentials().TypeBuilder(), ctx.Package(), "credentials", &e.Credentials); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserAccessKeyCreateOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {

	var fieldContentCredentials clientruntime.Content
	if v := e.Credentials; v != nil {
		fieldContentCredentials = clientruntime.NewContent()
		if err := v.Dehydrate(ctx.CopyWithContent(fieldContentCredentials)); err != nil {
			return clientruntime.NewPropertyDehydrationError("credentials", err)
		}
	}
	if fieldContentCredentials == nil {
		return clientruntime.NewPropertyRequiredError("credentials")
	}
	ctx.Content().SetProperty("credentials", fieldContentCredentials.Map())
	return nil
}

// StructPath returns StructPath
func (e *UserAccessKeyCreateOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserAccessKeyCreateOutput.Path()
}

// NewUserAccessKeyListInput creates a new UserAccessKeyListInput
func NewUserAccessKeyListInput() *UserAccessKeyListInput {
	s := &UserAccessKeyListInput{}
	return s
}

// UserAccessKeyListInput entity
type UserAccessKeyListInput struct {
	Username *string
}

// GetUsername returns the value for the field username
func (e *UserAccessKeyListInput) GetUsername() *string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *UserAccessKeyListInput) SetUsername(username *string) {
	e.Username = username
}

// Hydrate implements struct hydrate
func (e *UserAccessKeyListInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentOptionalStringProperty(ctx.Content(), "username", &e.Username); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserAccessKeyListInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("username", e.Username)
	return nil
}

// StructPath returns StructPath
func (e *UserAccessKeyListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserAccessKeyListInput.Path()
}

// NewUserAccessKeyListOutput creates a new UserAccessKeyListOutput
func NewUserAccessKeyListOutput() *UserAccessKeyListOutput {
	s := &UserAccessKeyListOutput{}
	return s
}

// UserAccessKeyListOutput entity
type UserAccessKeyListOutput struct {
	Credentials []*CredentialInfo
}

// GetCredentials returns the value for the field credentials
func (e *UserAccessKeyListOutput) GetCredentials() []*CredentialInfo {
	return e.Credentials
}

// SetCredentials sets the value for the field credentials
func (e *UserAccessKeyListOutput) SetCredentials(credentials []*CredentialInfo) {
	e.Credentials = credentials
}

// Hydrate implements struct hydrate
func (e *UserAccessKeyListOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentObjectArrayProperty(ctx.Content(), SpecularMeta().CredentialInfo().TypeBuilder(), ctx.Package(), "credentials", &e.Credentials); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserAccessKeyListOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	var fieldValuesCredentials []map[string]any

	if fsv := e.Credentials; fsv != nil {
		fieldValuesCredentials = make([]map[string]any, 0, len(fsv))
		for ix, fsvitem := range fsv {
			var fsvm map[string]any
			if fsvitem != nil {
				ct := clientruntime.NewContent()
				if err := fsvitem.Dehydrate(ctx.CopyWithContent(ct)); err != nil {
					return clientruntime.NewPropertyItemDehydrationError("credentials", ix, err)
				}
				fsvm = ct.Map()
			}
			if fsvm == nil {
				return clientruntime.NewArrayItemMissingError("credentials", ix)
			}
			fieldValuesCredentials = append(fieldValuesCredentials, fsvm)
		}
	}

	ctx.Content().SetProperty("credentials", fieldValuesCredentials)
	return nil
}

// StructPath returns StructPath
func (e *UserAccessKeyListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserAccessKeyListOutput.Path()
}

// NewUserAccessKeyDestroyInput creates a new UserAccessKeyDestroyInput
func NewUserAccessKeyDestroyInput() *UserAccessKeyDestroyInput {
	s := &UserAccessKeyDestroyInput{}
	return s
}

// UserAccessKeyDestroyInput entity
type UserAccessKeyDestroyInput struct {
	AccessKeyID string
	Username    string
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

// Hydrate implements struct hydrate
func (e *UserAccessKeyDestroyInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "accessKeyID", &e.AccessKeyID); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "username", &e.Username); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserAccessKeyDestroyInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("accessKeyID", e.AccessKeyID)
	ctx.Content().SetProperty("username", e.Username)
	return nil
}

// StructPath returns StructPath
func (e *UserAccessKeyDestroyInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserAccessKeyDestroyInput.Path()
}

// NewUserAccessKeyDestroyOutput creates a new UserAccessKeyDestroyOutput
func NewUserAccessKeyDestroyOutput() *UserAccessKeyDestroyOutput {
	s := &UserAccessKeyDestroyOutput{}
	return s
}

// UserAccessKeyDestroyOutput entity
type UserAccessKeyDestroyOutput struct {
}

// Hydrate implements struct hydrate
func (e *UserAccessKeyDestroyOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserAccessKeyDestroyOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	return nil
}

// StructPath returns StructPath
func (e *UserAccessKeyDestroyOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserAccessKeyDestroyOutput.Path()
}

// NewUserIdentityPolicyAttachInput creates a new UserIdentityPolicyAttachInput
func NewUserIdentityPolicyAttachInput() *UserIdentityPolicyAttachInput {
	s := &UserIdentityPolicyAttachInput{}
	return s
}

// UserIdentityPolicyAttachInput entity
type UserIdentityPolicyAttachInput struct {
	PolicyName string
	Username   string
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

// Hydrate implements struct hydrate
func (e *UserIdentityPolicyAttachInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "policyName", &e.PolicyName); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "username", &e.Username); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserIdentityPolicyAttachInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("policyName", e.PolicyName)
	ctx.Content().SetProperty("username", e.Username)
	return nil
}

// StructPath returns StructPath
func (e *UserIdentityPolicyAttachInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserIdentityPolicyAttachInput.Path()
}

// NewUserIdentityPolicyAttachOutput creates a new UserIdentityPolicyAttachOutput
func NewUserIdentityPolicyAttachOutput() *UserIdentityPolicyAttachOutput {
	s := &UserIdentityPolicyAttachOutput{}
	return s
}

// UserIdentityPolicyAttachOutput entity
type UserIdentityPolicyAttachOutput struct {
}

// Hydrate implements struct hydrate
func (e *UserIdentityPolicyAttachOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserIdentityPolicyAttachOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	return nil
}

// StructPath returns StructPath
func (e *UserIdentityPolicyAttachOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserIdentityPolicyAttachOutput.Path()
}

// NewUserIdentityPolicyListInput creates a new UserIdentityPolicyListInput
func NewUserIdentityPolicyListInput() *UserIdentityPolicyListInput {
	s := &UserIdentityPolicyListInput{}
	return s
}

// UserIdentityPolicyListInput entity
type UserIdentityPolicyListInput struct {
	Username string
}

// GetUsername returns the value for the field username
func (e *UserIdentityPolicyListInput) GetUsername() string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *UserIdentityPolicyListInput) SetUsername(username string) {
	e.Username = username
}

// Hydrate implements struct hydrate
func (e *UserIdentityPolicyListInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "username", &e.Username); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserIdentityPolicyListInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("username", e.Username)
	return nil
}

// StructPath returns StructPath
func (e *UserIdentityPolicyListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserIdentityPolicyListInput.Path()
}

// NewUserIdentityPolicyListOutput creates a new UserIdentityPolicyListOutput
func NewUserIdentityPolicyListOutput() *UserIdentityPolicyListOutput {
	s := &UserIdentityPolicyListOutput{}
	return s
}

// UserIdentityPolicyListOutput entity
type UserIdentityPolicyListOutput struct {
	Attachments []*IdentityPolicyAttachment
}

// GetAttachments returns the value for the field attachments
func (e *UserIdentityPolicyListOutput) GetAttachments() []*IdentityPolicyAttachment {
	return e.Attachments
}

// SetAttachments sets the value for the field attachments
func (e *UserIdentityPolicyListOutput) SetAttachments(attachments []*IdentityPolicyAttachment) {
	e.Attachments = attachments
}

// Hydrate implements struct hydrate
func (e *UserIdentityPolicyListOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentObjectArrayProperty(ctx.Content(), SpecularMeta().IdentityPolicyAttachment().TypeBuilder(), ctx.Package(), "attachments", &e.Attachments); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserIdentityPolicyListOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	var fieldValuesAttachments []map[string]any

	if fsv := e.Attachments; fsv != nil {
		fieldValuesAttachments = make([]map[string]any, 0, len(fsv))
		for ix, fsvitem := range fsv {
			var fsvm map[string]any
			if fsvitem != nil {
				ct := clientruntime.NewContent()
				if err := fsvitem.Dehydrate(ctx.CopyWithContent(ct)); err != nil {
					return clientruntime.NewPropertyItemDehydrationError("attachments", ix, err)
				}
				fsvm = ct.Map()
			}
			if fsvm == nil {
				return clientruntime.NewArrayItemMissingError("attachments", ix)
			}
			fieldValuesAttachments = append(fieldValuesAttachments, fsvm)
		}
	}

	ctx.Content().SetProperty("attachments", fieldValuesAttachments)
	return nil
}

// StructPath returns StructPath
func (e *UserIdentityPolicyListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserIdentityPolicyListOutput.Path()
}

// NewUserIdentityPolicyDetachInput creates a new UserIdentityPolicyDetachInput
func NewUserIdentityPolicyDetachInput() *UserIdentityPolicyDetachInput {
	s := &UserIdentityPolicyDetachInput{}
	return s
}

// UserIdentityPolicyDetachInput entity
type UserIdentityPolicyDetachInput struct {
	PolicyName string
	Username   string
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

// Hydrate implements struct hydrate
func (e *UserIdentityPolicyDetachInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "policyName", &e.PolicyName); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "username", &e.Username); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserIdentityPolicyDetachInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("policyName", e.PolicyName)
	ctx.Content().SetProperty("username", e.Username)
	return nil
}

// StructPath returns StructPath
func (e *UserIdentityPolicyDetachInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserIdentityPolicyDetachInput.Path()
}

// NewUserIdentityPolicyDetachOutput creates a new UserIdentityPolicyDetachOutput
func NewUserIdentityPolicyDetachOutput() *UserIdentityPolicyDetachOutput {
	s := &UserIdentityPolicyDetachOutput{}
	return s
}

// UserIdentityPolicyDetachOutput entity
type UserIdentityPolicyDetachOutput struct {
}

// Hydrate implements struct hydrate
func (e *UserIdentityPolicyDetachOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	return nil
}

// Dehydrate implements struct dehydrate
func (e *UserIdentityPolicyDetachOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	return nil
}

// StructPath returns StructPath
func (e *UserIdentityPolicyDetachOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserIdentityPolicyDetachOutput.Path()
}

// NewPolicyStructureProblem creates a new PolicyStructureProblem
func NewPolicyStructureProblem() *PolicyStructureProblem {
	s := &PolicyStructureProblem{}
	return s
}

// PolicyStructureProblem entity
type PolicyStructureProblem struct {
	Message string
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

// Hydrate implements struct hydrate
func (e *PolicyStructureProblem) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "message", &e.Message); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *PolicyStructureProblem) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("message", e.Message)
	return nil
}

// StructPath returns StructPath
func (e *PolicyStructureProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathPolicyStructureProblem.Path()
}

// NewIdentityPolicyCreateInput creates a new IdentityPolicyCreateInput
func NewIdentityPolicyCreateInput() *IdentityPolicyCreateInput {
	s := &IdentityPolicyCreateInput{}
	return s
}

// IdentityPolicyCreateInput entity
type IdentityPolicyCreateInput struct {
	Name       string
	Statements []*IdentityPolicyStatement
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

// Hydrate implements struct hydrate
func (e *IdentityPolicyCreateInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "name", &e.Name); err != nil {
		return err
	}
	if err := clientruntime.ContentObjectArrayProperty(ctx.Content(), SpecularMeta().IdentityPolicyStatement().TypeBuilder(), ctx.Package(), "statements", &e.Statements); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *IdentityPolicyCreateInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("name", e.Name)
	var fieldValuesStatements []map[string]any

	if fsv := e.Statements; fsv != nil {
		fieldValuesStatements = make([]map[string]any, 0, len(fsv))
		for ix, fsvitem := range fsv {
			var fsvm map[string]any
			if fsvitem != nil {
				ct := clientruntime.NewContent()
				if err := fsvitem.Dehydrate(ctx.CopyWithContent(ct)); err != nil {
					return clientruntime.NewPropertyItemDehydrationError("statements", ix, err)
				}
				fsvm = ct.Map()
			}
			if fsvm == nil {
				return clientruntime.NewArrayItemMissingError("statements", ix)
			}
			fieldValuesStatements = append(fieldValuesStatements, fsvm)
		}
	}

	ctx.Content().SetProperty("statements", fieldValuesStatements)
	return nil
}

// StructPath returns StructPath
func (e *IdentityPolicyCreateInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyCreateInput.Path()
}

// NewIdentityPolicyCreateOutput creates a new IdentityPolicyCreateOutput
func NewIdentityPolicyCreateOutput() *IdentityPolicyCreateOutput {
	s := &IdentityPolicyCreateOutput{}
	return s
}

// IdentityPolicyCreateOutput entity
type IdentityPolicyCreateOutput struct {
	Policy *IdentityPolicy
}

// GetPolicy returns the value for the field policy
func (e *IdentityPolicyCreateOutput) GetPolicy() *IdentityPolicy {
	return e.Policy
}

// SetPolicy sets the value for the field policy
func (e *IdentityPolicyCreateOutput) SetPolicy(policy *IdentityPolicy) {
	e.Policy = policy
}

// Hydrate implements struct hydrate
func (e *IdentityPolicyCreateOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentObjectProperty(ctx.Content(), SpecularMeta().IdentityPolicy().TypeBuilder(), ctx.Package(), "policy", &e.Policy); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *IdentityPolicyCreateOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {

	var fieldContentPolicy clientruntime.Content
	if v := e.Policy; v != nil {
		fieldContentPolicy = clientruntime.NewContent()
		if err := v.Dehydrate(ctx.CopyWithContent(fieldContentPolicy)); err != nil {
			return clientruntime.NewPropertyDehydrationError("policy", err)
		}
	}
	if fieldContentPolicy == nil {
		return clientruntime.NewPropertyRequiredError("policy")
	}
	ctx.Content().SetProperty("policy", fieldContentPolicy.Map())
	return nil
}

// StructPath returns StructPath
func (e *IdentityPolicyCreateOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyCreateOutput.Path()
}

// NewIdentityPolicyListInput creates a new IdentityPolicyListInput
func NewIdentityPolicyListInput() *IdentityPolicyListInput {
	s := &IdentityPolicyListInput{}
	return s
}

// IdentityPolicyListInput entity
type IdentityPolicyListInput struct {
}

// Hydrate implements struct hydrate
func (e *IdentityPolicyListInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	return nil
}

// Dehydrate implements struct dehydrate
func (e *IdentityPolicyListInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	return nil
}

// StructPath returns StructPath
func (e *IdentityPolicyListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyListInput.Path()
}

// NewIdentityPolicyListOutput creates a new IdentityPolicyListOutput
func NewIdentityPolicyListOutput() *IdentityPolicyListOutput {
	s := &IdentityPolicyListOutput{}
	return s
}

// IdentityPolicyListOutput entity
type IdentityPolicyListOutput struct {
	Policies []*IdentityPolicy
}

// GetPolicies returns the value for the field policies
func (e *IdentityPolicyListOutput) GetPolicies() []*IdentityPolicy {
	return e.Policies
}

// SetPolicies sets the value for the field policies
func (e *IdentityPolicyListOutput) SetPolicies(policies []*IdentityPolicy) {
	e.Policies = policies
}

// Hydrate implements struct hydrate
func (e *IdentityPolicyListOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentObjectArrayProperty(ctx.Content(), SpecularMeta().IdentityPolicy().TypeBuilder(), ctx.Package(), "policies", &e.Policies); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *IdentityPolicyListOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	var fieldValuesPolicies []map[string]any

	if fsv := e.Policies; fsv != nil {
		fieldValuesPolicies = make([]map[string]any, 0, len(fsv))
		for ix, fsvitem := range fsv {
			var fsvm map[string]any
			if fsvitem != nil {
				ct := clientruntime.NewContent()
				if err := fsvitem.Dehydrate(ctx.CopyWithContent(ct)); err != nil {
					return clientruntime.NewPropertyItemDehydrationError("policies", ix, err)
				}
				fsvm = ct.Map()
			}
			if fsvm == nil {
				return clientruntime.NewArrayItemMissingError("policies", ix)
			}
			fieldValuesPolicies = append(fieldValuesPolicies, fsvm)
		}
	}

	ctx.Content().SetProperty("policies", fieldValuesPolicies)
	return nil
}

// StructPath returns StructPath
func (e *IdentityPolicyListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyListOutput.Path()
}

// NewIdentityPolicyRetrieveInput creates a new IdentityPolicyRetrieveInput
func NewIdentityPolicyRetrieveInput() *IdentityPolicyRetrieveInput {
	s := &IdentityPolicyRetrieveInput{}
	return s
}

// IdentityPolicyRetrieveInput entity
type IdentityPolicyRetrieveInput struct {
	Name string
}

// GetName returns the value for the field name
func (e *IdentityPolicyRetrieveInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *IdentityPolicyRetrieveInput) SetName(name string) {
	e.Name = name
}

// Hydrate implements struct hydrate
func (e *IdentityPolicyRetrieveInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "name", &e.Name); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *IdentityPolicyRetrieveInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("name", e.Name)
	return nil
}

// StructPath returns StructPath
func (e *IdentityPolicyRetrieveInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyRetrieveInput.Path()
}

// NewIdentityPolicyRetrieveOutput creates a new IdentityPolicyRetrieveOutput
func NewIdentityPolicyRetrieveOutput() *IdentityPolicyRetrieveOutput {
	s := &IdentityPolicyRetrieveOutput{}
	return s
}

// IdentityPolicyRetrieveOutput entity
type IdentityPolicyRetrieveOutput struct {
	Policy *IdentityPolicy
}

// GetPolicy returns the value for the field policy
func (e *IdentityPolicyRetrieveOutput) GetPolicy() *IdentityPolicy {
	return e.Policy
}

// SetPolicy sets the value for the field policy
func (e *IdentityPolicyRetrieveOutput) SetPolicy(policy *IdentityPolicy) {
	e.Policy = policy
}

// Hydrate implements struct hydrate
func (e *IdentityPolicyRetrieveOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentObjectProperty(ctx.Content(), SpecularMeta().IdentityPolicy().TypeBuilder(), ctx.Package(), "policy", &e.Policy); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *IdentityPolicyRetrieveOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {

	var fieldContentPolicy clientruntime.Content
	if v := e.Policy; v != nil {
		fieldContentPolicy = clientruntime.NewContent()
		if err := v.Dehydrate(ctx.CopyWithContent(fieldContentPolicy)); err != nil {
			return clientruntime.NewPropertyDehydrationError("policy", err)
		}
	}
	if fieldContentPolicy == nil {
		return clientruntime.NewPropertyRequiredError("policy")
	}
	ctx.Content().SetProperty("policy", fieldContentPolicy.Map())
	return nil
}

// StructPath returns StructPath
func (e *IdentityPolicyRetrieveOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyRetrieveOutput.Path()
}

// NewIdentityPolicyDestroyInput creates a new IdentityPolicyDestroyInput
func NewIdentityPolicyDestroyInput() *IdentityPolicyDestroyInput {
	s := &IdentityPolicyDestroyInput{}
	return s
}

// IdentityPolicyDestroyInput entity
type IdentityPolicyDestroyInput struct {
	Name string
}

// GetName returns the value for the field name
func (e *IdentityPolicyDestroyInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *IdentityPolicyDestroyInput) SetName(name string) {
	e.Name = name
}

// Hydrate implements struct hydrate
func (e *IdentityPolicyDestroyInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "name", &e.Name); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *IdentityPolicyDestroyInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("name", e.Name)
	return nil
}

// StructPath returns StructPath
func (e *IdentityPolicyDestroyInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyDestroyInput.Path()
}

// NewIdentityPolicyDestroyOutput creates a new IdentityPolicyDestroyOutput
func NewIdentityPolicyDestroyOutput() *IdentityPolicyDestroyOutput {
	s := &IdentityPolicyDestroyOutput{}
	return s
}

// IdentityPolicyDestroyOutput entity
type IdentityPolicyDestroyOutput struct {
}

// Hydrate implements struct hydrate
func (e *IdentityPolicyDestroyOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	return nil
}

// Dehydrate implements struct dehydrate
func (e *IdentityPolicyDestroyOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	return nil
}

// StructPath returns StructPath
func (e *IdentityPolicyDestroyOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyDestroyOutput.Path()
}

// NewIdentityPolicyUpdateInput creates a new IdentityPolicyUpdateInput
func NewIdentityPolicyUpdateInput() *IdentityPolicyUpdateInput {
	s := &IdentityPolicyUpdateInput{}
	return s
}

// IdentityPolicyUpdateInput entity
type IdentityPolicyUpdateInput struct {
	Name       string
	Statements []*IdentityPolicyStatement
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

// Hydrate implements struct hydrate
func (e *IdentityPolicyUpdateInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "name", &e.Name); err != nil {
		return err
	}
	if err := clientruntime.ContentObjectArrayProperty(ctx.Content(), SpecularMeta().IdentityPolicyStatement().TypeBuilder(), ctx.Package(), "statements", &e.Statements); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *IdentityPolicyUpdateInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("name", e.Name)
	var fieldValuesStatements []map[string]any

	if fsv := e.Statements; fsv != nil {
		fieldValuesStatements = make([]map[string]any, 0, len(fsv))
		for ix, fsvitem := range fsv {
			var fsvm map[string]any
			if fsvitem != nil {
				ct := clientruntime.NewContent()
				if err := fsvitem.Dehydrate(ctx.CopyWithContent(ct)); err != nil {
					return clientruntime.NewPropertyItemDehydrationError("statements", ix, err)
				}
				fsvm = ct.Map()
			}
			if fsvm == nil {
				return clientruntime.NewArrayItemMissingError("statements", ix)
			}
			fieldValuesStatements = append(fieldValuesStatements, fsvm)
		}
	}

	ctx.Content().SetProperty("statements", fieldValuesStatements)
	return nil
}

// StructPath returns StructPath
func (e *IdentityPolicyUpdateInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyUpdateInput.Path()
}

// NewIdentityPolicyUpdateOutput creates a new IdentityPolicyUpdateOutput
func NewIdentityPolicyUpdateOutput() *IdentityPolicyUpdateOutput {
	s := &IdentityPolicyUpdateOutput{}
	return s
}

// IdentityPolicyUpdateOutput entity
type IdentityPolicyUpdateOutput struct {
	Policy *IdentityPolicy
}

// GetPolicy returns the value for the field policy
func (e *IdentityPolicyUpdateOutput) GetPolicy() *IdentityPolicy {
	return e.Policy
}

// SetPolicy sets the value for the field policy
func (e *IdentityPolicyUpdateOutput) SetPolicy(policy *IdentityPolicy) {
	e.Policy = policy
}

// Hydrate implements struct hydrate
func (e *IdentityPolicyUpdateOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentObjectProperty(ctx.Content(), SpecularMeta().IdentityPolicy().TypeBuilder(), ctx.Package(), "policy", &e.Policy); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *IdentityPolicyUpdateOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {

	var fieldContentPolicy clientruntime.Content
	if v := e.Policy; v != nil {
		fieldContentPolicy = clientruntime.NewContent()
		if err := v.Dehydrate(ctx.CopyWithContent(fieldContentPolicy)); err != nil {
			return clientruntime.NewPropertyDehydrationError("policy", err)
		}
	}
	if fieldContentPolicy == nil {
		return clientruntime.NewPropertyRequiredError("policy")
	}
	ctx.Content().SetProperty("policy", fieldContentPolicy.Map())
	return nil
}

// StructPath returns StructPath
func (e *IdentityPolicyUpdateOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyUpdateOutput.Path()
}

// NewInternalAccessKeyUser creates a new InternalAccessKeyUser
func NewInternalAccessKeyUser() *InternalAccessKeyUser {
	s := &InternalAccessKeyUser{}
	return s
}

// InternalAccessKeyUser entity
type InternalAccessKeyUser struct {
	Username string
}

// GetUsername returns the value for the field username
func (e *InternalAccessKeyUser) GetUsername() string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *InternalAccessKeyUser) SetUsername(username string) {
	e.Username = username
}

// Hydrate implements struct hydrate
func (e *InternalAccessKeyUser) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "username", &e.Username); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *InternalAccessKeyUser) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("username", e.Username)
	return nil
}

// StructPath returns StructPath
func (e *InternalAccessKeyUser) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalAccessKeyUser.Path()
}

// NewInternalAccessKeyAccount creates a new InternalAccessKeyAccount
func NewInternalAccessKeyAccount() *InternalAccessKeyAccount {
	s := &InternalAccessKeyAccount{}
	return s
}

// InternalAccessKeyAccount entity
type InternalAccessKeyAccount struct {
	Name string
}

// GetName returns the value for the field name
func (e *InternalAccessKeyAccount) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *InternalAccessKeyAccount) SetName(name string) {
	e.Name = name
}

// Hydrate implements struct hydrate
func (e *InternalAccessKeyAccount) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "name", &e.Name); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *InternalAccessKeyAccount) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("name", e.Name)
	return nil
}

// StructPath returns StructPath
func (e *InternalAccessKeyAccount) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalAccessKeyAccount.Path()
}

// NewInternalAccessKey creates a new InternalAccessKey
func NewInternalAccessKey() *InternalAccessKey {
	s := &InternalAccessKey{}
	return s
}

// InternalAccessKey entity
type InternalAccessKey struct {
	AccessKeyID     string
	Account         *InternalAccessKeyAccount
	SecretAccessKey string
	User            *InternalAccessKeyUser
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

// Hydrate implements struct hydrate
func (e *InternalAccessKey) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "accessKeyID", &e.AccessKeyID); err != nil {
		return err
	}
	if err := clientruntime.ContentObjectProperty(ctx.Content(), SpecularMeta().InternalAccessKeyAccount().TypeBuilder(), ctx.Package(), "account", &e.Account); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "secretAccessKey", &e.SecretAccessKey); err != nil {
		return err
	}
	if err := clientruntime.ContentObjectProperty(ctx.Content(), SpecularMeta().InternalAccessKeyUser().TypeBuilder(), ctx.Package(), "user", &e.User); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *InternalAccessKey) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("accessKeyID", e.AccessKeyID)

	var fieldContentAccount clientruntime.Content
	if v := e.Account; v != nil {
		fieldContentAccount = clientruntime.NewContent()
		if err := v.Dehydrate(ctx.CopyWithContent(fieldContentAccount)); err != nil {
			return clientruntime.NewPropertyDehydrationError("account", err)
		}
	}
	if fieldContentAccount == nil {
		return clientruntime.NewPropertyRequiredError("account")
	}
	ctx.Content().SetProperty("account", fieldContentAccount.Map())
	ctx.Content().SetProperty("secretAccessKey", e.SecretAccessKey)

	var fieldContentUser clientruntime.Content
	if v := e.User; v != nil {
		fieldContentUser = clientruntime.NewContent()
		if err := v.Dehydrate(ctx.CopyWithContent(fieldContentUser)); err != nil {
			return clientruntime.NewPropertyDehydrationError("user", err)
		}
	}
	ctx.Content().SetProperty("user", fieldContentUser.Map())
	return nil
}

// StructPath returns StructPath
func (e *InternalAccessKey) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalAccessKey.Path()
}

// NewInternalQueryAssertionEntryValue creates a new InternalQueryAssertionEntryValue
func NewInternalQueryAssertionEntryValue() *InternalQueryAssertionEntryValue {
	s := &InternalQueryAssertionEntryValue{}
	return s
}

// InternalQueryAssertionEntryValue entity
type InternalQueryAssertionEntryValue struct {
	B *bool
	S *string
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

// Hydrate implements struct hydrate
func (e *InternalQueryAssertionEntryValue) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentOptionalBoolProperty(ctx.Content(), "b", &e.B); err != nil {
		return err
	}
	if err := clientruntime.ContentOptionalStringProperty(ctx.Content(), "s", &e.S); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *InternalQueryAssertionEntryValue) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("b", e.B)
	ctx.Content().SetProperty("s", e.S)
	return nil
}

// StructPath returns StructPath
func (e *InternalQueryAssertionEntryValue) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalQueryAssertionEntryValue.Path()
}

// NewInternalQueryAssertionEntry creates a new InternalQueryAssertionEntry
func NewInternalQueryAssertionEntry() *InternalQueryAssertionEntry {
	s := &InternalQueryAssertionEntry{}
	return s
}

// InternalQueryAssertionEntry entity
type InternalQueryAssertionEntry struct {
	Snippet *string
	Value   *InternalQueryAssertionEntryValue
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

// Hydrate implements struct hydrate
func (e *InternalQueryAssertionEntry) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentOptionalStringProperty(ctx.Content(), "snippet", &e.Snippet); err != nil {
		return err
	}
	if err := clientruntime.ContentObjectProperty(ctx.Content(), SpecularMeta().InternalQueryAssertionEntryValue().TypeBuilder(), ctx.Package(), "value", &e.Value); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *InternalQueryAssertionEntry) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("snippet", e.Snippet)

	var fieldContentValue clientruntime.Content
	if v := e.Value; v != nil {
		fieldContentValue = clientruntime.NewContent()
		if err := v.Dehydrate(ctx.CopyWithContent(fieldContentValue)); err != nil {
			return clientruntime.NewPropertyDehydrationError("value", err)
		}
	}
	ctx.Content().SetProperty("value", fieldContentValue.Map())
	return nil
}

// StructPath returns StructPath
func (e *InternalQueryAssertionEntry) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalQueryAssertionEntry.Path()
}

// NewInternalServicesAssertActionCallerForbiddenProblem creates a new InternalServicesAssertActionCallerForbiddenProblem
func NewInternalServicesAssertActionCallerForbiddenProblem() *InternalServicesAssertActionCallerForbiddenProblem {
	s := &InternalServicesAssertActionCallerForbiddenProblem{}
	return s
}

// InternalServicesAssertActionCallerForbiddenProblem entity
type InternalServicesAssertActionCallerForbiddenProblem struct {
	Message string
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

// Hydrate implements struct hydrate
func (e *InternalServicesAssertActionCallerForbiddenProblem) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "message", &e.Message); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *InternalServicesAssertActionCallerForbiddenProblem) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("message", e.Message)
	return nil
}

// StructPath returns StructPath
func (e *InternalServicesAssertActionCallerForbiddenProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalServicesAssertActionCallerForbiddenProblem.Path()
}

// NewInternalServicesValidateAccessKeyInput creates a new InternalServicesValidateAccessKeyInput
func NewInternalServicesValidateAccessKeyInput() *InternalServicesValidateAccessKeyInput {
	s := &InternalServicesValidateAccessKeyInput{}
	return s
}

// InternalServicesValidateAccessKeyInput entity
type InternalServicesValidateAccessKeyInput struct {
	RequestAccessKeyID string
}

// GetRequestAccessKeyID returns the value for the field requestAccessKeyID
func (e *InternalServicesValidateAccessKeyInput) GetRequestAccessKeyID() string {
	return e.RequestAccessKeyID
}

// SetRequestAccessKeyID sets the value for the field requestAccessKeyID
func (e *InternalServicesValidateAccessKeyInput) SetRequestAccessKeyID(requestAccessKeyID string) {
	e.RequestAccessKeyID = requestAccessKeyID
}

// Hydrate implements struct hydrate
func (e *InternalServicesValidateAccessKeyInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "requestAccessKeyID", &e.RequestAccessKeyID); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *InternalServicesValidateAccessKeyInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("requestAccessKeyID", e.RequestAccessKeyID)
	return nil
}

// StructPath returns StructPath
func (e *InternalServicesValidateAccessKeyInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalServicesValidateAccessKeyInput.Path()
}

// NewInternalServicesValidateAccessKeyOutput creates a new InternalServicesValidateAccessKeyOutput
func NewInternalServicesValidateAccessKeyOutput() *InternalServicesValidateAccessKeyOutput {
	s := &InternalServicesValidateAccessKeyOutput{}
	return s
}

// InternalServicesValidateAccessKeyOutput entity
type InternalServicesValidateAccessKeyOutput struct {
	Key *InternalAccessKey
}

// GetKey returns the value for the field key
func (e *InternalServicesValidateAccessKeyOutput) GetKey() *InternalAccessKey {
	return e.Key
}

// SetKey sets the value for the field key
func (e *InternalServicesValidateAccessKeyOutput) SetKey(key *InternalAccessKey) {
	e.Key = key
}

// Hydrate implements struct hydrate
func (e *InternalServicesValidateAccessKeyOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentObjectProperty(ctx.Content(), SpecularMeta().InternalAccessKey().TypeBuilder(), ctx.Package(), "key", &e.Key); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *InternalServicesValidateAccessKeyOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {

	var fieldContentKey clientruntime.Content
	if v := e.Key; v != nil {
		fieldContentKey = clientruntime.NewContent()
		if err := v.Dehydrate(ctx.CopyWithContent(fieldContentKey)); err != nil {
			return clientruntime.NewPropertyDehydrationError("key", err)
		}
	}
	if fieldContentKey == nil {
		return clientruntime.NewPropertyRequiredError("key")
	}
	ctx.Content().SetProperty("key", fieldContentKey.Map())
	return nil
}

// StructPath returns StructPath
func (e *InternalServicesValidateAccessKeyOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalServicesValidateAccessKeyOutput.Path()
}

// NewInternalServicesValidateAccessKeyInvalidAccessKeyProblem creates a new InternalServicesValidateAccessKeyInvalidAccessKeyProblem
func NewInternalServicesValidateAccessKeyInvalidAccessKeyProblem() *InternalServicesValidateAccessKeyInvalidAccessKeyProblem {
	s := &InternalServicesValidateAccessKeyInvalidAccessKeyProblem{}
	return s
}

// InternalServicesValidateAccessKeyInvalidAccessKeyProblem entity
type InternalServicesValidateAccessKeyInvalidAccessKeyProblem struct {
	Message string
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

// Hydrate implements struct hydrate
func (e *InternalServicesValidateAccessKeyInvalidAccessKeyProblem) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "message", &e.Message); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *InternalServicesValidateAccessKeyInvalidAccessKeyProblem) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("message", e.Message)
	return nil
}

// StructPath returns StructPath
func (e *InternalServicesValidateAccessKeyInvalidAccessKeyProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalServicesValidateAccessKeyInvalidAccessKeyProblem.Path()
}

// NewInternalServicesAssertActionInput creates a new InternalServicesAssertActionInput
func NewInternalServicesAssertActionInput() *InternalServicesAssertActionInput {
	s := &InternalServicesAssertActionInput{}
	return s
}

// InternalServicesAssertActionInput entity
type InternalServicesAssertActionInput struct {
	CallerAccessKeyID string
	CallerAction      string
	CallerRegion      string
	CallerResourceDRN string
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

// GetCallerRegion returns the value for the field callerRegion
func (e *InternalServicesAssertActionInput) GetCallerRegion() string {
	return e.CallerRegion
}

// SetCallerRegion sets the value for the field callerRegion
func (e *InternalServicesAssertActionInput) SetCallerRegion(callerRegion string) {
	e.CallerRegion = callerRegion
}

// GetCallerResourceDRN returns the value for the field callerResourceDRN
func (e *InternalServicesAssertActionInput) GetCallerResourceDRN() string {
	return e.CallerResourceDRN
}

// SetCallerResourceDRN sets the value for the field callerResourceDRN
func (e *InternalServicesAssertActionInput) SetCallerResourceDRN(callerResourceDRN string) {
	e.CallerResourceDRN = callerResourceDRN
}

// Hydrate implements struct hydrate
func (e *InternalServicesAssertActionInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "callerAccessKeyID", &e.CallerAccessKeyID); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "callerAction", &e.CallerAction); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "callerRegion", &e.CallerRegion); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "callerResourceDRN", &e.CallerResourceDRN); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *InternalServicesAssertActionInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("callerAccessKeyID", e.CallerAccessKeyID)
	ctx.Content().SetProperty("callerAction", e.CallerAction)
	ctx.Content().SetProperty("callerRegion", e.CallerRegion)
	ctx.Content().SetProperty("callerResourceDRN", e.CallerResourceDRN)
	return nil
}

// StructPath returns StructPath
func (e *InternalServicesAssertActionInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalServicesAssertActionInput.Path()
}

// NewInternalServicesAssertActionOutput creates a new InternalServicesAssertActionOutput
func NewInternalServicesAssertActionOutput() *InternalServicesAssertActionOutput {
	s := &InternalServicesAssertActionOutput{}
	return s
}

// InternalServicesAssertActionOutput entity
type InternalServicesAssertActionOutput struct {
}

// Hydrate implements struct hydrate
func (e *InternalServicesAssertActionOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	return nil
}

// Dehydrate implements struct dehydrate
func (e *InternalServicesAssertActionOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	return nil
}

// StructPath returns StructPath
func (e *InternalServicesAssertActionOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalServicesAssertActionOutput.Path()
}

// NewInternalServicesAssertQueryInput creates a new InternalServicesAssertQueryInput
func NewInternalServicesAssertQueryInput() *InternalServicesAssertQueryInput {
	s := &InternalServicesAssertQueryInput{}
	return s
}

// InternalServicesAssertQueryInput entity
type InternalServicesAssertQueryInput struct {
	CallerAccessKeyID             string
	CallerAction                  string
	CallerRegion                  string
	CallerResourceDRNReplacements string
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

// GetCallerRegion returns the value for the field callerRegion
func (e *InternalServicesAssertQueryInput) GetCallerRegion() string {
	return e.CallerRegion
}

// SetCallerRegion sets the value for the field callerRegion
func (e *InternalServicesAssertQueryInput) SetCallerRegion(callerRegion string) {
	e.CallerRegion = callerRegion
}

// GetCallerResourceDRNReplacements returns the value for the field callerResourceDRNReplacements
func (e *InternalServicesAssertQueryInput) GetCallerResourceDRNReplacements() string {
	return e.CallerResourceDRNReplacements
}

// SetCallerResourceDRNReplacements sets the value for the field callerResourceDRNReplacements
func (e *InternalServicesAssertQueryInput) SetCallerResourceDRNReplacements(callerResourceDRNReplacements string) {
	e.CallerResourceDRNReplacements = callerResourceDRNReplacements
}

// Hydrate implements struct hydrate
func (e *InternalServicesAssertQueryInput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "callerAccessKeyID", &e.CallerAccessKeyID); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "callerAction", &e.CallerAction); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "callerRegion", &e.CallerRegion); err != nil {
		return err
	}
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "callerResourceDRNReplacements", &e.CallerResourceDRNReplacements); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *InternalServicesAssertQueryInput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("callerAccessKeyID", e.CallerAccessKeyID)
	ctx.Content().SetProperty("callerAction", e.CallerAction)
	ctx.Content().SetProperty("callerRegion", e.CallerRegion)
	ctx.Content().SetProperty("callerResourceDRNReplacements", e.CallerResourceDRNReplacements)
	return nil
}

// StructPath returns StructPath
func (e *InternalServicesAssertQueryInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalServicesAssertQueryInput.Path()
}

// NewInternalServicesAssertQueryOutput creates a new InternalServicesAssertQueryOutput
func NewInternalServicesAssertQueryOutput() *InternalServicesAssertQueryOutput {
	s := &InternalServicesAssertQueryOutput{}
	return s
}

// InternalServicesAssertQueryOutput entity
type InternalServicesAssertQueryOutput struct {
	Entries []*InternalQueryAssertionEntry
}

// GetEntries returns the value for the field entries
func (e *InternalServicesAssertQueryOutput) GetEntries() []*InternalQueryAssertionEntry {
	return e.Entries
}

// SetEntries sets the value for the field entries
func (e *InternalServicesAssertQueryOutput) SetEntries(entries []*InternalQueryAssertionEntry) {
	e.Entries = entries
}

// Hydrate implements struct hydrate
func (e *InternalServicesAssertQueryOutput) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentObjectArrayProperty(ctx.Content(), SpecularMeta().InternalQueryAssertionEntry().TypeBuilder(), ctx.Package(), "entries", &e.Entries); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *InternalServicesAssertQueryOutput) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	var fieldValuesEntries []map[string]any

	if fsv := e.Entries; fsv != nil {
		fieldValuesEntries = make([]map[string]any, 0, len(fsv))
		for ix, fsvitem := range fsv {
			var fsvm map[string]any
			if fsvitem != nil {
				ct := clientruntime.NewContent()
				if err := fsvitem.Dehydrate(ctx.CopyWithContent(ct)); err != nil {
					return clientruntime.NewPropertyItemDehydrationError("entries", ix, err)
				}
				fsvm = ct.Map()
			}
			if fsvm == nil {
				return clientruntime.NewArrayItemMissingError("entries", ix)
			}
			fieldValuesEntries = append(fieldValuesEntries, fsvm)
		}
	}

	ctx.Content().SetProperty("entries", fieldValuesEntries)
	return nil
}

// StructPath returns StructPath
func (e *InternalServicesAssertQueryOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalServicesAssertQueryOutput.Path()
}

// NewInternalServicesAssertQueryReplacementProblem creates a new InternalServicesAssertQueryReplacementProblem
func NewInternalServicesAssertQueryReplacementProblem() *InternalServicesAssertQueryReplacementProblem {
	s := &InternalServicesAssertQueryReplacementProblem{}
	return s
}

// InternalServicesAssertQueryReplacementProblem entity
type InternalServicesAssertQueryReplacementProblem struct {
	Message string
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

// Hydrate implements struct hydrate
func (e *InternalServicesAssertQueryReplacementProblem) Hydrate(ctx *clientruntime.HydratationContext) error {
	if err := clientruntime.ContentRequireStringProperty(ctx.Content(), "message", &e.Message); err != nil {
		return err
	}
	return nil
}

// Dehydrate implements struct dehydrate
func (e *InternalServicesAssertQueryReplacementProblem) Dehydrate(ctx *clientruntime.DehydrationContext) (err error) {
	ctx.Content().SetProperty("message", e.Message)
	return nil
}

// StructPath returns StructPath
func (e *InternalServicesAssertQueryReplacementProblem) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInternalServicesAssertQueryReplacementProblem.Path()
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
	localSpecularMeta.structPathPolicyStructureProblem, err = pk.NewType(
		"PolicyStructureProblem",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewPolicyStructureProblem()
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
	localSpecularMeta.structPathInternalAccessKeyUser, err = pk.NewType(
		"InternalAccessKeyUser",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInternalAccessKeyUser()
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

	op.SetInput(SpecularMeta().AccountCreateInput())
	op.SetOutput(SpecularMeta().AccountCreateInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblem())
	op.RegisterProblemType(SpecularMeta().AccountCreateInvalidNameProblem())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resAccount.NewOperation("AssumeIdentity")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountAssumeIdentityInput())
	op.SetOutput(SpecularMeta().AccountAssumeIdentityInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblem())

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

	op.SetInput(SpecularMeta().AccountSSOBeginAuthenticationInput())
	op.SetOutput(SpecularMeta().AccountSSOBeginAuthenticationInput())
	op.RegisterProblemType(SpecularMeta().SSOProviderUnavailableProblem())
	op.RegisterProblemType(SpecularMeta().AccountSSOBeginAuthenticationParameterProblem())

	op, err = resAccountSSO.NewOperation("CompleteAuthentication")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountSSOCompleteAuthenticationInput())
	op.SetOutput(SpecularMeta().AccountSSOCompleteAuthenticationInput())
	op.RegisterProblemType(SpecularMeta().AccountSSOCompleteAuthenticationInvalidFlowProblem())

	op, err = resAccountSSO.NewOperation("GetProviders")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountSSOGetProvidersInput())
	op.SetOutput(SpecularMeta().AccountSSOGetProvidersInput())

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

	op.SetInput(SpecularMeta().AccountSSOAutoJoinPolicyCreateInput())
	op.SetOutput(SpecularMeta().AccountSSOAutoJoinPolicyCreateInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblem())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resAccountSSOAutoJoinPolicy.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountSSOAutoJoinPolicyListInput())
	op.SetOutput(SpecularMeta().AccountSSOAutoJoinPolicyListInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblem())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resAccountSSOAutoJoinPolicy.NewOperation("Enable")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountSSOAutoJoinPolicyEnableInput())
	op.SetOutput(SpecularMeta().AccountSSOAutoJoinPolicyEnableInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblem())

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

	op.SetInput(SpecularMeta().UserCreateInput())
	op.SetOutput(SpecularMeta().UserCreateInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblem())
	op.RegisterProblemType(SpecularMeta().InvalidUsernameProblem())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUser.NewOperation("Get")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserGetInput())
	op.SetOutput(SpecularMeta().UserGetInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(SpecularMeta().UserGetUserNotAvailableProblem())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUser.NewOperation("Destroy")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserDestroyInput())
	op.SetOutput(SpecularMeta().UserDestroyInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblem())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUser.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserListInput())
	op.SetOutput(SpecularMeta().UserListInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblem())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUser.NewOperation("MemberAccounts")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserMemberAccountsInput())
	op.SetOutput(SpecularMeta().UserMemberAccountsInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())

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

	op.SetInput(SpecularMeta().UserAccessKeyCreateInput())
	op.SetOutput(SpecularMeta().UserAccessKeyCreateInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblem())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUserAccessKey.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserAccessKeyListInput())
	op.SetOutput(SpecularMeta().UserAccessKeyListInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblem())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUserAccessKey.NewOperation("Destroy")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserAccessKeyDestroyInput())
	op.SetOutput(SpecularMeta().UserAccessKeyDestroyInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblem())

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

	op.SetInput(SpecularMeta().UserIdentityPolicyAttachInput())
	op.SetOutput(SpecularMeta().UserIdentityPolicyAttachInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblem())
	op.RegisterProblemType(SpecularMeta().PolicyNotFoundProblem())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUserIdentityPolicy.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserIdentityPolicyListInput())
	op.SetOutput(SpecularMeta().UserIdentityPolicyListInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblem())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUserIdentityPolicy.NewOperation("Detach")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserIdentityPolicyDetachInput())
	op.SetOutput(SpecularMeta().UserIdentityPolicyDetachInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblem())
	op.RegisterProblemType(SpecularMeta().PolicyNotFoundProblem())

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

	op.SetInput(SpecularMeta().IdentityPolicyCreateInput())
	op.SetOutput(SpecularMeta().IdentityPolicyCreateInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblem())
	op.RegisterProblemType(SpecularMeta().PolicyStructureProblem())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resIdentityPolicy.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().IdentityPolicyListInput())
	op.SetOutput(SpecularMeta().IdentityPolicyListInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblem())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resIdentityPolicy.NewOperation("Retrieve")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().IdentityPolicyRetrieveInput())
	op.SetOutput(SpecularMeta().IdentityPolicyRetrieveInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblem())
	op.RegisterProblemType(SpecularMeta().PolicyNotFoundProblem())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resIdentityPolicy.NewOperation("Destroy")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().IdentityPolicyDestroyInput())
	op.SetOutput(SpecularMeta().IdentityPolicyDestroyInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblem())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resIdentityPolicy.NewOperation("Update")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().IdentityPolicyUpdateInput())
	op.SetOutput(SpecularMeta().IdentityPolicyUpdateInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblem())
	op.RegisterProblemType(SpecularMeta().PolicyStructureProblem())

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

	op.SetInput(SpecularMeta().InternalServicesValidateAccessKeyInput())
	op.SetOutput(SpecularMeta().InternalServicesValidateAccessKeyInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblem())
	op.RegisterProblemType(SpecularMeta().InternalServicesValidateAccessKeyInvalidAccessKeyProblem())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resInternalServices.NewOperation("AssertAction")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().InternalServicesAssertActionInput())
	op.SetOutput(SpecularMeta().InternalServicesAssertActionInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblem())
	op.RegisterProblemType(SpecularMeta().InternalServicesAssertActionCallerForbiddenProblem())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resInternalServices.NewOperation("AssertQuery")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().InternalServicesAssertQueryInput())
	op.SetOutput(SpecularMeta().InternalServicesAssertQueryInput())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedProblem())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenProblem())
	op.RegisterProblemType(SpecularMeta().InternalServicesAssertActionCallerForbiddenProblem())
	op.RegisterProblemType(SpecularMeta().InternalServicesAssertQueryReplacementProblem())

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

// Create entity
// Creates a new Account with the given name
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

// AssumeIdentity entity
// Assume identity in the given account
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

// BeginAuthentication entity
// Returns the public link to the login page for the given SSO provider
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

// CompleteAuthentication entity
// Authorizes the user with the given code and returns a token
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

// GetProviders entity
// Returns the providers available for the given account
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

// Create entity
// Creates an auto-join policy in the account
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

// List entity
// Lists policies in the account that allow users in other accounts to join automatically
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

// Enable entity
// Enables the auto-join policy
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

// Create entity
// Creates a new user with the given username in the current account
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

// Get entity
// Returns information about the current user or the user with the given username
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

// Destroy entity
// Destroys the user with the given username
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

// List entity
// Returns the list of users in the current account
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

// MemberAccounts entity
// Returns the accounts the user is member of
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

// Create entity
// Creates a new access key for the given user
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

// List entity
// Returns the list of access keys for the given user or the current user
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

// Destroy entity
// Destroy an access key for the given user
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

// Attach entity
// Attaches an identity policy to a user
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

// List entity
// List returns the identity policies of a user
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

// Detach entity
// Deatach an identity policy from a user
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

// Create entity
// Creates a new identity policy
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

// List entity
// Retrieves list of identity policies
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

// Retrieve entity
// Retrieves identity policy by name
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

// Destroy entity
// Destroys an identity policy
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

// Update entity
// Updates an identity policy
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
	transport         clientruntime.Transport
	res               *clientruntime.Resource
	validateAccessKey *clientruntime.Operation
	assertAction      *clientruntime.Operation
	assertQuery       *clientruntime.Operation
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
	r.assertAction = res.FindOperation("AssertAction")
	r.assertQuery = res.FindOperation("AssertQuery")
	return r, nil
}

// ValidateAccessKey entity
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

// AssertAction entity
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

// AssertQuery entity
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

// Client is the main client of the package
type Client struct {
	transport      clientruntime.Transport
	pk             *clientruntime.Package
	Account        *AccountResource
	User           *UserResource
	IdentityPolicy *IdentityPolicyResource
	Internal       *InternalResource
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
	c.IdentityPolicy, err = newIdentityPolicyResource(transport, pk)
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
	mod                                                                *clientruntime.Package
	structPathUserInformationSSOProvider                               *clientruntime.StructDefinition
	structPathUserInformationSSOProfile                                *clientruntime.StructDefinition
	structPathUserInformationSSO                                       *clientruntime.StructDefinition
	structPathUserInformation                                          *clientruntime.StructDefinition
	structPathCredentials                                              *clientruntime.StructDefinition
	structPathSSOProviderUnavailableProblem                            *clientruntime.StructDefinition
	structPathSSOFlow                                                  *clientruntime.StructDefinition
	structPathAccount                                                  *clientruntime.StructDefinition
	structPathAccountSSOProvider                                       *clientruntime.StructDefinition
	structPathPolicyNotFoundProblem                                    *clientruntime.StructDefinition
	structPathAccountSSOAutoJoinPolicy                                 *clientruntime.StructDefinition
	structPathAccountCreateInput                                       *clientruntime.StructDefinition
	structPathAccountCreateOutput                                      *clientruntime.StructDefinition
	structPathAccountCreateInvalidNameProblem                          *clientruntime.StructDefinition
	structPathAccountAssumeIdentityInput                               *clientruntime.StructDefinition
	structPathAccountAssumeIdentityOutput                              *clientruntime.StructDefinition
	structPathAccountSSOBeginAuthenticationInput                       *clientruntime.StructDefinition
	structPathAccountSSOBeginAuthenticationOutput                      *clientruntime.StructDefinition
	structPathAccountSSOBeginAuthenticationParameterProblem            *clientruntime.StructDefinition
	structPathAccountSSOCompleteAuthenticationInput                    *clientruntime.StructDefinition
	structPathAccountSSOCompleteAuthenticationOutput                   *clientruntime.StructDefinition
	structPathAccountSSOCompleteAuthenticationInvalidFlowProblem       *clientruntime.StructDefinition
	structPathAccountSSOGetProvidersInput                              *clientruntime.StructDefinition
	structPathAccountSSOGetProvidersOutput                             *clientruntime.StructDefinition
	structPathAccountSSOAutoJoinPolicyCreateInput                      *clientruntime.StructDefinition
	structPathAccountSSOAutoJoinPolicyCreateOutput                     *clientruntime.StructDefinition
	structPathAccountSSOAutoJoinPolicyListInput                        *clientruntime.StructDefinition
	structPathAccountSSOAutoJoinPolicyListOutput                       *clientruntime.StructDefinition
	structPathAccountSSOAutoJoinPolicyEnableInput                      *clientruntime.StructDefinition
	structPathAccountSSOAutoJoinPolicyEnableOutput                     *clientruntime.StructDefinition
	structPathInvalidUsernameProblem                                   *clientruntime.StructDefinition
	structPathMemberAccount                                            *clientruntime.StructDefinition
	structPathCredentialInfo                                           *clientruntime.StructDefinition
	structPathIdentityPolicyStatement                                  *clientruntime.StructDefinition
	structPathIdentityPolicy                                           *clientruntime.StructDefinition
	structPathIdentityPolicyAttachment                                 *clientruntime.StructDefinition
	structPathUserCreateInput                                          *clientruntime.StructDefinition
	structPathUserCreateOutput                                         *clientruntime.StructDefinition
	structPathUserGetInput                                             *clientruntime.StructDefinition
	structPathUserGetOutput                                            *clientruntime.StructDefinition
	structPathUserGetUserNotAvailableProblem                           *clientruntime.StructDefinition
	structPathUserDestroyInput                                         *clientruntime.StructDefinition
	structPathUserDestroyOutput                                        *clientruntime.StructDefinition
	structPathUserListInput                                            *clientruntime.StructDefinition
	structPathUserListOutput                                           *clientruntime.StructDefinition
	structPathUserMemberAccountsInput                                  *clientruntime.StructDefinition
	structPathUserMemberAccountsOutput                                 *clientruntime.StructDefinition
	structPathUserAccessKeyCreateInput                                 *clientruntime.StructDefinition
	structPathUserAccessKeyCreateOutput                                *clientruntime.StructDefinition
	structPathUserAccessKeyListInput                                   *clientruntime.StructDefinition
	structPathUserAccessKeyListOutput                                  *clientruntime.StructDefinition
	structPathUserAccessKeyDestroyInput                                *clientruntime.StructDefinition
	structPathUserAccessKeyDestroyOutput                               *clientruntime.StructDefinition
	structPathUserIdentityPolicyAttachInput                            *clientruntime.StructDefinition
	structPathUserIdentityPolicyAttachOutput                           *clientruntime.StructDefinition
	structPathUserIdentityPolicyListInput                              *clientruntime.StructDefinition
	structPathUserIdentityPolicyListOutput                             *clientruntime.StructDefinition
	structPathUserIdentityPolicyDetachInput                            *clientruntime.StructDefinition
	structPathUserIdentityPolicyDetachOutput                           *clientruntime.StructDefinition
	structPathPolicyStructureProblem                                   *clientruntime.StructDefinition
	structPathIdentityPolicyCreateInput                                *clientruntime.StructDefinition
	structPathIdentityPolicyCreateOutput                               *clientruntime.StructDefinition
	structPathIdentityPolicyListInput                                  *clientruntime.StructDefinition
	structPathIdentityPolicyListOutput                                 *clientruntime.StructDefinition
	structPathIdentityPolicyRetrieveInput                              *clientruntime.StructDefinition
	structPathIdentityPolicyRetrieveOutput                             *clientruntime.StructDefinition
	structPathIdentityPolicyDestroyInput                               *clientruntime.StructDefinition
	structPathIdentityPolicyDestroyOutput                              *clientruntime.StructDefinition
	structPathIdentityPolicyUpdateInput                                *clientruntime.StructDefinition
	structPathIdentityPolicyUpdateOutput                               *clientruntime.StructDefinition
	structPathInternalAccessKeyUser                                    *clientruntime.StructDefinition
	structPathInternalAccessKeyAccount                                 *clientruntime.StructDefinition
	structPathInternalAccessKey                                        *clientruntime.StructDefinition
	structPathInternalQueryAssertionEntryValue                         *clientruntime.StructDefinition
	structPathInternalQueryAssertionEntry                              *clientruntime.StructDefinition
	structPathInternalServicesAssertActionCallerForbiddenProblem       *clientruntime.StructDefinition
	structPathInternalServicesValidateAccessKeyInput                   *clientruntime.StructDefinition
	structPathInternalServicesValidateAccessKeyOutput                  *clientruntime.StructDefinition
	structPathInternalServicesValidateAccessKeyInvalidAccessKeyProblem *clientruntime.StructDefinition
	structPathInternalServicesAssertActionInput                        *clientruntime.StructDefinition
	structPathInternalServicesAssertActionOutput                       *clientruntime.StructDefinition
	structPathInternalServicesAssertQueryInput                         *clientruntime.StructDefinition
	structPathInternalServicesAssertQueryOutput                        *clientruntime.StructDefinition
	structPathInternalServicesAssertQueryReplacementProblem            *clientruntime.StructDefinition
}

// Module returns the module definition
func (m *SpecularMetaInfo) Module() *clientruntime.Package {
	return m.mod
}

// UserInformationSSOProvider allows easy access to structure
func (m *SpecularMetaInfo) UserInformationSSOProvider() *clientruntime.StructDefinition {
	return m.structPathUserInformationSSOProvider
}

// UserInformationSSOProfile allows easy access to structure
func (m *SpecularMetaInfo) UserInformationSSOProfile() *clientruntime.StructDefinition {
	return m.structPathUserInformationSSOProfile
}

// UserInformationSSO allows easy access to structure
func (m *SpecularMetaInfo) UserInformationSSO() *clientruntime.StructDefinition {
	return m.structPathUserInformationSSO
}

// UserInformation allows easy access to structure
func (m *SpecularMetaInfo) UserInformation() *clientruntime.StructDefinition {
	return m.structPathUserInformation
}

// Credentials allows easy access to structure
func (m *SpecularMetaInfo) Credentials() *clientruntime.StructDefinition {
	return m.structPathCredentials
}

// SSOProviderUnavailableProblem allows easy access to structure
func (m *SpecularMetaInfo) SSOProviderUnavailableProblem() *clientruntime.StructDefinition {
	return m.structPathSSOProviderUnavailableProblem
}

// SSOFlow allows easy access to structure
func (m *SpecularMetaInfo) SSOFlow() *clientruntime.StructDefinition {
	return m.structPathSSOFlow
}

// Account allows easy access to structure
func (m *SpecularMetaInfo) Account() *clientruntime.StructDefinition {
	return m.structPathAccount
}

// AccountSSOProvider allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOProvider() *clientruntime.StructDefinition {
	return m.structPathAccountSSOProvider
}

// PolicyNotFoundProblem allows easy access to structure
func (m *SpecularMetaInfo) PolicyNotFoundProblem() *clientruntime.StructDefinition {
	return m.structPathPolicyNotFoundProblem
}

// AccountSSOAutoJoinPolicy allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOAutoJoinPolicy() *clientruntime.StructDefinition {
	return m.structPathAccountSSOAutoJoinPolicy
}

// AccountCreateInput allows easy access to structure
func (m *SpecularMetaInfo) AccountCreateInput() *clientruntime.StructDefinition {
	return m.structPathAccountCreateInput
}

// AccountCreateOutput allows easy access to structure
func (m *SpecularMetaInfo) AccountCreateOutput() *clientruntime.StructDefinition {
	return m.structPathAccountCreateOutput
}

// AccountCreateInvalidNameProblem allows easy access to structure
func (m *SpecularMetaInfo) AccountCreateInvalidNameProblem() *clientruntime.StructDefinition {
	return m.structPathAccountCreateInvalidNameProblem
}

// AccountAssumeIdentityInput allows easy access to structure
func (m *SpecularMetaInfo) AccountAssumeIdentityInput() *clientruntime.StructDefinition {
	return m.structPathAccountAssumeIdentityInput
}

// AccountAssumeIdentityOutput allows easy access to structure
func (m *SpecularMetaInfo) AccountAssumeIdentityOutput() *clientruntime.StructDefinition {
	return m.structPathAccountAssumeIdentityOutput
}

// AccountSSOBeginAuthenticationInput allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOBeginAuthenticationInput() *clientruntime.StructDefinition {
	return m.structPathAccountSSOBeginAuthenticationInput
}

// AccountSSOBeginAuthenticationOutput allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOBeginAuthenticationOutput() *clientruntime.StructDefinition {
	return m.structPathAccountSSOBeginAuthenticationOutput
}

// AccountSSOBeginAuthenticationParameterProblem allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOBeginAuthenticationParameterProblem() *clientruntime.StructDefinition {
	return m.structPathAccountSSOBeginAuthenticationParameterProblem
}

// AccountSSOCompleteAuthenticationInput allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOCompleteAuthenticationInput() *clientruntime.StructDefinition {
	return m.structPathAccountSSOCompleteAuthenticationInput
}

// AccountSSOCompleteAuthenticationOutput allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOCompleteAuthenticationOutput() *clientruntime.StructDefinition {
	return m.structPathAccountSSOCompleteAuthenticationOutput
}

// AccountSSOCompleteAuthenticationInvalidFlowProblem allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOCompleteAuthenticationInvalidFlowProblem() *clientruntime.StructDefinition {
	return m.structPathAccountSSOCompleteAuthenticationInvalidFlowProblem
}

// AccountSSOGetProvidersInput allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOGetProvidersInput() *clientruntime.StructDefinition {
	return m.structPathAccountSSOGetProvidersInput
}

// AccountSSOGetProvidersOutput allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOGetProvidersOutput() *clientruntime.StructDefinition {
	return m.structPathAccountSSOGetProvidersOutput
}

// AccountSSOAutoJoinPolicyCreateInput allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOAutoJoinPolicyCreateInput() *clientruntime.StructDefinition {
	return m.structPathAccountSSOAutoJoinPolicyCreateInput
}

// AccountSSOAutoJoinPolicyCreateOutput allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOAutoJoinPolicyCreateOutput() *clientruntime.StructDefinition {
	return m.structPathAccountSSOAutoJoinPolicyCreateOutput
}

// AccountSSOAutoJoinPolicyListInput allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOAutoJoinPolicyListInput() *clientruntime.StructDefinition {
	return m.structPathAccountSSOAutoJoinPolicyListInput
}

// AccountSSOAutoJoinPolicyListOutput allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOAutoJoinPolicyListOutput() *clientruntime.StructDefinition {
	return m.structPathAccountSSOAutoJoinPolicyListOutput
}

// AccountSSOAutoJoinPolicyEnableInput allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOAutoJoinPolicyEnableInput() *clientruntime.StructDefinition {
	return m.structPathAccountSSOAutoJoinPolicyEnableInput
}

// AccountSSOAutoJoinPolicyEnableOutput allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOAutoJoinPolicyEnableOutput() *clientruntime.StructDefinition {
	return m.structPathAccountSSOAutoJoinPolicyEnableOutput
}

// InvalidUsernameProblem allows easy access to structure
func (m *SpecularMetaInfo) InvalidUsernameProblem() *clientruntime.StructDefinition {
	return m.structPathInvalidUsernameProblem
}

// MemberAccount allows easy access to structure
func (m *SpecularMetaInfo) MemberAccount() *clientruntime.StructDefinition {
	return m.structPathMemberAccount
}

// CredentialInfo allows easy access to structure
func (m *SpecularMetaInfo) CredentialInfo() *clientruntime.StructDefinition {
	return m.structPathCredentialInfo
}

// IdentityPolicyStatement allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyStatement() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyStatement
}

// IdentityPolicy allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicy() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicy
}

// IdentityPolicyAttachment allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyAttachment() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyAttachment
}

// UserCreateInput allows easy access to structure
func (m *SpecularMetaInfo) UserCreateInput() *clientruntime.StructDefinition {
	return m.structPathUserCreateInput
}

// UserCreateOutput allows easy access to structure
func (m *SpecularMetaInfo) UserCreateOutput() *clientruntime.StructDefinition {
	return m.structPathUserCreateOutput
}

// UserGetInput allows easy access to structure
func (m *SpecularMetaInfo) UserGetInput() *clientruntime.StructDefinition {
	return m.structPathUserGetInput
}

// UserGetOutput allows easy access to structure
func (m *SpecularMetaInfo) UserGetOutput() *clientruntime.StructDefinition {
	return m.structPathUserGetOutput
}

// UserGetUserNotAvailableProblem allows easy access to structure
func (m *SpecularMetaInfo) UserGetUserNotAvailableProblem() *clientruntime.StructDefinition {
	return m.structPathUserGetUserNotAvailableProblem
}

// UserDestroyInput allows easy access to structure
func (m *SpecularMetaInfo) UserDestroyInput() *clientruntime.StructDefinition {
	return m.structPathUserDestroyInput
}

// UserDestroyOutput allows easy access to structure
func (m *SpecularMetaInfo) UserDestroyOutput() *clientruntime.StructDefinition {
	return m.structPathUserDestroyOutput
}

// UserListInput allows easy access to structure
func (m *SpecularMetaInfo) UserListInput() *clientruntime.StructDefinition {
	return m.structPathUserListInput
}

// UserListOutput allows easy access to structure
func (m *SpecularMetaInfo) UserListOutput() *clientruntime.StructDefinition {
	return m.structPathUserListOutput
}

// UserMemberAccountsInput allows easy access to structure
func (m *SpecularMetaInfo) UserMemberAccountsInput() *clientruntime.StructDefinition {
	return m.structPathUserMemberAccountsInput
}

// UserMemberAccountsOutput allows easy access to structure
func (m *SpecularMetaInfo) UserMemberAccountsOutput() *clientruntime.StructDefinition {
	return m.structPathUserMemberAccountsOutput
}

// UserAccessKeyCreateInput allows easy access to structure
func (m *SpecularMetaInfo) UserAccessKeyCreateInput() *clientruntime.StructDefinition {
	return m.structPathUserAccessKeyCreateInput
}

// UserAccessKeyCreateOutput allows easy access to structure
func (m *SpecularMetaInfo) UserAccessKeyCreateOutput() *clientruntime.StructDefinition {
	return m.structPathUserAccessKeyCreateOutput
}

// UserAccessKeyListInput allows easy access to structure
func (m *SpecularMetaInfo) UserAccessKeyListInput() *clientruntime.StructDefinition {
	return m.structPathUserAccessKeyListInput
}

// UserAccessKeyListOutput allows easy access to structure
func (m *SpecularMetaInfo) UserAccessKeyListOutput() *clientruntime.StructDefinition {
	return m.structPathUserAccessKeyListOutput
}

// UserAccessKeyDestroyInput allows easy access to structure
func (m *SpecularMetaInfo) UserAccessKeyDestroyInput() *clientruntime.StructDefinition {
	return m.structPathUserAccessKeyDestroyInput
}

// UserAccessKeyDestroyOutput allows easy access to structure
func (m *SpecularMetaInfo) UserAccessKeyDestroyOutput() *clientruntime.StructDefinition {
	return m.structPathUserAccessKeyDestroyOutput
}

// UserIdentityPolicyAttachInput allows easy access to structure
func (m *SpecularMetaInfo) UserIdentityPolicyAttachInput() *clientruntime.StructDefinition {
	return m.structPathUserIdentityPolicyAttachInput
}

// UserIdentityPolicyAttachOutput allows easy access to structure
func (m *SpecularMetaInfo) UserIdentityPolicyAttachOutput() *clientruntime.StructDefinition {
	return m.structPathUserIdentityPolicyAttachOutput
}

// UserIdentityPolicyListInput allows easy access to structure
func (m *SpecularMetaInfo) UserIdentityPolicyListInput() *clientruntime.StructDefinition {
	return m.structPathUserIdentityPolicyListInput
}

// UserIdentityPolicyListOutput allows easy access to structure
func (m *SpecularMetaInfo) UserIdentityPolicyListOutput() *clientruntime.StructDefinition {
	return m.structPathUserIdentityPolicyListOutput
}

// UserIdentityPolicyDetachInput allows easy access to structure
func (m *SpecularMetaInfo) UserIdentityPolicyDetachInput() *clientruntime.StructDefinition {
	return m.structPathUserIdentityPolicyDetachInput
}

// UserIdentityPolicyDetachOutput allows easy access to structure
func (m *SpecularMetaInfo) UserIdentityPolicyDetachOutput() *clientruntime.StructDefinition {
	return m.structPathUserIdentityPolicyDetachOutput
}

// PolicyStructureProblem allows easy access to structure
func (m *SpecularMetaInfo) PolicyStructureProblem() *clientruntime.StructDefinition {
	return m.structPathPolicyStructureProblem
}

// IdentityPolicyCreateInput allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyCreateInput() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyCreateInput
}

// IdentityPolicyCreateOutput allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyCreateOutput() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyCreateOutput
}

// IdentityPolicyListInput allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyListInput() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyListInput
}

// IdentityPolicyListOutput allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyListOutput() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyListOutput
}

// IdentityPolicyRetrieveInput allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyRetrieveInput() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyRetrieveInput
}

// IdentityPolicyRetrieveOutput allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyRetrieveOutput() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyRetrieveOutput
}

// IdentityPolicyDestroyInput allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyDestroyInput() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyDestroyInput
}

// IdentityPolicyDestroyOutput allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyDestroyOutput() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyDestroyOutput
}

// IdentityPolicyUpdateInput allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyUpdateInput() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyUpdateInput
}

// IdentityPolicyUpdateOutput allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyUpdateOutput() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyUpdateOutput
}

// InternalAccessKeyUser allows easy access to structure
func (m *SpecularMetaInfo) InternalAccessKeyUser() *clientruntime.StructDefinition {
	return m.structPathInternalAccessKeyUser
}

// InternalAccessKeyAccount allows easy access to structure
func (m *SpecularMetaInfo) InternalAccessKeyAccount() *clientruntime.StructDefinition {
	return m.structPathInternalAccessKeyAccount
}

// InternalAccessKey allows easy access to structure
func (m *SpecularMetaInfo) InternalAccessKey() *clientruntime.StructDefinition {
	return m.structPathInternalAccessKey
}

// InternalQueryAssertionEntryValue allows easy access to structure
func (m *SpecularMetaInfo) InternalQueryAssertionEntryValue() *clientruntime.StructDefinition {
	return m.structPathInternalQueryAssertionEntryValue
}

// InternalQueryAssertionEntry allows easy access to structure
func (m *SpecularMetaInfo) InternalQueryAssertionEntry() *clientruntime.StructDefinition {
	return m.structPathInternalQueryAssertionEntry
}

// InternalServicesAssertActionCallerForbiddenProblem allows easy access to structure
func (m *SpecularMetaInfo) InternalServicesAssertActionCallerForbiddenProblem() *clientruntime.StructDefinition {
	return m.structPathInternalServicesAssertActionCallerForbiddenProblem
}

// InternalServicesValidateAccessKeyInput allows easy access to structure
func (m *SpecularMetaInfo) InternalServicesValidateAccessKeyInput() *clientruntime.StructDefinition {
	return m.structPathInternalServicesValidateAccessKeyInput
}

// InternalServicesValidateAccessKeyOutput allows easy access to structure
func (m *SpecularMetaInfo) InternalServicesValidateAccessKeyOutput() *clientruntime.StructDefinition {
	return m.structPathInternalServicesValidateAccessKeyOutput
}

// InternalServicesValidateAccessKeyInvalidAccessKeyProblem allows easy access to structure
func (m *SpecularMetaInfo) InternalServicesValidateAccessKeyInvalidAccessKeyProblem() *clientruntime.StructDefinition {
	return m.structPathInternalServicesValidateAccessKeyInvalidAccessKeyProblem
}

// InternalServicesAssertActionInput allows easy access to structure
func (m *SpecularMetaInfo) InternalServicesAssertActionInput() *clientruntime.StructDefinition {
	return m.structPathInternalServicesAssertActionInput
}

// InternalServicesAssertActionOutput allows easy access to structure
func (m *SpecularMetaInfo) InternalServicesAssertActionOutput() *clientruntime.StructDefinition {
	return m.structPathInternalServicesAssertActionOutput
}

// InternalServicesAssertQueryInput allows easy access to structure
func (m *SpecularMetaInfo) InternalServicesAssertQueryInput() *clientruntime.StructDefinition {
	return m.structPathInternalServicesAssertQueryInput
}

// InternalServicesAssertQueryOutput allows easy access to structure
func (m *SpecularMetaInfo) InternalServicesAssertQueryOutput() *clientruntime.StructDefinition {
	return m.structPathInternalServicesAssertQueryOutput
}

// InternalServicesAssertQueryReplacementProblem allows easy access to structure
func (m *SpecularMetaInfo) InternalServicesAssertQueryReplacementProblem() *clientruntime.StructDefinition {
	return m.structPathInternalServicesAssertQueryReplacementProblem
}

var localSpecularMeta *SpecularMetaInfo = &SpecularMetaInfo{}

// SpecularMeta returns metadata of the specular module
func SpecularMeta() *SpecularMetaInfo {
	return localSpecularMeta
}
