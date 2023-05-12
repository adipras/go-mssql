package area

import (
	"database/sql"
	"go-mssql/internal/domain"
)

type AreaRepository struct {
	db *sql.DB
}

func NewAreaRepository(db *sql.DB) *AreaRepository {
	return &AreaRepository{
		db: db,
	}
}

// func (r *AreaRepository) CreateArea(area domain.IArea) error {

// }
func (r *AreaRepository) ListAreas() (*[]domain.Areas, error) {
	var areas []domain.Areas
	rows, err := r.db.Query(
		`select * from area where ar_oc is not null and ar_latitude is not null and ar_longitude is not null`,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var area domain.Areas
		err2 := rows.Scan(

			&area.AreaID,
			&area.AreaCode,
			&area.AreaName,
			&area.AreaAW,
			&area.AreaOC,
			&area.AreaLat,
			&area.AreaLong,
		)
		if err2 != nil {
			return nil, err2
		}
		areas = append(areas, area)
	}

	return &areas, nil
}
