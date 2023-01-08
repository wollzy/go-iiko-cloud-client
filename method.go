package go_iiko_cloud

import (
	"github.com/dimonrus/gorest"
	"github.com/dimonrus/porterr"
	"net/http"
)

// Функция получения информации о клиенте
func (s Service) Info(form *CustomerInfoRequest) (*CustomerInfoResponse, porterr.IError) {
	route := "/customer/info"
	item := &CustomerInfoResponse{}
	response := gorest.JsonResponse{
		Data: item,
	}
	_, err := s.EnsureJSON(http.MethodPost, route, nil, form, &response)
	if err != nil {
		return nil, err.(*porterr.PortError)
	}
	return item, nil
}

// Функция создания/обновления информации о клиенте
func (s Service) CreateOrUpdate(form *CreateOrUpdateRequest) (*CreateOrUpdateResponse, porterr.IError) {
	route := "/customer/create_or_update"
	item := &CreateOrUpdateResponse{}
	response := gorest.JsonResponse{
		Data: item,
	}
	_, err := s.EnsureJSON(http.MethodPost, route, nil, form, &response)
	if err != nil {
		return nil, err.(*porterr.PortError)
	}
	return item, nil
}
