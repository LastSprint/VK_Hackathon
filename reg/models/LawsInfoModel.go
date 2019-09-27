package reg

type LawsInfoModel struct {
	UseAlcohol          *string `json:"useAlcohol"`
	UsePsychotropic     *string `json:"usePsychotropic"`
	UseCigarets         *string `json:"useCigarets"`
	UseDrags            *string `json:"useDrags"`
	HaveACrimeRecords   *string `json:"haveACrimeRecords"`
	HaveParentalRights  *string `json:"haveParentalRights"`
	ApplyReportsRight   *bool   `json:"applyReportsRight"`
	ApplyNoPrivacyRight *bool   `json:"applyNoPrivacyRight"`
	WhereYouFindInfo    *string `json:"whereYouFindInfo"`
}
