package customer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserDB struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string             `bson:"username,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty"`
	Create_date  string             `bson:"create_date,omitempty"`
}

func ConnectDB() *mongo.Collection {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("DB_MONGO_URI")
	if uri == "" {
		log.Fatal("You must set your 'DB_MONGO_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	} else {

		// fmt.Printf("Connect DB from %s\n", uri)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// coll := client.Database("service_customer").Collection("customer")

	database := client.Database("service_customer")
	Collection := database.Collection("customer")

	return Collection
}

func Adduser(username string, phone_number string) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("DB_MONGO_URI")
	if uri == "" {
		log.Fatal("You must set your 'DB_MONGO_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	} else {

		fmt.Printf("Connect DB from %s\n", uri)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	type Customer struct {
		Username     string
		Phone_number string
		Create_date  string
	}

	dt := time.Now()
	fmt.Println("Current date and time is: ", dt.String())

	coll := client.Database("service_customer").Collection("customer")
	doc := Customer{Username: username, Phone_number: phone_number, Create_date: dt.String()}
	result, err := coll.InsertOne(context.TODO(), doc)

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	if err != nil {
		panic(err)
	}
}

func Getuser(username string) {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("DB_MONGO_URI")
	if uri == "" {
		log.Fatal("You must set your 'DB_MONGO_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	} else {

		// fmt.Printf("Connect DB from %s\n", uri)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("service_customer").Collection("customer")
	// Username := "aassdd123456"

	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// fmt.Printf("No document was found with the Username %s\n", username)
		return
	}
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
	// return json.Unmarshal(jsonData)
	// return jsonData
}

func GetAll() {

	coll := ConnectDB()

	// coll := client.Database("service_customer").Collection("customer")
	// Username := "aassdd123456"

	/*
		filter := bson.D{{}}
		result, _ := client.ListDatabaseNames(context.TODO(), filter)
		fmt.Printf("%+v\n", result)
	*/

	/*
		var result bson.M
		err = coll.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&result)
		if err == mongo.ErrNoDocuments {
			// fmt.Printf("No document was found with the Username %s\n", username)
			return
		}
		if err != nil {
			panic(err)
		}*/

	// var user []UserDB
	// cursor, err := coll.Find(ctx, bson.M{"create_date": bson.D{{"$gt", "25"}}})

	findOptions := options.Find()
	cursor, err := coll.Find(context.TODO(), nil, findOptions)

	if err != nil {
		fmt.Println(err) // prints 'document is nil'
	}

	// if err = cursor.All(ctx, &user); err != nil {
	// 	panic(err)
	// }
	fmt.Println(cursor)

	// jsonData, err := json.MarshalIndent(result, "", "    ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", jsonData)

}

// Find multiple documents
func FindRecords() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	//Get database settings from env file
	//dbUser := os.Getenv("db_username")
	//dbPass := os.Getenv("db_pass")
	dbName := os.Getenv("DB_MONGO_NAME")
	docCollection := "retailMembers"

	uri := os.Getenv("MONGODB_URI")

	// dbHost := os.Getenv("MONGODB_URI")
	// dbPort := os.Getenv("db_port")
	// dbEngine := os.Getenv("db_type")

	//set client options
	clientOptions := options.Client().ApplyURI(uri)
	//connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	//check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Connected to " + dbEngine)
	db := client.Database(dbName).Collection(docCollection)

	//find records
	//pass these options to the Find method
	findOptions := options.Find()
	//Set the limit of the number of record to find
	findOptions.SetLimit(5)
	//Define an array in which you can store the decoded documents
	var results []UserDB

	//Passing the bson.D{{}} as the filter matches  documents in the collection
	cur, err := db.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	//Finding multiple documents returns a cursor
	//Iterate through the cursor allows us to decode documents one at a time

	for cur.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem UserDB
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("documents: %s\n", elem)

		results = append(results, elem)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	//Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents: %+v\n", results)

}
