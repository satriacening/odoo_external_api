package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	xmlrpc "github.com/kolo/xmlrpc"
)

var (
	url      = "http://localhost:8116"
	db       = "16hris"
	username = "odoo"
	password = "odoo"
)

func count(models *xmlrpc.Client, uid int64) {
	var counter int64
	if err := models.Call("execute_kw", []any{
		db, uid, password,
		"res.partner", "search_count",
		[]any{[]any{
			[]any{"active", "=", true},
		}},
	}, &counter); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Counter", counter)
}
func searchRead(models *xmlrpc.Client, uid int64) {
	var recordFields []map[string]any
	if err := models.Call("execute_kw", []any{
		db, uid, password,
		"res.partner", "search_read",
		[]any{[]any{
			[]any{"active", "=", true},
		}},
		map[string]any{
			"fields": []string{"name", "function"},
			// "limit":  5,
		},
	}, &recordFields); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ini Searc read", recordFields)
}

func createdata(models *xmlrpc.Client) {
	// fmt.Println("Ini create date")
	// models, err := xmlrpc.NewClient(fmt.Sprintf("%s/xmlrpc/2/object", url), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	uid := 2
	name := "Customer "
	total := 2
	for i := 1; i <= total; i++ {
		res := name + strconv.Itoa(i)
		var id int64
		if err := models.Call("execute_kw", []any{
			db, uid, password,
			"res.partner", "create",
			[]map[string]string{
				{"name": res, "function": "Create from external API"},
			},
		}, &id); err != nil {
			log.Fatal(err)
		}
	}
}

func deleteData(models *xmlrpc.Client) {
	//delete record
	var delete bool
	for i := 75; i < 1000; i++ {
		if err := models.Call("execute_kw", []any{
			db, 2, password,
			"res.partner", "unlink",
			[]any{
				[]int64{int64(i)},
			},
		}, &delete); err != nil {
			log.Fatal(err)
		}
	}

}

func main() {
	// Start waktu untuk menghitung berapa lama program di jalankan
	startTime := time.Now()

	// endpoint menyediakan meta-call yang tidak memerlukan autentikasi,
	client, err := xmlrpc.NewClient(fmt.Sprintf("%s/xmlrpc/2/common", url), nil)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("ini client", client)

	common := map[string]any{}
	if err := client.Call("version", nil, &common); err != nil {
		log.Fatal(err)
	}
	// fmt.Println("ini common", common)

	var uid int64
	if err := client.Call("authenticate", []any{
		db, username, password,
		map[string]any{},
	}, &uid); err != nil {
		log.Fatal(err)
	}
	// fmt.Println("ini uid", uid)

	models, err := xmlrpc.NewClient(fmt.Sprintf("%s/xmlrpc/2/object", url), nil)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("ini models", models)

	// createdata(models)
	// deleteData(models)
	// searchRead(models, uid)
	count(models, uid)

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	fmt.Printf("Waktu eksekusi program: %s\n", executionTime)

	// var result bool
	// if err := models.Call("execute_kw", []any{
	// 	db, uid, password,
	// 	"res.partner", "check_access_rights",
	// 	[]string{"read"},
	// 	map[string]bool{"raise_exception": false},
	// }, &result); err != nil {
	// 	log.Fatal(err)
	// }
	// // fmt.Println("ini result", result)

	// var records []int64
	// if err := models.Call("execute_kw", []any{
	// 	db, uid, password,
	// 	"res.partner", "search",
	// 	[]any{[]any{
	// 		[]any{"is_company", "=", true},
	// 	}},
	// }, &records); err != nil {
	// 	log.Fatal(err)
	// }
	// // fmt.Println("ini records", reflect.TypeOf(records), records)

	// var recordFields []map[string]any
	// if err := models.Call("execute_kw", []any{
	// 	db, uid, password,
	// 	"res.partner", "search_read",
	// 	[]any{[]any{
	// 		[]any{"is_company", "=", true},
	// 	}},
	// 	map[string]any{
	// 		"fields": []string{"name", "country_id", "comment"},
	// 	},
	// }, &recordFields); err != nil {
	// 	log.Fatal(err)
	// }
	// // fmt.Println("ini recordFileds", recordFields)

	// var recordFields1 []map[string]any
	// if err := models.Call("execute_kw", []any{
	// 	db, uid, password,
	// 	"res.partner", "search_read",
	// 	[]any{[]any{
	// 		[]any{"is_company", "=", true},
	// 	}},
	// 	map[string]any{
	// 		"fields": []string{"name", "country_id", "comment"},
	// 		"limit":  5,
	// 	},
	// }, &recordFields1); err != nil {
	// 	log.Fatal(err)
	// }
	// // fmt.Println("ini recordFileds1", recordFields1)

	// recordFields2 := map[string]interface{}{}
	// if err := models.Call("execute_kw", []any{
	// 	db, uid, password,
	// 	"res.partner", "fields_get",
	// 	[]any{},
	// 	map[string][]string{
	// 		"attributes": {"string", "help", "type"},
	// 	},
	// }, &recordFields2); err != nil {
	// 	log.Fatal(err)
	// }
	// // fmt.Println("ini recordfields2", recordFields2)
	// // jsonString, err := json.Marshal(recordFields2)
	// // fmt.Println(">>>", string(jsonString))

	// var ids []int64
	// if err := models.Call("execute_kw", []any{
	// 	db, uid, password,
	// 	"res.partner", "search",
	// 	[]any{[]any{
	// 		[]any{"is_company", "=", true},
	// 	}},
	// 	map[string]int64{"limit": 1},
	// }, &ids); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("ini ids", ids)

	// // hanya bisa membaca 1 record
	// var recordFields3 []map[string]any
	// if err := models.Call("execute_kw", []any{
	// 	db, uid, password,
	// 	"res.partner", "read",
	// 	ids,
	// 	map[string][]string{
	// 		"fields": {"name", "country_id", "comment"},
	// 	},
	// }, &recordFields3); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("ini recordfields3", recordFields3)

	// create record
	// var id int64
	// if err := models.Call("execute_kw", []any{
	// 	db, uid, password,
	// 	"res.partner", "create",
	// 	[]map[string]string{
	// 		{"name": "Tony", "function": "Tukang scam"},
	// 	},
	// }, &id); err != nil {
	// 	log.Fatal(err)
	// }

	// update
	// var uid int32 // user id
	// var update bool
	// if err := models.Call("execute_kw", []any{
	// 	db, uid, password,
	// 	"res.partner", "write",
	// 	[]any{
	// 		[]int64{uid},
	// 		map[string]string{"name": "Newer partner"},
	// 	},
	// }, &update); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("ini update", update)

	// //delete record
	// var delete bool
	// if err := models.Call("execute_kw", []any{
	// 	db, uid, password,
	// 	"res.partner", "unlink",
	// 	[]any{
	// 		[]int64{uid},
	// 	},
	// }, &delete); err != nil {
	// 	log.Fatal(err)
	// }
	// // check if the deleted record is still in the database
	// var uid int32
	// var record []any
	// if err := models.Call("execute_kw", []any{
	// 	db, uid, password,
	// 	"res.partner", "search",
	// 	[]any{[]any{
	// 		[]any{"id", "=", 74},
	// 	}},
	// }, &record); err != nil {
	// 	log.Fatal(err)
	// }

	// Hitung durasi waktu untuk mendapatkan waktu eksekusi

}
