package persist

import (
	"context"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"learngo/simple_spider/model"
	"log"
)

func ItemSaver(index string) (chan model.Item, error){
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://192.168.110.104:9200"),
	)
	if err != nil {
		return nil, err
	}

	out := make(chan model.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("ItemSaver got"+
				" item#%d: %v", itemCount, item)
			itemCount++
			err := save(client, item, index)
			if err != nil {
				log.Printf("ItemSaver error: "+
					"save item: %v, %v", item, err)
			}
		}
	}()
	return out, nil
}

func save(client *elastic.Client, item model.Item, index string) error {
	if item.Type == "" {
		return errors.New("must supply Type filed")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type)

	if item.Id != "" || item.Id != "0" {
		indexService.Id(item.Id)
	}

	_, err := indexService.
		BodyJson(item).
		Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
