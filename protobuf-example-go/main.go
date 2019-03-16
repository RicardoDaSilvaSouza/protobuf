package main

import (
	"fmt"
	"io/ioutil"
	addressbook "protobuf-example-go/src/addressbook"
	complexpb "protobuf-example-go/src/complex"
	enumpb "protobuf-example-go/src/enum_example"
	simplepb "protobuf-example-go/src/simple"
	"time"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"

	"github.com/golang/protobuf/jsonpb"

	"github.com/golang/protobuf/proto"
)

func main() {
	sm := generateSample()

	if err := writeToFile("test_pb.bin", sm); err == nil {

		if newSm, err := readFromFile("test_pb.bin", &simplepb.SimpleMessage{}); err == nil {

			fmt.Println("Old:", sm)
			fmt.Println("new:", *newSm)

			smJSON := toJSON(sm)

			fmt.Println("JSON:", smJSON)

			fmt.Println("From JSON:", fromJSON(smJSON))
		}
	}

	fmt.Println("ENUM:", generateEnum())
	fmt.Println("Complex:", generateComplex())

	doAddressbook()
}

func doAddressbook() {
	adMessage := addressbook.AddressBook{
		People: []*addressbook.Person{
			&addressbook.Person{
				Id:    1,
				Name:  "Ricardo",
				Email: "ricardo@email.com",
				Phones: []*addressbook.Person_PhoneNumber{
					&addressbook.Person_PhoneNumber{
						Number: "18765435647",
						Type:   addressbook.Person_MOBILE,
					},
				},
				LastUpdated: &timestamp.Timestamp{
					Nanos: int32(time.Now().UnixNano()),
				},
			},
		},
	}

	fmt.Println("AddressBook before write:", adMessage)
	if err := writeToFile("addressbook_exercise.bin", &adMessage); err == nil {

		if readMessage, rErr := readFromFile("addressbook_exercise.bin", &addressbook.AddressBook{}); rErr == nil {
			fmt.Println("AddressBook afer read:", *readMessage)
		}
	}
}

func generateComplex() complexpb.ComplexMessage {
	return complexpb.ComplexMessage{
		Msg: &complexpb.DummyMessage{
			Id:   1,
			Name: "Name",
		},
		Dummies: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Name",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Name",
			},
		},
	}
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}

	json, _ := marshaler.MarshalToString(pb)
	return json
}

func fromJSON(message string) proto.Message {
	pb := new(simplepb.SimpleMessage)
	jsonpb.UnmarshalString(message, pb)
	return pb
}

func writeToFile(fileName string, pb proto.Message) error {
	bytes, mErr := proto.Marshal(pb)
	if mErr != nil {
		fmt.Println("Error:", mErr)
		return mErr
	}

	wErr := ioutil.WriteFile(fileName, bytes, 0644)
	if wErr != nil {
		fmt.Println("Error:", wErr)
	}

	return nil
}

func readFromFile(fileName string, pb proto.Message) (*proto.Message, error) {
	if fileBytes, err := ioutil.ReadFile(fileName); err == nil {
		if uErr := proto.Unmarshal(fileBytes, pb); uErr != nil {
			fmt.Println("Error:", uErr)
			return nil, uErr
		}
	} else {
		fmt.Println("Error:", err)
		return nil, err
	}
	return &pb, nil
}

func generateSample() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:       1,
		IsSimple: true,
		Name:     "Example",
		SampleList: []int32{
			1,
			2,
		},
	}

	return &sm
}

func generateEnum() enumpb.EnumMessage {
	return enumpb.EnumMessage{
		Id:           78,
		DayOfTheWeek: enumpb.DayOfTheWeek_MONDAY,
	}
}
