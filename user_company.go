package bunq

type Alias struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Name  string `json:"name"`
}

type AvatarImage struct {
	AttachmentPublicUUID string `json:"attachment_public_uuid"`
	ContentType          string `json:"content_type"`
	Height               int    `json:"height"`
	Width                int    `json:"width"`
}
type UBO struct {
	Name        string `json:"name"`
	DateOfBirth Time   `json:"date_of_birth"`
	Nationality string `json:"nationality"`
}
type NotificationFilter struct {
	NotificationDeliveryMethod string `json:"notification_delivery_method"`
	NotificationTarget         string `json:"notification_target"`
	Category                   string `json:"category"`
}

type Address struct {
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
	POBox       string `json:"po_box"`
	PostalCode  string `json:"postal_code"`
	City        string `json:"city"`
	Country     string `json:"country"`
}

type Avatar struct {
	UUID       string        `json:"uuid"`
	AnchorUUID string        `json:"anchor_uuid"`
	Image      []AvatarImage `json:"image"`
}

type DirectorAlias struct {
	UUID           string `json:"uuid"`
	Avatar         Avatar `json:"avatar"`
	PublicNickName string `json:"public_nick_name"`
	DisplayName    string `json:"display_name"`
	Country        string `json:"country"`
}

type Limit struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type UserCompany struct {
	ID                                 int                  `json:"id"`
	CreatedAt                          Time                 `json:"created"`
	UpdatedAt                          Time                 `json:"updated"`
	PublicUUID                         string               `json:"public_uuid"`
	Name                               string               `json:"name"`
	DisplayName                        string               `json:"display_name"`
	PublicNickName                     string               `json:"public_nick_name"`
	Alias                              []Alias              `json:"alias"`
	ChamberOfCommerceNumber            string               `json:"chamber_of_commerce_number"`
	TypeOfBusinessEntity               string               `json:"type_of_business_entity"`
	SectorOfIndustry                   string               `json:"sector_of_industry"`
	CounterBankIBAN                    string               `json:"counter_bank_IBAN"`
	Avatar                             Avatar               `json:"avatar"`
	AddressMain                        Address              `json:"address_main"`
	AddressPostal                      Address              `json:"address_postal"`
	VersionTermsOfService              string               `json:"version_terms_of_service"`
	DirectorAlias                      DirectorAlias        `json:"director_alias"`
	Language                           string               `json:"language"`
	Region                             string               `json:"region"`
	UBO                                []UBO                `json:"UBO"`
	Status                             string               `json:"status"`
	SubStatus                          string               `json:"sub_status"`
	SessionTimeout                     int                  `json:"session_timeout"`
	DailyLimitWithoutConfirmationLogin Limit                `json:"daily_limit_without_confirmation_login"`
	NotificationFilters                []NotificationFilter `json:"notification_filters"`
}
