package reg

type FamilyStatusModel int

const (
	Married       FamilyStatusModel = 0
	CivilMarriage FamilyStatusModel = 1
	Divorced      FamilyStatusModel = 2
	Dowager       FamilyStatusModel = 3
	NotMarried    FamilyStatusModel = 4
)
