package campaign

import (
	"errors"
	"senderEmails/internal/contracts"
	internalerrors "senderEmails/internal/internal-errors"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(createCampaign contracts.CreateCampaign) (*Campaign, error) {
	
	campaign, domainError := NewCampaign(createCampaign.Name, createCampaign.Content, createCampaign.Emails)
	if domainError != nil {
		return nil, errors.New(domainError.Error())
	}

	_, repositoryError := s.Repository.Create(campaign)

	if repositoryError != nil {
		return nil, internalerrors.ErrInternal
	}

	return campaign, nil

}


func (s *Service) Get() []Campaign {
	campaigns := s.Repository.Get()

	return campaigns
}


func (s *Service) GetById(id string) (contracts.GetCampaignByIdResponse, error) {
	campaign, err := s.Repository.GetById(id)

	if err != nil {
		return contracts.GetCampaignByIdResponse{}, errors.New("campanha não encontrada")
	}

	formatCampaignResponse := contracts.GetCampaignByIdResponse{
		ID:      campaign.ID,
		Name:    campaign.Name,
		Content: campaign.Content,
		Status:  campaign.Status,
	}

	return formatCampaignResponse, nil
}



func (s *Service) Cancel(id string) error {
	campaign, foundErr := s.Repository.GetById(id)

	if foundErr != nil {
		return errors.New("campanha não encontrada")
	}

	if campaign.Status != PendingStatus {
		return errors.New("campanha não pode ser cancelada")
	}

	campaign.Cancel()
	_, updateErr := s.Repository.Update(campaign)
	if updateErr != nil {
		return errors.New("erro ao cancelar campanha " + updateErr.Error())
	}

	return nil
}