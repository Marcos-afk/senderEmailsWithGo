package campaign

import (
	"errors"
	"senderEmails/internal/contracts"
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
		return nil, errors.New(repositoryError.Error())
	}

	return campaign, nil

}