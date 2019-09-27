package reg

type CheckStatusRequestModel struct {
	Value *string `json:"value"`
}

func (model *CheckStatusRequestModel) IsValid() bool {
	return model.Value != nil
}

type CheckStatusResponseModel struct {
	Status FormStatusModel `json:"status"`
}
