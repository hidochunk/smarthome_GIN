package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"smarthome_GIN/database"
)

type Device struct {
	ID                primitive.ObjectID `bson:"_id"`
	DeviceType        string             `bson:"device_type,omitempty"`
	DeviceName        string             `bson:"device_name,omitempty"`
	CommunicationType string             `bson:"communication_type,omitempty"`
	Topic             string             `bson:"topic,omitempty"`
	Functions         []Function         `bson:"functions,omitempty"`
}

func GetDeviceByDeviceType(DeviceType string) []*Device {
	var devices []*Device
	DB := database.DB
	coll := DB.Database("smarthome").Collection("device")
	cursor, err := coll.Find(context.TODO(), bson.D{{"device_type", DeviceType}})
	if err != nil {
		panic(err)
	}

	for cursor.Next(context.TODO()) {
		var temp Device
		err = cursor.Decode(&temp)
		if err != nil {
			panic(err)
		}
		devices = append(devices, &temp)
	}

	return devices
}

func GetDeviceByCommunicationType(CommunicationType string) []*Device {
	var devices []*Device
	DB := database.DB
	coll := DB.Database("smarthome").Collection("device")
	cursor, err := coll.Find(context.TODO(), bson.D{{"communication_type", CommunicationType}})
	if err != nil {
		panic(err)
	}

	for cursor.Next(context.TODO()) {
		var temp Device
		err = cursor.Decode(&temp)
		if err != nil {
			panic(err)
		}
		devices = append(devices, &temp)
	}

	return devices
}
