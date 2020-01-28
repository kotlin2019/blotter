package mongo

import (
	"context"
	"github.com/OhYee/blotter/output"
	"github.com/OhYee/rainbow/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientOptions = options.Client().ApplyURI("mongodb://127.0.0.1:27017")

type countResult struct {
	Count int `bson:"count"`
}

func Find(databaseName string, collectionName string, filter interface{},
	opt *options.FindOptions, res interface{}) (err error) {
	defer func() {
		if err != nil {
			err = errors.NewErr(err)
		}
	}()

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return
	}
	defer client.Disconnect(context.TODO())

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return
	}

	collection := client.Database(databaseName).Collection(collectionName)
	cur, err := collection.Find(context.TODO(), filter, opt)
	if err != nil {
		return
	}
	defer cur.Close(context.TODO())

	err = cur.All(context.TODO(), res)
	return
}

func Aggregate(databaseName string, collectionName string, pipeline interface{},
	opt *options.AggregateOptions, res interface{}) (total int, err error) {
	defer func() {
		if err != nil {
			err = errors.NewErr(err)
		}
	}()

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return
	}
	defer client.Disconnect(context.TODO())

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return
	}

	collection := client.Database(databaseName).Collection(collectionName)
	cur, err := collection.Aggregate(context.TODO(), pipeline, opt)
	if err != nil {
		return
	}
	defer cur.Close(context.TODO())
	if err = cur.All(context.TODO(), res); err != nil {
		return
	}

	count := countResult{}
	countPipeline, err := pipelineTruncated(pipeline)
	countPipeline = append(countPipeline, bson.M{"$count": "count"})
	if err != nil {
		return
	}
	output.Debug("%+v", countPipeline)
	countCur, err := collection.Aggregate(context.TODO(), countPipeline, opt)
	if err != nil {
		return
	}
	defer countCur.Close(context.TODO())
	if countCur.Next(context.TODO()) {
		if err = countCur.Decode(&count); err != nil {
			return
		}
	}
	total = count.Count

	return
}

func bsonFormat(b interface{}) (bb []bson.M, err error) {
	switch b.(type) {
	case bson.D:
		bb = []bson.M{b.(bson.D).Map()}
	case []bson.E:
		bb = []bson.M{bson.D(b.([]bson.E)).Map()}
	case bson.E:
		bb = []bson.M{bson.D([]bson.E{b.(bson.E)}).Map()}
	case bson.M:
		bb = []bson.M{b.(bson.M)}
	case map[string]interface{}:
		bb = []bson.M{bson.M(b.(map[string]interface{}))}
	case []bson.M:
		bb = b.([]bson.M)
	case []map[string]interface{}:
		m := b.([]map[string]interface{})
		bb = make([]bson.M, len(m))
		for idx, data := range m {
			bb[idx] = bson.M(data)
		}
	default:
		err = errors.New("Can format bson: %+v", b)
		bb = []bson.M{}
	}
	return
}

func pipelineTruncated(pipeline interface{}) (res []bson.M, err error) {
	m, err := bsonFormat(pipeline)
	end := -1
	for i := len(m) - 1; i >= 0; i-- {
		output.Debug("%+v", m[i])
		if _, exist := m[i]["$limit"]; exist {
			continue
		}
		if _, exist := m[i]["$skip"]; exist {
			continue
		}
		end = i
		break
	}
	res = m[0 : end+1]
	return
}