package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
	jsontime "github.com/liamylian/jsontime/v2/v2"
)

const CUS_TIME_FORMAT = "2006-01-02 15:04:05"

// func (ts TimeStamp) MarshalJSON() ([]byte, error) {
// 	t := ts
// 	if y := t.Year(); y < 0 || y >= 10000 {
// 		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
// 	}
// 	b := make([]byte, 0, len(CUS_TIME_FORMAT)+2)
// 	b = append(b, '"')
// 	b = t.AppendFormat(b, CUS_TIME_FORMAT)
// 	b = append(b, '"')
// 	return b, nil
// }

// func (ts *time.Time) UnmarshalJSON(data []byte) error {
// 	if string(data) == "null" {
// 		return nil
// 	}
// 	parseTime, err := time.Parse(`"`+CUS_TIME_FORMAT+`"`, string(data))
// 	if err != nil {
// 		return err
// 	}
// 	*ts = TimeStamp(parseTime)
// 	return nil
// }

// func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
// 	// you can now parse b as thoroughly as you want

// 	s := strings.Trim(string(b), "\"")
// 	if s == "null" {
// 		ct.Timer = time.Timer{}
// 		return
// 	}
// 	ct.Timer, err = time.Parse(`"`+CUS_TIME_FORMAT+`"`, s)
// 	return
// }

type Deposit struct {
	Id        int64     `json:"id"`
	Email     string    `json:"email"`
	Txid      string    `json:"txid"`
	Currency  string    `json:"currency"`
	Amount    float64   `json:"amount" orm:"digits(12);decimals(2)"`
	Status    string    `json:"status"`
	CreatedAt Time      `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" orm:"auto_now;type(datetime)"`
}

type Time string

func init() {
	// Need to register model in init
	jsontime.SetDefaultTimeFormat(time.RFC3339, time.Local)
	orm.RegisterModel(new(Deposit))
}

// AddDeposit insert a new Deposit into database and returns
// last inserted Id on success.
func AddDeposit(m *Deposit) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func SeedDeposit(m *[]Deposit) (n int64, err error) {
	o := orm.NewOrm()
	n, err = o.InsertMulti(100, m)
	return
}

// GetDepositById retrieves Deposit by Id. Returns error if
// Id doesn't exist
func GetDepositById(id int64) (v *Deposit, err error) {
	o := orm.NewOrm()
	v = &Deposit{Id: id}
	if err = o.QueryTable(new(Deposit)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDeposit retrieves all Deposit matches certain condition. Returns empty list if
// no records exist
func GetAllDeposit(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Deposit))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("error: invalid order. must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("error: invalid order. must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}

	} else {
		if len(order) != 0 {
			return nil, errors.New("error: unused 'order' fields")
		}
	}

	var l []Deposit
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateDeposit updates Deposit by Id and returns error if
// the record to be updated doesn't exist
func UpdateDepositById(m *Deposit) (err error) {
	o := orm.NewOrm()
	v := Deposit{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTeste deletes Teste by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDeposit(id int64) (err error) {
	o := orm.NewOrm()
	v := Deposit{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Deposit{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// DeleteTeste deletes Teste by Id and returns error if
// the record to be deleted doesn't exist
func UpdateStatus(id int64, status string) (err error) {
	o := orm.NewOrm()
	v := Deposit{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		v.Status = status
		var num int64
		if num, err = o.Update(&v); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}
