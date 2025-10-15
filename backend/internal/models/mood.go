package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MoodQuestion represents a question in the mood quiz
// Admin creates questions with simple text options
// AI analyzes the selected answers and recommends products on its own
type MoodQuestion struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Text      string             `bson:"text" json:"text"`
	Type      string             `bson:"type" json:"type"` // single, multiple
	Options   []QuestionOption   `bson:"options" json:"options"`
	Order     int                `bson:"order" json:"order"`
	IsActive  bool               `bson:"isActive" json:"isActive"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`
}

// QuestionOption represents a simple text option (no tags, no predefined logic)
type QuestionOption struct {
	Value string `bson:"value" json:"value"`
	Label string `bson:"label" json:"label"`
}

// AISession represents a mood quiz session
type AISession struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID          string             `bson:"userId,omitempty" json:"userId"`
	SessionID       string             `bson:"sessionId,omitempty" json:"sessionId"`
	Answers         []MoodAnswer       `bson:"answers" json:"answers"`
	Recommendations []string           `bson:"recommendations" json:"recommendations"` // product slugs
	Reasoning       string             `bson:"reasoning" json:"reasoning"`
	CreatedAt       time.Time          `bson:"createdAt" json:"createdAt"`
}

// MoodAnswer represents a user's answer (just the selected option values)
type MoodAnswer struct {
	QuestionID      string   `bson:"questionId" json:"questionId"`
	SelectedOptions []string `bson:"selectedOptions" json:"selectedOptions"` // option values
	QuestionText    string   `bson:"questionText,omitempty" json:"questionText,omitempty"`
	AnswerLabels    []string `bson:"answerLabels,omitempty" json:"answerLabels,omitempty"` // human-readable answers
}
