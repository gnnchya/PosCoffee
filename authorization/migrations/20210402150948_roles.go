package migrations

import (
	"log"

	"github.com/uniplaces/carbon"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type roles struct {
	ID          string    `bson:"_id"`
	Name        *Lang     `bson:"name"`
	Permissions []*string `bson:"permissions"`
	CreatedAt   int64     `bson:"createdAt"`
	UpdatedAt   int64     `bson:"updatedAt"`
}

func init() {
	/*
	 * up = migrate function
	 * down = rollback function
	 */
	var coll = "roles"
	var up migrate.MigrationFunc = func(db *mongo.Database) error {
		filter := bson.M{"_id": "1100000000000000001"}
		cnt, _ := db.Collection(coll).CountDocuments(ctx, filter)
		if cnt == 0 {
			permID1 := "1000000000000000001"
			permID2 := "1000000000000000002"
			permID3 := "1000000000000000003"
			permID4 := "1000000000000000004"
			permID5 := "1000000000000000005"
			permID6 := "1000000000000000006"
			permID7 := "1000000000000000007"
			permID8 := "1000000000000000008"
			data := roles{
				ID: "1100000000000000001",
				Name: &Lang{
					Th: "ผู้ดูแลระบบ",
					En: "Administrator",
				},
				Permissions: []*string{
					&permID1,
					&permID2,
					&permID3,
					&permID4,
					&permID5,
					&permID6,
					&permID7,
					&permID8,
				},
				CreatedAt: carbon.Now().Unix(),
				UpdatedAt: carbon.Now().Unix(),
			}
			_, err := db.Collection(coll).InsertOne(ctx, data)
			if err != nil {
				return err
			}
			log.Println("- Migrate roles administrator success")
		}
		return nil
	}
	var down migrate.MigrationFunc = func(db *mongo.Database) error {
		filter := bson.M{"_id": "1100000000000000001"}
		_, err := db.Collection(coll).DeleteOne(ctx, filter)
		if err != nil {
			return err
		}
		log.Println("- Rollback roles administrator success")
		return nil
	}

	if err := migrate.Register(up, down); err != nil {
		log.Fatal(err)
	}
}
