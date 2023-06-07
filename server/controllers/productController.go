package controller

import (
	"context"
	"golang-chain-management/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = database.OpenCollection(database.Client, "product")

type Product struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `json:"name" validate:"required,min=2,max=100"`
	Price     float64            `json:"price" validate:"required"`
	Quantity  int                `json:"quantity" validate:"required"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

func GetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var products []Product
		cursor, err := productCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing products"})
			return
		}

		defer cancel()
		if err = cursor.All(ctx, &products); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while decoding products"})
			return
		}

		c.JSON(http.StatusOK, products)
	}
}

func GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		productID := c.Param("product_id")
		var product Product
		err := productCollection.FindOne(ctx, bson.M{"_id": productID}).Decode(&product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while finding the product"})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, product)
	}
}

func CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		// Get the user ID from the request context (assuming it was set during authentication)
		userID, _ := c.Get("uid")

		var product Product
		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		product.CreatedAt = time.Now()
		product.UpdatedAt = time.Now()

		// Insert the product into the product collection
		result, err := productCollection.InsertOne(ctx, product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating the product"})
			return
		}

		// Get the product ID from the inserted result
		productID := result.InsertedID.(primitive.ObjectID)

		// Update the user document with the new product ID
		userCollection := database.OpenCollection(database.Client, "user")
		_, err = userCollection.UpdateOne(
			ctx,
			bson.M{"_id": userID},
			bson.M{"$push": bson.M{"products": productID}},
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating the user"})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		productID := c.Param("product_id")
		var updatedProduct Product
		if err := c.BindJSON(&updatedProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updatedProduct.UpdatedAt = time.Now()

		result, err := productCollection.UpdateOne(ctx, bson.M{"_id": productID}, bson.M{"$set": updatedProduct})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating the product"})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		productID := c.Param("product_id")

		result, err := productCollection.DeleteOne(ctx, bson.M{"_id": productID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while deleting the product"})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}
