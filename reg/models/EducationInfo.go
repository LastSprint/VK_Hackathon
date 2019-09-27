package reg

type EducationTypeModel int

const (
	School         EducationTypeModel = 0
	University     EducationTypeModel = 1
	SpecialCourses EducationTypeModel = 2
)

type EducationInfoModel struct {
	EducationType EducationTypeModel `json:"educationType"`
	OrgName       string             `json:"orgName"`
	Degree        string             `json:"degree"`
}
