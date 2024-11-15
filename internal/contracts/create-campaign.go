package contracts

type CreateCampaign struct {
	Name    string
	Content string
	Emails  []string
}