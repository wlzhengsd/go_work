/**
功用：全量迁移数据
起因：研究go多协程性能，与java进行对比
结论：程序运行在数据库服务器，同一个库不同表之间的数据迁移[表结构一致]，[windows 4核超频共8核，10个协程，每次读取1000条记录，每100条写库，20万条记录从A到B耗时14616毫秒，平均13333条/秒，性能最好]
问题：测试机4核超频共8核，所有参数如下代码配置性能不错，协程数增加反而运行时间延迟，没有明白什么原因，估计是总共8核，开太多协程造成资源更大消耗，比如开到200个，程序几乎停滞，不插入数据
     共8核开启10个协程效果很好，貌似协程数应该接近核数，性能更好
	 db连接数max_connections可以开启多一些，如果少了程序会异常，提示连接数不够
	 10个协程，每次读取1000条记录，每100条写库，不管怎么修改其中的任意参数都得不能超过这个组合的性能，始终没有闹明白[这个参数组合，是开发程序的时候最初设置的，没有想到居然是最佳组合]

参考资料
1、golang go-sql-drive mysql连接池的实现
http://www.01happy.com/golang-go-sql-drive-mysql-connection-pooling/
2、用Go-SQL-Driver访问mysql数据库
http://lavafree.iteye.com/blog/1827140
*/
package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:111111@tcp(192.168.88.129:3306)/go_demo?charset=utf8")
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(500)
	db.Ping()
}

//获取主键
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

//func getIds() []int {
//	ids := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//	return ids
//}
func toDb(ids []string) {
	var buf bytes.Buffer
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
	buf.WriteString(insql)

	for rows.Next() {
		rnum++
		if rnum != 1 {
			buf.WriteString(",")
		}
		//将行数据保存到record字典
		err = rows.Scan(&order_id, &user_id, &order_status, &consignee, &country, &province, &city, &district, &area, &address, &zipcode, &tel, &email, &best_time, &postscript, &invoice_title, &invoice_type, &invoice_company_code, &express_id, &pay_id, &pay_bank, &pickup_id, &currency, &goods_amount, &imprest, &shipment_id, &shipment_expense, &tax, &weight, &express_sn, &express_update_time, &add_time, &close_time, &ttl, &p_order_id, &related_order_id, &complete_time, &trade_no, &notes, &order_type, &order_flow, &wms_acc, &push_time, &country_name, &wms_name, &sales_type, &order_from, &invoice_stat, &push_stat, &lock_status, &discount, &pay_time, &wms_acc_time, &from_ip, &channel, &consignee_c, &tel_c, &address_c, &tel_idx, &is_cod, &update_time, &consignee_idx, &est_out_of_wh_time, &est_delivered_time, &est_delivered_time_l, &stockout_status)
		//checkErr(err)
		if rnum == 500 {
			//insql = insql + "(" + strconv.Itoa(order_id) + "," + strconv.Itoa(user_id) + "," + strconv.Itoa(order_status) + ",'" + consignee + "'," + strconv.Itoa(country) + "," + strconv.Itoa(province) + "," + strconv.Itoa(city) + "," + strconv.Itoa(district) + "," + strconv.Itoa(area) + ",'" + address + "','" + zipcode + "','" + tel + "','" + email + "'," + strconv.Itoa(best_time) + ",'" + postscript + "','" + invoice_title + "'," + strconv.Itoa(invoice_type) + ",'" + invoice_company_code + "'," + strconv.Itoa(express_id) + "," + strconv.Itoa(pay_id) + ",'" + pay_bank + "'," + strconv.Itoa(pickup_id) + ",'" + currency + "'," + strconv.FormatFloat(goods_amount, 'f', 2, 64) + "," + strconv.FormatFloat(imprest, 'f', 2, 64) + "," + strconv.Itoa(shipment_id) + "," + strconv.FormatFloat(shipment_expense, 'f', 2, 64) + "," + strconv.FormatFloat(tax, 'f', 2, 64) + "," + strconv.FormatFloat(weight, 'f', 2, 64) + ",'" + express_sn + "'," + strconv.Itoa(express_update_time) + "," + strconv.Itoa(add_time) + "," + strconv.Itoa(close_time) + "," + strconv.Itoa(ttl) + "," + strconv.Itoa(p_order_id) + "," + strconv.Itoa(related_order_id) + "," + strconv.Itoa(complete_time) + ",'" + trade_no + "','" + notes + "'," + strconv.Itoa(order_type) + "," + strconv.Itoa(order_flow) + ",'" + wms_acc + "'," + strconv.Itoa(push_time) + ",'" + country_name + "','" + wms_name + "'," + strconv.Itoa(sales_type) + "," + strconv.Itoa(order_from) + "," + strconv.Itoa(invoice_stat) + "," + strconv.Itoa(push_stat) + "," + strconv.Itoa(lock_status) + "," + strconv.FormatFloat(discount, 'f', 2, 64) + "," + strconv.Itoa(pay_time) + "," + strconv.Itoa(wms_acc_time) + ",'" + from_ip + "','" + channel + "','" + consignee_c + "','" + tel_c + "','" + address_c + "','" + tel_idx + "'," + strconv.Itoa(is_cod) + ",'" + update_time + "','" + consignee_idx + "'," + strconv.Itoa(est_out_of_wh_time) + "," + strconv.Itoa(est_delivered_time) + "," + strconv.Itoa(est_delivered_time_l) + "," + strconv.Itoa(stockout_status) + ")"
			//fmt.Println(insql)
			buf.WriteString("(")
			buf.WriteString(strconv.Itoa(order_id))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(user_id))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(order_status))
			buf.WriteString(",'")
			buf.WriteString(consignee)
			buf.WriteString("',")
			buf.WriteString(strconv.Itoa(country))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(province))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(city))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(district))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(area))
			buf.WriteString(",'")
			buf.WriteString(address)
			buf.WriteString("','")
			buf.WriteString(zipcode)
			buf.WriteString("','")
			buf.WriteString(tel)
			buf.WriteString("','")
			buf.WriteString(email)
			buf.WriteString("',")
			buf.WriteString(strconv.Itoa(best_time))
			buf.WriteString(",'")
			buf.WriteString(postscript)
			buf.WriteString("','")
			buf.WriteString(invoice_title)
			buf.WriteString("',")
			buf.WriteString(strconv.Itoa(invoice_type))
			buf.WriteString(",'")
			buf.WriteString(invoice_company_code)
			buf.WriteString("',")
			buf.WriteString(strconv.Itoa(express_id))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(pay_id))
			buf.WriteString(",'")
			buf.WriteString(pay_bank)
			buf.WriteString("',")
			buf.WriteString(strconv.Itoa(pickup_id))
			buf.WriteString(",'")
			buf.WriteString(currency)
			buf.WriteString("',")
			buf.WriteString(strconv.FormatFloat(goods_amount, 'f', 2, 64))
			buf.WriteString(",")
			buf.WriteString(strconv.FormatFloat(imprest, 'f', 2, 64))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(shipment_id))
			buf.WriteString(",")
			buf.WriteString(strconv.FormatFloat(shipment_expense, 'f', 2, 64))
			buf.WriteString(",")
			buf.WriteString(strconv.FormatFloat(tax, 'f', 2, 64))
			buf.WriteString(",")
			buf.WriteString(strconv.FormatFloat(weight, 'f', 2, 64))
			buf.WriteString(",'")
			buf.WriteString(express_sn)
			buf.WriteString("',")
			buf.WriteString(strconv.Itoa(express_update_time))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(add_time))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(close_time))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(ttl))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(p_order_id))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(related_order_id))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(complete_time))
			buf.WriteString(",'")
			buf.WriteString(trade_no)
			buf.WriteString("','")
			buf.WriteString(notes)
			buf.WriteString("',")
			buf.WriteString(strconv.Itoa(order_type))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(order_flow))
			buf.WriteString(",'")
			buf.WriteString(wms_acc)
			buf.WriteString("',")
			buf.WriteString(strconv.Itoa(push_time))
			buf.WriteString(",'")
			buf.WriteString(country_name)
			buf.WriteString("','")
			buf.WriteString(wms_name)
			buf.WriteString("',")
			buf.WriteString(strconv.Itoa(sales_type))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(order_from))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(invoice_stat))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(push_stat))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(lock_status))
			buf.WriteString(",")
			buf.WriteString(strconv.FormatFloat(discount, 'f', 2, 64))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(pay_time))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(wms_acc_time))
			buf.WriteString(",'")
			buf.WriteString(from_ip)
			buf.WriteString("','")
			buf.WriteString(channel)
			buf.WriteString("','")
			buf.WriteString(consignee_c)
			buf.WriteString("','")
			buf.WriteString(tel_c)
			buf.WriteString("','")
			buf.WriteString(address_c)
			buf.WriteString("','")
			buf.WriteString(tel_idx)
			buf.WriteString("',")
			buf.WriteString(strconv.Itoa(is_cod))
			buf.WriteString(",'")
			buf.WriteString(update_time)
			buf.WriteString("','")
			buf.WriteString(consignee_idx)
			buf.WriteString("',")
			buf.WriteString(strconv.Itoa(est_out_of_wh_time))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(est_delivered_time))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(est_delivered_time_l))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(stockout_status))
			buf.WriteString(")")

			db.Exec(buf.String())
			rnum = 0
			buf.Reset()
			buf.WriteString(insql)
			//insql = "insert into xm_order_bak(order_id,user_id,order_status,consignee,country,province,city,district,area,address,zipcode,tel,email,best_time,postscript,invoice_title,invoice_type,invoice_company_code,express_id,pay_id,pay_bank,pickup_id,currency,goods_amount,imprest,shipment_id,shipment_expense,tax,weight,express_sn,express_update_time,add_time,close_time,ttl,p_order_id,related_order_id,complete_time,trade_no,notes,order_type,order_flow,wms_acc,push_time,country_name,wms_name,sales_type,order_from,invoice_stat,push_stat,lock_status,discount,pay_time,wms_acc_time,from_ip,channel,consignee_c,tel_c,address_c,tel_idx,is_cod,update_time,consignee_idx,est_out_of_wh_time,est_delivered_time,est_delivered_time_l,stockout_status) values"
		} else {
			//insql = insql + "(" + strconv.Itoa(order_id) + "," + strconv.Itoa(user_id) + "," + strconv.Itoa(order_status) + ",'" + consignee + "'," + strconv.Itoa(country) + "," + strconv.Itoa(province) + "," + strconv.Itoa(city) + "," + strconv.Itoa(district) + "," + strconv.Itoa(area) + ",'" + address + "','" + zipcode + "','" + tel + "','" + email + "'," + strconv.Itoa(best_time) + ",'" + postscript + "','" + invoice_title + "'," + strconv.Itoa(invoice_type) + ",'" + invoice_company_code + "'," + strconv.Itoa(express_id) + "," + strconv.Itoa(pay_id) + ",'" + pay_bank + "'," + strconv.Itoa(pickup_id) + ",'" + currency + "'," + strconv.FormatFloat(goods_amount, 'f', 2, 64) + "," + strconv.FormatFloat(imprest, 'f', 2, 64) + "," + strconv.Itoa(shipment_id) + "," + strconv.FormatFloat(shipment_expense, 'f', 2, 64) + "," + strconv.FormatFloat(tax, 'f', 2, 64) + "," + strconv.FormatFloat(weight, 'f', 2, 64) + ",'" + express_sn + "'," + strconv.Itoa(express_update_time) + "," + strconv.Itoa(add_time) + "," + strconv.Itoa(close_time) + "," + strconv.Itoa(ttl) + "," + strconv.Itoa(p_order_id) + "," + strconv.Itoa(related_order_id) + "," + strconv.Itoa(complete_time) + ",'" + trade_no + "','" + notes + "'," + strconv.Itoa(order_type) + "," + strconv.Itoa(order_flow) + ",'" + wms_acc + "'," + strconv.Itoa(push_time) + ",'" + country_name + "','" + wms_name + "'," + strconv.Itoa(sales_type) + "," + strconv.Itoa(order_from) + "," + strconv.Itoa(invoice_stat) + "," + strconv.Itoa(push_stat) + "," + strconv.Itoa(lock_status) + "," + strconv.FormatFloat(discount, 'f', 2, 64) + "," + strconv.Itoa(pay_time) + "," + strconv.Itoa(wms_acc_time) + ",'" + from_ip + "','" + channel + "','" + consignee_c + "','" + tel_c + "','" + address_c + "','" + tel_idx + "'," + strconv.Itoa(is_cod) + ",'" + update_time + "','" + consignee_idx + "'," + strconv.Itoa(est_out_of_wh_time) + "," + strconv.Itoa(est_delivered_time) + "," + strconv.Itoa(est_delivered_time_l) + "," + strconv.Itoa(stockout_status) + "),"
			buf.WriteString("(")
			buf.WriteString(strconv.Itoa(order_id))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(user_id))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(order_status))
			buf.WriteString(",'")
			buf.WriteString(consignee)
			buf.WriteString("',")
			buf.WriteString(strconv.Itoa(country))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(province))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(city))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(district))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(area))
			buf.WriteString(",'")
			buf.WriteString(address)
			buf.WriteString("','")
			buf.WriteString(zipcode)
			buf.WriteString("','")
			buf.WriteString(tel)
			buf.WriteString("','")
			buf.WriteString(email)
			buf.WriteString("',")
			buf.WriteString(strconv.Itoa(best_time))
			buf.WriteString(",'")
			buf.WriteString(postscript)
			buf.WriteString("','")
			buf.WriteString(invoice_title)
			buf.WriteString("',")
			buf.WriteString(strconv.Itoa(invoice_type))
			buf.WriteString(",'")
			buf.WriteString(invoice_company_code)
			buf.WriteString("',")
			buf.WriteString(strconv.Itoa(express_id))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(pay_id))
			buf.WriteString(",'")
			buf.WriteString(pay_bank)
			buf.WriteString("',")
			buf.WriteString(strconv.Itoa(pickup_id))
			buf.WriteString(",'")
			buf.WriteString(currency)
			buf.WriteString("',")
			buf.WriteString(strconv.FormatFloat(goods_amount, 'f', 2, 64))
			buf.WriteString(",")
			buf.WriteString(strconv.FormatFloat(imprest, 'f', 2, 64))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(shipment_id))
			buf.WriteString(",")
			buf.WriteString(strconv.FormatFloat(shipment_expense, 'f', 2, 64))
			buf.WriteString(",")
			buf.WriteString(strconv.FormatFloat(tax, 'f', 2, 64))
			buf.WriteString(",")
			buf.WriteString(strconv.FormatFloat(weight, 'f', 2, 64))
			buf.WriteString(",'")
			buf.WriteString(express_sn)
			buf.WriteString("',")
			buf.WriteString(strconv.Itoa(express_update_time))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(add_time))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(close_time))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(ttl))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(p_order_id))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(related_order_id))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(complete_time))
			buf.WriteString(",'")
			buf.WriteString(trade_no)
			buf.WriteString("','")
			buf.WriteString(notes)
			buf.WriteString("',")
			buf.WriteString(strconv.Itoa(order_type))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(order_flow))
			buf.WriteString(",'")
			buf.WriteString(wms_acc)
			buf.WriteString("',")
			buf.WriteString(strconv.Itoa(push_time))
			buf.WriteString(",'")
			buf.WriteString(country_name)
			buf.WriteString("','")
			buf.WriteString(wms_name)
			buf.WriteString("',")
			buf.WriteString(strconv.Itoa(sales_type))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(order_from))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(invoice_stat))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(push_stat))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(lock_status))
			buf.WriteString(",")
			buf.WriteString(strconv.FormatFloat(discount, 'f', 2, 64))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(pay_time))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(wms_acc_time))
			buf.WriteString(",'")
			buf.WriteString(from_ip)
			buf.WriteString("','")
			buf.WriteString(channel)
			buf.WriteString("','")
			buf.WriteString(consignee_c)
			buf.WriteString("','")
			buf.WriteString(tel_c)
			buf.WriteString("','")
			buf.WriteString(address_c)
			buf.WriteString("','")
			buf.WriteString(tel_idx)
			buf.WriteString("',")
			buf.WriteString(strconv.Itoa(is_cod))
			buf.WriteString(",'")
			buf.WriteString(update_time)
			buf.WriteString("','")
			buf.WriteString(consignee_idx)
			buf.WriteString("',")
			buf.WriteString(strconv.Itoa(est_out_of_wh_time))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(est_delivered_time))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(est_delivered_time_l))
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(stockout_status))
			buf.WriteString(")")
		}
	}

	if rnum != 0 {
		db.Exec(buf.String())
	}

	//fmt.Println(record)
	//fmt.Println(strings.TrimRight(insql, ","))
}

//工作池
func worker(id int, jobs <-chan []string, results chan<- int) {
	for j := range jobs {
		//time.Sleep(time.Second * 2)
		//fmt.Println("worker", id, "processing job", 1)
		//fmt.Println(j)
		toDb(j)
		results <- 1
	}
}

func main() {
	//记录开始时间
	start := time.Now()
	//fmt.Println("hello,world!")
	//ids := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//ids := getIds()

	//fmt.Println(ids[:2])

	// 为了使用我们的工作池，我们需要发送工作和接受工作的结果，
	// 这里我们定义两个通道，一个jobs，一个results
	wnum := 10                        //worker协程数
	jobs := make(chan []string, wnum) //超过wnum就不会在写入，可以写入wnum个，即使没有<-read
	results := make(chan int, wnum)

	// 这里启动1000个worker协程，一开始的时候worker阻塞执行，因为
	// jobs通道里面还没有工作任务
	for w := 1; w <= wnum; w++ {
		go worker(w, jobs, results)
	}
	//获取主键
	ids := getIds()
	idnum := len(ids)
	pnum := 2000         //每次处理的id数
	tnum := idnum / pnum //总次数
	if idnum%pnum > 0 {
		tnum = tnum + 1
	}
	// 这里我们发送9个任务，然后关闭通道，告知任务发送完成
	for j := 1; j <= tnum; j++ {
		if j > wnum {
			<-results
		}
		enum := pnum * j
		if enum > idnum {
			enum = idnum
		}
		jobs <- ids[(j-1)*pnum : enum]
		//fmt.Println(j)
	}
	close(jobs)

	// 然后我们从results里面获得结果
	if tnum < wnum {
		wnum = tnum
	}
	for a := 1; a <= wnum; a++ {
		<-results
		//fmt.Println("result")
	}
	end := time.Now()
	//输出执行时间，单位为毫秒。
	fmt.Println(end.Sub(start).Nanoseconds() / 1000000)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
