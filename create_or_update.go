package go_iiko_cloud

import "github.com/wollzy/iiko-go"

type CreateOrUpdateRequest struct {
	iiko.CreateOrUpdateRequest
	ExternalId uint `json:"external_id"`
}
type CreateOrUpdateResponse struct {
	iiko.CreateOrUpdateResponse
}

// Создать форму запроса информации о клиенте
func NewCreateOrUpdateRequest(externalId uint) *CreateOrUpdateRequest {
	return &CreateOrUpdateRequest{
		CreateOrUpdateRequest: iiko.CreateOrUpdateRequest{},
		ExternalId:            externalId,
	}
}

func (c *CreateOrUpdateRequest) WithId(id string) *CreateOrUpdateRequest {
	c.Id = &id
	return c
}

func (c *CreateOrUpdateRequest) WithPhone(phone string) *CreateOrUpdateRequest {
	c.Phone = &phone
	return c
}

func (c *CreateOrUpdateRequest) WithName(name string) *CreateOrUpdateRequest {
	c.Name = &name
	return c
}

func (c *CreateOrUpdateRequest) WithMiddleName(middleName string) *CreateOrUpdateRequest {
	c.MiddleName = &middleName
	return c
}

func (c *CreateOrUpdateRequest) WithBirthday(birthday string) *CreateOrUpdateRequest {
	c.Birthday = &birthday
	return c
}

func (c *CreateOrUpdateRequest) WithEmail(email string) *CreateOrUpdateRequest {
	c.Email = &email
	return c
}

func (c *CreateOrUpdateRequest) WithSex(sex int) *CreateOrUpdateRequest {
	if sex >= int(iiko.SexTypeNotSpecified) && sex <= int(iiko.SexTypeFemale) {
		c.Sex = iiko.SexType(sex)
	} else {
		c.Sex = iiko.SexTypeNotSpecified
	}
	return c
}

func (c *CreateOrUpdateRequest) WithConsentStatus(consentStatus int) *CreateOrUpdateRequest {
	if consentStatus >= int(iiko.ConsentStatusUnknown) && consentStatus <= int(iiko.ConsentStatusRevoked) {
		c.ConsentStatus = iiko.ConsentStatus(consentStatus)
	} else {
		c.ConsentStatus = iiko.ConsentStatusUnknown
	}
	return c
}

func (c *CreateOrUpdateRequest) WithReferrerId(referrerId string) *CreateOrUpdateRequest {
	c.ReferrerId = &referrerId
	return c
}

func (c *CreateOrUpdateRequest) WithUserData(userData string) *CreateOrUpdateRequest {
	c.UserData = &userData
	return c
}
