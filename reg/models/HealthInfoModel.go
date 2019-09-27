package reg

import "reflect"

type HealthStatusModel int

const (
	Great  HealthStatusModel = 0
	Good   HealthStatusModel = 1
	Middle HealthStatusModel = 2
	Bad    HealthStatusModel = 3
)

type HealthInfoModel struct {
	HealthStatus *HealthStatusModel `json:"healthStatus"`
	Desease      *string            `json:"desease"`
}

func (model *HealthInfoModel) IsValid() bool {
	refl := reflect.ValueOf(model)

	for i := 0; i < refl.NumField(); i++ {
		if refl.Field(i).IsNil() {
			return false
		}
	}

	return true
}
