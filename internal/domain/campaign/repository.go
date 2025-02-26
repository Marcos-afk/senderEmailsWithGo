package campaign

type Repository interface {
	Get() []Campaign
	GetCampaignsToBeSent() []Campaign
	GetById(id string) (*Campaign, error)
	Create(campaign *Campaign) (*Campaign, error)
	Update(campaign *Campaign) (*Campaign, error)
	Delete(id string) error
}