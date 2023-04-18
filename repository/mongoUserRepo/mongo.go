package mongoUserRepo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	cfg        inputConfig
	ctx        context.Context
	client     *mongo.Client
	collection *mongo.Collection
}

type inputConfig struct {
	Host           string `yaml:"Host"`
	Port           string `yaml:"Port"`
	DbName         string `yaml:"DbName"`
	CollectionName string `yaml:"CollectionName"`
}

func NewDB(cfg inputConfig, ctx context.Context) (*DB, error) {
	clientOptions := options.Client().ApplyURI("mongodb://" + cfg.Host + ":" + cfg.Port)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &DB{
		cfg:        cfg,
		ctx:        ctx,
		client:     client,
		collection: client.Database(cfg.DbName).Collection(cfg.CollectionName),
	}, nil

}
