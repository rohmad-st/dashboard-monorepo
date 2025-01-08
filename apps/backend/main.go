package main

import (
	"context"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

// Key represents a data model for KEYS
type Key struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

// KeyService handles business logic for keys
type KeyService interface {
	ListKeys(ctx context.Context) ([]Key, error)
	GetKey(ctx context.Context, id string) (*Key, error)
	UpdateKey(ctx context.Context, id string, key *Key) error
	DeleteKey(ctx context.Context, id string) error
}

// FirebaseKeyService implements KeyService using Firebase
type FirebaseKeyService struct {
	db *db.Client
}

func NewFirebaseKeyService(client *db.Client) *FirebaseKeyService {
	return &FirebaseKeyService{db: client}
}

func (s *FirebaseKeyService) ListKeys(ctx context.Context) ([]Key, error) {
	var keys []Key
	if err := s.db.NewRef("KEYS").Get(ctx, &keys); err != nil {
		return nil, err
	}
	return keys, nil
}

func (s *FirebaseKeyService) GetKey(ctx context.Context, id string) (*Key, error) {
	var key Key
	if err := s.db.NewRef("KEYS/"+id).Get(ctx, &key); err != nil {
		return nil, err
	}
	return &key, nil
}

func (s *FirebaseKeyService) UpdateKey(ctx context.Context, id string, key *Key) error {
	return s.db.NewRef("KEYS/"+id).Set(ctx, key)
}

func (s *FirebaseKeyService) DeleteKey(ctx context.Context, id string) error {
	return s.db.NewRef("KEYS/" + id).Delete(ctx)
}

// KeyHandler handles HTTP requests related to keys
type KeyHandler struct {
	service KeyService
}

func NewKeyHandler(service KeyService) *KeyHandler {
	return &KeyHandler{service: service}
}

func (h *KeyHandler) ListKeys(c *gin.Context) {
	ctx := c.Request.Context()
	keys, err := h.service.ListKeys(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, keys)
}

func (h *KeyHandler) GetKey(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()
	key, err := h.service.GetKey(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
		return
	}
	c.JSON(http.StatusOK, key)
}

func (h *KeyHandler) UpdateKey(c *gin.Context) {
	id := c.Param("id")
	var key Key
	if err := c.ShouldBindJSON(&key); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	ctx := c.Request.Context()
	if err := h.service.UpdateKey(ctx, id, &key); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Key updated successfully"})
}

func (h *KeyHandler) DeleteKey(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()
	if err := h.service.DeleteKey(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Key deleted successfully"})
}

func main() {
	// Load .env file if present
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Get Database URL from environment variables
	databaseURL := os.Getenv("FIREBASE_DATABASE_URL")
	if databaseURL == "" {
		log.Fatalf("FIREBASE_DATABASE_URL environment variable is not set")
	}

	// Please provide your firebase crendentials in json format
	opt := option.WithCredentialsFile("./firebase-credentials.json")
	app, err := firebase.NewApp(context.Background(), &firebase.Config{
		DatabaseURL: databaseURL,
	}, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
	}
	dbClient, err := app.Database(context.Background())
	if err != nil {
		log.Fatalf("Error initializing Firebase Database client: %v", err)
	}

	service := NewFirebaseKeyService(dbClient)
	handler := NewKeyHandler(service)

	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/keys", handler.ListKeys)
		api.GET("/keys/:id", handler.GetKey)
		api.PUT("/keys/:id", handler.UpdateKey)
		api.DELETE("/keys/:id", handler.DeleteKey)
	}

	log.Println("Server running at http://localhost:8000")
	log.Fatal(r.Run(":8000"))
}
