package main

import (
	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)

var client *mongo.Client
var jwtKey []byte

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	ID       string `json:"id"` // Add this line
	jwt.StandardClaims
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string
	Password string
}

type Todo struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Text   string
	UserID primitive.ObjectID
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtKey = []byte(os.Getenv("JWT_SECRET"))

	// Environment variables for database connection
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")

	// MongoDB connection URI
	uri := "mongodb://127.0.0.1:27017/todo"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/signup", Signup).Methods("POST")
	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/todos", CreateTodo).Methods("POST")
	r.HandleFunc("/todos", GetTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", DeleteTodo).Methods("DELETE")

	// Add CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	log.Println("Server is running in port 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}

func Signup(w http.ResponseWriter, r *http.Request) {
	log.Println("Post request sent to /signup")
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := User{Username: creds.Username, Password: string(hashedPassword)}
	collection := client.Database("todoapp").Collection("users")
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	jsonResponse := map[string]string{"message": "User created successfully"}
	json.NewEncoder(w).Encode(jsonResponse)
}

func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Post request sent to /login")
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	collection := client.Database("todoapp").Collection("users")
	var user User
	err := collection.FindOne(context.TODO(), bson.M{"username": creds.Username}).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)) != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: creds.Username,
		ID:       user.ID.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResponse := map[string]string{"message": "User logged in successfully", "jwt": tokenString}
	json.NewEncoder(w).Encode(jsonResponse)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("Post request sent to /todos")
	tokenStr := r.Header.Get("Authorization")
	if tokenStr == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, err := primitive.ObjectIDFromHex(claims.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	todo.UserID = userID
	todosCollection := client.Database("todoapp").Collection("todos")
	_, err = todosCollection.InsertOne(context.TODO(), todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type Response struct {
	Todos []Todo `json:"todos"`
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	log.Println("Get request sent to /todos")

	collection := client.Database("todoapp").Collection("todos")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Error querying todos:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var todos []Todo
	for cursor.Next(context.TODO()) {
		var todo Todo
		cursor.Decode(&todo)
		todos = append(todos, todo)
	}
	log.Println(todos)

	response := Response{Todos: todos}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Error encoding JSON:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete request sent to /todos/{id}")

	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		log.Println("Invalid ID format:", vars["id"], err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("Parsed ID:", id)

	collection := client.Database("todoapp").Collection("todos")

	res, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Println("Error deleting todo:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if res.DeletedCount == 0 {
		log.Println("No todo found to delete with id:", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	log.Println("Todo deleted successfully with id:", id)
	w.WriteHeader(http.StatusOK)
}
