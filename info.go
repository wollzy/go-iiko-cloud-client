package go_iiko_cloud

import "github.com/wollzy/iiko-go"

type CustomerInfoRequest struct {
	iiko.CustomerInfoRequest
	ExternalId uint `json:"external_id"`
}
type CustomerInfoResponse struct {
	iiko.CustomerInfoResponse
}

// Создать форму запроса информации о клиенте
func NewCustomerInfoRequest(externalId uint) *CustomerInfoRequest {
	return &CustomerInfoRequest{
		CustomerInfoRequest: iiko.CustomerInfoRequest{},
		ExternalId:          externalId,
	}
}

func (c *CustomerInfoRequest) WithId(id string) *CustomerInfoRequest {
	c.Id = &id
	c.Type = iiko.CustomerInfoTypePhone
	return c
}

func (c *CustomerInfoRequest) WithPhone(phone string) *CustomerInfoRequest {
	c.Phone = &phone
	c.Type = iiko.CustomerInfoTypeId
	return c
}

func (c *CustomerInfoRequest) WithCardTrack(cardTrack string) *CustomerInfoRequest {
	c.CardTrack = &cardTrack
	c.Type = iiko.CustomerInfoTypeCardTrack
	return c
}

func (c *CustomerInfoRequest) WithCardNumber(cardNumber string) *CustomerInfoRequest {
	c.CardNumber = &cardNumber
	c.Type = iiko.CustomerInfoTypeCardNumber
	return c
}

func (c *CustomerInfoRequest) WithEmail(email string) *CustomerInfoRequest {
	c.Email = &email
	c.Type = iiko.CustomerInfoTypeEmail
	return c
}
