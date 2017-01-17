package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lycam-dev/lycamplus-go-sdk/lycamplus"
)

//
// constant
//
const (
	// AppKey .
	AppKey = "488ITUGN1G"

	// AppSecret .
	AppSecret = "z1oyx55jNQEXeRUu1iltfINZegWuGx"

	// MasterSecret .
	MasterSecret = "9O1MZJ5UJwnuZky3tUBiZFPAlDJNs2"
)

//
// global variable
//
var aLycamPlus *lycamplus.LycamPlus

//
// main function.
//
func main() {

	// create LycamPlus instance
	aLycamPlus = lycamplus.NewLycamPlus(AppKey, AppSecret, MasterSecret)

	// http router
	router := httprouter.New()

	// User
	router.POST("/user", CreateUser)
	router.POST("/assume", UserAssume)

	// Stream
	router.POST("/stream", CreateStream)
	router.PUT("/stream", UpdateStream)
	router.GET("/stream/:streamId", GetStreamByID)
	router.GET("/streams", GetStreamList)
	router.POST("/search_by_location", SearchStreamByLocation)
	router.DELETE("/stream", DeleteStream)

	log.Fatal(http.ListenAndServe(":8080", router))
}

//
//
// http router
//
//

// CreateUser router handle.
func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// model
	userRequestModel := lycamplus.UserRequestModel{
		UserName: "zhangsan00",
	}

	// request
	userResponseModel, err := aLycamPlus.UserInstance.Create(&userRequestModel)

	// result
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 400)
	} else {
		fmt.Fprintf(w, "the result: %s \n", userResponseModel.UUID)
	}
}

// UserAssume router handle.
func UserAssume(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	UUID := "3725d420-dc71-11e6-b191-5f7a2ebf06ef"

	tokenResponseModel, err := aLycamPlus.UserInstance.Assume(UUID)

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 400)
	} else {
		fmt.Fprintf(w, "the result: %s \n", tokenResponseModel.Token.AccessToken)
	}
}

// CreateStream router handle.
func CreateStream(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	streamRequestModel := lycamplus.StreamRequestModel{
		Title: "test stream",
	}

	streamResponseModel, err := aLycamPlus.StreamInstance.Create(&streamRequestModel)

	// result
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 400)
	} else {
		fmt.Fprintf(w, "the result: %s \n", streamResponseModel.StreamID)
	}
}

// UpdateStream router handle.
func UpdateStream(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	streamID := "b7d87ea0-dc72-11e6-98af-bb17f4293ffa"

	streamRequestModel := lycamplus.StreamRequestModel{
		Title:       "test_stream",
		Description: "no Description",
	}

	streamResponseModel, err := aLycamPlus.StreamInstance.Update(streamID, &streamRequestModel)

	// result
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 400)
	} else {
		fmt.Fprintf(w, "the result: %s \n", streamResponseModel.StreamID)
	}
}

// GetStreamByID router handle.
func GetStreamByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	streamID := params.ByName("streamId")

	streamResponseModel, err := aLycamPlus.StreamInstance.Show(streamID)

	// result
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 400)
	} else {
		fmt.Fprintf(w, "the result: %s \n", streamResponseModel.Title)
	}

}

// GetStreamList router handle.
func GetStreamList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	streamResponseModelList, err := aLycamPlus.StreamInstance.List(nil)

	// result
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 400)
	} else {
		fmt.Fprintf(w, "the result: %d \n", streamResponseModelList.TotalItems)
	}

}

// SearchStreamByLocation router handle.
func SearchStreamByLocation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	locationModel := lycamplus.LocationModel{
		Lon:    90,
		Lat:    90,
		Radius: 100,
	}

	streamResponseModelList, err := aLycamPlus.StreamInstance.SearchByLocation(&locationModel)

	// result
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 400)
	} else {
		fmt.Fprintf(w, "the result: %d \n", streamResponseModelList.TotalItems)
	}

}

// DeleteStream router handle.
func DeleteStream(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	streamID := "b7d87ea0-dc72-11e6-98af-bb17f4293ffa"

	successModel, err := aLycamPlus.StreamInstance.Delete(streamID)

	// result
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 400)
	} else {
		fmt.Fprintf(w, "the result: %t \n", successModel.Success)
	}
}
