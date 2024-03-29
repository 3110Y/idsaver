package service

import (
	"context"
	utlits "github.com/3110Y/cc-utlits"
	"github.com/3110Y/idsaver/internal/application/dto"
	"github.com/3110Y/idsaver/internal/application/mapping"
	"github.com/3110Y/idsaver/internal/infrastructure/repository"
	"github.com/google/uuid"
)

type IdService struct {
	IdMapper     *mapping.IdMapper
	IdRepository *repository.IdRepository
}

func NewIdService(idMapper *mapping.IdMapper, idRepository *repository.IdRepository) *IdService {
	return &IdService{IdMapper: idMapper, IdRepository: idRepository}
}

func (i *IdService) Add(ctx context.Context, dto dto.IdDTO) (id *string, err error) {
	entity := i.IdMapper.DTOToEntity(dto)
	entity.Id = uuid.New().String()
	_, err = i.IdRepository.Add(ctx, entity)
	return
}

func (i *IdService) Item(ctx context.Context, id string) (dto *dto.IdDTO, err error) {
	entity, err := i.IdRepository.Item(ctx, id)
	dto = utlits.Pointer(i.IdMapper.EntityToDTO(*entity))
	return
}
