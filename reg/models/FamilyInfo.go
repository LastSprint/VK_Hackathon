package reg

type FamilyInfoModel struct {
	Status        *FamilyStatusModel `json:"status"`
	PartnerName   *string            `json:"partnerName"`
	PartnerSex    *Sex               `json:"partnerSex"`
	PartnerOld    *int               `json:"partnerOld"`
	Relationships *string            `json:"relationships"`
}

func (model *FamilyInfoModel) IsValid() bool {
	return model.Status != nil
}
