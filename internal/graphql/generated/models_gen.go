// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"fmt"
	"io"
	"strconv"
)

type Activity struct {
	ActivityID           string `json:"activityID"`
	ActivityType         string `json:"activityType"`
	DateTime             string `json:"dateTime"`
	CommunicationChannel string `json:"communicationChannel"`
	ContentNotes         string `json:"contentNotes"`
	ParticipantDetails   string `json:"participantDetails"`
	FollowUpActions      string `json:"followUpActions"`
	LeadID               string `json:"leadId"`
}

type AuthPayload struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type Campaign struct {
	CampaignID       string  `json:"campaignID"`
	CampaignName     string  `json:"campaignName"`
	CampaignCountry  string  `json:"campaignCountry"`
	CampaignRegion   string  `json:"campaignRegion"`
	IndustryTargeted string  `json:"industryTargeted"`
	Users            []*User `json:"users"`
	Leads            []*Lead `json:"leads"`
}

type CampaignFilter struct {
	CampaignName    *string `json:"campaignName,omitempty"`
	CampaignCountry *string `json:"campaignCountry,omitempty"`
}

type CampaignPage struct {
	Items      []*Campaign `json:"items"`
	TotalCount int32       `json:"totalCount"`
}

type CampaignSortInput struct {
	Field CampaignSortField `json:"field"`
	Order SortOrder         `json:"order"`
}

type Contact struct {
	ContactID   string  `json:"contactID"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
	VendorID    string  `json:"vendorId"`
	Name        string  `json:"name"`
	Email       *string `json:"email,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

type CreateActivityInput struct {
	ActivityType         string `json:"activityType"`
	DateTime             string `json:"dateTime"`
	CommunicationChannel string `json:"communicationChannel"`
	ContentNotes         string `json:"contentNotes"`
	ParticipantDetails   string `json:"participantDetails"`
	FollowUpActions      string `json:"followUpActions"`
	LeadID               string `json:"leadId"`
}

type CreateCampaignInput struct {
	CampaignName     string `json:"campaignName"`
	CampaignCountry  string `json:"campaignCountry"`
	CampaignRegion   string `json:"campaignRegion"`
	IndustryTargeted string `json:"industryTargeted"`
}

type CreateCaseStudyInput struct {
	ProjectName     string `json:"projectName"`
	ClientName      string `json:"clientName"`
	TechStack       string `json:"techStack"`
	ProjectDuration string `json:"projectDuration"`
	KeyOutcomes     string `json:"keyOutcomes"`
	IndustryTarget  string `json:"industryTarget"`
	Tags            string `json:"tags"`
	Document        string `json:"document"`
}

type CreateDealInput struct {
	DealName            string     `json:"dealName"`
	LeadID              string     `json:"leadID"`
	DealStartDate       string     `json:"dealStartDate"`
	DealEndDate         string     `json:"dealEndDate"`
	ProjectRequirements string     `json:"ProjectRequirements"`
	DealAmount          string     `json:"dealAmount"`
	DealStatus          DealStatus `json:"dealStatus"`
}

type CreateLeadInput struct {
	FirstName          string       `json:"firstName"`
	LastName           string       `json:"lastName"`
	Email              string       `json:"email"`
	LinkedIn           string       `json:"linkedIn"`
	Country            string       `json:"country"`
	Phone              string       `json:"phone"`
	LeadSource         string       `json:"leadSource"`
	InitialContactDate string       `json:"initialContactDate"`
	LeadAssignedTo     string       `json:"leadAssignedTo"`
	LeadStage          LeadStage    `json:"leadStage"`
	LeadNotes          string       `json:"leadNotes"`
	LeadPriority       LeadPriority `json:"leadPriority"`
	OrganizationID     string       `json:"organizationID"`
	CampaignID         string       `json:"campaignID"`
}

type CreateLeadWithActivityInput struct {
	Firstname            string       `json:"firstname"`
	Lastname             string       `json:"lastname"`
	Email                string       `json:"email"`
	LinkedIn             string       `json:"linkedIn"`
	Country              string       `json:"country"`
	Phone                string       `json:"phone"`
	LeadSource           string       `json:"leadSource"`
	InitialContactDate   string       `json:"initialContactDate"`
	LeadAssignedTo       string       `json:"leadAssignedTo"`
	LeadStage            LeadStage    `json:"leadStage"`
	LeadNotes            string       `json:"leadNotes"`
	LeadPriority         LeadPriority `json:"leadPriority"`
	OrganizationID       string       `json:"organizationID"`
	CampaignID           string       `json:"campaignID"`
	ActivityType         string       `json:"activityType"`
	DateTime             string       `json:"dateTime"`
	CommunicationChannel string       `json:"communicationChannel"`
	ContentNotes         string       `json:"contentNotes"`
	ParticipantDetails   string       `json:"participantDetails"`
	FollowUpActions      string       `json:"followUpActions"`
}

type CreateOrganizationInput struct {
	OrganizationName    string  `json:"organizationName"`
	OrganizationEmail   string  `json:"organizationEmail"`
	OrganizationWebsite *string `json:"organizationWebsite,omitempty"`
	City                string  `json:"city"`
	Country             string  `json:"country"`
	NoOfEmployees       string  `json:"noOfEmployees"`
	AnnualRevenue       string  `json:"annualRevenue"`
}

type CreateResourceProfileInput struct {
	Type               ResourceType   `json:"type"`
	FirstName          string         `json:"firstName"`
	LastName           string         `json:"lastName"`
	TotalExperience    float64        `json:"totalExperience"`
	ContactInformation string         `json:"contactInformation"`
	GoogleDriveLink    *string        `json:"googleDriveLink,omitempty"`
	Status             ResourceStatus `json:"status"`
	VendorID           *string        `json:"vendorId,omitempty"`
	SkillIDs           []string       `json:"skillIDs,omitempty"`
	PastProjectIds     []string       `json:"pastProjectIds,omitempty"`
}

type CreateUserInput struct {
	GoogleID *string  `json:"googleId,omitempty"`
	Name     string   `json:"name"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
	Phone    *string  `json:"phone,omitempty"`
	Role     UserRole `json:"role"`
}

type CreateVendorInput struct {
	CompanyName     string       `json:"companyName"`
	Status          VendorStatus `json:"status"`
	PaymentTerms    PaymentTerms `json:"paymentTerms"`
	Address         string       `json:"address"`
	GstOrVatDetails *string      `json:"gstOrVatDetails,omitempty"`
	Notes           *string      `json:"notes,omitempty"`
	SkillIDs        []string     `json:"skillIDs,omitempty"`
}

type Deal struct {
	DealID              string `json:"dealID"`
	DealName            string `json:"dealName"`
	LeadID              string `json:"leadID"`
	DealStartDate       string `json:"dealStartDate"`
	DealEndDate         string `json:"dealEndDate"`
	ProjectRequirements string `json:"ProjectRequirements"`
	DealAmount          string `json:"dealAmount"`
	DealStatus          string `json:"dealStatus"`
}

type Lead struct {
	LeadID             string        `json:"leadID"`
	FirstName          string        `json:"firstName"`
	LastName           string        `json:"lastName"`
	Email              string        `json:"email"`
	LinkedIn           string        `json:"linkedIn"`
	Country            string        `json:"country"`
	Phone              string        `json:"phone"`
	LeadSource         string        `json:"leadSource"`
	InitialContactDate string        `json:"initialContactDate"`
	LeadCreatedBy      *User         `json:"leadCreatedBy"`
	LeadAssignedTo     *User         `json:"leadAssignedTo"`
	LeadStage          string        `json:"leadStage"`
	LeadNotes          string        `json:"leadNotes"`
	LeadPriority       string        `json:"leadPriority"`
	Organization       *Organization `json:"organization"`
	Campaign           *Campaign     `json:"campaign"`
	Activities         []*Activity   `json:"activities"`
}

type LeadFilter struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

type LeadPage struct {
	Items      []*Lead `json:"items"`
	TotalCount int32   `json:"totalCount"`
}

type LeadSortInput struct {
	Field LeadSortField `json:"field"`
	Order SortOrder     `json:"order"`
}

type Mutation struct {
}

type Organization struct {
	OrganizationID      string  `json:"organizationID"`
	OrganizationName    string  `json:"organizationName"`
	OrganizationEmail   string  `json:"organizationEmail"`
	OrganizationWebsite *string `json:"organizationWebsite,omitempty"`
	City                string  `json:"city"`
	Country             string  `json:"country"`
	NoOfEmployees       string  `json:"noOfEmployees"`
	AnnualRevenue       string  `json:"annualRevenue"`
	Leads               []*Lead `json:"leads"`
}

type PaginationInput struct {
	Page     int32 `json:"page"`
	PageSize int32 `json:"pageSize"`
}

type PastProject struct {
	PastProjectID     string  `json:"pastProjectID"`
	CreatedAt         string  `json:"createdAt"`
	UpdatedAt         string  `json:"updatedAt"`
	ResourceProfileID string  `json:"resourceProfileId"`
	ProjectName       string  `json:"projectName"`
	Description       *string `json:"description,omitempty"`
}

type PerformanceRating struct {
	PerformanceRatingsID string  `json:"performanceRatingsID"`
	CreatedAt            string  `json:"createdAt"`
	UpdatedAt            string  `json:"updatedAt"`
	VendorID             string  `json:"vendorId"`
	Rating               int32   `json:"rating"`
	Review               *string `json:"review,omitempty"`
}

type Query struct {
}

type ResourceProfile struct {
	ResourceProfileID  string         `json:"resourceProfileID"`
	CreatedAt          string         `json:"createdAt"`
	UpdatedAt          string         `json:"updatedAt"`
	Type               ResourceType   `json:"type"`
	FirstName          string         `json:"firstName"`
	LastName           string         `json:"lastName"`
	TotalExperience    float64        `json:"totalExperience"`
	ContactInformation string         `json:"contactInformation"`
	GoogleDriveLink    *string        `json:"googleDriveLink,omitempty"`
	Status             ResourceStatus `json:"status"`
	VendorID           string         `json:"vendorId"`
	Vendor             *Vendor        `json:"vendor,omitempty"`
	Skills             []*Skill       `json:"skills"`
	PastProjects       []*PastProject `json:"pastProjects"`
}

type ResourceProfileFilter struct {
	Type               *ResourceType   `json:"type,omitempty"`
	FirstName          *string         `json:"firstName,omitempty"`
	LastName           *string         `json:"lastName,omitempty"`
	TotalExperienceMin *float64        `json:"totalExperienceMin,omitempty"`
	TotalExperienceMax *float64        `json:"totalExperienceMax,omitempty"`
	Status             *ResourceStatus `json:"status,omitempty"`
	VendorID           *string         `json:"vendorId,omitempty"`
	SkillIDs           []string        `json:"skillIDs,omitempty"`
	Search             *string         `json:"search,omitempty"`
}

type ResourceProfilePage struct {
	Items      []*ResourceProfile `json:"items"`
	TotalCount int32              `json:"totalCount"`
}

type ResourceProfileSortInput struct {
	Field ResourceProfileSortField `json:"field"`
	Order SortOrder                `json:"order"`
}

type Skill struct {
	SkillID     string  `json:"skillID"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type UpdateActivityInput struct {
	ActivityType         *string `json:"activityType,omitempty"`
	DateTime             *string `json:"dateTime,omitempty"`
	CommunicationChannel *string `json:"communicationChannel,omitempty"`
	ContentNotes         *string `json:"contentNotes,omitempty"`
	ParticipantDetails   *string `json:"participantDetails,omitempty"`
	FollowUpActions      *string `json:"followUpActions,omitempty"`
}

type UpdateCaseStudyInput struct {
	ProjectName     string `json:"projectName"`
	ClientName      string `json:"clientName"`
	TechStack       string `json:"techStack"`
	ProjectDuration string `json:"projectDuration"`
	KeyOutcomes     string `json:"keyOutcomes"`
	IndustryTarget  string `json:"industryTarget"`
	Tags            string `json:"tags"`
	Document        string `json:"document"`
}

type UpdateLeadInput struct {
	FirstName          *string      `json:"firstName,omitempty"`
	LastName           *string      `json:"lastName,omitempty"`
	Email              string       `json:"email"`
	LinkedIn           *string      `json:"linkedIn,omitempty"`
	Country            *string      `json:"country,omitempty"`
	Phone              *string      `json:"phone,omitempty"`
	LeadSource         string       `json:"leadSource"`
	InitialContactDate string       `json:"initialContactDate"`
	LeadAssignedTo     string       `json:"leadAssignedTo"`
	LeadStage          LeadStage    `json:"leadStage"`
	LeadNotes          string       `json:"leadNotes"`
	LeadPriority       LeadPriority `json:"leadPriority"`
	OrganizationID     string       `json:"organizationID"`
	CampaignID         string       `json:"campaignID"`
}

type UpdateResourceProfileInput struct {
	Type               *ResourceType   `json:"type,omitempty"`
	FirstName          *string         `json:"firstName,omitempty"`
	LastName           *string         `json:"lastName,omitempty"`
	TotalExperience    *float64        `json:"totalExperience,omitempty"`
	ContactInformation *string         `json:"contactInformation,omitempty"`
	GoogleDriveLink    *string         `json:"googleDriveLink,omitempty"`
	Status             *ResourceStatus `json:"status,omitempty"`
	VendorID           *string         `json:"vendorId,omitempty"`
	SkillIDs           []string        `json:"skillIDs,omitempty"`
	PastProjectIds     []string        `json:"pastProjectIds,omitempty"`
}

type UpdateUserInput struct {
	Name  *string   `json:"name,omitempty"`
	Email *string   `json:"email,omitempty"`
	Phone *string   `json:"phone,omitempty"`
	Role  *UserRole `json:"role,omitempty"`
}

type UpdateVendorInput struct {
	CompanyName     *string       `json:"companyName,omitempty"`
	Status          *VendorStatus `json:"status,omitempty"`
	PaymentTerms    *PaymentTerms `json:"paymentTerms,omitempty"`
	Address         *string       `json:"address,omitempty"`
	GstOrVatDetails *string       `json:"gstOrVatDetails,omitempty"`
	Notes           *string       `json:"notes,omitempty"`
	SkillIDs        []string      `json:"skillIDs,omitempty"`
}

type User struct {
	UserID    string      `json:"userID"`
	GoogleID  *string     `json:"googleId,omitempty"`
	Name      string      `json:"name"`
	Email     string      `json:"email"`
	Phone     string      `json:"phone"`
	Role      string      `json:"role"`
	Password  string      `json:"password"`
	Campaigns []*Campaign `json:"campaigns"`
}

type UserFilter struct {
	Name   *string   `json:"name,omitempty"`
	Email  *string   `json:"email,omitempty"`
	Role   *UserRole `json:"role,omitempty"`
	Search *string   `json:"search,omitempty"`
}

type UserPage struct {
	Items      []*User `json:"items"`
	TotalCount int32   `json:"totalCount"`
}

type UserSortInput struct {
	Field UserSortField `json:"field"`
	Order SortOrder     `json:"order"`
}

type Vendor struct {
	VendorID           string               `json:"vendorID"`
	CreatedAt          string               `json:"createdAt"`
	UpdatedAt          string               `json:"updatedAt"`
	CompanyName        string               `json:"companyName"`
	Status             VendorStatus         `json:"status"`
	PaymentTerms       PaymentTerms         `json:"paymentTerms"`
	Address            string               `json:"address"`
	GstOrVatDetails    *string              `json:"gstOrVatDetails,omitempty"`
	Notes              *string              `json:"notes,omitempty"`
	ContactList        []*Contact           `json:"contactList"`
	Skills             []*Skill             `json:"skills"`
	PerformanceRatings []*PerformanceRating `json:"performanceRatings"`
	Resources          []*ResourceProfile   `json:"resources"`
}

type VendorFilter struct {
	CompanyName  *string       `json:"companyName,omitempty"`
	Status       *VendorStatus `json:"status,omitempty"`
	PaymentTerms *PaymentTerms `json:"paymentTerms,omitempty"`
	Search       *string       `json:"search,omitempty"`
	SkillIDs     []string      `json:"skillIDs,omitempty"`
}

type VendorPage struct {
	Items      []*Vendor `json:"items"`
	TotalCount int32     `json:"totalCount"`
}

type VendorSortInput struct {
	Field VendorSortField `json:"field"`
	Order SortOrder       `json:"order"`
}

type CaseStudy struct {
	CaseStudyID     string `json:"caseStudyID"`
	ProjectName     string `json:"projectName"`
	ClientName      string `json:"clientName"`
	TechStack       string `json:"techStack"`
	ProjectDuration string `json:"projectDuration"`
	KeyOutcomes     string `json:"keyOutcomes"`
	IndustryTarget  string `json:"industryTarget"`
	Tags            string `json:"tags"`
	Document        string `json:"document"`
}

type CampaignSortField string

const (
	CampaignSortFieldCampaignName CampaignSortField = "CAMPAIGN_NAME"
	CampaignSortFieldCreatedAt    CampaignSortField = "CREATED_AT"
)

var AllCampaignSortField = []CampaignSortField{
	CampaignSortFieldCampaignName,
	CampaignSortFieldCreatedAt,
}

func (e CampaignSortField) IsValid() bool {
	switch e {
	case CampaignSortFieldCampaignName, CampaignSortFieldCreatedAt:
		return true
	}
	return false
}

func (e CampaignSortField) String() string {
	return string(e)
}

func (e *CampaignSortField) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CampaignSortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CampaignSortField", str)
	}
	return nil
}

func (e CampaignSortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type LeadPriority string

const (
	LeadPriorityHigh   LeadPriority = "HIGH"
	LeadPriorityMedium LeadPriority = "MEDIUM"
	LeadPriorityLow    LeadPriority = "LOW"
)

var AllLeadPriority = []LeadPriority{
	LeadPriorityHigh,
	LeadPriorityMedium,
	LeadPriorityLow,
}

func (e LeadPriority) IsValid() bool {
	switch e {
	case LeadPriorityHigh, LeadPriorityMedium, LeadPriorityLow:
		return true
	}
	return false
}

func (e LeadPriority) String() string {
	return string(e)
}

func (e *LeadPriority) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LeadPriority(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LeadPriority", str)
	}
	return nil
}

func (e LeadPriority) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type LeadSortField string

const (
	LeadSortFieldFirstName LeadSortField = "FIRST_NAME"
	LeadSortFieldLastName  LeadSortField = "LAST_NAME"
	LeadSortFieldEmail     LeadSortField = "EMAIL"
	LeadSortFieldCreatedAt LeadSortField = "CREATED_AT"
)

var AllLeadSortField = []LeadSortField{
	LeadSortFieldFirstName,
	LeadSortFieldLastName,
	LeadSortFieldEmail,
	LeadSortFieldCreatedAt,
}

func (e LeadSortField) IsValid() bool {
	switch e {
	case LeadSortFieldFirstName, LeadSortFieldLastName, LeadSortFieldEmail, LeadSortFieldCreatedAt:
		return true
	}
	return false
}

func (e LeadSortField) String() string {
	return string(e)
}

func (e *LeadSortField) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LeadSortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LeadSortField", str)
	}
	return nil
}

func (e LeadSortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type LeadStage string

const (
	LeadStageNew        LeadStage = "NEW"
	LeadStageInProgress LeadStage = "IN_PROGRESS"
	LeadStageFollowUp   LeadStage = "FOLLOW_UP"
	LeadStageClosedWon  LeadStage = "CLOSED_WON"
	LeadStageClosedLost LeadStage = "CLOSED_LOST"
)

var AllLeadStage = []LeadStage{
	LeadStageNew,
	LeadStageInProgress,
	LeadStageFollowUp,
	LeadStageClosedWon,
	LeadStageClosedLost,
}

func (e LeadStage) IsValid() bool {
	switch e {
	case LeadStageNew, LeadStageInProgress, LeadStageFollowUp, LeadStageClosedWon, LeadStageClosedLost:
		return true
	}
	return false
}

func (e LeadStage) String() string {
	return string(e)
}

func (e *LeadStage) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LeadStage(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LeadStage", str)
	}
	return nil
}

func (e LeadStage) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PaymentTerms string

const (
	PaymentTermsNet30 PaymentTerms = "NET_30"
	PaymentTermsNet60 PaymentTerms = "NET_60"
	PaymentTermsNet90 PaymentTerms = "NET_90"
)

var AllPaymentTerms = []PaymentTerms{
	PaymentTermsNet30,
	PaymentTermsNet60,
	PaymentTermsNet90,
}

func (e PaymentTerms) IsValid() bool {
	switch e {
	case PaymentTermsNet30, PaymentTermsNet60, PaymentTermsNet90:
		return true
	}
	return false
}

func (e PaymentTerms) String() string {
	return string(e)
}

func (e *PaymentTerms) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PaymentTerms(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PaymentTerms", str)
	}
	return nil
}

func (e PaymentTerms) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ResourceProfileSortField string

const (
	ResourceProfileSortFieldCreatedAt       ResourceProfileSortField = "createdAt"
	ResourceProfileSortFieldUpdatedAt       ResourceProfileSortField = "updatedAt"
	ResourceProfileSortFieldFirstName       ResourceProfileSortField = "firstName"
	ResourceProfileSortFieldLastName        ResourceProfileSortField = "lastName"
	ResourceProfileSortFieldTotalExperience ResourceProfileSortField = "totalExperience"
	ResourceProfileSortFieldStatus          ResourceProfileSortField = "status"
)

var AllResourceProfileSortField = []ResourceProfileSortField{
	ResourceProfileSortFieldCreatedAt,
	ResourceProfileSortFieldUpdatedAt,
	ResourceProfileSortFieldFirstName,
	ResourceProfileSortFieldLastName,
	ResourceProfileSortFieldTotalExperience,
	ResourceProfileSortFieldStatus,
}

func (e ResourceProfileSortField) IsValid() bool {
	switch e {
	case ResourceProfileSortFieldCreatedAt, ResourceProfileSortFieldUpdatedAt, ResourceProfileSortFieldFirstName, ResourceProfileSortFieldLastName, ResourceProfileSortFieldTotalExperience, ResourceProfileSortFieldStatus:
		return true
	}
	return false
}

func (e ResourceProfileSortField) String() string {
	return string(e)
}

func (e *ResourceProfileSortField) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ResourceProfileSortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ResourceProfileSortField", str)
	}
	return nil
}

func (e ResourceProfileSortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ResourceStatus string

const (
	ResourceStatusActive   ResourceStatus = "ACTIVE"
	ResourceStatusInactive ResourceStatus = "INACTIVE"
	ResourceStatusOnBench  ResourceStatus = "ON_BENCH"
)

var AllResourceStatus = []ResourceStatus{
	ResourceStatusActive,
	ResourceStatusInactive,
	ResourceStatusOnBench,
}

func (e ResourceStatus) IsValid() bool {
	switch e {
	case ResourceStatusActive, ResourceStatusInactive, ResourceStatusOnBench:
		return true
	}
	return false
}

func (e ResourceStatus) String() string {
	return string(e)
}

func (e *ResourceStatus) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ResourceStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ResourceStatus", str)
	}
	return nil
}

func (e ResourceStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ResourceType string

const (
	ResourceTypeConsultant ResourceType = "CONSULTANT"
	ResourceTypeFreelancer ResourceType = "FREELANCER"
	ResourceTypeContractor ResourceType = "CONTRACTOR"
	ResourceTypeEmployee   ResourceType = "EMPLOYEE"
)

var AllResourceType = []ResourceType{
	ResourceTypeConsultant,
	ResourceTypeFreelancer,
	ResourceTypeContractor,
	ResourceTypeEmployee,
}

func (e ResourceType) IsValid() bool {
	switch e {
	case ResourceTypeConsultant, ResourceTypeFreelancer, ResourceTypeContractor, ResourceTypeEmployee:
		return true
	}
	return false
}

func (e ResourceType) String() string {
	return string(e)
}

func (e *ResourceType) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ResourceType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ResourceType", str)
	}
	return nil
}

func (e ResourceType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SortOrder string

const (
	SortOrderAsc  SortOrder = "ASC"
	SortOrderDesc SortOrder = "DESC"
)

var AllSortOrder = []SortOrder{
	SortOrderAsc,
	SortOrderDesc,
}

func (e SortOrder) IsValid() bool {
	switch e {
	case SortOrderAsc, SortOrderDesc:
		return true
	}
	return false
}

func (e SortOrder) String() string {
	return string(e)
}

func (e *SortOrder) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortOrder(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortOrder", str)
	}
	return nil
}

func (e SortOrder) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserRole string

const (
	UserRoleAdmin          UserRole = "ADMIN"
	UserRoleSalesExecutive UserRole = "SALES_EXECUTIVE"
	UserRoleManager        UserRole = "MANAGER"
)

var AllUserRole = []UserRole{
	UserRoleAdmin,
	UserRoleSalesExecutive,
	UserRoleManager,
}

func (e UserRole) IsValid() bool {
	switch e {
	case UserRoleAdmin, UserRoleSalesExecutive, UserRoleManager:
		return true
	}
	return false
}

func (e UserRole) String() string {
	return string(e)
}

func (e *UserRole) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserRole", str)
	}
	return nil
}

func (e UserRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserSortField string

const (
	UserSortFieldCreatedAt UserSortField = "createdAt"
	UserSortFieldName      UserSortField = "name"
	UserSortFieldEmail     UserSortField = "email"
	UserSortFieldRole      UserSortField = "role"
)

var AllUserSortField = []UserSortField{
	UserSortFieldCreatedAt,
	UserSortFieldName,
	UserSortFieldEmail,
	UserSortFieldRole,
}

func (e UserSortField) IsValid() bool {
	switch e {
	case UserSortFieldCreatedAt, UserSortFieldName, UserSortFieldEmail, UserSortFieldRole:
		return true
	}
	return false
}

func (e UserSortField) String() string {
	return string(e)
}

func (e *UserSortField) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserSortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserSortField", str)
	}
	return nil
}

func (e UserSortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type VendorSortField string

const (
	VendorSortFieldCreatedAt   VendorSortField = "createdAt"
	VendorSortFieldUpdatedAt   VendorSortField = "updatedAt"
	VendorSortFieldCompanyName VendorSortField = "companyName"
	VendorSortFieldStatus      VendorSortField = "status"
)

var AllVendorSortField = []VendorSortField{
	VendorSortFieldCreatedAt,
	VendorSortFieldUpdatedAt,
	VendorSortFieldCompanyName,
	VendorSortFieldStatus,
}

func (e VendorSortField) IsValid() bool {
	switch e {
	case VendorSortFieldCreatedAt, VendorSortFieldUpdatedAt, VendorSortFieldCompanyName, VendorSortFieldStatus:
		return true
	}
	return false
}

func (e VendorSortField) String() string {
	return string(e)
}

func (e *VendorSortField) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = VendorSortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid VendorSortField", str)
	}
	return nil
}

func (e VendorSortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type VendorStatus string

const (
	VendorStatusActive    VendorStatus = "ACTIVE"
	VendorStatusInactive  VendorStatus = "INACTIVE"
	VendorStatusPreferred VendorStatus = "PREFERRED"
)

var AllVendorStatus = []VendorStatus{
	VendorStatusActive,
	VendorStatusInactive,
	VendorStatusPreferred,
}

func (e VendorStatus) IsValid() bool {
	switch e {
	case VendorStatusActive, VendorStatusInactive, VendorStatusPreferred:
		return true
	}
	return false
}

func (e VendorStatus) String() string {
	return string(e)
}

func (e *VendorStatus) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = VendorStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid VendorStatus", str)
	}
	return nil
}

func (e VendorStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DealStatus string

const (
	DealStatusStarted   DealStatus = "STARTED"
	DealStatusPending   DealStatus = "PENDING"
	DealStatusCompleted DealStatus = "COMPLETED"
)

var AllDealStatus = []DealStatus{
	DealStatusStarted,
	DealStatusPending,
	DealStatusCompleted,
}

func (e DealStatus) IsValid() bool {
	switch e {
	case DealStatusStarted, DealStatusPending, DealStatusCompleted:
		return true
	}
	return false
}

func (e DealStatus) String() string {
	return string(e)
}

func (e *DealStatus) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DealStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid dealStatus", str)
	}
	return nil
}

func (e DealStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
