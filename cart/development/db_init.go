package main

import (
	"context"
	goxid "github.com/touchtechnologies-product/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Sp struct {
	Name      		string   `bson:"name" json:"name" validate:"required"`
	ActualName 		string   `bson:"actual_name" json:"actual_name"`
	ActualLastName  string   `bson:"actual_lastname" json:"actual_lastname"`
	Gender     		string   `bson:"gender" json:"gender"`
	BirthDate  		int64    `bson:"birth_date" json:"birth_date"`
	Height     		int      `bson:"height" json:"height" validate:"gte=0"`
	SuperPower 		[]string `bson:"super_power" json:"super_power"`
	Alive      		bool     `bson:"alive" json:"alive"`
	Universe 		string	 `bson:"universe" json:"universe"`
	Movies			[]string `bson:"movies" json:"movies"`
	Enemies			[]string `bson:"enemies" json:"enemies"`
	FamilyMember	[]string `bson:"family_member" json:"family_member"`
	About			string	 `bson:"about" json:"about"`
}

var Sp_list =  []Sp{
	{"Spider-Man", "Peter", "Parker", "Male", 997401600, 178, []string{"Web-shooting"}, true, "Marvel", []string{"Spiderman", "The Avengers"}, []string{"Globlin", "Doctor Octopus"}, []string{"Richard Parker", "Mary Parker"}, "A boy who has been bitten by a spider and become superhero."},
	{"Batman", "Bruce", "Wayne", "Male", 261619200, 188, []string{"Rich"}, true, "DC", []string{"Batman", "Justice League", "The Dark Knight"}, []string{"Joker", "Superman"}, []string{"Tim Drake", "Cassandra Cain"}, "A rich man who want to be a superhero."},
	{"Superman", "Clark", "Kent", "Male", 230169600, 191, []string{"Flight", "Strength"}, true, "DC", []string{"Superman", "Man of Steel", "Justice League"}, []string{"Batman", "Justice league"}, []string{"Kara Kent", "Linda Danvers"}, "A alien who come from Krypton and become superhero in the earth."},
	{"Wonder woman", "Diana", "Prince", "Female", -908236800, 178, []string{"Agility" , "Strength"}, true, "DC", []string{"Wonder Woman", "Justice League"}, []string{"Doctor Poison"}, []string{"Donna Troy", "Miss America"}, "A girl from the mystery island and become a superhero in the real world."},
	{"Doctor Strange", "Stephen", "Strange", "Male", -1234569600, 183, []string{"Magic"}, true, "Marvel", []string{"Doctor Strange", "The Avengers"}, []string{"Baron Karl Amadeus Mordo", "Thanos"}, []string{"Donna Strange"}, "A doctor who has a lot of maditation and become a superhero."},
	{"Iron man", "Tony", "Stark", "Male", 12787200, 185, []string{"Genius", "super-suit"}, false, "Marvel", []string{"Iron Man", "The Avengers"}, []string{"Mandarin", "Doctor Doom"}, []string{"Howard Stark", "Maria Stark"}, "A rich man who become superhero in the super suit."},
	{"Black Widow", "Natasha", "Romanoff", "Female", 469929600, 170, []string{"Expert spy"}, false, "Marvel", []string{"Captain America", "The Avengers"}, []string{"Thanos", "Hulk"}, []string{"Melina Vostokoff", "Yelena Belova"}, "A Shield's spy agent who become a superhero."},
	{"Scarlet Witch", "Wanda", "Maximoff",  "Female", 192758400, 170, []string{"Energy manipulation"}, true, "Marvel", []string{"Captain America", "The Avengers"}, []string{"Iron man", "Ultron"}, []string{"Marya Maximoff", "Natalya Maximoff"}, "A witch who become superhero."},
	{"Harley Quinn", "Harleen", "Quinzel", "Female",  929836800, 168, []string{"Immunity", "Strength"}, true, "DC", []string{"Suicide Squad", "Birds of Prey"}, []string{"Batman", "Brimstone"}, []string{"Delia Quinn"}, "A doctor who become crazy and in love with Joker."},
	{"Captain America", "Steve", "Rogers", "Male", -1625097600, 188, []string{"Immunity", "Strength"}, true, "Marvel", []string{"Captain America", "The Avengers"}, []string{"Red Skull", "Thanos"}, []string{"Michael Rogers"}, "A man who become a superhero by the government experiment."},
}

func main(){
	uri := "mongodb://touch:touchja@localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	collection := client.Database("superhero").Collection("lists")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}


	defer client.Disconnect(ctx)
	for _ ,v := range Sp_list{
		if err != nil {
			log.Fatal(err)
		}
		initID := goxid.New()
		idGen := initID.Gen()
		_, err = collection.InsertOne(ctx, bson.D{
			{"_id", idGen},
			{"name", v.Name},
			{"actual_name", v.ActualName},
			{"gender", v.Gender},
			{"birth_date", v.BirthDate},
			{"height", v.Height},
			{"super_power", v.SuperPower},
			{"alive", v.Alive},
		})

		if err != nil {
			log.Fatal(err)
		}

	}
	_, err = collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.M{
				"name": 1,
			},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.M{
				"actual_name": 1,
			},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}