package comment

import (
	"github.com/OhYee/goutils/time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TypeDB comment type in database
type TypeDB struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Email   string             `json:"email" bson:"email"`
	Avatar  string             `json:"avatar" bson:"avatar"`
	Time    int64              `json:"time" bson:"time"`
	Raw     string             `json:"raw" bson:"raw"`
	Content string             `json:"content" bson:"content"`
	Reply   primitive.ObjectID `json:"reply" bson:"reply"`
	URL     string             `json:"url" bson:"url"`
	Ad      bool               `json:"ad" bson:"ad"`
	Show    bool               `json:"show" bson:"show"`
	Recv    bool               `json:"recv" bson:"recv"`
}

// ToComment transfer CommentDB to comment
func (cm *TypeDB) ToComment() *Type {
	return &Type{
		ID:       cm.ID.Hex(),
		Email:    cm.Email,
		Avatar:   cm.Avatar,
		Time:     time.ToString(cm.Time),
		Content:  cm.Content,
		Reply:    cm.Reply,
		Children: []*Type{},
		Ad:       cm.Ad,
		Show:     cm.Show,
		Recv:     cm.Recv,
	}
}

// Type type
type Type struct {
	ID       string             `json:"id" bson:"_id"`
	Email    string             `json:"email" bson:"email"`
	Avatar   string             `json:"avatar" bson:"avatar"`
	Time     string             `json:"time" bson:"time"`
	Content  string             `json:"content" bson:"content"`
	Reply    primitive.ObjectID `json:"reply" bson:"reply"`
	Children []*Type            `json:"children" bson:"children"`
	Ad       bool               `json:"ad" bson:"ad"`
	Show     bool               `json:"show" bson:"show"`
	Recv     bool               `json:"recv" bson:"recv"`
}

// AdminDB comment database Type
type AdminDB struct {
	TypeDB       `bson:",inline"`
	ReplyComment TypeDB `json:"reply_comment" bson:"reply_comment"`
	Title        string `json:"title" bson:"title"`
}

// ToAdmin transfer AdminDB to Admin
func (comment AdminDB) ToAdmin() *Admin {
	return &Admin{
		Type:         *comment.TypeDB.ToComment(),
		ReplyComment: *comment.ReplyComment.ToComment(),
		Title:        comment.Title,
		URL:          comment.TypeDB.URL,
	}
}

// Admin comment Type
type Admin struct {
	Type         `bson:",inline"`
	ReplyComment Type   `json:"reply_comment" bson:"reply_comment"`
	URL          string `json:"url" bson:"url"`
	Title        string `json:"title" bson:"title"`
}

// Info base info of comment
type Info struct {
	Title string `json:"title" bson:"title"`
	Recv  bool   `json:"recv" bson:"recv"`
	Email string `json:"email" bson:"email"`
}
