package mongodb

import (
	"reflect"
	"strings"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *Repository) clone(origin interface{}) (clone interface{}, err error) {
	newClone := reflect.New(reflect.TypeOf(origin).Elem()).Interface()
	return newClone, copier.Copy(newClone, origin)
}

func (repo *Repository) makeFilters(filters []string) (bsonFilters bson.M) {
	bsonFilters = bson.M{}
	for _, v := range filters {
		slFilter := strings.Split(v, ":")
		key := slFilter[0]
		operations := slFilter[1]

		switch operations {
		case "ne":
			bsonFilters[key] = bson.M{"$ne": slFilter[2]}
			break
		case "like":
			bsonFilters[key] = bson.M{
				"$regex":   slFilter[2],
				"$options": "i",
			}
			break
		case "eq":
			bsonFilters[key] = slFilter[2]
			break
		case "eqInt":
			bsonFilters[key] = slFilter[2]
			break
		case "isNull":
			bsonFilters[key] = nil
			break
		case "isNotNull":
			bsonFilters[key] = bson.M{"$ne": nil}
			break
		case "id":
			oid, _ := primitive.ObjectIDFromHex(slFilter[2])
			bsonFilters[key] = oid
			break
		default:
			bsonFilters[key] = slFilter[2]
			break
		}
	}

	return bsonFilters
}

func (repo *Repository) makeSorts(sorts []string) (bsonSorts bson.M) {
	bsonSorts = bson.M{}

	for _, v := range sorts {
		slFilter := strings.Split(v, ":")
		field := slFilter[0]
		order := slFilter[1]
		bsonSorts[field] = -1
		if order == "asc" {
			bsonSorts[field] = 1
		}
	}

	return bsonSorts
}

func (repo *Repository) makePagingOpts(page int, perPage int) (opts *options.FindOptions) {
	skip := (page - 1) * perPage
	opts = options.Find()
	opts.SetSkip(int64(skip))

	if perPage > 0 {
		opts.SetLimit(int64(perPage))
	}

	return opts
}
