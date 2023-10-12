package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func checkIndexExisted(client *elastic.Client, index string) error {
	ctx := context.TODO()
	indices := []string{index}

	existService := elastic.NewIndicesExistsService(client)
	existService.Index(indices)

	exist, err := existService.Do(ctx)

	if err != nil {
		message := fmt.Sprintf("NewIndicesExistsService.Do() %s", err.Error())
		return errors.New(message)
	} else if !exist {
		fmt.Println("nOh no! The index", index, "doesn't exist.")
		fmt.Println("Create the index, and then run the Go script again")
		if _, err = client.CreateIndex(index).Do(ctx); err != nil {
			return err
		} else {
			return nil
		}
	} else if exist {
		return nil
	} else {
		return nil
	}
}
