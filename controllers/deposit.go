package controllers

import (
	"cripto-moedas/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/client/httplib"
	beego "github.com/beego/beego/v2/server/web"
)

//  DepositController operations for Deposit
type DepositController struct {
	beego.Controller
}

// URLMapping ...
func (c *DepositController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Deposit
// @Param	body		body 	models.Deposit	true		"body for Deposit content"
// @Success 201 {int} models.Deposit
// @Failure 403 body is empty
// @router / [post]
func (c *DepositController) Post() {
	var v models.Deposit
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddDeposit(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Deposit by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Deposit
// @Failure 403 :id is empty
// @router /:id [get]
func (c *DepositController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetDepositById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Deposit
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Deposit
// @Failure 403
// @router / [get]
func (c *DepositController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllDeposit(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Deposit
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Deposit	true		"body for Deposit content"
// @Success 200 {object} models.Deposit
// @Failure 403 :id is not int
// @router /:id [put]
func (c *DepositController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Deposit{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateDepositById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Deposit
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *DepositController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteDeposit(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Seed deposits
// @Description Seed deposits
// @Success 200 {string} seed success!
// @Failure 403 error to seed
// @router /seed [get]
func (c *DepositController) SeedDeposit() {

	result := &[]models.Deposit{}

	err := httplib.Get("https://testapimockstratum2021.free.beeceptor.com/api/transactions").ToJSON(result)

	if err == nil {
		if n, err := models.SeedDeposit(result); err == nil {
			c.Data["json"] = n
		} else {
			c.Data["json"] = err.Error()
		}
		c.Data["json"] = result
	} else {
		c.Data["json"] = err.Error()
	}

	c.ServeJSON()
}

// Delete ...
// @Title Update status deposit
// @Description Use to updade status
// @Param	id		path 	string	true		"The id you want to update"
// @Success 200 {string} updade status success!
// @Failure 403 error updade status
// @router /status/:id [get]
func (c *DepositController) UpdateStatus() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)

	type Data struct {
		Status string `json:"status`
	}

	type Body struct {
		Status string `json:"status"`
		Data   Data   `json:"data"`
	}

	result := &Body{}

	s, err := httplib.Get(`https://testapimockstratum2021.free.beeceptor.com/api/transaction/` + idStr).String()
	if err == nil {
		t := strings.Replace(s, "]", "", -1)
		b := []byte(t)
		err := json.Unmarshal(b, &result)
		if err != nil {
			c.Data["json"] = err.Error()
		}
		if err := models.UpdateStatus(id, result.Data.Status); err == nil {
			c.Data["json"] = `Update deposit: ` + idStr + `status: ` + result.Data.Status
		} else {
			c.Data["json"] = err.Error()
		}
		c.Data["json"] = result
	} else {
		c.Data["json"] = err.Error()
	}

	c.ServeJSON()
}
