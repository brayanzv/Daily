package bd

import (
	"context"
	"time"

	"github.com/brayanzv/Daily/models"
)

func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("Daily")
	col := db.Collection("relation")

	_, err := col.DeleteOne(ctx, t)

	if err != nil {
		return false, err
	}
	return true, nil
}
