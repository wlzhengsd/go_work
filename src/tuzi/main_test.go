//数据库连接池测试
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:111111@tcp(192.168.88.129:3306)/go_demo?charset=utf8")
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}

func getIds() []string {
	rows, err := db.Query("SELECT order_id FROM xm_order")
	defer rows.Close()
	checkErr(err)

	var order_id string
	ids := []string{}
	for rows.Next() {
		err := rows.Scan(&order_id)
		checkErr(err)
		ids = append(ids, order_id)
	}

	return ids
}

func toDb(ids []string) {
	//idstr := []string{}
	//for _, i := range ids {
	//idstr = append(idstr, strconv.Itoa(i))
	//}
	rows, err := db.Query("SELECT order_id,user_id,order_status,consignee,country,province,city,district,area,address,zipcode,tel,email,best_time,postscript,invoice_title,invoice_type,invoice_company_code,express_id,pay_id,pay_bank,pickup_id,currency,goods_amount,imprest,shipment_id,shipment_expense,tax,weight,express_sn,express_update_time,add_time,close_time,ttl,p_order_id,related_order_id,complete_time,trade_no,notes,order_type,order_flow,wms_acc,push_time,country_name,wms_name,sales_type,order_from,invoice_stat,push_stat,lock_status,discount,pay_time,wms_acc_time,from_ip,channel,consignee_c,tel_c,address_c,tel_idx,is_cod,update_time,consignee_idx,est_out_of_wh_time,est_delivered_time,est_delivered_time_l,stockout_status FROM xm_order where order_id in (" + strings.Join(ids, ",") + ")")
	defer rows.Close()
	checkErr(err)

	var order_id, user_id, order_status, country, province, city, district, area, best_time, invoice_type, express_id, pay_id, pickup_id, shipment_id, express_update_time, add_time, close_time, ttl, p_order_id, related_order_id, complete_time, order_type, order_flow, push_time, sales_type, order_from, invoice_stat, push_stat, lock_status, pay_time, wms_acc_time, is_cod, est_out_of_wh_time, est_delivered_time, est_delivered_time_l, stockout_status int
	var consignee, address, zipcode, tel, email, postscript, invoice_title, invoice_company_code, pay_bank, currency, express_sn, trade_no, notes, wms_acc, country_name, wms_name, from_ip, channel, consignee_c, tel_c, address_c, tel_idx, update_time, consignee_idx string
	var goods_amount, imprest, shipment_expense, tax, weight, discount float64
	var insql string
	rnum := 0
	insql = "insert into xm_order_bak(order_id,user_id,order_status,consignee,country,province,city,district,area,address,zipcode,tel,email,best_time,postscript,invoice_title,invoice_type,invoice_company_code,express_id,pay_id,pay_bank,pickup_id,currency,goods_amount,imprest,shipment_id,shipment_expense,tax,weight,express_sn,express_update_time,add_time,close_time,ttl,p_order_id,related_order_id,complete_time,trade_no,notes,order_type,order_flow,wms_acc,push_time,country_name,wms_name,sales_type,order_from,invoice_stat,push_stat,lock_status,discount,pay_time,wms_acc_time,from_ip,channel,consignee_c,tel_c,address_c,tel_idx,is_cod,update_time,consignee_idx,est_out_of_wh_time,est_delivered_time,est_delivered_time_l,stockout_status) values"

	for rows.Next() {
		rnum++
		//将行数据保存到record字典
		err = rows.Scan(&order_id, &user_id, &order_status, &consignee, &country, &province, &city, &district, &area, &address, &zipcode, &tel, &email, &best_time, &postscript, &invoice_title, &invoice_type, &invoice_company_code, &express_id, &pay_id, &pay_bank, &pickup_id, &currency, &goods_amount, &imprest, &shipment_id, &shipment_expense, &tax, &weight, &express_sn, &express_update_time, &add_time, &close_time, &ttl, &p_order_id, &related_order_id, &complete_time, &trade_no, &notes, &order_type, &order_flow, &wms_acc, &push_time, &country_name, &wms_name, &sales_type, &order_from, &invoice_stat, &push_stat, &lock_status, &discount, &pay_time, &wms_acc_time, &from_ip, &channel, &consignee_c, &tel_c, &address_c, &tel_idx, &is_cod, &update_time, &consignee_idx, &est_out_of_wh_time, &est_delivered_time, &est_delivered_time_l, &stockout_status)
		checkErr(err)
		if rnum == 100 {
			insql = insql + "(" + strconv.Itoa(order_id) + "," + strconv.Itoa(user_id) + "," + strconv.Itoa(order_status) + ",'" + consignee + "'," + strconv.Itoa(country) + "," + strconv.Itoa(province) + "," + strconv.Itoa(city) + "," + strconv.Itoa(district) + "," + strconv.Itoa(area) + ",'" + address + "','" + zipcode + "','" + tel + "','" + email + "'," + strconv.Itoa(best_time) + ",'" + postscript + "','" + invoice_title + "'," + strconv.Itoa(invoice_type) + ",'" + invoice_company_code + "'," + strconv.Itoa(express_id) + "," + strconv.Itoa(pay_id) + ",'" + pay_bank + "'," + strconv.Itoa(pickup_id) + ",'" + currency + "'," + strconv.FormatFloat(goods_amount, 'f', 2, 64) + "," + strconv.FormatFloat(imprest, 'f', 2, 64) + "," + strconv.Itoa(shipment_id) + "," + strconv.FormatFloat(shipment_expense, 'f', 2, 64) + "," + strconv.FormatFloat(tax, 'f', 2, 64) + "," + strconv.FormatFloat(weight, 'f', 2, 64) + ",'" + express_sn + "'," + strconv.Itoa(express_update_time) + "," + strconv.Itoa(add_time) + "," + strconv.Itoa(close_time) + "," + strconv.Itoa(ttl) + "," + strconv.Itoa(p_order_id) + "," + strconv.Itoa(related_order_id) + "," + strconv.Itoa(complete_time) + ",'" + trade_no + "','" + notes + "'," + strconv.Itoa(order_type) + "," + strconv.Itoa(order_flow) + ",'" + wms_acc + "'," + strconv.Itoa(push_time) + ",'" + country_name + "','" + wms_name + "'," + strconv.Itoa(sales_type) + "," + strconv.Itoa(order_from) + "," + strconv.Itoa(invoice_stat) + "," + strconv.Itoa(push_stat) + "," + strconv.Itoa(lock_status) + "," + strconv.FormatFloat(discount, 'f', 2, 64) + "," + strconv.Itoa(pay_time) + "," + strconv.Itoa(wms_acc_time) + ",'" + from_ip + "','" + channel + "','" + consignee_c + "','" + tel_c + "','" + address_c + "','" + tel_idx + "'," + strconv.Itoa(is_cod) + ",'" + update_time + "','" + consignee_idx + "'," + strconv.Itoa(est_out_of_wh_time) + "," + strconv.Itoa(est_delivered_time) + "," + strconv.Itoa(est_delivered_time_l) + "," + strconv.Itoa(stockout_status) + ")"
			//fmt.Println(insql)
			db.Exec(insql)
			rnum = 0
			insql = "insert into xm_order_bak(order_id,user_id,order_status,consignee,country,province,city,district,area,address,zipcode,tel,email,best_time,postscript,invoice_title,invoice_type,invoice_company_code,express_id,pay_id,pay_bank,pickup_id,currency,goods_amount,imprest,shipment_id,shipment_expense,tax,weight,express_sn,express_update_time,add_time,close_time,ttl,p_order_id,related_order_id,complete_time,trade_no,notes,order_type,order_flow,wms_acc,push_time,country_name,wms_name,sales_type,order_from,invoice_stat,push_stat,lock_status,discount,pay_time,wms_acc_time,from_ip,channel,consignee_c,tel_c,address_c,tel_idx,is_cod,update_time,consignee_idx,est_out_of_wh_time,est_delivered_time,est_delivered_time_l,stockout_status) values"
		} else {
			insql = insql + "(" + strconv.Itoa(order_id) + "," + strconv.Itoa(user_id) + "," + strconv.Itoa(order_status) + ",'" + consignee + "'," + strconv.Itoa(country) + "," + strconv.Itoa(province) + "," + strconv.Itoa(city) + "," + strconv.Itoa(district) + "," + strconv.Itoa(area) + ",'" + address + "','" + zipcode + "','" + tel + "','" + email + "'," + strconv.Itoa(best_time) + ",'" + postscript + "','" + invoice_title + "'," + strconv.Itoa(invoice_type) + ",'" + invoice_company_code + "'," + strconv.Itoa(express_id) + "," + strconv.Itoa(pay_id) + ",'" + pay_bank + "'," + strconv.Itoa(pickup_id) + ",'" + currency + "'," + strconv.FormatFloat(goods_amount, 'f', 2, 64) + "," + strconv.FormatFloat(imprest, 'f', 2, 64) + "," + strconv.Itoa(shipment_id) + "," + strconv.FormatFloat(shipment_expense, 'f', 2, 64) + "," + strconv.FormatFloat(tax, 'f', 2, 64) + "," + strconv.FormatFloat(weight, 'f', 2, 64) + ",'" + express_sn + "'," + strconv.Itoa(express_update_time) + "," + strconv.Itoa(add_time) + "," + strconv.Itoa(close_time) + "," + strconv.Itoa(ttl) + "," + strconv.Itoa(p_order_id) + "," + strconv.Itoa(related_order_id) + "," + strconv.Itoa(complete_time) + ",'" + trade_no + "','" + notes + "'," + strconv.Itoa(order_type) + "," + strconv.Itoa(order_flow) + ",'" + wms_acc + "'," + strconv.Itoa(push_time) + ",'" + country_name + "','" + wms_name + "'," + strconv.Itoa(sales_type) + "," + strconv.Itoa(order_from) + "," + strconv.Itoa(invoice_stat) + "," + strconv.Itoa(push_stat) + "," + strconv.Itoa(lock_status) + "," + strconv.FormatFloat(discount, 'f', 2, 64) + "," + strconv.Itoa(pay_time) + "," + strconv.Itoa(wms_acc_time) + ",'" + from_ip + "','" + channel + "','" + consignee_c + "','" + tel_c + "','" + address_c + "','" + tel_idx + "'," + strconv.Itoa(is_cod) + ",'" + update_time + "','" + consignee_idx + "'," + strconv.Itoa(est_out_of_wh_time) + "," + strconv.Itoa(est_delivered_time) + "," + strconv.Itoa(est_delivered_time_l) + "," + strconv.Itoa(stockout_status) + "),"
		}
	}

	if rnum != 0 {
		db.Exec(insql)
	}

	//fmt.Println(record)
	//fmt.Println(strings.TrimRight(insql, ","))
}

func inDb() {

	var order_id, user_id, order_status, country, province, city, district, area, best_time, invoice_type, express_id, pay_id, pickup_id, goods_amount, imprest, shipment_id, shipment_expense, tax, weight, express_update_time, add_time, close_time, ttl, p_order_id, related_order_id, complete_time, order_type, order_flow, wms_acc, push_time, sales_type, order_from, invoice_stat, push_stat, lock_status, discount, pay_time, wms_acc_time, is_cod, est_out_of_wh_time, est_delivered_time, est_delivered_time_l, stockout_status int
	var consignee, address, zipcode, tel, email, postscript, invoice_title, invoice_company_code, pay_bank, currency, express_sn, trade_no, notes, country_name, wms_name, from_ip, channel, consignee_c, tel_c, address_c, tel_idx, update_time, consignee_idx string
	var insql string
	var rnum int
	order_id = 1
	user_id = 1
	order_status = 1
	country = 1
	province = 1
	city = 1
	district = 1
	area = 1
	best_time = 1
	invoice_type = 1
	express_id = 1
	pay_id = 1
	pickup_id = 1
	goods_amount = 1
	imprest = 1
	shipment_id = 1
	shipment_expense = 1
	tax = 1
	weight = 1
	express_update_time = 1
	add_time = 1
	close_time = 1
	ttl = 1
	p_order_id = 1
	related_order_id = 1
	complete_time = 1
	order_type = 1
	order_flow = 1
	wms_acc = 1
	push_time = 1
	sales_type = 1
	order_from = 1
	invoice_stat = 1
	push_stat = 1
	lock_status = 1
	discount = 1
	pay_time = 1
	wms_acc_time = 1
	is_cod = 1
	est_out_of_wh_time = 1
	est_delivered_time = 1
	est_delivered_time_l = 1
	stockout_status = 1

	consignee = "test"
	address = "test"
	zipcode = "test"
	tel = "test"
	email = "test"
	postscript = "test"
	invoice_title = "test"
	invoice_company_code = "test"
	pay_bank = "test"
	currency = "test"
	express_sn = "test"
	trade_no = "test"
	notes = "test"
	country_name = "test"
	wms_name = "test"
	from_ip = "test"
	channel = "test"
	consignee_c = "test"
	tel_c = "test"
	address_c = "test"
	tel_idx = "test"
	update_time = "test"
	consignee_idx = "test"

	insql = "insert into xm_order(order_id,user_id,order_status,consignee,country,province,city,district,area,address,zipcode,tel,email,best_time,postscript,invoice_title,invoice_type,invoice_company_code,express_id,pay_id,pay_bank,pickup_id,currency,goods_amount,imprest,shipment_id,shipment_expense,tax,weight,express_sn,express_update_time,add_time,close_time,ttl,p_order_id,related_order_id,complete_time,trade_no,notes,order_type,order_flow,wms_acc,push_time,country_name,wms_name,sales_type,order_from,invoice_stat,push_stat,lock_status,discount,pay_time,wms_acc_time,from_ip,channel,consignee_c,tel_c,address_c,tel_idx,is_cod,update_time,consignee_idx,est_out_of_wh_time,est_delivered_time,est_delivered_time_l,stockout_status) values"
	rnum = 0
	for w := 100001; w <= 200000; w++ {
		order_id = w
		rnum++
		if rnum == 100 {
			//fmt.Print(rnum)
			insql = insql + "(" + strconv.Itoa(order_id) + "," + strconv.Itoa(user_id) + "," + strconv.Itoa(order_status) + ",'" + consignee + "'," + strconv.Itoa(country) + "," + strconv.Itoa(province) + "," + strconv.Itoa(city) + "," + strconv.Itoa(district) + "," + strconv.Itoa(area) + ",'" + address + "','" + zipcode + "','" + tel + "','" + email + "'," + strconv.Itoa(best_time) + ",'" + postscript + "','" + invoice_title + "'," + strconv.Itoa(invoice_type) + ",'" + invoice_company_code + "'," + strconv.Itoa(express_id) + "," + strconv.Itoa(pay_id) + ",'" + pay_bank + "'," + strconv.Itoa(pickup_id) + ",'" + currency + "'," + strconv.Itoa(goods_amount) + "," + strconv.Itoa(imprest) + "," + strconv.Itoa(shipment_id) + "," + strconv.Itoa(shipment_expense) + "," + strconv.Itoa(tax) + "," + strconv.Itoa(weight) + ",'" + express_sn + "'," + strconv.Itoa(express_update_time) + "," + strconv.Itoa(add_time) + "," + strconv.Itoa(close_time) + "," + strconv.Itoa(ttl) + "," + strconv.Itoa(p_order_id) + "," + strconv.Itoa(related_order_id) + "," + strconv.Itoa(complete_time) + ",'" + trade_no + "','" + notes + "'," + strconv.Itoa(order_type) + "," + strconv.Itoa(order_flow) + "," + strconv.Itoa(wms_acc) + "," + strconv.Itoa(push_time) + ",'" + country_name + "','" + wms_name + "'," + strconv.Itoa(sales_type) + "," + strconv.Itoa(order_from) + "," + strconv.Itoa(invoice_stat) + "," + strconv.Itoa(push_stat) + "," + strconv.Itoa(lock_status) + "," + strconv.Itoa(discount) + "," + strconv.Itoa(pay_time) + "," + strconv.Itoa(wms_acc_time) + ",'" + from_ip + "','" + channel + "','" + consignee_c + "','" + tel_c + "','" + address_c + "','" + tel_idx + "'," + strconv.Itoa(is_cod) + ",'" + update_time + "','" + consignee_idx + "'," + strconv.Itoa(est_out_of_wh_time) + "," + strconv.Itoa(est_delivered_time) + "," + strconv.Itoa(est_delivered_time_l) + "," + strconv.Itoa(stockout_status) + ")"
			//fmt.Println(insql)
			db.Exec(insql)
			rnum = 0
			insql = "insert into xm_order(order_id,user_id,order_status,consignee,country,province,city,district,area,address,zipcode,tel,email,best_time,postscript,invoice_title,invoice_type,invoice_company_code,express_id,pay_id,pay_bank,pickup_id,currency,goods_amount,imprest,shipment_id,shipment_expense,tax,weight,express_sn,express_update_time,add_time,close_time,ttl,p_order_id,related_order_id,complete_time,trade_no,notes,order_type,order_flow,wms_acc,push_time,country_name,wms_name,sales_type,order_from,invoice_stat,push_stat,lock_status,discount,pay_time,wms_acc_time,from_ip,channel,consignee_c,tel_c,address_c,tel_idx,is_cod,update_time,consignee_idx,est_out_of_wh_time,est_delivered_time,est_delivered_time_l,stockout_status) values"
		} else {
			insql = insql + "(" + strconv.Itoa(order_id) + "," + strconv.Itoa(user_id) + "," + strconv.Itoa(order_status) + ",'" + consignee + "'," + strconv.Itoa(country) + "," + strconv.Itoa(province) + "," + strconv.Itoa(city) + "," + strconv.Itoa(district) + "," + strconv.Itoa(area) + ",'" + address + "','" + zipcode + "','" + tel + "','" + email + "'," + strconv.Itoa(best_time) + ",'" + postscript + "','" + invoice_title + "'," + strconv.Itoa(invoice_type) + ",'" + invoice_company_code + "'," + strconv.Itoa(express_id) + "," + strconv.Itoa(pay_id) + ",'" + pay_bank + "'," + strconv.Itoa(pickup_id) + ",'" + currency + "'," + strconv.Itoa(goods_amount) + "," + strconv.Itoa(imprest) + "," + strconv.Itoa(shipment_id) + "," + strconv.Itoa(shipment_expense) + "," + strconv.Itoa(tax) + "," + strconv.Itoa(weight) + ",'" + express_sn + "'," + strconv.Itoa(express_update_time) + "," + strconv.Itoa(add_time) + "," + strconv.Itoa(close_time) + "," + strconv.Itoa(ttl) + "," + strconv.Itoa(p_order_id) + "," + strconv.Itoa(related_order_id) + "," + strconv.Itoa(complete_time) + ",'" + trade_no + "','" + notes + "'," + strconv.Itoa(order_type) + "," + strconv.Itoa(order_flow) + "," + strconv.Itoa(wms_acc) + "," + strconv.Itoa(push_time) + ",'" + country_name + "','" + wms_name + "'," + strconv.Itoa(sales_type) + "," + strconv.Itoa(order_from) + "," + strconv.Itoa(invoice_stat) + "," + strconv.Itoa(push_stat) + "," + strconv.Itoa(lock_status) + "," + strconv.Itoa(discount) + "," + strconv.Itoa(pay_time) + "," + strconv.Itoa(wms_acc_time) + ",'" + from_ip + "','" + channel + "','" + consignee_c + "','" + tel_c + "','" + address_c + "','" + tel_idx + "'," + strconv.Itoa(is_cod) + ",'" + update_time + "','" + consignee_idx + "'," + strconv.Itoa(est_out_of_wh_time) + "," + strconv.Itoa(est_delivered_time) + "," + strconv.Itoa(est_delivered_time_l) + "," + strconv.Itoa(stockout_status) + "),"
		}
	}

	if rnum != 1 {
		db.Exec(insql)
	}

	//fmt.Println(record)
	//fmt.Println(strings.TrimRight(insql, ","))
}

func main() {
	//记录开始时间
	start := time.Now()
	ids := getIds()
	//fmt.Println(ids)
	//ids := []int{1}
	toDb(ids)
	//inDb()
	//记录结束时间
	end := time.Now()
	//输出执行时间，单位为毫秒。
	fmt.Println(end.Sub(start).Nanoseconds() / 1000000)
}

func test() {
	//startHttpServer()
	rows, err := db.Query("SELECT * FROM webdemo_admin limit 1")
	defer rows.Close()
	checkErr(err)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]string)
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
	}

	fmt.Println(record)
	//fmt.Fprintln(w, "finish")
}

func startHttpServer() {
	http.HandleFunc("/pool", pool)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func pool(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM user limit 1")
	defer rows.Close()
	checkErr(err)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]string)
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
	}

	fmt.Println(record)
	fmt.Fprintln(w, "finish")
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
