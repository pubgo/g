package certification

import "github.com/pubgo/g/models"

type Certification struct {
	models.BaseModel

	InUid                string `json:"in_uid"`
	CompanyName          string `json:"company_name"`
	RegisteredNumber     string `json:"registered_number"`
	BusinessLicenseImgId string `json:"business_license_img_id"`
	CertificationImgId   string `json:"certification_img_id"`
	CompanyDomain        string `json:"company_domain"`
	Remark               string `json:"remark"`
}
