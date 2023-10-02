package hasura

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type CasaResponse string
type DepositeResponse string
type GLResponse string

type Data struct {
	ResponseSaving  CasaResponse     `json:"abcsInqSaving"`
	ResponseGiro    CasaResponse     `json:"abcsInqGiro"`
	ResponseDeposit DepositeResponse `json:"abcsInqDeposit"`
	ResponseGL      GLResponse       `json:"abcsInqGL"`
}

func TestMapper(t *testing.T) {
	// var dt = ResponseGraphQL.Data;
	dt := Data{
		ResponseSaving:  "ini response saving", // ada datanya
		ResponseGiro:    "ini response giro",   // ada datanya
		ResponseDeposit: "",
		ResponseGL:      "",
	}

	responseGraphql := reflect.ValueOf(dt)
	fmt.Println(responseGraphql)

	for i := 0; i < responseGraphql.Type().NumField(); i++ { // looping semua data response dt
		if responseGraphql.Field(i).String() != "" { // cek jika datanya tidak "" (string null/string kosong)
			propertyName := responseGraphql.Type().Field(i).Name                // get nama property
			propertyType := responseGraphql.Type().Field(i).Type                // get tipe data property -> CasaResponse
			jsonPropertyName := responseGraphql.Type().Field(i).Tag.Get("json") // get json name -> abcsInqSaving

			if strings.HasPrefix(jsonPropertyName, "abcs") { // jika jsonNamenya 'abcs...' (TLLOG)
				// mapping ke TlLog
				// var dataTlLog = CustomeObjectMapper.Map<propertyType, ABCS_T_TLLOG>(responsGraphql.Field(i))

				// insert ke tllog
				// tlLogRepo.Insert(dataTlLog)
				fmt.Println(propertyName, "tipe data", propertyType)

			} else { // jika jsonName nya 'mbase...'
				// mapping ke MbLog
				// insert ke mblog
			}
		} else {
			continue // data null (bukan resolver yang dipilih)
		}
	}
}
