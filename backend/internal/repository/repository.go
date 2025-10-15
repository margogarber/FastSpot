package repository

import (
	"context"
	"time"

	"github.com/fastspot/backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repositories holds all repository instances
type Repositories struct {
	Users         *UserRepository
	Categories    *CategoryRepository
	Products      *ProductRepository
	Promotions    *PromotionRepository
	Carts         *CartRepository
	Orders        *OrderRepository
	MoodQuestions *MoodQuestionRepository
	AISessions    *AISessionRepository
}

// User Repository
type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{collection: db.Collection("users")}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Category Repository
type CategoryRepository struct {
	collection *mongo.Collection
}

func NewCategoryRepository(db *mongo.Database) *CategoryRepository {
	return &CategoryRepository{collection: db.Collection("categories")}
}

func (r *CategoryRepository) FindAll(ctx context.Context, activeOnly bool) ([]*models.Category, error) {
	filter := bson.M{}
	if activeOnly {
		filter["isActive"] = true
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var categories []*models.Category
	if err = cursor.All(ctx, &categories); err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) FindBySlug(ctx context.Context, slug string) (*models.Category, error) {
	var category models.Category
	err := r.collection.FindOne(ctx, bson.M{"slug": slug}).Decode(&category)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Category, error) {
	var category models.Category
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&category)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) Create(ctx context.Context, category *models.Category) (*models.Category, error) {
	result, err := r.collection.InsertOne(ctx, category)
	if err != nil {
		return nil, err
	}
	category.ID = result.InsertedID.(primitive.ObjectID)
	return category, nil
}

func (r *CategoryRepository) Update(ctx context.Context, id primitive.ObjectID, category *models.Category) (*models.Category, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": category}

	result := r.collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		return nil, result.Err()
	}

	var updated models.Category
	if err := result.Decode(&updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *CategoryRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

// Product Repository
type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) *ProductRepository {
	return &ProductRepository{collection: db.Collection("products")}
}

func (r *ProductRepository) FindAll(ctx context.Context, filter bson.M) ([]*models.Product, error) {
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*models.Product
	if err = cursor.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) FindBySlug(ctx context.Context, slug string) (*models.Product, error) {
	var product models.Product
	err := r.collection.FindOne(ctx, bson.M{"slug": slug, "isActive": true}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Product, error) {
	var product models.Product
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Create(ctx context.Context, product *models.Product) (*models.Product, error) {
	result, err := r.collection.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}
	product.ID = result.InsertedID.(primitive.ObjectID)
	return product, nil
}

func (r *ProductRepository) Update(ctx context.Context, id primitive.ObjectID, product *models.Product) (*models.Product, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": product}

	result := r.collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		return nil, result.Err()
	}

	var updated models.Product
	if err := result.Decode(&updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *ProductRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

// Promotion Repository
type PromotionRepository struct {
	collection *mongo.Collection
}

func NewPromotionRepository(db *mongo.Database) *PromotionRepository {
	return &PromotionRepository{collection: db.Collection("promotions")}
}

func (r *PromotionRepository) FindAll(ctx context.Context, activeOnly bool) ([]*models.Promotion, error) {
	filter := bson.M{}
	if activeOnly {
		filter["isActive"] = true
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var promotions []*models.Promotion
	if err = cursor.All(ctx, &promotions); err != nil {
		return nil, err
	}
	return promotions, nil
}

func (r *PromotionRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Promotion, error) {
	var promotion models.Promotion
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&promotion)
	if err != nil {
		return nil, err
	}
	return &promotion, nil
}

func (r *PromotionRepository) Create(ctx context.Context, promotion *models.Promotion) (*models.Promotion, error) {
	result, err := r.collection.InsertOne(ctx, promotion)
	if err != nil {
		return nil, err
	}
	promotion.ID = result.InsertedID.(primitive.ObjectID)
	return promotion, nil
}

func (r *PromotionRepository) Update(ctx context.Context, id primitive.ObjectID, promotion *models.Promotion) (*models.Promotion, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": promotion}

	result := r.collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		return nil, result.Err()
	}

	var updated models.Promotion
	if err := result.Decode(&updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *PromotionRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

// Cart Repository
type CartRepository struct {
	collection *mongo.Collection
}

func NewCartRepository(db *mongo.Database) *CartRepository {
	return &CartRepository{collection: db.Collection("carts")}
}

func (r *CartRepository) FindByUserID(ctx context.Context, userID string) (*models.Cart, error) {
	var cart models.Cart
	err := r.collection.FindOne(ctx, bson.M{"userId": userID}).Decode(&cart)
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *CartRepository) FindBySessionID(ctx context.Context, sessionID string) (*models.Cart, error) {
	var cart models.Cart
	err := r.collection.FindOne(ctx, bson.M{"sessionId": sessionID}).Decode(&cart)
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *CartRepository) Create(ctx context.Context, cart *models.Cart) error {
	result, err := r.collection.InsertOne(ctx, cart)
	if err != nil {
		return err
	}
	cart.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *CartRepository) Update(ctx context.Context, cart *models.Cart) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": cart.ID},
		bson.M{"$set": cart},
	)
	return err
}

func (r *CartRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

// Order Repository
type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) *OrderRepository {
	return &OrderRepository{collection: db.Collection("orders")}
}

func (r *OrderRepository) Create(ctx context.Context, order *models.Order) error {
	result, err := r.collection.InsertOne(ctx, order)
	if err != nil {
		return err
	}
	order.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *OrderRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Order, error) {
	var order models.Order
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) FindByUserID(ctx context.Context, userID string) ([]models.Order, error) {
	var orders []models.Order
	cursor, err := r.collection.Find(ctx, bson.M{"userId": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) FindBySessionID(ctx context.Context, sessionID string) ([]models.Order, error) {
	var orders []models.Order
	cursor, err := r.collection.Find(ctx, bson.M{"sessionId": sessionID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) Update(ctx context.Context, order *models.Order) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": order.ID},
		bson.M{"$set": order},
	)
	return err
}

func (r *OrderRepository) UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"status": status, "updatedAt": time.Now()}},
	)
	return err
}

func (r *OrderRepository) FindAll(ctx context.Context, filter bson.M) ([]models.Order, error) {
	var orders []models.Order
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

// MoodQuestion Repository
type MoodQuestionRepository struct {
	collection *mongo.Collection
}

func NewMoodQuestionRepository(db *mongo.Database) *MoodQuestionRepository {
	return &MoodQuestionRepository{collection: db.Collection("mood_questions")}
}

func (r *MoodQuestionRepository) FindAll(ctx context.Context) ([]models.MoodQuestion, error) {
	var questions []models.MoodQuestion
	cursor, err := r.collection.Find(ctx, bson.M{"isActive": true})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &questions); err != nil {
		return nil, err
	}
	return questions, nil
}

func (r *MoodQuestionRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.MoodQuestion, error) {
	var question models.MoodQuestion
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&question)
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func (r *MoodQuestionRepository) Create(ctx context.Context, question *models.MoodQuestion) (*models.MoodQuestion, error) {
	result, err := r.collection.InsertOne(ctx, question)
	if err != nil {
		return nil, err
	}
	question.ID = result.InsertedID.(primitive.ObjectID)
	return question, nil
}

func (r *MoodQuestionRepository) Update(ctx context.Context, id primitive.ObjectID, question *models.MoodQuestion) (*models.MoodQuestion, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": question}

	result := r.collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		return nil, result.Err()
	}

	var updated models.MoodQuestion
	if err := result.Decode(&updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *MoodQuestionRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

// AISession Repository
type AISessionRepository struct {
	collection *mongo.Collection
}

func NewAISessionRepository(db *mongo.Database) *AISessionRepository {
	return &AISessionRepository{collection: db.Collection("ai_sessions")}
}

func (r *AISessionRepository) Create(ctx context.Context, session *models.AISession) error {
	_, err := r.collection.InsertOne(ctx, session)
	return err
}

func (r *AISessionRepository) FindByUserID(ctx context.Context, userID string) ([]models.AISession, error) {
	var sessions []models.AISession
	cursor, err := r.collection.Find(ctx, bson.M{"userId": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &sessions); err != nil {
		return nil, err
	}
	return sessions, nil
}
