package data

import (
	"context"
	"fmt"
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

func GetMovimentById(id int) (models.Moviment, error) {

	returnMoviment := models.Moviment{}

	ctx := context.Background()
	rows, err := configurations.DB.QueryContext(ctx, "SELECT * FROM Moviments WHERE ID="+fmt.Sprintf("%d", id))
	if err != nil {
		return returnMoviment, fmt.Errorf("error on: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		mov := models.Moviment{}
		rows.Scan(&mov.Id, &mov.Value)

		returnMoviment = mov
	}

	return returnMoviment, nil
}

func PutMoviment(m *models.Moviment) (models.Moviment, error) {

	requestMoviment := *m

	query := "UPDATE Moviments SET value=" + fmt.Sprintf("%f", requestMoviment.Value) + " WHERE ID=" + fmt.Sprintf("%d", requestMoviment.Id)

	ctx := context.Background()
	rows, err := configurations.DB.QueryContext(ctx, query)
	if err != nil {
		return requestMoviment, fmt.Errorf("error on: %v", err)
	}
	defer rows.Close()

	returnMoviment, err := GetMovimentById(int(requestMoviment.Id))
	if err != nil {
		return requestMoviment, fmt.Errorf("error on: %v", err)
	}

	return returnMoviment, nil
}

func PostMoviment(m *models.Moviment) (models.Moviment, error) {

	requestMoviment := *m

	query := "INSERT INTO Moviments (value) VALUES (" + fmt.Sprintf("%f", requestMoviment.Value) + ")"

	ctx := context.Background()
	rows, err := configurations.DB.QueryContext(ctx, query)
	if err != nil {
		return requestMoviment, fmt.Errorf("error on: %v", err)
	}
	defer rows.Close()

	returnMoviment, err := GetMovimentById(int(requestMoviment.Id))
	if err != nil {
		return requestMoviment, fmt.Errorf("error on: %v", err)
	}

	return returnMoviment, nil
}

func DeleteMoviment(m *models.Moviment) error {

	requestMoviment := *m

	query := "DELETE FROM Moviments WHERE ID=" + fmt.Sprintf("%d", requestMoviment.Id)

	ctx := context.Background()
	rows, err := configurations.DB.QueryContext(ctx, query)
	if err != nil {
		return fmt.Errorf("error on: %v", err)
	}
	defer rows.Close()

	return nil
}
