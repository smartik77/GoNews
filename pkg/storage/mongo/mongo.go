package mongo

import (
	"cmd/pkg/storage"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const (
	databaseName   = "testdb" // Имя учебной БД
	collectionName = "news"   // Имя коллекции в учебной БД
)

func GetAllPosts(c *mongo.Client) ([]storage.Post, error) {
	collection := c.Database(databaseName).Collection(collectionName)
	filter := bson.D{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("error getting all posts: %v", err)
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(cur, context.Background())
	var posts []storage.Post
	for cur.Next(context.Background()) {
		var post storage.Post
		err := cur.Decode(&post)
		if err != nil {
			return nil, fmt.Errorf("error decoding post: %v", err)
		}
		posts = append(posts, post)
	}
	return posts, cur.Err()
}

func AddPost(c *mongo.Client, post storage.Post) error {
	collection := c.Database(databaseName).Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		return fmt.Errorf("error inserting post: %v", err)
	}
	fmt.Println("Added new post with ID:", post.ID)
	return nil
}

func DeletePost(c *mongo.Client, post storage.Post) error {
	collection := c.Database(databaseName).Collection(collectionName)
	_, err := collection.DeleteOne(context.Background(), post)
	if err != nil {
		return fmt.Errorf("error deleting post: %v", err)
	}
	fmt.Println("Deleted post with ID:", post.ID)
	return nil
}

func UpdatePost(c *mongo.Client, post storage.Post) error {
	collection := c.Database(databaseName).Collection(collectionName)
	filter := bson.D{{"_id", post.ID}}
	update := bson.D{
		{"$set", bson.D{}},
	}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("error updating post: %v", err)
	}
	fmt.Println("Updated post with ID:", post.ID)
	return nil
}
