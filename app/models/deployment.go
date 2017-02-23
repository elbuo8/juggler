package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Deployment struct {
	ID                 string        `bson:"id" json:"id"`
	mID                bson.ObjectId `bson:"_id" json:"_id"`
	Environment        string        `bson:"name" json:"name"`
	Region             string        `bson:"region" json:"region"`
	AMIID              string        `bson:"amiId" json:"amiId"`
	Active             bool          `bson:"active" json:"active"`
	ServiceID          string        `bson:"serviceId" json:"serviceId"`
	Type               string        `bson:"type" json:"type"`
	PreviousDeployment string        `bson:"previousDeployment,omitempty" json:"previousDeployment,omitempty"`
	LCName             string        `bson:"lcName" json:"lcName"`
	ASGName            string        `bson:"asgName" json:"asgName"`
	LBName             string        `bson:"lbName" json:"lbName"`
}

/*
Register a service.
Discover Service last/first deployment.
On deploy, find active one and proceed.
*/
