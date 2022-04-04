package middleware

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/OksidGen/todoapp/internal/model"
)

func GetAllTasks(c *gin.Context) {
	cursor, err := taskCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var tasks []model.Task
	for cursor.Next(context.TODO()) {
		var task model.Task
		err := cursor.Decode(&task)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	cursor.Close(context.Background())
	// c.JSON(http.StatusOK, gin.H{"All": tasks})
	c.HTML(http.StatusOK, "index.html", gin.H{"tasks": tasks})
}

func CreateTask(c *gin.Context) {
	text := c.PostForm("text")
	task := model.Task{
		ID:     primitive.NewObjectID(),
		Text:   text,
		Status: false,
	}
	_, err := taskCollection.InsertOne(context.TODO(), task)
	if err != nil {
		log.Fatal("Could not create Task: ", err)
	}
	// c.JSON(http.StatusFound, gin.H{"result": result.InsertedID.(primitive.ObjectID).Hex()})
	c.Redirect(http.StatusFound, "/")
}

func UpdateTask(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.PostForm("id"))
	text := c.PostForm("text")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"text": text}}
	_, err := taskCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	// c.JSON(http.StatusFound, gin.H{"Modified count: ": result.MatchedCount})
	c.Redirect(http.StatusFound, "/")
}

func UpdateStatusTask(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.PostForm("id"))
	status, _ := strconv.ParseBool(c.PostForm("status"))
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": !status}}
	_, err := taskCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	// c.JSON(http.StatusFound, gin.H{"Modified count: ": result.MatchedCount})
	c.Redirect(http.StatusFound, "/")

}

func DeleteTask(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.PostForm("id"))
	filter := bson.M{"_id": id}
	_, err := taskCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	// c.JSON(http.StatusFound, gin.H{"Deleted count": result.DeletedCount})
	c.Redirect(http.StatusFound, "/")
}

func DeleteAllTask(c *gin.Context) {
	_, err := taskCollection.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	// c.JSON(http.StatusFound, gin.H{"Deleted count": result.DeletedCount})
	c.Redirect(http.StatusFound, "/")
}
