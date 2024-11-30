package campaign

type Repository interface {
	Get() []Campaign
	GetById(id string) (*Campaign, error)
	Create(campaign *Campaign) (Campaign, error)
}