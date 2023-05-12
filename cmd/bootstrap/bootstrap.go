package bootstrap

import (
	"context"
	"crypto/tls"
	"database/sql"
	"fmt"
	"go-mssql/internal/infra/database/sql/area"
	"go-mssql/internal/infra/server"
	"go-mssql/internal/services"

	_ "github.com/denisenkom/go-mssqldb"
	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/denisenkom/go-mssqldb/msdsn"
)

func Start() error {
	mySqlURI := fmt.Sprintf("server=10.62.175.91;user id=dit-opd;password=oPd!_tid;port=1433;database=jovice;trustservercertificate=true")
	cfg, _, _ := msdsn.Parse(mySqlURI)
	cfg.TLSConfig.MinVersion = tls.VersionTLS10
	conn := mssql.NewConnectorConfig(cfg)
	db := sql.OpenDB(conn)

	// if err != nil {
	// 	return err
	// }

	areaRepository := area.NewAreaRepository(db)
	areaService := services.NewAreaService(areaRepository)

	server := server.New(context.Background(), "localhost", 3000, areaService)

	return server.Run()
}
