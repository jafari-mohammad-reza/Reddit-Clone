package common

import "go.mongodb.org/mongo-driver/bson/primitive"

func StringToObjectId(id string) (primitive.ObjectID, error) {
	obi, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return obi, nil
}
