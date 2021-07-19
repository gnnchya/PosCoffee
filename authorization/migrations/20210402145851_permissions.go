package migrations

import (
	"log"

	"github.com/uniplaces/carbon"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type permissions struct {
	ID        string `bson:"_id"`
	Method    string `bson:"method"`
	Endpoint  string `bson:"endpoint"`
	CreatedAt int64  `bson:"createdAt"`
	UpdatedAt int64  `bson:"updatedAt"`
}

func init() {
	/*
	 * up = migrate function
	 * down = rollback function
	 */
	var coll = "permissions"

	migrateData := [8][3]string{
		{"1000000000000000001", "GET", "/permissions"},
		{"1000000000000000002", "POST", "/permissions"},
		{"1000000000000000003", "PUT", "/permissions"},
		{"1000000000000000004", "DELETE", "/permissions"},
		{"1000000000000000005", "GET", "/roles"},
		{"1000000000000000006", "POST", "/roles"},
		{"1000000000000000007", "PUT", "/roles"},
		{"1000000000000000008", "DELETE", "/roles"},
	}

	var up migrate.MigrationFunc = func(db *mongo.Database) error {
		for _, data := range migrateData {
			filter := bson.M{"_id": data[0]}
			cnt, _ := db.Collection(coll).CountDocuments(ctx, filter)
			if cnt == 0 {
				permissionData := permissions{
					ID:        data[0],
					Method:    data[1],
					Endpoint:  data[2],
					CreatedAt: carbon.Now().Unix(),
					UpdatedAt: carbon.Now().Unix(),
				}
				_, err := db.Collection(coll).InsertOne(ctx, permissionData)
				if err != nil {
					return err
				}
				log.Println("- Migrate permission", data[1], data[2], "success")
			}
		}
		return nil
	}
	var down migrate.MigrationFunc = func(db *mongo.Database) error {
		for _, data := range migrateData {
			filter := bson.M{"_id": data[0]}
			_, err := db.Collection(coll).DeleteOne(ctx, filter)
			if err != nil {
				return err
			}
			log.Println("- Rollback permission", data[1], data[2], "success")
		}
		return nil
	}

	if err := migrate.Register(up, down); err != nil {
		log.Fatal(err)
	}
}
