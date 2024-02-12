package main

import (
	"fmt"
	"log"
	"time"

	xmlrpc "github.com/kolo/xmlrpc"
)

var (
	url      = "http://localhost:8116"
	db       = "16hris"
	username = "odoo"
	password = "odoo"
)

func main() {
	// Start waktu untuk menghitung berapa lama program di jalankan
	startTime := time.Now()

	// endpoint menyediakan meta-call yang tidak memerlukan autentikasi,
	client, err := xmlrpc.NewClient(fmt.Sprintf("%s/xmlrpc/2/common", url), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ini client", client)

	common := map[string]any{}
	if err := client.Call("version", nil, &common); err != nil {
		log.Fatal(err)
	}
	fmt.Println("ini common", common)

	var uid int64
	if err := client.Call("authenticate", []any{
		db, username, password,
		map[string]any{},
	}, &uid); err != nil {
		log.Fatal(err)
	}
	fmt.Println("ini uid", uid)

	models, err := xmlrpc.NewClient(fmt.Sprintf("%s/xmlrpc/2/object", url), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ini models", models)

	var result bool
	if err := models.Call("execute_kw", []any{
		db, uid, password,
		"res.partner", "check_access_rights",
		[]string{"read"},
		map[string]bool{"raise_exception": false},
	}, &result); err != nil {
		log.Fatal(err)
	}
	fmt.Println("ini result", result)

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
	fmt.Println("ini records", records)

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
	fmt.Println("ini recordFileds", recordFields)

	var recordFields1 []map[string]any
	if err := models.Call("execute_kw", []any{
		db, uid, password,
		"res.partner", "search_read",
		[]any{[]any{
			[]any{"is_company", "=", true},
		}},
		map[string]any{
			"fields": []string{"name", "country_id", "comment"},
			"limit":  5,
		},
	}, &recordFields1); err != nil {
		log.Fatal(err)
	}
	fmt.Println("ini recordFileds1", recordFields1)

	// Hitung durasi waktu untuk mendapatkan waktu eksekusi
	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	fmt.Printf("Waktu eksekusi program: %s\n", executionTime)
}
