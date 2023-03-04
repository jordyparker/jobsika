package storage

import (
	"github.com/elhmn/jobsika/pkg/models/v1beta"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (db DB) GetJobOffers(page string, limit string, jobtitle string, company string, isRemote string) (v1beta.JobOffersResponse, error) {
	offset, limitInt := Paginate(page, limit)
	var nbHits int64

	// Build the query
	query := db.queryJobOffers().Order("jb.createdat DESC")

	if jobtitle != "" {
		query = query.Where("j.title LIKE ?", "%"+jobtitle+"%")
	}

	if company != "" {
		query = query.Where("jb.company_name LIKE ?", "%"+company+"%")
	}

	switch isRemote {
	case "true":
		query = query.Where("jb.is_remote = ?", true)
	case "false":
		query = query.Where("jb.is_remote = ?", false)
	}

	rows, err := query.Count(&nbHits).Offset(offset).Limit(limitInt).Rows()
	if err != nil {
		return v1beta.JobOffersResponse{}, err
	}

	jobOffers := make([]v1beta.JobOfferPresenter, 0)
	for rows.Next() {
		j := v1beta.JobOfferPresenter{}
		err := db.c.ScanRows(rows, &j)
		if err != nil {
			return v1beta.JobOffersResponse{}, err
		}
		jobOffers = append(jobOffers, j)
	}

	resp := v1beta.JobOffersResponse{
		Hits:   jobOffers,
		NbHits: nbHits,
		Offset: int64(offset),
		Limit:  int64(limitInt),
	}

	return resp, nil
}

// PostJobOffer post new job offer
func (db DB) PostJobOffer(query v1beta.OfferPostQuery) (*v1beta.JobOffer, error) {
	offer := v1beta.JobOffer{
		CompanyName:       query.CompanyName,
		CompanyEmail:      query.CompanyEmail,
		IsRemote:          query.IsRemote,
		Location:          query.Location,
		Department:        query.Department,
		SalaryRangeMin:    query.SalaryRangeMin,
		SalaryRangeMax:    query.SalaryRangeMax,
		Description:       query.Description,
		Benefits:          query.Benefits,
		HowToApply:        query.HowToApply,
		ApplyUrl:          query.ApplyUrl,
		ApplyEmailAddress: query.ApplyEmailAddress,
		ApplyPhoneNumber:  query.ApplyPhoneNumber,
		Tags:              query.Tags,
	}

	err := db.c.Transaction(func(tx *gorm.DB) error {
		jobTitle, err := postJobTitle(tx, query.JobTitle)
		if err != nil {
			log.Error(err)
			return err
		}

		offer.TitleID = jobTitle.ID
		res := tx.Table("job_offers").Create(&offer)
		if res.Error != nil {
			log.Error(res.Error)
			return res.Error
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &offer, nil
}

func (db DB) queryJobOffers() *gorm.DB {
	return db.c.Table("job_offers as jb").Select(`
		jb.id,
		jb.createdat,
		jb.updatedat,
		jb.company_name,
		jb.company_email,
		jb.title_id,
		jb.is_remote,
		jb.location,
		jb.department,
		jb.salary_range_min,
		jb.salary_range_max,
		jb.description,
		jb.benefits,
		jb.how_to_apply,
		jb.apply_url,
		jb.apply_email_address,
		jb.apply_phone_number,
		jb.tags,
		jt.title as job_title
	`).
		Joins("left join jobtitles as jt on jb.title_id = jt.id")
}

// GetJobOfferById get job offers by id
func (db DB) GetJobOfferById(id int64) (*v1beta.JobOffer, error) {
	var offer v1beta.JobOffer
	res := db.queryJobOffers().Where("jb.id = ?", id).First(&offer)
	if res.Error != nil {
		log.Error(res.Error)
		return nil, res.Error
	}

	return &offer, nil
}
