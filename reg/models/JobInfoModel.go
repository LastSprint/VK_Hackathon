package reg

import "reflect"

type JobInfoModel struct {
	OrgName          *string `json:"orgName"`
	Contacts         *string `json:"contacts"`
	Position         *string `json:"position"`
	Responsibilities *string `json:"responsibilities"`
	WorkTimeTable    *string `json:"workTimeTable"`
}

func (model *JobInfoModel) IsValid() bool {
	refl := reflect.ValueOf(model)

	for i := 0; i < refl.NumField(); i++ {
		if refl.Field(i).IsNil() {
			return false
		}
	}

	return true
}
