package campaign

import (
	"errors"
	"senderEmails/internal/contracts"
	"senderEmails/internal/infrastructure/providers"
	internalerrors "senderEmails/internal/internal-errors"
)

type Service interface {
	Get () []Campaign
	GetById(id string) (*contracts.GetCampaignByIdResponse, error)
	Create(campaign contracts.CreateCampaign) (*Campaign, error)
	Cancel(id string) error
	Delete(id string) error
	Start(id string) error
	SendMailAndUpdateStatus(SendMailToCampaign *Campaign) (bool, error)
}

type ServiceImp struct {
	Repository Repository
	MailProvider providers.MailProvider
}

func (s *ServiceImp) Create(createCampaign contracts.CreateCampaign) (*Campaign, error) {
	
	campaign, domainError := NewCampaign(createCampaign.Name, 
		createCampaign.Content,
		createCampaign.CreatedBy, 
		createCampaign.Emails)

	if domainError != nil {
		return nil, errors.New(domainError.Error())
	}

	_, repositoryError := s.Repository.Create(campaign)

	if repositoryError != nil {
		return nil, internalerrors.ErrInternal
	}

	return campaign, nil

}


func (s *ServiceImp) Get() []Campaign {
	campaigns := s.Repository.Get()

	return campaigns
}


func (s *ServiceImp) GetById(id string) (*contracts.GetCampaignByIdResponse, error) {
	campaign, err := s.Repository.GetById(id)

	if err != nil {
		return nil, errors.New("campanha não encontrada")
	}

	formatCampaignResponse := contracts.GetCampaignByIdResponse{
		ID:      campaign.ID,
		Name:    campaign.Name,
		Content: campaign.Content,
		Status:  campaign.Status,
	}

	return &formatCampaignResponse, nil
}


func (s *ServiceImp) Cancel(id string) error {
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

func (s *ServiceImp) Delete(id string) error {
	_, foundErr := s.Repository.GetById(id)

	if foundErr != nil {
		return errors.New("campanha não encontrada")
	}

	deleteErr := s.Repository.Delete(id)
	if deleteErr != nil {
		return errors.New("erro ao deletar campanha " + deleteErr.Error())
	}

	return nil
}

func (s *ServiceImp) SendMailAndUpdateStatus(SendMailToCampaign *Campaign) (bool, error) {
	contacts := make([]string, len(SendMailToCampaign.Contacts))
	for i, contact := range SendMailToCampaign.Contacts {
    contacts[i] = contact.Email
	}

	sendEmailErr := s.MailProvider.SendMail(contracts.SendMailRequest{
		To: contacts,
		Subject: SendMailToCampaign.Name,
		Message: SendMailToCampaign.Content,
	})
		
	if sendEmailErr != nil {
			SendMailToCampaign.Failed()
	}	else{ 
			SendMailToCampaign.Sent()
	}

  _, updateErr:=	s.Repository.Update(SendMailToCampaign)
	if updateErr != nil {
		return false, errors.New("erro ao enviar campanha " + updateErr.Error())
	}

	return sendEmailErr == nil, nil
}

func (s *ServiceImp) Start(id string) error {
	campaign, foundErr := s.Repository.GetById(id)

	if foundErr != nil {
		return errors.New("campanha não encontrada")
	}

	if campaign.Status != PendingStatus {
		return errors.New("campanha não pode ser iniciada")
	}

	campaign.Started()

	_, updateErr := s.Repository.Update(campaign)
	if updateErr != nil {
		return errors.New("erro ao enviar campanha " + updateErr.Error())
	}


	go s.SendMailAndUpdateStatus(campaign)
	
	return nil
}