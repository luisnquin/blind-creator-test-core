package creators

import (
	"database/sql"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
	"go-backend-challenge/core-models-private-library/models"
)

const (
	OnboardingStatusAccountData             = "ACCOUNT_DATA"
	OnboardingStatusMetaLinkAccount         = "META_LINK_ACCOUNT"
	OnboardingStatusMetaRates               = "META_RATES"
	OnboardingStatusTiktokLinkAccount       = "TIKTOK_LINK_ACCOUNT"
	OnboardingStatusTiktokRates             = "TIKTOK_RATES"
	OnboardingStatusOtherSocialNetworks     = "OTHER_SOCIAL_NETWORKS"
	OnboardingStatusOtherSocialNetworkRates = "OTHER_SOCIAL_NETWORK_RATES"
	OnboardingStatusWizardCompleted         = "WIZARD_COMPLETED"
	OnboardingStatusCreatorSiteCreated      = "CREATOR_SITE_CREATED"
)

type Creator struct {
	models.BaseModelIDInt
	Username               sql.NullString `gorm:"column:username" sql:"type:varchar(250); not null; unique_index" json:"username"`
	Name                   sql.NullString `gorm:"column:name" sql:"type:varchar(250);" json:"name"`
	Avatar                 sql.NullString `gorm:"column:avatar" sql:"type:varchar(250);" json:"avatar"`
	Email                  sql.NullString `gorm:"column:email" sql:"type:varchar(250);" json:"email"`
	PhoneNumber            sql.NullString `gorm:"column:phone_number" sql:"type:varchar(250);" json:"phone_number"`
	PhoneNumberCountryCode sql.NullString `gorm:"column:phone_number_country_code" sql:"type:varchar(250);" json:"phone_number_country_code"`
	Gender                 sql.NullString `gorm:"column:gender" sql:"type:varchar(250);" json:"gender"`
	BirthDate              sql.NullTime   `gorm:"column:birth_date"  json:"birth_date"`
	Categories             pq.StringArray `gorm:"column:categories" sql:"type:TEXT[];" json:"categories"`
	Languages              pq.StringArray `gorm:"column:languages" sql:"type:TEXT[];" json:"languages"`
	Location               postgres.Jsonb `gorm:"column:location;type:jsonb;default:'{}'" json:"location"`
	OnboardingStatus       string         `gorm:"column:onboarding_status" sql:"type:varchar(25)" json:"onboarding_status"`
	OnboardingVersion      string         `gorm:"column:onboarding_version" sql:"type:varchar(30);" json:"onboarding_version"`
	// RELATIONS
	AgencyId uint `gorm:"column:agency_id" sql:"type:integer;" json:"agency_id"`
}

func (model Creator) TableName() string {
	return "creators"
}
