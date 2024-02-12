package main

import (
	"fmt"
	"log"
	"reflect"
	"time"

	xmlrpc "github.com/kolo/xmlrpc"
)

var (
	url      = "http://localhost:8116"
	db       = "16hris"
	username = "odoo"
	password = "odoo"
)

// type FieldInfo struct {
//     String    string `xmlrpc:"string"`
//     Help      string `xmlrpc:"help"`
//     Type      string `xmlrpc:"type"`
//     // Tambahkan lebih banyak field jika diperlukan
// }



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
	fmt.Println("ini records", reflect.TypeOf(records), records)

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

	recordFields2 := map[string]interface{}{}
	if err := models.Call("execute_kw", []any{
		db, uid, password,
		"res.partner", "fields_get",
		[]any{},
		map[string][]string{
			"attributes": {"string", "help", "type"},
		},
	}, &recordFields2); err != nil {
		log.Fatal(err)
	}
	fmt.Println("ini recordfields2", recordFields2)
	
	for index, value := range recordFields2 {
		// for key, val := range value{

		// 	switch v := val.(type) {
		// 	case string:
		// 		fmt.Printf("%s is a string: %s\n", key, v)
		// 	case int:
		// 		fmt.Printf("%s is an int: %d\n", key, v)
		// 	case bool:
		// 		fmt.Printf("%s is a boolean: %t\n", key, v)
		// 	default:
		// 		fmt.Printf("%s has an unknown type\n", key)
		// 	}
		// }
		fmt.Println("=>",index,value)

	}
	// for i:= 1; i < len(recordFields2); i++{
	// 	fmt.Println("---", recordFields2[i])
	// }


	// Hitung durasi waktu untuk mendapatkan waktu eksekusi
	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	fmt.Printf("Waktu eksekusi program: %s\n", executionTime)
}


// map[
// 	__last_update:map[string:Last Modified on type:datetime] active:map[string:Active type:boolean] active_lang_count:map[string:Active Lang Count type:integer] activity_date_deadline:map[string:Next Activity Deadline type:date] activity_exception_decoration:map[help:Type of the exception activity on record. string:Activity Exception Decoration type:selection] activity_exception_icon:map[help:Icon to indicate an exception activity. string:Icon type:char] activity_ids:map[string:Activities type:one2many] 
// 	activity_state:map[help:Status based on activities
// Overdue: Due date is already passed
// Today: Activity date is today
// Planned: Future activities. string:Activity State type:selection] activity_summary:map[string:Next Activity Summary type:char] activity_type_icon:map[help:Font awesome icon e.g. fa-tasks string:Activity Type Icon type:char] activity_type_id:map[string:Next Activity Type type:many2one] activity_user_id:map[string:Responsible User type:many2one] additional_info:map[string:Additional info type:char] avatar_1024:map[string:Avatar 1024 type:binary] avatar_128:map[string:Avatar 128 type:binary] avatar_1920:map[string:Avatar type:binary] avatar_256:map[string:Avatar 256 type:binary] avatar_512:map[string:Avatar 512 type:binary] bank_ids:map[string:Banks type:one2many] barcode:map[help:Use a barcode to identify this contact. string:Barcode type:char] category_id:map[string:Tags type:many2many] channel_ids:map[string:Channels type:many2many] child_ids:map[string:Contact type:one2many] city:map[string:City type:char] color:map[string:Color Index type:integer] comment:map[string:Notes type:html] commercial_company_name:map[string:Company Name Entity type:char] commercial_partner_id:map[string:Commercial Entity type:many2one] company_id:map[string:Company type:many2one] company_name:map[string:Company Name type:char] company_registry:map[help:The registry number of the company. Use it if it is different from the Tax ID. It must be unique across all partners of a same country string:Company ID type:char] company_type:map[string:Company Type type:selection] contact_address:map[string:Complete Address type:char] contact_address_complete:map[string:Contact Address Complete type:char] country_code:map[help:The ISO country code in two chars.
// You can use this field for quick search. string:Country Code type:char] country_id:map[string:Country type:many2one] create_date:map[string:Created on type:datetime] create_uid:map[string:Created by type:many2one] date:map[string:Date type:date] display_name:map[string:Display Name type:char] email:map[string:Email type:char] email_formatted:map[help:Format email address "Name <email@domain>" string:Formatted Email type:char] email_normalized:map[help:This field is used to search on email address as the primary email field can contain more than strictly an email address. string:Normalized Email type:char] employee:map[help:Check this box if this contact is an Employee. string:Employee type:boolean] employee_ids:map[help:Related employees based on their private address string:Employees type:one2many] employees_count:map[string:Employees Count type:integer] function:map[string:Job Position type:char] has_message:map[string:Has Message type:boolean] id:map[string:ID type:integer] im_status:map[string:IM Status type:char] image_1024:map[string:Image 1024 type:binary] image_128:map[string:Image 128 type:binary] image_1920:map[string:Image type:binary] image_256:map[string:Image 256 type:binary] image_512:map[string:Image 512 type:binary] image_medium:map[string:Medium-sized image type:binary] industry_id:map[string:Industry type:many2one] is_blacklisted:map[help:If the email address is on the blacklist, the contact won't receive mass mailing anymore, from any list string:Blacklist type:boolean] is_company:map[help:Check if the contact is a company, otherwise it is a person string:Is a Company type:boolean] is_public:map[string:Is Public type:boolean] lang:map[help:All the emails and documents sent to this contact will be translated in this language. string:Language type:selection] message_attachment_count:map[string:Attachment Count type:integer] message_bounce:map[help:Counter of the number of bounced emails for this contact string:Bounce type:integer] message_follower_ids:map[string:Followers type:one2many] message_has_error:map[help:If checked, some messages have a delivery error. string:Message Delivery error type:boolean] message_has_error_counter:map[help:Number of messages with delivery error string:Number of errors type:integer] message_has_sms_error:map[help:If checked, some messages have a delivery error. string:SMS Delivery error type:boolean] message_ids:map[string:Messages type:one2many] message_is_follower:map[string:Is Follower type:boolean] message_main_attachment_id:map[string:Main Attachment type:many2one] message_needaction:map[help:If checked, new messages require your attention. string:Action Needed type:boolean] message_needaction_counter:map[help:Number of messages requiring action string:Number of Actions type:integer] message_partner_ids:map[string:Followers (Partners) type:many2many] mobile:map[string:Mobile type:char] mobile_blacklisted:map[help:Indicates if a blacklisted sanitized phone number is a mobile number. Helps distinguish which number is blacklisted             when there is both a mobile and phone field in a model. string:Blacklisted Phone Is Mobile type:boolean] my_activity_date_deadline:map[string:My Activity Deadline type:date] name:map[string:Name type:char] ocn_token:map[help:Used for sending notification to registered devices string:OCN Token type:char] parent_id:map[string:Related Company type:many2one] parent_name:map[string:Parent name type:char] partner_gid:map[string:Company database ID type:integer] partner_latitude:map[string:Geo Latitude type:float] partner_longitude:map[string:Geo Longitude type:float] partner_share:map[help:Either customer (not a user), either shared user. Indicated the current partner is a customer without access or with a limited access created for sharing data. string:Share Partner type:boolean] phone:map[string:Phone type:char] phone_blacklisted:map[help:Indicates if a blacklisted sanitized phone number is a phone number. Helps distinguish which number is blacklisted             when there is both a mobile and phone field in a model. string:Blacklisted Phone is Phone type:boolean] phone_mobile_search:map[string:Phone/Mobile type:char] phone_sanitized:map[help:Field used to store sanitized phone number. Helps speeding up searches and comparisons. string:Sanitized Number type:char] phone_sanitized_blacklisted:map[help:If the sanitized phone number is on the blacklist, the contact won't receive mass mailing sms anymore, from any list string:Phone Blacklisted type:boolean] ref:map[string:Reference type:char] same_company_registry_partner_id:map[string:Partner with same Company Registry type:many2one] same_vat_partner_id:map[string:Partner with same Tax ID type:many2one] self:map[string:Self type:many2one] signup_expiration:map[string:Signup Expiration type:datetime] signup_token:map[string:Signup Token type:char] signup_type:map[string:Signup Token Type type:char] signup_url:map[string:Signup URL type:char] signup_valid:map[string:Signup Token is Valid type:boolean] state_id:map[string:State type:many2one] street:map[string:Street type:char] street2:map[string:Street2 type:char] title:map[string:Title type:many2one] translated_display_name:map[string:Translated Display Name type:char] type:map[help:- Contact: Use this to organize the contact details of employees of a given company (e.g. CEO, CFO, ...).
// - Invoice Address : Preferred address for all invoices. Selected by default when you invoice an order that belongs to this company.
// - Delivery Address : Preferred address for all deliveries. Selected by default when you deliver an order that belongs to this company.
// - Private: Private addresses are only visible by authorized users and contain sensitive data (employee home addresses, ...).
// - Other: Other address for the company (e.g. subsidiary, ...) string:Address Type type:selection] tz:map[help:When printing documents and exporting/importing data, time values are computed according to this timezone.
// If the timezone is not set, UTC (Coordinated Universal Time) is used.
// Anywhere else, time values are computed according to the time offset of your web client. string:Timezone type:selection] tz_offset:map[string:Timezone offset type:char] user_id:map[help:The internal user in charge of this contact. string:Salesperson type:many2one] user_ids:map[string:Users type:one2many] vat:map[help:The Tax Identification Number. Values here will be validated based on the country format. You can use '/' to indicate that the partner is not subject to tax. string:Tax ID type:char] website:map[string:Website Link type:char] write_date:map[string:Last Updated on type:datetime] write_uid:map[string:Last Updated by type:many2one] zip:map[string:Zip type:char]]