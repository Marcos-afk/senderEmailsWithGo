package campaign

type Repository interface {
	Get() []Campaign
	Create(campaign *Campaign) (Campaign, error)
}