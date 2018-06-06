package models

import "gopkg.in/mgo.v2/bson"

// Dialogs type model
type Dialogs struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Dialog    string        `json:"dialog" bson:"dialog"`
	Tags      string        `json:"tags" bson:"tags"`
	Lang      string        `json:"lang" bson:"lang"`
	User      string        `json:"user" bson:"user"`
	Status    string        `json:"status" bson:"status"`
	Timestamp string        `json:"timestamp" bson:"timestamp"`
}

// Pics type model
type Pics struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	DialogID  string        `json:"dialog_id" bson:"dialog_id"`
	Pic       string        `json:"pic" bson:"pic"`
	User      string        `json:"user" bson:"user"`
	Status    string        `json:"status" bson:"status"`
	Timestamp string        `json:"timestamp" bson:"timestamp"`
}

// Videos type model
type Videos struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	DialogID  string        `json:"dialog_id" bson:"dialog_id"`
	Video     string        `json:"video" bson:"video"`
	User      string        `json:"user" bson:"user"`
	Status    string        `json:"status" bson:"status"`
	Timestamp string        `json:"timestamp" bson:"timestamp"`
}

// Voices type model
type Voices struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	DialogID  string        `json:"dialog_id" bson:"dialog_id"`
	Voice     string        `json:"voice" bson:"voice"`
	User      string        `json:"user" bson:"user"`
	Status    string        `json:"status" bson:"status"`
	Timestamp string        `json:"timestamp" bson:"timestamp"`
}

// Favorites type model
type Favorites struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	DialogType   string        `json:"dialog_type" bson:"dialog_type"`
	DialogTypeID string        `json:"dialog_type_id" bson:"dialog_type_id"`
	User         string        `json:"user" bson:"user"`
	Status       string        `json:"status" bson:"status"`
	Timestamp    string        `json:"timestamp" bson:"timestamp"`
}

// Reactions type model
type Reactions struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	DialogType   string        `json:"dialog_type" bson:"dialog_type"`
	DialogTypeID string        `json:"dialog_type_id" bson:"dialog_type_id"`
	ReactionType string        `json:"reaction_type" bson:"reaction_type"`
	Reaction     string        `json:"reaction" bson:"reaction"`
	User         string        `json:"user" bson:"user"`
	Status       string        `json:"status" bson:"status"`
	Timestamp    string        `json:"timestamp" bson:"timestamp"`
}
