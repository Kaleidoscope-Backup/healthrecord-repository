package service

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/mongo-lib/mserver"
)

/*==========================================================================================
Comment service
==========================================================================================*/

//CommentService is for creating comment
type CommentService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

//NewCommentService creates a new Comment service that has all calls to the database, queries and mutations via the Data Access Layer
func NewCommentService(dal mserver.DataAccessLayer, log *logging.Logger) *CommentService {
	return &CommentService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

//FindByID ..
func (u *CommentService) FindByID(id string) (*model.Comment, error) {

	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching comment from Mongo
	c, err := u.dal.Get(id, &model.Comment{})
	if err != nil {
		return nil, err
	}

	var comment *model.Comment
	bsonBytes, _ := bson.Marshal(c)
	bson.Unmarshal(bsonBytes, &comment)

	return comment, nil
}

//FindByExternalID ...
func (u *CommentService) FindByExternalID(externalID string) (*model.Comment, error) {
	if externalID == "" {
		return nil, errors.New("Missing parameter")
	}

	var params map[string]string
	params = map[string]string{}

	params["externalID"] = externalID

	//find the matching product id Record (if any) from Mongo
	ccArr, err := FindRecords(&params, &model.Comment{}, u.dal)
	if err != nil {
		return nil, err
	}

	for _, pr := range ccArr {
		var comment *model.Comment
		bsonBytes, _ := bson.Marshal(pr)
		bson.Unmarshal(bsonBytes, &comment)
		return comment, nil
	}

	return nil, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

//CreateCommentOnComment will create a new comment in Mongo using the Data Access Layer
func (u *CommentService) CreateCommentOnComment(externalID string, comment *model.Comment) (*model.Comment, error) {

	// Get the comment to be updated
	commentOnComment, err := u.FindByExternalID(externalID)
	if err != nil {
		return nil, err
	}

	// Update the comment
	comments := []model.Comment{}
	if commentOnComment.Comments != nil {
		comments = *commentOnComment.Comments
	}

	comments = append(comments, *comment)
	commentOnComment.Comments = &comments
	result, err := u.dal.Put(commentOnComment.Id, commentOnComment)
	if err != nil || result != true {
		return nil, err
	}

	return commentOnComment, nil
}

//CreateComment will create a new comment in Mongo using the Data Access Layer
func (u *CommentService) CreateComment(comment *model.Comment) (*model.Comment, error) {
	id, err := u.dal.Post(comment)
	if err != nil {
		return nil, err
	}

	comment.Id = id
	return comment, nil
}
