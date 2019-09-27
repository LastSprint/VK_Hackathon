package reg

import "reflect"

type ProgramInfo struct {
	Type              *MentorType `json:"type"`
	Reason            *string     `json:"reason"`
	Recomendation     *string     `json:"recomendation"`
	SelfRecomendation *string     `json:"selfRecomendation"`
	WhyYouWantToHelp  *string     `json:"whyYouWantToHelp"`
	Child             *ChildType  `json:"childType"`
}

type ChildType struct {
	Old             *int    `json:"old"`
	SexT            *Sex    `json:"sex,omitempty"`
	Requirements    *string `json:"requirements"`
	VisitsFrequency *string `json:"visitsFrequency"`
	IfChildIsBroken *bool   `json:"ifChildIsBroken"`
}

func (model *ProgramInfo) IsValid() bool {
	refl := reflect.ValueOf(model)

	for i := 0; i < refl.NumField(); i++ {
		if refl.Field(i).IsNil() {
			return false
		}
	}

	return true
}

func (model *ChildType) IsValid() bool {
	refl := reflect.ValueOf(model)

	for i := 0; i < refl.NumField(); i++ {
		if refl.Field(i).IsNil() {

			if refl.Type().Field(i).Name == "SexT" {
				continue
			}

			return false
		}
	}

	return true
}
