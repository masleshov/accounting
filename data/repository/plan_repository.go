package repository

// PlanRepository represents type which can manage data related with financial plans
type PlanRepository struct {
	*Repository
}

// NewPlanRepository creates new instance of PlanRepository type
func NewPlanRepository() *PlanRepository {
	return &PlanRepository{
		Repository: NewRepository(),
	}
}

// func (pr *PlanRepository) InsertPlan()

func selectAllPlansQuery() string {
	return "Select " +
		"From ACCOUNTING.Plan"
}

func insertPlanQuery() string {
	return "Insert into ACCOUNTING.Plan(User_ID, Date_Beg, Date_End, Plan_No) " +
		"Values($1, $2, $3, $4) " +
		"Returning Plan_ID"
}

func updatePlanQuery() string {
	return "Update ACCOUNTING.Plan " +
		"Set User_ID = $2 " +
		"  , Date_Beg = $3" +
		"  , Date_End = $4" +
		"  , Plan_No = $5" +
		"Where Plan_ID = $1"
}

func deletePlanQuery() string {
	return "Delete From ACCOUNTING.Plan " +
		"Where Plan_ID = $1"
}
