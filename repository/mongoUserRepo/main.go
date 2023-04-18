package mongoUserRepo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"modularis/models"
)

type Repository struct {
	mgo *DB
}

func NewRepository(cfg inputConfig, ctx context.Context) (*Repository, error) {
	db, err := NewDB(cfg, ctx)
	if err != nil {
		return nil, err
	}
	return &Repository{
		mgo: db,
	}, nil
}

func (r Repository) CreateUser(user *models.User) error {
	bytes, err := bson.Marshal(user)
	if err != nil {
		return err
	}
	_, err = r.mgo.collection.InsertOne(r.mgo.ctx, bytes)
	return err
}

func (r Repository) RemoveUSer(login string) error {
	_, err := r.mgo.collection.DeleteOne(r.mgo.ctx, bson.M{"_id": login})
	return err
}

func (r Repository) GetUserInfo(login string) (*models.User, error) {
	res := r.mgo.collection.FindOne(r.mgo.ctx, bson.M{"_id": login})
	var usr models.User
	err := res.Decode(&usr)
	return &usr, err
}
