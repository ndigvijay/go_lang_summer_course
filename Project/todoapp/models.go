// models.go
package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty"`
    Username string             `bson:"username"`
    Password string             `bson:"password"`
}

type Todo struct {
    ID     primitive.ObjectID `bson:"_id,omitempty"`
    UserID primitive.ObjectID `bson:"user_id"`
    Title  string             `bson:"title"`
    Status string             `bson:"status"`
}
