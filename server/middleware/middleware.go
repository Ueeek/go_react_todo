package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
    "time"
    "sort"

	"server/models"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB connection string
const connectionString = "mongodb://localhost:27017"
//const connectionString = "Connection String"

// Database Name
const dbName = "test"

// Collection name
const collName = "todolist"

// collection object/instance
var collection *mongo.Collection

// create connection with mongo db
func init() {

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created!")
}

func Signup(w http.ResponseWriter, r *http.Request){
    fmt.Println("sign up called")
}

func Login(w http.ResponseWriter, r *http.Request){
    fmt.Println("login called")
}

// GetAllTask get all the task route
func GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

    fmt.Println("get_params=>",r.URL.Query())
    queries:=r.URL.Query()
    sort_by :=queries.Get("sort_by")
    sort_order :=queries.Get("sort_order")
    show_option :=queries.Get("show_option")



	payload := getAllTask(sort_by,sort_order,show_option)
	json.NewEncoder(w).Encode(payload)
}

// CreateTask create task route
func CreateTask(w http.ResponseWriter, r *http.Request) {
    //Response header
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task models.ToDoList


    _= json.NewDecoder(r.Body).Decode(&task)
	fmt.Println("insert",task, r.Body)
    fmt.Println("time",time.Now())
    task.Date=time.Now()
    task.Status=false
	insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}

// TaskComplete update task route
func TaskComplete(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	taskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// UndoTask undo the complete task route
func UndoTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	undoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// DeleteTask delete one task route
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	deleteOneTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
	json.NewEncoder(w).Encode("Task not found")

}
func DeleteDoneTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	count := deleteDoneTask()
	json.NewEncoder(w).Encode(count)
	json.NewEncoder(w).Encode("Task not found")
}


// DeleteAllTask delete all tasks route
func DeleteAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	count := deleteAllTask()
	json.NewEncoder(w).Encode(count)
	json.NewEncoder(w).Encode("Task not found")

}

// get all task from the DB and return it
//func getAllTask(sort_by string) []primitive.M {
func getAllTask(sort_by string, sort_order string, show_option string) []models.ToDoList {


	cur, err := collection.Find(context.Background(), bson.D{{}})

    if show_option=="" || show_option=="all"{
	    cur, err = collection.Find(context.Background(), bson.D{{}})
    } else {
        fmt.Println("show option is setted:",show_option)
        fmt.Println(show_option=="done")
	    cur, err = collection.Find(context.Background(), bson.D{{"status",show_option=="done"}})
    }

	if err != nil {
		log.Fatal(err)
	}

	//var results []primitive.M
    var results []models.ToDoList
	for cur.Next(context.Background()) {
		//var result bson.M
		var result models.ToDoList
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}


    switch sort_by{
        case "date":
            if sort_order=="acs"{
                sort.SliceStable(results,func(i,j int) bool{ return results[i].Date.Before(results[j].Date)})
            } else{
                sort.SliceStable(results,func(i,j int) bool{ return results[i].Date.After(results[j].Date)})
            }

        case "task":
            if sort_order=="acs"{
                sort.SliceStable(results,func(i,j int) bool{ return results[i].Task < results[j].Task})
            } else{
                sort.SliceStable(results,func(i,j int) bool{ return results[i].Task > results[j].Task})
            }
        case "status":
            if sort_order=="acs"{
                sort.SliceStable(results,func(i,j int) bool{ return B2i(results[i].Status) < B2i(results[j].Status)})
            } else{
                sort.SliceStable(results,func(i,j int) bool{ return B2i(results[i].Status) > B2i(results[j].Status)})
            }
        default:
            //donathing

    }

	cur.Close(context.Background())
	return results
}

func B2i(b bool) int8{
    if b{
        return 1
    }
    return 0
}

// Insert one task in the DB
func insertOneTask(task models.ToDoList) {
	insertResult, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

// task complete method, update task's status to true
func taskComplete(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// task undo method, update task's status to false
func undoTask(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// delete one task from the DB, delete by ID
func deleteOneTask(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
}

// delete all the tasks from the DB
func deleteAllTask() int64 {
	d, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
	return d.DeletedCount
}

func deleteDoneTask() int64 {
	d, err := collection.DeleteMany(context.Background(), bson.D{{"status",true}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deletet Done Document", d.DeletedCount)
	return d.DeletedCount
}


