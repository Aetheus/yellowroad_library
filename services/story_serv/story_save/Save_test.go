package story_save

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"encoding/json"
	"reflect"
)

func TestSave(t *testing.T) {
	Convey("Given a valid JSON string in a Save struct", t, func(){
		validJsonString := `{
								"Name" : "Martha Stewart",
								"Class" : "Archer",
								"HP"    : 50,
								"Inventory" : {
									"minor_potion_healing" : { "quantity" : 1 }
								},
								"Morale" : 100
							}`

		initialSaveData := Save{JsonString:validJsonString}

		Convey("Creating an encoded save string should work", func (){
			encodedSaveString, err := initialSaveData.EncodedSaveString()

			So(err,ShouldBeNil)
			So(len(encodedSaveString),ShouldBeGreaterThan, 0)

			Convey("Given a valid encoded save string, decoding it to a Save struct should work", func (){
				decodedSave, err := DecodeSaveString(encodedSaveString)

				So(err, ShouldBeNil)

				Convey("The decoded save struct should have a valid JsonString that has all the same values as before", func (){
					var person1 interface{}
					var person2 interface{}

					unmarshallErr1 := json.Unmarshal([]byte(initialSaveData.JsonString), &person1)
					unmarshallErr2 := json.Unmarshal([]byte(decodedSave.JsonString), &person2)

					So(unmarshallErr1, ShouldBeNil)
					So(unmarshallErr2, ShouldBeNil)

					So(reflect.DeepEqual(person1,person2), ShouldBeTrue)
				})
			})

		})

		Convey("Applying a valid patch to the save should produce an appropriate resulting JSON string", func (){
			expectedResultJsonString := `	{
												"Name" : "Martha Stewart",
												"Class" : "Archer",
												"HP"    : 25,
												"Inventory" : {
													"minor_potion_healing" : { "quantity" : 1 }
												},
												"Morale" : 50
											}`

			var expectedResultDocument interface{}
			unmarshalErr := json.Unmarshal([]byte(expectedResultJsonString), &expectedResultDocument)
			So(unmarshalErr, ShouldBeNil)

			err := initialSaveData.ApplyEffect(`
				{
					"/Morale" : {
						"op" : "INCR",
						"arg" : -50
					},
					"/HP" : {
						"op" : "INCR",
						"arg" : -25
					}
				}
			`)
			So(err, ShouldBeNil)

			var actualResultDocument interface{}
			unmarshalErr = json.Unmarshal([]byte(initialSaveData.JsonString), &actualResultDocument)
			So(unmarshalErr, ShouldBeNil)

			So(reflect.DeepEqual(expectedResultDocument,actualResultDocument), ShouldBeTrue)
		})

		Convey("Given a valid (JSON Schema) requirement that the Save fulfills", func(){
			jsonSchema := `
				{
					"type": "object",
					"required": [
						"Name",
						"Class"
					]
				}
			`

			err := initialSaveData.ValidateRequirements(jsonSchema)
			So(err, ShouldBeNil)
		})

		Convey("Given a valid (JSON Schema) requirement that the Save does not fulfill", func(){
			jsonSchema := `
				{
					"type": "object",
					"required": [
						"KillCount",
					]
				}
			`

			err := initialSaveData.ValidateRequirements(jsonSchema)
			So(err, ShouldNotBeNil)

		})

		Convey("Given an improperly formatted requirement", func(){
			jsonSchema := `
				{ "doorknob" : "pasta"
			`

			err := initialSaveData.ValidateRequirements(jsonSchema)
			So(err, ShouldNotBeNil)
		})
	})

}