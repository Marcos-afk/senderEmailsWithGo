package contracts

type CreateCampaignRequest struct {
	Name    string
	Content string
	Emails  []string
}
type CreateCampaign struct {
	Name      string
	Content   string
	CreatedBy string
	Emails    []string
}