package es

import (
	"github.com/spf13/viper"
	"log"

	"github.com/olivere/elastic"
)

var client *elastic.Client

func GetElastic() *elastic.Client {

	if client == nil {
		if err := Init(); err != nil {
			log.Println("ElasticInit:", err)
		}
	}

	return client
}

func Init() (err error) {

	elasticURL := viper.GetString("elastic_url")
	client, err = elastic.NewClient(elastic.SetURL(elasticURL), elastic.SetSniff(false))

	if err != nil {
		return err
	}

	return nil
}
