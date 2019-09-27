package reg

type FormStatusModel int

const (
	New      FormStatusModel = 0
	Viewed   FormStatusModel = 1
	Called   FormStatusModel = 2
	Approves FormStatusModel = 3
)

type FormModel struct {
	LawsInf               *LawsInfoModel          `json:"lawsInfo"`
	ProgramInf            *ProgramInfo            `json:"programInfo"`
	WorkWithChilredExpInf *WorkWithChilredExpInfo `json:"workWithChilredExpInfo"`
	FamilyInfo            *FamilyInfoModel        `json:"familyInfo"`
	Hobby                 *string                 `json:"hobbyInfo"`
	JobInfo               *JobInfoModel           `json:"jobInfo"`
	EducationInfo         *EducationInfoModel     `json:"educationInfo"`
	HealthInfo            *HealthInfoModel        `json:"healthInfo"`
	FormStatus            *FormStatusModel        `json:"formStatus,omitempty"`
	UserInfo              *UserInfoModel          `json:"userInfo"`
}

func (model *FormModel) IsValid() bool {
	return model.LawsInf != nil &&
		model.ProgramInf != nil &&
		model.FamilyInfo != nil &&
		model.JobInfo != nil &&
		model.HealthInfo != nil &&
		model.Hobby != nil &&
		model.UserInfo != nil
}
