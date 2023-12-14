package main

import (
	"fmt"
	"log"
	"time"

	xmlrpc "github.com/kolo/xmlrpc"
)

var (
	url      = "http://localhost:8916"
	db       = "16enterprise"
	username = "odoo"
	password = "odoo"
)

func testmain() {
	startTime := time.Now()
	client, err := xmlrpc.NewClient(fmt.Sprintf("%s/xmlrpc/2/common", url), nil)
	if err != nil {
		log.Fatal(err)
	}
	common := map[string]any{}
	if err := client.Call("version", nil, &common); err != nil {
		log.Fatal(err)
	}

	var uid int64
	if err := client.Call("authenticate", []any{
		db, username, password,
		map[string]any{},
	}, &uid); err != nil {
		log.Fatal(err)
	}
	models, err := xmlrpc.NewClient(fmt.Sprintf("%s/xmlrpc/2/object", url), nil)
	if err != nil {
		log.Fatal(err)
	}
	var result bool
	if err := models.Call("execute_kw", []any{
		db, uid, password,
		"res.partner", "check_access_rights",
		[]string{"read"},
		map[string]bool{"raise_exception": false},
	}, &result); err != nil {
		log.Fatal(err)
	}
	var records []int64
	if err := models.Call("execute_kw", []any{
		db, uid, password,
		"res.partner", "search",
		[]any{[]any{
			[]any{"is_company", "=", true},
		}},
	}, &records); err != nil {
		log.Fatal(err)
	}
	fmt.Println("ini data", records)
	endTime := time.Now()

	var recordFields []map[string]any
	if err := models.Call("execute_kw", []any{
		db, uid, password,
		"res.partner", "search_read",
		[]any{[]any{
			[]any{"is_company", "=", true},
		}},
		map[string]any{
			"fields": []string{"name", "country_id", "comment"},
		},
	}, &recordFields); err != nil {
		log.Fatal(err)
	}

	// for index, value := range recordFields {
	// 	fmt.Printf("Data : %d\n", index)
	// 	for _, val := range value {
	// 		fmt.Println("---------", val)
	// 	}
	// }

	// Hitung durasi waktu untuk mendapatkan waktu eksekusi
	executionTime := endTime.Sub(startTime)
	fmt.Printf("Waktu eksekusi program: %s\n", executionTime)
}
