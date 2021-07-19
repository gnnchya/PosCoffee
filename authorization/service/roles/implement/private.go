package implement

import (
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gnnchya/PosCoffee/authorize/domain"
)

func makeRoleIDFilters(roleID string) (filters []string) {
	return []string{
		fmt.Sprintf("_id:eq:%s", roleID),
		//"deletedAt:isNull",
	}
}

func makeDeletedAtIsNullFilterString() (filter string) {
	return "deletedAt:isNull"
}

func makePipelineRead(filters *[]string) (pipeline bson.A) {
	return bson.A{
		bson.M{
			"$match": makeFilters(*filters),
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "permissions",
				"localField":   "permissions",
				"foreignField": "_id",
				"as":           "permissions",
			},
		},
		bson.M{
			"$project": bson.M{
				"_id":  0,
				"id":   "$_id",
				"name": 1,
				"permissions": bson.M{
					"$map": bson.M{
						"input": "$permissions",
						"as":    "item",
						"in": bson.M{
							"id":       "$$item._id",
							"method":   "$$item.method",
							"endpoint": "$$item.endpoint",
						},
					},
				},
				"createdAt": 1,
				"updatedAt": 1,
			},
		},
	}
}

func makePipelineList(opt *domain.PageOption) (pipeline bson.A) {
	return bson.A{
		bson.M{
			"$match": makeFilters(opt.Filters),
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "permissions",
				"localField":   "permissions",
				"foreignField": "_id",
				"as":           "permissions",
			},
		},
		bson.M{
			"$project": bson.M{
				"_id":  0,
				"id":   "$_id",
				"name": 1,
				"permissions": bson.M{
					"$map": bson.M{
						"input": "$permissions",
						"as":    "item",
						"in": bson.M{
							"id":       "$$item._id",
							"method":   "$$item.method",
							"endpoint": "$$item.endpoint",
						},
					},
				},
				"createdAt": 1,
				"updatedAt": 1,
			},
		},
		bson.M{
			"$sort": makeSorts(opt.Sorts),
		},
		bson.M{
			"$skip": (opt.Page - 1) * opt.PerPage,
		},
		bson.M{
			"$limit": opt.PerPage,
		},
	}
}

func makeFilters(filters []string) (bsonFilters bson.M) {
	bsonFilters = bson.M{}
	for _, v := range filters {
		slFilter := strings.Split(v, ":")
		key := slFilter[0]
		operations := slFilter[1]

		switch operations {
		case "eq":
			bsonFilters[key] = slFilter[2]
			break
		case "isNull":
			bsonFilters[key] = nil
			break
		default:
			bsonFilters[key] = slFilter[2]
			break
		}
	}

	return bsonFilters
}

func makeSorts(sorts []string) (bsonSorts bson.M) {
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