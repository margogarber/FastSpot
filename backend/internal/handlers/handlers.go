package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/fastspot/backend/configs"
	"github.com/fastspot/backend/internal/models"
	"github.com/fastspot/backend/internal/repository"
	"github.com/fastspot/backend/internal/services/ai"
	"github.com/fastspot/backend/internal/services/payments"
	"github.com/fastspot/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Auth Handler
type AuthHandler struct {
	repos  *repository.Repositories
	config *configs.Config
}

func NewAuthHandler(repos *repository.Repositories, config *configs.Config) *AuthHandler {
	return &AuthHandler{repos: repos, config: config}
}

func (h *AuthHandler) CreateGuestSession(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Guest session handled by frontend"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Find user by email
	user, err := h.repos.Users.FindByEmail(ctx, req.Email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Check password
	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID.Hex(), user.Email, user.Role, h.config.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   token,
		"user": gin.H{
			"id":    user.ID.Hex(),
			"email": user.Email,
			"name":  user.Name,
			"role":  user.Role,
		},
	})
}

func (h *AuthHandler) Register(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Registration not available"})
}

func (h *AuthHandler) GetCurrentUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Logout handled by frontend"})
}

// Category Handler
type CategoryHandler struct {
	repos *repository.Repositories
}

func NewCategoryHandler(repos *repository.Repositories) *CategoryHandler {
	return &CategoryHandler{repos: repos}
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	activeOnly := c.Query("active") == "true" || c.Query("active") == ""

	categories, err := h.repos.Categories.FindAll(ctx, activeOnly)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"categories": categories}})
}

func (h *CategoryHandler) GetBySlug(c *gin.Context) {
	slug := c.Param("slug")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	category, err := h.repos.Categories.FindBySlug(ctx, slug)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": category})
}

// GetByID returns a single category by ID (Admin)
func (h *CategoryHandler) GetByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idParam := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	category, err := h.repos.Categories.FindByID(ctx, objectID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": category})
}

// Create creates a new category (Admin)
func (h *CategoryHandler) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	createdCategory, err := h.repos.Categories.Create(ctx, &category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": createdCategory})
}

// Update updates an existing category (Admin)
func (h *CategoryHandler) Update(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idParam := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var updates models.Category
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	updates.UpdatedAt = time.Now()
	updates.ID = objectID

	updatedCategory, err := h.repos.Categories.Update(ctx, objectID, &updates)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": updatedCategory})
}

// Delete deletes a category (Admin)
func (h *CategoryHandler) Delete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idParam := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	err = h.repos.Categories.Delete(ctx, objectID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Category deleted successfully"})
}

// Product Handler
type ProductHandler struct {
	repos *repository.Repositories
}

func NewProductHandler(repos *repository.Repositories) *ProductHandler {
	return &ProductHandler{repos: repos}
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"isActive": true}

	// Filter by category if provided
	if categorySlug := c.Query("category"); categorySlug != "" {
		// Find category by slug to get its ID
		category, err := h.repos.Categories.FindBySlug(ctx, categorySlug)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				// Category not found - return empty list
				c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"products": []*models.Product{}}})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch category"})
			return
		}
		filter["categoryId"] = category.ID
	}

	// Search by name if provided
	if search := c.Query("search"); search != "" {
		filter["name"] = bson.M{"$regex": search, "$options": "i"}
	}

	products, err := h.repos.Products.FindAll(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"products": products}})
}

func (h *ProductHandler) GetBySlug(c *gin.Context) {
	slug := c.Param("slug")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	product, err := h.repos.Products.FindBySlug(ctx, slug)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": product})
}

// GetByID returns a single product by ID (Admin)
func (h *ProductHandler) GetByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idParam := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := h.repos.Products.FindByID(ctx, objectID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": product})
}

// Create creates a new product (Admin)
func (h *ProductHandler) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	// Set timestamps
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	// Convert categoryId string to ObjectID if provided
	if product.CategoryID.IsZero() && c.PostForm("categoryId") != "" {
		categoryID, err := primitive.ObjectIDFromHex(c.PostForm("categoryId"))
		if err == nil {
			product.CategoryID = categoryID
		}
	}

	createdProduct, err := h.repos.Products.Create(ctx, &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": createdProduct})
}

// Update updates an existing product (Admin)
func (h *ProductHandler) Update(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idParam := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var updates models.Product
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	// Set update timestamp
	updates.UpdatedAt = time.Now()
	updates.ID = objectID

	updatedProduct, err := h.repos.Products.Update(ctx, objectID, &updates)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": updatedProduct})
}

// Delete deletes a product (Admin)
func (h *ProductHandler) Delete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idParam := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	err = h.repos.Products.Delete(ctx, objectID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Product deleted successfully"})
}

// Promotion Handler
type PromotionHandler struct {
	repos *repository.Repositories
}

func NewPromotionHandler(repos *repository.Repositories) *PromotionHandler {
	return &PromotionHandler{repos: repos}
}

func (h *PromotionHandler) GetAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	activeOnly := c.Query("active") == "true" || c.Query("active") == ""

	promotions, err := h.repos.Promotions.FindAll(ctx, activeOnly)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch promotions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"promotions": promotions}})
}

// GetByID returns a single promotion by ID
func (h *PromotionHandler) GetByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idParam := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid promotion ID"})
		return
	}

	promotion, err := h.repos.Promotions.FindByID(ctx, objectID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Promotion not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch promotion"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": promotion})
}

// Create creates a new promotion (Admin)
func (h *PromotionHandler) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var promotion models.Promotion
	if err := c.ShouldBindJSON(&promotion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	promotion.CreatedAt = time.Now()
	promotion.UpdatedAt = time.Now()

	createdPromotion, err := h.repos.Promotions.Create(ctx, &promotion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create promotion", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": createdPromotion})
}

// Update updates an existing promotion (Admin)
func (h *PromotionHandler) Update(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idParam := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid promotion ID"})
		return
	}

	var updates models.Promotion
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	updates.UpdatedAt = time.Now()
	updates.ID = objectID

	updatedPromotion, err := h.repos.Promotions.Update(ctx, objectID, &updates)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Promotion not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update promotion", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": updatedPromotion})
}

// Delete deletes a promotion (Admin)
func (h *PromotionHandler) Delete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idParam := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid promotion ID"})
		return
	}

	err = h.repos.Promotions.Delete(ctx, objectID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Promotion not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete promotion"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Promotion deleted successfully"})
}

// Cart Handler
type CartHandler struct {
	repos *repository.Repositories
}

func NewCartHandler(repos *repository.Repositories) *CartHandler {
	return &CartHandler{repos: repos}
}

// Get returns the current user's cart
func (h *CartHandler) Get(c *gin.Context) {
	ctx := c.Request.Context()

	// Get user/session ID from context (set by auth middleware)
	userID, hasUserID := c.Get("user_id")
	sessionID, hasSessionID := c.Get("session_id")

	var cart *models.Cart
	var err error

	// Try to find cart by user ID or session ID
	if hasUserID && userID != "" {
		cart, err = h.repos.Carts.FindByUserID(ctx, userID.(string))
	} else if hasSessionID && sessionID != "" {
		cart, err = h.repos.Carts.FindBySessionID(ctx, sessionID.(string))
	} else {
		// No user or session, return empty cart
		c.JSON(200, gin.H{
			"success": true,
			"data": gin.H{
				"items":    []models.CartItem{},
				"totalUSD": 0,
				"currency": "USD",
			},
		})
		return
	}

	// If cart not found, return empty cart
	if err != nil {
		c.JSON(200, gin.H{
			"success": true,
			"data": gin.H{
				"items":    []models.CartItem{},
				"totalUSD": 0,
				"currency": "USD",
			},
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    cart,
	})
}

// AddItem adds an item to the cart
func (h *CartHandler) AddItem(c *gin.Context) {
	var req struct {
		ProductID         string            `json:"productId" binding:"required"`
		Qty               int               `json:"qty"`
		ChosenIngredients []string          `json:"chosenIngredients"`
		ChosenOptions     map[string]string `json:"chosenOptions"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"success": false, "error": "Invalid request", "details": err.Error()})
		return
	}

	if req.Qty <= 0 {
		req.Qty = 1
	}

	ctx := c.Request.Context()

	// Get product details
	productOID, err := primitive.ObjectIDFromHex(req.ProductID)
	if err != nil {
		c.JSON(400, gin.H{"success": false, "error": "Invalid product ID"})
		return
	}

	product, err := h.repos.Products.FindByID(ctx, productOID)
	if err != nil {
		c.JSON(404, gin.H{"success": false, "error": "Product not found"})
		return
	}

	// Get or create cart
	userID, hasUserID := c.Get("user_id")
	sessionID, hasSessionID := c.Get("session_id")

	var cart *models.Cart

	if hasUserID && userID != "" {
		cart, err = h.repos.Carts.FindByUserID(ctx, userID.(string))
	} else if hasSessionID && sessionID != "" {
		cart, err = h.repos.Carts.FindBySessionID(ctx, sessionID.(string))
	}

	// Create new cart if not found
	if err != nil || cart == nil {
		cart = &models.Cart{
			Items:     []models.CartItem{},
			TotalUSD:  0,
			Currency:  "USD",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if hasUserID && userID != "" {
			cart.UserID = userID.(string)
		} else if hasSessionID && sessionID != "" {
			cart.SessionID = sessionID.(string)
		}

		if err := h.repos.Carts.Create(ctx, cart); err != nil {
			c.JSON(500, gin.H{"success": false, "error": "Failed to create cart"})
			return
		}
	}

	// Check if item already exists in cart
	itemIndex := -1
	for i, item := range cart.Items {
		if item.ProductID == productOID {
			itemIndex = i
			break
		}
	}

	// Calculate item price
	itemTotal := product.PriceUSD * float64(req.Qty)

	if itemIndex >= 0 {
		// Update existing item
		cart.Items[itemIndex].Qty += req.Qty
		cart.Items[itemIndex].TotalUSD = product.PriceUSD * float64(cart.Items[itemIndex].Qty)
		if req.ChosenIngredients != nil {
			cart.Items[itemIndex].ChosenIngredients = req.ChosenIngredients
		}
		if req.ChosenOptions != nil {
			cart.Items[itemIndex].ChosenOptions = req.ChosenOptions
		}
	} else {
		// Add new item
		newItem := models.CartItem{
			ProductID:         productOID,
			Name:              product.Name,
			Image:             product.Image,
			Qty:               req.Qty,
			UnitPriceUSD:      product.PriceUSD,
			TotalUSD:          itemTotal,
			ChosenIngredients: req.ChosenIngredients,
			ChosenOptions:     req.ChosenOptions,
		}
		cart.Items = append(cart.Items, newItem)
	}

	// Recalculate total
	cart.TotalUSD = 0
	for _, item := range cart.Items {
		cart.TotalUSD += item.TotalUSD
	}

	cart.UpdatedAt = time.Now()

	// Save cart
	if err := h.repos.Carts.Update(ctx, cart); err != nil {
		c.JSON(500, gin.H{"success": false, "error": "Failed to update cart"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    cart,
	})
}

// UpdateItem updates an item quantity in the cart
func (h *CartHandler) UpdateItem(c *gin.Context) {
	productID := c.Param("productId")

	var req struct {
		Qty               int               `json:"qty" binding:"required"`
		ChosenIngredients []string          `json:"chosenIngredients"`
		ChosenOptions     map[string]string `json:"chosenOptions"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"success": false, "error": "Invalid request"})
		return
	}

	ctx := c.Request.Context()
	productOID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		c.JSON(400, gin.H{"success": false, "error": "Invalid product ID"})
		return
	}

	// Get cart
	userID, hasUserID := c.Get("user_id")
	sessionID, hasSessionID := c.Get("session_id")

	var cart *models.Cart

	if hasUserID && userID != "" {
		cart, err = h.repos.Carts.FindByUserID(ctx, userID.(string))
	} else if hasSessionID && sessionID != "" {
		cart, err = h.repos.Carts.FindBySessionID(ctx, sessionID.(string))
	}

	if err != nil || cart == nil {
		c.JSON(404, gin.H{"success": false, "error": "Cart not found"})
		return
	}

	// Find and update item
	found := false
	for i, item := range cart.Items {
		if item.ProductID == productOID {
			cart.Items[i].Qty = req.Qty
			cart.Items[i].TotalUSD = item.UnitPriceUSD * float64(req.Qty)
			if req.ChosenIngredients != nil {
				cart.Items[i].ChosenIngredients = req.ChosenIngredients
			}
			if req.ChosenOptions != nil {
				cart.Items[i].ChosenOptions = req.ChosenOptions
			}
			found = true
			break
		}
	}

	if !found {
		c.JSON(404, gin.H{"success": false, "error": "Item not found in cart"})
		return
	}

	// Recalculate total
	cart.TotalUSD = 0
	for _, item := range cart.Items {
		cart.TotalUSD += item.TotalUSD
	}

	cart.UpdatedAt = time.Now()

	if err := h.repos.Carts.Update(ctx, cart); err != nil {
		c.JSON(500, gin.H{"success": false, "error": "Failed to update cart"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    cart,
	})
}

// RemoveItem removes an item from the cart
func (h *CartHandler) RemoveItem(c *gin.Context) {
	productID := c.Param("productId")

	ctx := c.Request.Context()
	productOID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		c.JSON(400, gin.H{"success": false, "error": "Invalid product ID"})
		return
	}

	// Get cart
	userID, hasUserID := c.Get("user_id")
	sessionID, hasSessionID := c.Get("session_id")

	var cart *models.Cart

	if hasUserID && userID != "" {
		cart, err = h.repos.Carts.FindByUserID(ctx, userID.(string))
	} else if hasSessionID && sessionID != "" {
		cart, err = h.repos.Carts.FindBySessionID(ctx, sessionID.(string))
	}

	if err != nil || cart == nil {
		c.JSON(404, gin.H{"success": false, "error": "Cart not found"})
		return
	}

	// Remove item
	newItems := []models.CartItem{}
	for _, item := range cart.Items {
		if item.ProductID != productOID {
			newItems = append(newItems, item)
		}
	}

	cart.Items = newItems

	// Recalculate total
	cart.TotalUSD = 0
	for _, item := range cart.Items {
		cart.TotalUSD += item.TotalUSD
	}

	cart.UpdatedAt = time.Now()

	if err := h.repos.Carts.Update(ctx, cart); err != nil {
		c.JSON(500, gin.H{"success": false, "error": "Failed to update cart"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    cart,
	})
}

// Clear removes all items from the cart
func (h *CartHandler) Clear(c *gin.Context) {
	ctx := c.Request.Context()

	// Get cart
	userID, hasUserID := c.Get("user_id")
	sessionID, hasSessionID := c.Get("session_id")

	var cart *models.Cart
	var err error

	if hasUserID && userID != "" {
		cart, err = h.repos.Carts.FindByUserID(ctx, userID.(string))
	} else if hasSessionID && sessionID != "" {
		cart, err = h.repos.Carts.FindBySessionID(ctx, sessionID.(string))
	}

	if err != nil || cart == nil {
		c.JSON(200, gin.H{"success": true, "message": "Cart already empty"})
		return
	}

	cart.Items = []models.CartItem{}
	cart.TotalUSD = 0
	cart.UpdatedAt = time.Now()

	if err := h.repos.Carts.Update(ctx, cart); err != nil {
		c.JSON(500, gin.H{"success": false, "error": "Failed to clear cart"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    cart,
	})
}

// Order Handler
type OrderHandler struct {
	repos          *repository.Repositories
	paymentService payments.PaymentProvider
}

func NewOrderHandler(repos *repository.Repositories, paymentService payments.PaymentProvider) *OrderHandler {
	return &OrderHandler{repos: repos, paymentService: paymentService}
}

// Create creates a new order from cart
func (h *OrderHandler) Create(c *gin.Context) {
	var req struct {
		PaymentMethod string `json:"paymentMethod" binding:"required"` // card, applepay, googlepay, cash
		DeliveryType  string `json:"deliveryType" binding:"required"`  // pickup, delivery
		CustomerInfo  struct {
			Name  string `json:"name" binding:"required"`
			Email string `json:"email"`
			Phone string `json:"phone" binding:"required"`
		} `json:"customerInfo" binding:"required"`
		DeliveryAddress *struct {
			Street  string `json:"street"`
			City    string `json:"city"`
			ZipCode string `json:"zipCode"`
			Notes   string `json:"notes"`
		} `json:"deliveryAddress"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"success": false, "error": "Invalid request", "details": err.Error()})
		return
	}

	ctx := c.Request.Context()

	// Validate delivery type and address
	if req.DeliveryType == "delivery" && req.DeliveryAddress == nil {
		c.JSON(400, gin.H{"success": false, "error": "Delivery address is required for delivery orders"})
		return
	}

	// Get cart
	userID, hasUserID := c.Get("user_id")
	sessionID, hasSessionID := c.Get("session_id")

	var cart *models.Cart
	var err error

	if hasUserID && userID != "" {
		cart, err = h.repos.Carts.FindByUserID(ctx, userID.(string))
	} else if hasSessionID && sessionID != "" {
		cart, err = h.repos.Carts.FindBySessionID(ctx, sessionID.(string))
	} else {
		c.JSON(400, gin.H{"success": false, "error": "No cart found"})
		return
	}

	if err != nil || cart == nil || len(cart.Items) == 0 {
		c.JSON(400, gin.H{"success": false, "error": "Cart is empty"})
		return
	}

	// Generate order number
	orderNumber := fmt.Sprintf("ORD-%d-%s", time.Now().Unix(), primitive.NewObjectID().Hex()[:6])

	// Convert cart items to order items
	orderItems := make([]models.OrderItem, len(cart.Items))
	for i, item := range cart.Items {
		orderItems[i] = models.OrderItem(item)
	}

	// Create order
	order := &models.Order{
		OrderNumber: orderNumber,
		Items:       orderItems,
		TotalUSD:    cart.TotalUSD,
		Currency:    "USD",
		Status:      "new",
		Payment: models.Payment{
			Method: req.PaymentMethod,
			Status: "pending",
		},
		Delivery: models.Delivery{
			Type: req.DeliveryType,
			Tracking: []models.TrackingEvent{
				{
					Timestamp: time.Now(),
					Status:    "Order received",
					Note:      "Your order has been received and is being processed",
				},
			},
		},
		CustomerInfo: models.CustomerInfo{
			Name:  req.CustomerInfo.Name,
			Email: req.CustomerInfo.Email,
			Phone: req.CustomerInfo.Phone,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Add user/session ID
	if hasUserID && userID != "" {
		order.UserID = userID.(string)
	} else if hasSessionID && sessionID != "" {
		order.SessionID = sessionID.(string)
	}

	// Add delivery address if needed
	if req.DeliveryType == "delivery" && req.DeliveryAddress != nil {
		order.Delivery.Address = models.DeliveryAddress{
			Street:  req.DeliveryAddress.Street,
			City:    req.DeliveryAddress.City,
			ZipCode: req.DeliveryAddress.ZipCode,
			Notes:   req.DeliveryAddress.Notes,
		}
		// Set ETA (example: 30-45 minutes)
		order.Delivery.ETA = time.Now().Add(40 * time.Minute)
	} else {
		// Pickup ETA (example: 15-20 minutes)
		order.Delivery.ETA = time.Now().Add(18 * time.Minute)
	}

	// Process payment (stub for now)
	paymentResult, err := h.paymentService.ProcessPayment(order.TotalUSD, order.Currency)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": "Payment processing failed"})
		return
	}

	order.Payment.Status = paymentResult.Status
	order.Payment.TxnID = paymentResult.TransactionID

	// Save order
	if err := h.repos.Orders.Create(ctx, order); err != nil {
		c.JSON(500, gin.H{"success": false, "error": "Failed to create order"})
		return
	}

	// Clear cart after successful order
	cart.Items = []models.CartItem{}
	cart.TotalUSD = 0
	cart.UpdatedAt = time.Now()
	_ = h.repos.Carts.Update(ctx, cart)

	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"order": order,
		},
	})
}
func (h *OrderHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	// Get user/session ID from context
	userID, hasUserID := c.Get("user_id")
	sessionID, hasSessionID := c.Get("session_id")

	var orders []models.Order
	var err error

	if hasUserID && userID != "" {
		orders, err = h.repos.Orders.FindByUserID(ctx, userID.(string))
	} else if hasSessionID && sessionID != "" {
		orders, err = h.repos.Orders.FindBySessionID(ctx, sessionID.(string))
	} else {
		c.JSON(200, gin.H{"success": true, "data": gin.H{"orders": []models.Order{}}})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": "Failed to fetch orders"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"orders": orders,
		},
	})
}

func (h *OrderHandler) GetByID(c *gin.Context) {
	orderID := c.Param("id")

	ctx := c.Request.Context()
	orderOID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		c.JSON(400, gin.H{"success": false, "error": "Invalid order ID"})
		return
	}

	order, err := h.repos.Orders.FindByID(ctx, orderOID)
	if err != nil {
		c.JSON(404, gin.H{"success": false, "error": "Order not found"})
		return
	}

	// Verify ownership
	userID, hasUserID := c.Get("user_id")
	sessionID, hasSessionID := c.Get("session_id")

	isOwner := false
	if hasUserID && userID != "" && order.UserID == userID.(string) {
		isOwner = true
	} else if hasSessionID && sessionID != "" && order.SessionID == sessionID.(string) {
		isOwner = true
	}

	if !isOwner {
		c.JSON(403, gin.H{"success": false, "error": "Access denied"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    order,
	})
}
func (h *OrderHandler) Cancel(c *gin.Context) { c.JSON(200, gin.H{"success": true}) }
func (h *OrderHandler) GetAllAdmin(c *gin.Context) {
	c.JSON(200, gin.H{"success": true, "data": gin.H{"orders": []gin.H{}}})
}
func (h *OrderHandler) UpdateStatus(c *gin.Context) { c.JSON(200, gin.H{"success": true}) }

// Mood Handler
type MoodHandler struct {
	repos         *repository.Repositories
	geminiService *ai.GeminiService
}

func NewMoodHandler(repos *repository.Repositories, geminiService *ai.GeminiService) *MoodHandler {
	return &MoodHandler{repos: repos, geminiService: geminiService}
}

// GetQuestions returns all active mood quiz questions
func (h *MoodHandler) GetQuestions(c *gin.Context) {
	questions, err := h.repos.MoodQuestions.FindAll(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": "Failed to fetch questions"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"questions": questions,
		},
	})
}

// GetRecommendations processes mood quiz answers and returns AI recommendations
func (h *MoodHandler) GetRecommendations(c *gin.Context) {
	var req struct {
		Answers []models.MoodAnswer `json:"answers" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"success": false, "error": "Invalid request", "details": err.Error()})
		return
	}

	ctx := c.Request.Context()

	// Get all active products from database with full details
	products, err := h.repos.Products.FindAll(ctx, bson.M{"isActive": true})
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": "Failed to fetch products"})
		return
	}

	// Get categories for additional context
	categories, err := h.repos.Categories.FindAll(ctx, true)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": "Failed to fetch categories"})
		return
	}

	// Get mood questions for context
	questions, err := h.repos.MoodQuestions.FindAll(ctx)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": "Failed to fetch questions"})
		return
	}

	// Convert to detailed map format for AI service with FULL product information
	productMaps := make([]map[string]interface{}, len(products))
	for i, p := range products {
		// Extract ingredient names
		ingredientNames := make([]string, len(p.Ingredients))
		for j, ing := range p.Ingredients {
			ingredientNames[j] = ing.Label
		}

		// Extract option details
		options := make([]map[string]interface{}, len(p.Options))
		for j, opt := range p.Options {
			choices := make([]string, len(opt.Choices))
			for k, choice := range opt.Choices {
				choices[k] = choice.Label
			}
			options[j] = map[string]interface{}{
				"type":    opt.Label,
				"choices": choices,
			}
		}

		// Find category name
		categoryName := "Unknown"
		for _, cat := range categories {
			if cat.ID == p.CategoryID {
				categoryName = cat.Name
				break
			}
		}

		productMaps[i] = map[string]interface{}{
			"name":        p.Name,
			"slug":        p.Slug,
			"description": p.Description,
			"category":    categoryName,
			"tags":        p.Tags,
			"priceUSD":    p.PriceUSD,
			"ingredients": ingredientNames,
			"options":     options,
		}
	}

	questionMaps := make([]map[string]interface{}, len(questions))
	for i, q := range questions {
		// Extract question options (no tags - AI analyzes text only)
		optionDetails := make([]map[string]interface{}, len(q.Options))
		for j, opt := range q.Options {
			optionDetails[j] = map[string]interface{}{
				"value": opt.Value,
				"label": opt.Label,
			}
		}

		questionMaps[i] = map[string]interface{}{
			"id":      q.ID.Hex(),
			"text":    q.Text,
			"options": optionDetails,
		}
	}

	// Convert request answers to AI service format
	aiAnswers := make([]ai.MoodAnswer, len(req.Answers))
	for i, a := range req.Answers {
		aiAnswers[i] = ai.MoodAnswer{
			QuestionID:      a.QuestionID,
			SelectedOptions: a.SelectedOptions,
		}
	}

	// Get recommendations from Gemini
	recommendation, err := h.geminiService.GetRecommendations(aiAnswers, questionMaps, productMaps)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get AI recommendations",
		})
		return
	}

	// Save AI session
	session := &models.AISession{
		Answers:         req.Answers,
		Recommendations: recommendation.ProductSlugs,
		Reasoning:       recommendation.Reasoning,
		CreatedAt:       time.Now(),
	}

	// Add user/session ID if available
	if userID, exists := c.Get("user_id"); exists {
		session.UserID = userID.(string)
	}
	if sessionID, exists := c.Get("session_id"); exists {
		session.SessionID = sessionID.(string)
	}

	_ = h.repos.AISessions.Create(ctx, session)

	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"recommendations": recommendation.ProductSlugs,
			"reasoning":       recommendation.Reasoning,
		},
	})
}

// GetAllQuestions returns all mood questions (Admin)
func (h *MoodHandler) GetAllQuestions(c *gin.Context) {
	ctx := c.Request.Context()
	questions, err := h.repos.MoodQuestions.FindAll(ctx)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": "Failed to fetch questions"})
		return
	}
	c.JSON(200, gin.H{"success": true, "data": gin.H{"questions": questions}})
}

// GetQuestionByID returns a single mood question by ID (Admin)
func (h *MoodHandler) GetQuestionByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idParam := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
		return
	}

	question, err := h.repos.MoodQuestions.FindByID(ctx, objectID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch question"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": question})
}

// CreateQuestion creates a new mood question (Admin)
func (h *MoodHandler) CreateQuestion(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var question models.MoodQuestion
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	question.CreatedAt = time.Now()
	question.UpdatedAt = time.Now()

	createdQuestion, err := h.repos.MoodQuestions.Create(ctx, &question)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create question", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": createdQuestion})
}

// UpdateQuestion updates an existing mood question (Admin)
func (h *MoodHandler) UpdateQuestion(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idParam := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
		return
	}

	var updates models.MoodQuestion
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	updates.UpdatedAt = time.Now()
	updates.ID = objectID

	updatedQuestion, err := h.repos.MoodQuestions.Update(ctx, objectID, &updates)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update question", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": updatedQuestion})
}

// DeleteQuestion deletes a mood question (Admin)
func (h *MoodHandler) DeleteQuestion(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idParam := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
		return
	}

	err = h.repos.MoodQuestions.Delete(ctx, objectID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete question"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Question deleted successfully"})
}

// Admin Handler
type AdminHandler struct {
	repos *repository.Repositories
}

func NewAdminHandler(repos *repository.Repositories) *AdminHandler {
	return &AdminHandler{repos: repos}
}

func (h *AdminHandler) GetAnalytics(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"totalOrders":     0,
			"totalRevenue":    0,
			"productsCount":   11,
			"categoriesCount": 4,
			"promotionsCount": 3,
		},
	})
}
