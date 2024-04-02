package main

import (
	"fmt"
	"os"

	"github.com/ibrahimker/golang-praisindo/protobuf/model"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	user1 := &model.User{
		Id:       "u001",
		Name:     "Sylvana Windrunner",
		Password: "f0r Th3 H0rD3",
		Gender:   model.UserGender_FEMALE,
	}

	garage1 := &model.Garage{
		Id:   "g001",
		Name: "Kalimdor",
		Coordinate: &model.GarageCoordinate{
			Latitude:  23.2212847,
			Longitude: 53.22033123,
		},
	}

	garageList := &model.GarageList{
		List: []*model.Garage{
			garage1,
		},
	}

	garageListByUser := &model.GarageListByUser{
		List: map[string]*model.GarageList{
			user1.Id: garageList,
		},
	}

	// =========== original
	fmt.Printf("# ==== Original\n       %#v \n", garageListByUser)
	// =========== as string
	fmt.Printf("# ==== As String\n       %s \n", garageListByUser.String())
	// =========== as json string
	jsonb, err1 := protojson.Marshal(garageListByUser)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}
	fmt.Printf("# ==== As JSON String\n       %s \n", string(jsonb))
}
