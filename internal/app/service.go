package app

import (
	"context"
	"errors"
	"slices"

	"github.com/ryvasa/test-dna/internal/dto"
	"github.com/ryvasa/test-dna/internal/model"
)

var Data []*model.Language

type ServiceInterface interface {
	GetPalindrome(ctx context.Context, text string) string
	GetAll(ctx context.Context) ([]*model.Language, error)
	GetById(ctx context.Context, id int) (*model.Language, error)
	Create(ctx context.Context, data *dto.RequestDTO) (*model.Language, error)
	Update(ctx context.Context, id int, data *dto.RequestDTO) (*model.Language, error)
	Delete(ctx context.Context, id int) error
}

type Service struct {
}

func NewService() ServiceInterface {
	return &Service{}
}

func (s *Service) GetPalindrome(ctx context.Context, text string) string {
	for i := range text[:len(text)/2] {
		if text[i] != text[len(text)-i-1] {
			return "Not Palindrome"
		}
	}
	return "Palindrome"
}

func (s *Service) GetAll(ctx context.Context) ([]*model.Language, error) {
	return Data, nil
}

func (s *Service) GetById(ctx context.Context, id int) (*model.Language, error) {
	if id < 0 || id >= len(Data) {
		return nil, errors.New("not found")
	}
	var res *model.Language
	for i := range Data {
		if i == id {
			res = Data[i]
			break
		}
	}
	if res == nil {
		return nil, errors.New("not found")
	}
	return res, nil
}

func (s *Service) Create(ctx context.Context, data *dto.RequestDTO) (*model.Language, error) {
	newData := model.Language{
		Language:       data.Language,
		Appeared:       data.Appeared,
		Created:        data.Created,
		Function:       data.Function,
		ObjectOriented: data.ObjectOriented,
		Relation: model.Relation{
			InfluencedBy: data.Relation.InfluencedBy,
			Influences:   data.Relation.Influences,
		},
	}

	Data = append(Data, &newData)
	return &newData, nil
}
func (s *Service) Update(ctx context.Context, id int, data *dto.RequestDTO) (*model.Language, error) {
	if id < 0 || id >= len(Data) {
		return nil, errors.New("not found")
	}

	Data[id].Language = data.Language
	Data[id].Appeared = data.Appeared
	Data[id].Created = data.Created
	Data[id].Function = data.Function
	Data[id].ObjectOriented = data.ObjectOriented
	Data[id].Relation.InfluencedBy = data.Relation.InfluencedBy
	Data[id].Relation.Influences = data.Relation.Influences

	return Data[id], nil
}

func (s *Service) Delete(ctx context.Context, id int) error {
	_, err := s.GetById(ctx, id)
	if err != nil {
		return err
	}
	Data = slices.Delete(Data, id, id+1)
	return nil
}
