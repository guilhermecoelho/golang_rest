package data

import (
	"context"
	"gorest/configurations"
	"gorest/models"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func GetMoviment() models.Moviments {

	ctx := context.Background()
	rows, err := configurations.DB.QueryContext(ctx, "SELECT * FROM Moviments")
	if err != nil {
		log.Fatal("Error on consulting: ", err.Error())
	}
	defer rows.Close()

	returnMoviment := []*models.Moviment{}

	for rows.Next() {
		mov := models.Moviment{}
		rows.Scan(&mov.Id, &mov.Value)

		movJson := &models.Moviment{Id: mov.Id, Value: mov.Value}

		returnMoviment = append(returnMoviment, movJson)
	}

	return returnMoviment
}
