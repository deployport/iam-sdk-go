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

// NewServiceCatalogAction creates a new ServiceCatalogAction
func NewServiceCatalogAction() *ServiceCatalogAction {
	s := &ServiceCatalogAction{}
	s.InitializeDefaults()
	return s
}

// ServiceCatalogAction - Service Catalog: global reference data describing each service's actions and
// resource types, used by the console's visual/JSON identity-policy builder.
type ServiceCatalogAction struct {
	// resource-type keys this action can target (scopes the resource picker)
	AppliesToResourceTypes []string `json:"appliesToResourceTypes,omitempty" yaml:"appliesToResourceTypes,omitempty"`
	// true when this action belongs in the account's baseline group, granted over
	// every resource. False, the default, means it does not
	Baseline    bool    `json:"baseline,omitempty" yaml:"baseline,omitempty"`
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
	// short action name, e.g. "CreateUser"; the wire action is "<namespace>:<name>"
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetAppliesToResourceTypes returns the value for the field appliesToResourceTypes
func (e *ServiceCatalogAction) GetAppliesToResourceTypes() []string {
	return e.AppliesToResourceTypes
}

// SetAppliesToResourceTypes sets the value for the field appliesToResourceTypes
func (e *ServiceCatalogAction) SetAppliesToResourceTypes(appliesToResourceTypes []string) {
	e.AppliesToResourceTypes = appliesToResourceTypes
}

// GetBaseline returns the value for the field baseline
func (e *ServiceCatalogAction) GetBaseline() bool {
	return e.Baseline
}

// SetBaseline sets the value for the field baseline
func (e *ServiceCatalogAction) SetBaseline(baseline bool) {
	e.Baseline = baseline
}

// GetDescription returns the value for the field description
func (e *ServiceCatalogAction) GetDescription() *string {
	return e.Description
}

// SetDescription sets the value for the field description
func (e *ServiceCatalogAction) SetDescription(description *string) {
	e.Description = description
}

// GetName returns the value for the field name
func (e *ServiceCatalogAction) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *ServiceCatalogAction) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *ServiceCatalogAction) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathServiceCatalogAction.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *ServiceCatalogAction) InitializeDefaults() {
}

// serviceCatalogActionAlias is defined to help pre and post JSON marshaling without recursive loops
type serviceCatalogActionAlias ServiceCatalogAction

// UnmarshalJSON implements json.Unmarshaler
func (e *ServiceCatalogAction) UnmarshalJSON(data []byte) error {
	var alias serviceCatalogActionAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*ServiceCatalogAction)(&alias)).InitializeDefaults()
	*e = ServiceCatalogAction(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e ServiceCatalogAction) MarshalJSON() ([]byte, error) {
	alias := serviceCatalogActionAlias(e)
	return json.Marshal(alias)
}

// NewServiceCatalogResourceType creates a new ServiceCatalogResourceType
func NewServiceCatalogResourceType() *ServiceCatalogResourceType {
	s := &ServiceCatalogResourceType{}
	s.InitializeDefaults()
	return s
}

// ServiceCatalogResourceType struct
type ServiceCatalogResourceType struct {
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
	// resource-builder template, e.g. "iam:User(${UserName})"
	DrnFormat string `json:"drnFormat,omitempty" yaml:"drnFormat,omitempty"`
	// DRN function path, e.g. "User" or "User.AccessKey"
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
	// parent resource-type key (unset for a root type)
	Parent *string `json:"parent,omitempty" yaml:"parent,omitempty"`
}

// GetDescription returns the value for the field description
func (e *ServiceCatalogResourceType) GetDescription() *string {
	return e.Description
}

// SetDescription sets the value for the field description
func (e *ServiceCatalogResourceType) SetDescription(description *string) {
	e.Description = description
}

// GetDrnFormat returns the value for the field drnFormat
func (e *ServiceCatalogResourceType) GetDrnFormat() string {
	return e.DrnFormat
}

// SetDrnFormat sets the value for the field drnFormat
func (e *ServiceCatalogResourceType) SetDrnFormat(drnFormat string) {
	e.DrnFormat = drnFormat
}

// GetKey returns the value for the field key
func (e *ServiceCatalogResourceType) GetKey() string {
	return e.Key
}

// SetKey sets the value for the field key
func (e *ServiceCatalogResourceType) SetKey(key string) {
	e.Key = key
}

// GetParent returns the value for the field parent
func (e *ServiceCatalogResourceType) GetParent() *string {
	return e.Parent
}

// SetParent sets the value for the field parent
func (e *ServiceCatalogResourceType) SetParent(parent *string) {
	e.Parent = parent
}

// StructPath returns StructPath
func (e *ServiceCatalogResourceType) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathServiceCatalogResourceType.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *ServiceCatalogResourceType) InitializeDefaults() {
}

// serviceCatalogResourceTypeAlias is defined to help pre and post JSON marshaling without recursive loops
type serviceCatalogResourceTypeAlias ServiceCatalogResourceType

// UnmarshalJSON implements json.Unmarshaler
func (e *ServiceCatalogResourceType) UnmarshalJSON(data []byte) error {
	var alias serviceCatalogResourceTypeAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*ServiceCatalogResourceType)(&alias)).InitializeDefaults()
	*e = ServiceCatalogResourceType(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e ServiceCatalogResourceType) MarshalJSON() ([]byte, error) {
	alias := serviceCatalogResourceTypeAlias(e)
	return json.Marshal(alias)
}

// NewServiceCatalogData creates a new ServiceCatalogData
func NewServiceCatalogData() *ServiceCatalogData {
	s := &ServiceCatalogData{}
	s.InitializeDefaults()
	return s
}

// ServiceCatalogData - ServiceCatalogData is a service's full authorization surface for one namespace
// (the RegisterCatalog input).
type ServiceCatalogData struct {
	Actions                       []*ServiceCatalogAction       `json:"actions,omitempty" yaml:"actions,omitempty"`
	Description                   *string                       `json:"description,omitempty" yaml:"description,omitempty"`
	DisplayName                   *string                       `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	Namespace                     string                        `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	ResourceTypes                 []*ServiceCatalogResourceType `json:"resourceTypes,omitempty" yaml:"resourceTypes,omitempty"`
	ServiceLinkedRolePolicyHuJSON []byte                        `json:"serviceLinkedRolePolicyHuJSON,omitempty" yaml:"serviceLinkedRolePolicyHuJSON,omitempty"`
}

// GetActions returns the value for the field actions
func (e *ServiceCatalogData) GetActions() []*ServiceCatalogAction {
	return e.Actions
}

// SetActions sets the value for the field actions
func (e *ServiceCatalogData) SetActions(actions []*ServiceCatalogAction) {
	e.Actions = actions
}

// GetDescription returns the value for the field description
func (e *ServiceCatalogData) GetDescription() *string {
	return e.Description
}

// SetDescription sets the value for the field description
func (e *ServiceCatalogData) SetDescription(description *string) {
	e.Description = description
}

// GetDisplayName returns the value for the field displayName
func (e *ServiceCatalogData) GetDisplayName() *string {
	return e.DisplayName
}

// SetDisplayName sets the value for the field displayName
func (e *ServiceCatalogData) SetDisplayName(displayName *string) {
	e.DisplayName = displayName
}

// GetNamespace returns the value for the field namespace
func (e *ServiceCatalogData) GetNamespace() string {
	return e.Namespace
}

// SetNamespace sets the value for the field namespace
func (e *ServiceCatalogData) SetNamespace(namespace string) {
	e.Namespace = namespace
}

// GetResourceTypes returns the value for the field resourceTypes
func (e *ServiceCatalogData) GetResourceTypes() []*ServiceCatalogResourceType {
	return e.ResourceTypes
}

// SetResourceTypes sets the value for the field resourceTypes
func (e *ServiceCatalogData) SetResourceTypes(resourceTypes []*ServiceCatalogResourceType) {
	e.ResourceTypes = resourceTypes
}

// GetServiceLinkedRolePolicyHuJSON returns the value for the field serviceLinkedRolePolicyHuJSON
func (e *ServiceCatalogData) GetServiceLinkedRolePolicyHuJSON() []byte {
	return e.ServiceLinkedRolePolicyHuJSON
}

// SetServiceLinkedRolePolicyHuJSON sets the value for the field serviceLinkedRolePolicyHuJSON
func (e *ServiceCatalogData) SetServiceLinkedRolePolicyHuJSON(serviceLinkedRolePolicyHuJSON []byte) {
	e.ServiceLinkedRolePolicyHuJSON = serviceLinkedRolePolicyHuJSON
}

// StructPath returns StructPath
func (e *ServiceCatalogData) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathServiceCatalogData.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *ServiceCatalogData) InitializeDefaults() {
}

// serviceCatalogDataAlias is defined to help pre and post JSON marshaling without recursive loops
type serviceCatalogDataAlias ServiceCatalogData

// UnmarshalJSON implements json.Unmarshaler
func (e *ServiceCatalogData) UnmarshalJSON(data []byte) error {
	var alias serviceCatalogDataAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*ServiceCatalogData)(&alias)).InitializeDefaults()
	*e = ServiceCatalogData(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e ServiceCatalogData) MarshalJSON() ([]byte, error) {
	alias := serviceCatalogDataAlias(e)
	return json.Marshal(alias)
}

// NewServiceCatalogInfo creates a new ServiceCatalogInfo
func NewServiceCatalogInfo() *ServiceCatalogInfo {
	s := &ServiceCatalogInfo{}
	s.InitializeDefaults()
	return s
}

// ServiceCatalogInfo - ServiceCatalogInfo is a published catalog together with its publish metadata.
type ServiceCatalogInfo struct {
	Actions                       []*ServiceCatalogAction       `json:"actions,omitempty" yaml:"actions,omitempty"`
	Description                   *string                       `json:"description,omitempty" yaml:"description,omitempty"`
	DisplayName                   *string                       `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	Namespace                     string                        `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	ResourceTypes                 []*ServiceCatalogResourceType `json:"resourceTypes,omitempty" yaml:"resourceTypes,omitempty"`
	Revision                      string                        `json:"revision,omitempty" yaml:"revision,omitempty"`
	ServiceLinkedRolePolicyHuJSON []byte                        `json:"serviceLinkedRolePolicyHuJSON,omitempty" yaml:"serviceLinkedRolePolicyHuJSON,omitempty"`
	// when the catalog revision was published.
	// This value is always set. An absent value indicates a server fault, not a state.
	UpdatedAt *time.Time `json:"updatedAt,omitempty" yaml:"updatedAt,omitempty"`
}

// GetActions returns the value for the field actions
func (e *ServiceCatalogInfo) GetActions() []*ServiceCatalogAction {
	return e.Actions
}

// SetActions sets the value for the field actions
func (e *ServiceCatalogInfo) SetActions(actions []*ServiceCatalogAction) {
	e.Actions = actions
}

// GetDescription returns the value for the field description
func (e *ServiceCatalogInfo) GetDescription() *string {
	return e.Description
}

// SetDescription sets the value for the field description
func (e *ServiceCatalogInfo) SetDescription(description *string) {
	e.Description = description
}

// GetDisplayName returns the value for the field displayName
func (e *ServiceCatalogInfo) GetDisplayName() *string {
	return e.DisplayName
}

// SetDisplayName sets the value for the field displayName
func (e *ServiceCatalogInfo) SetDisplayName(displayName *string) {
	e.DisplayName = displayName
}

// GetNamespace returns the value for the field namespace
func (e *ServiceCatalogInfo) GetNamespace() string {
	return e.Namespace
}

// SetNamespace sets the value for the field namespace
func (e *ServiceCatalogInfo) SetNamespace(namespace string) {
	e.Namespace = namespace
}

// GetResourceTypes returns the value for the field resourceTypes
func (e *ServiceCatalogInfo) GetResourceTypes() []*ServiceCatalogResourceType {
	return e.ResourceTypes
}

// SetResourceTypes sets the value for the field resourceTypes
func (e *ServiceCatalogInfo) SetResourceTypes(resourceTypes []*ServiceCatalogResourceType) {
	e.ResourceTypes = resourceTypes
}

// GetRevision returns the value for the field revision
func (e *ServiceCatalogInfo) GetRevision() string {
	return e.Revision
}

// SetRevision sets the value for the field revision
func (e *ServiceCatalogInfo) SetRevision(revision string) {
	e.Revision = revision
}

// GetServiceLinkedRolePolicyHuJSON returns the value for the field serviceLinkedRolePolicyHuJSON
func (e *ServiceCatalogInfo) GetServiceLinkedRolePolicyHuJSON() []byte {
	return e.ServiceLinkedRolePolicyHuJSON
}

// SetServiceLinkedRolePolicyHuJSON sets the value for the field serviceLinkedRolePolicyHuJSON
func (e *ServiceCatalogInfo) SetServiceLinkedRolePolicyHuJSON(serviceLinkedRolePolicyHuJSON []byte) {
	e.ServiceLinkedRolePolicyHuJSON = serviceLinkedRolePolicyHuJSON
}

// GetUpdatedAt returns the value for the field updatedAt
func (e *ServiceCatalogInfo) GetUpdatedAt() *time.Time {
	return e.UpdatedAt
}

// SetUpdatedAt sets the value for the field updatedAt
func (e *ServiceCatalogInfo) SetUpdatedAt(updatedAt *time.Time) {
	e.UpdatedAt = updatedAt
}

// StructPath returns StructPath
func (e *ServiceCatalogInfo) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathServiceCatalogInfo.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *ServiceCatalogInfo) InitializeDefaults() {
}

// serviceCatalogInfoAlias is defined to help pre and post JSON marshaling without recursive loops
type serviceCatalogInfoAlias ServiceCatalogInfo

// UnmarshalJSON implements json.Unmarshaler
func (e *ServiceCatalogInfo) UnmarshalJSON(data []byte) error {
	var alias serviceCatalogInfoAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*ServiceCatalogInfo)(&alias)).InitializeDefaults()
	*e = ServiceCatalogInfo(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e ServiceCatalogInfo) MarshalJSON() ([]byte, error) {
	alias := serviceCatalogInfoAlias(e)
	return json.Marshal(alias)
}

// NewServiceCatalogSummary creates a new ServiceCatalogSummary
func NewServiceCatalogSummary() *ServiceCatalogSummary {
	s := &ServiceCatalogSummary{}
	s.InitializeDefaults()
	return s
}

// ServiceCatalogSummary - ServiceCatalogSummary is a compact listing entry for the namespace picker.
type ServiceCatalogSummary struct {
	ActionCount       int32   `json:"actionCount,omitempty" yaml:"actionCount,omitempty"`
	Description       *string `json:"description,omitempty" yaml:"description,omitempty"`
	DisplayName       *string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	Namespace         string  `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	ResourceTypeCount int32   `json:"resourceTypeCount,omitempty" yaml:"resourceTypeCount,omitempty"`
	Revision          string  `json:"revision,omitempty" yaml:"revision,omitempty"`
	// when the catalog revision was published.
	// This value is always set. An absent value indicates a server fault, not a state.
	UpdatedAt *time.Time `json:"updatedAt,omitempty" yaml:"updatedAt,omitempty"`
}

// GetActionCount returns the value for the field actionCount
func (e *ServiceCatalogSummary) GetActionCount() int32 {
	return e.ActionCount
}

// SetActionCount sets the value for the field actionCount
func (e *ServiceCatalogSummary) SetActionCount(actionCount int32) {
	e.ActionCount = actionCount
}

// GetDescription returns the value for the field description
func (e *ServiceCatalogSummary) GetDescription() *string {
	return e.Description
}

// SetDescription sets the value for the field description
func (e *ServiceCatalogSummary) SetDescription(description *string) {
	e.Description = description
}

// GetDisplayName returns the value for the field displayName
func (e *ServiceCatalogSummary) GetDisplayName() *string {
	return e.DisplayName
}

// SetDisplayName sets the value for the field displayName
func (e *ServiceCatalogSummary) SetDisplayName(displayName *string) {
	e.DisplayName = displayName
}

// GetNamespace returns the value for the field namespace
func (e *ServiceCatalogSummary) GetNamespace() string {
	return e.Namespace
}

// SetNamespace sets the value for the field namespace
func (e *ServiceCatalogSummary) SetNamespace(namespace string) {
	e.Namespace = namespace
}

// GetResourceTypeCount returns the value for the field resourceTypeCount
func (e *ServiceCatalogSummary) GetResourceTypeCount() int32 {
	return e.ResourceTypeCount
}

// SetResourceTypeCount sets the value for the field resourceTypeCount
func (e *ServiceCatalogSummary) SetResourceTypeCount(resourceTypeCount int32) {
	e.ResourceTypeCount = resourceTypeCount
}

// GetRevision returns the value for the field revision
func (e *ServiceCatalogSummary) GetRevision() string {
	return e.Revision
}

// SetRevision sets the value for the field revision
func (e *ServiceCatalogSummary) SetRevision(revision string) {
	e.Revision = revision
}

// GetUpdatedAt returns the value for the field updatedAt
func (e *ServiceCatalogSummary) GetUpdatedAt() *time.Time {
	return e.UpdatedAt
}

// SetUpdatedAt sets the value for the field updatedAt
func (e *ServiceCatalogSummary) SetUpdatedAt(updatedAt *time.Time) {
	e.UpdatedAt = updatedAt
}

// StructPath returns StructPath
func (e *ServiceCatalogSummary) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathServiceCatalogSummary.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *ServiceCatalogSummary) InitializeDefaults() {
}

// serviceCatalogSummaryAlias is defined to help pre and post JSON marshaling without recursive loops
type serviceCatalogSummaryAlias ServiceCatalogSummary

// UnmarshalJSON implements json.Unmarshaler
func (e *ServiceCatalogSummary) UnmarshalJSON(data []byte) error {
	var alias serviceCatalogSummaryAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*ServiceCatalogSummary)(&alias)).InitializeDefaults()
	*e = ServiceCatalogSummary(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e ServiceCatalogSummary) MarshalJSON() ([]byte, error) {
	alias := serviceCatalogSummaryAlias(e)
	return json.Marshal(alias)
}

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

// NewManagedByService creates a new ManagedByService
func NewManagedByService() *ManagedByService {
	s := &ManagedByService{}
	s.InitializeDefaults()
	return s
}

// ManagedByService - Marks a resource as owned by a service. Present only when managed; its presence is
// the read-only "managed" signal for the console banner.
type ManagedByService struct {
	// the managing service principal FQDN, e.g. "uplink.deployport.io"
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}

// GetService returns the value for the field service
func (e *ManagedByService) GetService() string {
	return e.Service
}

// SetService sets the value for the field service
func (e *ManagedByService) SetService(service string) {
	e.Service = service
}

// StructPath returns StructPath
func (e *ManagedByService) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathManagedByService.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *ManagedByService) InitializeDefaults() {
}

// managedByServiceAlias is defined to help pre and post JSON marshaling without recursive loops
type managedByServiceAlias ManagedByService

// UnmarshalJSON implements json.Unmarshaler
func (e *ManagedByService) UnmarshalJSON(data []byte) error {
	var alias managedByServiceAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*ManagedByService)(&alias)).InitializeDefaults()
	*e = ManagedByService(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e ManagedByService) MarshalJSON() ([]byte, error) {
	alias := managedByServiceAlias(e)
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
	// false when the user is suspended: it cannot authenticate and its existing
	// credentials are inert until it is re-enabled. Reversible, unlike Destroy.
	Active bool `json:"active,omitempty" yaml:"active,omitempty"`
	// when the user was created
	// This value is always set. An absent value indicates a server fault, not a state.
	CreatedAt   *time.Time `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
	Description string     `json:"description,omitempty" yaml:"description,omitempty"`
	// DRN of this user, e.g. iam:User(johan)
	Drn string `json:"drn,omitempty" yaml:"drn,omitempty"`
	// present only for a service-managed user: the account may view it but not
	// mutate it or its access keys
	ManagedBy *ManagedByService `json:"managedBy,omitempty" yaml:"managedBy,omitempty"`
	// when the user comes from SSO, this field is populated with the extra information
	Sso      *UserInformationSSO `json:"sso,omitempty" yaml:"sso,omitempty"`
	Username string              `json:"username,omitempty" yaml:"username,omitempty"`
}

// GetActive returns the value for the field active
func (e *UserInformation) GetActive() bool {
	return e.Active
}

// SetActive sets the value for the field active
func (e *UserInformation) SetActive(active bool) {
	e.Active = active
}

// GetCreatedAt returns the value for the field createdAt
func (e *UserInformation) GetCreatedAt() *time.Time {
	return e.CreatedAt
}

// SetCreatedAt sets the value for the field createdAt
func (e *UserInformation) SetCreatedAt(createdAt *time.Time) {
	e.CreatedAt = createdAt
}

// GetDescription returns the value for the field description
func (e *UserInformation) GetDescription() string {
	return e.Description
}

// SetDescription sets the value for the field description
func (e *UserInformation) SetDescription(description string) {
	e.Description = description
}

// GetDrn returns the value for the field drn
func (e *UserInformation) GetDrn() string {
	return e.Drn
}

// SetDrn sets the value for the field drn
func (e *UserInformation) SetDrn(drn string) {
	e.Drn = drn
}

// GetManagedBy returns the value for the field managedBy
func (e *UserInformation) GetManagedBy() *ManagedByService {
	return e.ManagedBy
}

// SetManagedBy sets the value for the field managedBy
func (e *UserInformation) SetManagedBy(managedBy *ManagedByService) {
	e.ManagedBy = managedBy
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
	// when the role was created
	// This value is always set. An absent value indicates a server fault, not a state.
	CreatedAt   *time.Time `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
	Description string     `json:"description,omitempty" yaml:"description,omitempty"`
	// DRN of this role, e.g. iam:Role(deployer)
	Drn string `json:"drn,omitempty" yaml:"drn,omitempty"`
	// present only for a service-linked role: the account may view it but not
	// mutate it, its attachments, or its access keys
	ManagedBy *ManagedByService `json:"managedBy,omitempty" yaml:"managedBy,omitempty"`
	Name      string            `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetCreatedAt returns the value for the field createdAt
func (e *RoleInformation) GetCreatedAt() *time.Time {
	return e.CreatedAt
}

// SetCreatedAt sets the value for the field createdAt
func (e *RoleInformation) SetCreatedAt(createdAt *time.Time) {
	e.CreatedAt = createdAt
}

// GetDescription returns the value for the field description
func (e *RoleInformation) GetDescription() string {
	return e.Description
}

// SetDescription sets the value for the field description
func (e *RoleInformation) SetDescription(description string) {
	e.Description = description
}

// GetDrn returns the value for the field drn
func (e *RoleInformation) GetDrn() string {
	return e.Drn
}

// SetDrn sets the value for the field drn
func (e *RoleInformation) SetDrn(drn string) {
	e.Drn = drn
}

// GetManagedBy returns the value for the field managedBy
func (e *RoleInformation) GetManagedBy() *ManagedByService {
	return e.ManagedBy
}

// SetManagedBy sets the value for the field managedBy
func (e *RoleInformation) SetManagedBy(managedBy *ManagedByService) {
	e.ManagedBy = managedBy
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

// NewGroupInformation creates a new GroupInformation
func NewGroupInformation() *GroupInformation {
	s := &GroupInformation{}
	s.InitializeDefaults()
	return s
}

// GroupInformation - A group: an account-scoped, named collection of users. Identity policies attached
// to a group are granted to every member (unioned into the member's permissions).
type GroupInformation struct {
	// when the group was created
	// This value is always set. An absent value indicates a server fault, not a state.
	CreatedAt   *time.Time `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
	Description string     `json:"description,omitempty" yaml:"description,omitempty"`
	// DRN of this group, e.g. iam:Group(engineering)
	Drn string `json:"drn,omitempty" yaml:"drn,omitempty"`
	// present only for a service-managed group: the account may view it but not
	// mutate it (members, policy attachments, destroy)
	ManagedBy *ManagedByService `json:"managedBy,omitempty" yaml:"managedBy,omitempty"`
	Name      string            `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetCreatedAt returns the value for the field createdAt
func (e *GroupInformation) GetCreatedAt() *time.Time {
	return e.CreatedAt
}

// SetCreatedAt sets the value for the field createdAt
func (e *GroupInformation) SetCreatedAt(createdAt *time.Time) {
	e.CreatedAt = createdAt
}

// GetDescription returns the value for the field description
func (e *GroupInformation) GetDescription() string {
	return e.Description
}

// SetDescription sets the value for the field description
func (e *GroupInformation) SetDescription(description string) {
	e.Description = description
}

// GetDrn returns the value for the field drn
func (e *GroupInformation) GetDrn() string {
	return e.Drn
}

// SetDrn sets the value for the field drn
func (e *GroupInformation) SetDrn(drn string) {
	e.Drn = drn
}

// GetManagedBy returns the value for the field managedBy
func (e *GroupInformation) GetManagedBy() *ManagedByService {
	return e.ManagedBy
}

// SetManagedBy sets the value for the field managedBy
func (e *GroupInformation) SetManagedBy(managedBy *ManagedByService) {
	e.ManagedBy = managedBy
}

// GetName returns the value for the field name
func (e *GroupInformation) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *GroupInformation) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *GroupInformation) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathGroupInformation.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *GroupInformation) InitializeDefaults() {
}

// groupInformationAlias is defined to help pre and post JSON marshaling without recursive loops
type groupInformationAlias GroupInformation

// UnmarshalJSON implements json.Unmarshaler
func (e *GroupInformation) UnmarshalJSON(data []byte) error {
	var alias groupInformationAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*GroupInformation)(&alias)).InitializeDefaults()
	*e = GroupInformation(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e GroupInformation) MarshalJSON() ([]byte, error) {
	alias := groupInformationAlias(e)
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
	AccessKeyID string `json:"accessKeyID,omitempty" yaml:"accessKeyID,omitempty"`
	// when the credential expires. null for permanent access keys
	ExpiresAt       *time.Time `json:"expiresAt,omitempty" yaml:"expiresAt,omitempty"`
	SecretAccessKey string     `json:"secretAccessKey,omitempty" yaml:"secretAccessKey,omitempty"`
}

// GetAccessKeyID returns the value for the field accessKeyID
func (e *Credentials) GetAccessKeyID() string {
	return e.AccessKeyID
}

// SetAccessKeyID sets the value for the field accessKeyID
func (e *Credentials) SetAccessKeyID(accessKeyID string) {
	e.AccessKeyID = accessKeyID
}

// GetExpiresAt returns the value for the field expiresAt
func (e *Credentials) GetExpiresAt() *time.Time {
	return e.ExpiresAt
}

// SetExpiresAt sets the value for the field expiresAt
func (e *Credentials) SetExpiresAt(expiresAt *time.Time) {
	e.ExpiresAt = expiresAt
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

// NewSSOProviderUnavailableError creates a new SSOProviderUnavailableError
func NewSSOProviderUnavailableError() *SSOProviderUnavailableError {
	s := &SSOProviderUnavailableError{}
	s.InitializeDefaults()
	return s
}

// SSOProviderUnavailableError struct
type SSOProviderUnavailableError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *SSOProviderUnavailableError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [SSOProviderUnavailableError]
func (e *SSOProviderUnavailableError) Is(err error) bool {
	_, ok := err.(*SSOProviderUnavailableError)
	return ok
}

// IsSSOProviderUnavailableError indicates whether the given error chain contains an error of type [SSOProviderUnavailableError]
func IsSSOProviderUnavailableError(err error) bool {
	return errors.Is(err, &SSOProviderUnavailableError{})
}

// GetMessage returns the value for the field message
func (e *SSOProviderUnavailableError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *SSOProviderUnavailableError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *SSOProviderUnavailableError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathSSOProviderUnavailableError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *SSOProviderUnavailableError) InitializeDefaults() {
}

// sSOProviderUnavailableErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type sSOProviderUnavailableErrorAlias SSOProviderUnavailableError

// UnmarshalJSON implements json.Unmarshaler
func (e *SSOProviderUnavailableError) UnmarshalJSON(data []byte) error {
	var alias sSOProviderUnavailableErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*SSOProviderUnavailableError)(&alias)).InitializeDefaults()
	*e = SSOProviderUnavailableError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e SSOProviderUnavailableError) MarshalJSON() ([]byte, error) {
	alias := sSOProviderUnavailableErrorAlias(e)
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
	BrowseURL                 string `json:"browseURL,omitempty" yaml:"browseURL,omitempty"`
	CompletionIntervalSeconds int32  `json:"completionIntervalSeconds,omitempty" yaml:"completionIntervalSeconds,omitempty"`
	// when the sign-in flow expires.
	// This value is always set. An absent value indicates a server fault, not a state.
	ExpiresAt *time.Time `json:"expiresAt,omitempty" yaml:"expiresAt,omitempty"`
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
func (e *SSOFlow) GetExpiresAt() *time.Time {
	return e.ExpiresAt
}

// SetExpiresAt sets the value for the field expiresAt
func (e *SSOFlow) SetExpiresAt(expiresAt *time.Time) {
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

// NewRegionEndpoint creates a new RegionEndpoint
func NewRegionEndpoint() *RegionEndpoint {
	s := &RegionEndpoint{}
	s.InitializeDefaults()
	return s
}

// RegionEndpoint - A single named service endpoint for a region, e.g. {name: "api", url: "https://iam.us-east-2.api.deployport.io"}.
type RegionEndpoint struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	Url  string `json:"url,omitempty" yaml:"url,omitempty"`
}

// GetName returns the value for the field name
func (e *RegionEndpoint) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *RegionEndpoint) SetName(name string) {
	e.Name = name
}

// GetUrl returns the value for the field url
func (e *RegionEndpoint) GetUrl() string {
	return e.Url
}

// SetUrl sets the value for the field url
func (e *RegionEndpoint) SetUrl(url string) {
	e.Url = url
}

// StructPath returns StructPath
func (e *RegionEndpoint) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRegionEndpoint.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RegionEndpoint) InitializeDefaults() {
}

// regionEndpointAlias is defined to help pre and post JSON marshaling without recursive loops
type regionEndpointAlias RegionEndpoint

// UnmarshalJSON implements json.Unmarshaler
func (e *RegionEndpoint) UnmarshalJSON(data []byte) error {
	var alias regionEndpointAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RegionEndpoint)(&alias)).InitializeDefaults()
	*e = RegionEndpoint(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RegionEndpoint) MarshalJSON() ([]byte, error) {
	alias := regionEndpointAlias(e)
	return json.Marshal(alias)
}

// NewRegionInfo creates a new RegionInfo
func NewRegionInfo() *RegionInfo {
	s := &RegionInfo{}
	s.InitializeDefaults()
	return s
}

// RegionInfo - Metadata about an available region and its service endpoints.
type RegionInfo struct {
	// Global console origin for this region, e.g. "https://global.us-east-2.console.deployport.io".
	ConsoleURL string            `json:"consoleURL,omitempty" yaml:"consoleURL,omitempty"`
	Endpoints  []*RegionEndpoint `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
	Slug       string            `json:"slug,omitempty" yaml:"slug,omitempty"`
}

// GetConsoleURL returns the value for the field consoleURL
func (e *RegionInfo) GetConsoleURL() string {
	return e.ConsoleURL
}

// SetConsoleURL sets the value for the field consoleURL
func (e *RegionInfo) SetConsoleURL(consoleURL string) {
	e.ConsoleURL = consoleURL
}

// GetEndpoints returns the value for the field endpoints
func (e *RegionInfo) GetEndpoints() []*RegionEndpoint {
	return e.Endpoints
}

// SetEndpoints sets the value for the field endpoints
func (e *RegionInfo) SetEndpoints(endpoints []*RegionEndpoint) {
	e.Endpoints = endpoints
}

// GetSlug returns the value for the field slug
func (e *RegionInfo) GetSlug() string {
	return e.Slug
}

// SetSlug sets the value for the field slug
func (e *RegionInfo) SetSlug(slug string) {
	e.Slug = slug
}

// StructPath returns StructPath
func (e *RegionInfo) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRegionInfo.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RegionInfo) InitializeDefaults() {
}

// regionInfoAlias is defined to help pre and post JSON marshaling without recursive loops
type regionInfoAlias RegionInfo

// UnmarshalJSON implements json.Unmarshaler
func (e *RegionInfo) UnmarshalJSON(data []byte) error {
	var alias regionInfoAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RegionInfo)(&alias)).InitializeDefaults()
	*e = RegionInfo(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RegionInfo) MarshalJSON() ([]byte, error) {
	alias := regionInfoAlias(e)
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

// NewPolicyNotFoundError creates a new PolicyNotFoundError
func NewPolicyNotFoundError() *PolicyNotFoundError {
	s := &PolicyNotFoundError{}
	s.InitializeDefaults()
	return s
}

// PolicyNotFoundError struct
type PolicyNotFoundError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *PolicyNotFoundError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [PolicyNotFoundError]
func (e *PolicyNotFoundError) Is(err error) bool {
	_, ok := err.(*PolicyNotFoundError)
	return ok
}

// IsPolicyNotFoundError indicates whether the given error chain contains an error of type [PolicyNotFoundError]
func IsPolicyNotFoundError(err error) bool {
	return errors.Is(err, &PolicyNotFoundError{})
}

// GetMessage returns the value for the field message
func (e *PolicyNotFoundError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *PolicyNotFoundError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *PolicyNotFoundError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathPolicyNotFoundError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *PolicyNotFoundError) InitializeDefaults() {
}

// policyNotFoundErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type policyNotFoundErrorAlias PolicyNotFoundError

// UnmarshalJSON implements json.Unmarshaler
func (e *PolicyNotFoundError) UnmarshalJSON(data []byte) error {
	var alias policyNotFoundErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*PolicyNotFoundError)(&alias)).InitializeDefaults()
	*e = PolicyNotFoundError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e PolicyNotFoundError) MarshalJSON() ([]byte, error) {
	alias := policyNotFoundErrorAlias(e)
	return json.Marshal(alias)
}

// NewInvitation creates a new Invitation
func NewInvitation() *Invitation {
	s := &Invitation{}
	s.InitializeDefaults()
	return s
}

// Invitation - An invitation for a person (by email) to become a member of an account.
type Invitation struct {
	// when the invitation was created.
	// This value is always set. An absent value indicates a server fault, not a state.
	CreatedAt *time.Time `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
	// invited email address
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
	// when the invitation expires
	// This value is always set. An absent value indicates a server fault, not a state.
	ExpiresAt *time.Time `json:"expiresAt,omitempty" yaml:"expiresAt,omitempty"`
	// groups (by name) the member joins when they accept
	GroupNames []string `json:"groupNames,omitempty" yaml:"groupNames,omitempty"`
	// unique identifier of the invitation (opaque)
	Id string `json:"id,omitempty" yaml:"id,omitempty"`
	// identity policies (by name) attached to the member when they accept
	PolicyNames []string `json:"policyNames,omitempty" yaml:"policyNames,omitempty"`
	// optional SSO provider hint (e.g. google/github) the invitee should use
	ProviderName *string `json:"providerName,omitempty" yaml:"providerName,omitempty"`
	// lifecycle status: pending | accepted | revoked | declined
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}

// GetCreatedAt returns the value for the field createdAt
func (e *Invitation) GetCreatedAt() *time.Time {
	return e.CreatedAt
}

// SetCreatedAt sets the value for the field createdAt
func (e *Invitation) SetCreatedAt(createdAt *time.Time) {
	e.CreatedAt = createdAt
}

// GetEmail returns the value for the field email
func (e *Invitation) GetEmail() string {
	return e.Email
}

// SetEmail sets the value for the field email
func (e *Invitation) SetEmail(email string) {
	e.Email = email
}

// GetExpiresAt returns the value for the field expiresAt
func (e *Invitation) GetExpiresAt() *time.Time {
	return e.ExpiresAt
}

// SetExpiresAt sets the value for the field expiresAt
func (e *Invitation) SetExpiresAt(expiresAt *time.Time) {
	e.ExpiresAt = expiresAt
}

// GetGroupNames returns the value for the field groupNames
func (e *Invitation) GetGroupNames() []string {
	return e.GroupNames
}

// SetGroupNames sets the value for the field groupNames
func (e *Invitation) SetGroupNames(groupNames []string) {
	e.GroupNames = groupNames
}

// GetId returns the value for the field id
func (e *Invitation) GetId() string {
	return e.Id
}

// SetId sets the value for the field id
func (e *Invitation) SetId(id string) {
	e.Id = id
}

// GetPolicyNames returns the value for the field policyNames
func (e *Invitation) GetPolicyNames() []string {
	return e.PolicyNames
}

// SetPolicyNames sets the value for the field policyNames
func (e *Invitation) SetPolicyNames(policyNames []string) {
	e.PolicyNames = policyNames
}

// GetProviderName returns the value for the field providerName
func (e *Invitation) GetProviderName() *string {
	return e.ProviderName
}

// SetProviderName sets the value for the field providerName
func (e *Invitation) SetProviderName(providerName *string) {
	e.ProviderName = providerName
}

// GetStatus returns the value for the field status
func (e *Invitation) GetStatus() string {
	return e.Status
}

// SetStatus sets the value for the field status
func (e *Invitation) SetStatus(status string) {
	e.Status = status
}

// StructPath returns StructPath
func (e *Invitation) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvitation.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *Invitation) InitializeDefaults() {
}

// invitationAlias is defined to help pre and post JSON marshaling without recursive loops
type invitationAlias Invitation

// UnmarshalJSON implements json.Unmarshaler
func (e *Invitation) UnmarshalJSON(data []byte) error {
	var alias invitationAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*Invitation)(&alias)).InitializeDefaults()
	*e = Invitation(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e Invitation) MarshalJSON() ([]byte, error) {
	alias := invitationAlias(e)
	return json.Marshal(alias)
}

// NewInvalidInvitationError creates a new InvalidInvitationError
func NewInvalidInvitationError() *InvalidInvitationError {
	s := &InvalidInvitationError{}
	s.InitializeDefaults()
	return s
}

// InvalidInvitationError - The invitation token/id is unknown, already used, revoked, or expired.
type InvalidInvitationError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InvalidInvitationError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InvalidInvitationError]
func (e *InvalidInvitationError) Is(err error) bool {
	_, ok := err.(*InvalidInvitationError)
	return ok
}

// IsInvalidInvitationError indicates whether the given error chain contains an error of type [InvalidInvitationError]
func IsInvalidInvitationError(err error) bool {
	return errors.Is(err, &InvalidInvitationError{})
}

// GetMessage returns the value for the field message
func (e *InvalidInvitationError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InvalidInvitationError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InvalidInvitationError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvalidInvitationError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvalidInvitationError) InitializeDefaults() {
}

// invalidInvitationErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type invalidInvitationErrorAlias InvalidInvitationError

// UnmarshalJSON implements json.Unmarshaler
func (e *InvalidInvitationError) UnmarshalJSON(data []byte) error {
	var alias invalidInvitationErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvalidInvitationError)(&alias)).InitializeDefaults()
	*e = InvalidInvitationError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvalidInvitationError) MarshalJSON() ([]byte, error) {
	alias := invalidInvitationErrorAlias(e)
	return json.Marshal(alias)
}

// NewInvitationNotFoundError creates a new InvitationNotFoundError
func NewInvitationNotFoundError() *InvitationNotFoundError {
	s := &InvitationNotFoundError{}
	s.InitializeDefaults()
	return s
}

// InvitationNotFoundError - No pending invitation with the given id exists in the caller's account.
type InvitationNotFoundError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InvitationNotFoundError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InvitationNotFoundError]
func (e *InvitationNotFoundError) Is(err error) bool {
	_, ok := err.(*InvitationNotFoundError)
	return ok
}

// IsInvitationNotFoundError indicates whether the given error chain contains an error of type [InvitationNotFoundError]
func IsInvitationNotFoundError(err error) bool {
	return errors.Is(err, &InvitationNotFoundError{})
}

// GetMessage returns the value for the field message
func (e *InvitationNotFoundError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InvitationNotFoundError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InvitationNotFoundError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvitationNotFoundError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvitationNotFoundError) InitializeDefaults() {
}

// invitationNotFoundErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type invitationNotFoundErrorAlias InvitationNotFoundError

// UnmarshalJSON implements json.Unmarshaler
func (e *InvitationNotFoundError) UnmarshalJSON(data []byte) error {
	var alias invitationNotFoundErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvitationNotFoundError)(&alias)).InitializeDefaults()
	*e = InvitationNotFoundError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvitationNotFoundError) MarshalJSON() ([]byte, error) {
	alias := invitationNotFoundErrorAlias(e)
	return json.Marshal(alias)
}

// NewInvitationPreview creates a new InvitationPreview
func NewInvitationPreview() *InvitationPreview {
	s := &InvitationPreview{}
	s.InitializeDefaults()
	return s
}

// InvitationPreview - A safe, invitee-facing view of an invitation resolved from its token, so the
// frontend can show what the invitation is before the invitee signs in/accepts.
type InvitationPreview struct {
	// the account the invitee is invited to join
	AccountName string `json:"accountName,omitempty" yaml:"accountName,omitempty"`
	// the invited email address
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
	// when the invitation expires
	// This value is always set. An absent value indicates a server fault, not a state.
	ExpiresAt *time.Time `json:"expiresAt,omitempty" yaml:"expiresAt,omitempty"`
	// the inviter's username (best-effort; may be empty)
	InvitedBy string `json:"invitedBy,omitempty" yaml:"invitedBy,omitempty"`
	// lifecycle status: pending | accepted | revoked | declined | expired
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}

// GetAccountName returns the value for the field accountName
func (e *InvitationPreview) GetAccountName() string {
	return e.AccountName
}

// SetAccountName sets the value for the field accountName
func (e *InvitationPreview) SetAccountName(accountName string) {
	e.AccountName = accountName
}

// GetEmail returns the value for the field email
func (e *InvitationPreview) GetEmail() string {
	return e.Email
}

// SetEmail sets the value for the field email
func (e *InvitationPreview) SetEmail(email string) {
	e.Email = email
}

// GetExpiresAt returns the value for the field expiresAt
func (e *InvitationPreview) GetExpiresAt() *time.Time {
	return e.ExpiresAt
}

// SetExpiresAt sets the value for the field expiresAt
func (e *InvitationPreview) SetExpiresAt(expiresAt *time.Time) {
	e.ExpiresAt = expiresAt
}

// GetInvitedBy returns the value for the field invitedBy
func (e *InvitationPreview) GetInvitedBy() string {
	return e.InvitedBy
}

// SetInvitedBy sets the value for the field invitedBy
func (e *InvitationPreview) SetInvitedBy(invitedBy string) {
	e.InvitedBy = invitedBy
}

// GetStatus returns the value for the field status
func (e *InvitationPreview) GetStatus() string {
	return e.Status
}

// SetStatus sets the value for the field status
func (e *InvitationPreview) SetStatus(status string) {
	e.Status = status
}

// StructPath returns StructPath
func (e *InvitationPreview) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvitationPreview.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvitationPreview) InitializeDefaults() {
}

// invitationPreviewAlias is defined to help pre and post JSON marshaling without recursive loops
type invitationPreviewAlias InvitationPreview

// UnmarshalJSON implements json.Unmarshaler
func (e *InvitationPreview) UnmarshalJSON(data []byte) error {
	var alias invitationPreviewAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvitationPreview)(&alias)).InitializeDefaults()
	*e = InvitationPreview(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvitationPreview) MarshalJSON() ([]byte, error) {
	alias := invitationPreviewAlias(e)
	return json.Marshal(alias)
}

// NewOIDCProvider creates a new OIDCProvider
func NewOIDCProvider() *OIDCProvider {
	s := &OIDCProvider{}
	s.InitializeDefaults()
	return s
}

// OIDCProvider - --- OIDC federation (AssumeRoleWithWebIdentity) ---
// A registered external OIDC identity provider used for web-identity role federation.
type OIDCProvider struct {
	Audiences []string `json:"audiences,omitempty" yaml:"audiences,omitempty"`
	// DRN of this provider, e.g. iam:OIDCProvider(github-actions)
	Drn       string `json:"drn,omitempty" yaml:"drn,omitempty"`
	IssuerURL string `json:"issuerURL,omitempty" yaml:"issuerURL,omitempty"`
	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetAudiences returns the value for the field audiences
func (e *OIDCProvider) GetAudiences() []string {
	return e.Audiences
}

// SetAudiences sets the value for the field audiences
func (e *OIDCProvider) SetAudiences(audiences []string) {
	e.Audiences = audiences
}

// GetDrn returns the value for the field drn
func (e *OIDCProvider) GetDrn() string {
	return e.Drn
}

// SetDrn sets the value for the field drn
func (e *OIDCProvider) SetDrn(drn string) {
	e.Drn = drn
}

// GetIssuerURL returns the value for the field issuerURL
func (e *OIDCProvider) GetIssuerURL() string {
	return e.IssuerURL
}

// SetIssuerURL sets the value for the field issuerURL
func (e *OIDCProvider) SetIssuerURL(issuerURL string) {
	e.IssuerURL = issuerURL
}

// GetName returns the value for the field name
func (e *OIDCProvider) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *OIDCProvider) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *OIDCProvider) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathOIDCProvider.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *OIDCProvider) InitializeDefaults() {
}

// oIDCProviderAlias is defined to help pre and post JSON marshaling without recursive loops
type oIDCProviderAlias OIDCProvider

// UnmarshalJSON implements json.Unmarshaler
func (e *OIDCProvider) UnmarshalJSON(data []byte) error {
	var alias oIDCProviderAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*OIDCProvider)(&alias)).InitializeDefaults()
	*e = OIDCProvider(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e OIDCProvider) MarshalJSON() ([]byte, error) {
	alias := oIDCProviderAlias(e)
	return json.Marshal(alias)
}

// NewInvalidOIDCProviderError creates a new InvalidOIDCProviderError
func NewInvalidOIDCProviderError() *InvalidOIDCProviderError {
	s := &InvalidOIDCProviderError{}
	s.InitializeDefaults()
	return s
}

// InvalidOIDCProviderError struct
type InvalidOIDCProviderError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InvalidOIDCProviderError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InvalidOIDCProviderError]
func (e *InvalidOIDCProviderError) Is(err error) bool {
	_, ok := err.(*InvalidOIDCProviderError)
	return ok
}

// IsInvalidOIDCProviderError indicates whether the given error chain contains an error of type [InvalidOIDCProviderError]
func IsInvalidOIDCProviderError(err error) bool {
	return errors.Is(err, &InvalidOIDCProviderError{})
}

// GetMessage returns the value for the field message
func (e *InvalidOIDCProviderError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InvalidOIDCProviderError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InvalidOIDCProviderError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvalidOIDCProviderError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvalidOIDCProviderError) InitializeDefaults() {
}

// invalidOIDCProviderErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type invalidOIDCProviderErrorAlias InvalidOIDCProviderError

// UnmarshalJSON implements json.Unmarshaler
func (e *InvalidOIDCProviderError) UnmarshalJSON(data []byte) error {
	var alias invalidOIDCProviderErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvalidOIDCProviderError)(&alias)).InitializeDefaults()
	*e = InvalidOIDCProviderError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvalidOIDCProviderError) MarshalJSON() ([]byte, error) {
	alias := invalidOIDCProviderErrorAlias(e)
	return json.Marshal(alias)
}

// NewInvalidOIDCIssuerError creates a new InvalidOIDCIssuerError
func NewInvalidOIDCIssuerError() *InvalidOIDCIssuerError {
	s := &InvalidOIDCIssuerError{}
	s.InitializeDefaults()
	return s
}

// InvalidOIDCIssuerError - The issuer URL is malformed or is not a reachable OIDC provider (discovery failed).
type InvalidOIDCIssuerError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InvalidOIDCIssuerError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InvalidOIDCIssuerError]
func (e *InvalidOIDCIssuerError) Is(err error) bool {
	_, ok := err.(*InvalidOIDCIssuerError)
	return ok
}

// IsInvalidOIDCIssuerError indicates whether the given error chain contains an error of type [InvalidOIDCIssuerError]
func IsInvalidOIDCIssuerError(err error) bool {
	return errors.Is(err, &InvalidOIDCIssuerError{})
}

// GetMessage returns the value for the field message
func (e *InvalidOIDCIssuerError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InvalidOIDCIssuerError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InvalidOIDCIssuerError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvalidOIDCIssuerError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvalidOIDCIssuerError) InitializeDefaults() {
}

// invalidOIDCIssuerErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type invalidOIDCIssuerErrorAlias InvalidOIDCIssuerError

// UnmarshalJSON implements json.Unmarshaler
func (e *InvalidOIDCIssuerError) UnmarshalJSON(data []byte) error {
	var alias invalidOIDCIssuerErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvalidOIDCIssuerError)(&alias)).InitializeDefaults()
	*e = InvalidOIDCIssuerError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvalidOIDCIssuerError) MarshalJSON() ([]byte, error) {
	alias := invalidOIDCIssuerErrorAlias(e)
	return json.Marshal(alias)
}

// NewOIDCProviderNotFoundError creates a new OIDCProviderNotFoundError
func NewOIDCProviderNotFoundError() *OIDCProviderNotFoundError {
	s := &OIDCProviderNotFoundError{}
	s.InitializeDefaults()
	return s
}

// OIDCProviderNotFoundError struct
type OIDCProviderNotFoundError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *OIDCProviderNotFoundError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [OIDCProviderNotFoundError]
func (e *OIDCProviderNotFoundError) Is(err error) bool {
	_, ok := err.(*OIDCProviderNotFoundError)
	return ok
}

// IsOIDCProviderNotFoundError indicates whether the given error chain contains an error of type [OIDCProviderNotFoundError]
func IsOIDCProviderNotFoundError(err error) bool {
	return errors.Is(err, &OIDCProviderNotFoundError{})
}

// GetMessage returns the value for the field message
func (e *OIDCProviderNotFoundError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *OIDCProviderNotFoundError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *OIDCProviderNotFoundError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathOIDCProviderNotFoundError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *OIDCProviderNotFoundError) InitializeDefaults() {
}

// oIDCProviderNotFoundErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type oIDCProviderNotFoundErrorAlias OIDCProviderNotFoundError

// UnmarshalJSON implements json.Unmarshaler
func (e *OIDCProviderNotFoundError) UnmarshalJSON(data []byte) error {
	var alias oIDCProviderNotFoundErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*OIDCProviderNotFoundError)(&alias)).InitializeDefaults()
	*e = OIDCProviderNotFoundError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e OIDCProviderNotFoundError) MarshalJSON() ([]byte, error) {
	alias := oIDCProviderNotFoundErrorAlias(e)
	return json.Marshal(alias)
}

// NewOIDCProviderInUseError creates a new OIDCProviderInUseError
func NewOIDCProviderInUseError() *OIDCProviderInUseError {
	s := &OIDCProviderInUseError{}
	s.InitializeDefaults()
	return s
}

// OIDCProviderInUseError struct
type OIDCProviderInUseError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *OIDCProviderInUseError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [OIDCProviderInUseError]
func (e *OIDCProviderInUseError) Is(err error) bool {
	_, ok := err.(*OIDCProviderInUseError)
	return ok
}

// IsOIDCProviderInUseError indicates whether the given error chain contains an error of type [OIDCProviderInUseError]
func IsOIDCProviderInUseError(err error) bool {
	return errors.Is(err, &OIDCProviderInUseError{})
}

// GetMessage returns the value for the field message
func (e *OIDCProviderInUseError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *OIDCProviderInUseError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *OIDCProviderInUseError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathOIDCProviderInUseError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *OIDCProviderInUseError) InitializeDefaults() {
}

// oIDCProviderInUseErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type oIDCProviderInUseErrorAlias OIDCProviderInUseError

// UnmarshalJSON implements json.Unmarshaler
func (e *OIDCProviderInUseError) UnmarshalJSON(data []byte) error {
	var alias oIDCProviderInUseErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*OIDCProviderInUseError)(&alias)).InitializeDefaults()
	*e = OIDCProviderInUseError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e OIDCProviderInUseError) MarshalJSON() ([]byte, error) {
	alias := oIDCProviderInUseErrorAlias(e)
	return json.Marshal(alias)
}

// NewTrustPolicyStatement creates a new TrustPolicyStatement
func NewTrustPolicyStatement() *TrustPolicyStatement {
	s := &TrustPolicyStatement{}
	s.InitializeDefaults()
	return s
}

// TrustPolicyStatement - A trust-policy statement (Allow or Deny; Deny overrides). Each constraints entry is a clause of space-separated qualifier atoms such as github-actions:repository(deployport-myapp). Within a clause, different qualifiers are ANDed and a repeated qualifier ORs its values. Clauses are ORed. OIDC principals require at least one clause; IAM principals may omit constraints.
type TrustPolicyStatement struct {
	Constraints []string             `json:"constraints,omitempty" yaml:"constraints,omitempty"`
	Effect      TrustStatementEffect `json:"effect,omitempty" yaml:"effect,omitempty"`
}

// GetConstraints returns the value for the field constraints
func (e *TrustPolicyStatement) GetConstraints() []string {
	return e.Constraints
}

// SetConstraints sets the value for the field constraints
func (e *TrustPolicyStatement) SetConstraints(constraints []string) {
	e.Constraints = constraints
}

// GetEffect returns the value for the field effect
func (e *TrustPolicyStatement) GetEffect() TrustStatementEffect {
	return e.Effect
}

// SetEffect sets the value for the field effect
func (e *TrustPolicyStatement) SetEffect(effect TrustStatementEffect) {
	e.Effect = effect
}

// StructPath returns StructPath
func (e *TrustPolicyStatement) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathTrustPolicyStatement.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *TrustPolicyStatement) InitializeDefaults() {
	if e.Effect == "" {
		e.Effect = TrustStatementEffectAllow
	}
}

// trustPolicyStatementAlias is defined to help pre and post JSON marshaling without recursive loops
type trustPolicyStatementAlias TrustPolicyStatement

// UnmarshalJSON implements json.Unmarshaler
func (e *TrustPolicyStatement) UnmarshalJSON(data []byte) error {
	var alias trustPolicyStatementAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*TrustPolicyStatement)(&alias)).InitializeDefaults()
	*e = TrustPolicyStatement(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e TrustPolicyStatement) MarshalJSON() ([]byte, error) {
	alias := trustPolicyStatementAlias(e)
	if alias.Effect == TrustStatementEffectAllow {
		alias.Effect = ""
	}
	return json.Marshal(alias)
}

// NewTrustPolicy creates a new TrustPolicy
func NewTrustPolicy() *TrustPolicy {
	s := &TrustPolicy{}
	s.InitializeDefaults()
	return s
}

// TrustPolicy - A reusable, attachable trust policy authorizing a single principal to assume a role
// under the given conditions. Attach it to roles via Role.TrustPolicy.Attach.
type TrustPolicy struct {
	// when the trust policy was created.
	// This value is always set. An absent value indicates a server fault, not a state.
	CreatedAt *time.Time `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
	// DRN of this trust policy itself, e.g. iam:TrustPolicy(gha-main). Distinct from principal below.
	Drn string `json:"drn,omitempty" yaml:"drn,omitempty"`
	// present only for a service-linked role's trust policy: read-only to the account
	ManagedBy *ManagedByService `json:"managedBy,omitempty" yaml:"managedBy,omitempty"`
	Name      string            `json:"name,omitempty" yaml:"name,omitempty"`
	// DRN of the trusted principal, e.g. iam:OIDCProvider(github-actions)
	Principal string `json:"principal,omitempty" yaml:"principal,omitempty"`
	// OR of statements; a token is admitted if any statement holds
	Statements []*TrustPolicyStatement `json:"statements,omitempty" yaml:"statements,omitempty"`
}

// GetCreatedAt returns the value for the field createdAt
func (e *TrustPolicy) GetCreatedAt() *time.Time {
	return e.CreatedAt
}

// SetCreatedAt sets the value for the field createdAt
func (e *TrustPolicy) SetCreatedAt(createdAt *time.Time) {
	e.CreatedAt = createdAt
}

// GetDrn returns the value for the field drn
func (e *TrustPolicy) GetDrn() string {
	return e.Drn
}

// SetDrn sets the value for the field drn
func (e *TrustPolicy) SetDrn(drn string) {
	e.Drn = drn
}

// GetManagedBy returns the value for the field managedBy
func (e *TrustPolicy) GetManagedBy() *ManagedByService {
	return e.ManagedBy
}

// SetManagedBy sets the value for the field managedBy
func (e *TrustPolicy) SetManagedBy(managedBy *ManagedByService) {
	e.ManagedBy = managedBy
}

// GetName returns the value for the field name
func (e *TrustPolicy) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *TrustPolicy) SetName(name string) {
	e.Name = name
}

// GetPrincipal returns the value for the field principal
func (e *TrustPolicy) GetPrincipal() string {
	return e.Principal
}

// SetPrincipal sets the value for the field principal
func (e *TrustPolicy) SetPrincipal(principal string) {
	e.Principal = principal
}

// GetStatements returns the value for the field statements
func (e *TrustPolicy) GetStatements() []*TrustPolicyStatement {
	return e.Statements
}

// SetStatements sets the value for the field statements
func (e *TrustPolicy) SetStatements(statements []*TrustPolicyStatement) {
	e.Statements = statements
}

// StructPath returns StructPath
func (e *TrustPolicy) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathTrustPolicy.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *TrustPolicy) InitializeDefaults() {
}

// trustPolicyAlias is defined to help pre and post JSON marshaling without recursive loops
type trustPolicyAlias TrustPolicy

// UnmarshalJSON implements json.Unmarshaler
func (e *TrustPolicy) UnmarshalJSON(data []byte) error {
	var alias trustPolicyAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*TrustPolicy)(&alias)).InitializeDefaults()
	*e = TrustPolicy(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e TrustPolicy) MarshalJSON() ([]byte, error) {
	alias := trustPolicyAlias(e)
	return json.Marshal(alias)
}

// NewTrustPolicyAttachment creates a new TrustPolicyAttachment
func NewTrustPolicyAttachment() *TrustPolicyAttachment {
	s := &TrustPolicyAttachment{}
	s.InitializeDefaults()
	return s
}

// TrustPolicyAttachment - A trust policy attached to a target (a role), including when it was attached.
type TrustPolicyAttachment struct {
	// when the policy was attached to the target.
	// This value is always set. An absent value indicates a server fault, not a state.
	CreatedAt *time.Time `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
	// DRN of the target the policy is attached to, e.g. iam:Role(deployer)
	TargetDRN       string `json:"targetDRN,omitempty" yaml:"targetDRN,omitempty"`
	TrustPolicyName string `json:"trustPolicyName,omitempty" yaml:"trustPolicyName,omitempty"`
}

// GetCreatedAt returns the value for the field createdAt
func (e *TrustPolicyAttachment) GetCreatedAt() *time.Time {
	return e.CreatedAt
}

// SetCreatedAt sets the value for the field createdAt
func (e *TrustPolicyAttachment) SetCreatedAt(createdAt *time.Time) {
	e.CreatedAt = createdAt
}

// GetTargetDRN returns the value for the field targetDRN
func (e *TrustPolicyAttachment) GetTargetDRN() string {
	return e.TargetDRN
}

// SetTargetDRN sets the value for the field targetDRN
func (e *TrustPolicyAttachment) SetTargetDRN(targetDRN string) {
	e.TargetDRN = targetDRN
}

// GetTrustPolicyName returns the value for the field trustPolicyName
func (e *TrustPolicyAttachment) GetTrustPolicyName() string {
	return e.TrustPolicyName
}

// SetTrustPolicyName sets the value for the field trustPolicyName
func (e *TrustPolicyAttachment) SetTrustPolicyName(trustPolicyName string) {
	e.TrustPolicyName = trustPolicyName
}

// StructPath returns StructPath
func (e *TrustPolicyAttachment) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathTrustPolicyAttachment.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *TrustPolicyAttachment) InitializeDefaults() {
}

// trustPolicyAttachmentAlias is defined to help pre and post JSON marshaling without recursive loops
type trustPolicyAttachmentAlias TrustPolicyAttachment

// UnmarshalJSON implements json.Unmarshaler
func (e *TrustPolicyAttachment) UnmarshalJSON(data []byte) error {
	var alias trustPolicyAttachmentAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*TrustPolicyAttachment)(&alias)).InitializeDefaults()
	*e = TrustPolicyAttachment(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e TrustPolicyAttachment) MarshalJSON() ([]byte, error) {
	alias := trustPolicyAttachmentAlias(e)
	return json.Marshal(alias)
}

// NewInvalidTrustPolicyError creates a new InvalidTrustPolicyError
func NewInvalidTrustPolicyError() *InvalidTrustPolicyError {
	s := &InvalidTrustPolicyError{}
	s.InitializeDefaults()
	return s
}

// InvalidTrustPolicyError struct
type InvalidTrustPolicyError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InvalidTrustPolicyError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InvalidTrustPolicyError]
func (e *InvalidTrustPolicyError) Is(err error) bool {
	_, ok := err.(*InvalidTrustPolicyError)
	return ok
}

// IsInvalidTrustPolicyError indicates whether the given error chain contains an error of type [InvalidTrustPolicyError]
func IsInvalidTrustPolicyError(err error) bool {
	return errors.Is(err, &InvalidTrustPolicyError{})
}

// GetMessage returns the value for the field message
func (e *InvalidTrustPolicyError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InvalidTrustPolicyError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InvalidTrustPolicyError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvalidTrustPolicyError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvalidTrustPolicyError) InitializeDefaults() {
}

// invalidTrustPolicyErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type invalidTrustPolicyErrorAlias InvalidTrustPolicyError

// UnmarshalJSON implements json.Unmarshaler
func (e *InvalidTrustPolicyError) UnmarshalJSON(data []byte) error {
	var alias invalidTrustPolicyErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvalidTrustPolicyError)(&alias)).InitializeDefaults()
	*e = InvalidTrustPolicyError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvalidTrustPolicyError) MarshalJSON() ([]byte, error) {
	alias := invalidTrustPolicyErrorAlias(e)
	return json.Marshal(alias)
}

// NewTrustPolicyNotFoundError creates a new TrustPolicyNotFoundError
func NewTrustPolicyNotFoundError() *TrustPolicyNotFoundError {
	s := &TrustPolicyNotFoundError{}
	s.InitializeDefaults()
	return s
}

// TrustPolicyNotFoundError struct
type TrustPolicyNotFoundError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *TrustPolicyNotFoundError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [TrustPolicyNotFoundError]
func (e *TrustPolicyNotFoundError) Is(err error) bool {
	_, ok := err.(*TrustPolicyNotFoundError)
	return ok
}

// IsTrustPolicyNotFoundError indicates whether the given error chain contains an error of type [TrustPolicyNotFoundError]
func IsTrustPolicyNotFoundError(err error) bool {
	return errors.Is(err, &TrustPolicyNotFoundError{})
}

// GetMessage returns the value for the field message
func (e *TrustPolicyNotFoundError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *TrustPolicyNotFoundError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *TrustPolicyNotFoundError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathTrustPolicyNotFoundError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *TrustPolicyNotFoundError) InitializeDefaults() {
}

// trustPolicyNotFoundErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type trustPolicyNotFoundErrorAlias TrustPolicyNotFoundError

// UnmarshalJSON implements json.Unmarshaler
func (e *TrustPolicyNotFoundError) UnmarshalJSON(data []byte) error {
	var alias trustPolicyNotFoundErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*TrustPolicyNotFoundError)(&alias)).InitializeDefaults()
	*e = TrustPolicyNotFoundError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e TrustPolicyNotFoundError) MarshalJSON() ([]byte, error) {
	alias := trustPolicyNotFoundErrorAlias(e)
	return json.Marshal(alias)
}

// NewInvalidWebIdentityTokenError creates a new InvalidWebIdentityTokenError
func NewInvalidWebIdentityTokenError() *InvalidWebIdentityTokenError {
	s := &InvalidWebIdentityTokenError{}
	s.InitializeDefaults()
	return s
}

// InvalidWebIdentityTokenError - Raised by Role.AssumeWithWebIdentity for any invalid, expired, untrusted, or
// unmatched web-identity token. Generic on purpose.
type InvalidWebIdentityTokenError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InvalidWebIdentityTokenError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InvalidWebIdentityTokenError]
func (e *InvalidWebIdentityTokenError) Is(err error) bool {
	_, ok := err.(*InvalidWebIdentityTokenError)
	return ok
}

// IsInvalidWebIdentityTokenError indicates whether the given error chain contains an error of type [InvalidWebIdentityTokenError]
func IsInvalidWebIdentityTokenError(err error) bool {
	return errors.Is(err, &InvalidWebIdentityTokenError{})
}

// GetMessage returns the value for the field message
func (e *InvalidWebIdentityTokenError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InvalidWebIdentityTokenError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InvalidWebIdentityTokenError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvalidWebIdentityTokenError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvalidWebIdentityTokenError) InitializeDefaults() {
}

// invalidWebIdentityTokenErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type invalidWebIdentityTokenErrorAlias InvalidWebIdentityTokenError

// UnmarshalJSON implements json.Unmarshaler
func (e *InvalidWebIdentityTokenError) UnmarshalJSON(data []byte) error {
	var alias invalidWebIdentityTokenErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvalidWebIdentityTokenError)(&alias)).InitializeDefaults()
	*e = InvalidWebIdentityTokenError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvalidWebIdentityTokenError) MarshalJSON() ([]byte, error) {
	alias := invalidWebIdentityTokenErrorAlias(e)
	return json.Marshal(alias)
}

// NewInvalidPrincipalDRNError creates a new InvalidPrincipalDRNError
func NewInvalidPrincipalDRNError() *InvalidPrincipalDRNError {
	s := &InvalidPrincipalDRNError{}
	s.InitializeDefaults()
	return s
}

// InvalidPrincipalDRNError - Raised when a principal DRN is malformed or incomplete, e.g. missing the
// required account(<name>) qualifier, or not targeting an iam:Role(<name>).
type InvalidPrincipalDRNError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InvalidPrincipalDRNError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InvalidPrincipalDRNError]
func (e *InvalidPrincipalDRNError) Is(err error) bool {
	_, ok := err.(*InvalidPrincipalDRNError)
	return ok
}

// IsInvalidPrincipalDRNError indicates whether the given error chain contains an error of type [InvalidPrincipalDRNError]
func IsInvalidPrincipalDRNError(err error) bool {
	return errors.Is(err, &InvalidPrincipalDRNError{})
}

// GetMessage returns the value for the field message
func (e *InvalidPrincipalDRNError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InvalidPrincipalDRNError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InvalidPrincipalDRNError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvalidPrincipalDRNError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvalidPrincipalDRNError) InitializeDefaults() {
}

// invalidPrincipalDRNErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type invalidPrincipalDRNErrorAlias InvalidPrincipalDRNError

// UnmarshalJSON implements json.Unmarshaler
func (e *InvalidPrincipalDRNError) UnmarshalJSON(data []byte) error {
	var alias invalidPrincipalDRNErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvalidPrincipalDRNError)(&alias)).InitializeDefaults()
	*e = InvalidPrincipalDRNError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvalidPrincipalDRNError) MarshalJSON() ([]byte, error) {
	alias := invalidPrincipalDRNErrorAlias(e)
	return json.Marshal(alias)
}

// NewInvalidHandoffCodeError creates a new InvalidHandoffCodeError
func NewInvalidHandoffCodeError() *InvalidHandoffCodeError {
	s := &InvalidHandoffCodeError{}
	s.InitializeDefaults()
	return s
}

// InvalidHandoffCodeError - Raised by Session.CompleteHandoff for any code that cannot be redeemed: malformed, unknown, expired, already consumed, issued by a session that has since been revoked, or no longer permitted.
// Generic on purpose, so a caller learns nothing beyond "invalid", and a redemption that fails this way has still consumed the code.
type InvalidHandoffCodeError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InvalidHandoffCodeError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InvalidHandoffCodeError]
func (e *InvalidHandoffCodeError) Is(err error) bool {
	_, ok := err.(*InvalidHandoffCodeError)
	return ok
}

// IsInvalidHandoffCodeError indicates whether the given error chain contains an error of type [InvalidHandoffCodeError]
func IsInvalidHandoffCodeError(err error) bool {
	return errors.Is(err, &InvalidHandoffCodeError{})
}

// GetMessage returns the value for the field message
func (e *InvalidHandoffCodeError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InvalidHandoffCodeError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InvalidHandoffCodeError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvalidHandoffCodeError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvalidHandoffCodeError) InitializeDefaults() {
}

// invalidHandoffCodeErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type invalidHandoffCodeErrorAlias InvalidHandoffCodeError

// UnmarshalJSON implements json.Unmarshaler
func (e *InvalidHandoffCodeError) UnmarshalJSON(data []byte) error {
	var alias invalidHandoffCodeErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvalidHandoffCodeError)(&alias)).InitializeDefaults()
	*e = InvalidHandoffCodeError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvalidHandoffCodeError) MarshalJSON() ([]byte, error) {
	alias := invalidHandoffCodeErrorAlias(e)
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

// NewAccountCreateInvalidNameError creates a new AccountCreateInvalidNameError
func NewAccountCreateInvalidNameError() *AccountCreateInvalidNameError {
	s := &AccountCreateInvalidNameError{}
	s.InitializeDefaults()
	return s
}

// AccountCreateInvalidNameError struct
type AccountCreateInvalidNameError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *AccountCreateInvalidNameError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [AccountCreateInvalidNameError]
func (e *AccountCreateInvalidNameError) Is(err error) bool {
	_, ok := err.(*AccountCreateInvalidNameError)
	return ok
}

// IsAccountCreateInvalidNameError indicates whether the given error chain contains an error of type [AccountCreateInvalidNameError]
func IsAccountCreateInvalidNameError(err error) bool {
	return errors.Is(err, &AccountCreateInvalidNameError{})
}

// GetMessage returns the value for the field message
func (e *AccountCreateInvalidNameError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *AccountCreateInvalidNameError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *AccountCreateInvalidNameError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountCreateInvalidNameError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountCreateInvalidNameError) InitializeDefaults() {
}

// accountCreateInvalidNameErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type accountCreateInvalidNameErrorAlias AccountCreateInvalidNameError

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountCreateInvalidNameError) UnmarshalJSON(data []byte) error {
	var alias accountCreateInvalidNameErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountCreateInvalidNameError)(&alias)).InitializeDefaults()
	*e = AccountCreateInvalidNameError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountCreateInvalidNameError) MarshalJSON() ([]byte, error) {
	alias := accountCreateInvalidNameErrorAlias(e)
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
	// idle/sliding window in seconds (see Role.Assume). Optional; defaults
	// to 57600 (16h) for an interactive SSO session, 3600 (1h) otherwise.
	DurationSeconds *int32 `json:"durationSeconds,omitempty" yaml:"durationSeconds,omitempty"`
	// optional hard max-lifetime in seconds, clamped to the user's max.
	MaxLifetimeSeconds *int32 `json:"maxLifetimeSeconds,omitempty" yaml:"maxLifetimeSeconds,omitempty"`
}

// GetAccountName returns the value for the field accountName
func (e *AccountAssumeIdentityInput) GetAccountName() string {
	return e.AccountName
}

// SetAccountName sets the value for the field accountName
func (e *AccountAssumeIdentityInput) SetAccountName(accountName string) {
	e.AccountName = accountName
}

// GetDurationSeconds returns the value for the field durationSeconds
func (e *AccountAssumeIdentityInput) GetDurationSeconds() *int32 {
	return e.DurationSeconds
}

// SetDurationSeconds sets the value for the field durationSeconds
func (e *AccountAssumeIdentityInput) SetDurationSeconds(durationSeconds *int32) {
	e.DurationSeconds = durationSeconds
}

// GetMaxLifetimeSeconds returns the value for the field maxLifetimeSeconds
func (e *AccountAssumeIdentityInput) GetMaxLifetimeSeconds() *int32 {
	return e.MaxLifetimeSeconds
}

// SetMaxLifetimeSeconds sets the value for the field maxLifetimeSeconds
func (e *AccountAssumeIdentityInput) SetMaxLifetimeSeconds(maxLifetimeSeconds *int32) {
	e.MaxLifetimeSeconds = maxLifetimeSeconds
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

// NewAccountBeginAssumeIdentityInput creates a new AccountBeginAssumeIdentityInput
func NewAccountBeginAssumeIdentityInput() *AccountBeginAssumeIdentityInput {
	s := &AccountBeginAssumeIdentityInput{}
	s.InitializeDefaults()
	return s
}

// AccountBeginAssumeIdentityInput struct
type AccountBeginAssumeIdentityInput struct {
	AccountName string `json:"accountName,omitempty" yaml:"accountName,omitempty"`
}

// GetAccountName returns the value for the field accountName
func (e *AccountBeginAssumeIdentityInput) GetAccountName() string {
	return e.AccountName
}

// SetAccountName sets the value for the field accountName
func (e *AccountBeginAssumeIdentityInput) SetAccountName(accountName string) {
	e.AccountName = accountName
}

// StructPath returns StructPath
func (e *AccountBeginAssumeIdentityInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountBeginAssumeIdentityInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountBeginAssumeIdentityInput) InitializeDefaults() {
}

// accountBeginAssumeIdentityInputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountBeginAssumeIdentityInputAlias AccountBeginAssumeIdentityInput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountBeginAssumeIdentityInput) UnmarshalJSON(data []byte) error {
	var alias accountBeginAssumeIdentityInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountBeginAssumeIdentityInput)(&alias)).InitializeDefaults()
	*e = AccountBeginAssumeIdentityInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountBeginAssumeIdentityInput) MarshalJSON() ([]byte, error) {
	alias := accountBeginAssumeIdentityInputAlias(e)
	return json.Marshal(alias)
}

// NewAccountBeginAssumeIdentityOutput creates a new AccountBeginAssumeIdentityOutput
func NewAccountBeginAssumeIdentityOutput() *AccountBeginAssumeIdentityOutput {
	s := &AccountBeginAssumeIdentityOutput{}
	s.InitializeDefaults()
	return s
}

// AccountBeginAssumeIdentityOutput struct
type AccountBeginAssumeIdentityOutput struct {
	// single-use opaque code; redeem within a short window via CompleteAssumeIdentity
	Code string `json:"code,omitempty" yaml:"code,omitempty"`
}

// GetCode returns the value for the field code
func (e *AccountBeginAssumeIdentityOutput) GetCode() string {
	return e.Code
}

// SetCode sets the value for the field code
func (e *AccountBeginAssumeIdentityOutput) SetCode(code string) {
	e.Code = code
}

// StructPath returns StructPath
func (e *AccountBeginAssumeIdentityOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountBeginAssumeIdentityOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountBeginAssumeIdentityOutput) InitializeDefaults() {
}

// accountBeginAssumeIdentityOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountBeginAssumeIdentityOutputAlias AccountBeginAssumeIdentityOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountBeginAssumeIdentityOutput) UnmarshalJSON(data []byte) error {
	var alias accountBeginAssumeIdentityOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountBeginAssumeIdentityOutput)(&alias)).InitializeDefaults()
	*e = AccountBeginAssumeIdentityOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountBeginAssumeIdentityOutput) MarshalJSON() ([]byte, error) {
	alias := accountBeginAssumeIdentityOutputAlias(e)
	return json.Marshal(alias)
}

// NewAccountCompleteAssumeIdentityInput creates a new AccountCompleteAssumeIdentityInput
func NewAccountCompleteAssumeIdentityInput() *AccountCompleteAssumeIdentityInput {
	s := &AccountCompleteAssumeIdentityInput{}
	s.InitializeDefaults()
	return s
}

// AccountCompleteAssumeIdentityInput struct
type AccountCompleteAssumeIdentityInput struct {
	AccountName string `json:"accountName,omitempty" yaml:"accountName,omitempty"`
	Code        string `json:"code,omitempty" yaml:"code,omitempty"`
}

// GetAccountName returns the value for the field accountName
func (e *AccountCompleteAssumeIdentityInput) GetAccountName() string {
	return e.AccountName
}

// SetAccountName sets the value for the field accountName
func (e *AccountCompleteAssumeIdentityInput) SetAccountName(accountName string) {
	e.AccountName = accountName
}

// GetCode returns the value for the field code
func (e *AccountCompleteAssumeIdentityInput) GetCode() string {
	return e.Code
}

// SetCode sets the value for the field code
func (e *AccountCompleteAssumeIdentityInput) SetCode(code string) {
	e.Code = code
}

// StructPath returns StructPath
func (e *AccountCompleteAssumeIdentityInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountCompleteAssumeIdentityInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountCompleteAssumeIdentityInput) InitializeDefaults() {
}

// accountCompleteAssumeIdentityInputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountCompleteAssumeIdentityInputAlias AccountCompleteAssumeIdentityInput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountCompleteAssumeIdentityInput) UnmarshalJSON(data []byte) error {
	var alias accountCompleteAssumeIdentityInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountCompleteAssumeIdentityInput)(&alias)).InitializeDefaults()
	*e = AccountCompleteAssumeIdentityInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountCompleteAssumeIdentityInput) MarshalJSON() ([]byte, error) {
	alias := accountCompleteAssumeIdentityInputAlias(e)
	return json.Marshal(alias)
}

// NewAccountCompleteAssumeIdentityOutput creates a new AccountCompleteAssumeIdentityOutput
func NewAccountCompleteAssumeIdentityOutput() *AccountCompleteAssumeIdentityOutput {
	s := &AccountCompleteAssumeIdentityOutput{}
	s.InitializeDefaults()
	return s
}

// AccountCompleteAssumeIdentityOutput struct
type AccountCompleteAssumeIdentityOutput struct {
	Credentials *Credentials `json:"credentials,omitempty" yaml:"credentials,omitempty"`
}

// GetCredentials returns the value for the field credentials
func (e *AccountCompleteAssumeIdentityOutput) GetCredentials() *Credentials {
	return e.Credentials
}

// SetCredentials sets the value for the field credentials
func (e *AccountCompleteAssumeIdentityOutput) SetCredentials(credentials *Credentials) {
	e.Credentials = credentials
}

// StructPath returns StructPath
func (e *AccountCompleteAssumeIdentityOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountCompleteAssumeIdentityOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountCompleteAssumeIdentityOutput) InitializeDefaults() {
}

// accountCompleteAssumeIdentityOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountCompleteAssumeIdentityOutputAlias AccountCompleteAssumeIdentityOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountCompleteAssumeIdentityOutput) UnmarshalJSON(data []byte) error {
	var alias accountCompleteAssumeIdentityOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountCompleteAssumeIdentityOutput)(&alias)).InitializeDefaults()
	*e = AccountCompleteAssumeIdentityOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountCompleteAssumeIdentityOutput) MarshalJSON() ([]byte, error) {
	alias := accountCompleteAssumeIdentityOutputAlias(e)
	return json.Marshal(alias)
}

// NewAccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError creates a new AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError
func NewAccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError() *AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError {
	s := &AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError{}
	s.InitializeDefaults()
	return s
}

// AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError struct
type AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError]
func (e *AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError) Is(err error) bool {
	_, ok := err.(*AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError)
	return ok
}

// IsAccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError indicates whether the given error chain contains an error of type [AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError]
func IsAccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError(err error) bool {
	return errors.Is(err, &AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError{})
}

// GetMessage returns the value for the field message
func (e *AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError) InitializeDefaults() {
}

// accountCompleteAssumeIdentityInvalidAssumeIdentityCodeErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type accountCompleteAssumeIdentityInvalidAssumeIdentityCodeErrorAlias AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError) UnmarshalJSON(data []byte) error {
	var alias accountCompleteAssumeIdentityInvalidAssumeIdentityCodeErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError)(&alias)).InitializeDefaults()
	*e = AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError) MarshalJSON() ([]byte, error) {
	alias := accountCompleteAssumeIdentityInvalidAssumeIdentityCodeErrorAlias(e)
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
	// optional sliding idle window, in seconds, for the credential this
	// login mints. Applied ONE-WAY: a value larger than the default is
	// clamped down to it, so a client can only tighten its own session.
	// Omit to get the default (57600, 16h).
	IdleWindowSeconds *int32 `json:"idleWindowSeconds,omitempty" yaml:"idleWindowSeconds,omitempty"`
	// optional hard cap, in seconds, on the total lifetime of the
	// credential this login mints, regardless of activity or keep-alive.
	// One-way like idleWindowSeconds: it can only shorten, never extend.
	// Omit to get the default (604800, 7d).
	MaxSessionLifetimeSeconds *int32 `json:"maxSessionLifetimeSeconds,omitempty" yaml:"maxSessionLifetimeSeconds,omitempty"`
	ProviderName              string `json:"providerName,omitempty" yaml:"providerName,omitempty"`
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

// GetIdleWindowSeconds returns the value for the field idleWindowSeconds
func (e *AccountSSOBeginAuthenticationInput) GetIdleWindowSeconds() *int32 {
	return e.IdleWindowSeconds
}

// SetIdleWindowSeconds sets the value for the field idleWindowSeconds
func (e *AccountSSOBeginAuthenticationInput) SetIdleWindowSeconds(idleWindowSeconds *int32) {
	e.IdleWindowSeconds = idleWindowSeconds
}

// GetMaxSessionLifetimeSeconds returns the value for the field maxSessionLifetimeSeconds
func (e *AccountSSOBeginAuthenticationInput) GetMaxSessionLifetimeSeconds() *int32 {
	return e.MaxSessionLifetimeSeconds
}

// SetMaxSessionLifetimeSeconds sets the value for the field maxSessionLifetimeSeconds
func (e *AccountSSOBeginAuthenticationInput) SetMaxSessionLifetimeSeconds(maxSessionLifetimeSeconds *int32) {
	e.MaxSessionLifetimeSeconds = maxSessionLifetimeSeconds
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

// NewAccountSSOBeginAuthenticationParameterError creates a new AccountSSOBeginAuthenticationParameterError
func NewAccountSSOBeginAuthenticationParameterError() *AccountSSOBeginAuthenticationParameterError {
	s := &AccountSSOBeginAuthenticationParameterError{}
	s.InitializeDefaults()
	return s
}

// AccountSSOBeginAuthenticationParameterError struct
type AccountSSOBeginAuthenticationParameterError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *AccountSSOBeginAuthenticationParameterError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [AccountSSOBeginAuthenticationParameterError]
func (e *AccountSSOBeginAuthenticationParameterError) Is(err error) bool {
	_, ok := err.(*AccountSSOBeginAuthenticationParameterError)
	return ok
}

// IsAccountSSOBeginAuthenticationParameterError indicates whether the given error chain contains an error of type [AccountSSOBeginAuthenticationParameterError]
func IsAccountSSOBeginAuthenticationParameterError(err error) bool {
	return errors.Is(err, &AccountSSOBeginAuthenticationParameterError{})
}

// GetMessage returns the value for the field message
func (e *AccountSSOBeginAuthenticationParameterError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *AccountSSOBeginAuthenticationParameterError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *AccountSSOBeginAuthenticationParameterError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOBeginAuthenticationParameterError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountSSOBeginAuthenticationParameterError) InitializeDefaults() {
}

// accountSSOBeginAuthenticationParameterErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type accountSSOBeginAuthenticationParameterErrorAlias AccountSSOBeginAuthenticationParameterError

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountSSOBeginAuthenticationParameterError) UnmarshalJSON(data []byte) error {
	var alias accountSSOBeginAuthenticationParameterErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountSSOBeginAuthenticationParameterError)(&alias)).InitializeDefaults()
	*e = AccountSSOBeginAuthenticationParameterError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountSSOBeginAuthenticationParameterError) MarshalJSON() ([]byte, error) {
	alias := accountSSOBeginAuthenticationParameterErrorAlias(e)
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

// NewAccountSSOCompleteAuthenticationInvalidFlowError creates a new AccountSSOCompleteAuthenticationInvalidFlowError
func NewAccountSSOCompleteAuthenticationInvalidFlowError() *AccountSSOCompleteAuthenticationInvalidFlowError {
	s := &AccountSSOCompleteAuthenticationInvalidFlowError{}
	s.InitializeDefaults()
	return s
}

// AccountSSOCompleteAuthenticationInvalidFlowError struct
type AccountSSOCompleteAuthenticationInvalidFlowError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *AccountSSOCompleteAuthenticationInvalidFlowError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [AccountSSOCompleteAuthenticationInvalidFlowError]
func (e *AccountSSOCompleteAuthenticationInvalidFlowError) Is(err error) bool {
	_, ok := err.(*AccountSSOCompleteAuthenticationInvalidFlowError)
	return ok
}

// IsAccountSSOCompleteAuthenticationInvalidFlowError indicates whether the given error chain contains an error of type [AccountSSOCompleteAuthenticationInvalidFlowError]
func IsAccountSSOCompleteAuthenticationInvalidFlowError(err error) bool {
	return errors.Is(err, &AccountSSOCompleteAuthenticationInvalidFlowError{})
}

// GetMessage returns the value for the field message
func (e *AccountSSOCompleteAuthenticationInvalidFlowError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *AccountSSOCompleteAuthenticationInvalidFlowError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *AccountSSOCompleteAuthenticationInvalidFlowError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountSSOCompleteAuthenticationInvalidFlowError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountSSOCompleteAuthenticationInvalidFlowError) InitializeDefaults() {
}

// accountSSOCompleteAuthenticationInvalidFlowErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type accountSSOCompleteAuthenticationInvalidFlowErrorAlias AccountSSOCompleteAuthenticationInvalidFlowError

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountSSOCompleteAuthenticationInvalidFlowError) UnmarshalJSON(data []byte) error {
	var alias accountSSOCompleteAuthenticationInvalidFlowErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountSSOCompleteAuthenticationInvalidFlowError)(&alias)).InitializeDefaults()
	*e = AccountSSOCompleteAuthenticationInvalidFlowError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountSSOCompleteAuthenticationInvalidFlowError) MarshalJSON() ([]byte, error) {
	alias := accountSSOCompleteAuthenticationInvalidFlowErrorAlias(e)
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

// NewAccountOIDCProviderCreateInput creates a new AccountOIDCProviderCreateInput
func NewAccountOIDCProviderCreateInput() *AccountOIDCProviderCreateInput {
	s := &AccountOIDCProviderCreateInput{}
	s.InitializeDefaults()
	return s
}

// AccountOIDCProviderCreateInput struct
type AccountOIDCProviderCreateInput struct {
	// Allowlist of accepted token aud values. Opaque identifiers, not
	// URLs; e.g. iam.deployport.io. See the operation docs.
	Audiences []string `json:"audiences,omitempty" yaml:"audiences,omitempty"`
	IssuerURL string   `json:"issuerURL,omitempty" yaml:"issuerURL,omitempty"`
	Name      string   `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetAudiences returns the value for the field audiences
func (e *AccountOIDCProviderCreateInput) GetAudiences() []string {
	return e.Audiences
}

// SetAudiences sets the value for the field audiences
func (e *AccountOIDCProviderCreateInput) SetAudiences(audiences []string) {
	e.Audiences = audiences
}

// GetIssuerURL returns the value for the field issuerURL
func (e *AccountOIDCProviderCreateInput) GetIssuerURL() string {
	return e.IssuerURL
}

// SetIssuerURL sets the value for the field issuerURL
func (e *AccountOIDCProviderCreateInput) SetIssuerURL(issuerURL string) {
	e.IssuerURL = issuerURL
}

// GetName returns the value for the field name
func (e *AccountOIDCProviderCreateInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *AccountOIDCProviderCreateInput) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *AccountOIDCProviderCreateInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountOIDCProviderCreateInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountOIDCProviderCreateInput) InitializeDefaults() {
}

// accountOIDCProviderCreateInputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountOIDCProviderCreateInputAlias AccountOIDCProviderCreateInput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountOIDCProviderCreateInput) UnmarshalJSON(data []byte) error {
	var alias accountOIDCProviderCreateInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountOIDCProviderCreateInput)(&alias)).InitializeDefaults()
	*e = AccountOIDCProviderCreateInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountOIDCProviderCreateInput) MarshalJSON() ([]byte, error) {
	alias := accountOIDCProviderCreateInputAlias(e)
	return json.Marshal(alias)
}

// NewAccountOIDCProviderCreateOutput creates a new AccountOIDCProviderCreateOutput
func NewAccountOIDCProviderCreateOutput() *AccountOIDCProviderCreateOutput {
	s := &AccountOIDCProviderCreateOutput{}
	s.InitializeDefaults()
	return s
}

// AccountOIDCProviderCreateOutput struct
type AccountOIDCProviderCreateOutput struct {
	Provider *OIDCProvider `json:"provider,omitempty" yaml:"provider,omitempty"`
}

// GetProvider returns the value for the field provider
func (e *AccountOIDCProviderCreateOutput) GetProvider() *OIDCProvider {
	return e.Provider
}

// SetProvider sets the value for the field provider
func (e *AccountOIDCProviderCreateOutput) SetProvider(provider *OIDCProvider) {
	e.Provider = provider
}

// StructPath returns StructPath
func (e *AccountOIDCProviderCreateOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountOIDCProviderCreateOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountOIDCProviderCreateOutput) InitializeDefaults() {
}

// accountOIDCProviderCreateOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountOIDCProviderCreateOutputAlias AccountOIDCProviderCreateOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountOIDCProviderCreateOutput) UnmarshalJSON(data []byte) error {
	var alias accountOIDCProviderCreateOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountOIDCProviderCreateOutput)(&alias)).InitializeDefaults()
	*e = AccountOIDCProviderCreateOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountOIDCProviderCreateOutput) MarshalJSON() ([]byte, error) {
	alias := accountOIDCProviderCreateOutputAlias(e)
	return json.Marshal(alias)
}

// NewAccountOIDCProviderSetAudiencesInput creates a new AccountOIDCProviderSetAudiencesInput
func NewAccountOIDCProviderSetAudiencesInput() *AccountOIDCProviderSetAudiencesInput {
	s := &AccountOIDCProviderSetAudiencesInput{}
	s.InitializeDefaults()
	return s
}

// AccountOIDCProviderSetAudiencesInput struct
type AccountOIDCProviderSetAudiencesInput struct {
	Audiences []string `json:"audiences,omitempty" yaml:"audiences,omitempty"`
	Name      string   `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetAudiences returns the value for the field audiences
func (e *AccountOIDCProviderSetAudiencesInput) GetAudiences() []string {
	return e.Audiences
}

// SetAudiences sets the value for the field audiences
func (e *AccountOIDCProviderSetAudiencesInput) SetAudiences(audiences []string) {
	e.Audiences = audiences
}

// GetName returns the value for the field name
func (e *AccountOIDCProviderSetAudiencesInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *AccountOIDCProviderSetAudiencesInput) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *AccountOIDCProviderSetAudiencesInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountOIDCProviderSetAudiencesInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountOIDCProviderSetAudiencesInput) InitializeDefaults() {
}

// accountOIDCProviderSetAudiencesInputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountOIDCProviderSetAudiencesInputAlias AccountOIDCProviderSetAudiencesInput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountOIDCProviderSetAudiencesInput) UnmarshalJSON(data []byte) error {
	var alias accountOIDCProviderSetAudiencesInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountOIDCProviderSetAudiencesInput)(&alias)).InitializeDefaults()
	*e = AccountOIDCProviderSetAudiencesInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountOIDCProviderSetAudiencesInput) MarshalJSON() ([]byte, error) {
	alias := accountOIDCProviderSetAudiencesInputAlias(e)
	return json.Marshal(alias)
}

// NewAccountOIDCProviderSetAudiencesOutput creates a new AccountOIDCProviderSetAudiencesOutput
func NewAccountOIDCProviderSetAudiencesOutput() *AccountOIDCProviderSetAudiencesOutput {
	s := &AccountOIDCProviderSetAudiencesOutput{}
	s.InitializeDefaults()
	return s
}

// AccountOIDCProviderSetAudiencesOutput struct
type AccountOIDCProviderSetAudiencesOutput struct {
	Provider *OIDCProvider `json:"provider,omitempty" yaml:"provider,omitempty"`
}

// GetProvider returns the value for the field provider
func (e *AccountOIDCProviderSetAudiencesOutput) GetProvider() *OIDCProvider {
	return e.Provider
}

// SetProvider sets the value for the field provider
func (e *AccountOIDCProviderSetAudiencesOutput) SetProvider(provider *OIDCProvider) {
	e.Provider = provider
}

// StructPath returns StructPath
func (e *AccountOIDCProviderSetAudiencesOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountOIDCProviderSetAudiencesOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountOIDCProviderSetAudiencesOutput) InitializeDefaults() {
}

// accountOIDCProviderSetAudiencesOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountOIDCProviderSetAudiencesOutputAlias AccountOIDCProviderSetAudiencesOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountOIDCProviderSetAudiencesOutput) UnmarshalJSON(data []byte) error {
	var alias accountOIDCProviderSetAudiencesOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountOIDCProviderSetAudiencesOutput)(&alias)).InitializeDefaults()
	*e = AccountOIDCProviderSetAudiencesOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountOIDCProviderSetAudiencesOutput) MarshalJSON() ([]byte, error) {
	alias := accountOIDCProviderSetAudiencesOutputAlias(e)
	return json.Marshal(alias)
}

// NewAccountOIDCProviderListInput creates a new AccountOIDCProviderListInput
func NewAccountOIDCProviderListInput() *AccountOIDCProviderListInput {
	s := &AccountOIDCProviderListInput{}
	s.InitializeDefaults()
	return s
}

// AccountOIDCProviderListInput struct
type AccountOIDCProviderListInput struct {
}

// StructPath returns StructPath
func (e *AccountOIDCProviderListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountOIDCProviderListInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountOIDCProviderListInput) InitializeDefaults() {
}

// accountOIDCProviderListInputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountOIDCProviderListInputAlias AccountOIDCProviderListInput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountOIDCProviderListInput) UnmarshalJSON(data []byte) error {
	var alias accountOIDCProviderListInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountOIDCProviderListInput)(&alias)).InitializeDefaults()
	*e = AccountOIDCProviderListInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountOIDCProviderListInput) MarshalJSON() ([]byte, error) {
	alias := accountOIDCProviderListInputAlias(e)
	return json.Marshal(alias)
}

// NewAccountOIDCProviderListOutput creates a new AccountOIDCProviderListOutput
func NewAccountOIDCProviderListOutput() *AccountOIDCProviderListOutput {
	s := &AccountOIDCProviderListOutput{}
	s.InitializeDefaults()
	return s
}

// AccountOIDCProviderListOutput struct
type AccountOIDCProviderListOutput struct {
	Providers []*OIDCProvider `json:"providers,omitempty" yaml:"providers,omitempty"`
}

// GetProviders returns the value for the field providers
func (e *AccountOIDCProviderListOutput) GetProviders() []*OIDCProvider {
	return e.Providers
}

// SetProviders sets the value for the field providers
func (e *AccountOIDCProviderListOutput) SetProviders(providers []*OIDCProvider) {
	e.Providers = providers
}

// StructPath returns StructPath
func (e *AccountOIDCProviderListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountOIDCProviderListOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountOIDCProviderListOutput) InitializeDefaults() {
}

// accountOIDCProviderListOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountOIDCProviderListOutputAlias AccountOIDCProviderListOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountOIDCProviderListOutput) UnmarshalJSON(data []byte) error {
	var alias accountOIDCProviderListOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountOIDCProviderListOutput)(&alias)).InitializeDefaults()
	*e = AccountOIDCProviderListOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountOIDCProviderListOutput) MarshalJSON() ([]byte, error) {
	alias := accountOIDCProviderListOutputAlias(e)
	return json.Marshal(alias)
}

// NewAccountOIDCProviderDeleteInput creates a new AccountOIDCProviderDeleteInput
func NewAccountOIDCProviderDeleteInput() *AccountOIDCProviderDeleteInput {
	s := &AccountOIDCProviderDeleteInput{}
	s.InitializeDefaults()
	return s
}

// AccountOIDCProviderDeleteInput struct
type AccountOIDCProviderDeleteInput struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetName returns the value for the field name
func (e *AccountOIDCProviderDeleteInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *AccountOIDCProviderDeleteInput) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *AccountOIDCProviderDeleteInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountOIDCProviderDeleteInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountOIDCProviderDeleteInput) InitializeDefaults() {
}

// accountOIDCProviderDeleteInputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountOIDCProviderDeleteInputAlias AccountOIDCProviderDeleteInput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountOIDCProviderDeleteInput) UnmarshalJSON(data []byte) error {
	var alias accountOIDCProviderDeleteInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountOIDCProviderDeleteInput)(&alias)).InitializeDefaults()
	*e = AccountOIDCProviderDeleteInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountOIDCProviderDeleteInput) MarshalJSON() ([]byte, error) {
	alias := accountOIDCProviderDeleteInputAlias(e)
	return json.Marshal(alias)
}

// NewAccountOIDCProviderDeleteOutput creates a new AccountOIDCProviderDeleteOutput
func NewAccountOIDCProviderDeleteOutput() *AccountOIDCProviderDeleteOutput {
	s := &AccountOIDCProviderDeleteOutput{}
	s.InitializeDefaults()
	return s
}

// AccountOIDCProviderDeleteOutput struct
type AccountOIDCProviderDeleteOutput struct {
}

// StructPath returns StructPath
func (e *AccountOIDCProviderDeleteOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountOIDCProviderDeleteOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountOIDCProviderDeleteOutput) InitializeDefaults() {
}

// accountOIDCProviderDeleteOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountOIDCProviderDeleteOutputAlias AccountOIDCProviderDeleteOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountOIDCProviderDeleteOutput) UnmarshalJSON(data []byte) error {
	var alias accountOIDCProviderDeleteOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountOIDCProviderDeleteOutput)(&alias)).InitializeDefaults()
	*e = AccountOIDCProviderDeleteOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountOIDCProviderDeleteOutput) MarshalJSON() ([]byte, error) {
	alias := accountOIDCProviderDeleteOutputAlias(e)
	return json.Marshal(alias)
}

// NewAccountOIDCProviderTrustPolicyListInput creates a new AccountOIDCProviderTrustPolicyListInput
func NewAccountOIDCProviderTrustPolicyListInput() *AccountOIDCProviderTrustPolicyListInput {
	s := &AccountOIDCProviderTrustPolicyListInput{}
	s.InitializeDefaults()
	return s
}

// AccountOIDCProviderTrustPolicyListInput struct
type AccountOIDCProviderTrustPolicyListInput struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetName returns the value for the field name
func (e *AccountOIDCProviderTrustPolicyListInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *AccountOIDCProviderTrustPolicyListInput) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *AccountOIDCProviderTrustPolicyListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountOIDCProviderTrustPolicyListInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountOIDCProviderTrustPolicyListInput) InitializeDefaults() {
}

// accountOIDCProviderTrustPolicyListInputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountOIDCProviderTrustPolicyListInputAlias AccountOIDCProviderTrustPolicyListInput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountOIDCProviderTrustPolicyListInput) UnmarshalJSON(data []byte) error {
	var alias accountOIDCProviderTrustPolicyListInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountOIDCProviderTrustPolicyListInput)(&alias)).InitializeDefaults()
	*e = AccountOIDCProviderTrustPolicyListInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountOIDCProviderTrustPolicyListInput) MarshalJSON() ([]byte, error) {
	alias := accountOIDCProviderTrustPolicyListInputAlias(e)
	return json.Marshal(alias)
}

// NewAccountOIDCProviderTrustPolicyListOutput creates a new AccountOIDCProviderTrustPolicyListOutput
func NewAccountOIDCProviderTrustPolicyListOutput() *AccountOIDCProviderTrustPolicyListOutput {
	s := &AccountOIDCProviderTrustPolicyListOutput{}
	s.InitializeDefaults()
	return s
}

// AccountOIDCProviderTrustPolicyListOutput struct
type AccountOIDCProviderTrustPolicyListOutput struct {
	TrustPolicies []*TrustPolicy `json:"trustPolicies,omitempty" yaml:"trustPolicies,omitempty"`
}

// GetTrustPolicies returns the value for the field trustPolicies
func (e *AccountOIDCProviderTrustPolicyListOutput) GetTrustPolicies() []*TrustPolicy {
	return e.TrustPolicies
}

// SetTrustPolicies sets the value for the field trustPolicies
func (e *AccountOIDCProviderTrustPolicyListOutput) SetTrustPolicies(trustPolicies []*TrustPolicy) {
	e.TrustPolicies = trustPolicies
}

// StructPath returns StructPath
func (e *AccountOIDCProviderTrustPolicyListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathAccountOIDCProviderTrustPolicyListOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *AccountOIDCProviderTrustPolicyListOutput) InitializeDefaults() {
}

// accountOIDCProviderTrustPolicyListOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type accountOIDCProviderTrustPolicyListOutputAlias AccountOIDCProviderTrustPolicyListOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *AccountOIDCProviderTrustPolicyListOutput) UnmarshalJSON(data []byte) error {
	var alias accountOIDCProviderTrustPolicyListOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*AccountOIDCProviderTrustPolicyListOutput)(&alias)).InitializeDefaults()
	*e = AccountOIDCProviderTrustPolicyListOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e AccountOIDCProviderTrustPolicyListOutput) MarshalJSON() ([]byte, error) {
	alias := accountOIDCProviderTrustPolicyListOutputAlias(e)
	return json.Marshal(alias)
}

// NewRegionListInput creates a new RegionListInput
func NewRegionListInput() *RegionListInput {
	s := &RegionListInput{}
	s.InitializeDefaults()
	return s
}

// RegionListInput struct
type RegionListInput struct {
}

// StructPath returns StructPath
func (e *RegionListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRegionListInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RegionListInput) InitializeDefaults() {
}

// regionListInputAlias is defined to help pre and post JSON marshaling without recursive loops
type regionListInputAlias RegionListInput

// UnmarshalJSON implements json.Unmarshaler
func (e *RegionListInput) UnmarshalJSON(data []byte) error {
	var alias regionListInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RegionListInput)(&alias)).InitializeDefaults()
	*e = RegionListInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RegionListInput) MarshalJSON() ([]byte, error) {
	alias := regionListInputAlias(e)
	return json.Marshal(alias)
}

// NewRegionListOutput creates a new RegionListOutput
func NewRegionListOutput() *RegionListOutput {
	s := &RegionListOutput{}
	s.InitializeDefaults()
	return s
}

// RegionListOutput struct
type RegionListOutput struct {
	Regions []*RegionInfo `json:"regions,omitempty" yaml:"regions,omitempty"`
}

// GetRegions returns the value for the field regions
func (e *RegionListOutput) GetRegions() []*RegionInfo {
	return e.Regions
}

// SetRegions sets the value for the field regions
func (e *RegionListOutput) SetRegions(regions []*RegionInfo) {
	e.Regions = regions
}

// StructPath returns StructPath
func (e *RegionListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRegionListOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RegionListOutput) InitializeDefaults() {
}

// regionListOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type regionListOutputAlias RegionListOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *RegionListOutput) UnmarshalJSON(data []byte) error {
	var alias regionListOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RegionListOutput)(&alias)).InitializeDefaults()
	*e = RegionListOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RegionListOutput) MarshalJSON() ([]byte, error) {
	alias := regionListOutputAlias(e)
	return json.Marshal(alias)
}

// NewInvalidUsernameError creates a new InvalidUsernameError
func NewInvalidUsernameError() *InvalidUsernameError {
	s := &InvalidUsernameError{}
	s.InitializeDefaults()
	return s
}

// InvalidUsernameError struct
type InvalidUsernameError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InvalidUsernameError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InvalidUsernameError]
func (e *InvalidUsernameError) Is(err error) bool {
	_, ok := err.(*InvalidUsernameError)
	return ok
}

// IsInvalidUsernameError indicates whether the given error chain contains an error of type [InvalidUsernameError]
func IsInvalidUsernameError(err error) bool {
	return errors.Is(err, &InvalidUsernameError{})
}

// GetMessage returns the value for the field message
func (e *InvalidUsernameError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InvalidUsernameError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InvalidUsernameError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvalidUsernameError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvalidUsernameError) InitializeDefaults() {
}

// invalidUsernameErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type invalidUsernameErrorAlias InvalidUsernameError

// UnmarshalJSON implements json.Unmarshaler
func (e *InvalidUsernameError) UnmarshalJSON(data []byte) error {
	var alias invalidUsernameErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvalidUsernameError)(&alias)).InitializeDefaults()
	*e = InvalidUsernameError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvalidUsernameError) MarshalJSON() ([]byte, error) {
	alias := invalidUsernameErrorAlias(e)
	return json.Marshal(alias)
}

// NewInvalidRoleNameError creates a new InvalidRoleNameError
func NewInvalidRoleNameError() *InvalidRoleNameError {
	s := &InvalidRoleNameError{}
	s.InitializeDefaults()
	return s
}

// InvalidRoleNameError struct
type InvalidRoleNameError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InvalidRoleNameError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InvalidRoleNameError]
func (e *InvalidRoleNameError) Is(err error) bool {
	_, ok := err.(*InvalidRoleNameError)
	return ok
}

// IsInvalidRoleNameError indicates whether the given error chain contains an error of type [InvalidRoleNameError]
func IsInvalidRoleNameError(err error) bool {
	return errors.Is(err, &InvalidRoleNameError{})
}

// GetMessage returns the value for the field message
func (e *InvalidRoleNameError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InvalidRoleNameError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InvalidRoleNameError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvalidRoleNameError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvalidRoleNameError) InitializeDefaults() {
}

// invalidRoleNameErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type invalidRoleNameErrorAlias InvalidRoleNameError

// UnmarshalJSON implements json.Unmarshaler
func (e *InvalidRoleNameError) UnmarshalJSON(data []byte) error {
	var alias invalidRoleNameErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvalidRoleNameError)(&alias)).InitializeDefaults()
	*e = InvalidRoleNameError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvalidRoleNameError) MarshalJSON() ([]byte, error) {
	alias := invalidRoleNameErrorAlias(e)
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
	// when the user was last inside this account, either by assuming an identity in
	// it or by signing straight into it. unset if they were never in it.
	// accounts are returned never-entered first (newest membership first), then the
	// rest, most recent first.
	LastEnteredAt *time.Time `json:"lastEnteredAt,omitempty" yaml:"lastEnteredAt,omitempty"`
}

// GetAccountName returns the value for the field accountName
func (e *MemberAccount) GetAccountName() string {
	return e.AccountName
}

// SetAccountName sets the value for the field accountName
func (e *MemberAccount) SetAccountName(accountName string) {
	e.AccountName = accountName
}

// GetLastEnteredAt returns the value for the field lastEnteredAt
func (e *MemberAccount) GetLastEnteredAt() *time.Time {
	return e.LastEnteredAt
}

// SetLastEnteredAt sets the value for the field lastEnteredAt
func (e *MemberAccount) SetLastEnteredAt(lastEnteredAt *time.Time) {
	e.LastEnteredAt = lastEnteredAt
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
	// when the access key was created
	// This value is always set. An absent value indicates a server fault, not a state.
	CreatedAt *time.Time `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
	// when the access key was last used to authenticate, unset if has never been used
	LastActivityAt *time.Time `json:"lastActivityAt,omitempty" yaml:"lastActivityAt,omitempty"`
}

// GetAccessKeyID returns the value for the field accessKeyID
func (e *CredentialInfo) GetAccessKeyID() string {
	return e.AccessKeyID
}

// SetAccessKeyID sets the value for the field accessKeyID
func (e *CredentialInfo) SetAccessKeyID(accessKeyID string) {
	e.AccessKeyID = accessKeyID
}

// GetCreatedAt returns the value for the field createdAt
func (e *CredentialInfo) GetCreatedAt() *time.Time {
	return e.CreatedAt
}

// SetCreatedAt sets the value for the field createdAt
func (e *CredentialInfo) SetCreatedAt(createdAt *time.Time) {
	e.CreatedAt = createdAt
}

// GetLastActivityAt returns the value for the field lastActivityAt
func (e *CredentialInfo) GetLastActivityAt() *time.Time {
	return e.LastActivityAt
}

// SetLastActivityAt sets the value for the field lastActivityAt
func (e *CredentialInfo) SetLastActivityAt(lastActivityAt *time.Time) {
	e.LastActivityAt = lastActivityAt
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
	Builtin bool `json:"builtin,omitempty" yaml:"builtin,omitempty"`
	// when the user-defined policy was created
	CreatedAt        *time.Time `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
	DefinitionHuJSON []byte     `json:"definitionHuJSON,omitempty" yaml:"definitionHuJSON,omitempty"`
	// DRN of this policy, e.g. iam:IdentityPolicy(read-only)
	Drn string `json:"drn,omitempty" yaml:"drn,omitempty"`
	// present only for a service-provisioned policy: read-only to the account
	ManagedBy  *ManagedByService          `json:"managedBy,omitempty" yaml:"managedBy,omitempty"`
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

// GetCreatedAt returns the value for the field createdAt
func (e *IdentityPolicy) GetCreatedAt() *time.Time {
	return e.CreatedAt
}

// SetCreatedAt sets the value for the field createdAt
func (e *IdentityPolicy) SetCreatedAt(createdAt *time.Time) {
	e.CreatedAt = createdAt
}

// GetDefinitionHuJSON returns the value for the field definitionHuJSON
func (e *IdentityPolicy) GetDefinitionHuJSON() []byte {
	return e.DefinitionHuJSON
}

// SetDefinitionHuJSON sets the value for the field definitionHuJSON
func (e *IdentityPolicy) SetDefinitionHuJSON(definitionHuJSON []byte) {
	e.DefinitionHuJSON = definitionHuJSON
}

// GetDrn returns the value for the field drn
func (e *IdentityPolicy) GetDrn() string {
	return e.Drn
}

// SetDrn sets the value for the field drn
func (e *IdentityPolicy) SetDrn(drn string) {
	e.Drn = drn
}

// GetManagedBy returns the value for the field managedBy
func (e *IdentityPolicy) GetManagedBy() *ManagedByService {
	return e.ManagedBy
}

// SetManagedBy sets the value for the field managedBy
func (e *IdentityPolicy) SetManagedBy(managedBy *ManagedByService) {
	e.ManagedBy = managedBy
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
	// when the policy was attached to the target
	// This value is always set. An absent value indicates a server fault, not a state.
	CreatedAt  *time.Time `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
	PolicyName string     `json:"policyName,omitempty" yaml:"policyName,omitempty"`
	TargetDRN  string     `json:"targetDRN,omitempty" yaml:"targetDRN,omitempty"`
}

// GetCreatedAt returns the value for the field createdAt
func (e *IdentityPolicyAttachment) GetCreatedAt() *time.Time {
	return e.CreatedAt
}

// SetCreatedAt sets the value for the field createdAt
func (e *IdentityPolicyAttachment) SetCreatedAt(createdAt *time.Time) {
	e.CreatedAt = createdAt
}

// GetPolicyName returns the value for the field policyName
func (e *IdentityPolicyAttachment) GetPolicyName() string {
	return e.PolicyName
}

// SetPolicyName sets the value for the field policyName
func (e *IdentityPolicyAttachment) SetPolicyName(policyName string) {
	e.PolicyName = policyName
}

// GetTargetDRN returns the value for the field targetDRN
func (e *IdentityPolicyAttachment) GetTargetDRN() string {
	return e.TargetDRN
}

// SetTargetDRN sets the value for the field targetDRN
func (e *IdentityPolicyAttachment) SetTargetDRN(targetDRN string) {
	e.TargetDRN = targetDRN
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

// NewIdentityPolicyAttachmentInfo creates a new IdentityPolicyAttachmentInfo
func NewIdentityPolicyAttachmentInfo() *IdentityPolicyAttachmentInfo {
	s := &IdentityPolicyAttachmentInfo{}
	s.InitializeDefaults()
	return s
}

// IdentityPolicyAttachmentInfo - Describes a policy attached to a user or role, including when it was attached.
type IdentityPolicyAttachmentInfo struct {
	// whether the attached policy is a builtin policy
	Builtin bool `json:"builtin,omitempty" yaml:"builtin,omitempty"`
	// when the policy was attached to the user or role
	// This value is always set. An absent value indicates a server fault, not a state.
	CreatedAt  *time.Time `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
	PolicyName string     `json:"policyName,omitempty" yaml:"policyName,omitempty"`
}

// GetBuiltin returns the value for the field builtin
func (e *IdentityPolicyAttachmentInfo) GetBuiltin() bool {
	return e.Builtin
}

// SetBuiltin sets the value for the field builtin
func (e *IdentityPolicyAttachmentInfo) SetBuiltin(builtin bool) {
	e.Builtin = builtin
}

// GetCreatedAt returns the value for the field createdAt
func (e *IdentityPolicyAttachmentInfo) GetCreatedAt() *time.Time {
	return e.CreatedAt
}

// SetCreatedAt sets the value for the field createdAt
func (e *IdentityPolicyAttachmentInfo) SetCreatedAt(createdAt *time.Time) {
	e.CreatedAt = createdAt
}

// GetPolicyName returns the value for the field policyName
func (e *IdentityPolicyAttachmentInfo) GetPolicyName() string {
	return e.PolicyName
}

// SetPolicyName sets the value for the field policyName
func (e *IdentityPolicyAttachmentInfo) SetPolicyName(policyName string) {
	e.PolicyName = policyName
}

// StructPath returns StructPath
func (e *IdentityPolicyAttachmentInfo) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyAttachmentInfo.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *IdentityPolicyAttachmentInfo) InitializeDefaults() {
}

// identityPolicyAttachmentInfoAlias is defined to help pre and post JSON marshaling without recursive loops
type identityPolicyAttachmentInfoAlias IdentityPolicyAttachmentInfo

// UnmarshalJSON implements json.Unmarshaler
func (e *IdentityPolicyAttachmentInfo) UnmarshalJSON(data []byte) error {
	var alias identityPolicyAttachmentInfoAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*IdentityPolicyAttachmentInfo)(&alias)).InitializeDefaults()
	*e = IdentityPolicyAttachmentInfo(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e IdentityPolicyAttachmentInfo) MarshalJSON() ([]byte, error) {
	alias := identityPolicyAttachmentInfoAlias(e)
	return json.Marshal(alias)
}

// NewUserNotFoundError creates a new UserNotFoundError
func NewUserNotFoundError() *UserNotFoundError {
	s := &UserNotFoundError{}
	s.InitializeDefaults()
	return s
}

// UserNotFoundError struct
type UserNotFoundError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *UserNotFoundError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [UserNotFoundError]
func (e *UserNotFoundError) Is(err error) bool {
	_, ok := err.(*UserNotFoundError)
	return ok
}

// IsUserNotFoundError indicates whether the given error chain contains an error of type [UserNotFoundError]
func IsUserNotFoundError(err error) bool {
	return errors.Is(err, &UserNotFoundError{})
}

// GetMessage returns the value for the field message
func (e *UserNotFoundError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *UserNotFoundError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *UserNotFoundError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserNotFoundError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserNotFoundError) InitializeDefaults() {
}

// userNotFoundErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type userNotFoundErrorAlias UserNotFoundError

// UnmarshalJSON implements json.Unmarshaler
func (e *UserNotFoundError) UnmarshalJSON(data []byte) error {
	var alias userNotFoundErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserNotFoundError)(&alias)).InitializeDefaults()
	*e = UserNotFoundError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserNotFoundError) MarshalJSON() ([]byte, error) {
	alias := userNotFoundErrorAlias(e)
	return json.Marshal(alias)
}

// NewCannotDisableSelfError creates a new CannotDisableSelfError
func NewCannotDisableSelfError() *CannotDisableSelfError {
	s := &CannotDisableSelfError{}
	s.InitializeDefaults()
	return s
}

// CannotDisableSelfError - Occurs when a caller tries to suspend (deactivate) their own user, which would
// lock them out of the account with no way to restore access.
type CannotDisableSelfError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *CannotDisableSelfError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [CannotDisableSelfError]
func (e *CannotDisableSelfError) Is(err error) bool {
	_, ok := err.(*CannotDisableSelfError)
	return ok
}

// IsCannotDisableSelfError indicates whether the given error chain contains an error of type [CannotDisableSelfError]
func IsCannotDisableSelfError(err error) bool {
	return errors.Is(err, &CannotDisableSelfError{})
}

// GetMessage returns the value for the field message
func (e *CannotDisableSelfError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *CannotDisableSelfError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *CannotDisableSelfError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathCannotDisableSelfError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *CannotDisableSelfError) InitializeDefaults() {
}

// cannotDisableSelfErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type cannotDisableSelfErrorAlias CannotDisableSelfError

// UnmarshalJSON implements json.Unmarshaler
func (e *CannotDisableSelfError) UnmarshalJSON(data []byte) error {
	var alias cannotDisableSelfErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*CannotDisableSelfError)(&alias)).InitializeDefaults()
	*e = CannotDisableSelfError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e CannotDisableSelfError) MarshalJSON() ([]byte, error) {
	alias := cannotDisableSelfErrorAlias(e)
	return json.Marshal(alias)
}

// NewRoleNotFoundError creates a new RoleNotFoundError
func NewRoleNotFoundError() *RoleNotFoundError {
	s := &RoleNotFoundError{}
	s.InitializeDefaults()
	return s
}

// RoleNotFoundError struct
type RoleNotFoundError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *RoleNotFoundError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [RoleNotFoundError]
func (e *RoleNotFoundError) Is(err error) bool {
	_, ok := err.(*RoleNotFoundError)
	return ok
}

// IsRoleNotFoundError indicates whether the given error chain contains an error of type [RoleNotFoundError]
func IsRoleNotFoundError(err error) bool {
	return errors.Is(err, &RoleNotFoundError{})
}

// GetMessage returns the value for the field message
func (e *RoleNotFoundError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *RoleNotFoundError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *RoleNotFoundError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleNotFoundError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleNotFoundError) InitializeDefaults() {
}

// roleNotFoundErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type roleNotFoundErrorAlias RoleNotFoundError

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleNotFoundError) UnmarshalJSON(data []byte) error {
	var alias roleNotFoundErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleNotFoundError)(&alias)).InitializeDefaults()
	*e = RoleNotFoundError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleNotFoundError) MarshalJSON() ([]byte, error) {
	alias := roleNotFoundErrorAlias(e)
	return json.Marshal(alias)
}

// NewGroupNotFoundError creates a new GroupNotFoundError
func NewGroupNotFoundError() *GroupNotFoundError {
	s := &GroupNotFoundError{}
	s.InitializeDefaults()
	return s
}

// GroupNotFoundError struct
type GroupNotFoundError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *GroupNotFoundError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [GroupNotFoundError]
func (e *GroupNotFoundError) Is(err error) bool {
	_, ok := err.(*GroupNotFoundError)
	return ok
}

// IsGroupNotFoundError indicates whether the given error chain contains an error of type [GroupNotFoundError]
func IsGroupNotFoundError(err error) bool {
	return errors.Is(err, &GroupNotFoundError{})
}

// GetMessage returns the value for the field message
func (e *GroupNotFoundError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *GroupNotFoundError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *GroupNotFoundError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathGroupNotFoundError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *GroupNotFoundError) InitializeDefaults() {
}

// groupNotFoundErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type groupNotFoundErrorAlias GroupNotFoundError

// UnmarshalJSON implements json.Unmarshaler
func (e *GroupNotFoundError) UnmarshalJSON(data []byte) error {
	var alias groupNotFoundErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*GroupNotFoundError)(&alias)).InitializeDefaults()
	*e = GroupNotFoundError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e GroupNotFoundError) MarshalJSON() ([]byte, error) {
	alias := groupNotFoundErrorAlias(e)
	return json.Marshal(alias)
}

// NewInvalidGroupNameError creates a new InvalidGroupNameError
func NewInvalidGroupNameError() *InvalidGroupNameError {
	s := &InvalidGroupNameError{}
	s.InitializeDefaults()
	return s
}

// InvalidGroupNameError struct
type InvalidGroupNameError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InvalidGroupNameError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InvalidGroupNameError]
func (e *InvalidGroupNameError) Is(err error) bool {
	_, ok := err.(*InvalidGroupNameError)
	return ok
}

// IsInvalidGroupNameError indicates whether the given error chain contains an error of type [InvalidGroupNameError]
func IsInvalidGroupNameError(err error) bool {
	return errors.Is(err, &InvalidGroupNameError{})
}

// GetMessage returns the value for the field message
func (e *InvalidGroupNameError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InvalidGroupNameError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InvalidGroupNameError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvalidGroupNameError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvalidGroupNameError) InitializeDefaults() {
}

// invalidGroupNameErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type invalidGroupNameErrorAlias InvalidGroupNameError

// UnmarshalJSON implements json.Unmarshaler
func (e *InvalidGroupNameError) UnmarshalJSON(data []byte) error {
	var alias invalidGroupNameErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvalidGroupNameError)(&alias)).InitializeDefaults()
	*e = InvalidGroupNameError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvalidGroupNameError) MarshalJSON() ([]byte, error) {
	alias := invalidGroupNameErrorAlias(e)
	return json.Marshal(alias)
}

// NewInvalidGroupFilterError creates a new InvalidGroupFilterError
func NewInvalidGroupFilterError() *InvalidGroupFilterError {
	s := &InvalidGroupFilterError{}
	s.InitializeDefaults()
	return s
}

// InvalidGroupFilterError - Occurs when Group.List is asked for a filter combination no group can satisfy,
// such as managed false together with a managedByService.
type InvalidGroupFilterError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InvalidGroupFilterError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InvalidGroupFilterError]
func (e *InvalidGroupFilterError) Is(err error) bool {
	_, ok := err.(*InvalidGroupFilterError)
	return ok
}

// IsInvalidGroupFilterError indicates whether the given error chain contains an error of type [InvalidGroupFilterError]
func IsInvalidGroupFilterError(err error) bool {
	return errors.Is(err, &InvalidGroupFilterError{})
}

// GetMessage returns the value for the field message
func (e *InvalidGroupFilterError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InvalidGroupFilterError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InvalidGroupFilterError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvalidGroupFilterError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvalidGroupFilterError) InitializeDefaults() {
}

// invalidGroupFilterErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type invalidGroupFilterErrorAlias InvalidGroupFilterError

// UnmarshalJSON implements json.Unmarshaler
func (e *InvalidGroupFilterError) UnmarshalJSON(data []byte) error {
	var alias invalidGroupFilterErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvalidGroupFilterError)(&alias)).InitializeDefaults()
	*e = InvalidGroupFilterError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvalidGroupFilterError) MarshalJSON() ([]byte, error) {
	alias := invalidGroupFilterErrorAlias(e)
	return json.Marshal(alias)
}

// NewInvalidUserFilterError creates a new InvalidUserFilterError
func NewInvalidUserFilterError() *InvalidUserFilterError {
	s := &InvalidUserFilterError{}
	s.InitializeDefaults()
	return s
}

// InvalidUserFilterError - Occurs when User.List is asked for a filter combination no user can satisfy,
// such as managed false together with a managedByService.
type InvalidUserFilterError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InvalidUserFilterError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InvalidUserFilterError]
func (e *InvalidUserFilterError) Is(err error) bool {
	_, ok := err.(*InvalidUserFilterError)
	return ok
}

// IsInvalidUserFilterError indicates whether the given error chain contains an error of type [InvalidUserFilterError]
func IsInvalidUserFilterError(err error) bool {
	return errors.Is(err, &InvalidUserFilterError{})
}

// GetMessage returns the value for the field message
func (e *InvalidUserFilterError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InvalidUserFilterError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InvalidUserFilterError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvalidUserFilterError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvalidUserFilterError) InitializeDefaults() {
}

// invalidUserFilterErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type invalidUserFilterErrorAlias InvalidUserFilterError

// UnmarshalJSON implements json.Unmarshaler
func (e *InvalidUserFilterError) UnmarshalJSON(data []byte) error {
	var alias invalidUserFilterErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvalidUserFilterError)(&alias)).InitializeDefaults()
	*e = InvalidUserFilterError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvalidUserFilterError) MarshalJSON() ([]byte, error) {
	alias := invalidUserFilterErrorAlias(e)
	return json.Marshal(alias)
}

// NewIdentityInUseError creates a new IdentityInUseError
func NewIdentityInUseError() *IdentityInUseError {
	s := &IdentityInUseError{}
	s.InitializeDefaults()
	return s
}

// IdentityInUseError - Occurs when an identity is in use and cannot be deleted because it is attached to a resource
type IdentityInUseError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *IdentityInUseError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [IdentityInUseError]
func (e *IdentityInUseError) Is(err error) bool {
	_, ok := err.(*IdentityInUseError)
	return ok
}

// IsIdentityInUseError indicates whether the given error chain contains an error of type [IdentityInUseError]
func IsIdentityInUseError(err error) bool {
	return errors.Is(err, &IdentityInUseError{})
}

// GetMessage returns the value for the field message
func (e *IdentityInUseError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *IdentityInUseError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *IdentityInUseError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityInUseError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *IdentityInUseError) InitializeDefaults() {
}

// identityInUseErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type identityInUseErrorAlias IdentityInUseError

// UnmarshalJSON implements json.Unmarshaler
func (e *IdentityInUseError) UnmarshalJSON(data []byte) error {
	var alias identityInUseErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*IdentityInUseError)(&alias)).InitializeDefaults()
	*e = IdentityInUseError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e IdentityInUseError) MarshalJSON() ([]byte, error) {
	alias := identityInUseErrorAlias(e)
	return json.Marshal(alias)
}

// NewUserSSHKey creates a new UserSSHKey
func NewUserSSHKey() *UserSSHKey {
	s := &UserSSHKey{}
	s.InitializeDefaults()
	return s
}

// UserSSHKey - One public key carried by one person.
type UserSSHKey struct {
	// the comment OpenSSH already carries at the end of the line, e.g. "johan@laptop". It is a
	// LABEL a person reads and NEVER an identifier: two keys may carry the same one
	Comment   *string    `json:"comment,omitempty" yaml:"comment,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
	// the identity, derived from publicKey and unable to drift from it
	Fingerprint string `json:"fingerprint,omitempty" yaml:"fingerprint,omitempty"`
	// the public key in OpenSSH authorized_keys form
	PublicKey string `json:"publicKey,omitempty" yaml:"publicKey,omitempty"`
	// what this key is called, at most 160 characters. Never empty: with nothing stored it is
	// the OpenSSH comment, and with no comment either it is the fingerprint
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
}

// GetComment returns the value for the field comment
func (e *UserSSHKey) GetComment() *string {
	return e.Comment
}

// SetComment sets the value for the field comment
func (e *UserSSHKey) SetComment(comment *string) {
	e.Comment = comment
}

// GetCreatedAt returns the value for the field createdAt
func (e *UserSSHKey) GetCreatedAt() *time.Time {
	return e.CreatedAt
}

// SetCreatedAt sets the value for the field createdAt
func (e *UserSSHKey) SetCreatedAt(createdAt *time.Time) {
	e.CreatedAt = createdAt
}

// GetFingerprint returns the value for the field fingerprint
func (e *UserSSHKey) GetFingerprint() string {
	return e.Fingerprint
}

// SetFingerprint sets the value for the field fingerprint
func (e *UserSSHKey) SetFingerprint(fingerprint string) {
	e.Fingerprint = fingerprint
}

// GetPublicKey returns the value for the field publicKey
func (e *UserSSHKey) GetPublicKey() string {
	return e.PublicKey
}

// SetPublicKey sets the value for the field publicKey
func (e *UserSSHKey) SetPublicKey(publicKey string) {
	e.PublicKey = publicKey
}

// GetTitle returns the value for the field title
func (e *UserSSHKey) GetTitle() string {
	return e.Title
}

// SetTitle sets the value for the field title
func (e *UserSSHKey) SetTitle(title string) {
	e.Title = title
}

// StructPath returns StructPath
func (e *UserSSHKey) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserSSHKey.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserSSHKey) InitializeDefaults() {
}

// userSSHKeyAlias is defined to help pre and post JSON marshaling without recursive loops
type userSSHKeyAlias UserSSHKey

// UnmarshalJSON implements json.Unmarshaler
func (e *UserSSHKey) UnmarshalJSON(data []byte) error {
	var alias userSSHKeyAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserSSHKey)(&alias)).InitializeDefaults()
	*e = UserSSHKey(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserSSHKey) MarshalJSON() ([]byte, error) {
	alias := userSSHKeyAlias(e)
	return json.Marshal(alias)
}

// NewInvalidSSHPublicKeyError creates a new InvalidSSHPublicKeyError
func NewInvalidSSHPublicKeyError() *InvalidSSHPublicKeyError {
	s := &InvalidSSHPublicKeyError{}
	s.InitializeDefaults()
	return s
}

// InvalidSSHPublicKeyError - Occurs when the submitted bytes are not a public key this service will store. The message names what is wrong with the key.
type InvalidSSHPublicKeyError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InvalidSSHPublicKeyError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InvalidSSHPublicKeyError]
func (e *InvalidSSHPublicKeyError) Is(err error) bool {
	_, ok := err.(*InvalidSSHPublicKeyError)
	return ok
}

// IsInvalidSSHPublicKeyError indicates whether the given error chain contains an error of type [InvalidSSHPublicKeyError]
func IsInvalidSSHPublicKeyError(err error) bool {
	return errors.Is(err, &InvalidSSHPublicKeyError{})
}

// GetMessage returns the value for the field message
func (e *InvalidSSHPublicKeyError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InvalidSSHPublicKeyError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InvalidSSHPublicKeyError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvalidSSHPublicKeyError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvalidSSHPublicKeyError) InitializeDefaults() {
}

// invalidSSHPublicKeyErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type invalidSSHPublicKeyErrorAlias InvalidSSHPublicKeyError

// UnmarshalJSON implements json.Unmarshaler
func (e *InvalidSSHPublicKeyError) UnmarshalJSON(data []byte) error {
	var alias invalidSSHPublicKeyErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvalidSSHPublicKeyError)(&alias)).InitializeDefaults()
	*e = InvalidSSHPublicKeyError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvalidSSHPublicKeyError) MarshalJSON() ([]byte, error) {
	alias := invalidSSHPublicKeyErrorAlias(e)
	return json.Marshal(alias)
}

// NewSSHKeyAlreadyExistsError creates a new SSHKeyAlreadyExistsError
func NewSSHKeyAlreadyExistsError() *SSHKeyAlreadyExistsError {
	s := &SSHKeyAlreadyExistsError{}
	s.InitializeDefaults()
	return s
}

// SSHKeyAlreadyExistsError - Occurs when this public key is already registered, whether on a user in this account, including the caller's own, or in another account entirely.
// UNIQUENESS IS PLATFORM-WIDE AND NOT PER USER, AND IT IS A SECURITY PROPERTY RATHER THAN TIDINESS. The authorized_keys marker carries the USER, so one key on two users writes two lines with the same bytes and different owners. Revoking one leaves the other, and whoever holds that private key still has a shell on every machine the second user can reach. An administrator who removed somebody would believe they had.
type SSHKeyAlreadyExistsError struct {
	Fingerprint string `json:"fingerprint,omitempty" yaml:"fingerprint,omitempty"`
	Message     string `json:"message,omitempty" yaml:"message,omitempty"`
	// the user who already holds it, so the refusal is actionable rather than only true. ABSENT when another account registered the key, because naming a person the caller cannot see would disclose more than it helps
	Username *string `json:"username,omitempty" yaml:"username,omitempty"`
}

// Error implements the error interface
func (e *SSHKeyAlreadyExistsError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [SSHKeyAlreadyExistsError]
func (e *SSHKeyAlreadyExistsError) Is(err error) bool {
	_, ok := err.(*SSHKeyAlreadyExistsError)
	return ok
}

// IsSSHKeyAlreadyExistsError indicates whether the given error chain contains an error of type [SSHKeyAlreadyExistsError]
func IsSSHKeyAlreadyExistsError(err error) bool {
	return errors.Is(err, &SSHKeyAlreadyExistsError{})
}

// GetFingerprint returns the value for the field fingerprint
func (e *SSHKeyAlreadyExistsError) GetFingerprint() string {
	return e.Fingerprint
}

// SetFingerprint sets the value for the field fingerprint
func (e *SSHKeyAlreadyExistsError) SetFingerprint(fingerprint string) {
	e.Fingerprint = fingerprint
}

// GetMessage returns the value for the field message
func (e *SSHKeyAlreadyExistsError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *SSHKeyAlreadyExistsError) SetMessage(message string) {
	e.Message = message
}

// GetUsername returns the value for the field username
func (e *SSHKeyAlreadyExistsError) GetUsername() *string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *SSHKeyAlreadyExistsError) SetUsername(username *string) {
	e.Username = username
}

// StructPath returns StructPath
func (e *SSHKeyAlreadyExistsError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathSSHKeyAlreadyExistsError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *SSHKeyAlreadyExistsError) InitializeDefaults() {
}

// sSHKeyAlreadyExistsErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type sSHKeyAlreadyExistsErrorAlias SSHKeyAlreadyExistsError

// UnmarshalJSON implements json.Unmarshaler
func (e *SSHKeyAlreadyExistsError) UnmarshalJSON(data []byte) error {
	var alias sSHKeyAlreadyExistsErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*SSHKeyAlreadyExistsError)(&alias)).InitializeDefaults()
	*e = SSHKeyAlreadyExistsError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e SSHKeyAlreadyExistsError) MarshalJSON() ([]byte, error) {
	alias := sSHKeyAlreadyExistsErrorAlias(e)
	return json.Marshal(alias)
}

// NewInvalidSSHKeyTitleError creates a new InvalidSSHKeyTitleError
func NewInvalidSSHKeyTitleError() *InvalidSSHKeyTitleError {
	s := &InvalidSSHKeyTitleError{}
	s.InitializeDefaults()
	return s
}

// InvalidSSHKeyTitleError - Occurs when the submitted title is not one this service will store. The message names what is wrong with it.
type InvalidSSHKeyTitleError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InvalidSSHKeyTitleError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InvalidSSHKeyTitleError]
func (e *InvalidSSHKeyTitleError) Is(err error) bool {
	_, ok := err.(*InvalidSSHKeyTitleError)
	return ok
}

// IsInvalidSSHKeyTitleError indicates whether the given error chain contains an error of type [InvalidSSHKeyTitleError]
func IsInvalidSSHKeyTitleError(err error) bool {
	return errors.Is(err, &InvalidSSHKeyTitleError{})
}

// GetMessage returns the value for the field message
func (e *InvalidSSHKeyTitleError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InvalidSSHKeyTitleError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InvalidSSHKeyTitleError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvalidSSHKeyTitleError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvalidSSHKeyTitleError) InitializeDefaults() {
}

// invalidSSHKeyTitleErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type invalidSSHKeyTitleErrorAlias InvalidSSHKeyTitleError

// UnmarshalJSON implements json.Unmarshaler
func (e *InvalidSSHKeyTitleError) UnmarshalJSON(data []byte) error {
	var alias invalidSSHKeyTitleErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvalidSSHKeyTitleError)(&alias)).InitializeDefaults()
	*e = InvalidSSHKeyTitleError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvalidSSHKeyTitleError) MarshalJSON() ([]byte, error) {
	alias := invalidSSHKeyTitleErrorAlias(e)
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
	// optional hard cap on the lifetime of any credential assumed for this
	// user, in seconds. Bounds total life regardless of activity/keep-alive.
	// Omit to leave it unconfigured, in which case the cap is the one applied by the path that mints the credential (see Account.AssumeIdentity and Account.SSO.BeginAuthentication).
	MaxSessionLifetimeSeconds *int32 `json:"maxSessionLifetimeSeconds,omitempty" yaml:"maxSessionLifetimeSeconds,omitempty"`
	Username                  string `json:"username,omitempty" yaml:"username,omitempty"`
}

// GetDescription returns the value for the field description
func (e *UserCreateInput) GetDescription() string {
	return e.Description
}

// SetDescription sets the value for the field description
func (e *UserCreateInput) SetDescription(description string) {
	e.Description = description
}

// GetMaxSessionLifetimeSeconds returns the value for the field maxSessionLifetimeSeconds
func (e *UserCreateInput) GetMaxSessionLifetimeSeconds() *int32 {
	return e.MaxSessionLifetimeSeconds
}

// SetMaxSessionLifetimeSeconds sets the value for the field maxSessionLifetimeSeconds
func (e *UserCreateInput) SetMaxSessionLifetimeSeconds(maxSessionLifetimeSeconds *int32) {
	e.MaxSessionLifetimeSeconds = maxSessionLifetimeSeconds
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

// NewUserGetUserNotAvailableError creates a new UserGetUserNotAvailableError
func NewUserGetUserNotAvailableError() *UserGetUserNotAvailableError {
	s := &UserGetUserNotAvailableError{}
	s.InitializeDefaults()
	return s
}

// UserGetUserNotAvailableError struct
type UserGetUserNotAvailableError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *UserGetUserNotAvailableError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [UserGetUserNotAvailableError]
func (e *UserGetUserNotAvailableError) Is(err error) bool {
	_, ok := err.(*UserGetUserNotAvailableError)
	return ok
}

// IsUserGetUserNotAvailableError indicates whether the given error chain contains an error of type [UserGetUserNotAvailableError]
func IsUserGetUserNotAvailableError(err error) bool {
	return errors.Is(err, &UserGetUserNotAvailableError{})
}

// GetMessage returns the value for the field message
func (e *UserGetUserNotAvailableError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *UserGetUserNotAvailableError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *UserGetUserNotAvailableError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserGetUserNotAvailableError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserGetUserNotAvailableError) InitializeDefaults() {
}

// userGetUserNotAvailableErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type userGetUserNotAvailableErrorAlias UserGetUserNotAvailableError

// UnmarshalJSON implements json.Unmarshaler
func (e *UserGetUserNotAvailableError) UnmarshalJSON(data []byte) error {
	var alias userGetUserNotAvailableErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserGetUserNotAvailableError)(&alias)).InitializeDefaults()
	*e = UserGetUserNotAvailableError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserGetUserNotAvailableError) MarshalJSON() ([]byte, error) {
	alias := userGetUserNotAvailableErrorAlias(e)
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

// NewUserSetActiveInput creates a new UserSetActiveInput
func NewUserSetActiveInput() *UserSetActiveInput {
	s := &UserSetActiveInput{}
	s.InitializeDefaults()
	return s
}

// UserSetActiveInput struct
type UserSetActiveInput struct {
	Active   bool   `json:"active,omitempty" yaml:"active,omitempty"`
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}

// GetActive returns the value for the field active
func (e *UserSetActiveInput) GetActive() bool {
	return e.Active
}

// SetActive sets the value for the field active
func (e *UserSetActiveInput) SetActive(active bool) {
	e.Active = active
}

// GetUsername returns the value for the field username
func (e *UserSetActiveInput) GetUsername() string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *UserSetActiveInput) SetUsername(username string) {
	e.Username = username
}

// StructPath returns StructPath
func (e *UserSetActiveInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserSetActiveInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserSetActiveInput) InitializeDefaults() {
}

// userSetActiveInputAlias is defined to help pre and post JSON marshaling without recursive loops
type userSetActiveInputAlias UserSetActiveInput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserSetActiveInput) UnmarshalJSON(data []byte) error {
	var alias userSetActiveInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserSetActiveInput)(&alias)).InitializeDefaults()
	*e = UserSetActiveInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserSetActiveInput) MarshalJSON() ([]byte, error) {
	alias := userSetActiveInputAlias(e)
	return json.Marshal(alias)
}

// NewUserSetActiveOutput creates a new UserSetActiveOutput
func NewUserSetActiveOutput() *UserSetActiveOutput {
	s := &UserSetActiveOutput{}
	s.InitializeDefaults()
	return s
}

// UserSetActiveOutput struct
type UserSetActiveOutput struct {
	User *UserInformation `json:"user,omitempty" yaml:"user,omitempty"`
}

// GetUser returns the value for the field user
func (e *UserSetActiveOutput) GetUser() *UserInformation {
	return e.User
}

// SetUser sets the value for the field user
func (e *UserSetActiveOutput) SetUser(user *UserInformation) {
	e.User = user
}

// StructPath returns StructPath
func (e *UserSetActiveOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserSetActiveOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserSetActiveOutput) InitializeDefaults() {
}

// userSetActiveOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type userSetActiveOutputAlias UserSetActiveOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserSetActiveOutput) UnmarshalJSON(data []byte) error {
	var alias userSetActiveOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserSetActiveOutput)(&alias)).InitializeDefaults()
	*e = UserSetActiveOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserSetActiveOutput) MarshalJSON() ([]byte, error) {
	alias := userSetActiveOutputAlias(e)
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
	// opaque page cursor: a previous response's nextCursor, verbatim
	Cursor *string `json:"cursor,omitempty" yaml:"cursor,omitempty"`
	// page size, 1 to 200; 50 when absent
	Limit *int32 `json:"limit,omitempty" yaml:"limit,omitempty"`
	// when true, only service-managed users; when false, only users the account
	// authored. Absent returns both, which is today's behaviour and keeps every
	// existing caller correct
	Managed *bool `json:"managed,omitempty" yaml:"managed,omitempty"`
	// the managing service principal FQDN, e.g. "uplink.deployport.io", to
	// return only that service's users. Absent applies no service filter
	ManagedByService *string `json:"managedByService,omitempty" yaml:"managedByService,omitempty"`
}

// GetCursor returns the value for the field cursor
func (e *UserListInput) GetCursor() *string {
	return e.Cursor
}

// SetCursor sets the value for the field cursor
func (e *UserListInput) SetCursor(cursor *string) {
	e.Cursor = cursor
}

// GetLimit returns the value for the field limit
func (e *UserListInput) GetLimit() *int32 {
	return e.Limit
}

// SetLimit sets the value for the field limit
func (e *UserListInput) SetLimit(limit *int32) {
	e.Limit = limit
}

// GetManaged returns the value for the field managed
func (e *UserListInput) GetManaged() *bool {
	return e.Managed
}

// SetManaged sets the value for the field managed
func (e *UserListInput) SetManaged(managed *bool) {
	e.Managed = managed
}

// GetManagedByService returns the value for the field managedByService
func (e *UserListInput) GetManagedByService() *string {
	return e.ManagedByService
}

// SetManagedByService sets the value for the field managedByService
func (e *UserListInput) SetManagedByService(managedByService *string) {
	e.ManagedByService = managedByService
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
	// pass to the next call's cursor; absent on the last page
	NextCursor *string            `json:"nextCursor,omitempty" yaml:"nextCursor,omitempty"`
	Users      []*UserInformation `json:"users,omitempty" yaml:"users,omitempty"`
}

// GetNextCursor returns the value for the field nextCursor
func (e *UserListOutput) GetNextCursor() *string {
	return e.NextCursor
}

// SetNextCursor sets the value for the field nextCursor
func (e *UserListOutput) SetNextCursor(nextCursor *string) {
	e.NextCursor = nextCursor
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
	Attachments []*IdentityPolicyAttachmentInfo `json:"attachments,omitempty" yaml:"attachments,omitempty"`
}

// GetAttachments returns the value for the field attachments
func (e *UserIdentityPolicyListOutput) GetAttachments() []*IdentityPolicyAttachmentInfo {
	return e.Attachments
}

// SetAttachments sets the value for the field attachments
func (e *UserIdentityPolicyListOutput) SetAttachments(attachments []*IdentityPolicyAttachmentInfo) {
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

// NewUserSSHKeyCreateInput creates a new UserSSHKeyCreateInput
func NewUserSSHKeyCreateInput() *UserSSHKeyCreateInput {
	s := &UserSSHKeyCreateInput{}
	s.InitializeDefaults()
	return s
}

// UserSSHKeyCreateInput struct
type UserSSHKeyCreateInput struct {
	// the public key, in OpenSSH authorized_keys form
	PublicKey string `json:"publicKey,omitempty" yaml:"publicKey,omitempty"`
	// what to call this key, at most 160 characters. Absent derives it from the
	// public key's own comment
	Title *string `json:"title,omitempty" yaml:"title,omitempty"`
	// whose key this is. ABSENT MEANS THE CALLING USER, which is the idiom
	// AccessKey.List already established. Naming somebody else is how an
	// administrator seeds a key at onboarding, or for a person whose SSO is broken
	Username *string `json:"username,omitempty" yaml:"username,omitempty"`
}

// GetPublicKey returns the value for the field publicKey
func (e *UserSSHKeyCreateInput) GetPublicKey() string {
	return e.PublicKey
}

// SetPublicKey sets the value for the field publicKey
func (e *UserSSHKeyCreateInput) SetPublicKey(publicKey string) {
	e.PublicKey = publicKey
}

// GetTitle returns the value for the field title
func (e *UserSSHKeyCreateInput) GetTitle() *string {
	return e.Title
}

// SetTitle sets the value for the field title
func (e *UserSSHKeyCreateInput) SetTitle(title *string) {
	e.Title = title
}

// GetUsername returns the value for the field username
func (e *UserSSHKeyCreateInput) GetUsername() *string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *UserSSHKeyCreateInput) SetUsername(username *string) {
	e.Username = username
}

// StructPath returns StructPath
func (e *UserSSHKeyCreateInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserSSHKeyCreateInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserSSHKeyCreateInput) InitializeDefaults() {
}

// userSSHKeyCreateInputAlias is defined to help pre and post JSON marshaling without recursive loops
type userSSHKeyCreateInputAlias UserSSHKeyCreateInput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserSSHKeyCreateInput) UnmarshalJSON(data []byte) error {
	var alias userSSHKeyCreateInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserSSHKeyCreateInput)(&alias)).InitializeDefaults()
	*e = UserSSHKeyCreateInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserSSHKeyCreateInput) MarshalJSON() ([]byte, error) {
	alias := userSSHKeyCreateInputAlias(e)
	return json.Marshal(alias)
}

// NewUserSSHKeyCreateOutput creates a new UserSSHKeyCreateOutput
func NewUserSSHKeyCreateOutput() *UserSSHKeyCreateOutput {
	s := &UserSSHKeyCreateOutput{}
	s.InitializeDefaults()
	return s
}

// UserSSHKeyCreateOutput struct
type UserSSHKeyCreateOutput struct {
	SshKey *UserSSHKey `json:"sshKey,omitempty" yaml:"sshKey,omitempty"`
}

// GetSshKey returns the value for the field sshKey
func (e *UserSSHKeyCreateOutput) GetSshKey() *UserSSHKey {
	return e.SshKey
}

// SetSshKey sets the value for the field sshKey
func (e *UserSSHKeyCreateOutput) SetSshKey(sshKey *UserSSHKey) {
	e.SshKey = sshKey
}

// StructPath returns StructPath
func (e *UserSSHKeyCreateOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserSSHKeyCreateOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserSSHKeyCreateOutput) InitializeDefaults() {
}

// userSSHKeyCreateOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type userSSHKeyCreateOutputAlias UserSSHKeyCreateOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserSSHKeyCreateOutput) UnmarshalJSON(data []byte) error {
	var alias userSSHKeyCreateOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserSSHKeyCreateOutput)(&alias)).InitializeDefaults()
	*e = UserSSHKeyCreateOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserSSHKeyCreateOutput) MarshalJSON() ([]byte, error) {
	alias := userSSHKeyCreateOutputAlias(e)
	return json.Marshal(alias)
}

// NewUserSSHKeyListInput creates a new UserSSHKeyListInput
func NewUserSSHKeyListInput() *UserSSHKeyListInput {
	s := &UserSSHKeyListInput{}
	s.InitializeDefaults()
	return s
}

// UserSSHKeyListInput struct
type UserSSHKeyListInput struct {
	Username *string `json:"username,omitempty" yaml:"username,omitempty"`
}

// GetUsername returns the value for the field username
func (e *UserSSHKeyListInput) GetUsername() *string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *UserSSHKeyListInput) SetUsername(username *string) {
	e.Username = username
}

// StructPath returns StructPath
func (e *UserSSHKeyListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserSSHKeyListInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserSSHKeyListInput) InitializeDefaults() {
}

// userSSHKeyListInputAlias is defined to help pre and post JSON marshaling without recursive loops
type userSSHKeyListInputAlias UserSSHKeyListInput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserSSHKeyListInput) UnmarshalJSON(data []byte) error {
	var alias userSSHKeyListInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserSSHKeyListInput)(&alias)).InitializeDefaults()
	*e = UserSSHKeyListInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserSSHKeyListInput) MarshalJSON() ([]byte, error) {
	alias := userSSHKeyListInputAlias(e)
	return json.Marshal(alias)
}

// NewUserSSHKeyListOutput creates a new UserSSHKeyListOutput
func NewUserSSHKeyListOutput() *UserSSHKeyListOutput {
	s := &UserSSHKeyListOutput{}
	s.InitializeDefaults()
	return s
}

// UserSSHKeyListOutput struct
type UserSSHKeyListOutput struct {
	SshKeys []*UserSSHKey `json:"sshKeys,omitempty" yaml:"sshKeys,omitempty"`
}

// GetSshKeys returns the value for the field sshKeys
func (e *UserSSHKeyListOutput) GetSshKeys() []*UserSSHKey {
	return e.SshKeys
}

// SetSshKeys sets the value for the field sshKeys
func (e *UserSSHKeyListOutput) SetSshKeys(sshKeys []*UserSSHKey) {
	e.SshKeys = sshKeys
}

// StructPath returns StructPath
func (e *UserSSHKeyListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserSSHKeyListOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserSSHKeyListOutput) InitializeDefaults() {
}

// userSSHKeyListOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type userSSHKeyListOutputAlias UserSSHKeyListOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserSSHKeyListOutput) UnmarshalJSON(data []byte) error {
	var alias userSSHKeyListOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserSSHKeyListOutput)(&alias)).InitializeDefaults()
	*e = UserSSHKeyListOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserSSHKeyListOutput) MarshalJSON() ([]byte, error) {
	alias := userSSHKeyListOutputAlias(e)
	return json.Marshal(alias)
}

// NewUserSSHKeyDestroyInput creates a new UserSSHKeyDestroyInput
func NewUserSSHKeyDestroyInput() *UserSSHKeyDestroyInput {
	s := &UserSSHKeyDestroyInput{}
	s.InitializeDefaults()
	return s
}

// UserSSHKeyDestroyInput struct
type UserSSHKeyDestroyInput struct {
	Fingerprint string `json:"fingerprint,omitempty" yaml:"fingerprint,omitempty"`
	// absent means the calling user, as on Create and List
	Username *string `json:"username,omitempty" yaml:"username,omitempty"`
}

// GetFingerprint returns the value for the field fingerprint
func (e *UserSSHKeyDestroyInput) GetFingerprint() string {
	return e.Fingerprint
}

// SetFingerprint sets the value for the field fingerprint
func (e *UserSSHKeyDestroyInput) SetFingerprint(fingerprint string) {
	e.Fingerprint = fingerprint
}

// GetUsername returns the value for the field username
func (e *UserSSHKeyDestroyInput) GetUsername() *string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *UserSSHKeyDestroyInput) SetUsername(username *string) {
	e.Username = username
}

// StructPath returns StructPath
func (e *UserSSHKeyDestroyInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserSSHKeyDestroyInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserSSHKeyDestroyInput) InitializeDefaults() {
}

// userSSHKeyDestroyInputAlias is defined to help pre and post JSON marshaling without recursive loops
type userSSHKeyDestroyInputAlias UserSSHKeyDestroyInput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserSSHKeyDestroyInput) UnmarshalJSON(data []byte) error {
	var alias userSSHKeyDestroyInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserSSHKeyDestroyInput)(&alias)).InitializeDefaults()
	*e = UserSSHKeyDestroyInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserSSHKeyDestroyInput) MarshalJSON() ([]byte, error) {
	alias := userSSHKeyDestroyInputAlias(e)
	return json.Marshal(alias)
}

// NewUserSSHKeyDestroyOutput creates a new UserSSHKeyDestroyOutput
func NewUserSSHKeyDestroyOutput() *UserSSHKeyDestroyOutput {
	s := &UserSSHKeyDestroyOutput{}
	s.InitializeDefaults()
	return s
}

// UserSSHKeyDestroyOutput struct
type UserSSHKeyDestroyOutput struct {
}

// StructPath returns StructPath
func (e *UserSSHKeyDestroyOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserSSHKeyDestroyOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserSSHKeyDestroyOutput) InitializeDefaults() {
}

// userSSHKeyDestroyOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type userSSHKeyDestroyOutputAlias UserSSHKeyDestroyOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserSSHKeyDestroyOutput) UnmarshalJSON(data []byte) error {
	var alias userSSHKeyDestroyOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserSSHKeyDestroyOutput)(&alias)).InitializeDefaults()
	*e = UserSSHKeyDestroyOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserSSHKeyDestroyOutput) MarshalJSON() ([]byte, error) {
	alias := userSSHKeyDestroyOutputAlias(e)
	return json.Marshal(alias)
}

// NewUserSSHKeySetTitleInput creates a new UserSSHKeySetTitleInput
func NewUserSSHKeySetTitleInput() *UserSSHKeySetTitleInput {
	s := &UserSSHKeySetTitleInput{}
	s.InitializeDefaults()
	return s
}

// UserSSHKeySetTitleInput struct
type UserSSHKeySetTitleInput struct {
	Fingerprint string `json:"fingerprint,omitempty" yaml:"fingerprint,omitempty"`
	// absent clears it, so the title falls back to the public key's own comment
	Title *string `json:"title,omitempty" yaml:"title,omitempty"`
	// absent means the calling user, as on Create, List and Destroy
	Username *string `json:"username,omitempty" yaml:"username,omitempty"`
}

// GetFingerprint returns the value for the field fingerprint
func (e *UserSSHKeySetTitleInput) GetFingerprint() string {
	return e.Fingerprint
}

// SetFingerprint sets the value for the field fingerprint
func (e *UserSSHKeySetTitleInput) SetFingerprint(fingerprint string) {
	e.Fingerprint = fingerprint
}

// GetTitle returns the value for the field title
func (e *UserSSHKeySetTitleInput) GetTitle() *string {
	return e.Title
}

// SetTitle sets the value for the field title
func (e *UserSSHKeySetTitleInput) SetTitle(title *string) {
	e.Title = title
}

// GetUsername returns the value for the field username
func (e *UserSSHKeySetTitleInput) GetUsername() *string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *UserSSHKeySetTitleInput) SetUsername(username *string) {
	e.Username = username
}

// StructPath returns StructPath
func (e *UserSSHKeySetTitleInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserSSHKeySetTitleInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserSSHKeySetTitleInput) InitializeDefaults() {
}

// userSSHKeySetTitleInputAlias is defined to help pre and post JSON marshaling without recursive loops
type userSSHKeySetTitleInputAlias UserSSHKeySetTitleInput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserSSHKeySetTitleInput) UnmarshalJSON(data []byte) error {
	var alias userSSHKeySetTitleInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserSSHKeySetTitleInput)(&alias)).InitializeDefaults()
	*e = UserSSHKeySetTitleInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserSSHKeySetTitleInput) MarshalJSON() ([]byte, error) {
	alias := userSSHKeySetTitleInputAlias(e)
	return json.Marshal(alias)
}

// NewUserSSHKeySetTitleOutput creates a new UserSSHKeySetTitleOutput
func NewUserSSHKeySetTitleOutput() *UserSSHKeySetTitleOutput {
	s := &UserSSHKeySetTitleOutput{}
	s.InitializeDefaults()
	return s
}

// UserSSHKeySetTitleOutput struct
type UserSSHKeySetTitleOutput struct {
	SshKey *UserSSHKey `json:"sshKey,omitempty" yaml:"sshKey,omitempty"`
}

// GetSshKey returns the value for the field sshKey
func (e *UserSSHKeySetTitleOutput) GetSshKey() *UserSSHKey {
	return e.SshKey
}

// SetSshKey sets the value for the field sshKey
func (e *UserSSHKeySetTitleOutput) SetSshKey(sshKey *UserSSHKey) {
	e.SshKey = sshKey
}

// StructPath returns StructPath
func (e *UserSSHKeySetTitleOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserSSHKeySetTitleOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserSSHKeySetTitleOutput) InitializeDefaults() {
}

// userSSHKeySetTitleOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type userSSHKeySetTitleOutputAlias UserSSHKeySetTitleOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserSSHKeySetTitleOutput) UnmarshalJSON(data []byte) error {
	var alias userSSHKeySetTitleOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserSSHKeySetTitleOutput)(&alias)).InitializeDefaults()
	*e = UserSSHKeySetTitleOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserSSHKeySetTitleOutput) MarshalJSON() ([]byte, error) {
	alias := userSSHKeySetTitleOutputAlias(e)
	return json.Marshal(alias)
}

// NewUserGroupListInput creates a new UserGroupListInput
func NewUserGroupListInput() *UserGroupListInput {
	s := &UserGroupListInput{}
	s.InitializeDefaults()
	return s
}

// UserGroupListInput struct
type UserGroupListInput struct {
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}

// GetUsername returns the value for the field username
func (e *UserGroupListInput) GetUsername() string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *UserGroupListInput) SetUsername(username string) {
	e.Username = username
}

// StructPath returns StructPath
func (e *UserGroupListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserGroupListInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserGroupListInput) InitializeDefaults() {
}

// userGroupListInputAlias is defined to help pre and post JSON marshaling without recursive loops
type userGroupListInputAlias UserGroupListInput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserGroupListInput) UnmarshalJSON(data []byte) error {
	var alias userGroupListInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserGroupListInput)(&alias)).InitializeDefaults()
	*e = UserGroupListInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserGroupListInput) MarshalJSON() ([]byte, error) {
	alias := userGroupListInputAlias(e)
	return json.Marshal(alias)
}

// NewUserGroupListOutput creates a new UserGroupListOutput
func NewUserGroupListOutput() *UserGroupListOutput {
	s := &UserGroupListOutput{}
	s.InitializeDefaults()
	return s
}

// UserGroupListOutput struct
type UserGroupListOutput struct {
	Groups []*GroupInformation `json:"groups,omitempty" yaml:"groups,omitempty"`
}

// GetGroups returns the value for the field groups
func (e *UserGroupListOutput) GetGroups() []*GroupInformation {
	return e.Groups
}

// SetGroups sets the value for the field groups
func (e *UserGroupListOutput) SetGroups(groups []*GroupInformation) {
	e.Groups = groups
}

// StructPath returns StructPath
func (e *UserGroupListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathUserGroupListOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *UserGroupListOutput) InitializeDefaults() {
}

// userGroupListOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type userGroupListOutputAlias UserGroupListOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *UserGroupListOutput) UnmarshalJSON(data []byte) error {
	var alias userGroupListOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*UserGroupListOutput)(&alias)).InitializeDefaults()
	*e = UserGroupListOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e UserGroupListOutput) MarshalJSON() ([]byte, error) {
	alias := userGroupListOutputAlias(e)
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

// NewPolicyStructureError creates a new PolicyStructureError
func NewPolicyStructureError() *PolicyStructureError {
	s := &PolicyStructureError{}
	s.InitializeDefaults()
	return s
}

// PolicyStructureError struct
type PolicyStructureError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *PolicyStructureError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [PolicyStructureError]
func (e *PolicyStructureError) Is(err error) bool {
	_, ok := err.(*PolicyStructureError)
	return ok
}

// IsPolicyStructureError indicates whether the given error chain contains an error of type [PolicyStructureError]
func IsPolicyStructureError(err error) bool {
	return errors.Is(err, &PolicyStructureError{})
}

// GetMessage returns the value for the field message
func (e *PolicyStructureError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *PolicyStructureError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *PolicyStructureError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathPolicyStructureError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *PolicyStructureError) InitializeDefaults() {
}

// policyStructureErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type policyStructureErrorAlias PolicyStructureError

// UnmarshalJSON implements json.Unmarshaler
func (e *PolicyStructureError) UnmarshalJSON(data []byte) error {
	var alias policyStructureErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*PolicyStructureError)(&alias)).InitializeDefaults()
	*e = PolicyStructureError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e PolicyStructureError) MarshalJSON() ([]byte, error) {
	alias := policyStructureErrorAlias(e)
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
	// optional hard cap on the lifetime of any credential assumed for this
	// role, in seconds. Bounds total life regardless of activity/keep-alive.
	// Omit to use the system default (43200, 12h).
	MaxSessionLifetimeSeconds *int32 `json:"maxSessionLifetimeSeconds,omitempty" yaml:"maxSessionLifetimeSeconds,omitempty"`
	Name                      string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetDescription returns the value for the field description
func (e *RoleCreateInput) GetDescription() string {
	return e.Description
}

// SetDescription sets the value for the field description
func (e *RoleCreateInput) SetDescription(description string) {
	e.Description = description
}

// GetMaxSessionLifetimeSeconds returns the value for the field maxSessionLifetimeSeconds
func (e *RoleCreateInput) GetMaxSessionLifetimeSeconds() *int32 {
	return e.MaxSessionLifetimeSeconds
}

// SetMaxSessionLifetimeSeconds sets the value for the field maxSessionLifetimeSeconds
func (e *RoleCreateInput) SetMaxSessionLifetimeSeconds(maxSessionLifetimeSeconds *int32) {
	e.MaxSessionLifetimeSeconds = maxSessionLifetimeSeconds
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
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetName returns the value for the field name
func (e *RoleGetInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *RoleGetInput) SetName(name string) {
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
	// DRN of a target role in an explicitly named (possibly different)
	// account, e.g. "account(acme) iam:Role(deployer)". Used for
	// cross-account assume, e.g. a service assuming a customer-account role
	// whose trust policy admits iam:Service(<fqdn>). The account(<name>)
	// qualifier is required and the resource must be an iam:Role(<name>); an
	// optional region qualifier must match this IAM region. Provide exactly
	// one of name or drn.
	Drn *string `json:"drn,omitempty" yaml:"drn,omitempty"`
	// idle/sliding window in seconds for the assumed credential. Activity
	// (or Session.KeepAlive) slides the expiry forward by this much, up to
	// the hard cap. min 900 (15m), default 3600 (1h), max 43200 (12h).
	DurationSeconds int32 `json:"durationSeconds,omitempty" yaml:"durationSeconds,omitempty"`
	// optional inline policy to attach to the assumed role
	InlinePolicy *InlinePolicy `json:"inlinePolicy,omitempty" yaml:"inlinePolicy,omitempty"`
	// optional hard max-lifetime in seconds. Bounds total life regardless
	// of activity/keep-alive. Clamped to the role's configured max; never
	// higher. Omit to use the role's configured max.
	MaxLifetimeSeconds *int32 `json:"maxLifetimeSeconds,omitempty" yaml:"maxLifetimeSeconds,omitempty"`
	// name of the role to assume in the caller's own account. Provide exactly
	// one of name or drn.
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetDrn returns the value for the field drn
func (e *RoleAssumeInput) GetDrn() *string {
	return e.Drn
}

// SetDrn sets the value for the field drn
func (e *RoleAssumeInput) SetDrn(drn *string) {
	e.Drn = drn
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

// GetMaxLifetimeSeconds returns the value for the field maxLifetimeSeconds
func (e *RoleAssumeInput) GetMaxLifetimeSeconds() *int32 {
	return e.MaxLifetimeSeconds
}

// SetMaxLifetimeSeconds sets the value for the field maxLifetimeSeconds
func (e *RoleAssumeInput) SetMaxLifetimeSeconds(maxLifetimeSeconds *int32) {
	e.MaxLifetimeSeconds = maxLifetimeSeconds
}

// GetName returns the value for the field name
func (e *RoleAssumeInput) GetName() *string {
	return e.Name
}

// SetName sets the value for the field name
func (e *RoleAssumeInput) SetName(name *string) {
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

// NewRoleAssumeWithWebIdentityInput creates a new RoleAssumeWithWebIdentityInput
func NewRoleAssumeWithWebIdentityInput() *RoleAssumeWithWebIdentityInput {
	s := &RoleAssumeWithWebIdentityInput{}
	s.InitializeDefaults()
	return s
}

// RoleAssumeWithWebIdentityInput struct
type RoleAssumeWithWebIdentityInput struct {
	// DRN of the target principal to assume, identifying both the account
	// and the role, e.g. "account(account1) iam:Role(my-role)". The
	// account(<name>) qualifier is required and the resource must be an
	// iam:Role(<name>); an incomplete or non-role DRN is rejected.
	Drn string `json:"drn,omitempty" yaml:"drn,omitempty"`
	// idle/sliding window in seconds for the assumed credential (see
	// Role.Assume). Optional; defaults to 3600 (1h).
	DurationSeconds *int32 `json:"durationSeconds,omitempty" yaml:"durationSeconds,omitempty"`
	// optional inline policy to narrow the assumed session (must be a subset)
	InlinePolicy *InlinePolicy `json:"inlinePolicy,omitempty" yaml:"inlinePolicy,omitempty"`
	// optional hard max-lifetime in seconds, clamped to the role's max.
	MaxLifetimeSeconds *int32 `json:"maxLifetimeSeconds,omitempty" yaml:"maxLifetimeSeconds,omitempty"`
	// the external OIDC JWT
	WebIdentityToken string `json:"webIdentityToken,omitempty" yaml:"webIdentityToken,omitempty"`
}

// GetDrn returns the value for the field drn
func (e *RoleAssumeWithWebIdentityInput) GetDrn() string {
	return e.Drn
}

// SetDrn sets the value for the field drn
func (e *RoleAssumeWithWebIdentityInput) SetDrn(drn string) {
	e.Drn = drn
}

// GetDurationSeconds returns the value for the field durationSeconds
func (e *RoleAssumeWithWebIdentityInput) GetDurationSeconds() *int32 {
	return e.DurationSeconds
}

// SetDurationSeconds sets the value for the field durationSeconds
func (e *RoleAssumeWithWebIdentityInput) SetDurationSeconds(durationSeconds *int32) {
	e.DurationSeconds = durationSeconds
}

// GetInlinePolicy returns the value for the field inlinePolicy
func (e *RoleAssumeWithWebIdentityInput) GetInlinePolicy() *InlinePolicy {
	return e.InlinePolicy
}

// SetInlinePolicy sets the value for the field inlinePolicy
func (e *RoleAssumeWithWebIdentityInput) SetInlinePolicy(inlinePolicy *InlinePolicy) {
	e.InlinePolicy = inlinePolicy
}

// GetMaxLifetimeSeconds returns the value for the field maxLifetimeSeconds
func (e *RoleAssumeWithWebIdentityInput) GetMaxLifetimeSeconds() *int32 {
	return e.MaxLifetimeSeconds
}

// SetMaxLifetimeSeconds sets the value for the field maxLifetimeSeconds
func (e *RoleAssumeWithWebIdentityInput) SetMaxLifetimeSeconds(maxLifetimeSeconds *int32) {
	e.MaxLifetimeSeconds = maxLifetimeSeconds
}

// GetWebIdentityToken returns the value for the field webIdentityToken
func (e *RoleAssumeWithWebIdentityInput) GetWebIdentityToken() string {
	return e.WebIdentityToken
}

// SetWebIdentityToken sets the value for the field webIdentityToken
func (e *RoleAssumeWithWebIdentityInput) SetWebIdentityToken(webIdentityToken string) {
	e.WebIdentityToken = webIdentityToken
}

// StructPath returns StructPath
func (e *RoleAssumeWithWebIdentityInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleAssumeWithWebIdentityInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleAssumeWithWebIdentityInput) InitializeDefaults() {
}

// roleAssumeWithWebIdentityInputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleAssumeWithWebIdentityInputAlias RoleAssumeWithWebIdentityInput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleAssumeWithWebIdentityInput) UnmarshalJSON(data []byte) error {
	var alias roleAssumeWithWebIdentityInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleAssumeWithWebIdentityInput)(&alias)).InitializeDefaults()
	*e = RoleAssumeWithWebIdentityInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleAssumeWithWebIdentityInput) MarshalJSON() ([]byte, error) {
	alias := roleAssumeWithWebIdentityInputAlias(e)
	return json.Marshal(alias)
}

// NewRoleAssumeWithWebIdentityOutput creates a new RoleAssumeWithWebIdentityOutput
func NewRoleAssumeWithWebIdentityOutput() *RoleAssumeWithWebIdentityOutput {
	s := &RoleAssumeWithWebIdentityOutput{}
	s.InitializeDefaults()
	return s
}

// RoleAssumeWithWebIdentityOutput struct
type RoleAssumeWithWebIdentityOutput struct {
	Credentials *Credentials `json:"credentials,omitempty" yaml:"credentials,omitempty"`
}

// GetCredentials returns the value for the field credentials
func (e *RoleAssumeWithWebIdentityOutput) GetCredentials() *Credentials {
	return e.Credentials
}

// SetCredentials sets the value for the field credentials
func (e *RoleAssumeWithWebIdentityOutput) SetCredentials(credentials *Credentials) {
	e.Credentials = credentials
}

// StructPath returns StructPath
func (e *RoleAssumeWithWebIdentityOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleAssumeWithWebIdentityOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleAssumeWithWebIdentityOutput) InitializeDefaults() {
}

// roleAssumeWithWebIdentityOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleAssumeWithWebIdentityOutputAlias RoleAssumeWithWebIdentityOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleAssumeWithWebIdentityOutput) UnmarshalJSON(data []byte) error {
	var alias roleAssumeWithWebIdentityOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleAssumeWithWebIdentityOutput)(&alias)).InitializeDefaults()
	*e = RoleAssumeWithWebIdentityOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleAssumeWithWebIdentityOutput) MarshalJSON() ([]byte, error) {
	alias := roleAssumeWithWebIdentityOutputAlias(e)
	return json.Marshal(alias)
}

// NewRoleAccessKeyListInput creates a new RoleAccessKeyListInput
func NewRoleAccessKeyListInput() *RoleAccessKeyListInput {
	s := &RoleAccessKeyListInput{}
	s.InitializeDefaults()
	return s
}

// RoleAccessKeyListInput struct
type RoleAccessKeyListInput struct {
	RoleName string `json:"roleName,omitempty" yaml:"roleName,omitempty"`
}

// GetRoleName returns the value for the field roleName
func (e *RoleAccessKeyListInput) GetRoleName() string {
	return e.RoleName
}

// SetRoleName sets the value for the field roleName
func (e *RoleAccessKeyListInput) SetRoleName(roleName string) {
	e.RoleName = roleName
}

// StructPath returns StructPath
func (e *RoleAccessKeyListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleAccessKeyListInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleAccessKeyListInput) InitializeDefaults() {
}

// roleAccessKeyListInputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleAccessKeyListInputAlias RoleAccessKeyListInput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleAccessKeyListInput) UnmarshalJSON(data []byte) error {
	var alias roleAccessKeyListInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleAccessKeyListInput)(&alias)).InitializeDefaults()
	*e = RoleAccessKeyListInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleAccessKeyListInput) MarshalJSON() ([]byte, error) {
	alias := roleAccessKeyListInputAlias(e)
	return json.Marshal(alias)
}

// NewRoleAccessKeyListOutput creates a new RoleAccessKeyListOutput
func NewRoleAccessKeyListOutput() *RoleAccessKeyListOutput {
	s := &RoleAccessKeyListOutput{}
	s.InitializeDefaults()
	return s
}

// RoleAccessKeyListOutput struct
type RoleAccessKeyListOutput struct {
	Credentials []*CredentialInfo `json:"credentials,omitempty" yaml:"credentials,omitempty"`
}

// GetCredentials returns the value for the field credentials
func (e *RoleAccessKeyListOutput) GetCredentials() []*CredentialInfo {
	return e.Credentials
}

// SetCredentials sets the value for the field credentials
func (e *RoleAccessKeyListOutput) SetCredentials(credentials []*CredentialInfo) {
	e.Credentials = credentials
}

// StructPath returns StructPath
func (e *RoleAccessKeyListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleAccessKeyListOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleAccessKeyListOutput) InitializeDefaults() {
}

// roleAccessKeyListOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleAccessKeyListOutputAlias RoleAccessKeyListOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleAccessKeyListOutput) UnmarshalJSON(data []byte) error {
	var alias roleAccessKeyListOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleAccessKeyListOutput)(&alias)).InitializeDefaults()
	*e = RoleAccessKeyListOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleAccessKeyListOutput) MarshalJSON() ([]byte, error) {
	alias := roleAccessKeyListOutputAlias(e)
	return json.Marshal(alias)
}

// NewRoleAccessKeyDestroyInput creates a new RoleAccessKeyDestroyInput
func NewRoleAccessKeyDestroyInput() *RoleAccessKeyDestroyInput {
	s := &RoleAccessKeyDestroyInput{}
	s.InitializeDefaults()
	return s
}

// RoleAccessKeyDestroyInput struct
type RoleAccessKeyDestroyInput struct {
	AccessKeyID string `json:"accessKeyID,omitempty" yaml:"accessKeyID,omitempty"`
	RoleName    string `json:"roleName,omitempty" yaml:"roleName,omitempty"`
}

// GetAccessKeyID returns the value for the field accessKeyID
func (e *RoleAccessKeyDestroyInput) GetAccessKeyID() string {
	return e.AccessKeyID
}

// SetAccessKeyID sets the value for the field accessKeyID
func (e *RoleAccessKeyDestroyInput) SetAccessKeyID(accessKeyID string) {
	e.AccessKeyID = accessKeyID
}

// GetRoleName returns the value for the field roleName
func (e *RoleAccessKeyDestroyInput) GetRoleName() string {
	return e.RoleName
}

// SetRoleName sets the value for the field roleName
func (e *RoleAccessKeyDestroyInput) SetRoleName(roleName string) {
	e.RoleName = roleName
}

// StructPath returns StructPath
func (e *RoleAccessKeyDestroyInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleAccessKeyDestroyInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleAccessKeyDestroyInput) InitializeDefaults() {
}

// roleAccessKeyDestroyInputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleAccessKeyDestroyInputAlias RoleAccessKeyDestroyInput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleAccessKeyDestroyInput) UnmarshalJSON(data []byte) error {
	var alias roleAccessKeyDestroyInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleAccessKeyDestroyInput)(&alias)).InitializeDefaults()
	*e = RoleAccessKeyDestroyInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleAccessKeyDestroyInput) MarshalJSON() ([]byte, error) {
	alias := roleAccessKeyDestroyInputAlias(e)
	return json.Marshal(alias)
}

// NewRoleAccessKeyDestroyOutput creates a new RoleAccessKeyDestroyOutput
func NewRoleAccessKeyDestroyOutput() *RoleAccessKeyDestroyOutput {
	s := &RoleAccessKeyDestroyOutput{}
	s.InitializeDefaults()
	return s
}

// RoleAccessKeyDestroyOutput struct
type RoleAccessKeyDestroyOutput struct {
}

// StructPath returns StructPath
func (e *RoleAccessKeyDestroyOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleAccessKeyDestroyOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleAccessKeyDestroyOutput) InitializeDefaults() {
}

// roleAccessKeyDestroyOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleAccessKeyDestroyOutputAlias RoleAccessKeyDestroyOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleAccessKeyDestroyOutput) UnmarshalJSON(data []byte) error {
	var alias roleAccessKeyDestroyOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleAccessKeyDestroyOutput)(&alias)).InitializeDefaults()
	*e = RoleAccessKeyDestroyOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleAccessKeyDestroyOutput) MarshalJSON() ([]byte, error) {
	alias := roleAccessKeyDestroyOutputAlias(e)
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
	Attachments []*IdentityPolicyAttachmentInfo `json:"attachments,omitempty" yaml:"attachments,omitempty"`
}

// GetAttachments returns the value for the field attachments
func (e *RoleIdentityPolicyListOutput) GetAttachments() []*IdentityPolicyAttachmentInfo {
	return e.Attachments
}

// SetAttachments sets the value for the field attachments
func (e *RoleIdentityPolicyListOutput) SetAttachments(attachments []*IdentityPolicyAttachmentInfo) {
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

// NewRoleTrustPolicyAttachInput creates a new RoleTrustPolicyAttachInput
func NewRoleTrustPolicyAttachInput() *RoleTrustPolicyAttachInput {
	s := &RoleTrustPolicyAttachInput{}
	s.InitializeDefaults()
	return s
}

// RoleTrustPolicyAttachInput struct
type RoleTrustPolicyAttachInput struct {
	RoleName        string `json:"roleName,omitempty" yaml:"roleName,omitempty"`
	TrustPolicyName string `json:"trustPolicyName,omitempty" yaml:"trustPolicyName,omitempty"`
}

// GetRoleName returns the value for the field roleName
func (e *RoleTrustPolicyAttachInput) GetRoleName() string {
	return e.RoleName
}

// SetRoleName sets the value for the field roleName
func (e *RoleTrustPolicyAttachInput) SetRoleName(roleName string) {
	e.RoleName = roleName
}

// GetTrustPolicyName returns the value for the field trustPolicyName
func (e *RoleTrustPolicyAttachInput) GetTrustPolicyName() string {
	return e.TrustPolicyName
}

// SetTrustPolicyName sets the value for the field trustPolicyName
func (e *RoleTrustPolicyAttachInput) SetTrustPolicyName(trustPolicyName string) {
	e.TrustPolicyName = trustPolicyName
}

// StructPath returns StructPath
func (e *RoleTrustPolicyAttachInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleTrustPolicyAttachInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleTrustPolicyAttachInput) InitializeDefaults() {
}

// roleTrustPolicyAttachInputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleTrustPolicyAttachInputAlias RoleTrustPolicyAttachInput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleTrustPolicyAttachInput) UnmarshalJSON(data []byte) error {
	var alias roleTrustPolicyAttachInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleTrustPolicyAttachInput)(&alias)).InitializeDefaults()
	*e = RoleTrustPolicyAttachInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleTrustPolicyAttachInput) MarshalJSON() ([]byte, error) {
	alias := roleTrustPolicyAttachInputAlias(e)
	return json.Marshal(alias)
}

// NewRoleTrustPolicyAttachOutput creates a new RoleTrustPolicyAttachOutput
func NewRoleTrustPolicyAttachOutput() *RoleTrustPolicyAttachOutput {
	s := &RoleTrustPolicyAttachOutput{}
	s.InitializeDefaults()
	return s
}

// RoleTrustPolicyAttachOutput struct
type RoleTrustPolicyAttachOutput struct {
}

// StructPath returns StructPath
func (e *RoleTrustPolicyAttachOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleTrustPolicyAttachOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleTrustPolicyAttachOutput) InitializeDefaults() {
}

// roleTrustPolicyAttachOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleTrustPolicyAttachOutputAlias RoleTrustPolicyAttachOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleTrustPolicyAttachOutput) UnmarshalJSON(data []byte) error {
	var alias roleTrustPolicyAttachOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleTrustPolicyAttachOutput)(&alias)).InitializeDefaults()
	*e = RoleTrustPolicyAttachOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleTrustPolicyAttachOutput) MarshalJSON() ([]byte, error) {
	alias := roleTrustPolicyAttachOutputAlias(e)
	return json.Marshal(alias)
}

// NewRoleTrustPolicyListInput creates a new RoleTrustPolicyListInput
func NewRoleTrustPolicyListInput() *RoleTrustPolicyListInput {
	s := &RoleTrustPolicyListInput{}
	s.InitializeDefaults()
	return s
}

// RoleTrustPolicyListInput struct
type RoleTrustPolicyListInput struct {
	RoleName string `json:"roleName,omitempty" yaml:"roleName,omitempty"`
}

// GetRoleName returns the value for the field roleName
func (e *RoleTrustPolicyListInput) GetRoleName() string {
	return e.RoleName
}

// SetRoleName sets the value for the field roleName
func (e *RoleTrustPolicyListInput) SetRoleName(roleName string) {
	e.RoleName = roleName
}

// StructPath returns StructPath
func (e *RoleTrustPolicyListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleTrustPolicyListInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleTrustPolicyListInput) InitializeDefaults() {
}

// roleTrustPolicyListInputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleTrustPolicyListInputAlias RoleTrustPolicyListInput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleTrustPolicyListInput) UnmarshalJSON(data []byte) error {
	var alias roleTrustPolicyListInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleTrustPolicyListInput)(&alias)).InitializeDefaults()
	*e = RoleTrustPolicyListInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleTrustPolicyListInput) MarshalJSON() ([]byte, error) {
	alias := roleTrustPolicyListInputAlias(e)
	return json.Marshal(alias)
}

// NewRoleTrustPolicyListOutput creates a new RoleTrustPolicyListOutput
func NewRoleTrustPolicyListOutput() *RoleTrustPolicyListOutput {
	s := &RoleTrustPolicyListOutput{}
	s.InitializeDefaults()
	return s
}

// RoleTrustPolicyListOutput struct
type RoleTrustPolicyListOutput struct {
	TrustPolicies []*TrustPolicy `json:"trustPolicies,omitempty" yaml:"trustPolicies,omitempty"`
}

// GetTrustPolicies returns the value for the field trustPolicies
func (e *RoleTrustPolicyListOutput) GetTrustPolicies() []*TrustPolicy {
	return e.TrustPolicies
}

// SetTrustPolicies sets the value for the field trustPolicies
func (e *RoleTrustPolicyListOutput) SetTrustPolicies(trustPolicies []*TrustPolicy) {
	e.TrustPolicies = trustPolicies
}

// StructPath returns StructPath
func (e *RoleTrustPolicyListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleTrustPolicyListOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleTrustPolicyListOutput) InitializeDefaults() {
}

// roleTrustPolicyListOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleTrustPolicyListOutputAlias RoleTrustPolicyListOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleTrustPolicyListOutput) UnmarshalJSON(data []byte) error {
	var alias roleTrustPolicyListOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleTrustPolicyListOutput)(&alias)).InitializeDefaults()
	*e = RoleTrustPolicyListOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleTrustPolicyListOutput) MarshalJSON() ([]byte, error) {
	alias := roleTrustPolicyListOutputAlias(e)
	return json.Marshal(alias)
}

// NewRoleTrustPolicyDetachInput creates a new RoleTrustPolicyDetachInput
func NewRoleTrustPolicyDetachInput() *RoleTrustPolicyDetachInput {
	s := &RoleTrustPolicyDetachInput{}
	s.InitializeDefaults()
	return s
}

// RoleTrustPolicyDetachInput struct
type RoleTrustPolicyDetachInput struct {
	RoleName        string `json:"roleName,omitempty" yaml:"roleName,omitempty"`
	TrustPolicyName string `json:"trustPolicyName,omitempty" yaml:"trustPolicyName,omitempty"`
}

// GetRoleName returns the value for the field roleName
func (e *RoleTrustPolicyDetachInput) GetRoleName() string {
	return e.RoleName
}

// SetRoleName sets the value for the field roleName
func (e *RoleTrustPolicyDetachInput) SetRoleName(roleName string) {
	e.RoleName = roleName
}

// GetTrustPolicyName returns the value for the field trustPolicyName
func (e *RoleTrustPolicyDetachInput) GetTrustPolicyName() string {
	return e.TrustPolicyName
}

// SetTrustPolicyName sets the value for the field trustPolicyName
func (e *RoleTrustPolicyDetachInput) SetTrustPolicyName(trustPolicyName string) {
	e.TrustPolicyName = trustPolicyName
}

// StructPath returns StructPath
func (e *RoleTrustPolicyDetachInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleTrustPolicyDetachInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleTrustPolicyDetachInput) InitializeDefaults() {
}

// roleTrustPolicyDetachInputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleTrustPolicyDetachInputAlias RoleTrustPolicyDetachInput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleTrustPolicyDetachInput) UnmarshalJSON(data []byte) error {
	var alias roleTrustPolicyDetachInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleTrustPolicyDetachInput)(&alias)).InitializeDefaults()
	*e = RoleTrustPolicyDetachInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleTrustPolicyDetachInput) MarshalJSON() ([]byte, error) {
	alias := roleTrustPolicyDetachInputAlias(e)
	return json.Marshal(alias)
}

// NewRoleTrustPolicyDetachOutput creates a new RoleTrustPolicyDetachOutput
func NewRoleTrustPolicyDetachOutput() *RoleTrustPolicyDetachOutput {
	s := &RoleTrustPolicyDetachOutput{}
	s.InitializeDefaults()
	return s
}

// RoleTrustPolicyDetachOutput struct
type RoleTrustPolicyDetachOutput struct {
}

// StructPath returns StructPath
func (e *RoleTrustPolicyDetachOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathRoleTrustPolicyDetachOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *RoleTrustPolicyDetachOutput) InitializeDefaults() {
}

// roleTrustPolicyDetachOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type roleTrustPolicyDetachOutputAlias RoleTrustPolicyDetachOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *RoleTrustPolicyDetachOutput) UnmarshalJSON(data []byte) error {
	var alias roleTrustPolicyDetachOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*RoleTrustPolicyDetachOutput)(&alias)).InitializeDefaults()
	*e = RoleTrustPolicyDetachOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e RoleTrustPolicyDetachOutput) MarshalJSON() ([]byte, error) {
	alias := roleTrustPolicyDetachOutputAlias(e)
	return json.Marshal(alias)
}

// NewGroupCreateInput creates a new GroupCreateInput
func NewGroupCreateInput() *GroupCreateInput {
	s := &GroupCreateInput{}
	s.InitializeDefaults()
	return s
}

// GroupCreateInput struct
type GroupCreateInput struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetDescription returns the value for the field description
func (e *GroupCreateInput) GetDescription() string {
	return e.Description
}

// SetDescription sets the value for the field description
func (e *GroupCreateInput) SetDescription(description string) {
	e.Description = description
}

// GetName returns the value for the field name
func (e *GroupCreateInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *GroupCreateInput) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *GroupCreateInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathGroupCreateInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *GroupCreateInput) InitializeDefaults() {
}

// groupCreateInputAlias is defined to help pre and post JSON marshaling without recursive loops
type groupCreateInputAlias GroupCreateInput

// UnmarshalJSON implements json.Unmarshaler
func (e *GroupCreateInput) UnmarshalJSON(data []byte) error {
	var alias groupCreateInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*GroupCreateInput)(&alias)).InitializeDefaults()
	*e = GroupCreateInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e GroupCreateInput) MarshalJSON() ([]byte, error) {
	alias := groupCreateInputAlias(e)
	return json.Marshal(alias)
}

// NewGroupCreateOutput creates a new GroupCreateOutput
func NewGroupCreateOutput() *GroupCreateOutput {
	s := &GroupCreateOutput{}
	s.InitializeDefaults()
	return s
}

// GroupCreateOutput struct
type GroupCreateOutput struct {
	Group *GroupInformation `json:"group,omitempty" yaml:"group,omitempty"`
}

// GetGroup returns the value for the field group
func (e *GroupCreateOutput) GetGroup() *GroupInformation {
	return e.Group
}

// SetGroup sets the value for the field group
func (e *GroupCreateOutput) SetGroup(group *GroupInformation) {
	e.Group = group
}

// StructPath returns StructPath
func (e *GroupCreateOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathGroupCreateOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *GroupCreateOutput) InitializeDefaults() {
}

// groupCreateOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type groupCreateOutputAlias GroupCreateOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *GroupCreateOutput) UnmarshalJSON(data []byte) error {
	var alias groupCreateOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*GroupCreateOutput)(&alias)).InitializeDefaults()
	*e = GroupCreateOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e GroupCreateOutput) MarshalJSON() ([]byte, error) {
	alias := groupCreateOutputAlias(e)
	return json.Marshal(alias)
}

// NewGroupDestroyInput creates a new GroupDestroyInput
func NewGroupDestroyInput() *GroupDestroyInput {
	s := &GroupDestroyInput{}
	s.InitializeDefaults()
	return s
}

// GroupDestroyInput struct
type GroupDestroyInput struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetName returns the value for the field name
func (e *GroupDestroyInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *GroupDestroyInput) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *GroupDestroyInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathGroupDestroyInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *GroupDestroyInput) InitializeDefaults() {
}

// groupDestroyInputAlias is defined to help pre and post JSON marshaling without recursive loops
type groupDestroyInputAlias GroupDestroyInput

// UnmarshalJSON implements json.Unmarshaler
func (e *GroupDestroyInput) UnmarshalJSON(data []byte) error {
	var alias groupDestroyInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*GroupDestroyInput)(&alias)).InitializeDefaults()
	*e = GroupDestroyInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e GroupDestroyInput) MarshalJSON() ([]byte, error) {
	alias := groupDestroyInputAlias(e)
	return json.Marshal(alias)
}

// NewGroupDestroyOutput creates a new GroupDestroyOutput
func NewGroupDestroyOutput() *GroupDestroyOutput {
	s := &GroupDestroyOutput{}
	s.InitializeDefaults()
	return s
}

// GroupDestroyOutput struct
type GroupDestroyOutput struct {
}

// StructPath returns StructPath
func (e *GroupDestroyOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathGroupDestroyOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *GroupDestroyOutput) InitializeDefaults() {
}

// groupDestroyOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type groupDestroyOutputAlias GroupDestroyOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *GroupDestroyOutput) UnmarshalJSON(data []byte) error {
	var alias groupDestroyOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*GroupDestroyOutput)(&alias)).InitializeDefaults()
	*e = GroupDestroyOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e GroupDestroyOutput) MarshalJSON() ([]byte, error) {
	alias := groupDestroyOutputAlias(e)
	return json.Marshal(alias)
}

// NewGroupGetInput creates a new GroupGetInput
func NewGroupGetInput() *GroupGetInput {
	s := &GroupGetInput{}
	s.InitializeDefaults()
	return s
}

// GroupGetInput struct
type GroupGetInput struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetName returns the value for the field name
func (e *GroupGetInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *GroupGetInput) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *GroupGetInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathGroupGetInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *GroupGetInput) InitializeDefaults() {
}

// groupGetInputAlias is defined to help pre and post JSON marshaling without recursive loops
type groupGetInputAlias GroupGetInput

// UnmarshalJSON implements json.Unmarshaler
func (e *GroupGetInput) UnmarshalJSON(data []byte) error {
	var alias groupGetInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*GroupGetInput)(&alias)).InitializeDefaults()
	*e = GroupGetInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e GroupGetInput) MarshalJSON() ([]byte, error) {
	alias := groupGetInputAlias(e)
	return json.Marshal(alias)
}

// NewGroupGetOutput creates a new GroupGetOutput
func NewGroupGetOutput() *GroupGetOutput {
	s := &GroupGetOutput{}
	s.InitializeDefaults()
	return s
}

// GroupGetOutput struct
type GroupGetOutput struct {
	Group *GroupInformation `json:"group,omitempty" yaml:"group,omitempty"`
}

// GetGroup returns the value for the field group
func (e *GroupGetOutput) GetGroup() *GroupInformation {
	return e.Group
}

// SetGroup sets the value for the field group
func (e *GroupGetOutput) SetGroup(group *GroupInformation) {
	e.Group = group
}

// StructPath returns StructPath
func (e *GroupGetOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathGroupGetOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *GroupGetOutput) InitializeDefaults() {
}

// groupGetOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type groupGetOutputAlias GroupGetOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *GroupGetOutput) UnmarshalJSON(data []byte) error {
	var alias groupGetOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*GroupGetOutput)(&alias)).InitializeDefaults()
	*e = GroupGetOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e GroupGetOutput) MarshalJSON() ([]byte, error) {
	alias := groupGetOutputAlias(e)
	return json.Marshal(alias)
}

// NewGroupListInput creates a new GroupListInput
func NewGroupListInput() *GroupListInput {
	s := &GroupListInput{}
	s.InitializeDefaults()
	return s
}

// GroupListInput struct
type GroupListInput struct {
	// opaque page cursor: a previous response's nextCursor, verbatim
	Cursor *string `json:"cursor,omitempty" yaml:"cursor,omitempty"`
	// page size, 1 to 200; 50 when absent
	Limit *int32 `json:"limit,omitempty" yaml:"limit,omitempty"`
	// when true, only service-managed groups; when false, only groups the account
	// authored. Absent returns both, which is today's behaviour and keeps every
	// existing caller correct
	Managed *bool `json:"managed,omitempty" yaml:"managed,omitempty"`
	// the managing service principal FQDN, e.g. "uplink.deployport.io", to
	// return only that service's groups. Absent applies no service filter
	ManagedByService *string `json:"managedByService,omitempty" yaml:"managedByService,omitempty"`
}

// GetCursor returns the value for the field cursor
func (e *GroupListInput) GetCursor() *string {
	return e.Cursor
}

// SetCursor sets the value for the field cursor
func (e *GroupListInput) SetCursor(cursor *string) {
	e.Cursor = cursor
}

// GetLimit returns the value for the field limit
func (e *GroupListInput) GetLimit() *int32 {
	return e.Limit
}

// SetLimit sets the value for the field limit
func (e *GroupListInput) SetLimit(limit *int32) {
	e.Limit = limit
}

// GetManaged returns the value for the field managed
func (e *GroupListInput) GetManaged() *bool {
	return e.Managed
}

// SetManaged sets the value for the field managed
func (e *GroupListInput) SetManaged(managed *bool) {
	e.Managed = managed
}

// GetManagedByService returns the value for the field managedByService
func (e *GroupListInput) GetManagedByService() *string {
	return e.ManagedByService
}

// SetManagedByService sets the value for the field managedByService
func (e *GroupListInput) SetManagedByService(managedByService *string) {
	e.ManagedByService = managedByService
}

// StructPath returns StructPath
func (e *GroupListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathGroupListInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *GroupListInput) InitializeDefaults() {
}

// groupListInputAlias is defined to help pre and post JSON marshaling without recursive loops
type groupListInputAlias GroupListInput

// UnmarshalJSON implements json.Unmarshaler
func (e *GroupListInput) UnmarshalJSON(data []byte) error {
	var alias groupListInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*GroupListInput)(&alias)).InitializeDefaults()
	*e = GroupListInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e GroupListInput) MarshalJSON() ([]byte, error) {
	alias := groupListInputAlias(e)
	return json.Marshal(alias)
}

// NewGroupListOutput creates a new GroupListOutput
func NewGroupListOutput() *GroupListOutput {
	s := &GroupListOutput{}
	s.InitializeDefaults()
	return s
}

// GroupListOutput struct
type GroupListOutput struct {
	Groups []*GroupInformation `json:"groups,omitempty" yaml:"groups,omitempty"`
	// pass to the next call's cursor; absent on the last page
	NextCursor *string `json:"nextCursor,omitempty" yaml:"nextCursor,omitempty"`
}

// GetGroups returns the value for the field groups
func (e *GroupListOutput) GetGroups() []*GroupInformation {
	return e.Groups
}

// SetGroups sets the value for the field groups
func (e *GroupListOutput) SetGroups(groups []*GroupInformation) {
	e.Groups = groups
}

// GetNextCursor returns the value for the field nextCursor
func (e *GroupListOutput) GetNextCursor() *string {
	return e.NextCursor
}

// SetNextCursor sets the value for the field nextCursor
func (e *GroupListOutput) SetNextCursor(nextCursor *string) {
	e.NextCursor = nextCursor
}

// StructPath returns StructPath
func (e *GroupListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathGroupListOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *GroupListOutput) InitializeDefaults() {
}

// groupListOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type groupListOutputAlias GroupListOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *GroupListOutput) UnmarshalJSON(data []byte) error {
	var alias groupListOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*GroupListOutput)(&alias)).InitializeDefaults()
	*e = GroupListOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e GroupListOutput) MarshalJSON() ([]byte, error) {
	alias := groupListOutputAlias(e)
	return json.Marshal(alias)
}

// NewGroupMemberAddInput creates a new GroupMemberAddInput
func NewGroupMemberAddInput() *GroupMemberAddInput {
	s := &GroupMemberAddInput{}
	s.InitializeDefaults()
	return s
}

// GroupMemberAddInput struct
type GroupMemberAddInput struct {
	GroupName string `json:"groupName,omitempty" yaml:"groupName,omitempty"`
	Username  string `json:"username,omitempty" yaml:"username,omitempty"`
}

// GetGroupName returns the value for the field groupName
func (e *GroupMemberAddInput) GetGroupName() string {
	return e.GroupName
}

// SetGroupName sets the value for the field groupName
func (e *GroupMemberAddInput) SetGroupName(groupName string) {
	e.GroupName = groupName
}

// GetUsername returns the value for the field username
func (e *GroupMemberAddInput) GetUsername() string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *GroupMemberAddInput) SetUsername(username string) {
	e.Username = username
}

// StructPath returns StructPath
func (e *GroupMemberAddInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathGroupMemberAddInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *GroupMemberAddInput) InitializeDefaults() {
}

// groupMemberAddInputAlias is defined to help pre and post JSON marshaling without recursive loops
type groupMemberAddInputAlias GroupMemberAddInput

// UnmarshalJSON implements json.Unmarshaler
func (e *GroupMemberAddInput) UnmarshalJSON(data []byte) error {
	var alias groupMemberAddInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*GroupMemberAddInput)(&alias)).InitializeDefaults()
	*e = GroupMemberAddInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e GroupMemberAddInput) MarshalJSON() ([]byte, error) {
	alias := groupMemberAddInputAlias(e)
	return json.Marshal(alias)
}

// NewGroupMemberAddOutput creates a new GroupMemberAddOutput
func NewGroupMemberAddOutput() *GroupMemberAddOutput {
	s := &GroupMemberAddOutput{}
	s.InitializeDefaults()
	return s
}

// GroupMemberAddOutput struct
type GroupMemberAddOutput struct {
}

// StructPath returns StructPath
func (e *GroupMemberAddOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathGroupMemberAddOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *GroupMemberAddOutput) InitializeDefaults() {
}

// groupMemberAddOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type groupMemberAddOutputAlias GroupMemberAddOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *GroupMemberAddOutput) UnmarshalJSON(data []byte) error {
	var alias groupMemberAddOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*GroupMemberAddOutput)(&alias)).InitializeDefaults()
	*e = GroupMemberAddOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e GroupMemberAddOutput) MarshalJSON() ([]byte, error) {
	alias := groupMemberAddOutputAlias(e)
	return json.Marshal(alias)
}

// NewGroupMemberRemoveInput creates a new GroupMemberRemoveInput
func NewGroupMemberRemoveInput() *GroupMemberRemoveInput {
	s := &GroupMemberRemoveInput{}
	s.InitializeDefaults()
	return s
}

// GroupMemberRemoveInput struct
type GroupMemberRemoveInput struct {
	GroupName string `json:"groupName,omitempty" yaml:"groupName,omitempty"`
	Username  string `json:"username,omitempty" yaml:"username,omitempty"`
}

// GetGroupName returns the value for the field groupName
func (e *GroupMemberRemoveInput) GetGroupName() string {
	return e.GroupName
}

// SetGroupName sets the value for the field groupName
func (e *GroupMemberRemoveInput) SetGroupName(groupName string) {
	e.GroupName = groupName
}

// GetUsername returns the value for the field username
func (e *GroupMemberRemoveInput) GetUsername() string {
	return e.Username
}

// SetUsername sets the value for the field username
func (e *GroupMemberRemoveInput) SetUsername(username string) {
	e.Username = username
}

// StructPath returns StructPath
func (e *GroupMemberRemoveInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathGroupMemberRemoveInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *GroupMemberRemoveInput) InitializeDefaults() {
}

// groupMemberRemoveInputAlias is defined to help pre and post JSON marshaling without recursive loops
type groupMemberRemoveInputAlias GroupMemberRemoveInput

// UnmarshalJSON implements json.Unmarshaler
func (e *GroupMemberRemoveInput) UnmarshalJSON(data []byte) error {
	var alias groupMemberRemoveInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*GroupMemberRemoveInput)(&alias)).InitializeDefaults()
	*e = GroupMemberRemoveInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e GroupMemberRemoveInput) MarshalJSON() ([]byte, error) {
	alias := groupMemberRemoveInputAlias(e)
	return json.Marshal(alias)
}

// NewGroupMemberRemoveOutput creates a new GroupMemberRemoveOutput
func NewGroupMemberRemoveOutput() *GroupMemberRemoveOutput {
	s := &GroupMemberRemoveOutput{}
	s.InitializeDefaults()
	return s
}

// GroupMemberRemoveOutput struct
type GroupMemberRemoveOutput struct {
}

// StructPath returns StructPath
func (e *GroupMemberRemoveOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathGroupMemberRemoveOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *GroupMemberRemoveOutput) InitializeDefaults() {
}

// groupMemberRemoveOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type groupMemberRemoveOutputAlias GroupMemberRemoveOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *GroupMemberRemoveOutput) UnmarshalJSON(data []byte) error {
	var alias groupMemberRemoveOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*GroupMemberRemoveOutput)(&alias)).InitializeDefaults()
	*e = GroupMemberRemoveOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e GroupMemberRemoveOutput) MarshalJSON() ([]byte, error) {
	alias := groupMemberRemoveOutputAlias(e)
	return json.Marshal(alias)
}

// NewGroupMemberListInput creates a new GroupMemberListInput
func NewGroupMemberListInput() *GroupMemberListInput {
	s := &GroupMemberListInput{}
	s.InitializeDefaults()
	return s
}

// GroupMemberListInput struct
type GroupMemberListInput struct {
	GroupName string `json:"groupName,omitempty" yaml:"groupName,omitempty"`
}

// GetGroupName returns the value for the field groupName
func (e *GroupMemberListInput) GetGroupName() string {
	return e.GroupName
}

// SetGroupName sets the value for the field groupName
func (e *GroupMemberListInput) SetGroupName(groupName string) {
	e.GroupName = groupName
}

// StructPath returns StructPath
func (e *GroupMemberListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathGroupMemberListInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *GroupMemberListInput) InitializeDefaults() {
}

// groupMemberListInputAlias is defined to help pre and post JSON marshaling without recursive loops
type groupMemberListInputAlias GroupMemberListInput

// UnmarshalJSON implements json.Unmarshaler
func (e *GroupMemberListInput) UnmarshalJSON(data []byte) error {
	var alias groupMemberListInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*GroupMemberListInput)(&alias)).InitializeDefaults()
	*e = GroupMemberListInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e GroupMemberListInput) MarshalJSON() ([]byte, error) {
	alias := groupMemberListInputAlias(e)
	return json.Marshal(alias)
}

// NewGroupMemberListOutput creates a new GroupMemberListOutput
func NewGroupMemberListOutput() *GroupMemberListOutput {
	s := &GroupMemberListOutput{}
	s.InitializeDefaults()
	return s
}

// GroupMemberListOutput struct
type GroupMemberListOutput struct {
	Usernames []string `json:"usernames,omitempty" yaml:"usernames,omitempty"`
}

// GetUsernames returns the value for the field usernames
func (e *GroupMemberListOutput) GetUsernames() []string {
	return e.Usernames
}

// SetUsernames sets the value for the field usernames
func (e *GroupMemberListOutput) SetUsernames(usernames []string) {
	e.Usernames = usernames
}

// StructPath returns StructPath
func (e *GroupMemberListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathGroupMemberListOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *GroupMemberListOutput) InitializeDefaults() {
}

// groupMemberListOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type groupMemberListOutputAlias GroupMemberListOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *GroupMemberListOutput) UnmarshalJSON(data []byte) error {
	var alias groupMemberListOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*GroupMemberListOutput)(&alias)).InitializeDefaults()
	*e = GroupMemberListOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e GroupMemberListOutput) MarshalJSON() ([]byte, error) {
	alias := groupMemberListOutputAlias(e)
	return json.Marshal(alias)
}

// NewGroupIdentityPolicyAttachInput creates a new GroupIdentityPolicyAttachInput
func NewGroupIdentityPolicyAttachInput() *GroupIdentityPolicyAttachInput {
	s := &GroupIdentityPolicyAttachInput{}
	s.InitializeDefaults()
	return s
}

// GroupIdentityPolicyAttachInput struct
type GroupIdentityPolicyAttachInput struct {
	GroupName  string `json:"groupName,omitempty" yaml:"groupName,omitempty"`
	PolicyName string `json:"policyName,omitempty" yaml:"policyName,omitempty"`
}

// GetGroupName returns the value for the field groupName
func (e *GroupIdentityPolicyAttachInput) GetGroupName() string {
	return e.GroupName
}

// SetGroupName sets the value for the field groupName
func (e *GroupIdentityPolicyAttachInput) SetGroupName(groupName string) {
	e.GroupName = groupName
}

// GetPolicyName returns the value for the field policyName
func (e *GroupIdentityPolicyAttachInput) GetPolicyName() string {
	return e.PolicyName
}

// SetPolicyName sets the value for the field policyName
func (e *GroupIdentityPolicyAttachInput) SetPolicyName(policyName string) {
	e.PolicyName = policyName
}

// StructPath returns StructPath
func (e *GroupIdentityPolicyAttachInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathGroupIdentityPolicyAttachInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *GroupIdentityPolicyAttachInput) InitializeDefaults() {
}

// groupIdentityPolicyAttachInputAlias is defined to help pre and post JSON marshaling without recursive loops
type groupIdentityPolicyAttachInputAlias GroupIdentityPolicyAttachInput

// UnmarshalJSON implements json.Unmarshaler
func (e *GroupIdentityPolicyAttachInput) UnmarshalJSON(data []byte) error {
	var alias groupIdentityPolicyAttachInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*GroupIdentityPolicyAttachInput)(&alias)).InitializeDefaults()
	*e = GroupIdentityPolicyAttachInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e GroupIdentityPolicyAttachInput) MarshalJSON() ([]byte, error) {
	alias := groupIdentityPolicyAttachInputAlias(e)
	return json.Marshal(alias)
}

// NewGroupIdentityPolicyAttachOutput creates a new GroupIdentityPolicyAttachOutput
func NewGroupIdentityPolicyAttachOutput() *GroupIdentityPolicyAttachOutput {
	s := &GroupIdentityPolicyAttachOutput{}
	s.InitializeDefaults()
	return s
}

// GroupIdentityPolicyAttachOutput struct
type GroupIdentityPolicyAttachOutput struct {
}

// StructPath returns StructPath
func (e *GroupIdentityPolicyAttachOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathGroupIdentityPolicyAttachOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *GroupIdentityPolicyAttachOutput) InitializeDefaults() {
}

// groupIdentityPolicyAttachOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type groupIdentityPolicyAttachOutputAlias GroupIdentityPolicyAttachOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *GroupIdentityPolicyAttachOutput) UnmarshalJSON(data []byte) error {
	var alias groupIdentityPolicyAttachOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*GroupIdentityPolicyAttachOutput)(&alias)).InitializeDefaults()
	*e = GroupIdentityPolicyAttachOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e GroupIdentityPolicyAttachOutput) MarshalJSON() ([]byte, error) {
	alias := groupIdentityPolicyAttachOutputAlias(e)
	return json.Marshal(alias)
}

// NewGroupIdentityPolicyDetachInput creates a new GroupIdentityPolicyDetachInput
func NewGroupIdentityPolicyDetachInput() *GroupIdentityPolicyDetachInput {
	s := &GroupIdentityPolicyDetachInput{}
	s.InitializeDefaults()
	return s
}

// GroupIdentityPolicyDetachInput struct
type GroupIdentityPolicyDetachInput struct {
	GroupName  string `json:"groupName,omitempty" yaml:"groupName,omitempty"`
	PolicyName string `json:"policyName,omitempty" yaml:"policyName,omitempty"`
}

// GetGroupName returns the value for the field groupName
func (e *GroupIdentityPolicyDetachInput) GetGroupName() string {
	return e.GroupName
}

// SetGroupName sets the value for the field groupName
func (e *GroupIdentityPolicyDetachInput) SetGroupName(groupName string) {
	e.GroupName = groupName
}

// GetPolicyName returns the value for the field policyName
func (e *GroupIdentityPolicyDetachInput) GetPolicyName() string {
	return e.PolicyName
}

// SetPolicyName sets the value for the field policyName
func (e *GroupIdentityPolicyDetachInput) SetPolicyName(policyName string) {
	e.PolicyName = policyName
}

// StructPath returns StructPath
func (e *GroupIdentityPolicyDetachInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathGroupIdentityPolicyDetachInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *GroupIdentityPolicyDetachInput) InitializeDefaults() {
}

// groupIdentityPolicyDetachInputAlias is defined to help pre and post JSON marshaling without recursive loops
type groupIdentityPolicyDetachInputAlias GroupIdentityPolicyDetachInput

// UnmarshalJSON implements json.Unmarshaler
func (e *GroupIdentityPolicyDetachInput) UnmarshalJSON(data []byte) error {
	var alias groupIdentityPolicyDetachInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*GroupIdentityPolicyDetachInput)(&alias)).InitializeDefaults()
	*e = GroupIdentityPolicyDetachInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e GroupIdentityPolicyDetachInput) MarshalJSON() ([]byte, error) {
	alias := groupIdentityPolicyDetachInputAlias(e)
	return json.Marshal(alias)
}

// NewGroupIdentityPolicyDetachOutput creates a new GroupIdentityPolicyDetachOutput
func NewGroupIdentityPolicyDetachOutput() *GroupIdentityPolicyDetachOutput {
	s := &GroupIdentityPolicyDetachOutput{}
	s.InitializeDefaults()
	return s
}

// GroupIdentityPolicyDetachOutput struct
type GroupIdentityPolicyDetachOutput struct {
}

// StructPath returns StructPath
func (e *GroupIdentityPolicyDetachOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathGroupIdentityPolicyDetachOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *GroupIdentityPolicyDetachOutput) InitializeDefaults() {
}

// groupIdentityPolicyDetachOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type groupIdentityPolicyDetachOutputAlias GroupIdentityPolicyDetachOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *GroupIdentityPolicyDetachOutput) UnmarshalJSON(data []byte) error {
	var alias groupIdentityPolicyDetachOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*GroupIdentityPolicyDetachOutput)(&alias)).InitializeDefaults()
	*e = GroupIdentityPolicyDetachOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e GroupIdentityPolicyDetachOutput) MarshalJSON() ([]byte, error) {
	alias := groupIdentityPolicyDetachOutputAlias(e)
	return json.Marshal(alias)
}

// NewGroupIdentityPolicyListInput creates a new GroupIdentityPolicyListInput
func NewGroupIdentityPolicyListInput() *GroupIdentityPolicyListInput {
	s := &GroupIdentityPolicyListInput{}
	s.InitializeDefaults()
	return s
}

// GroupIdentityPolicyListInput struct
type GroupIdentityPolicyListInput struct {
	GroupName string `json:"groupName,omitempty" yaml:"groupName,omitempty"`
}

// GetGroupName returns the value for the field groupName
func (e *GroupIdentityPolicyListInput) GetGroupName() string {
	return e.GroupName
}

// SetGroupName sets the value for the field groupName
func (e *GroupIdentityPolicyListInput) SetGroupName(groupName string) {
	e.GroupName = groupName
}

// StructPath returns StructPath
func (e *GroupIdentityPolicyListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathGroupIdentityPolicyListInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *GroupIdentityPolicyListInput) InitializeDefaults() {
}

// groupIdentityPolicyListInputAlias is defined to help pre and post JSON marshaling without recursive loops
type groupIdentityPolicyListInputAlias GroupIdentityPolicyListInput

// UnmarshalJSON implements json.Unmarshaler
func (e *GroupIdentityPolicyListInput) UnmarshalJSON(data []byte) error {
	var alias groupIdentityPolicyListInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*GroupIdentityPolicyListInput)(&alias)).InitializeDefaults()
	*e = GroupIdentityPolicyListInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e GroupIdentityPolicyListInput) MarshalJSON() ([]byte, error) {
	alias := groupIdentityPolicyListInputAlias(e)
	return json.Marshal(alias)
}

// NewGroupIdentityPolicyListOutput creates a new GroupIdentityPolicyListOutput
func NewGroupIdentityPolicyListOutput() *GroupIdentityPolicyListOutput {
	s := &GroupIdentityPolicyListOutput{}
	s.InitializeDefaults()
	return s
}

// GroupIdentityPolicyListOutput struct
type GroupIdentityPolicyListOutput struct {
	Attachments []*IdentityPolicyAttachmentInfo `json:"attachments,omitempty" yaml:"attachments,omitempty"`
}

// GetAttachments returns the value for the field attachments
func (e *GroupIdentityPolicyListOutput) GetAttachments() []*IdentityPolicyAttachmentInfo {
	return e.Attachments
}

// SetAttachments sets the value for the field attachments
func (e *GroupIdentityPolicyListOutput) SetAttachments(attachments []*IdentityPolicyAttachmentInfo) {
	e.Attachments = attachments
}

// StructPath returns StructPath
func (e *GroupIdentityPolicyListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathGroupIdentityPolicyListOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *GroupIdentityPolicyListOutput) InitializeDefaults() {
}

// groupIdentityPolicyListOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type groupIdentityPolicyListOutputAlias GroupIdentityPolicyListOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *GroupIdentityPolicyListOutput) UnmarshalJSON(data []byte) error {
	var alias groupIdentityPolicyListOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*GroupIdentityPolicyListOutput)(&alias)).InitializeDefaults()
	*e = GroupIdentityPolicyListOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e GroupIdentityPolicyListOutput) MarshalJSON() ([]byte, error) {
	alias := groupIdentityPolicyListOutputAlias(e)
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
	DefinitionHuJSON []byte `json:"definitionHuJSON,omitempty" yaml:"definitionHuJSON,omitempty"`
}

// GetDefinitionHuJSON returns the value for the field definitionHuJSON
func (e *IdentityPolicyCreateInput) GetDefinitionHuJSON() []byte {
	return e.DefinitionHuJSON
}

// SetDefinitionHuJSON sets the value for the field definitionHuJSON
func (e *IdentityPolicyCreateInput) SetDefinitionHuJSON(definitionHuJSON []byte) {
	e.DefinitionHuJSON = definitionHuJSON
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
	DefinitionHuJSON []byte `json:"definitionHuJSON,omitempty" yaml:"definitionHuJSON,omitempty"`
}

// GetDefinitionHuJSON returns the value for the field definitionHuJSON
func (e *IdentityPolicyUpdateInput) GetDefinitionHuJSON() []byte {
	return e.DefinitionHuJSON
}

// SetDefinitionHuJSON sets the value for the field definitionHuJSON
func (e *IdentityPolicyUpdateInput) SetDefinitionHuJSON(definitionHuJSON []byte) {
	e.DefinitionHuJSON = definitionHuJSON
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

// NewIdentityPolicyAttachmentListInput creates a new IdentityPolicyAttachmentListInput
func NewIdentityPolicyAttachmentListInput() *IdentityPolicyAttachmentListInput {
	s := &IdentityPolicyAttachmentListInput{}
	s.InitializeDefaults()
	return s
}

// IdentityPolicyAttachmentListInput struct
type IdentityPolicyAttachmentListInput struct {
	PolicyName string `json:"policyName,omitempty" yaml:"policyName,omitempty"`
}

// GetPolicyName returns the value for the field policyName
func (e *IdentityPolicyAttachmentListInput) GetPolicyName() string {
	return e.PolicyName
}

// SetPolicyName sets the value for the field policyName
func (e *IdentityPolicyAttachmentListInput) SetPolicyName(policyName string) {
	e.PolicyName = policyName
}

// StructPath returns StructPath
func (e *IdentityPolicyAttachmentListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyAttachmentListInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *IdentityPolicyAttachmentListInput) InitializeDefaults() {
}

// identityPolicyAttachmentListInputAlias is defined to help pre and post JSON marshaling without recursive loops
type identityPolicyAttachmentListInputAlias IdentityPolicyAttachmentListInput

// UnmarshalJSON implements json.Unmarshaler
func (e *IdentityPolicyAttachmentListInput) UnmarshalJSON(data []byte) error {
	var alias identityPolicyAttachmentListInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*IdentityPolicyAttachmentListInput)(&alias)).InitializeDefaults()
	*e = IdentityPolicyAttachmentListInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e IdentityPolicyAttachmentListInput) MarshalJSON() ([]byte, error) {
	alias := identityPolicyAttachmentListInputAlias(e)
	return json.Marshal(alias)
}

// NewIdentityPolicyAttachmentListOutput creates a new IdentityPolicyAttachmentListOutput
func NewIdentityPolicyAttachmentListOutput() *IdentityPolicyAttachmentListOutput {
	s := &IdentityPolicyAttachmentListOutput{}
	s.InitializeDefaults()
	return s
}

// IdentityPolicyAttachmentListOutput struct
type IdentityPolicyAttachmentListOutput struct {
	Policies []*IdentityPolicyAttachment `json:"policies,omitempty" yaml:"policies,omitempty"`
}

// GetPolicies returns the value for the field policies
func (e *IdentityPolicyAttachmentListOutput) GetPolicies() []*IdentityPolicyAttachment {
	return e.Policies
}

// SetPolicies sets the value for the field policies
func (e *IdentityPolicyAttachmentListOutput) SetPolicies(policies []*IdentityPolicyAttachment) {
	e.Policies = policies
}

// StructPath returns StructPath
func (e *IdentityPolicyAttachmentListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathIdentityPolicyAttachmentListOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *IdentityPolicyAttachmentListOutput) InitializeDefaults() {
}

// identityPolicyAttachmentListOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type identityPolicyAttachmentListOutputAlias IdentityPolicyAttachmentListOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *IdentityPolicyAttachmentListOutput) UnmarshalJSON(data []byte) error {
	var alias identityPolicyAttachmentListOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*IdentityPolicyAttachmentListOutput)(&alias)).InitializeDefaults()
	*e = IdentityPolicyAttachmentListOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e IdentityPolicyAttachmentListOutput) MarshalJSON() ([]byte, error) {
	alias := identityPolicyAttachmentListOutputAlias(e)
	return json.Marshal(alias)
}

// NewTrustPolicyCreateInput creates a new TrustPolicyCreateInput
func NewTrustPolicyCreateInput() *TrustPolicyCreateInput {
	s := &TrustPolicyCreateInput{}
	s.InitializeDefaults()
	return s
}

// TrustPolicyCreateInput struct
type TrustPolicyCreateInput struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// DRN of the trusted principal, e.g. "iam:OIDCProvider(github-actions)"
	Principal  string                  `json:"principal,omitempty" yaml:"principal,omitempty"`
	Statements []*TrustPolicyStatement `json:"statements,omitempty" yaml:"statements,omitempty"`
}

// GetName returns the value for the field name
func (e *TrustPolicyCreateInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *TrustPolicyCreateInput) SetName(name string) {
	e.Name = name
}

// GetPrincipal returns the value for the field principal
func (e *TrustPolicyCreateInput) GetPrincipal() string {
	return e.Principal
}

// SetPrincipal sets the value for the field principal
func (e *TrustPolicyCreateInput) SetPrincipal(principal string) {
	e.Principal = principal
}

// GetStatements returns the value for the field statements
func (e *TrustPolicyCreateInput) GetStatements() []*TrustPolicyStatement {
	return e.Statements
}

// SetStatements sets the value for the field statements
func (e *TrustPolicyCreateInput) SetStatements(statements []*TrustPolicyStatement) {
	e.Statements = statements
}

// StructPath returns StructPath
func (e *TrustPolicyCreateInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathTrustPolicyCreateInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *TrustPolicyCreateInput) InitializeDefaults() {
}

// trustPolicyCreateInputAlias is defined to help pre and post JSON marshaling without recursive loops
type trustPolicyCreateInputAlias TrustPolicyCreateInput

// UnmarshalJSON implements json.Unmarshaler
func (e *TrustPolicyCreateInput) UnmarshalJSON(data []byte) error {
	var alias trustPolicyCreateInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*TrustPolicyCreateInput)(&alias)).InitializeDefaults()
	*e = TrustPolicyCreateInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e TrustPolicyCreateInput) MarshalJSON() ([]byte, error) {
	alias := trustPolicyCreateInputAlias(e)
	return json.Marshal(alias)
}

// NewTrustPolicyCreateOutput creates a new TrustPolicyCreateOutput
func NewTrustPolicyCreateOutput() *TrustPolicyCreateOutput {
	s := &TrustPolicyCreateOutput{}
	s.InitializeDefaults()
	return s
}

// TrustPolicyCreateOutput struct
type TrustPolicyCreateOutput struct {
	TrustPolicy *TrustPolicy `json:"trustPolicy,omitempty" yaml:"trustPolicy,omitempty"`
}

// GetTrustPolicy returns the value for the field trustPolicy
func (e *TrustPolicyCreateOutput) GetTrustPolicy() *TrustPolicy {
	return e.TrustPolicy
}

// SetTrustPolicy sets the value for the field trustPolicy
func (e *TrustPolicyCreateOutput) SetTrustPolicy(trustPolicy *TrustPolicy) {
	e.TrustPolicy = trustPolicy
}

// StructPath returns StructPath
func (e *TrustPolicyCreateOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathTrustPolicyCreateOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *TrustPolicyCreateOutput) InitializeDefaults() {
}

// trustPolicyCreateOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type trustPolicyCreateOutputAlias TrustPolicyCreateOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *TrustPolicyCreateOutput) UnmarshalJSON(data []byte) error {
	var alias trustPolicyCreateOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*TrustPolicyCreateOutput)(&alias)).InitializeDefaults()
	*e = TrustPolicyCreateOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e TrustPolicyCreateOutput) MarshalJSON() ([]byte, error) {
	alias := trustPolicyCreateOutputAlias(e)
	return json.Marshal(alias)
}

// NewTrustPolicyRetrieveInput creates a new TrustPolicyRetrieveInput
func NewTrustPolicyRetrieveInput() *TrustPolicyRetrieveInput {
	s := &TrustPolicyRetrieveInput{}
	s.InitializeDefaults()
	return s
}

// TrustPolicyRetrieveInput struct
type TrustPolicyRetrieveInput struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetName returns the value for the field name
func (e *TrustPolicyRetrieveInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *TrustPolicyRetrieveInput) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *TrustPolicyRetrieveInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathTrustPolicyRetrieveInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *TrustPolicyRetrieveInput) InitializeDefaults() {
}

// trustPolicyRetrieveInputAlias is defined to help pre and post JSON marshaling without recursive loops
type trustPolicyRetrieveInputAlias TrustPolicyRetrieveInput

// UnmarshalJSON implements json.Unmarshaler
func (e *TrustPolicyRetrieveInput) UnmarshalJSON(data []byte) error {
	var alias trustPolicyRetrieveInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*TrustPolicyRetrieveInput)(&alias)).InitializeDefaults()
	*e = TrustPolicyRetrieveInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e TrustPolicyRetrieveInput) MarshalJSON() ([]byte, error) {
	alias := trustPolicyRetrieveInputAlias(e)
	return json.Marshal(alias)
}

// NewTrustPolicyRetrieveOutput creates a new TrustPolicyRetrieveOutput
func NewTrustPolicyRetrieveOutput() *TrustPolicyRetrieveOutput {
	s := &TrustPolicyRetrieveOutput{}
	s.InitializeDefaults()
	return s
}

// TrustPolicyRetrieveOutput struct
type TrustPolicyRetrieveOutput struct {
	TrustPolicy *TrustPolicy `json:"trustPolicy,omitempty" yaml:"trustPolicy,omitempty"`
}

// GetTrustPolicy returns the value for the field trustPolicy
func (e *TrustPolicyRetrieveOutput) GetTrustPolicy() *TrustPolicy {
	return e.TrustPolicy
}

// SetTrustPolicy sets the value for the field trustPolicy
func (e *TrustPolicyRetrieveOutput) SetTrustPolicy(trustPolicy *TrustPolicy) {
	e.TrustPolicy = trustPolicy
}

// StructPath returns StructPath
func (e *TrustPolicyRetrieveOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathTrustPolicyRetrieveOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *TrustPolicyRetrieveOutput) InitializeDefaults() {
}

// trustPolicyRetrieveOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type trustPolicyRetrieveOutputAlias TrustPolicyRetrieveOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *TrustPolicyRetrieveOutput) UnmarshalJSON(data []byte) error {
	var alias trustPolicyRetrieveOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*TrustPolicyRetrieveOutput)(&alias)).InitializeDefaults()
	*e = TrustPolicyRetrieveOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e TrustPolicyRetrieveOutput) MarshalJSON() ([]byte, error) {
	alias := trustPolicyRetrieveOutputAlias(e)
	return json.Marshal(alias)
}

// NewTrustPolicyListInput creates a new TrustPolicyListInput
func NewTrustPolicyListInput() *TrustPolicyListInput {
	s := &TrustPolicyListInput{}
	s.InitializeDefaults()
	return s
}

// TrustPolicyListInput struct
type TrustPolicyListInput struct {
}

// StructPath returns StructPath
func (e *TrustPolicyListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathTrustPolicyListInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *TrustPolicyListInput) InitializeDefaults() {
}

// trustPolicyListInputAlias is defined to help pre and post JSON marshaling without recursive loops
type trustPolicyListInputAlias TrustPolicyListInput

// UnmarshalJSON implements json.Unmarshaler
func (e *TrustPolicyListInput) UnmarshalJSON(data []byte) error {
	var alias trustPolicyListInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*TrustPolicyListInput)(&alias)).InitializeDefaults()
	*e = TrustPolicyListInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e TrustPolicyListInput) MarshalJSON() ([]byte, error) {
	alias := trustPolicyListInputAlias(e)
	return json.Marshal(alias)
}

// NewTrustPolicyListOutput creates a new TrustPolicyListOutput
func NewTrustPolicyListOutput() *TrustPolicyListOutput {
	s := &TrustPolicyListOutput{}
	s.InitializeDefaults()
	return s
}

// TrustPolicyListOutput struct
type TrustPolicyListOutput struct {
	TrustPolicies []*TrustPolicy `json:"trustPolicies,omitempty" yaml:"trustPolicies,omitempty"`
}

// GetTrustPolicies returns the value for the field trustPolicies
func (e *TrustPolicyListOutput) GetTrustPolicies() []*TrustPolicy {
	return e.TrustPolicies
}

// SetTrustPolicies sets the value for the field trustPolicies
func (e *TrustPolicyListOutput) SetTrustPolicies(trustPolicies []*TrustPolicy) {
	e.TrustPolicies = trustPolicies
}

// StructPath returns StructPath
func (e *TrustPolicyListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathTrustPolicyListOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *TrustPolicyListOutput) InitializeDefaults() {
}

// trustPolicyListOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type trustPolicyListOutputAlias TrustPolicyListOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *TrustPolicyListOutput) UnmarshalJSON(data []byte) error {
	var alias trustPolicyListOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*TrustPolicyListOutput)(&alias)).InitializeDefaults()
	*e = TrustPolicyListOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e TrustPolicyListOutput) MarshalJSON() ([]byte, error) {
	alias := trustPolicyListOutputAlias(e)
	return json.Marshal(alias)
}

// NewTrustPolicyUpdateInput creates a new TrustPolicyUpdateInput
func NewTrustPolicyUpdateInput() *TrustPolicyUpdateInput {
	s := &TrustPolicyUpdateInput{}
	s.InitializeDefaults()
	return s
}

// TrustPolicyUpdateInput struct
type TrustPolicyUpdateInput struct {
	Name       string                  `json:"name,omitempty" yaml:"name,omitempty"`
	Principal  string                  `json:"principal,omitempty" yaml:"principal,omitempty"`
	Statements []*TrustPolicyStatement `json:"statements,omitempty" yaml:"statements,omitempty"`
}

// GetName returns the value for the field name
func (e *TrustPolicyUpdateInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *TrustPolicyUpdateInput) SetName(name string) {
	e.Name = name
}

// GetPrincipal returns the value for the field principal
func (e *TrustPolicyUpdateInput) GetPrincipal() string {
	return e.Principal
}

// SetPrincipal sets the value for the field principal
func (e *TrustPolicyUpdateInput) SetPrincipal(principal string) {
	e.Principal = principal
}

// GetStatements returns the value for the field statements
func (e *TrustPolicyUpdateInput) GetStatements() []*TrustPolicyStatement {
	return e.Statements
}

// SetStatements sets the value for the field statements
func (e *TrustPolicyUpdateInput) SetStatements(statements []*TrustPolicyStatement) {
	e.Statements = statements
}

// StructPath returns StructPath
func (e *TrustPolicyUpdateInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathTrustPolicyUpdateInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *TrustPolicyUpdateInput) InitializeDefaults() {
}

// trustPolicyUpdateInputAlias is defined to help pre and post JSON marshaling without recursive loops
type trustPolicyUpdateInputAlias TrustPolicyUpdateInput

// UnmarshalJSON implements json.Unmarshaler
func (e *TrustPolicyUpdateInput) UnmarshalJSON(data []byte) error {
	var alias trustPolicyUpdateInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*TrustPolicyUpdateInput)(&alias)).InitializeDefaults()
	*e = TrustPolicyUpdateInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e TrustPolicyUpdateInput) MarshalJSON() ([]byte, error) {
	alias := trustPolicyUpdateInputAlias(e)
	return json.Marshal(alias)
}

// NewTrustPolicyUpdateOutput creates a new TrustPolicyUpdateOutput
func NewTrustPolicyUpdateOutput() *TrustPolicyUpdateOutput {
	s := &TrustPolicyUpdateOutput{}
	s.InitializeDefaults()
	return s
}

// TrustPolicyUpdateOutput struct
type TrustPolicyUpdateOutput struct {
	TrustPolicy *TrustPolicy `json:"trustPolicy,omitempty" yaml:"trustPolicy,omitempty"`
}

// GetTrustPolicy returns the value for the field trustPolicy
func (e *TrustPolicyUpdateOutput) GetTrustPolicy() *TrustPolicy {
	return e.TrustPolicy
}

// SetTrustPolicy sets the value for the field trustPolicy
func (e *TrustPolicyUpdateOutput) SetTrustPolicy(trustPolicy *TrustPolicy) {
	e.TrustPolicy = trustPolicy
}

// StructPath returns StructPath
func (e *TrustPolicyUpdateOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathTrustPolicyUpdateOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *TrustPolicyUpdateOutput) InitializeDefaults() {
}

// trustPolicyUpdateOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type trustPolicyUpdateOutputAlias TrustPolicyUpdateOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *TrustPolicyUpdateOutput) UnmarshalJSON(data []byte) error {
	var alias trustPolicyUpdateOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*TrustPolicyUpdateOutput)(&alias)).InitializeDefaults()
	*e = TrustPolicyUpdateOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e TrustPolicyUpdateOutput) MarshalJSON() ([]byte, error) {
	alias := trustPolicyUpdateOutputAlias(e)
	return json.Marshal(alias)
}

// NewTrustPolicyDestroyInput creates a new TrustPolicyDestroyInput
func NewTrustPolicyDestroyInput() *TrustPolicyDestroyInput {
	s := &TrustPolicyDestroyInput{}
	s.InitializeDefaults()
	return s
}

// TrustPolicyDestroyInput struct
type TrustPolicyDestroyInput struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetName returns the value for the field name
func (e *TrustPolicyDestroyInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *TrustPolicyDestroyInput) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *TrustPolicyDestroyInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathTrustPolicyDestroyInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *TrustPolicyDestroyInput) InitializeDefaults() {
}

// trustPolicyDestroyInputAlias is defined to help pre and post JSON marshaling without recursive loops
type trustPolicyDestroyInputAlias TrustPolicyDestroyInput

// UnmarshalJSON implements json.Unmarshaler
func (e *TrustPolicyDestroyInput) UnmarshalJSON(data []byte) error {
	var alias trustPolicyDestroyInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*TrustPolicyDestroyInput)(&alias)).InitializeDefaults()
	*e = TrustPolicyDestroyInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e TrustPolicyDestroyInput) MarshalJSON() ([]byte, error) {
	alias := trustPolicyDestroyInputAlias(e)
	return json.Marshal(alias)
}

// NewTrustPolicyDestroyOutput creates a new TrustPolicyDestroyOutput
func NewTrustPolicyDestroyOutput() *TrustPolicyDestroyOutput {
	s := &TrustPolicyDestroyOutput{}
	s.InitializeDefaults()
	return s
}

// TrustPolicyDestroyOutput struct
type TrustPolicyDestroyOutput struct {
}

// StructPath returns StructPath
func (e *TrustPolicyDestroyOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathTrustPolicyDestroyOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *TrustPolicyDestroyOutput) InitializeDefaults() {
}

// trustPolicyDestroyOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type trustPolicyDestroyOutputAlias TrustPolicyDestroyOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *TrustPolicyDestroyOutput) UnmarshalJSON(data []byte) error {
	var alias trustPolicyDestroyOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*TrustPolicyDestroyOutput)(&alias)).InitializeDefaults()
	*e = TrustPolicyDestroyOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e TrustPolicyDestroyOutput) MarshalJSON() ([]byte, error) {
	alias := trustPolicyDestroyOutputAlias(e)
	return json.Marshal(alias)
}

// NewTrustPolicyAttachmentListInput creates a new TrustPolicyAttachmentListInput
func NewTrustPolicyAttachmentListInput() *TrustPolicyAttachmentListInput {
	s := &TrustPolicyAttachmentListInput{}
	s.InitializeDefaults()
	return s
}

// TrustPolicyAttachmentListInput struct
type TrustPolicyAttachmentListInput struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// GetName returns the value for the field name
func (e *TrustPolicyAttachmentListInput) GetName() string {
	return e.Name
}

// SetName sets the value for the field name
func (e *TrustPolicyAttachmentListInput) SetName(name string) {
	e.Name = name
}

// StructPath returns StructPath
func (e *TrustPolicyAttachmentListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathTrustPolicyAttachmentListInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *TrustPolicyAttachmentListInput) InitializeDefaults() {
}

// trustPolicyAttachmentListInputAlias is defined to help pre and post JSON marshaling without recursive loops
type trustPolicyAttachmentListInputAlias TrustPolicyAttachmentListInput

// UnmarshalJSON implements json.Unmarshaler
func (e *TrustPolicyAttachmentListInput) UnmarshalJSON(data []byte) error {
	var alias trustPolicyAttachmentListInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*TrustPolicyAttachmentListInput)(&alias)).InitializeDefaults()
	*e = TrustPolicyAttachmentListInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e TrustPolicyAttachmentListInput) MarshalJSON() ([]byte, error) {
	alias := trustPolicyAttachmentListInputAlias(e)
	return json.Marshal(alias)
}

// NewTrustPolicyAttachmentListOutput creates a new TrustPolicyAttachmentListOutput
func NewTrustPolicyAttachmentListOutput() *TrustPolicyAttachmentListOutput {
	s := &TrustPolicyAttachmentListOutput{}
	s.InitializeDefaults()
	return s
}

// TrustPolicyAttachmentListOutput struct
type TrustPolicyAttachmentListOutput struct {
	Attachments []*TrustPolicyAttachment `json:"attachments,omitempty" yaml:"attachments,omitempty"`
}

// GetAttachments returns the value for the field attachments
func (e *TrustPolicyAttachmentListOutput) GetAttachments() []*TrustPolicyAttachment {
	return e.Attachments
}

// SetAttachments sets the value for the field attachments
func (e *TrustPolicyAttachmentListOutput) SetAttachments(attachments []*TrustPolicyAttachment) {
	e.Attachments = attachments
}

// StructPath returns StructPath
func (e *TrustPolicyAttachmentListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathTrustPolicyAttachmentListOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *TrustPolicyAttachmentListOutput) InitializeDefaults() {
}

// trustPolicyAttachmentListOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type trustPolicyAttachmentListOutputAlias TrustPolicyAttachmentListOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *TrustPolicyAttachmentListOutput) UnmarshalJSON(data []byte) error {
	var alias trustPolicyAttachmentListOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*TrustPolicyAttachmentListOutput)(&alias)).InitializeDefaults()
	*e = TrustPolicyAttachmentListOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e TrustPolicyAttachmentListOutput) MarshalJSON() ([]byte, error) {
	alias := trustPolicyAttachmentListOutputAlias(e)
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
	// when the token expires.
	// This value is always set. An absent value indicates a server fault, not a state.
	ExpiresAt *time.Time `json:"expiresAt,omitempty" yaml:"expiresAt,omitempty"`
	Token     string     `json:"token,omitempty" yaml:"token,omitempty"`
}

// GetExpiresAt returns the value for the field expiresAt
func (e *ServiceBearerToken) GetExpiresAt() *time.Time {
	return e.ExpiresAt
}

// SetExpiresAt sets the value for the field expiresAt
func (e *ServiceBearerToken) SetExpiresAt(expiresAt *time.Time) {
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

// NewInvalidServiceBearerTokenDurationError creates a new InvalidServiceBearerTokenDurationError
func NewInvalidServiceBearerTokenDurationError() *InvalidServiceBearerTokenDurationError {
	s := &InvalidServiceBearerTokenDurationError{}
	s.InitializeDefaults()
	return s
}

// InvalidServiceBearerTokenDurationError struct
type InvalidServiceBearerTokenDurationError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InvalidServiceBearerTokenDurationError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InvalidServiceBearerTokenDurationError]
func (e *InvalidServiceBearerTokenDurationError) Is(err error) bool {
	_, ok := err.(*InvalidServiceBearerTokenDurationError)
	return ok
}

// IsInvalidServiceBearerTokenDurationError indicates whether the given error chain contains an error of type [InvalidServiceBearerTokenDurationError]
func IsInvalidServiceBearerTokenDurationError(err error) bool {
	return errors.Is(err, &InvalidServiceBearerTokenDurationError{})
}

// GetMessage returns the value for the field message
func (e *InvalidServiceBearerTokenDurationError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InvalidServiceBearerTokenDurationError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InvalidServiceBearerTokenDurationError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvalidServiceBearerTokenDurationError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvalidServiceBearerTokenDurationError) InitializeDefaults() {
}

// invalidServiceBearerTokenDurationErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type invalidServiceBearerTokenDurationErrorAlias InvalidServiceBearerTokenDurationError

// UnmarshalJSON implements json.Unmarshaler
func (e *InvalidServiceBearerTokenDurationError) UnmarshalJSON(data []byte) error {
	var alias invalidServiceBearerTokenDurationErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvalidServiceBearerTokenDurationError)(&alias)).InitializeDefaults()
	*e = InvalidServiceBearerTokenDurationError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvalidServiceBearerTokenDurationError) MarshalJSON() ([]byte, error) {
	alias := invalidServiceBearerTokenDurationErrorAlias(e)
	return json.Marshal(alias)
}

// NewInvalidServiceNameError creates a new InvalidServiceNameError
func NewInvalidServiceNameError() *InvalidServiceNameError {
	s := &InvalidServiceNameError{}
	s.InitializeDefaults()
	return s
}

// InvalidServiceNameError struct
type InvalidServiceNameError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *InvalidServiceNameError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [InvalidServiceNameError]
func (e *InvalidServiceNameError) Is(err error) bool {
	_, ok := err.(*InvalidServiceNameError)
	return ok
}

// IsInvalidServiceNameError indicates whether the given error chain contains an error of type [InvalidServiceNameError]
func IsInvalidServiceNameError(err error) bool {
	return errors.Is(err, &InvalidServiceNameError{})
}

// GetMessage returns the value for the field message
func (e *InvalidServiceNameError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *InvalidServiceNameError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *InvalidServiceNameError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvalidServiceNameError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvalidServiceNameError) InitializeDefaults() {
}

// invalidServiceNameErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type invalidServiceNameErrorAlias InvalidServiceNameError

// UnmarshalJSON implements json.Unmarshaler
func (e *InvalidServiceNameError) UnmarshalJSON(data []byte) error {
	var alias invalidServiceNameErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvalidServiceNameError)(&alias)).InitializeDefaults()
	*e = InvalidServiceNameError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvalidServiceNameError) MarshalJSON() ([]byte, error) {
	alias := invalidServiceNameErrorAlias(e)
	return json.Marshal(alias)
}

// NewInvitationCreateInput creates a new InvitationCreateInput
func NewInvitationCreateInput() *InvitationCreateInput {
	s := &InvitationCreateInput{}
	s.InitializeDefaults()
	return s
}

// InvitationCreateInput struct
type InvitationCreateInput struct {
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
	// groups the invitee joins on accept. Absent or empty is legal and means the
	// invitee joins with policies alone, which is the pre-2026-08 behaviour
	GroupNames []string `json:"groupNames,omitempty" yaml:"groupNames,omitempty"`
	// identity policies (by name) to attach to the member when they accept
	PolicyNames []string `json:"policyNames,omitempty" yaml:"policyNames,omitempty"`
	// optional SSO provider hint (e.g. google/github)
	ProviderName *string `json:"providerName,omitempty" yaml:"providerName,omitempty"`
}

// GetEmail returns the value for the field email
func (e *InvitationCreateInput) GetEmail() string {
	return e.Email
}

// SetEmail sets the value for the field email
func (e *InvitationCreateInput) SetEmail(email string) {
	e.Email = email
}

// GetGroupNames returns the value for the field groupNames
func (e *InvitationCreateInput) GetGroupNames() []string {
	return e.GroupNames
}

// SetGroupNames sets the value for the field groupNames
func (e *InvitationCreateInput) SetGroupNames(groupNames []string) {
	e.GroupNames = groupNames
}

// GetPolicyNames returns the value for the field policyNames
func (e *InvitationCreateInput) GetPolicyNames() []string {
	return e.PolicyNames
}

// SetPolicyNames sets the value for the field policyNames
func (e *InvitationCreateInput) SetPolicyNames(policyNames []string) {
	e.PolicyNames = policyNames
}

// GetProviderName returns the value for the field providerName
func (e *InvitationCreateInput) GetProviderName() *string {
	return e.ProviderName
}

// SetProviderName sets the value for the field providerName
func (e *InvitationCreateInput) SetProviderName(providerName *string) {
	e.ProviderName = providerName
}

// StructPath returns StructPath
func (e *InvitationCreateInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvitationCreateInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvitationCreateInput) InitializeDefaults() {
}

// invitationCreateInputAlias is defined to help pre and post JSON marshaling without recursive loops
type invitationCreateInputAlias InvitationCreateInput

// UnmarshalJSON implements json.Unmarshaler
func (e *InvitationCreateInput) UnmarshalJSON(data []byte) error {
	var alias invitationCreateInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvitationCreateInput)(&alias)).InitializeDefaults()
	*e = InvitationCreateInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvitationCreateInput) MarshalJSON() ([]byte, error) {
	alias := invitationCreateInputAlias(e)
	return json.Marshal(alias)
}

// NewInvitationCreateOutput creates a new InvitationCreateOutput
func NewInvitationCreateOutput() *InvitationCreateOutput {
	s := &InvitationCreateOutput{}
	s.InitializeDefaults()
	return s
}

// InvitationCreateOutput struct
type InvitationCreateOutput struct {
	Invitation *Invitation `json:"invitation,omitempty" yaml:"invitation,omitempty"`
}

// GetInvitation returns the value for the field invitation
func (e *InvitationCreateOutput) GetInvitation() *Invitation {
	return e.Invitation
}

// SetInvitation sets the value for the field invitation
func (e *InvitationCreateOutput) SetInvitation(invitation *Invitation) {
	e.Invitation = invitation
}

// StructPath returns StructPath
func (e *InvitationCreateOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvitationCreateOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvitationCreateOutput) InitializeDefaults() {
}

// invitationCreateOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type invitationCreateOutputAlias InvitationCreateOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *InvitationCreateOutput) UnmarshalJSON(data []byte) error {
	var alias invitationCreateOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvitationCreateOutput)(&alias)).InitializeDefaults()
	*e = InvitationCreateOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvitationCreateOutput) MarshalJSON() ([]byte, error) {
	alias := invitationCreateOutputAlias(e)
	return json.Marshal(alias)
}

// NewInvitationListInput creates a new InvitationListInput
func NewInvitationListInput() *InvitationListInput {
	s := &InvitationListInput{}
	s.InitializeDefaults()
	return s
}

// InvitationListInput struct
type InvitationListInput struct {
}

// StructPath returns StructPath
func (e *InvitationListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvitationListInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvitationListInput) InitializeDefaults() {
}

// invitationListInputAlias is defined to help pre and post JSON marshaling without recursive loops
type invitationListInputAlias InvitationListInput

// UnmarshalJSON implements json.Unmarshaler
func (e *InvitationListInput) UnmarshalJSON(data []byte) error {
	var alias invitationListInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvitationListInput)(&alias)).InitializeDefaults()
	*e = InvitationListInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvitationListInput) MarshalJSON() ([]byte, error) {
	alias := invitationListInputAlias(e)
	return json.Marshal(alias)
}

// NewInvitationListOutput creates a new InvitationListOutput
func NewInvitationListOutput() *InvitationListOutput {
	s := &InvitationListOutput{}
	s.InitializeDefaults()
	return s
}

// InvitationListOutput struct
type InvitationListOutput struct {
	Invitations []*Invitation `json:"invitations,omitempty" yaml:"invitations,omitempty"`
}

// GetInvitations returns the value for the field invitations
func (e *InvitationListOutput) GetInvitations() []*Invitation {
	return e.Invitations
}

// SetInvitations sets the value for the field invitations
func (e *InvitationListOutput) SetInvitations(invitations []*Invitation) {
	e.Invitations = invitations
}

// StructPath returns StructPath
func (e *InvitationListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvitationListOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvitationListOutput) InitializeDefaults() {
}

// invitationListOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type invitationListOutputAlias InvitationListOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *InvitationListOutput) UnmarshalJSON(data []byte) error {
	var alias invitationListOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvitationListOutput)(&alias)).InitializeDefaults()
	*e = InvitationListOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvitationListOutput) MarshalJSON() ([]byte, error) {
	alias := invitationListOutputAlias(e)
	return json.Marshal(alias)
}

// NewInvitationRevokeInput creates a new InvitationRevokeInput
func NewInvitationRevokeInput() *InvitationRevokeInput {
	s := &InvitationRevokeInput{}
	s.InitializeDefaults()
	return s
}

// InvitationRevokeInput struct
type InvitationRevokeInput struct {
	InvitationID string `json:"invitationID,omitempty" yaml:"invitationID,omitempty"`
}

// GetInvitationID returns the value for the field invitationID
func (e *InvitationRevokeInput) GetInvitationID() string {
	return e.InvitationID
}

// SetInvitationID sets the value for the field invitationID
func (e *InvitationRevokeInput) SetInvitationID(invitationID string) {
	e.InvitationID = invitationID
}

// StructPath returns StructPath
func (e *InvitationRevokeInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvitationRevokeInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvitationRevokeInput) InitializeDefaults() {
}

// invitationRevokeInputAlias is defined to help pre and post JSON marshaling without recursive loops
type invitationRevokeInputAlias InvitationRevokeInput

// UnmarshalJSON implements json.Unmarshaler
func (e *InvitationRevokeInput) UnmarshalJSON(data []byte) error {
	var alias invitationRevokeInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvitationRevokeInput)(&alias)).InitializeDefaults()
	*e = InvitationRevokeInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvitationRevokeInput) MarshalJSON() ([]byte, error) {
	alias := invitationRevokeInputAlias(e)
	return json.Marshal(alias)
}

// NewInvitationRevokeOutput creates a new InvitationRevokeOutput
func NewInvitationRevokeOutput() *InvitationRevokeOutput {
	s := &InvitationRevokeOutput{}
	s.InitializeDefaults()
	return s
}

// InvitationRevokeOutput struct
type InvitationRevokeOutput struct {
}

// StructPath returns StructPath
func (e *InvitationRevokeOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvitationRevokeOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvitationRevokeOutput) InitializeDefaults() {
}

// invitationRevokeOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type invitationRevokeOutputAlias InvitationRevokeOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *InvitationRevokeOutput) UnmarshalJSON(data []byte) error {
	var alias invitationRevokeOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvitationRevokeOutput)(&alias)).InitializeDefaults()
	*e = InvitationRevokeOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvitationRevokeOutput) MarshalJSON() ([]byte, error) {
	alias := invitationRevokeOutputAlias(e)
	return json.Marshal(alias)
}

// NewInvitationAcceptInput creates a new InvitationAcceptInput
func NewInvitationAcceptInput() *InvitationAcceptInput {
	s := &InvitationAcceptInput{}
	s.InitializeDefaults()
	return s
}

// InvitationAcceptInput struct
type InvitationAcceptInput struct {
	Token string `json:"token,omitempty" yaml:"token,omitempty"`
}

// GetToken returns the value for the field token
func (e *InvitationAcceptInput) GetToken() string {
	return e.Token
}

// SetToken sets the value for the field token
func (e *InvitationAcceptInput) SetToken(token string) {
	e.Token = token
}

// StructPath returns StructPath
func (e *InvitationAcceptInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvitationAcceptInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvitationAcceptInput) InitializeDefaults() {
}

// invitationAcceptInputAlias is defined to help pre and post JSON marshaling without recursive loops
type invitationAcceptInputAlias InvitationAcceptInput

// UnmarshalJSON implements json.Unmarshaler
func (e *InvitationAcceptInput) UnmarshalJSON(data []byte) error {
	var alias invitationAcceptInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvitationAcceptInput)(&alias)).InitializeDefaults()
	*e = InvitationAcceptInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvitationAcceptInput) MarshalJSON() ([]byte, error) {
	alias := invitationAcceptInputAlias(e)
	return json.Marshal(alias)
}

// NewInvitationAcceptOutput creates a new InvitationAcceptOutput
func NewInvitationAcceptOutput() *InvitationAcceptOutput {
	s := &InvitationAcceptOutput{}
	s.InitializeDefaults()
	return s
}

// InvitationAcceptOutput struct
type InvitationAcceptOutput struct {
	AccountName string `json:"accountName,omitempty" yaml:"accountName,omitempty"`
}

// GetAccountName returns the value for the field accountName
func (e *InvitationAcceptOutput) GetAccountName() string {
	return e.AccountName
}

// SetAccountName sets the value for the field accountName
func (e *InvitationAcceptOutput) SetAccountName(accountName string) {
	e.AccountName = accountName
}

// StructPath returns StructPath
func (e *InvitationAcceptOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvitationAcceptOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvitationAcceptOutput) InitializeDefaults() {
}

// invitationAcceptOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type invitationAcceptOutputAlias InvitationAcceptOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *InvitationAcceptOutput) UnmarshalJSON(data []byte) error {
	var alias invitationAcceptOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvitationAcceptOutput)(&alias)).InitializeDefaults()
	*e = InvitationAcceptOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvitationAcceptOutput) MarshalJSON() ([]byte, error) {
	alias := invitationAcceptOutputAlias(e)
	return json.Marshal(alias)
}

// NewInvitationInspectInput creates a new InvitationInspectInput
func NewInvitationInspectInput() *InvitationInspectInput {
	s := &InvitationInspectInput{}
	s.InitializeDefaults()
	return s
}

// InvitationInspectInput struct
type InvitationInspectInput struct {
	Token string `json:"token,omitempty" yaml:"token,omitempty"`
}

// GetToken returns the value for the field token
func (e *InvitationInspectInput) GetToken() string {
	return e.Token
}

// SetToken sets the value for the field token
func (e *InvitationInspectInput) SetToken(token string) {
	e.Token = token
}

// StructPath returns StructPath
func (e *InvitationInspectInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvitationInspectInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvitationInspectInput) InitializeDefaults() {
}

// invitationInspectInputAlias is defined to help pre and post JSON marshaling without recursive loops
type invitationInspectInputAlias InvitationInspectInput

// UnmarshalJSON implements json.Unmarshaler
func (e *InvitationInspectInput) UnmarshalJSON(data []byte) error {
	var alias invitationInspectInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvitationInspectInput)(&alias)).InitializeDefaults()
	*e = InvitationInspectInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvitationInspectInput) MarshalJSON() ([]byte, error) {
	alias := invitationInspectInputAlias(e)
	return json.Marshal(alias)
}

// NewInvitationInspectOutput creates a new InvitationInspectOutput
func NewInvitationInspectOutput() *InvitationInspectOutput {
	s := &InvitationInspectOutput{}
	s.InitializeDefaults()
	return s
}

// InvitationInspectOutput struct
type InvitationInspectOutput struct {
	Invitation *InvitationPreview `json:"invitation,omitempty" yaml:"invitation,omitempty"`
}

// GetInvitation returns the value for the field invitation
func (e *InvitationInspectOutput) GetInvitation() *InvitationPreview {
	return e.Invitation
}

// SetInvitation sets the value for the field invitation
func (e *InvitationInspectOutput) SetInvitation(invitation *InvitationPreview) {
	e.Invitation = invitation
}

// StructPath returns StructPath
func (e *InvitationInspectOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvitationInspectOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvitationInspectOutput) InitializeDefaults() {
}

// invitationInspectOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type invitationInspectOutputAlias InvitationInspectOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *InvitationInspectOutput) UnmarshalJSON(data []byte) error {
	var alias invitationInspectOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvitationInspectOutput)(&alias)).InitializeDefaults()
	*e = InvitationInspectOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvitationInspectOutput) MarshalJSON() ([]byte, error) {
	alias := invitationInspectOutputAlias(e)
	return json.Marshal(alias)
}

// NewInvitationDeclineInput creates a new InvitationDeclineInput
func NewInvitationDeclineInput() *InvitationDeclineInput {
	s := &InvitationDeclineInput{}
	s.InitializeDefaults()
	return s
}

// InvitationDeclineInput struct
type InvitationDeclineInput struct {
	Token string `json:"token,omitempty" yaml:"token,omitempty"`
}

// GetToken returns the value for the field token
func (e *InvitationDeclineInput) GetToken() string {
	return e.Token
}

// SetToken sets the value for the field token
func (e *InvitationDeclineInput) SetToken(token string) {
	e.Token = token
}

// StructPath returns StructPath
func (e *InvitationDeclineInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvitationDeclineInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvitationDeclineInput) InitializeDefaults() {
}

// invitationDeclineInputAlias is defined to help pre and post JSON marshaling without recursive loops
type invitationDeclineInputAlias InvitationDeclineInput

// UnmarshalJSON implements json.Unmarshaler
func (e *InvitationDeclineInput) UnmarshalJSON(data []byte) error {
	var alias invitationDeclineInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvitationDeclineInput)(&alias)).InitializeDefaults()
	*e = InvitationDeclineInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvitationDeclineInput) MarshalJSON() ([]byte, error) {
	alias := invitationDeclineInputAlias(e)
	return json.Marshal(alias)
}

// NewInvitationDeclineOutput creates a new InvitationDeclineOutput
func NewInvitationDeclineOutput() *InvitationDeclineOutput {
	s := &InvitationDeclineOutput{}
	s.InitializeDefaults()
	return s
}

// InvitationDeclineOutput struct
type InvitationDeclineOutput struct {
}

// StructPath returns StructPath
func (e *InvitationDeclineOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathInvitationDeclineOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *InvitationDeclineOutput) InitializeDefaults() {
}

// invitationDeclineOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type invitationDeclineOutputAlias InvitationDeclineOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *InvitationDeclineOutput) UnmarshalJSON(data []byte) error {
	var alias invitationDeclineOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*InvitationDeclineOutput)(&alias)).InitializeDefaults()
	*e = InvitationDeclineOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e InvitationDeclineOutput) MarshalJSON() ([]byte, error) {
	alias := invitationDeclineOutputAlias(e)
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

// NewSessionKeepAliveInput creates a new SessionKeepAliveInput
func NewSessionKeepAliveInput() *SessionKeepAliveInput {
	s := &SessionKeepAliveInput{}
	s.InitializeDefaults()
	return s
}

// SessionKeepAliveInput struct
type SessionKeepAliveInput struct {
}

// StructPath returns StructPath
func (e *SessionKeepAliveInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathSessionKeepAliveInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *SessionKeepAliveInput) InitializeDefaults() {
}

// sessionKeepAliveInputAlias is defined to help pre and post JSON marshaling without recursive loops
type sessionKeepAliveInputAlias SessionKeepAliveInput

// UnmarshalJSON implements json.Unmarshaler
func (e *SessionKeepAliveInput) UnmarshalJSON(data []byte) error {
	var alias sessionKeepAliveInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*SessionKeepAliveInput)(&alias)).InitializeDefaults()
	*e = SessionKeepAliveInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e SessionKeepAliveInput) MarshalJSON() ([]byte, error) {
	alias := sessionKeepAliveInputAlias(e)
	return json.Marshal(alias)
}

// NewSessionKeepAliveOutput creates a new SessionKeepAliveOutput
func NewSessionKeepAliveOutput() *SessionKeepAliveOutput {
	s := &SessionKeepAliveOutput{}
	s.InitializeDefaults()
	return s
}

// SessionKeepAliveOutput struct
type SessionKeepAliveOutput struct {
	// the new sliding deadline after this heartbeat (bounded by the cap).
	// Absent means this credential does not expire, which a permanent key does
	// not. It is a real answer, not a fault: do not read it as a deadline in
	// the past.
	ExpiresAt *time.Time `json:"expiresAt,omitempty" yaml:"expiresAt,omitempty"`
}

// GetExpiresAt returns the value for the field expiresAt
func (e *SessionKeepAliveOutput) GetExpiresAt() *time.Time {
	return e.ExpiresAt
}

// SetExpiresAt sets the value for the field expiresAt
func (e *SessionKeepAliveOutput) SetExpiresAt(expiresAt *time.Time) {
	e.ExpiresAt = expiresAt
}

// StructPath returns StructPath
func (e *SessionKeepAliveOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathSessionKeepAliveOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *SessionKeepAliveOutput) InitializeDefaults() {
}

// sessionKeepAliveOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type sessionKeepAliveOutputAlias SessionKeepAliveOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *SessionKeepAliveOutput) UnmarshalJSON(data []byte) error {
	var alias sessionKeepAliveOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*SessionKeepAliveOutput)(&alias)).InitializeDefaults()
	*e = SessionKeepAliveOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e SessionKeepAliveOutput) MarshalJSON() ([]byte, error) {
	alias := sessionKeepAliveOutputAlias(e)
	return json.Marshal(alias)
}

// NewSessionRevokeInput creates a new SessionRevokeInput
func NewSessionRevokeInput() *SessionRevokeInput {
	s := &SessionRevokeInput{}
	s.InitializeDefaults()
	return s
}

// SessionRevokeInput struct
type SessionRevokeInput struct {
}

// StructPath returns StructPath
func (e *SessionRevokeInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathSessionRevokeInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *SessionRevokeInput) InitializeDefaults() {
}

// sessionRevokeInputAlias is defined to help pre and post JSON marshaling without recursive loops
type sessionRevokeInputAlias SessionRevokeInput

// UnmarshalJSON implements json.Unmarshaler
func (e *SessionRevokeInput) UnmarshalJSON(data []byte) error {
	var alias sessionRevokeInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*SessionRevokeInput)(&alias)).InitializeDefaults()
	*e = SessionRevokeInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e SessionRevokeInput) MarshalJSON() ([]byte, error) {
	alias := sessionRevokeInputAlias(e)
	return json.Marshal(alias)
}

// NewSessionRevokeOutput creates a new SessionRevokeOutput
func NewSessionRevokeOutput() *SessionRevokeOutput {
	s := &SessionRevokeOutput{}
	s.InitializeDefaults()
	return s
}

// SessionRevokeOutput struct
type SessionRevokeOutput struct {
}

// StructPath returns StructPath
func (e *SessionRevokeOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathSessionRevokeOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *SessionRevokeOutput) InitializeDefaults() {
}

// sessionRevokeOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type sessionRevokeOutputAlias SessionRevokeOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *SessionRevokeOutput) UnmarshalJSON(data []byte) error {
	var alias sessionRevokeOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*SessionRevokeOutput)(&alias)).InitializeDefaults()
	*e = SessionRevokeOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e SessionRevokeOutput) MarshalJSON() ([]byte, error) {
	alias := sessionRevokeOutputAlias(e)
	return json.Marshal(alias)
}

// NewSessionIdentityInput creates a new SessionIdentityInput
func NewSessionIdentityInput() *SessionIdentityInput {
	s := &SessionIdentityInput{}
	s.InitializeDefaults()
	return s
}

// SessionIdentityInput struct
type SessionIdentityInput struct {
}

// StructPath returns StructPath
func (e *SessionIdentityInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathSessionIdentityInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *SessionIdentityInput) InitializeDefaults() {
}

// sessionIdentityInputAlias is defined to help pre and post JSON marshaling without recursive loops
type sessionIdentityInputAlias SessionIdentityInput

// UnmarshalJSON implements json.Unmarshaler
func (e *SessionIdentityInput) UnmarshalJSON(data []byte) error {
	var alias sessionIdentityInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*SessionIdentityInput)(&alias)).InitializeDefaults()
	*e = SessionIdentityInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e SessionIdentityInput) MarshalJSON() ([]byte, error) {
	alias := sessionIdentityInputAlias(e)
	return json.Marshal(alias)
}

// NewSessionIdentityOutput creates a new SessionIdentityOutput
func NewSessionIdentityOutput() *SessionIdentityOutput {
	s := &SessionIdentityOutput{}
	s.InitializeDefaults()
	return s
}

// SessionIdentityOutput struct
type SessionIdentityOutput struct {
	// the public id of the calling access key
	AccessKeyID string `json:"accessKeyID,omitempty" yaml:"accessKeyID,omitempty"`
	// the account the calling credential is scoped to
	AccountName string `json:"accountName,omitempty" yaml:"accountName,omitempty"`
	// present only when acting via a service-linked role
	ManagedBy *ManagedByService `json:"managedBy,omitempty" yaml:"managedBy,omitempty"`
	// DRN of the calling principal: iam:User(alice), iam:Role(deployer), or
	// iam:Service(uplink.deployport.io) when the caller is a service itself
	PrincipalDrn string `json:"principalDrn,omitempty" yaml:"principalDrn,omitempty"`
}

// GetAccessKeyID returns the value for the field accessKeyID
func (e *SessionIdentityOutput) GetAccessKeyID() string {
	return e.AccessKeyID
}

// SetAccessKeyID sets the value for the field accessKeyID
func (e *SessionIdentityOutput) SetAccessKeyID(accessKeyID string) {
	e.AccessKeyID = accessKeyID
}

// GetAccountName returns the value for the field accountName
func (e *SessionIdentityOutput) GetAccountName() string {
	return e.AccountName
}

// SetAccountName sets the value for the field accountName
func (e *SessionIdentityOutput) SetAccountName(accountName string) {
	e.AccountName = accountName
}

// GetManagedBy returns the value for the field managedBy
func (e *SessionIdentityOutput) GetManagedBy() *ManagedByService {
	return e.ManagedBy
}

// SetManagedBy sets the value for the field managedBy
func (e *SessionIdentityOutput) SetManagedBy(managedBy *ManagedByService) {
	e.ManagedBy = managedBy
}

// GetPrincipalDrn returns the value for the field principalDrn
func (e *SessionIdentityOutput) GetPrincipalDrn() string {
	return e.PrincipalDrn
}

// SetPrincipalDrn sets the value for the field principalDrn
func (e *SessionIdentityOutput) SetPrincipalDrn(principalDrn string) {
	e.PrincipalDrn = principalDrn
}

// StructPath returns StructPath
func (e *SessionIdentityOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathSessionIdentityOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *SessionIdentityOutput) InitializeDefaults() {
}

// sessionIdentityOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type sessionIdentityOutputAlias SessionIdentityOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *SessionIdentityOutput) UnmarshalJSON(data []byte) error {
	var alias sessionIdentityOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*SessionIdentityOutput)(&alias)).InitializeDefaults()
	*e = SessionIdentityOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e SessionIdentityOutput) MarshalJSON() ([]byte, error) {
	alias := sessionIdentityOutputAlias(e)
	return json.Marshal(alias)
}

// NewSessionBeginHandoffInput creates a new SessionBeginHandoffInput
func NewSessionBeginHandoffInput() *SessionBeginHandoffInput {
	s := &SessionBeginHandoffInput{}
	s.InitializeDefaults()
	return s
}

// SessionBeginHandoffInput struct
type SessionBeginHandoffInput struct {
}

// StructPath returns StructPath
func (e *SessionBeginHandoffInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathSessionBeginHandoffInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *SessionBeginHandoffInput) InitializeDefaults() {
}

// sessionBeginHandoffInputAlias is defined to help pre and post JSON marshaling without recursive loops
type sessionBeginHandoffInputAlias SessionBeginHandoffInput

// UnmarshalJSON implements json.Unmarshaler
func (e *SessionBeginHandoffInput) UnmarshalJSON(data []byte) error {
	var alias sessionBeginHandoffInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*SessionBeginHandoffInput)(&alias)).InitializeDefaults()
	*e = SessionBeginHandoffInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e SessionBeginHandoffInput) MarshalJSON() ([]byte, error) {
	alias := sessionBeginHandoffInputAlias(e)
	return json.Marshal(alias)
}

// NewSessionBeginHandoffOutput creates a new SessionBeginHandoffOutput
func NewSessionBeginHandoffOutput() *SessionBeginHandoffOutput {
	s := &SessionBeginHandoffOutput{}
	s.InitializeDefaults()
	return s
}

// SessionBeginHandoffOutput struct
type SessionBeginHandoffOutput struct {
	// single-use opaque code; redeem within a short window via CompleteHandoff
	Code string `json:"code,omitempty" yaml:"code,omitempty"`
}

// GetCode returns the value for the field code
func (e *SessionBeginHandoffOutput) GetCode() string {
	return e.Code
}

// SetCode sets the value for the field code
func (e *SessionBeginHandoffOutput) SetCode(code string) {
	e.Code = code
}

// StructPath returns StructPath
func (e *SessionBeginHandoffOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathSessionBeginHandoffOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *SessionBeginHandoffOutput) InitializeDefaults() {
}

// sessionBeginHandoffOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type sessionBeginHandoffOutputAlias SessionBeginHandoffOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *SessionBeginHandoffOutput) UnmarshalJSON(data []byte) error {
	var alias sessionBeginHandoffOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*SessionBeginHandoffOutput)(&alias)).InitializeDefaults()
	*e = SessionBeginHandoffOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e SessionBeginHandoffOutput) MarshalJSON() ([]byte, error) {
	alias := sessionBeginHandoffOutputAlias(e)
	return json.Marshal(alias)
}

// NewSessionCompleteHandoffInput creates a new SessionCompleteHandoffInput
func NewSessionCompleteHandoffInput() *SessionCompleteHandoffInput {
	s := &SessionCompleteHandoffInput{}
	s.InitializeDefaults()
	return s
}

// SessionCompleteHandoffInput struct
type SessionCompleteHandoffInput struct {
	Code string `json:"code,omitempty" yaml:"code,omitempty"`
}

// GetCode returns the value for the field code
func (e *SessionCompleteHandoffInput) GetCode() string {
	return e.Code
}

// SetCode sets the value for the field code
func (e *SessionCompleteHandoffInput) SetCode(code string) {
	e.Code = code
}

// StructPath returns StructPath
func (e *SessionCompleteHandoffInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathSessionCompleteHandoffInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *SessionCompleteHandoffInput) InitializeDefaults() {
}

// sessionCompleteHandoffInputAlias is defined to help pre and post JSON marshaling without recursive loops
type sessionCompleteHandoffInputAlias SessionCompleteHandoffInput

// UnmarshalJSON implements json.Unmarshaler
func (e *SessionCompleteHandoffInput) UnmarshalJSON(data []byte) error {
	var alias sessionCompleteHandoffInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*SessionCompleteHandoffInput)(&alias)).InitializeDefaults()
	*e = SessionCompleteHandoffInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e SessionCompleteHandoffInput) MarshalJSON() ([]byte, error) {
	alias := sessionCompleteHandoffInputAlias(e)
	return json.Marshal(alias)
}

// NewSessionCompleteHandoffOutput creates a new SessionCompleteHandoffOutput
func NewSessionCompleteHandoffOutput() *SessionCompleteHandoffOutput {
	s := &SessionCompleteHandoffOutput{}
	s.InitializeDefaults()
	return s
}

// SessionCompleteHandoffOutput struct
type SessionCompleteHandoffOutput struct {
	// the account the new session is scoped to
	AccountName string       `json:"accountName,omitempty" yaml:"accountName,omitempty"`
	Credentials *Credentials `json:"credentials,omitempty" yaml:"credentials,omitempty"`
	// DRN of the principal the new session acts as, e.g. iam:User(alice)
	PrincipalDrn string `json:"principalDrn,omitempty" yaml:"principalDrn,omitempty"`
}

// GetAccountName returns the value for the field accountName
func (e *SessionCompleteHandoffOutput) GetAccountName() string {
	return e.AccountName
}

// SetAccountName sets the value for the field accountName
func (e *SessionCompleteHandoffOutput) SetAccountName(accountName string) {
	e.AccountName = accountName
}

// GetCredentials returns the value for the field credentials
func (e *SessionCompleteHandoffOutput) GetCredentials() *Credentials {
	return e.Credentials
}

// SetCredentials sets the value for the field credentials
func (e *SessionCompleteHandoffOutput) SetCredentials(credentials *Credentials) {
	e.Credentials = credentials
}

// GetPrincipalDrn returns the value for the field principalDrn
func (e *SessionCompleteHandoffOutput) GetPrincipalDrn() string {
	return e.PrincipalDrn
}

// SetPrincipalDrn sets the value for the field principalDrn
func (e *SessionCompleteHandoffOutput) SetPrincipalDrn(principalDrn string) {
	e.PrincipalDrn = principalDrn
}

// StructPath returns StructPath
func (e *SessionCompleteHandoffOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathSessionCompleteHandoffOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *SessionCompleteHandoffOutput) InitializeDefaults() {
}

// sessionCompleteHandoffOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type sessionCompleteHandoffOutputAlias SessionCompleteHandoffOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *SessionCompleteHandoffOutput) UnmarshalJSON(data []byte) error {
	var alias sessionCompleteHandoffOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*SessionCompleteHandoffOutput)(&alias)).InitializeDefaults()
	*e = SessionCompleteHandoffOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e SessionCompleteHandoffOutput) MarshalJSON() ([]byte, error) {
	alias := sessionCompleteHandoffOutputAlias(e)
	return json.Marshal(alias)
}

// NewServiceCatalogListInput creates a new ServiceCatalogListInput
func NewServiceCatalogListInput() *ServiceCatalogListInput {
	s := &ServiceCatalogListInput{}
	s.InitializeDefaults()
	return s
}

// ServiceCatalogListInput struct
type ServiceCatalogListInput struct {
}

// StructPath returns StructPath
func (e *ServiceCatalogListInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathServiceCatalogListInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *ServiceCatalogListInput) InitializeDefaults() {
}

// serviceCatalogListInputAlias is defined to help pre and post JSON marshaling without recursive loops
type serviceCatalogListInputAlias ServiceCatalogListInput

// UnmarshalJSON implements json.Unmarshaler
func (e *ServiceCatalogListInput) UnmarshalJSON(data []byte) error {
	var alias serviceCatalogListInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*ServiceCatalogListInput)(&alias)).InitializeDefaults()
	*e = ServiceCatalogListInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e ServiceCatalogListInput) MarshalJSON() ([]byte, error) {
	alias := serviceCatalogListInputAlias(e)
	return json.Marshal(alias)
}

// NewServiceCatalogListOutput creates a new ServiceCatalogListOutput
func NewServiceCatalogListOutput() *ServiceCatalogListOutput {
	s := &ServiceCatalogListOutput{}
	s.InitializeDefaults()
	return s
}

// ServiceCatalogListOutput struct
type ServiceCatalogListOutput struct {
	Catalogs []*ServiceCatalogSummary `json:"catalogs,omitempty" yaml:"catalogs,omitempty"`
}

// GetCatalogs returns the value for the field catalogs
func (e *ServiceCatalogListOutput) GetCatalogs() []*ServiceCatalogSummary {
	return e.Catalogs
}

// SetCatalogs sets the value for the field catalogs
func (e *ServiceCatalogListOutput) SetCatalogs(catalogs []*ServiceCatalogSummary) {
	e.Catalogs = catalogs
}

// StructPath returns StructPath
func (e *ServiceCatalogListOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathServiceCatalogListOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *ServiceCatalogListOutput) InitializeDefaults() {
}

// serviceCatalogListOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type serviceCatalogListOutputAlias ServiceCatalogListOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *ServiceCatalogListOutput) UnmarshalJSON(data []byte) error {
	var alias serviceCatalogListOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*ServiceCatalogListOutput)(&alias)).InitializeDefaults()
	*e = ServiceCatalogListOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e ServiceCatalogListOutput) MarshalJSON() ([]byte, error) {
	alias := serviceCatalogListOutputAlias(e)
	return json.Marshal(alias)
}

// NewServiceCatalogGetInput creates a new ServiceCatalogGetInput
func NewServiceCatalogGetInput() *ServiceCatalogGetInput {
	s := &ServiceCatalogGetInput{}
	s.InitializeDefaults()
	return s
}

// ServiceCatalogGetInput struct
type ServiceCatalogGetInput struct {
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}

// GetNamespace returns the value for the field namespace
func (e *ServiceCatalogGetInput) GetNamespace() string {
	return e.Namespace
}

// SetNamespace sets the value for the field namespace
func (e *ServiceCatalogGetInput) SetNamespace(namespace string) {
	e.Namespace = namespace
}

// StructPath returns StructPath
func (e *ServiceCatalogGetInput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathServiceCatalogGetInput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *ServiceCatalogGetInput) InitializeDefaults() {
}

// serviceCatalogGetInputAlias is defined to help pre and post JSON marshaling without recursive loops
type serviceCatalogGetInputAlias ServiceCatalogGetInput

// UnmarshalJSON implements json.Unmarshaler
func (e *ServiceCatalogGetInput) UnmarshalJSON(data []byte) error {
	var alias serviceCatalogGetInputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*ServiceCatalogGetInput)(&alias)).InitializeDefaults()
	*e = ServiceCatalogGetInput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e ServiceCatalogGetInput) MarshalJSON() ([]byte, error) {
	alias := serviceCatalogGetInputAlias(e)
	return json.Marshal(alias)
}

// NewServiceCatalogGetOutput creates a new ServiceCatalogGetOutput
func NewServiceCatalogGetOutput() *ServiceCatalogGetOutput {
	s := &ServiceCatalogGetOutput{}
	s.InitializeDefaults()
	return s
}

// ServiceCatalogGetOutput struct
type ServiceCatalogGetOutput struct {
	Catalog *ServiceCatalogInfo `json:"catalog,omitempty" yaml:"catalog,omitempty"`
}

// GetCatalog returns the value for the field catalog
func (e *ServiceCatalogGetOutput) GetCatalog() *ServiceCatalogInfo {
	return e.Catalog
}

// SetCatalog sets the value for the field catalog
func (e *ServiceCatalogGetOutput) SetCatalog(catalog *ServiceCatalogInfo) {
	e.Catalog = catalog
}

// StructPath returns StructPath
func (e *ServiceCatalogGetOutput) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathServiceCatalogGetOutput.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *ServiceCatalogGetOutput) InitializeDefaults() {
}

// serviceCatalogGetOutputAlias is defined to help pre and post JSON marshaling without recursive loops
type serviceCatalogGetOutputAlias ServiceCatalogGetOutput

// UnmarshalJSON implements json.Unmarshaler
func (e *ServiceCatalogGetOutput) UnmarshalJSON(data []byte) error {
	var alias serviceCatalogGetOutputAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*ServiceCatalogGetOutput)(&alias)).InitializeDefaults()
	*e = ServiceCatalogGetOutput(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e ServiceCatalogGetOutput) MarshalJSON() ([]byte, error) {
	alias := serviceCatalogGetOutputAlias(e)
	return json.Marshal(alias)
}

// NewServiceCatalogGetServiceCatalogNotFoundError creates a new ServiceCatalogGetServiceCatalogNotFoundError
func NewServiceCatalogGetServiceCatalogNotFoundError() *ServiceCatalogGetServiceCatalogNotFoundError {
	s := &ServiceCatalogGetServiceCatalogNotFoundError{}
	s.InitializeDefaults()
	return s
}

// ServiceCatalogGetServiceCatalogNotFoundError struct
type ServiceCatalogGetServiceCatalogNotFoundError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Error implements the error interface
func (e *ServiceCatalogGetServiceCatalogNotFoundError) Error() string {
	return e.GetMessage()
}

// Is indicates whether the given error chain contains an error of type [ServiceCatalogGetServiceCatalogNotFoundError]
func (e *ServiceCatalogGetServiceCatalogNotFoundError) Is(err error) bool {
	_, ok := err.(*ServiceCatalogGetServiceCatalogNotFoundError)
	return ok
}

// IsServiceCatalogGetServiceCatalogNotFoundError indicates whether the given error chain contains an error of type [ServiceCatalogGetServiceCatalogNotFoundError]
func IsServiceCatalogGetServiceCatalogNotFoundError(err error) bool {
	return errors.Is(err, &ServiceCatalogGetServiceCatalogNotFoundError{})
}

// GetMessage returns the value for the field message
func (e *ServiceCatalogGetServiceCatalogNotFoundError) GetMessage() string {
	return e.Message
}

// SetMessage sets the value for the field message
func (e *ServiceCatalogGetServiceCatalogNotFoundError) SetMessage(message string) {
	e.Message = message
}

// StructPath returns StructPath
func (e *ServiceCatalogGetServiceCatalogNotFoundError) StructPath() clientruntime.StructPath {
	return *localSpecularMeta.structPathServiceCatalogGetServiceCatalogNotFoundError.Path()
}

// InitializeDefaults initializes the default values in the struct
func (e *ServiceCatalogGetServiceCatalogNotFoundError) InitializeDefaults() {
}

// serviceCatalogGetServiceCatalogNotFoundErrorAlias is defined to help pre and post JSON marshaling without recursive loops
type serviceCatalogGetServiceCatalogNotFoundErrorAlias ServiceCatalogGetServiceCatalogNotFoundError

// UnmarshalJSON implements json.Unmarshaler
func (e *ServiceCatalogGetServiceCatalogNotFoundError) UnmarshalJSON(data []byte) error {
	var alias serviceCatalogGetServiceCatalogNotFoundErrorAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	((*ServiceCatalogGetServiceCatalogNotFoundError)(&alias)).InitializeDefaults()
	*e = ServiceCatalogGetServiceCatalogNotFoundError(alias)
	return nil
}

// MarshalJSON implements json.Marshaler
func (e ServiceCatalogGetServiceCatalogNotFoundError) MarshalJSON() ([]byte, error) {
	alias := serviceCatalogGetServiceCatalogNotFoundErrorAlias(e)
	return json.Marshal(alias)
}

// TrustStatementEffect entity
// Effect of a trust-policy statement. Deny overrides Allow.
type TrustStatementEffect string

// TrustStatementEffectAllow = "allow"
const TrustStatementEffectAllow TrustStatementEffect = "allow"

// TrustStatementEffectDeny = "deny"
const TrustStatementEffectDeny TrustStatementEffect = "deny"

// String returns the string representation of the enum
func (en TrustStatementEffect) String() string {
	return string(en)
}

// Optional returns the optional value
func (en TrustStatementEffect) Optional() *TrustStatementEffect {
	c := en
	return &c
}

// TrustStatementEffectAllConstants returns all constants in enum TrustStatementEffect
func TrustStatementEffectAllConstants() []TrustStatementEffect {
	return []TrustStatementEffect{
		TrustStatementEffectAllow,
		TrustStatementEffectDeny,
	}
}

// PolicyStatementEffect entity
type PolicyStatementEffect string

// PolicyStatementEffectDeny = "deny"
const PolicyStatementEffectDeny PolicyStatementEffect = "deny"

// PolicyStatementEffectAllow = "allow"
const PolicyStatementEffectAllow PolicyStatementEffect = "allow"

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
	localSpecularMeta.structPathServiceCatalogAction, err = pk.NewType(
		"ServiceCatalogAction",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewServiceCatalogAction()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathServiceCatalogResourceType, err = pk.NewType(
		"ServiceCatalogResourceType",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewServiceCatalogResourceType()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathServiceCatalogData, err = pk.NewType(
		"ServiceCatalogData",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewServiceCatalogData()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathServiceCatalogInfo, err = pk.NewType(
		"ServiceCatalogInfo",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewServiceCatalogInfo()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathServiceCatalogSummary, err = pk.NewType(
		"ServiceCatalogSummary",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewServiceCatalogSummary()
		}),
	)
	if err != nil {
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
	localSpecularMeta.structPathManagedByService, err = pk.NewType(
		"ManagedByService",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewManagedByService()
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
	localSpecularMeta.structPathGroupInformation, err = pk.NewType(
		"GroupInformation",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewGroupInformation()
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
	localSpecularMeta.structPathSSOProviderUnavailableError, err = pk.NewType(
		"SSOProviderUnavailableError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewSSOProviderUnavailableError()
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
	localSpecularMeta.structPathRegionEndpoint, err = pk.NewType(
		"RegionEndpoint",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRegionEndpoint()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRegionInfo, err = pk.NewType(
		"RegionInfo",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRegionInfo()
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
	localSpecularMeta.structPathPolicyNotFoundError, err = pk.NewType(
		"PolicyNotFoundError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewPolicyNotFoundError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvitation, err = pk.NewType(
		"Invitation",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvitation()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvalidInvitationError, err = pk.NewType(
		"InvalidInvitationError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvalidInvitationError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvitationNotFoundError, err = pk.NewType(
		"InvitationNotFoundError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvitationNotFoundError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvitationPreview, err = pk.NewType(
		"InvitationPreview",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvitationPreview()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathOIDCProvider, err = pk.NewType(
		"OIDCProvider",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewOIDCProvider()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvalidOIDCProviderError, err = pk.NewType(
		"InvalidOIDCProviderError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvalidOIDCProviderError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvalidOIDCIssuerError, err = pk.NewType(
		"InvalidOIDCIssuerError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvalidOIDCIssuerError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathOIDCProviderNotFoundError, err = pk.NewType(
		"OIDCProviderNotFoundError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewOIDCProviderNotFoundError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathOIDCProviderInUseError, err = pk.NewType(
		"OIDCProviderInUseError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewOIDCProviderInUseError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathTrustPolicyStatement, err = pk.NewType(
		"TrustPolicyStatement",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewTrustPolicyStatement()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathTrustPolicy, err = pk.NewType(
		"TrustPolicy",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewTrustPolicy()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathTrustPolicyAttachment, err = pk.NewType(
		"TrustPolicyAttachment",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewTrustPolicyAttachment()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvalidTrustPolicyError, err = pk.NewType(
		"InvalidTrustPolicyError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvalidTrustPolicyError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathTrustPolicyNotFoundError, err = pk.NewType(
		"TrustPolicyNotFoundError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewTrustPolicyNotFoundError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvalidWebIdentityTokenError, err = pk.NewType(
		"InvalidWebIdentityTokenError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvalidWebIdentityTokenError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvalidPrincipalDRNError, err = pk.NewType(
		"InvalidPrincipalDRNError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvalidPrincipalDRNError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvalidHandoffCodeError, err = pk.NewType(
		"InvalidHandoffCodeError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvalidHandoffCodeError()
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
	localSpecularMeta.structPathAccountCreateInvalidNameError, err = pk.NewType(
		"AccountCreateInvalidNameError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountCreateInvalidNameError()
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
	localSpecularMeta.structPathAccountBeginAssumeIdentityInput, err = pk.NewType(
		"AccountBeginAssumeIdentityInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountBeginAssumeIdentityInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountBeginAssumeIdentityOutput, err = pk.NewType(
		"AccountBeginAssumeIdentityOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountBeginAssumeIdentityOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountCompleteAssumeIdentityInput, err = pk.NewType(
		"AccountCompleteAssumeIdentityInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountCompleteAssumeIdentityInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountCompleteAssumeIdentityOutput, err = pk.NewType(
		"AccountCompleteAssumeIdentityOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountCompleteAssumeIdentityOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError, err = pk.NewType(
		"AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError()
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
	localSpecularMeta.structPathAccountSSOBeginAuthenticationParameterError, err = pk.NewType(
		"AccountSSOBeginAuthenticationParameterError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountSSOBeginAuthenticationParameterError()
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
	localSpecularMeta.structPathAccountSSOCompleteAuthenticationInvalidFlowError, err = pk.NewType(
		"AccountSSOCompleteAuthenticationInvalidFlowError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountSSOCompleteAuthenticationInvalidFlowError()
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
	localSpecularMeta.structPathAccountOIDCProviderCreateInput, err = pk.NewType(
		"AccountOIDCProviderCreateInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountOIDCProviderCreateInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountOIDCProviderCreateOutput, err = pk.NewType(
		"AccountOIDCProviderCreateOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountOIDCProviderCreateOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountOIDCProviderSetAudiencesInput, err = pk.NewType(
		"AccountOIDCProviderSetAudiencesInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountOIDCProviderSetAudiencesInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountOIDCProviderSetAudiencesOutput, err = pk.NewType(
		"AccountOIDCProviderSetAudiencesOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountOIDCProviderSetAudiencesOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountOIDCProviderListInput, err = pk.NewType(
		"AccountOIDCProviderListInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountOIDCProviderListInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountOIDCProviderListOutput, err = pk.NewType(
		"AccountOIDCProviderListOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountOIDCProviderListOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountOIDCProviderDeleteInput, err = pk.NewType(
		"AccountOIDCProviderDeleteInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountOIDCProviderDeleteInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountOIDCProviderDeleteOutput, err = pk.NewType(
		"AccountOIDCProviderDeleteOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountOIDCProviderDeleteOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountOIDCProviderTrustPolicyListInput, err = pk.NewType(
		"AccountOIDCProviderTrustPolicyListInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountOIDCProviderTrustPolicyListInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathAccountOIDCProviderTrustPolicyListOutput, err = pk.NewType(
		"AccountOIDCProviderTrustPolicyListOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewAccountOIDCProviderTrustPolicyListOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRegionListInput, err = pk.NewType(
		"RegionListInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRegionListInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRegionListOutput, err = pk.NewType(
		"RegionListOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRegionListOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvalidUsernameError, err = pk.NewType(
		"InvalidUsernameError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvalidUsernameError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvalidRoleNameError, err = pk.NewType(
		"InvalidRoleNameError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvalidRoleNameError()
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
	localSpecularMeta.structPathIdentityPolicyAttachmentInfo, err = pk.NewType(
		"IdentityPolicyAttachmentInfo",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewIdentityPolicyAttachmentInfo()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserNotFoundError, err = pk.NewType(
		"UserNotFoundError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserNotFoundError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathCannotDisableSelfError, err = pk.NewType(
		"CannotDisableSelfError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewCannotDisableSelfError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleNotFoundError, err = pk.NewType(
		"RoleNotFoundError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleNotFoundError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathGroupNotFoundError, err = pk.NewType(
		"GroupNotFoundError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewGroupNotFoundError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvalidGroupNameError, err = pk.NewType(
		"InvalidGroupNameError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvalidGroupNameError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvalidGroupFilterError, err = pk.NewType(
		"InvalidGroupFilterError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvalidGroupFilterError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvalidUserFilterError, err = pk.NewType(
		"InvalidUserFilterError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvalidUserFilterError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathIdentityInUseError, err = pk.NewType(
		"IdentityInUseError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewIdentityInUseError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserSSHKey, err = pk.NewType(
		"UserSSHKey",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserSSHKey()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvalidSSHPublicKeyError, err = pk.NewType(
		"InvalidSSHPublicKeyError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvalidSSHPublicKeyError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathSSHKeyAlreadyExistsError, err = pk.NewType(
		"SSHKeyAlreadyExistsError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewSSHKeyAlreadyExistsError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvalidSSHKeyTitleError, err = pk.NewType(
		"InvalidSSHKeyTitleError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvalidSSHKeyTitleError()
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
	localSpecularMeta.structPathUserGetUserNotAvailableError, err = pk.NewType(
		"UserGetUserNotAvailableError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserGetUserNotAvailableError()
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
	localSpecularMeta.structPathUserSetActiveInput, err = pk.NewType(
		"UserSetActiveInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserSetActiveInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserSetActiveOutput, err = pk.NewType(
		"UserSetActiveOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserSetActiveOutput()
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
	localSpecularMeta.structPathUserSSHKeyCreateInput, err = pk.NewType(
		"UserSSHKeyCreateInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserSSHKeyCreateInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserSSHKeyCreateOutput, err = pk.NewType(
		"UserSSHKeyCreateOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserSSHKeyCreateOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserSSHKeyListInput, err = pk.NewType(
		"UserSSHKeyListInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserSSHKeyListInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserSSHKeyListOutput, err = pk.NewType(
		"UserSSHKeyListOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserSSHKeyListOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserSSHKeyDestroyInput, err = pk.NewType(
		"UserSSHKeyDestroyInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserSSHKeyDestroyInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserSSHKeyDestroyOutput, err = pk.NewType(
		"UserSSHKeyDestroyOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserSSHKeyDestroyOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserSSHKeySetTitleInput, err = pk.NewType(
		"UserSSHKeySetTitleInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserSSHKeySetTitleInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserSSHKeySetTitleOutput, err = pk.NewType(
		"UserSSHKeySetTitleOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserSSHKeySetTitleOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserGroupListInput, err = pk.NewType(
		"UserGroupListInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserGroupListInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathUserGroupListOutput, err = pk.NewType(
		"UserGroupListOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewUserGroupListOutput()
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
	localSpecularMeta.structPathPolicyStructureError, err = pk.NewType(
		"PolicyStructureError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewPolicyStructureError()
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
	localSpecularMeta.structPathRoleAssumeWithWebIdentityInput, err = pk.NewType(
		"RoleAssumeWithWebIdentityInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleAssumeWithWebIdentityInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleAssumeWithWebIdentityOutput, err = pk.NewType(
		"RoleAssumeWithWebIdentityOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleAssumeWithWebIdentityOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleAccessKeyListInput, err = pk.NewType(
		"RoleAccessKeyListInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleAccessKeyListInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleAccessKeyListOutput, err = pk.NewType(
		"RoleAccessKeyListOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleAccessKeyListOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleAccessKeyDestroyInput, err = pk.NewType(
		"RoleAccessKeyDestroyInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleAccessKeyDestroyInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleAccessKeyDestroyOutput, err = pk.NewType(
		"RoleAccessKeyDestroyOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleAccessKeyDestroyOutput()
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
	localSpecularMeta.structPathRoleTrustPolicyAttachInput, err = pk.NewType(
		"RoleTrustPolicyAttachInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleTrustPolicyAttachInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleTrustPolicyAttachOutput, err = pk.NewType(
		"RoleTrustPolicyAttachOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleTrustPolicyAttachOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleTrustPolicyListInput, err = pk.NewType(
		"RoleTrustPolicyListInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleTrustPolicyListInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleTrustPolicyListOutput, err = pk.NewType(
		"RoleTrustPolicyListOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleTrustPolicyListOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleTrustPolicyDetachInput, err = pk.NewType(
		"RoleTrustPolicyDetachInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleTrustPolicyDetachInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathRoleTrustPolicyDetachOutput, err = pk.NewType(
		"RoleTrustPolicyDetachOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewRoleTrustPolicyDetachOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathGroupCreateInput, err = pk.NewType(
		"GroupCreateInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewGroupCreateInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathGroupCreateOutput, err = pk.NewType(
		"GroupCreateOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewGroupCreateOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathGroupDestroyInput, err = pk.NewType(
		"GroupDestroyInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewGroupDestroyInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathGroupDestroyOutput, err = pk.NewType(
		"GroupDestroyOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewGroupDestroyOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathGroupGetInput, err = pk.NewType(
		"GroupGetInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewGroupGetInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathGroupGetOutput, err = pk.NewType(
		"GroupGetOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewGroupGetOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathGroupListInput, err = pk.NewType(
		"GroupListInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewGroupListInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathGroupListOutput, err = pk.NewType(
		"GroupListOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewGroupListOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathGroupMemberAddInput, err = pk.NewType(
		"GroupMemberAddInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewGroupMemberAddInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathGroupMemberAddOutput, err = pk.NewType(
		"GroupMemberAddOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewGroupMemberAddOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathGroupMemberRemoveInput, err = pk.NewType(
		"GroupMemberRemoveInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewGroupMemberRemoveInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathGroupMemberRemoveOutput, err = pk.NewType(
		"GroupMemberRemoveOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewGroupMemberRemoveOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathGroupMemberListInput, err = pk.NewType(
		"GroupMemberListInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewGroupMemberListInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathGroupMemberListOutput, err = pk.NewType(
		"GroupMemberListOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewGroupMemberListOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathGroupIdentityPolicyAttachInput, err = pk.NewType(
		"GroupIdentityPolicyAttachInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewGroupIdentityPolicyAttachInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathGroupIdentityPolicyAttachOutput, err = pk.NewType(
		"GroupIdentityPolicyAttachOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewGroupIdentityPolicyAttachOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathGroupIdentityPolicyDetachInput, err = pk.NewType(
		"GroupIdentityPolicyDetachInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewGroupIdentityPolicyDetachInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathGroupIdentityPolicyDetachOutput, err = pk.NewType(
		"GroupIdentityPolicyDetachOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewGroupIdentityPolicyDetachOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathGroupIdentityPolicyListInput, err = pk.NewType(
		"GroupIdentityPolicyListInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewGroupIdentityPolicyListInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathGroupIdentityPolicyListOutput, err = pk.NewType(
		"GroupIdentityPolicyListOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewGroupIdentityPolicyListOutput()
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
	localSpecularMeta.structPathIdentityPolicyAttachmentListInput, err = pk.NewType(
		"IdentityPolicyAttachmentListInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewIdentityPolicyAttachmentListInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathIdentityPolicyAttachmentListOutput, err = pk.NewType(
		"IdentityPolicyAttachmentListOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewIdentityPolicyAttachmentListOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathTrustPolicyCreateInput, err = pk.NewType(
		"TrustPolicyCreateInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewTrustPolicyCreateInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathTrustPolicyCreateOutput, err = pk.NewType(
		"TrustPolicyCreateOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewTrustPolicyCreateOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathTrustPolicyRetrieveInput, err = pk.NewType(
		"TrustPolicyRetrieveInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewTrustPolicyRetrieveInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathTrustPolicyRetrieveOutput, err = pk.NewType(
		"TrustPolicyRetrieveOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewTrustPolicyRetrieveOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathTrustPolicyListInput, err = pk.NewType(
		"TrustPolicyListInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewTrustPolicyListInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathTrustPolicyListOutput, err = pk.NewType(
		"TrustPolicyListOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewTrustPolicyListOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathTrustPolicyUpdateInput, err = pk.NewType(
		"TrustPolicyUpdateInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewTrustPolicyUpdateInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathTrustPolicyUpdateOutput, err = pk.NewType(
		"TrustPolicyUpdateOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewTrustPolicyUpdateOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathTrustPolicyDestroyInput, err = pk.NewType(
		"TrustPolicyDestroyInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewTrustPolicyDestroyInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathTrustPolicyDestroyOutput, err = pk.NewType(
		"TrustPolicyDestroyOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewTrustPolicyDestroyOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathTrustPolicyAttachmentListInput, err = pk.NewType(
		"TrustPolicyAttachmentListInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewTrustPolicyAttachmentListInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathTrustPolicyAttachmentListOutput, err = pk.NewType(
		"TrustPolicyAttachmentListOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewTrustPolicyAttachmentListOutput()
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
	localSpecularMeta.structPathInvalidServiceBearerTokenDurationError, err = pk.NewType(
		"InvalidServiceBearerTokenDurationError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvalidServiceBearerTokenDurationError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvalidServiceNameError, err = pk.NewType(
		"InvalidServiceNameError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvalidServiceNameError()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvitationCreateInput, err = pk.NewType(
		"InvitationCreateInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvitationCreateInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvitationCreateOutput, err = pk.NewType(
		"InvitationCreateOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvitationCreateOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvitationListInput, err = pk.NewType(
		"InvitationListInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvitationListInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvitationListOutput, err = pk.NewType(
		"InvitationListOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvitationListOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvitationRevokeInput, err = pk.NewType(
		"InvitationRevokeInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvitationRevokeInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvitationRevokeOutput, err = pk.NewType(
		"InvitationRevokeOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvitationRevokeOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvitationAcceptInput, err = pk.NewType(
		"InvitationAcceptInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvitationAcceptInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvitationAcceptOutput, err = pk.NewType(
		"InvitationAcceptOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvitationAcceptOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvitationInspectInput, err = pk.NewType(
		"InvitationInspectInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvitationInspectInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvitationInspectOutput, err = pk.NewType(
		"InvitationInspectOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvitationInspectOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvitationDeclineInput, err = pk.NewType(
		"InvitationDeclineInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvitationDeclineInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathInvitationDeclineOutput, err = pk.NewType(
		"InvitationDeclineOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewInvitationDeclineOutput()
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
	localSpecularMeta.structPathSessionKeepAliveInput, err = pk.NewType(
		"SessionKeepAliveInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewSessionKeepAliveInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathSessionKeepAliveOutput, err = pk.NewType(
		"SessionKeepAliveOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewSessionKeepAliveOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathSessionRevokeInput, err = pk.NewType(
		"SessionRevokeInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewSessionRevokeInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathSessionRevokeOutput, err = pk.NewType(
		"SessionRevokeOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewSessionRevokeOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathSessionIdentityInput, err = pk.NewType(
		"SessionIdentityInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewSessionIdentityInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathSessionIdentityOutput, err = pk.NewType(
		"SessionIdentityOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewSessionIdentityOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathSessionBeginHandoffInput, err = pk.NewType(
		"SessionBeginHandoffInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewSessionBeginHandoffInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathSessionBeginHandoffOutput, err = pk.NewType(
		"SessionBeginHandoffOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewSessionBeginHandoffOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathSessionCompleteHandoffInput, err = pk.NewType(
		"SessionCompleteHandoffInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewSessionCompleteHandoffInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathSessionCompleteHandoffOutput, err = pk.NewType(
		"SessionCompleteHandoffOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewSessionCompleteHandoffOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathServiceCatalogListInput, err = pk.NewType(
		"ServiceCatalogListInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewServiceCatalogListInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathServiceCatalogListOutput, err = pk.NewType(
		"ServiceCatalogListOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewServiceCatalogListOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathServiceCatalogGetInput, err = pk.NewType(
		"ServiceCatalogGetInput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewServiceCatalogGetInput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathServiceCatalogGetOutput, err = pk.NewType(
		"ServiceCatalogGetOutput",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewServiceCatalogGetOutput()
		}),
	)
	if err != nil {
		return nil, err
	}
	localSpecularMeta.structPathServiceCatalogGetServiceCatalogNotFoundError, err = pk.NewType(
		"ServiceCatalogGetServiceCatalogNotFoundError",
		clientruntime.TypeBuilder(func() clientruntime.Struct {
			return NewServiceCatalogGetServiceCatalogNotFoundError()
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
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().AccountCreateInvalidNameErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resAccount.NewOperation("AssumeIdentity")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountAssumeIdentityInputStruct())
	op.SetOutput(SpecularMeta().AccountAssumeIdentityInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resAccount.NewOperation("BeginAssumeIdentity")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountBeginAssumeIdentityInputStruct())
	op.SetOutput(SpecularMeta().AccountBeginAssumeIdentityInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resAccount.NewOperation("CompleteAssumeIdentity")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountCompleteAssumeIdentityInputStruct())
	op.SetOutput(SpecularMeta().AccountCompleteAssumeIdentityInputStruct())
	op.RegisterProblemType(SpecularMeta().AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeErrorStruct())

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
	op.RegisterProblemType(SpecularMeta().SSOProviderUnavailableErrorStruct())
	op.RegisterProblemType(SpecularMeta().AccountSSOBeginAuthenticationParameterErrorStruct())

	op, err = resAccountSSO.NewOperation("CompleteAuthentication")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountSSOCompleteAuthenticationInputStruct())
	op.SetOutput(SpecularMeta().AccountSSOCompleteAuthenticationInputStruct())
	op.RegisterProblemType(SpecularMeta().AccountSSOCompleteAuthenticationInvalidFlowErrorStruct())

	op, err = resAccountSSO.NewOperation("GetProviders")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountSSOGetProvidersInputStruct())
	op.SetOutput(SpecularMeta().AccountSSOGetProvidersInputStruct())

	// subresource AccountOIDCProvider
	resAccountOIDCProvider, err := resAccount.NewSubResource("OIDCProvider")
	if err != nil {
		return nil, err
	}
	_ = resAccountOIDCProvider

	op, err = resAccountOIDCProvider.NewOperation("Create")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountOIDCProviderCreateInputStruct())
	op.SetOutput(SpecularMeta().AccountOIDCProviderCreateInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidOIDCProviderErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidOIDCIssuerErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resAccountOIDCProvider.NewOperation("SetAudiences")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountOIDCProviderSetAudiencesInputStruct())
	op.SetOutput(SpecularMeta().AccountOIDCProviderSetAudiencesInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().OIDCProviderNotFoundErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidOIDCProviderErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resAccountOIDCProvider.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountOIDCProviderListInputStruct())
	op.SetOutput(SpecularMeta().AccountOIDCProviderListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resAccountOIDCProvider.NewOperation("Delete")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountOIDCProviderDeleteInputStruct())
	op.SetOutput(SpecularMeta().AccountOIDCProviderDeleteInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().OIDCProviderNotFoundErrorStruct())
	op.RegisterProblemType(SpecularMeta().OIDCProviderInUseErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})
	// subresource AccountOIDCProviderTrustPolicy
	resAccountOIDCProviderTrustPolicy, err := resAccountOIDCProvider.NewSubResource("TrustPolicy")
	if err != nil {
		return nil, err
	}
	_ = resAccountOIDCProviderTrustPolicy

	op, err = resAccountOIDCProviderTrustPolicy.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().AccountOIDCProviderTrustPolicyListInputStruct())
	op.SetOutput(SpecularMeta().AccountOIDCProviderTrustPolicyListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().OIDCProviderNotFoundErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	resRegion, err := pk.NewResource("Region")
	if err != nil {
		return nil, err
	}
	_ = resRegion

	op, err = resRegion.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().RegionListInputStruct())
	op.SetOutput(SpecularMeta().RegionListInputStruct())

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
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidUsernameErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUser.NewOperation("Get")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserGetInputStruct())
	op.SetOutput(SpecularMeta().UserGetInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(SpecularMeta().UserGetUserNotAvailableErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUser.NewOperation("Destroy")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserDestroyInputStruct())
	op.SetOutput(SpecularMeta().UserDestroyInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().IdentityInUseErrorStruct())
	op.RegisterProblemType(SpecularMeta().UserNotFoundErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUser.NewOperation("SetActive")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserSetActiveInputStruct())
	op.SetOutput(SpecularMeta().UserSetActiveInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().UserNotFoundErrorStruct())
	op.RegisterProblemType(SpecularMeta().CannotDisableSelfErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUser.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserListInputStruct())
	op.SetOutput(SpecularMeta().UserListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidUserFilterErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUser.NewOperation("MemberAccounts")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserMemberAccountsInputStruct())
	op.SetOutput(SpecularMeta().UserMemberAccountsInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())

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
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().UserNotFoundErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUserAccessKey.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserAccessKeyListInputStruct())
	op.SetOutput(SpecularMeta().UserAccessKeyListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUserAccessKey.NewOperation("Destroy")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserAccessKeyDestroyInputStruct())
	op.SetOutput(SpecularMeta().UserAccessKeyDestroyInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())

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
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().PolicyNotFoundErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUserIdentityPolicy.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserIdentityPolicyListInputStruct())
	op.SetOutput(SpecularMeta().UserIdentityPolicyListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUserIdentityPolicy.NewOperation("Detach")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserIdentityPolicyDetachInputStruct())
	op.SetOutput(SpecularMeta().UserIdentityPolicyDetachInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().PolicyNotFoundErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	// subresource UserSSHKey
	resUserSSHKey, err := resUser.NewSubResource("SSHKey")
	if err != nil {
		return nil, err
	}
	_ = resUserSSHKey

	op, err = resUserSSHKey.NewOperation("Create")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserSSHKeyCreateInputStruct())
	op.SetOutput(SpecularMeta().UserSSHKeyCreateInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().UserNotFoundErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidSSHPublicKeyErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidSSHKeyTitleErrorStruct())
	op.RegisterProblemType(SpecularMeta().SSHKeyAlreadyExistsErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUserSSHKey.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserSSHKeyListInputStruct())
	op.SetOutput(SpecularMeta().UserSSHKeyListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUserSSHKey.NewOperation("Destroy")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserSSHKeyDestroyInputStruct())
	op.SetOutput(SpecularMeta().UserSSHKeyDestroyInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resUserSSHKey.NewOperation("SetTitle")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserSSHKeySetTitleInputStruct())
	op.SetOutput(SpecularMeta().UserSSHKeySetTitleInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidSSHKeyTitleErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	// subresource UserGroup
	resUserGroup, err := resUser.NewSubResource("Group")
	if err != nil {
		return nil, err
	}
	_ = resUserGroup

	op, err = resUserGroup.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().UserGroupListInputStruct())
	op.SetOutput(SpecularMeta().UserGroupListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().UserNotFoundErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidUsernameErrorStruct())

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
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidRoleNameErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resRole.NewOperation("Get")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().RoleGetInputStruct())
	op.SetOutput(SpecularMeta().RoleGetInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resRole.NewOperation("Destroy")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().RoleDestroyInputStruct())
	op.SetOutput(SpecularMeta().RoleDestroyInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().IdentityInUseErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resRole.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().RoleListInputStruct())
	op.SetOutput(SpecularMeta().RoleListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resRole.NewOperation("Assume")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().RoleAssumeInputStruct())
	op.SetOutput(SpecularMeta().RoleAssumeInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().PolicyStructureErrorStruct())
	op.RegisterProblemType(SpecularMeta().RoleNotFoundErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidPrincipalDRNErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resRole.NewOperation("AssumeWithWebIdentity")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().RoleAssumeWithWebIdentityInputStruct())
	op.SetOutput(SpecularMeta().RoleAssumeWithWebIdentityInputStruct())
	op.RegisterProblemType(SpecularMeta().InvalidWebIdentityTokenErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidPrincipalDRNErrorStruct())
	op.RegisterProblemType(SpecularMeta().PolicyStructureErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())

	// subresource RoleAccessKey
	resRoleAccessKey, err := resRole.NewSubResource("AccessKey")
	if err != nil {
		return nil, err
	}
	_ = resRoleAccessKey

	op, err = resRoleAccessKey.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().RoleAccessKeyListInputStruct())
	op.SetOutput(SpecularMeta().RoleAccessKeyListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().RoleNotFoundErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resRoleAccessKey.NewOperation("Destroy")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().RoleAccessKeyDestroyInputStruct())
	op.SetOutput(SpecularMeta().RoleAccessKeyDestroyInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().RoleNotFoundErrorStruct())

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
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().PolicyNotFoundErrorStruct())
	op.RegisterProblemType(SpecularMeta().RoleNotFoundErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resRoleIdentityPolicy.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().RoleIdentityPolicyListInputStruct())
	op.SetOutput(SpecularMeta().RoleIdentityPolicyListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resRoleIdentityPolicy.NewOperation("Detach")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().RoleIdentityPolicyDetachInputStruct())
	op.SetOutput(SpecularMeta().RoleIdentityPolicyDetachInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().PolicyNotFoundErrorStruct())
	op.RegisterProblemType(SpecularMeta().RoleNotFoundErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	// subresource RoleTrustPolicy
	resRoleTrustPolicy, err := resRole.NewSubResource("TrustPolicy")
	if err != nil {
		return nil, err
	}
	_ = resRoleTrustPolicy

	op, err = resRoleTrustPolicy.NewOperation("Attach")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().RoleTrustPolicyAttachInputStruct())
	op.SetOutput(SpecularMeta().RoleTrustPolicyAttachInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().TrustPolicyNotFoundErrorStruct())
	op.RegisterProblemType(SpecularMeta().RoleNotFoundErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resRoleTrustPolicy.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().RoleTrustPolicyListInputStruct())
	op.SetOutput(SpecularMeta().RoleTrustPolicyListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().RoleNotFoundErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resRoleTrustPolicy.NewOperation("Detach")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().RoleTrustPolicyDetachInputStruct())
	op.SetOutput(SpecularMeta().RoleTrustPolicyDetachInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().TrustPolicyNotFoundErrorStruct())
	op.RegisterProblemType(SpecularMeta().RoleNotFoundErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	resGroup, err := pk.NewResource("Group")
	if err != nil {
		return nil, err
	}
	_ = resGroup

	op, err = resGroup.NewOperation("Create")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().GroupCreateInputStruct())
	op.SetOutput(SpecularMeta().GroupCreateInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidGroupNameErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resGroup.NewOperation("Destroy")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().GroupDestroyInputStruct())
	op.SetOutput(SpecularMeta().GroupDestroyInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().GroupNotFoundErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidGroupNameErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resGroup.NewOperation("Get")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().GroupGetInputStruct())
	op.SetOutput(SpecularMeta().GroupGetInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().GroupNotFoundErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidGroupNameErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resGroup.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().GroupListInputStruct())
	op.SetOutput(SpecularMeta().GroupListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidGroupFilterErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})
	// subresource GroupMember
	resGroupMember, err := resGroup.NewSubResource("Member")
	if err != nil {
		return nil, err
	}
	_ = resGroupMember

	op, err = resGroupMember.NewOperation("Add")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().GroupMemberAddInputStruct())
	op.SetOutput(SpecularMeta().GroupMemberAddInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().GroupNotFoundErrorStruct())
	op.RegisterProblemType(SpecularMeta().UserNotFoundErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidGroupNameErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidUsernameErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resGroupMember.NewOperation("Remove")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().GroupMemberRemoveInputStruct())
	op.SetOutput(SpecularMeta().GroupMemberRemoveInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().GroupNotFoundErrorStruct())
	op.RegisterProblemType(SpecularMeta().UserNotFoundErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidGroupNameErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidUsernameErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resGroupMember.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().GroupMemberListInputStruct())
	op.SetOutput(SpecularMeta().GroupMemberListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().GroupNotFoundErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidGroupNameErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	// subresource GroupIdentityPolicy
	resGroupIdentityPolicy, err := resGroup.NewSubResource("IdentityPolicy")
	if err != nil {
		return nil, err
	}
	_ = resGroupIdentityPolicy

	op, err = resGroupIdentityPolicy.NewOperation("Attach")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().GroupIdentityPolicyAttachInputStruct())
	op.SetOutput(SpecularMeta().GroupIdentityPolicyAttachInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().GroupNotFoundErrorStruct())
	op.RegisterProblemType(SpecularMeta().PolicyNotFoundErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidGroupNameErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resGroupIdentityPolicy.NewOperation("Detach")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().GroupIdentityPolicyDetachInputStruct())
	op.SetOutput(SpecularMeta().GroupIdentityPolicyDetachInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().GroupNotFoundErrorStruct())
	op.RegisterProblemType(SpecularMeta().PolicyNotFoundErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidGroupNameErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resGroupIdentityPolicy.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().GroupIdentityPolicyListInputStruct())
	op.SetOutput(SpecularMeta().GroupIdentityPolicyListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().GroupNotFoundErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidGroupNameErrorStruct())

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
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().PolicyStructureErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resIdentityPolicy.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().IdentityPolicyListInputStruct())
	op.SetOutput(SpecularMeta().IdentityPolicyListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resIdentityPolicy.NewOperation("Retrieve")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().IdentityPolicyRetrieveInputStruct())
	op.SetOutput(SpecularMeta().IdentityPolicyRetrieveInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().PolicyNotFoundErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resIdentityPolicy.NewOperation("Destroy")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().IdentityPolicyDestroyInputStruct())
	op.SetOutput(SpecularMeta().IdentityPolicyDestroyInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resIdentityPolicy.NewOperation("Update")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().IdentityPolicyUpdateInputStruct())
	op.SetOutput(SpecularMeta().IdentityPolicyUpdateInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().PolicyStructureErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})
	// subresource IdentityPolicyAttachment
	resIdentityPolicyAttachment, err := resIdentityPolicy.NewSubResource("Attachment")
	if err != nil {
		return nil, err
	}
	_ = resIdentityPolicyAttachment

	op, err = resIdentityPolicyAttachment.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().IdentityPolicyAttachmentListInputStruct())
	op.SetOutput(SpecularMeta().IdentityPolicyAttachmentListInputStruct())
	op.RegisterProblemType(SpecularMeta().PolicyNotFoundErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	resTrustPolicy, err := pk.NewResource("TrustPolicy")
	if err != nil {
		return nil, err
	}
	_ = resTrustPolicy

	op, err = resTrustPolicy.NewOperation("Create")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().TrustPolicyCreateInputStruct())
	op.SetOutput(SpecularMeta().TrustPolicyCreateInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidTrustPolicyErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resTrustPolicy.NewOperation("Retrieve")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().TrustPolicyRetrieveInputStruct())
	op.SetOutput(SpecularMeta().TrustPolicyRetrieveInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().TrustPolicyNotFoundErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resTrustPolicy.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().TrustPolicyListInputStruct())
	op.SetOutput(SpecularMeta().TrustPolicyListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resTrustPolicy.NewOperation("Update")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().TrustPolicyUpdateInputStruct())
	op.SetOutput(SpecularMeta().TrustPolicyUpdateInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidTrustPolicyErrorStruct())
	op.RegisterProblemType(SpecularMeta().TrustPolicyNotFoundErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resTrustPolicy.NewOperation("Destroy")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().TrustPolicyDestroyInputStruct())
	op.SetOutput(SpecularMeta().TrustPolicyDestroyInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().TrustPolicyNotFoundErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})
	// subresource TrustPolicyAttachment
	resTrustPolicyAttachment, err := resTrustPolicy.NewSubResource("Attachment")
	if err != nil {
		return nil, err
	}
	_ = resTrustPolicyAttachment

	op, err = resTrustPolicyAttachment.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().TrustPolicyAttachmentListInputStruct())
	op.SetOutput(SpecularMeta().TrustPolicyAttachmentListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().TrustPolicyNotFoundErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	resInvitation, err := pk.NewResource("Invitation")
	if err != nil {
		return nil, err
	}
	_ = resInvitation

	op, err = resInvitation.NewOperation("Create")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().InvitationCreateInputStruct())
	op.SetOutput(SpecularMeta().InvitationCreateInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidInvitationErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resInvitation.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().InvitationListInputStruct())
	op.SetOutput(SpecularMeta().InvitationListInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resInvitation.NewOperation("Revoke")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().InvitationRevokeInputStruct())
	op.SetOutput(SpecularMeta().InvitationRevokeInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvitationNotFoundErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resInvitation.NewOperation("Accept")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().InvitationAcceptInputStruct())
	op.SetOutput(SpecularMeta().InvitationAcceptInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidInvitationErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resInvitation.NewOperation("Inspect")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().InvitationInspectInputStruct())
	op.SetOutput(SpecularMeta().InvitationInspectInputStruct())
	op.RegisterProblemType(SpecularMeta().InvalidInvitationErrorStruct())

	op, err = resInvitation.NewOperation("Decline")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().InvitationDeclineInputStruct())
	op.SetOutput(SpecularMeta().InvitationDeclineInputStruct())
	op.RegisterProblemType(SpecularMeta().InvalidInvitationErrorStruct())

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
	op.RegisterProblemType(SpecularMeta().InvalidServiceBearerTokenDurationErrorStruct())
	op.RegisterProblemType(SpecularMeta().InvalidServiceNameErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	resSession, err := pk.NewResource("Session")
	if err != nil {
		return nil, err
	}
	_ = resSession

	op, err = resSession.NewOperation("KeepAlive")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().SessionKeepAliveInputStruct())
	op.SetOutput(SpecularMeta().SessionKeepAliveInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resSession.NewOperation("Revoke")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().SessionRevokeInputStruct())
	op.SetOutput(SpecularMeta().SessionRevokeInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resSession.NewOperation("Identity")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().SessionIdentityInputStruct())
	op.SetOutput(SpecularMeta().SessionIdentityInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resSession.NewOperation("BeginHandoff")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().SessionBeginHandoffInputStruct())
	op.SetOutput(SpecularMeta().SessionBeginHandoffInputStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().AccessDeniedErrorStruct())
	op.RegisterProblemType(godeployportcomapiservicescorelib.SpecularMeta().ForbiddenErrorStruct())

	op.AddAnnotation(&godeployportcomapiservicescorelib.SignedOperationV1{})

	op, err = resSession.NewOperation("CompleteHandoff")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().SessionCompleteHandoffInputStruct())
	op.SetOutput(SpecularMeta().SessionCompleteHandoffInputStruct())
	op.RegisterProblemType(SpecularMeta().InvalidHandoffCodeErrorStruct())

	resServiceCatalog, err := pk.NewResource("ServiceCatalog")
	if err != nil {
		return nil, err
	}
	_ = resServiceCatalog

	op, err = resServiceCatalog.NewOperation("List")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().ServiceCatalogListInputStruct())
	op.SetOutput(SpecularMeta().ServiceCatalogListInputStruct())

	op, err = resServiceCatalog.NewOperation("Get")
	if err != nil {
		return nil, err
	}

	op.SetInput(SpecularMeta().ServiceCatalogGetInputStruct())
	op.SetOutput(SpecularMeta().ServiceCatalogGetInputStruct())
	op.RegisterProblemType(SpecularMeta().ServiceCatalogGetServiceCatalogNotFoundErrorStruct())

	pk.AddAnnotation(&godeployportcomapiservicescorelib.ServiceSignatureV1{
		ServiceName: "iam",
	})
	return pk, nil
}

// AccountResourceClient is the AccountResourceClient resource client
type AccountResourceClient struct {
	transport              clientruntime.Transport
	res                    *clientruntime.Resource
	SSO                    *AccountSSOResourceClient
	OIDCProvider           *AccountOIDCProviderResourceClient
	create                 *clientruntime.Operation
	assumeIdentity         *clientruntime.Operation
	beginAssumeIdentity    *clientruntime.Operation
	completeAssumeIdentity *clientruntime.Operation
}

func newAccountResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*AccountResourceClient, error) {
	res := finder.FindResource("Account")
	r := &AccountResourceClient{
		transport: transport,
		res:       res,
	}
	var err error
	r.SSO, err = newAccountSSOResourceClient(transport, res)
	if err != nil {
		return nil, err
	}
	r.OIDCProvider, err = newAccountOIDCProviderResourceClient(transport, res)
	if err != nil {
		return nil, err
	}
	r.create = res.FindOperation("Create")
	r.assumeIdentity = res.FindOperation("AssumeIdentity")
	r.beginAssumeIdentity = res.FindOperation("BeginAssumeIdentity")
	r.completeAssumeIdentity = res.FindOperation("CompleteAssumeIdentity")
	return r, nil
}

// Create - Creates a new Account with the given name
// Create account only works for users of the "global" account
func (res *AccountResourceClient) Create(ctx context.Context, input *AccountCreateInput) (*AccountCreateOutput, error) {
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
func (res *AccountResourceClient) AssumeIdentity(ctx context.Context, input *AccountAssumeIdentityInput) (*AccountAssumeIdentityOutput, error) {
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

// BeginAssumeIdentity - Begins a cross-domain identity assumption. Authorizes the caller to assume
// into the given account and returns a single-use, short-lived opaque code to
// be redeemed via CompleteAssumeIdentity on the target domain. No credentials
// are minted until the code is redeemed.
func (res *AccountResourceClient) BeginAssumeIdentity(ctx context.Context, input *AccountBeginAssumeIdentityInput) (*AccountBeginAssumeIdentityOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.beginAssumeIdentity,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*AccountBeginAssumeIdentityOutput)
	return output, nil
}

// CompleteAssumeIdentity - Public: redeems a code minted by BeginAssumeIdentity for the assumed
// identity's credentials. Requires the target account name the code was issued
// for. No authentication (the redeeming domain has no credentials yet).
func (res *AccountResourceClient) CompleteAssumeIdentity(ctx context.Context, input *AccountCompleteAssumeIdentityInput) (*AccountCompleteAssumeIdentityOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.completeAssumeIdentity,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*AccountCompleteAssumeIdentityOutput)
	return output, nil
}

// AccountSSOResourceClient is the AccountSSOResourceClient resource client
type AccountSSOResourceClient struct {
	transport              clientruntime.Transport
	res                    *clientruntime.Resource
	beginAuthentication    *clientruntime.Operation
	completeAuthentication *clientruntime.Operation
	getProviders           *clientruntime.Operation
}

func newAccountSSOResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*AccountSSOResourceClient, error) {
	res := finder.FindResource("SSO")
	r := &AccountSSOResourceClient{
		transport: transport,
		res:       res,
	}
	r.beginAuthentication = res.FindOperation("BeginAuthentication")
	r.completeAuthentication = res.FindOperation("CompleteAuthentication")
	r.getProviders = res.FindOperation("GetProviders")
	return r, nil
}

// BeginAuthentication - Returns the public link to the login page for the given SSO provider
func (res *AccountSSOResourceClient) BeginAuthentication(ctx context.Context, input *AccountSSOBeginAuthenticationInput) (*AccountSSOBeginAuthenticationOutput, error) {
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
func (res *AccountSSOResourceClient) CompleteAuthentication(ctx context.Context, input *AccountSSOCompleteAuthenticationInput) (*AccountSSOCompleteAuthenticationOutput, error) {
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
func (res *AccountSSOResourceClient) GetProviders(ctx context.Context, input *AccountSSOGetProvidersInput) (*AccountSSOGetProvidersOutput, error) {
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

// AccountOIDCProviderResourceClient is the AccountOIDCProviderResourceClient resource client
type AccountOIDCProviderResourceClient struct {
	transport    clientruntime.Transport
	res          *clientruntime.Resource
	TrustPolicy  *AccountOIDCProviderTrustPolicyResourceClient
	create       *clientruntime.Operation
	setAudiences *clientruntime.Operation
	list         *clientruntime.Operation
	delete       *clientruntime.Operation
}

func newAccountOIDCProviderResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*AccountOIDCProviderResourceClient, error) {
	res := finder.FindResource("OIDCProvider")
	r := &AccountOIDCProviderResourceClient{
		transport: transport,
		res:       res,
	}
	var err error
	r.TrustPolicy, err = newAccountOIDCProviderTrustPolicyResourceClient(transport, res)
	if err != nil {
		return nil, err
	}
	r.create = res.FindOperation("Create")
	r.setAudiences = res.FindOperation("SetAudiences")
	r.list = res.FindOperation("List")
	r.delete = res.FindOperation("Delete")
	return r, nil
}

// Create - Registers an external OIDC identity provider in the current account. The
// issuer is probed via OIDC discovery to confirm it is reachable.
// audiences is the allowlist matched (exact, case-sensitive) against the
// token aud claim. It is an opaque identifier, not a URL: pick a stable
// value that identifies this IAM and set the same value as the token
// audience in your workload. We suggest iam.deployport.io.
// Requires permission action iam:CreateOIDCProvider over resource iam:OIDCProvider(<name>)
func (res *AccountOIDCProviderResourceClient) Create(ctx context.Context, input *AccountOIDCProviderCreateInput) (*AccountOIDCProviderCreateOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.create,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*AccountOIDCProviderCreateOutput)
	return output, nil
}

// SetAudiences - Replaces the audience allowlist of an existing OIDC provider. The issuer is
// immutable; only the audiences change. Trust policies reference the provider
// by name, so their attachments are unaffected.
// Requires permission action iam:UpdateOIDCProvider over resource iam:OIDCProvider(<name>)
func (res *AccountOIDCProviderResourceClient) SetAudiences(ctx context.Context, input *AccountOIDCProviderSetAudiencesInput) (*AccountOIDCProviderSetAudiencesOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.setAudiences,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*AccountOIDCProviderSetAudiencesOutput)
	return output, nil
}

// List - Lists the OIDC identity providers registered in the current account.
// Requires permission action iam:ListOIDCProviders over resource iam:OIDCProvider
func (res *AccountOIDCProviderResourceClient) List(ctx context.Context, input *AccountOIDCProviderListInput) (*AccountOIDCProviderListOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.list,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*AccountOIDCProviderListOutput)
	return output, nil
}

// Delete - Deletes an OIDC identity provider. Fails if a trust policy still references it.
// Requires permission action iam:DeleteOIDCProvider over resource iam:OIDCProvider(<name>)
func (res *AccountOIDCProviderResourceClient) Delete(ctx context.Context, input *AccountOIDCProviderDeleteInput) (*AccountOIDCProviderDeleteOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.delete,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*AccountOIDCProviderDeleteOutput)
	return output, nil
}

// AccountOIDCProviderTrustPolicyResourceClient is the AccountOIDCProviderTrustPolicyResourceClient resource client
type AccountOIDCProviderTrustPolicyResourceClient struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	list      *clientruntime.Operation
}

func newAccountOIDCProviderTrustPolicyResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*AccountOIDCProviderTrustPolicyResourceClient, error) {
	res := finder.FindResource("TrustPolicy")
	r := &AccountOIDCProviderTrustPolicyResourceClient{
		transport: transport,
		res:       res,
	}
	r.list = res.FindOperation("List")
	return r, nil
}

// List - Lists the trust policies whose principal is iam:OIDCProvider(<name>).
// Requires permission action iam:ListTrustPolicies over resource iam:OIDCProvider(<name>)
func (res *AccountOIDCProviderTrustPolicyResourceClient) List(ctx context.Context, input *AccountOIDCProviderTrustPolicyListInput) (*AccountOIDCProviderTrustPolicyListOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.list,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*AccountOIDCProviderTrustPolicyListOutput)
	return output, nil
}

// RegionResourceClient is the RegionResourceClient resource client
type RegionResourceClient struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	list      *clientruntime.Operation
}

func newRegionResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*RegionResourceClient, error) {
	res := finder.FindResource("Region")
	r := &RegionResourceClient{
		transport: transport,
		res:       res,
	}
	r.list = res.FindOperation("List")
	return r, nil
}

// List - Returns the list of available regions and their service endpoints.
// This is a public operation that does not require any authentication.
func (res *RegionResourceClient) List(ctx context.Context, input *RegionListInput) (*RegionListOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.list,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*RegionListOutput)
	return output, nil
}

// UserResourceClient is the UserResourceClient resource client
type UserResourceClient struct {
	transport      clientruntime.Transport
	res            *clientruntime.Resource
	AccessKey      *UserAccessKeyResourceClient
	IdentityPolicy *UserIdentityPolicyResourceClient
	SSHKey         *UserSSHKeyResourceClient
	Group          *UserGroupResourceClient
	create         *clientruntime.Operation
	get            *clientruntime.Operation
	destroy        *clientruntime.Operation
	setActive      *clientruntime.Operation
	list           *clientruntime.Operation
	memberAccounts *clientruntime.Operation
}

func newUserResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*UserResourceClient, error) {
	res := finder.FindResource("User")
	r := &UserResourceClient{
		transport: transport,
		res:       res,
	}
	var err error
	r.AccessKey, err = newUserAccessKeyResourceClient(transport, res)
	if err != nil {
		return nil, err
	}
	r.IdentityPolicy, err = newUserIdentityPolicyResourceClient(transport, res)
	if err != nil {
		return nil, err
	}
	r.SSHKey, err = newUserSSHKeyResourceClient(transport, res)
	if err != nil {
		return nil, err
	}
	r.Group, err = newUserGroupResourceClient(transport, res)
	if err != nil {
		return nil, err
	}
	r.create = res.FindOperation("Create")
	r.get = res.FindOperation("Get")
	r.destroy = res.FindOperation("Destroy")
	r.setActive = res.FindOperation("SetActive")
	r.list = res.FindOperation("List")
	r.memberAccounts = res.FindOperation("MemberAccounts")
	return r, nil
}

// Create - Creates a new user with the given username in the current account
// Requires permission action iam:CreateUser over resource iam:User(<username>)
func (res *UserResourceClient) Create(ctx context.Context, input *UserCreateInput) (*UserCreateOutput, error) {
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
// if the username is not provided, the current user is returned
// Requires permission action iam:GetRole over resource iam:Role(<name>)
func (res *UserResourceClient) Get(ctx context.Context, input *UserGetInput) (*UserGetOutput, error) {
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
func (res *UserResourceClient) Destroy(ctx context.Context, input *UserDestroyInput) (*UserDestroyOutput, error) {
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

// SetActive - Suspends or restores a user. A suspended (active:false) user cannot
// authenticate, its existing credentials go inert immediately, and any role
// session that originated from it stops working. Reversible: restoring it
// (active:true) brings its credentials back, unlike Destroy.
// Requires permission action iam:DisableUser (active:false) or iam:EnableUser
// (active:true) over resource iam:User(<username>)
func (res *UserResourceClient) SetActive(ctx context.Context, input *UserSetActiveInput) (*UserSetActiveOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.setActive,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*UserSetActiveOutput)
	return output, nil
}

// List - Returns the list of users in the current account
func (res *UserResourceClient) List(ctx context.Context, input *UserListInput) (*UserListOutput, error) {
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
func (res *UserResourceClient) MemberAccounts(ctx context.Context, input *UserMemberAccountsInput) (*UserMemberAccountsOutput, error) {
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

// UserAccessKeyResourceClient is the UserAccessKeyResourceClient resource client
type UserAccessKeyResourceClient struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	create    *clientruntime.Operation
	list      *clientruntime.Operation
	destroy   *clientruntime.Operation
}

func newUserAccessKeyResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*UserAccessKeyResourceClient, error) {
	res := finder.FindResource("AccessKey")
	r := &UserAccessKeyResourceClient{
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
func (res *UserAccessKeyResourceClient) Create(ctx context.Context, input *UserAccessKeyCreateInput) (*UserAccessKeyCreateOutput, error) {
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
func (res *UserAccessKeyResourceClient) List(ctx context.Context, input *UserAccessKeyListInput) (*UserAccessKeyListOutput, error) {
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
func (res *UserAccessKeyResourceClient) Destroy(ctx context.Context, input *UserAccessKeyDestroyInput) (*UserAccessKeyDestroyOutput, error) {
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

// UserIdentityPolicyResourceClient is the UserIdentityPolicyResourceClient resource client
type UserIdentityPolicyResourceClient struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	attach    *clientruntime.Operation
	list      *clientruntime.Operation
	detach    *clientruntime.Operation
}

func newUserIdentityPolicyResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*UserIdentityPolicyResourceClient, error) {
	res := finder.FindResource("IdentityPolicy")
	r := &UserIdentityPolicyResourceClient{
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
func (res *UserIdentityPolicyResourceClient) Attach(ctx context.Context, input *UserIdentityPolicyAttachInput) (*UserIdentityPolicyAttachOutput, error) {
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
func (res *UserIdentityPolicyResourceClient) List(ctx context.Context, input *UserIdentityPolicyListInput) (*UserIdentityPolicyListOutput, error) {
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
func (res *UserIdentityPolicyResourceClient) Detach(ctx context.Context, input *UserIdentityPolicyDetachInput) (*UserIdentityPolicyDetachOutput, error) {
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

// UserSSHKeyResourceClient is the UserSSHKeyResourceClient resource client
type UserSSHKeyResourceClient struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	create    *clientruntime.Operation
	list      *clientruntime.Operation
	destroy   *clientruntime.Operation
	setTitle  *clientruntime.Operation
}

func newUserSSHKeyResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*UserSSHKeyResourceClient, error) {
	res := finder.FindResource("SSHKey")
	r := &UserSSHKeyResourceClient{
		transport: transport,
		res:       res,
	}
	r.create = res.FindOperation("Create")
	r.list = res.FindOperation("List")
	r.destroy = res.FindOperation("Destroy")
	r.setTitle = res.FindOperation("SetTitle")
	return r, nil
}

// Create - Adds a public key to the given user
// Requires permission action iam:CreateUserSSHKey over resource iam:User(<username>).SSHKey
func (res *UserSSHKeyResourceClient) Create(ctx context.Context, input *UserSSHKeyCreateInput) (*UserSSHKeyCreateOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.create,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*UserSSHKeyCreateOutput)
	return output, nil
}

// List - Returns the SSH keys for the given user or the current user
// Requires permission action iam:ListUserSSHKeys over resource iam:User(<username>)
func (res *UserSSHKeyResourceClient) List(ctx context.Context, input *UserSSHKeyListInput) (*UserSSHKeyListOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.list,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*UserSSHKeyListOutput)
	return output, nil
}

// Destroy - Destroy an SSH key for the given user
// Requires permission action iam:DeleteUserSSHKey over resource iam:User(<username>).SSHKey(<fingerprint>)
func (res *UserSSHKeyResourceClient) Destroy(ctx context.Context, input *UserSSHKeyDestroyInput) (*UserSSHKeyDestroyOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.destroy,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*UserSSHKeyDestroyOutput)
	return output, nil
}

// SetTitle - Sets or clears what a person calls one of their keys
// Requires permission action iam:CreateUserSSHKey over resource iam:User(<username>).SSHKey(<fingerprint>)
func (res *UserSSHKeyResourceClient) SetTitle(ctx context.Context, input *UserSSHKeySetTitleInput) (*UserSSHKeySetTitleOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.setTitle,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*UserSSHKeySetTitleOutput)
	return output, nil
}

// UserGroupResourceClient is the UserGroupResourceClient resource client
type UserGroupResourceClient struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	list      *clientruntime.Operation
}

func newUserGroupResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*UserGroupResourceClient, error) {
	res := finder.FindResource("Group")
	r := &UserGroupResourceClient{
		transport: transport,
		res:       res,
	}
	r.list = res.FindOperation("List")
	return r, nil
}

// List - List returns the groups a user is a member of
// Requires permission action iam:ListGroups over resource iam:Group
func (res *UserGroupResourceClient) List(ctx context.Context, input *UserGroupListInput) (*UserGroupListOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.list,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*UserGroupListOutput)
	return output, nil
}

// RoleResourceClient is the RoleResourceClient resource client
type RoleResourceClient struct {
	transport             clientruntime.Transport
	res                   *clientruntime.Resource
	AccessKey             *RoleAccessKeyResourceClient
	IdentityPolicy        *RoleIdentityPolicyResourceClient
	TrustPolicy           *RoleTrustPolicyResourceClient
	create                *clientruntime.Operation
	get                   *clientruntime.Operation
	destroy               *clientruntime.Operation
	list                  *clientruntime.Operation
	assume                *clientruntime.Operation
	assumeWithWebIdentity *clientruntime.Operation
}

func newRoleResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*RoleResourceClient, error) {
	res := finder.FindResource("Role")
	r := &RoleResourceClient{
		transport: transport,
		res:       res,
	}
	var err error
	r.AccessKey, err = newRoleAccessKeyResourceClient(transport, res)
	if err != nil {
		return nil, err
	}
	r.IdentityPolicy, err = newRoleIdentityPolicyResourceClient(transport, res)
	if err != nil {
		return nil, err
	}
	r.TrustPolicy, err = newRoleTrustPolicyResourceClient(transport, res)
	if err != nil {
		return nil, err
	}
	r.create = res.FindOperation("Create")
	r.get = res.FindOperation("Get")
	r.destroy = res.FindOperation("Destroy")
	r.list = res.FindOperation("List")
	r.assume = res.FindOperation("Assume")
	r.assumeWithWebIdentity = res.FindOperation("AssumeWithWebIdentity")
	return r, nil
}

// Create - Creates a new role with the given name in the current account
// Requires permission action iam:CreateRole over resource iam:Role(<name>)
func (res *RoleResourceClient) Create(ctx context.Context, input *RoleCreateInput) (*RoleCreateOutput, error) {
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
func (res *RoleResourceClient) Get(ctx context.Context, input *RoleGetInput) (*RoleGetOutput, error) {
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
func (res *RoleResourceClient) Destroy(ctx context.Context, input *RoleDestroyInput) (*RoleDestroyOutput, error) {
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
func (res *RoleResourceClient) List(ctx context.Context, input *RoleListInput) (*RoleListOutput, error) {
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
// the caller needs to have permission of action iam:Assume over resource iam:Role(<roleName>)
func (res *RoleResourceClient) Assume(ctx context.Context, input *RoleAssumeInput) (*RoleAssumeOutput, error) {
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

// AssumeWithWebIdentity - Exchanges an external OIDC token for temporary credentials of the given role
// (AWS AssumeRoleWithWebIdentity). Public: the token itself is the proof of
// identity, no signed request is required. The role must have a trust policy
// attached whose principal and conditions admit the token's claims.
func (res *RoleResourceClient) AssumeWithWebIdentity(ctx context.Context, input *RoleAssumeWithWebIdentityInput) (*RoleAssumeWithWebIdentityOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.assumeWithWebIdentity,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*RoleAssumeWithWebIdentityOutput)
	return output, nil
}

// RoleAccessKeyResourceClient is the RoleAccessKeyResourceClient resource client
type RoleAccessKeyResourceClient struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	list      *clientruntime.Operation
	destroy   *clientruntime.Operation
}

func newRoleAccessKeyResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*RoleAccessKeyResourceClient, error) {
	res := finder.FindResource("AccessKey")
	r := &RoleAccessKeyResourceClient{
		transport: transport,
		res:       res,
	}
	r.list = res.FindOperation("List")
	r.destroy = res.FindOperation("Destroy")
	return r, nil
}

// List - Returns the list of access keys (assumed-role sessions) of the given role
// Requires permission action iam:ListAccessKeys over resource iam:Role(<name>)
func (res *RoleAccessKeyResourceClient) List(ctx context.Context, input *RoleAccessKeyListInput) (*RoleAccessKeyListOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.list,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*RoleAccessKeyListOutput)
	return output, nil
}

// Destroy - Destroys (revokes) an access key / assumed session of the given role.
// Requires permission action iam:DestroyRoleAccessKey over resource
// iam:Role(<name>).AccessKey(<accessKeyID>). Refused for a service-linked
// role: its sessions are managed by the owning service and cannot be revoked
// by the account.
func (res *RoleAccessKeyResourceClient) Destroy(ctx context.Context, input *RoleAccessKeyDestroyInput) (*RoleAccessKeyDestroyOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.destroy,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*RoleAccessKeyDestroyOutput)
	return output, nil
}

// RoleIdentityPolicyResourceClient is the RoleIdentityPolicyResourceClient resource client
type RoleIdentityPolicyResourceClient struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	attach    *clientruntime.Operation
	list      *clientruntime.Operation
	detach    *clientruntime.Operation
}

func newRoleIdentityPolicyResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*RoleIdentityPolicyResourceClient, error) {
	res := finder.FindResource("IdentityPolicy")
	r := &RoleIdentityPolicyResourceClient{
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
func (res *RoleIdentityPolicyResourceClient) Attach(ctx context.Context, input *RoleIdentityPolicyAttachInput) (*RoleIdentityPolicyAttachOutput, error) {
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
func (res *RoleIdentityPolicyResourceClient) List(ctx context.Context, input *RoleIdentityPolicyListInput) (*RoleIdentityPolicyListOutput, error) {
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
func (res *RoleIdentityPolicyResourceClient) Detach(ctx context.Context, input *RoleIdentityPolicyDetachInput) (*RoleIdentityPolicyDetachOutput, error) {
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

// RoleTrustPolicyResourceClient is the RoleTrustPolicyResourceClient resource client
type RoleTrustPolicyResourceClient struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	attach    *clientruntime.Operation
	list      *clientruntime.Operation
	detach    *clientruntime.Operation
}

func newRoleTrustPolicyResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*RoleTrustPolicyResourceClient, error) {
	res := finder.FindResource("TrustPolicy")
	r := &RoleTrustPolicyResourceClient{
		transport: transport,
		res:       res,
	}
	r.attach = res.FindOperation("Attach")
	r.list = res.FindOperation("List")
	r.detach = res.FindOperation("Detach")
	return r, nil
}

// Attach - Attaches a trust policy to a role. The role accepts web-identity assumption
// for any token admitted by an attached trust policy.
// Requires permission action iam:AttachRoleTrustPolicy over resource iam:Role(<name>).TrustPolicy(<trustPolicyName>)
func (res *RoleTrustPolicyResourceClient) Attach(ctx context.Context, input *RoleTrustPolicyAttachInput) (*RoleTrustPolicyAttachOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.attach,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*RoleTrustPolicyAttachOutput)
	return output, nil
}

// List - List returns the trust policies attached to a role.
// Requires permission action iam:ListRoleTrustPolicies over resource iam:Role(<rolename>).TrustPolicy
func (res *RoleTrustPolicyResourceClient) List(ctx context.Context, input *RoleTrustPolicyListInput) (*RoleTrustPolicyListOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.list,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*RoleTrustPolicyListOutput)
	return output, nil
}

// Detach - Detaches a trust policy from a role.
// Requires permission action iam:DetachRoleTrustPolicy over resource iam:Role(<name>).TrustPolicy(<trustPolicyName>)
func (res *RoleTrustPolicyResourceClient) Detach(ctx context.Context, input *RoleTrustPolicyDetachInput) (*RoleTrustPolicyDetachOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.detach,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*RoleTrustPolicyDetachOutput)
	return output, nil
}

// GroupResourceClient is the GroupResourceClient resource client
type GroupResourceClient struct {
	transport      clientruntime.Transport
	res            *clientruntime.Resource
	Member         *GroupMemberResourceClient
	IdentityPolicy *GroupIdentityPolicyResourceClient
	create         *clientruntime.Operation
	destroy        *clientruntime.Operation
	get            *clientruntime.Operation
	list           *clientruntime.Operation
}

func newGroupResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*GroupResourceClient, error) {
	res := finder.FindResource("Group")
	r := &GroupResourceClient{
		transport: transport,
		res:       res,
	}
	var err error
	r.Member, err = newGroupMemberResourceClient(transport, res)
	if err != nil {
		return nil, err
	}
	r.IdentityPolicy, err = newGroupIdentityPolicyResourceClient(transport, res)
	if err != nil {
		return nil, err
	}
	r.create = res.FindOperation("Create")
	r.destroy = res.FindOperation("Destroy")
	r.get = res.FindOperation("Get")
	r.list = res.FindOperation("List")
	return r, nil
}

// Create - Creates a new group in the current account.
// Requires permission action iam:CreateGroup over resource iam:Group(<name>)
func (res *GroupResourceClient) Create(ctx context.Context, input *GroupCreateInput) (*GroupCreateOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.create,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*GroupCreateOutput)
	return output, nil
}

// Destroy - Destroys a group. Cascades its memberships and policy attachments; the users
// and policies themselves are untouched.
// Requires permission action iam:DestroyGroup over resource iam:Group(<name>)
func (res *GroupResourceClient) Destroy(ctx context.Context, input *GroupDestroyInput) (*GroupDestroyOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.destroy,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*GroupDestroyOutput)
	return output, nil
}

// Get - Returns a single group by name.
// Requires permission action iam:GetGroup over resource iam:Group(<name>)
func (res *GroupResourceClient) Get(ctx context.Context, input *GroupGetInput) (*GroupGetOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.get,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*GroupGetOutput)
	return output, nil
}

// List - Lists the groups in the current account.
// Requires permission action iam:ListGroups over resource iam:Group
func (res *GroupResourceClient) List(ctx context.Context, input *GroupListInput) (*GroupListOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.list,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*GroupListOutput)
	return output, nil
}

// GroupMemberResourceClient is the GroupMemberResourceClient resource client
type GroupMemberResourceClient struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	add       *clientruntime.Operation
	remove    *clientruntime.Operation
	list      *clientruntime.Operation
}

func newGroupMemberResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*GroupMemberResourceClient, error) {
	res := finder.FindResource("Member")
	r := &GroupMemberResourceClient{
		transport: transport,
		res:       res,
	}
	r.add = res.FindOperation("Add")
	r.remove = res.FindOperation("Remove")
	r.list = res.FindOperation("List")
	return r, nil
}

// Add - Adds a user to the group.
// Requires permission action iam:AddGroupMember over resource iam:Group(<name>)
func (res *GroupMemberResourceClient) Add(ctx context.Context, input *GroupMemberAddInput) (*GroupMemberAddOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.add,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*GroupMemberAddOutput)
	return output, nil
}

// Remove - Removes a user from the group.
// Requires permission action iam:RemoveGroupMember over resource iam:Group(<name>)
func (res *GroupMemberResourceClient) Remove(ctx context.Context, input *GroupMemberRemoveInput) (*GroupMemberRemoveOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.remove,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*GroupMemberRemoveOutput)
	return output, nil
}

// List - Lists the usernames of a group's members.
// Requires permission action iam:ListGroups over resource iam:Group(<name>)
func (res *GroupMemberResourceClient) List(ctx context.Context, input *GroupMemberListInput) (*GroupMemberListOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.list,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*GroupMemberListOutput)
	return output, nil
}

// GroupIdentityPolicyResourceClient is the GroupIdentityPolicyResourceClient resource client
type GroupIdentityPolicyResourceClient struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	attach    *clientruntime.Operation
	detach    *clientruntime.Operation
	list      *clientruntime.Operation
}

func newGroupIdentityPolicyResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*GroupIdentityPolicyResourceClient, error) {
	res := finder.FindResource("IdentityPolicy")
	r := &GroupIdentityPolicyResourceClient{
		transport: transport,
		res:       res,
	}
	r.attach = res.FindOperation("Attach")
	r.detach = res.FindOperation("Detach")
	r.list = res.FindOperation("List")
	return r, nil
}

// Attach - Attaches an identity policy to the group.
// Requires permission action iam:AttachGroupIdentityPolicy over resource iam:Group(<name>).IdentityPolicy(<policyName>)
func (res *GroupIdentityPolicyResourceClient) Attach(ctx context.Context, input *GroupIdentityPolicyAttachInput) (*GroupIdentityPolicyAttachOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.attach,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*GroupIdentityPolicyAttachOutput)
	return output, nil
}

// Detach - Detaches an identity policy from the group.
// Requires permission action iam:DetachGroupIdentityPolicy over resource iam:Group(<name>).IdentityPolicy(<policyName>)
func (res *GroupIdentityPolicyResourceClient) Detach(ctx context.Context, input *GroupIdentityPolicyDetachInput) (*GroupIdentityPolicyDetachOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.detach,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*GroupIdentityPolicyDetachOutput)
	return output, nil
}

// List - Lists the identity policies attached to the group.
// Requires permission action iam:ListGroupPolicies over resource iam:Group(<name>).IdentityPolicy
func (res *GroupIdentityPolicyResourceClient) List(ctx context.Context, input *GroupIdentityPolicyListInput) (*GroupIdentityPolicyListOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.list,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*GroupIdentityPolicyListOutput)
	return output, nil
}

// IdentityPolicyResourceClient is the IdentityPolicyResourceClient resource client
type IdentityPolicyResourceClient struct {
	transport  clientruntime.Transport
	res        *clientruntime.Resource
	Attachment *IdentityPolicyAttachmentResourceClient
	create     *clientruntime.Operation
	list       *clientruntime.Operation
	retrieve   *clientruntime.Operation
	destroy    *clientruntime.Operation
	update     *clientruntime.Operation
}

func newIdentityPolicyResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*IdentityPolicyResourceClient, error) {
	res := finder.FindResource("IdentityPolicy")
	r := &IdentityPolicyResourceClient{
		transport: transport,
		res:       res,
	}
	var err error
	r.Attachment, err = newIdentityPolicyAttachmentResourceClient(transport, res)
	if err != nil {
		return nil, err
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
func (res *IdentityPolicyResourceClient) Create(ctx context.Context, input *IdentityPolicyCreateInput) (*IdentityPolicyCreateOutput, error) {
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
func (res *IdentityPolicyResourceClient) List(ctx context.Context, input *IdentityPolicyListInput) (*IdentityPolicyListOutput, error) {
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
func (res *IdentityPolicyResourceClient) Retrieve(ctx context.Context, input *IdentityPolicyRetrieveInput) (*IdentityPolicyRetrieveOutput, error) {
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
func (res *IdentityPolicyResourceClient) Destroy(ctx context.Context, input *IdentityPolicyDestroyInput) (*IdentityPolicyDestroyOutput, error) {
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
func (res *IdentityPolicyResourceClient) Update(ctx context.Context, input *IdentityPolicyUpdateInput) (*IdentityPolicyUpdateOutput, error) {
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

// IdentityPolicyAttachmentResourceClient is the IdentityPolicyAttachmentResourceClient resource client
type IdentityPolicyAttachmentResourceClient struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	list      *clientruntime.Operation
}

func newIdentityPolicyAttachmentResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*IdentityPolicyAttachmentResourceClient, error) {
	res := finder.FindResource("Attachment")
	r := &IdentityPolicyAttachmentResourceClient{
		transport: transport,
		res:       res,
	}
	r.list = res.FindOperation("List")
	return r, nil
}

// List - List attachments of an identity policy
// Requires permission action iam:ListIdentityPolicyAttachments over resource iam:IdentityPolicy(<name>)
func (res *IdentityPolicyAttachmentResourceClient) List(ctx context.Context, input *IdentityPolicyAttachmentListInput) (*IdentityPolicyAttachmentListOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.list,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*IdentityPolicyAttachmentListOutput)
	return output, nil
}

// TrustPolicyResourceClient is the TrustPolicyResourceClient resource client
type TrustPolicyResourceClient struct {
	transport  clientruntime.Transport
	res        *clientruntime.Resource
	Attachment *TrustPolicyAttachmentResourceClient
	create     *clientruntime.Operation
	retrieve   *clientruntime.Operation
	list       *clientruntime.Operation
	update     *clientruntime.Operation
	destroy    *clientruntime.Operation
}

func newTrustPolicyResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*TrustPolicyResourceClient, error) {
	res := finder.FindResource("TrustPolicy")
	r := &TrustPolicyResourceClient{
		transport: transport,
		res:       res,
	}
	var err error
	r.Attachment, err = newTrustPolicyAttachmentResourceClient(transport, res)
	if err != nil {
		return nil, err
	}
	r.create = res.FindOperation("Create")
	r.retrieve = res.FindOperation("Retrieve")
	r.list = res.FindOperation("List")
	r.update = res.FindOperation("Update")
	r.destroy = res.FindOperation("Destroy")
	return r, nil
}

// Create - Creates a reusable trust policy for a single principal.
// Requires permission action iam:CreateTrustPolicy over resource iam:TrustPolicy(<name>)
func (res *TrustPolicyResourceClient) Create(ctx context.Context, input *TrustPolicyCreateInput) (*TrustPolicyCreateOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.create,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*TrustPolicyCreateOutput)
	return output, nil
}

// Retrieve - Retrieves a trust policy by name.
func (res *TrustPolicyResourceClient) Retrieve(ctx context.Context, input *TrustPolicyRetrieveInput) (*TrustPolicyRetrieveOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.retrieve,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*TrustPolicyRetrieveOutput)
	return output, nil
}

// List - Lists the trust policies in the current account.
func (res *TrustPolicyResourceClient) List(ctx context.Context, input *TrustPolicyListInput) (*TrustPolicyListOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.list,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*TrustPolicyListOutput)
	return output, nil
}

// Update - Updates a trust policy's principal and statements.
// Requires permission action iam:UpdateTrustPolicy over resource iam:TrustPolicy(<name>)
func (res *TrustPolicyResourceClient) Update(ctx context.Context, input *TrustPolicyUpdateInput) (*TrustPolicyUpdateOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.update,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*TrustPolicyUpdateOutput)
	return output, nil
}

// Destroy - Destroys a trust policy. Any role attachments are removed.
// Requires permission action iam:DestroyTrustPolicy over resource iam:TrustPolicy(<name>)
func (res *TrustPolicyResourceClient) Destroy(ctx context.Context, input *TrustPolicyDestroyInput) (*TrustPolicyDestroyOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.destroy,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*TrustPolicyDestroyOutput)
	return output, nil
}

// TrustPolicyAttachmentResourceClient is the TrustPolicyAttachmentResourceClient resource client
type TrustPolicyAttachmentResourceClient struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	list      *clientruntime.Operation
}

func newTrustPolicyAttachmentResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*TrustPolicyAttachmentResourceClient, error) {
	res := finder.FindResource("Attachment")
	r := &TrustPolicyAttachmentResourceClient{
		transport: transport,
		res:       res,
	}
	r.list = res.FindOperation("List")
	return r, nil
}

// List - Lists the roles a trust policy is attached to.
// Requires permission action iam:ListTrustPolicies over resource iam:TrustPolicy(<name>)
func (res *TrustPolicyAttachmentResourceClient) List(ctx context.Context, input *TrustPolicyAttachmentListInput) (*TrustPolicyAttachmentListOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.list,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*TrustPolicyAttachmentListOutput)
	return output, nil
}

// InvitationResourceClient is the InvitationResourceClient resource client
type InvitationResourceClient struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	create    *clientruntime.Operation
	list      *clientruntime.Operation
	revoke    *clientruntime.Operation
	accept    *clientruntime.Operation
	inspect   *clientruntime.Operation
	decline   *clientruntime.Operation
}

func newInvitationResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*InvitationResourceClient, error) {
	res := finder.FindResource("Invitation")
	r := &InvitationResourceClient{
		transport: transport,
		res:       res,
	}
	r.create = res.FindOperation("Create")
	r.list = res.FindOperation("List")
	r.revoke = res.FindOperation("Revoke")
	r.accept = res.FindOperation("Accept")
	r.inspect = res.FindOperation("Inspect")
	r.decline = res.FindOperation("Decline")
	return r, nil
}

// Create - Invites a person by email to become a member of the caller's account.
// Requires permission action iam:CreateInvitation over resource iam:Invitation
func (res *InvitationResourceClient) Create(ctx context.Context, input *InvitationCreateInput) (*InvitationCreateOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.create,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*InvitationCreateOutput)
	return output, nil
}

// List - Lists pending invitations in the caller's account.
// Requires permission action iam:ListInvitations over resource iam:Invitation
func (res *InvitationResourceClient) List(ctx context.Context, input *InvitationListInput) (*InvitationListOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.list,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*InvitationListOutput)
	return output, nil
}

// Revoke - Revokes a pending invitation by id in the caller's account.
// Requires permission action iam:RevokeInvitation over resource iam:Invitation
func (res *InvitationResourceClient) Revoke(ctx context.Context, input *InvitationRevokeInput) (*InvitationRevokeOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.revoke,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*InvitationRevokeOutput)
	return output, nil
}

// Accept - Accepts an invitation. Called by the invitee after logging in via global
// SSO; the token comes from the emailed link. Provisions the invitee as a
// member of the inviting account and returns its name so the client can
// assume into it.
func (res *InvitationResourceClient) Accept(ctx context.Context, input *InvitationAcceptInput) (*InvitationAcceptOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.accept,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*InvitationAcceptOutput)
	return output, nil
}

// Inspect - Inspects an invitation by its token without accepting it, so the frontend
// can show the invitee what they've been invited to before they sign in and
// accept. Public (the token is the proof); rate-limited by client IP. Reports
// the status, so an already-used/revoked/expired invitation renders too.
func (res *InvitationResourceClient) Inspect(ctx context.Context, input *InvitationInspectInput) (*InvitationInspectOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.inspect,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*InvitationInspectOutput)
	return output, nil
}

// Decline - Declines an invitation by its token, closing it so it can no longer be
// accepted. Public (the token is the proof) and requires no sign-in. You
// shouldn't have to log into an account to reject it; rate-limited by client IP.
func (res *InvitationResourceClient) Decline(ctx context.Context, input *InvitationDeclineInput) (*InvitationDeclineOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.decline,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*InvitationDeclineOutput)
	return output, nil
}

// ServiceBearerTokenResourceClient is the ServiceBearerTokenResourceClient resource client
type ServiceBearerTokenResourceClient struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	get       *clientruntime.Operation
}

func newServiceBearerTokenResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*ServiceBearerTokenResourceClient, error) {
	res := finder.FindResource("ServiceBearerToken")
	r := &ServiceBearerTokenResourceClient{
		transport: transport,
		res:       res,
	}
	r.get = res.FindOperation("Get")
	return r, nil
}

// Get - Generates a temporary authorization token for accessing a a service or feature in a service that only supports bearer tokens
// The operations the token are limited to IAM permissions
// Requires permission to execute "iam:GetServiceBearerToken" over "*"
func (res *ServiceBearerTokenResourceClient) Get(ctx context.Context, input *ServiceBearerTokenGetInput) (*ServiceBearerTokenGetOutput, error) {
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

// SessionResourceClient is the SessionResourceClient resource client
type SessionResourceClient struct {
	transport       clientruntime.Transport
	res             *clientruntime.Resource
	keepAlive       *clientruntime.Operation
	revoke          *clientruntime.Operation
	identity        *clientruntime.Operation
	beginHandoff    *clientruntime.Operation
	completeHandoff *clientruntime.Operation
}

func newSessionResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*SessionResourceClient, error) {
	res := finder.FindResource("Session")
	r := &SessionResourceClient{
		transport: transport,
		res:       res,
	}
	r.keepAlive = res.FindOperation("KeepAlive")
	r.revoke = res.FindOperation("Revoke")
	r.identity = res.FindOperation("Identity")
	r.beginHandoff = res.FindOperation("BeginHandoff")
	r.completeHandoff = res.FindOperation("CompleteHandoff")
	return r, nil
}

// KeepAlive - Slides the calling credential's idle window forward (keep-alive heartbeat
// for idle periods; ordinary activity already slides it). Does not mint new
// credentials. The same access key stays valid, with a later expiry.
// Requires permission action iam:KeepAliveSession over resource iam:Session(<accessKeyID>)
func (res *SessionResourceClient) KeepAlive(ctx context.Context, input *SessionKeepAliveInput) (*SessionKeepAliveOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.keepAlive,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*SessionKeepAliveOutput)
	return output, nil
}

// Revoke - Revokes (deletes) the calling credential: self logout. The next request
// signed with this credential is rejected. Works for role and user credentials.
// Requires permission action iam:RevokeSession over resource iam:Session(<accessKeyID>)
func (res *SessionResourceClient) Revoke(ctx context.Context, input *SessionRevokeInput) (*SessionRevokeOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.revoke,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*SessionRevokeOutput)
	return output, nil
}

// Identity - Returns who the calling credential is: the account it is scoped to, its access
// key id, and the DRN of the calling principal. Works for user, role, and
// service (service-linked role) credentials; requires only that the request is
// signed. No extra permission is required.
func (res *SessionResourceClient) Identity(ctx context.Context, input *SessionIdentityInput) (*SessionIdentityOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.identity,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*SessionIdentityOutput)
	return output, nil
}

// BeginHandoff - Begins a handoff of this session to another device, for the SAME identity, and
// returns a single-use short-lived code. No credentials are minted until the code
// is redeemed.
//
// The name says handoff. The model is TWO SESSIONS, and nothing moves. This session
// keeps its own credential, its own address and its own audit trail, and the
// redeeming device gets its own of each. A session seen from two addresses cannot
// be told apart from a stolen one, so we never create one.
//
// Revoking this session revokes the session the code creates. This session merely
// expiring does not: expiry means this device went unused, which says nothing about
// the other one. The new session gets its own deadline at redemption time and keeps
// itself alive from then on.
//
// There is no accountName input. The redeeming device continues this session, in
// this account, as this principal, so there is no target to name.
// Requires permission action iam:BeginHandoffSession over resource iam:Session(<accessKeyID>)
func (res *SessionResourceClient) BeginHandoff(ctx context.Context, input *SessionBeginHandoffInput) (*SessionBeginHandoffOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.beginHandoff,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*SessionBeginHandoffOutput)
	return output, nil
}

// CompleteHandoff - Public: redeems a code minted by BeginHandoff and returns the credentials of a new
// session for the same identity. No signed request, because the redeeming device
// holds no credentials yet. The code is the only proof and it is consumed exactly
// once, so a failed redemption cannot be retried.
//
// accountName and principalDrn in the output are the authoritative identity of the
// new session. A client that carried an identity value alongside the code in a URL
// must treat that value as a display hint only, never as authentication.
func (res *SessionResourceClient) CompleteHandoff(ctx context.Context, input *SessionCompleteHandoffInput) (*SessionCompleteHandoffOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.completeHandoff,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*SessionCompleteHandoffOutput)
	return output, nil
}

// ServiceCatalogResourceClient is the ServiceCatalogResourceClient resource client
type ServiceCatalogResourceClient struct {
	transport clientruntime.Transport
	res       *clientruntime.Resource
	list      *clientruntime.Operation
	get       *clientruntime.Operation
}

func newServiceCatalogResourceClient(
	transport clientruntime.Transport,
	finder clientruntime.ResourceFinder,
) (*ServiceCatalogResourceClient, error) {
	res := finder.FindResource("ServiceCatalog")
	r := &ServiceCatalogResourceClient{
		transport: transport,
		res:       res,
	}
	r.list = res.FindOperation("List")
	r.get = res.FindOperation("Get")
	return r, nil
}

// List - Lists every published service catalog (the namespace picker).
func (res *ServiceCatalogResourceClient) List(ctx context.Context, input *ServiceCatalogListInput) (*ServiceCatalogListOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.list,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*ServiceCatalogListOutput)
	return output, nil
}

// Get - Returns the full catalog for a namespace.
func (res *ServiceCatalogResourceClient) Get(ctx context.Context, input *ServiceCatalogGetInput) (*ServiceCatalogGetOutput, error) {
	o, err := res.transport.Execute(ctx, &clientruntime.Request{
		Operation: res.get,
		Input:     input,
	})
	if err != nil {
		return nil, err
	}
	output := o.(*ServiceCatalogGetOutput)
	return output, nil
}

// Client is the main client of the API
type Client struct {
	transport clientruntime.Transport
	pk        *clientruntime.Package
	// Account operations
	Account *AccountResourceClient
	// Region operations
	Region *RegionResourceClient
	// User operations
	User *UserResourceClient
	// Role operations
	Role *RoleResourceClient
	// Group operations
	Group *GroupResourceClient
	// IdentityPolicy operations
	IdentityPolicy *IdentityPolicyResourceClient
	// TrustPolicy operations
	TrustPolicy *TrustPolicyResourceClient
	// Invitation - Invitations: invite people (by email) to become members of an account, and
	// accept an invitation once logged in. Create/List/Revoke act on the caller's own
	// account (resolved from the signed credential); Accept is called by the invitee
	// after logging in via global SSO.
	Invitation *InvitationResourceClient
	// ServiceBearerToken - Service Bearer Tokens
	ServiceBearerToken *ServiceBearerTokenResourceClient
	// Session - Self-management of the calling credential (the signing access key). The signed
	// operations here act on the caller's own credential, identified from the signed
	// request context (never from input), so a credential can only manage itself.
	// These are ordinary permissions, granted by the builtin policies
	// session-keepalive, session-revoke and session-handoff, and attached when an
	// identity is assumed.
	Session *SessionResourceClient
	// ServiceCatalog - ServiceCatalog is public/anonymous reference data: the full set of actions and
	// resource types each service supports, used to power identity-policy builders.
	// No signed request and no permission required.
	ServiceCatalog *ServiceCatalogResourceClient
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
		"https://iam.<region>.api.deployport.io/api",
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
	c.Account, err = newAccountResourceClient(transport, pk)
	if err != nil {
		return nil, err
	}
	c.Region, err = newRegionResourceClient(transport, pk)
	if err != nil {
		return nil, err
	}
	c.User, err = newUserResourceClient(transport, pk)
	if err != nil {
		return nil, err
	}
	c.Role, err = newRoleResourceClient(transport, pk)
	if err != nil {
		return nil, err
	}
	c.Group, err = newGroupResourceClient(transport, pk)
	if err != nil {
		return nil, err
	}
	c.IdentityPolicy, err = newIdentityPolicyResourceClient(transport, pk)
	if err != nil {
		return nil, err
	}
	c.TrustPolicy, err = newTrustPolicyResourceClient(transport, pk)
	if err != nil {
		return nil, err
	}
	c.Invitation, err = newInvitationResourceClient(transport, pk)
	if err != nil {
		return nil, err
	}
	c.ServiceBearerToken, err = newServiceBearerTokenResourceClient(transport, pk)
	if err != nil {
		return nil, err
	}
	c.Session, err = newSessionResourceClient(transport, pk)
	if err != nil {
		return nil, err
	}
	c.ServiceCatalog, err = newServiceCatalogResourceClient(transport, pk)
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
	mod                                                                   *clientruntime.Package
	structPathServiceCatalogAction                                        *clientruntime.StructDefinition
	structPathServiceCatalogResourceType                                  *clientruntime.StructDefinition
	structPathServiceCatalogData                                          *clientruntime.StructDefinition
	structPathServiceCatalogInfo                                          *clientruntime.StructDefinition
	structPathServiceCatalogSummary                                       *clientruntime.StructDefinition
	structPathUserInformationSSOProvider                                  *clientruntime.StructDefinition
	structPathUserInformationSSOProfile                                   *clientruntime.StructDefinition
	structPathUserInformationSSO                                          *clientruntime.StructDefinition
	structPathManagedByService                                            *clientruntime.StructDefinition
	structPathUserInformation                                             *clientruntime.StructDefinition
	structPathRoleInformation                                             *clientruntime.StructDefinition
	structPathGroupInformation                                            *clientruntime.StructDefinition
	structPathCredentials                                                 *clientruntime.StructDefinition
	structPathSSOProviderUnavailableError                                 *clientruntime.StructDefinition
	structPathSSOFlow                                                     *clientruntime.StructDefinition
	structPathAccount                                                     *clientruntime.StructDefinition
	structPathRegionEndpoint                                              *clientruntime.StructDefinition
	structPathRegionInfo                                                  *clientruntime.StructDefinition
	structPathAccountSSOProvider                                          *clientruntime.StructDefinition
	structPathPolicyNotFoundError                                         *clientruntime.StructDefinition
	structPathInvitation                                                  *clientruntime.StructDefinition
	structPathInvalidInvitationError                                      *clientruntime.StructDefinition
	structPathInvitationNotFoundError                                     *clientruntime.StructDefinition
	structPathInvitationPreview                                           *clientruntime.StructDefinition
	structPathOIDCProvider                                                *clientruntime.StructDefinition
	structPathInvalidOIDCProviderError                                    *clientruntime.StructDefinition
	structPathInvalidOIDCIssuerError                                      *clientruntime.StructDefinition
	structPathOIDCProviderNotFoundError                                   *clientruntime.StructDefinition
	structPathOIDCProviderInUseError                                      *clientruntime.StructDefinition
	structPathTrustPolicyStatement                                        *clientruntime.StructDefinition
	structPathTrustPolicy                                                 *clientruntime.StructDefinition
	structPathTrustPolicyAttachment                                       *clientruntime.StructDefinition
	structPathInvalidTrustPolicyError                                     *clientruntime.StructDefinition
	structPathTrustPolicyNotFoundError                                    *clientruntime.StructDefinition
	structPathInvalidWebIdentityTokenError                                *clientruntime.StructDefinition
	structPathInvalidPrincipalDRNError                                    *clientruntime.StructDefinition
	structPathInvalidHandoffCodeError                                     *clientruntime.StructDefinition
	structPathAccountCreateInput                                          *clientruntime.StructDefinition
	structPathAccountCreateOutput                                         *clientruntime.StructDefinition
	structPathAccountCreateInvalidNameError                               *clientruntime.StructDefinition
	structPathAccountAssumeIdentityInput                                  *clientruntime.StructDefinition
	structPathAccountAssumeIdentityOutput                                 *clientruntime.StructDefinition
	structPathAccountBeginAssumeIdentityInput                             *clientruntime.StructDefinition
	structPathAccountBeginAssumeIdentityOutput                            *clientruntime.StructDefinition
	structPathAccountCompleteAssumeIdentityInput                          *clientruntime.StructDefinition
	structPathAccountCompleteAssumeIdentityOutput                         *clientruntime.StructDefinition
	structPathAccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError *clientruntime.StructDefinition
	structPathAccountSSOBeginAuthenticationInput                          *clientruntime.StructDefinition
	structPathAccountSSOBeginAuthenticationOutput                         *clientruntime.StructDefinition
	structPathAccountSSOBeginAuthenticationParameterError                 *clientruntime.StructDefinition
	structPathAccountSSOCompleteAuthenticationInput                       *clientruntime.StructDefinition
	structPathAccountSSOCompleteAuthenticationOutput                      *clientruntime.StructDefinition
	structPathAccountSSOCompleteAuthenticationInvalidFlowError            *clientruntime.StructDefinition
	structPathAccountSSOGetProvidersInput                                 *clientruntime.StructDefinition
	structPathAccountSSOGetProvidersOutput                                *clientruntime.StructDefinition
	structPathAccountOIDCProviderCreateInput                              *clientruntime.StructDefinition
	structPathAccountOIDCProviderCreateOutput                             *clientruntime.StructDefinition
	structPathAccountOIDCProviderSetAudiencesInput                        *clientruntime.StructDefinition
	structPathAccountOIDCProviderSetAudiencesOutput                       *clientruntime.StructDefinition
	structPathAccountOIDCProviderListInput                                *clientruntime.StructDefinition
	structPathAccountOIDCProviderListOutput                               *clientruntime.StructDefinition
	structPathAccountOIDCProviderDeleteInput                              *clientruntime.StructDefinition
	structPathAccountOIDCProviderDeleteOutput                             *clientruntime.StructDefinition
	structPathAccountOIDCProviderTrustPolicyListInput                     *clientruntime.StructDefinition
	structPathAccountOIDCProviderTrustPolicyListOutput                    *clientruntime.StructDefinition
	structPathRegionListInput                                             *clientruntime.StructDefinition
	structPathRegionListOutput                                            *clientruntime.StructDefinition
	structPathInvalidUsernameError                                        *clientruntime.StructDefinition
	structPathInvalidRoleNameError                                        *clientruntime.StructDefinition
	structPathMemberAccount                                               *clientruntime.StructDefinition
	structPathCredentialInfo                                              *clientruntime.StructDefinition
	structPathIdentityPolicyStatement                                     *clientruntime.StructDefinition
	structPathIdentityPolicy                                              *clientruntime.StructDefinition
	structPathIdentityPolicyAttachment                                    *clientruntime.StructDefinition
	structPathIdentityPolicyAttachmentInfo                                *clientruntime.StructDefinition
	structPathUserNotFoundError                                           *clientruntime.StructDefinition
	structPathCannotDisableSelfError                                      *clientruntime.StructDefinition
	structPathRoleNotFoundError                                           *clientruntime.StructDefinition
	structPathGroupNotFoundError                                          *clientruntime.StructDefinition
	structPathInvalidGroupNameError                                       *clientruntime.StructDefinition
	structPathInvalidGroupFilterError                                     *clientruntime.StructDefinition
	structPathInvalidUserFilterError                                      *clientruntime.StructDefinition
	structPathIdentityInUseError                                          *clientruntime.StructDefinition
	structPathUserSSHKey                                                  *clientruntime.StructDefinition
	structPathInvalidSSHPublicKeyError                                    *clientruntime.StructDefinition
	structPathSSHKeyAlreadyExistsError                                    *clientruntime.StructDefinition
	structPathInvalidSSHKeyTitleError                                     *clientruntime.StructDefinition
	structPathUserCreateInput                                             *clientruntime.StructDefinition
	structPathUserCreateOutput                                            *clientruntime.StructDefinition
	structPathUserGetInput                                                *clientruntime.StructDefinition
	structPathUserGetOutput                                               *clientruntime.StructDefinition
	structPathUserGetUserNotAvailableError                                *clientruntime.StructDefinition
	structPathUserDestroyInput                                            *clientruntime.StructDefinition
	structPathUserDestroyOutput                                           *clientruntime.StructDefinition
	structPathUserSetActiveInput                                          *clientruntime.StructDefinition
	structPathUserSetActiveOutput                                         *clientruntime.StructDefinition
	structPathUserListInput                                               *clientruntime.StructDefinition
	structPathUserListOutput                                              *clientruntime.StructDefinition
	structPathUserMemberAccountsInput                                     *clientruntime.StructDefinition
	structPathUserMemberAccountsOutput                                    *clientruntime.StructDefinition
	structPathUserAccessKeyCreateInput                                    *clientruntime.StructDefinition
	structPathUserAccessKeyCreateOutput                                   *clientruntime.StructDefinition
	structPathUserAccessKeyListInput                                      *clientruntime.StructDefinition
	structPathUserAccessKeyListOutput                                     *clientruntime.StructDefinition
	structPathUserAccessKeyDestroyInput                                   *clientruntime.StructDefinition
	structPathUserAccessKeyDestroyOutput                                  *clientruntime.StructDefinition
	structPathUserIdentityPolicyAttachInput                               *clientruntime.StructDefinition
	structPathUserIdentityPolicyAttachOutput                              *clientruntime.StructDefinition
	structPathUserIdentityPolicyListInput                                 *clientruntime.StructDefinition
	structPathUserIdentityPolicyListOutput                                *clientruntime.StructDefinition
	structPathUserIdentityPolicyDetachInput                               *clientruntime.StructDefinition
	structPathUserIdentityPolicyDetachOutput                              *clientruntime.StructDefinition
	structPathUserSSHKeyCreateInput                                       *clientruntime.StructDefinition
	structPathUserSSHKeyCreateOutput                                      *clientruntime.StructDefinition
	structPathUserSSHKeyListInput                                         *clientruntime.StructDefinition
	structPathUserSSHKeyListOutput                                        *clientruntime.StructDefinition
	structPathUserSSHKeyDestroyInput                                      *clientruntime.StructDefinition
	structPathUserSSHKeyDestroyOutput                                     *clientruntime.StructDefinition
	structPathUserSSHKeySetTitleInput                                     *clientruntime.StructDefinition
	structPathUserSSHKeySetTitleOutput                                    *clientruntime.StructDefinition
	structPathUserGroupListInput                                          *clientruntime.StructDefinition
	structPathUserGroupListOutput                                         *clientruntime.StructDefinition
	structPathInlinePolicy                                                *clientruntime.StructDefinition
	structPathPolicyStructureError                                        *clientruntime.StructDefinition
	structPathRoleCreateInput                                             *clientruntime.StructDefinition
	structPathRoleCreateOutput                                            *clientruntime.StructDefinition
	structPathRoleGetInput                                                *clientruntime.StructDefinition
	structPathRoleGetOutput                                               *clientruntime.StructDefinition
	structPathRoleDestroyInput                                            *clientruntime.StructDefinition
	structPathRoleDestroyOutput                                           *clientruntime.StructDefinition
	structPathRoleListInput                                               *clientruntime.StructDefinition
	structPathRoleListOutput                                              *clientruntime.StructDefinition
	structPathRoleAssumeInput                                             *clientruntime.StructDefinition
	structPathRoleAssumeOutput                                            *clientruntime.StructDefinition
	structPathRoleAssumeWithWebIdentityInput                              *clientruntime.StructDefinition
	structPathRoleAssumeWithWebIdentityOutput                             *clientruntime.StructDefinition
	structPathRoleAccessKeyListInput                                      *clientruntime.StructDefinition
	structPathRoleAccessKeyListOutput                                     *clientruntime.StructDefinition
	structPathRoleAccessKeyDestroyInput                                   *clientruntime.StructDefinition
	structPathRoleAccessKeyDestroyOutput                                  *clientruntime.StructDefinition
	structPathRoleIdentityPolicyAttachInput                               *clientruntime.StructDefinition
	structPathRoleIdentityPolicyAttachOutput                              *clientruntime.StructDefinition
	structPathRoleIdentityPolicyListInput                                 *clientruntime.StructDefinition
	structPathRoleIdentityPolicyListOutput                                *clientruntime.StructDefinition
	structPathRoleIdentityPolicyDetachInput                               *clientruntime.StructDefinition
	structPathRoleIdentityPolicyDetachOutput                              *clientruntime.StructDefinition
	structPathRoleTrustPolicyAttachInput                                  *clientruntime.StructDefinition
	structPathRoleTrustPolicyAttachOutput                                 *clientruntime.StructDefinition
	structPathRoleTrustPolicyListInput                                    *clientruntime.StructDefinition
	structPathRoleTrustPolicyListOutput                                   *clientruntime.StructDefinition
	structPathRoleTrustPolicyDetachInput                                  *clientruntime.StructDefinition
	structPathRoleTrustPolicyDetachOutput                                 *clientruntime.StructDefinition
	structPathGroupCreateInput                                            *clientruntime.StructDefinition
	structPathGroupCreateOutput                                           *clientruntime.StructDefinition
	structPathGroupDestroyInput                                           *clientruntime.StructDefinition
	structPathGroupDestroyOutput                                          *clientruntime.StructDefinition
	structPathGroupGetInput                                               *clientruntime.StructDefinition
	structPathGroupGetOutput                                              *clientruntime.StructDefinition
	structPathGroupListInput                                              *clientruntime.StructDefinition
	structPathGroupListOutput                                             *clientruntime.StructDefinition
	structPathGroupMemberAddInput                                         *clientruntime.StructDefinition
	structPathGroupMemberAddOutput                                        *clientruntime.StructDefinition
	structPathGroupMemberRemoveInput                                      *clientruntime.StructDefinition
	structPathGroupMemberRemoveOutput                                     *clientruntime.StructDefinition
	structPathGroupMemberListInput                                        *clientruntime.StructDefinition
	structPathGroupMemberListOutput                                       *clientruntime.StructDefinition
	structPathGroupIdentityPolicyAttachInput                              *clientruntime.StructDefinition
	structPathGroupIdentityPolicyAttachOutput                             *clientruntime.StructDefinition
	structPathGroupIdentityPolicyDetachInput                              *clientruntime.StructDefinition
	structPathGroupIdentityPolicyDetachOutput                             *clientruntime.StructDefinition
	structPathGroupIdentityPolicyListInput                                *clientruntime.StructDefinition
	structPathGroupIdentityPolicyListOutput                               *clientruntime.StructDefinition
	structPathIdentityPolicyCreateInput                                   *clientruntime.StructDefinition
	structPathIdentityPolicyCreateOutput                                  *clientruntime.StructDefinition
	structPathIdentityPolicyListInput                                     *clientruntime.StructDefinition
	structPathIdentityPolicyListOutput                                    *clientruntime.StructDefinition
	structPathIdentityPolicyRetrieveInput                                 *clientruntime.StructDefinition
	structPathIdentityPolicyRetrieveOutput                                *clientruntime.StructDefinition
	structPathIdentityPolicyDestroyInput                                  *clientruntime.StructDefinition
	structPathIdentityPolicyDestroyOutput                                 *clientruntime.StructDefinition
	structPathIdentityPolicyUpdateInput                                   *clientruntime.StructDefinition
	structPathIdentityPolicyUpdateOutput                                  *clientruntime.StructDefinition
	structPathIdentityPolicyAttachmentListInput                           *clientruntime.StructDefinition
	structPathIdentityPolicyAttachmentListOutput                          *clientruntime.StructDefinition
	structPathTrustPolicyCreateInput                                      *clientruntime.StructDefinition
	structPathTrustPolicyCreateOutput                                     *clientruntime.StructDefinition
	structPathTrustPolicyRetrieveInput                                    *clientruntime.StructDefinition
	structPathTrustPolicyRetrieveOutput                                   *clientruntime.StructDefinition
	structPathTrustPolicyListInput                                        *clientruntime.StructDefinition
	structPathTrustPolicyListOutput                                       *clientruntime.StructDefinition
	structPathTrustPolicyUpdateInput                                      *clientruntime.StructDefinition
	structPathTrustPolicyUpdateOutput                                     *clientruntime.StructDefinition
	structPathTrustPolicyDestroyInput                                     *clientruntime.StructDefinition
	structPathTrustPolicyDestroyOutput                                    *clientruntime.StructDefinition
	structPathTrustPolicyAttachmentListInput                              *clientruntime.StructDefinition
	structPathTrustPolicyAttachmentListOutput                             *clientruntime.StructDefinition
	structPathServiceBearerToken                                          *clientruntime.StructDefinition
	structPathInvalidServiceBearerTokenDurationError                      *clientruntime.StructDefinition
	structPathInvalidServiceNameError                                     *clientruntime.StructDefinition
	structPathInvitationCreateInput                                       *clientruntime.StructDefinition
	structPathInvitationCreateOutput                                      *clientruntime.StructDefinition
	structPathInvitationListInput                                         *clientruntime.StructDefinition
	structPathInvitationListOutput                                        *clientruntime.StructDefinition
	structPathInvitationRevokeInput                                       *clientruntime.StructDefinition
	structPathInvitationRevokeOutput                                      *clientruntime.StructDefinition
	structPathInvitationAcceptInput                                       *clientruntime.StructDefinition
	structPathInvitationAcceptOutput                                      *clientruntime.StructDefinition
	structPathInvitationInspectInput                                      *clientruntime.StructDefinition
	structPathInvitationInspectOutput                                     *clientruntime.StructDefinition
	structPathInvitationDeclineInput                                      *clientruntime.StructDefinition
	structPathInvitationDeclineOutput                                     *clientruntime.StructDefinition
	structPathServiceBearerTokenGetInput                                  *clientruntime.StructDefinition
	structPathServiceBearerTokenGetOutput                                 *clientruntime.StructDefinition
	structPathSessionKeepAliveInput                                       *clientruntime.StructDefinition
	structPathSessionKeepAliveOutput                                      *clientruntime.StructDefinition
	structPathSessionRevokeInput                                          *clientruntime.StructDefinition
	structPathSessionRevokeOutput                                         *clientruntime.StructDefinition
	structPathSessionIdentityInput                                        *clientruntime.StructDefinition
	structPathSessionIdentityOutput                                       *clientruntime.StructDefinition
	structPathSessionBeginHandoffInput                                    *clientruntime.StructDefinition
	structPathSessionBeginHandoffOutput                                   *clientruntime.StructDefinition
	structPathSessionCompleteHandoffInput                                 *clientruntime.StructDefinition
	structPathSessionCompleteHandoffOutput                                *clientruntime.StructDefinition
	structPathServiceCatalogListInput                                     *clientruntime.StructDefinition
	structPathServiceCatalogListOutput                                    *clientruntime.StructDefinition
	structPathServiceCatalogGetInput                                      *clientruntime.StructDefinition
	structPathServiceCatalogGetOutput                                     *clientruntime.StructDefinition
	structPathServiceCatalogGetServiceCatalogNotFoundError                *clientruntime.StructDefinition
}

// Module returns the module definition
func (m *SpecularMetaInfo) Module() *clientruntime.Package {
	return m.mod
}

// ServiceCatalogActionStruct allows easy access to structure
func (m *SpecularMetaInfo) ServiceCatalogActionStruct() *clientruntime.StructDefinition {
	return m.structPathServiceCatalogAction
}

// ServiceCatalogResourceTypeStruct allows easy access to structure
func (m *SpecularMetaInfo) ServiceCatalogResourceTypeStruct() *clientruntime.StructDefinition {
	return m.structPathServiceCatalogResourceType
}

// ServiceCatalogDataStruct allows easy access to structure
func (m *SpecularMetaInfo) ServiceCatalogDataStruct() *clientruntime.StructDefinition {
	return m.structPathServiceCatalogData
}

// ServiceCatalogInfoStruct allows easy access to structure
func (m *SpecularMetaInfo) ServiceCatalogInfoStruct() *clientruntime.StructDefinition {
	return m.structPathServiceCatalogInfo
}

// ServiceCatalogSummaryStruct allows easy access to structure
func (m *SpecularMetaInfo) ServiceCatalogSummaryStruct() *clientruntime.StructDefinition {
	return m.structPathServiceCatalogSummary
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

// ManagedByServiceStruct allows easy access to structure
func (m *SpecularMetaInfo) ManagedByServiceStruct() *clientruntime.StructDefinition {
	return m.structPathManagedByService
}

// UserInformationStruct allows easy access to structure
func (m *SpecularMetaInfo) UserInformationStruct() *clientruntime.StructDefinition {
	return m.structPathUserInformation
}

// RoleInformationStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleInformationStruct() *clientruntime.StructDefinition {
	return m.structPathRoleInformation
}

// GroupInformationStruct allows easy access to structure
func (m *SpecularMetaInfo) GroupInformationStruct() *clientruntime.StructDefinition {
	return m.structPathGroupInformation
}

// CredentialsStruct allows easy access to structure
func (m *SpecularMetaInfo) CredentialsStruct() *clientruntime.StructDefinition {
	return m.structPathCredentials
}

// SSOProviderUnavailableErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) SSOProviderUnavailableErrorStruct() *clientruntime.StructDefinition {
	return m.structPathSSOProviderUnavailableError
}

// SSOFlowStruct allows easy access to structure
func (m *SpecularMetaInfo) SSOFlowStruct() *clientruntime.StructDefinition {
	return m.structPathSSOFlow
}

// AccountStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountStruct() *clientruntime.StructDefinition {
	return m.structPathAccount
}

// RegionEndpointStruct allows easy access to structure
func (m *SpecularMetaInfo) RegionEndpointStruct() *clientruntime.StructDefinition {
	return m.structPathRegionEndpoint
}

// RegionInfoStruct allows easy access to structure
func (m *SpecularMetaInfo) RegionInfoStruct() *clientruntime.StructDefinition {
	return m.structPathRegionInfo
}

// AccountSSOProviderStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOProviderStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOProvider
}

// PolicyNotFoundErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) PolicyNotFoundErrorStruct() *clientruntime.StructDefinition {
	return m.structPathPolicyNotFoundError
}

// InvitationStruct allows easy access to structure
func (m *SpecularMetaInfo) InvitationStruct() *clientruntime.StructDefinition {
	return m.structPathInvitation
}

// InvalidInvitationErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) InvalidInvitationErrorStruct() *clientruntime.StructDefinition {
	return m.structPathInvalidInvitationError
}

// InvitationNotFoundErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) InvitationNotFoundErrorStruct() *clientruntime.StructDefinition {
	return m.structPathInvitationNotFoundError
}

// InvitationPreviewStruct allows easy access to structure
func (m *SpecularMetaInfo) InvitationPreviewStruct() *clientruntime.StructDefinition {
	return m.structPathInvitationPreview
}

// OIDCProviderStruct allows easy access to structure
func (m *SpecularMetaInfo) OIDCProviderStruct() *clientruntime.StructDefinition {
	return m.structPathOIDCProvider
}

// InvalidOIDCProviderErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) InvalidOIDCProviderErrorStruct() *clientruntime.StructDefinition {
	return m.structPathInvalidOIDCProviderError
}

// InvalidOIDCIssuerErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) InvalidOIDCIssuerErrorStruct() *clientruntime.StructDefinition {
	return m.structPathInvalidOIDCIssuerError
}

// OIDCProviderNotFoundErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) OIDCProviderNotFoundErrorStruct() *clientruntime.StructDefinition {
	return m.structPathOIDCProviderNotFoundError
}

// OIDCProviderInUseErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) OIDCProviderInUseErrorStruct() *clientruntime.StructDefinition {
	return m.structPathOIDCProviderInUseError
}

// TrustPolicyStatementStruct allows easy access to structure
func (m *SpecularMetaInfo) TrustPolicyStatementStruct() *clientruntime.StructDefinition {
	return m.structPathTrustPolicyStatement
}

// TrustPolicyStruct allows easy access to structure
func (m *SpecularMetaInfo) TrustPolicyStruct() *clientruntime.StructDefinition {
	return m.structPathTrustPolicy
}

// TrustPolicyAttachmentStruct allows easy access to structure
func (m *SpecularMetaInfo) TrustPolicyAttachmentStruct() *clientruntime.StructDefinition {
	return m.structPathTrustPolicyAttachment
}

// InvalidTrustPolicyErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) InvalidTrustPolicyErrorStruct() *clientruntime.StructDefinition {
	return m.structPathInvalidTrustPolicyError
}

// TrustPolicyNotFoundErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) TrustPolicyNotFoundErrorStruct() *clientruntime.StructDefinition {
	return m.structPathTrustPolicyNotFoundError
}

// InvalidWebIdentityTokenErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) InvalidWebIdentityTokenErrorStruct() *clientruntime.StructDefinition {
	return m.structPathInvalidWebIdentityTokenError
}

// InvalidPrincipalDRNErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) InvalidPrincipalDRNErrorStruct() *clientruntime.StructDefinition {
	return m.structPathInvalidPrincipalDRNError
}

// InvalidHandoffCodeErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) InvalidHandoffCodeErrorStruct() *clientruntime.StructDefinition {
	return m.structPathInvalidHandoffCodeError
}

// AccountCreateInputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountCreateInputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountCreateInput
}

// AccountCreateOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountCreateOutputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountCreateOutput
}

// AccountCreateInvalidNameErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountCreateInvalidNameErrorStruct() *clientruntime.StructDefinition {
	return m.structPathAccountCreateInvalidNameError
}

// AccountAssumeIdentityInputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountAssumeIdentityInputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountAssumeIdentityInput
}

// AccountAssumeIdentityOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountAssumeIdentityOutputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountAssumeIdentityOutput
}

// AccountBeginAssumeIdentityInputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountBeginAssumeIdentityInputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountBeginAssumeIdentityInput
}

// AccountBeginAssumeIdentityOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountBeginAssumeIdentityOutputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountBeginAssumeIdentityOutput
}

// AccountCompleteAssumeIdentityInputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountCompleteAssumeIdentityInputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountCompleteAssumeIdentityInput
}

// AccountCompleteAssumeIdentityOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountCompleteAssumeIdentityOutputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountCompleteAssumeIdentityOutput
}

// AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountCompleteAssumeIdentityInvalidAssumeIdentityCodeErrorStruct() *clientruntime.StructDefinition {
	return m.structPathAccountCompleteAssumeIdentityInvalidAssumeIdentityCodeError
}

// AccountSSOBeginAuthenticationInputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOBeginAuthenticationInputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOBeginAuthenticationInput
}

// AccountSSOBeginAuthenticationOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOBeginAuthenticationOutputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOBeginAuthenticationOutput
}

// AccountSSOBeginAuthenticationParameterErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOBeginAuthenticationParameterErrorStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOBeginAuthenticationParameterError
}

// AccountSSOCompleteAuthenticationInputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOCompleteAuthenticationInputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOCompleteAuthenticationInput
}

// AccountSSOCompleteAuthenticationOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOCompleteAuthenticationOutputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOCompleteAuthenticationOutput
}

// AccountSSOCompleteAuthenticationInvalidFlowErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOCompleteAuthenticationInvalidFlowErrorStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOCompleteAuthenticationInvalidFlowError
}

// AccountSSOGetProvidersInputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOGetProvidersInputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOGetProvidersInput
}

// AccountSSOGetProvidersOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountSSOGetProvidersOutputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountSSOGetProvidersOutput
}

// AccountOIDCProviderCreateInputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountOIDCProviderCreateInputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountOIDCProviderCreateInput
}

// AccountOIDCProviderCreateOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountOIDCProviderCreateOutputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountOIDCProviderCreateOutput
}

// AccountOIDCProviderSetAudiencesInputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountOIDCProviderSetAudiencesInputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountOIDCProviderSetAudiencesInput
}

// AccountOIDCProviderSetAudiencesOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountOIDCProviderSetAudiencesOutputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountOIDCProviderSetAudiencesOutput
}

// AccountOIDCProviderListInputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountOIDCProviderListInputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountOIDCProviderListInput
}

// AccountOIDCProviderListOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountOIDCProviderListOutputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountOIDCProviderListOutput
}

// AccountOIDCProviderDeleteInputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountOIDCProviderDeleteInputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountOIDCProviderDeleteInput
}

// AccountOIDCProviderDeleteOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountOIDCProviderDeleteOutputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountOIDCProviderDeleteOutput
}

// AccountOIDCProviderTrustPolicyListInputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountOIDCProviderTrustPolicyListInputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountOIDCProviderTrustPolicyListInput
}

// AccountOIDCProviderTrustPolicyListOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) AccountOIDCProviderTrustPolicyListOutputStruct() *clientruntime.StructDefinition {
	return m.structPathAccountOIDCProviderTrustPolicyListOutput
}

// RegionListInputStruct allows easy access to structure
func (m *SpecularMetaInfo) RegionListInputStruct() *clientruntime.StructDefinition {
	return m.structPathRegionListInput
}

// RegionListOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) RegionListOutputStruct() *clientruntime.StructDefinition {
	return m.structPathRegionListOutput
}

// InvalidUsernameErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) InvalidUsernameErrorStruct() *clientruntime.StructDefinition {
	return m.structPathInvalidUsernameError
}

// InvalidRoleNameErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) InvalidRoleNameErrorStruct() *clientruntime.StructDefinition {
	return m.structPathInvalidRoleNameError
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

// IdentityPolicyAttachmentInfoStruct allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyAttachmentInfoStruct() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyAttachmentInfo
}

// UserNotFoundErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) UserNotFoundErrorStruct() *clientruntime.StructDefinition {
	return m.structPathUserNotFoundError
}

// CannotDisableSelfErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) CannotDisableSelfErrorStruct() *clientruntime.StructDefinition {
	return m.structPathCannotDisableSelfError
}

// RoleNotFoundErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleNotFoundErrorStruct() *clientruntime.StructDefinition {
	return m.structPathRoleNotFoundError
}

// GroupNotFoundErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) GroupNotFoundErrorStruct() *clientruntime.StructDefinition {
	return m.structPathGroupNotFoundError
}

// InvalidGroupNameErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) InvalidGroupNameErrorStruct() *clientruntime.StructDefinition {
	return m.structPathInvalidGroupNameError
}

// InvalidGroupFilterErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) InvalidGroupFilterErrorStruct() *clientruntime.StructDefinition {
	return m.structPathInvalidGroupFilterError
}

// InvalidUserFilterErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) InvalidUserFilterErrorStruct() *clientruntime.StructDefinition {
	return m.structPathInvalidUserFilterError
}

// IdentityInUseErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) IdentityInUseErrorStruct() *clientruntime.StructDefinition {
	return m.structPathIdentityInUseError
}

// UserSSHKeyStruct allows easy access to structure
func (m *SpecularMetaInfo) UserSSHKeyStruct() *clientruntime.StructDefinition {
	return m.structPathUserSSHKey
}

// InvalidSSHPublicKeyErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) InvalidSSHPublicKeyErrorStruct() *clientruntime.StructDefinition {
	return m.structPathInvalidSSHPublicKeyError
}

// SSHKeyAlreadyExistsErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) SSHKeyAlreadyExistsErrorStruct() *clientruntime.StructDefinition {
	return m.structPathSSHKeyAlreadyExistsError
}

// InvalidSSHKeyTitleErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) InvalidSSHKeyTitleErrorStruct() *clientruntime.StructDefinition {
	return m.structPathInvalidSSHKeyTitleError
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

// UserGetUserNotAvailableErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) UserGetUserNotAvailableErrorStruct() *clientruntime.StructDefinition {
	return m.structPathUserGetUserNotAvailableError
}

// UserDestroyInputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserDestroyInputStruct() *clientruntime.StructDefinition {
	return m.structPathUserDestroyInput
}

// UserDestroyOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserDestroyOutputStruct() *clientruntime.StructDefinition {
	return m.structPathUserDestroyOutput
}

// UserSetActiveInputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserSetActiveInputStruct() *clientruntime.StructDefinition {
	return m.structPathUserSetActiveInput
}

// UserSetActiveOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserSetActiveOutputStruct() *clientruntime.StructDefinition {
	return m.structPathUserSetActiveOutput
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

// UserSSHKeyCreateInputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserSSHKeyCreateInputStruct() *clientruntime.StructDefinition {
	return m.structPathUserSSHKeyCreateInput
}

// UserSSHKeyCreateOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserSSHKeyCreateOutputStruct() *clientruntime.StructDefinition {
	return m.structPathUserSSHKeyCreateOutput
}

// UserSSHKeyListInputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserSSHKeyListInputStruct() *clientruntime.StructDefinition {
	return m.structPathUserSSHKeyListInput
}

// UserSSHKeyListOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserSSHKeyListOutputStruct() *clientruntime.StructDefinition {
	return m.structPathUserSSHKeyListOutput
}

// UserSSHKeyDestroyInputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserSSHKeyDestroyInputStruct() *clientruntime.StructDefinition {
	return m.structPathUserSSHKeyDestroyInput
}

// UserSSHKeyDestroyOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserSSHKeyDestroyOutputStruct() *clientruntime.StructDefinition {
	return m.structPathUserSSHKeyDestroyOutput
}

// UserSSHKeySetTitleInputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserSSHKeySetTitleInputStruct() *clientruntime.StructDefinition {
	return m.structPathUserSSHKeySetTitleInput
}

// UserSSHKeySetTitleOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserSSHKeySetTitleOutputStruct() *clientruntime.StructDefinition {
	return m.structPathUserSSHKeySetTitleOutput
}

// UserGroupListInputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserGroupListInputStruct() *clientruntime.StructDefinition {
	return m.structPathUserGroupListInput
}

// UserGroupListOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) UserGroupListOutputStruct() *clientruntime.StructDefinition {
	return m.structPathUserGroupListOutput
}

// InlinePolicyStruct allows easy access to structure
func (m *SpecularMetaInfo) InlinePolicyStruct() *clientruntime.StructDefinition {
	return m.structPathInlinePolicy
}

// PolicyStructureErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) PolicyStructureErrorStruct() *clientruntime.StructDefinition {
	return m.structPathPolicyStructureError
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

// RoleAssumeWithWebIdentityInputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleAssumeWithWebIdentityInputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleAssumeWithWebIdentityInput
}

// RoleAssumeWithWebIdentityOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleAssumeWithWebIdentityOutputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleAssumeWithWebIdentityOutput
}

// RoleAccessKeyListInputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleAccessKeyListInputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleAccessKeyListInput
}

// RoleAccessKeyListOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleAccessKeyListOutputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleAccessKeyListOutput
}

// RoleAccessKeyDestroyInputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleAccessKeyDestroyInputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleAccessKeyDestroyInput
}

// RoleAccessKeyDestroyOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleAccessKeyDestroyOutputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleAccessKeyDestroyOutput
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

// RoleTrustPolicyAttachInputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleTrustPolicyAttachInputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleTrustPolicyAttachInput
}

// RoleTrustPolicyAttachOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleTrustPolicyAttachOutputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleTrustPolicyAttachOutput
}

// RoleTrustPolicyListInputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleTrustPolicyListInputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleTrustPolicyListInput
}

// RoleTrustPolicyListOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleTrustPolicyListOutputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleTrustPolicyListOutput
}

// RoleTrustPolicyDetachInputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleTrustPolicyDetachInputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleTrustPolicyDetachInput
}

// RoleTrustPolicyDetachOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) RoleTrustPolicyDetachOutputStruct() *clientruntime.StructDefinition {
	return m.structPathRoleTrustPolicyDetachOutput
}

// GroupCreateInputStruct allows easy access to structure
func (m *SpecularMetaInfo) GroupCreateInputStruct() *clientruntime.StructDefinition {
	return m.structPathGroupCreateInput
}

// GroupCreateOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) GroupCreateOutputStruct() *clientruntime.StructDefinition {
	return m.structPathGroupCreateOutput
}

// GroupDestroyInputStruct allows easy access to structure
func (m *SpecularMetaInfo) GroupDestroyInputStruct() *clientruntime.StructDefinition {
	return m.structPathGroupDestroyInput
}

// GroupDestroyOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) GroupDestroyOutputStruct() *clientruntime.StructDefinition {
	return m.structPathGroupDestroyOutput
}

// GroupGetInputStruct allows easy access to structure
func (m *SpecularMetaInfo) GroupGetInputStruct() *clientruntime.StructDefinition {
	return m.structPathGroupGetInput
}

// GroupGetOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) GroupGetOutputStruct() *clientruntime.StructDefinition {
	return m.structPathGroupGetOutput
}

// GroupListInputStruct allows easy access to structure
func (m *SpecularMetaInfo) GroupListInputStruct() *clientruntime.StructDefinition {
	return m.structPathGroupListInput
}

// GroupListOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) GroupListOutputStruct() *clientruntime.StructDefinition {
	return m.structPathGroupListOutput
}

// GroupMemberAddInputStruct allows easy access to structure
func (m *SpecularMetaInfo) GroupMemberAddInputStruct() *clientruntime.StructDefinition {
	return m.structPathGroupMemberAddInput
}

// GroupMemberAddOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) GroupMemberAddOutputStruct() *clientruntime.StructDefinition {
	return m.structPathGroupMemberAddOutput
}

// GroupMemberRemoveInputStruct allows easy access to structure
func (m *SpecularMetaInfo) GroupMemberRemoveInputStruct() *clientruntime.StructDefinition {
	return m.structPathGroupMemberRemoveInput
}

// GroupMemberRemoveOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) GroupMemberRemoveOutputStruct() *clientruntime.StructDefinition {
	return m.structPathGroupMemberRemoveOutput
}

// GroupMemberListInputStruct allows easy access to structure
func (m *SpecularMetaInfo) GroupMemberListInputStruct() *clientruntime.StructDefinition {
	return m.structPathGroupMemberListInput
}

// GroupMemberListOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) GroupMemberListOutputStruct() *clientruntime.StructDefinition {
	return m.structPathGroupMemberListOutput
}

// GroupIdentityPolicyAttachInputStruct allows easy access to structure
func (m *SpecularMetaInfo) GroupIdentityPolicyAttachInputStruct() *clientruntime.StructDefinition {
	return m.structPathGroupIdentityPolicyAttachInput
}

// GroupIdentityPolicyAttachOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) GroupIdentityPolicyAttachOutputStruct() *clientruntime.StructDefinition {
	return m.structPathGroupIdentityPolicyAttachOutput
}

// GroupIdentityPolicyDetachInputStruct allows easy access to structure
func (m *SpecularMetaInfo) GroupIdentityPolicyDetachInputStruct() *clientruntime.StructDefinition {
	return m.structPathGroupIdentityPolicyDetachInput
}

// GroupIdentityPolicyDetachOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) GroupIdentityPolicyDetachOutputStruct() *clientruntime.StructDefinition {
	return m.structPathGroupIdentityPolicyDetachOutput
}

// GroupIdentityPolicyListInputStruct allows easy access to structure
func (m *SpecularMetaInfo) GroupIdentityPolicyListInputStruct() *clientruntime.StructDefinition {
	return m.structPathGroupIdentityPolicyListInput
}

// GroupIdentityPolicyListOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) GroupIdentityPolicyListOutputStruct() *clientruntime.StructDefinition {
	return m.structPathGroupIdentityPolicyListOutput
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

// IdentityPolicyAttachmentListInputStruct allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyAttachmentListInputStruct() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyAttachmentListInput
}

// IdentityPolicyAttachmentListOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) IdentityPolicyAttachmentListOutputStruct() *clientruntime.StructDefinition {
	return m.structPathIdentityPolicyAttachmentListOutput
}

// TrustPolicyCreateInputStruct allows easy access to structure
func (m *SpecularMetaInfo) TrustPolicyCreateInputStruct() *clientruntime.StructDefinition {
	return m.structPathTrustPolicyCreateInput
}

// TrustPolicyCreateOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) TrustPolicyCreateOutputStruct() *clientruntime.StructDefinition {
	return m.structPathTrustPolicyCreateOutput
}

// TrustPolicyRetrieveInputStruct allows easy access to structure
func (m *SpecularMetaInfo) TrustPolicyRetrieveInputStruct() *clientruntime.StructDefinition {
	return m.structPathTrustPolicyRetrieveInput
}

// TrustPolicyRetrieveOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) TrustPolicyRetrieveOutputStruct() *clientruntime.StructDefinition {
	return m.structPathTrustPolicyRetrieveOutput
}

// TrustPolicyListInputStruct allows easy access to structure
func (m *SpecularMetaInfo) TrustPolicyListInputStruct() *clientruntime.StructDefinition {
	return m.structPathTrustPolicyListInput
}

// TrustPolicyListOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) TrustPolicyListOutputStruct() *clientruntime.StructDefinition {
	return m.structPathTrustPolicyListOutput
}

// TrustPolicyUpdateInputStruct allows easy access to structure
func (m *SpecularMetaInfo) TrustPolicyUpdateInputStruct() *clientruntime.StructDefinition {
	return m.structPathTrustPolicyUpdateInput
}

// TrustPolicyUpdateOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) TrustPolicyUpdateOutputStruct() *clientruntime.StructDefinition {
	return m.structPathTrustPolicyUpdateOutput
}

// TrustPolicyDestroyInputStruct allows easy access to structure
func (m *SpecularMetaInfo) TrustPolicyDestroyInputStruct() *clientruntime.StructDefinition {
	return m.structPathTrustPolicyDestroyInput
}

// TrustPolicyDestroyOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) TrustPolicyDestroyOutputStruct() *clientruntime.StructDefinition {
	return m.structPathTrustPolicyDestroyOutput
}

// TrustPolicyAttachmentListInputStruct allows easy access to structure
func (m *SpecularMetaInfo) TrustPolicyAttachmentListInputStruct() *clientruntime.StructDefinition {
	return m.structPathTrustPolicyAttachmentListInput
}

// TrustPolicyAttachmentListOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) TrustPolicyAttachmentListOutputStruct() *clientruntime.StructDefinition {
	return m.structPathTrustPolicyAttachmentListOutput
}

// ServiceBearerTokenStruct allows easy access to structure
func (m *SpecularMetaInfo) ServiceBearerTokenStruct() *clientruntime.StructDefinition {
	return m.structPathServiceBearerToken
}

// InvalidServiceBearerTokenDurationErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) InvalidServiceBearerTokenDurationErrorStruct() *clientruntime.StructDefinition {
	return m.structPathInvalidServiceBearerTokenDurationError
}

// InvalidServiceNameErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) InvalidServiceNameErrorStruct() *clientruntime.StructDefinition {
	return m.structPathInvalidServiceNameError
}

// InvitationCreateInputStruct allows easy access to structure
func (m *SpecularMetaInfo) InvitationCreateInputStruct() *clientruntime.StructDefinition {
	return m.structPathInvitationCreateInput
}

// InvitationCreateOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) InvitationCreateOutputStruct() *clientruntime.StructDefinition {
	return m.structPathInvitationCreateOutput
}

// InvitationListInputStruct allows easy access to structure
func (m *SpecularMetaInfo) InvitationListInputStruct() *clientruntime.StructDefinition {
	return m.structPathInvitationListInput
}

// InvitationListOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) InvitationListOutputStruct() *clientruntime.StructDefinition {
	return m.structPathInvitationListOutput
}

// InvitationRevokeInputStruct allows easy access to structure
func (m *SpecularMetaInfo) InvitationRevokeInputStruct() *clientruntime.StructDefinition {
	return m.structPathInvitationRevokeInput
}

// InvitationRevokeOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) InvitationRevokeOutputStruct() *clientruntime.StructDefinition {
	return m.structPathInvitationRevokeOutput
}

// InvitationAcceptInputStruct allows easy access to structure
func (m *SpecularMetaInfo) InvitationAcceptInputStruct() *clientruntime.StructDefinition {
	return m.structPathInvitationAcceptInput
}

// InvitationAcceptOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) InvitationAcceptOutputStruct() *clientruntime.StructDefinition {
	return m.structPathInvitationAcceptOutput
}

// InvitationInspectInputStruct allows easy access to structure
func (m *SpecularMetaInfo) InvitationInspectInputStruct() *clientruntime.StructDefinition {
	return m.structPathInvitationInspectInput
}

// InvitationInspectOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) InvitationInspectOutputStruct() *clientruntime.StructDefinition {
	return m.structPathInvitationInspectOutput
}

// InvitationDeclineInputStruct allows easy access to structure
func (m *SpecularMetaInfo) InvitationDeclineInputStruct() *clientruntime.StructDefinition {
	return m.structPathInvitationDeclineInput
}

// InvitationDeclineOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) InvitationDeclineOutputStruct() *clientruntime.StructDefinition {
	return m.structPathInvitationDeclineOutput
}

// ServiceBearerTokenGetInputStruct allows easy access to structure
func (m *SpecularMetaInfo) ServiceBearerTokenGetInputStruct() *clientruntime.StructDefinition {
	return m.structPathServiceBearerTokenGetInput
}

// ServiceBearerTokenGetOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) ServiceBearerTokenGetOutputStruct() *clientruntime.StructDefinition {
	return m.structPathServiceBearerTokenGetOutput
}

// SessionKeepAliveInputStruct allows easy access to structure
func (m *SpecularMetaInfo) SessionKeepAliveInputStruct() *clientruntime.StructDefinition {
	return m.structPathSessionKeepAliveInput
}

// SessionKeepAliveOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) SessionKeepAliveOutputStruct() *clientruntime.StructDefinition {
	return m.structPathSessionKeepAliveOutput
}

// SessionRevokeInputStruct allows easy access to structure
func (m *SpecularMetaInfo) SessionRevokeInputStruct() *clientruntime.StructDefinition {
	return m.structPathSessionRevokeInput
}

// SessionRevokeOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) SessionRevokeOutputStruct() *clientruntime.StructDefinition {
	return m.structPathSessionRevokeOutput
}

// SessionIdentityInputStruct allows easy access to structure
func (m *SpecularMetaInfo) SessionIdentityInputStruct() *clientruntime.StructDefinition {
	return m.structPathSessionIdentityInput
}

// SessionIdentityOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) SessionIdentityOutputStruct() *clientruntime.StructDefinition {
	return m.structPathSessionIdentityOutput
}

// SessionBeginHandoffInputStruct allows easy access to structure
func (m *SpecularMetaInfo) SessionBeginHandoffInputStruct() *clientruntime.StructDefinition {
	return m.structPathSessionBeginHandoffInput
}

// SessionBeginHandoffOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) SessionBeginHandoffOutputStruct() *clientruntime.StructDefinition {
	return m.structPathSessionBeginHandoffOutput
}

// SessionCompleteHandoffInputStruct allows easy access to structure
func (m *SpecularMetaInfo) SessionCompleteHandoffInputStruct() *clientruntime.StructDefinition {
	return m.structPathSessionCompleteHandoffInput
}

// SessionCompleteHandoffOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) SessionCompleteHandoffOutputStruct() *clientruntime.StructDefinition {
	return m.structPathSessionCompleteHandoffOutput
}

// ServiceCatalogListInputStruct allows easy access to structure
func (m *SpecularMetaInfo) ServiceCatalogListInputStruct() *clientruntime.StructDefinition {
	return m.structPathServiceCatalogListInput
}

// ServiceCatalogListOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) ServiceCatalogListOutputStruct() *clientruntime.StructDefinition {
	return m.structPathServiceCatalogListOutput
}

// ServiceCatalogGetInputStruct allows easy access to structure
func (m *SpecularMetaInfo) ServiceCatalogGetInputStruct() *clientruntime.StructDefinition {
	return m.structPathServiceCatalogGetInput
}

// ServiceCatalogGetOutputStruct allows easy access to structure
func (m *SpecularMetaInfo) ServiceCatalogGetOutputStruct() *clientruntime.StructDefinition {
	return m.structPathServiceCatalogGetOutput
}

// ServiceCatalogGetServiceCatalogNotFoundErrorStruct allows easy access to structure
func (m *SpecularMetaInfo) ServiceCatalogGetServiceCatalogNotFoundErrorStruct() *clientruntime.StructDefinition {
	return m.structPathServiceCatalogGetServiceCatalogNotFoundError
}

var localSpecularMeta *SpecularMetaInfo = &SpecularMetaInfo{}

// SpecularMeta returns metadata of the specular module
func SpecularMeta() *SpecularMetaInfo {
	return localSpecularMeta
}
