package campaign

type Repository interface {
	Create(campaign *Campaign) (Campaign, error)
}