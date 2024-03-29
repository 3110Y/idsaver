package mapping

import (
	utlits "github.com/3110Y/cc-utlits"
	"github.com/3110Y/idsaver/internal/application/dto"
	"github.com/3110Y/idsaver/internal/domain/entity"
)

type IdMapper struct {
}

func NewIdMapper() *IdMapper {
	return &IdMapper{}
}

func (i IdMapper) EntityToDTO(entity entity.IDEntity) dto.IdDTO {
	return dto.IdDTO{
		Id:       &entity.Id,
		Uid:      entity.Uid,
		Len:      entity.Len,
		Allow:    entity.Allow,
		CreateAt: &entity.CreateAt,
		UpdateAt: &entity.UpdateAt,
	}
}

func (i IdMapper) DTOToEntity(dto dto.IdDTO) entity.IDEntity {
	return entity.IDEntity{
		Id:       utlits.ValueOrDefault(dto.Id),
		Uid:      dto.Uid,
		Len:      dto.Len,
		Allow:    dto.Allow,
		CreateAt: utlits.ValueOrDefault(dto.CreateAt),
		UpdateAt: utlits.ValueOrDefault(dto.UpdateAt),
	}
}
