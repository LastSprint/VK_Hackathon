package reg

import "reflect"

type UserInfoModel struct {
	FirstName      *string `json:"firstName"`
	SecondName     *string `json:"secondName"`
	LastName       *string `json:"lastName"`
	RegisterAdress *string `json:"registerAdress"`
	LivingAddress  *string `json:"livingAddress"`
	Phone          *string `json:"phone"`
	Email          *string `json:"email"`
	Birthday       *string `json:"birthday"`
	Religion       *string `json:"religion"`
	IsRFCitizen    *string `json:"isRFCitizen"`
}

func (model *UserInfoModel) IsValid() bool {
	refl := reflect.ValueOf(model)

	for i := 0; i < refl.NumField(); i++ {
		if refl.Field(i).IsNil() {
			if refl.Type().Field(i).Name == "LivingAddress" && !refl.FieldByName("RegisterAdress").IsNil() {
				continue
			}
			return false
		}
	}

	return true
}
