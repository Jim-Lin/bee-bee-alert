package db

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/olivere/elastic"

	"github.com/Jim-Lin/bee-bee-alert/backend/model"
	"github.com/Jim-Lin/bee-bee-alert/backend/util"
)

const (
	FIELD    = "name"
	ES_INDEX = "bee"
	ES_TYPE  = "prod"
)

func MostLike(prod model.Prod) []byte {
	client, err := elastic.NewClient(
		elastic.SetURL(util.GetConfig().EsUrl),
		elastic.SetSniff(false))
	util.CheckError(err)
	defer client.Stop()

	ctx := context.Background()
	mltq := elastic.NewMoreLikeThisQuery().
		LikeText(prod.Name).
		Field(FIELD).
		MaxQueryTerms(4).
		MinTermFreq(1).
		MinDocFreq(1).
		MaxDocFreq(1).
		MinimumShouldMatch("3")
	res, err := client.Search().
		Index(ES_INDEX).
		Type(ES_TYPE).
		Query(mltq).
		Pretty(true).
		Do(ctx)
	util.CheckError(err)

	log.Printf("Query took %d milliseconds\n", res.TookInMillis)
	if res.Hits.TotalHits > 0 {
		log.Printf("Found a total of %d prods\n", res.Hits.TotalHits)

		if *res.Hits.MaxScore > 1 {
			var likeProd model.Prod
			err := json.Unmarshal(*res.Hits.Hits[0].Source, &likeProd)
			util.CheckError(err)

			if likeProd.Price > prod.Price {
				log.Print("Too expensive\n")

				go (&util.MailTemplate{
					Subject: "[Warning] Too expensive!",
					Msg:     "honestbee: \r\n" + likeProd.Name + "\r\n$" + strconv.Itoa(likeProd.Price) + "\r\n" + likeProd.Url + "\r\n======\r\n" + prod.Name + "\r\n$" + strconv.Itoa(prod.Price) + "\r\n" + prod.Url + "\r\n",
				}).GetMail().Notify()
			} else {
				jsonBytes, err := json.Marshal(likeProd)
				util.CheckError(err)

				return jsonBytes
			}
		}
	} else {
		log.Print("Not found\n")
	}

	jsonBytes, err := json.Marshal(model.Prod{})
	util.CheckError(err)

	return jsonBytes
}
