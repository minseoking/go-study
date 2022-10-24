package utils

import (
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConvertUrlParamToMap(queries string) map[string]string {
	var out map[string]string
	out = map[string]string{}
	params := strings.Split(queries, "&")
	for _, part := range params {
		for k, r := range part {
			if r == '=' {
				out[part[:k]] = part[k+1:]
			}
		}
	}

	return out
}

func ConvertUrlParamToBson(c *fiber.Ctx) bson.M {
	queries := c.Request().URI().QueryArgs().String()
	decodedQueries, _ := url.QueryUnescape(queries)
	queryMap := ConvertUrlParamToMap(decodedQueries)
	var filters []bson.M

	to := c.Query("to")
	from := c.Query("from")
	if to != "" && from != "" {

		layout := "2006-01-02"
		fromDate, _ := time.Parse(layout, from)
		toDate, _ := time.Parse(layout, to)
		t1 := time.Date(fromDate.Year(), fromDate.Month(), fromDate.Day(), 00, 01, 01, 01, fromDate.Location())
		t2 := time.Date(toDate.Year(), toDate.Month(), toDate.Day(), 23, 59, 59, 59, toDate.Location())
		filters = append(filters, bson.M{
			"createDate": bson.M{
				"$gte": primitive.NewDateTimeFromTime(t1),
				"$lte": primitive.NewDateTimeFromTime(t2),
			},
		})
	}

	for k, v := range queryMap {
		if k == "to" || k == "from" || k == "limit" || k == "skip" || v == "" {
			continue
		}
		filters = append(filters, bson.M{k: bson.M{
			"$regex": primitive.Regex{
				Pattern: v,
				Options: "i",
			},
		}})
	}

	query := bson.M{}
	if len(filters) > 0 {
		query = bson.M{
			"$and": filters,
		}
	}

	return query
}

func ConvertUrlParamToBsonOption(c *fiber.Ctx) *options.FindOptions {
	queries := c.Request().URI().QueryArgs().String()
	queryMap := ConvertUrlParamToMap(queries)

	options := options.Find()

	limit, _ := strconv.ParseInt(queryMap["limit"], 10, 64)
	skip, _ := strconv.ParseInt(queryMap["skip"], 10, 64)
	options.SetLimit(limit)
	options.SetSkip(skip * limit)

	return options
}
