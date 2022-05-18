package domain

import (
	//"context"
	//"fmt"

	//"practise/User/dto"
	"context"

	"go.mongodb.org/mongo-driver/bson"

	//"go.mongodb.org/mongo-driver/bson"
	"github.com/ehs/cmd/employee/dto"
	e "github.com/ehs/pkg/err"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type DbhttpUserRepo struct {
	dbclient     *mongo.Client
	dbcollection *mongo.Collection
}

func (d DbhttpUserRepo) CreateUser(emp dto.Emp_details) (*dto.Emp_details, *e.AppError) {
	emp.AccessToken = uuid.NewString()

	_, err := d.dbcollection.InsertOne(context.TODO(), emp)

	if err != nil {
		return nil, e.NewBadRequest("error occurred while creating document from db " + err.Error())
	}
	return &emp, nil
}
func (d DbhttpUserRepo) GetUserByFilter(field, value string) (*dto.Emp_details, *e.AppError) {
	var details dto.Emp_details
	Filter := bson.M{field: value}

	err := d.dbcollection.FindOne(context.Background(), Filter).Decode(&details)
	if err == mongo.ErrNoDocuments {

		return nil, e.NewFound("no document found")
	}

	if err != nil {

		return nil, e.NewBadRequest("request failed while getting document from db with error: " + err.Error())
	}

	return &details, nil
}
func (d DbhttpUserRepo) CreateRoom() *e.AppError {
	return nil;
}
func NewDbHttppuser(dbcollection *mongo.Collection, dbclient *mongo.Client) DbhttpUserRepo {
	return DbhttpUserRepo{dbclient, dbcollection}
}
