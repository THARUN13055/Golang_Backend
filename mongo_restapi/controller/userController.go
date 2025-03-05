package controller

import "tharun13055/utils"

var userCollection *mongo.Collection

func init(){
	client := utils.ConnectDB()
	userclient := client.Databae("curd").Collection("user")

}