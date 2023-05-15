package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/luisnquin/blind-creator/test-core/internal/model"

	"github.com/jinzhu/gorm"
	"github.com/luisnquin/blind-creator-test-core-models/models/campaign_creator_social_network_actions"
	"github.com/luisnquin/blind-creator-test-core-models/models/campaigns"
	"github.com/luisnquin/blind-creator-test-core-models/models/companies"
	"github.com/luisnquin/blind-creator-test-core-models/models/creator_social_networks"
	"github.com/luisnquin/blind-creator-test-core-models/models/user_agency_relations"
	"github.com/luisnquin/blind-creator-test-core-models/models/users"
	utils "github.com/luisnquin/blind-creator-test-core-utils"
)

type AgenciesDbRepository struct {
	*gorm.DB
}

func (c AgenciesDbRepository) GetUserById(userId uint) (users.User, error) {
	var user users.User

	return user, c.Table("users").Where("id = ?", userId).First(&user).Error
}

func (c AgenciesDbRepository) GetCampaignById(
	id uint,
) (
	campaigns.Campaign,
	error,
) {
	var campaign campaigns.Campaign

	err := c.Table("campaigns").
		Where("id = ?", id).
		First(&campaign).Error
	return campaign, err
}

func (c AgenciesDbRepository) GetCompanyById(
	id uint,
) (
	companies.Company,
	error,
) {
	var company companies.Company

	err := c.Table("companies").
		Where("id = ?", id).
		First(&company).Error
	return company, err
}

func (c AgenciesDbRepository) GetSocialNetworkById(
	id uint,
) (
	creator_social_networks.CreatorSocialNetworkAccount,
	error,
) {
	var creatorSocialNetworkAccount creator_social_networks.CreatorSocialNetworkAccount

	err := c.Table("creator_social_networks").
		Where("id = ?", id).
		First(&creatorSocialNetworkAccount).Error
	return creatorSocialNetworkAccount, err
}

func (c AgenciesDbRepository) CreateCampaign(
	u campaigns.Campaign,
) (
	campaigns.Campaign,
	error,
) {
	c.Exec(
		fmt.Sprintf(
			"SELECT setval('%s_id_seq', (select max(id) from %s) + 1, FALSE);",
			"campaigns",
			"campaigns",
		),
	)
	if err := c.Create(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (c AgenciesDbRepository) CreateCampaignCreatorSocialNetworkAction(
	u campaign_creator_social_network_actions.CampaignCreatorSocialNetworkActions,
) (
	campaign_creator_social_network_actions.CampaignCreatorSocialNetworkActions,
	error,
) {
	if !u.CreatedAt.Valid {
		u.CreatedAt = sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		}
	}

	if !u.UpdatedAt.Valid {
		u.UpdatedAt = sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		}
	}

	if u.DraftContentStatus == "" {
		u.DraftContentStatus = campaign_creator_social_network_actions.ContentStatusPending
	}

	if u.FinalContentStatus == "" {
		u.FinalContentStatus = campaign_creator_social_network_actions.ContentStatusPending
	}

	c.Exec(
		fmt.Sprintf(
			"SELECT setval('%s_id_seq', (select max(id) from %s) + 1, FALSE);",
			"campaign_creator_social_network_actions",
			"campaign_creator_social_network_actions",
		),
	)
	if err := c.Create(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (a AgenciesDbRepository) ListAgencyCampaigns(
	pagination utils.GormPaginationData,
	agencyId uint,
	keywordSearch string,
) (
	utils.GormPaginationData,
	error,
) {
	if pagination.Sort == "" {
		pagination.Sort = "id desc"
	}

	var results []model.ListCampaignsResponseModel

	selectColumns := `
		campaigns.id as campaign_id,
		campaigns.created_at as campaign_created_at,
		campaigns.updated_at as campaign_updated_at,
		campaigns.deleted_at as campaign_deleted_at,
		campaigns.name as campaign_name,
		campaigns.initial_date as campaign_initial_date,
		campaigns.final_date as campaign_final_date,
		campaigns.budget as campaign_budget,
		campaigns.currency as campaign_currency,
		campaigns.agency_id as campaign_agency_id,
		campaigns.manager_id as campaign_manager_id,
		campaigns.company_id as campaign_company_id,
		campaigns.bundle_id as campaign_bundle_id,
		users.first_name as manager_name,
		users.email as manager_email,
		companies.name as company_name,
		companies.email as company_email
	`

	tx := a.Table("campaigns").
		Select(selectColumns).
		Joins("join users on users.id = campaigns.manager_id").
		Joins("join companies on companies.id = companies.manager_id").
		Where("campaigns.agency_id = ?", agencyId).
		Where("campaigns.deleted_at is null")

	if keywordSearch != "" {
		tx = tx.
			Where(
				`campaigns.name ilike ? 
			or users.first_name ilike ?
			or companies.name ilike ?`,
				"%"+keywordSearch+"%",
				"%"+keywordSearch+"%",
				"%"+keywordSearch+"%",
			)
	}

	pagination.Sort = "campaigns.id desc"

	tx = tx.Scopes(utils.Paginate(results, &pagination, tx)).
		Scan(&results)

	pagination.Rows = results

	if pagination.Rows == nil {
		pagination.Rows = make([]interface{}, 0)
		return utils.GormPaginationData{}, tx.Error
	}

	return pagination, tx.Error
}

func (c AgenciesDbRepository) IsAnManagerAgencyRelation(
	agencyId uint,
	managerId uint,
) (
	bool,
	error,
) {
	var relationExists bool
	var res user_agency_relations.UserAgencyRelation

	err := c.DB.
		Table("user_agency_relations").
		Where("user_id = ?", managerId).
		Where("agency_id = ?", agencyId).
		First(&res).Error

	if res.UserId != 0 {
		relationExists = true
	}

	return relationExists, err
}

func (c AgenciesDbRepository) IsAnAgencyCompanyRelation(
	agencyId uint,
	companyId uint,
) (
	bool,
	error,
) {
	var relationExists bool
	var company companies.Company

	err := c.DB.
		Table(company.TableName()).
		Where("id = ?", companyId).
		Where("agency_id = ?", agencyId).
		First(&company).Error

	if company.ID != 0 {
		relationExists = true
	}

	return relationExists, err
}

func (c AgenciesDbRepository) UpdateCampaign(
	u campaigns.Campaign,
) (
	campaigns.Campaign,
	error,
) {
	err := c.Save(&u).Error
	return u, err
}

func (c AgenciesDbRepository) SearchCampaignsByQuery(q string) (
	[]campaign_creator_social_network_actions.CampaignCreatorSocialNetworkActions, error,
) {
	var campaigns []campaign_creator_social_network_actions.CampaignCreatorSocialNetworkActions

	fmt.Printf("c.DB.Table(\"campaign_creator_social_network_actions\").Where(\"code_name = ?\", q).QueryExpr(): %v\n", c.DB.Table("campaign_creator_social_network_actions").Where("code_name = ?", q).QueryExpr())

	return campaigns, c.DB.Table("campaign_creator_social_network_actions").
		Where("code_name LIKE ?", "%"+strings.ToUpper(q)+"%").Find(&campaigns).Error
}

// campaign_creator_social_network_actions

func (c AgenciesDbRepository) CampaignExistsByID(id uint) (bool, error) {
	return c.rowExistsInTable("campaigns", id)
}

func (c AgenciesDbRepository) CreatorExistsByID(id uint) (bool, error) {
	return c.rowExistsInTable("creators", id)
}

func (c AgenciesDbRepository) CreatorHasSocialNetwork(creatorId, socialNetworkId uint) (bool, error) {
	var result rowWithFound

	err := c.Table("creator_social_networks").Select("COUNT(*) > 0 AS Found").
		Where("id = ? AND creator_id = ?", socialNetworkId, creatorId).Find(&result).Error

	return result.Found, err
}

func (c AgenciesDbRepository) rowExistsInTable(tableName string, id uint) (bool, error) {
	var result rowWithFound

	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE id = ?) AS Found", tableName)

	err := c.Raw(query, id).Scan(&result).Error

	return result.Found, err
}
