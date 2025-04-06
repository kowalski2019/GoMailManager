package models

type AutoReplyRequest struct {
	Enabled   bool   `json:"enabled"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Subject   string `json:"subject"`
	Message   string `json:"message"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type EmailRuleRequest struct {
	Id             int64  `json:"id"`
	UserEmail      string `json:"user_email"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Condition      string `json:"condition"`
	ConditionValue string `json:"conditionValue"`
	Action         string `json:"action"`
	ActionValue    string `json:"actionValue"`
}

type UserGetMailboxesRequest struct {
	Username string `json:"username"`
}
