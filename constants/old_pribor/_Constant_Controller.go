package constants

import (
	"fmt"
	"errors"
	"encoding/json"
	"reflect"
	"strings"
	"context"
	
	"osbe"
	"osbe/fields"
	"osbe/srv"
	"osbe/socket"
	"osbe/model"
	"osbe/response"
	
//	"github.com/jackc/pgx/v4"
)

const RESP_ER_NOT_FOUND = 1000

//Controller
type Constant_Controller struct {
	PublicMethods osbe.PublicMethodCollection	
}

func (c *Constant_Controller) GetID() osbe.ControllerID {
	return osbe.ControllerID("Constant")
}

func (c *Constant_Controller) InitPublicMethods() {

	Constant_Model_init()

	c.PublicMethods = make(osbe.PublicMethodCollection)
	
	//************************** method get_object *************************************
	c.PublicMethods["get_object"] = &Constant_Controller_get_object{
		ModelMetadata: Constant_keys_md,
	}
	
	//************************** method get_list *************************************
	c.PublicMethods["get_list"] = &Constant_Controller_get_list{
		ModelMetadata: model.Cond_Model_md,
	}
	
	//************************** method get_values *************************************
	c.PublicMethods["get_values"] = &Constant_Controller_get_values{
		ModelMetadata: Constant_get_values_md,
	}

	//************************** method set_value *************************************
	c.PublicMethods["set_value"] = &Constant_Controller_set_value{
		ModelMetadata: Constant_set_value_md,
	}
	
}

func (c *Constant_Controller) GetPublicMethod(publicMethodID osbe.PublicMethodID) (pm osbe.PublicMethod, err error) {
	pm, ok := c.PublicMethods[publicMethodID]
	if !ok {
		err = errors.New(fmt.Sprintf(osbe.ER_CONTOLLER_METH_NOT_DEFINED, string(publicMethodID), string(c.GetID())))
	}
	
	return
}

//**************************************************************************************
//Public method: get_list
type Constant_Controller_get_list struct {
	ModelMetadata fields.FieldCollection
	EventList osbe.PublicMethodEventList
}

func (pm *Constant_Controller_get_list) GetEventList() osbe.PublicMethodEventList {
	return pm.EventList
}

func (pm *Constant_Controller_get_list) GetModelMetadata() fields.FieldCollection {
	return pm.ModelMetadata
}

func (c *Constant_Controller_get_list) GetID() osbe.PublicMethodID {
	return osbe.PublicMethodID("get_list")
}

//Public method Unmarshal to structure
func (pm *Constant_Controller_get_list) Unmarshal(payload []byte) (res reflect.Value, err error) {

	//argument structrure
	argv := &model.Controller_get_list_argv{}
	
	err = json.Unmarshal(payload, argv)
	if err != nil {
		return 
	}
	
	res = reflect.ValueOf(&argv.Argv).Elem()
	
	return
}

//Method implemenation
func (pm *Constant_Controller_get_list) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetListOnArgs(app, resp, rfltArgs, ConstantList_md, &ConstantList{})
}

//*******************************************************************************************************
//Public method: get_object
type Constant_Controller_get_object struct {
	ModelMetadata fields.FieldCollection
	EventList osbe.PublicMethodEventList
}

func (pm *Constant_Controller_get_object) GetEventList() osbe.PublicMethodEventList {
	return pm.EventList
}

func (pm *Constant_Controller_get_object) GetModelMetadata() fields.FieldCollection {
	return pm.ModelMetadata
}

func (c *Constant_Controller_get_object) GetID() osbe.PublicMethodID {
	return osbe.PublicMethodID("get_object")
}

//Public method Unmarshal to structure
func (pm *Constant_Controller_get_object) Unmarshal(payload []byte) (res reflect.Value, err error) {

	//argument structrure
	argv := &Constant_keys_argv{}
	
	err = json.Unmarshal(payload, argv)
	if err != nil {
		return 
	}
	
	res = reflect.ValueOf(&argv.Argv).Elem()
	
	return
}

//Method implemenation
func (pm *Constant_Controller_get_object) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {
	return osbe.GetObjectOnArgs(app, resp, rfltArgs, &ConstantList{})
}

//*******************************************************************************************************
//Public method: get_values
type Constant_Controller_get_values struct {
	ModelMetadata fields.FieldCollection
	EventList osbe.PublicMethodEventList
}

func (pm *Constant_Controller_get_values) GetEventList() osbe.PublicMethodEventList {
	return pm.EventList
}

func (pm *Constant_Controller_get_values) GetModelMetadata() fields.FieldCollection {
	return pm.ModelMetadata
}

func (c *Constant_Controller_get_values) GetID() osbe.PublicMethodID {
	return osbe.PublicMethodID("get_values")
}

//Public method Unmarshal to structure
func (pm *Constant_Controller_get_values) Unmarshal(payload []byte) (res reflect.Value, err error) {

	//argument structrure
	argv := &Constant_get_values_argv{}
	
	err = json.Unmarshal(payload, argv)
	if err != nil {
		return 
	}
	
	res = reflect.ValueOf(&argv.Argv).Elem()
	
	return
}

//Method implemenation
func (pm *Constant_Controller_get_values) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {

	args := rfltArgs.Interface().(*Constant_get_values)
	
	fld_sep := osbe.ArgsFieldSep(rfltArgs)
	ids_str := strings.Split(args.Id_list.GetValue(), fld_sep)
	query := ""
	for _, id := range ids_str {
		if query != "" {
			query += " UNION ALL "
		}
		query += fmt.Sprintf(`SELECT
			'%s' AS id,
			const_%s_val()::text AS val,
			(SELECT c.val_type FROM const_%s c) AS val_type`,
			id, id, id);		
	}
	if query != "" {
		pool_conn, pm_err := app.GetSecondaryPoolConn()
		if pm_err != nil {
			return pm_err
		}
		defer pool_conn.Release()
		conn := pool_conn.Conn()
	
		if err := osbe.AddQueryResult(resp, &ConstantValue{}, query, "", nil, conn); err != nil {
			return err
		}
	}
	
	return nil
}

//*******************************************************************************************************
//Public method: set_value
type Constant_Controller_set_value struct {
	ModelMetadata fields.FieldCollection
	EventList osbe.PublicMethodEventList
}

func (pm *Constant_Controller_set_value) GetEventList() osbe.PublicMethodEventList {
	return pm.EventList
}

func (pm *Constant_Controller_set_value) GetModelMetadata() fields.FieldCollection {
	return pm.ModelMetadata
}

func (c *Constant_Controller_set_value) GetID() osbe.PublicMethodID {
	return osbe.PublicMethodID("set_value")
}

//Public method Unmarshal to structure
func (pm *Constant_Controller_set_value) Unmarshal(payload []byte) (res reflect.Value, err error) {

	//argument structrure
	argv := &Constant_set_value_argv{}
	
	err = json.Unmarshal(payload, argv)
	if err != nil {
		return 
	}
	
	res = reflect.ValueOf(&argv.Argv).Elem()
	
	return
}

//Method implemenation
func (pm *Constant_Controller_set_value) Run(app osbe.Applicationer, serv srv.Server, sock socket.ClientSocketer, resp *response.Response, rfltArgs reflect.Value) error {

	args := rfltArgs.Interface().(*Constant_set_value)
	id := args.Id.GetValue()

	pool_conn, pm_err := app.GetPrimaryPoolConn()
	if pm_err != nil {
		return pm_err
	}
	defer pool_conn.Release()
	conn := pool_conn.Conn()
	
	found := false
	const_list := []string{"doc_per_page_count","nds,reserve","zip,base_height","empty_space_rate","cable_chan_rate","part_margin","work_margin"}
	for _, c := range const_list {
		if c == id {
			found = true
			break
		}
	}
	if !found {
		return osbe.NewPublicMethodError(RESP_ER_NOT_FOUND, fmt.Sprintf("?????????????????? %s ???? ??????????????",id))
	}
	_, err := conn.Exec(context.Background(), fmt.Sprintf(`SELECT const_%s_set_val(%s)`, id, args.Val.GetValue()))
	if err != nil {
		return osbe.NewPublicMethodError(response.RESP_ER_INTERNAL, fmt.Sprintf("pgx.Conn.Exec(): %v",err))
	}
	
	return nil
}


