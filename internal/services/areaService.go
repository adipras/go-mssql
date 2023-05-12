package services

import "go-mssql/internal/domain"

type AreaService struct {
	areaRepository domain.AreaRepository
}

func NewAreaService(areaRepository domain.AreaRepository) AreaService {
	return AreaService{
		areaRepository: areaRepository,
	}
}

func (s AreaService) ListAreas() (*[]domain.Areas, error) {
	return s.areaRepository.ListAreas()
}
