package testing

import (
	"fmt"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"testing"
	"time"
)

func TestOleConnectivity(t *testing.T) {

	ole.CoInitialize(0)

	ea_app_object, _ := oleutil.CreateObject("EA.App")

	ea_repository_object_i_dispatch, _ := ea_app_object.QueryInterface(ole.IID_IDispatch)

	connection_string := "sparx_ea --- DBType=4;Connect=Provider=MSDASQL.1;Persist Security Info=False;Data Source=sparx_ea;LazyLoad=1;"

	ea_repository_property := oleutil.MustGetProperty(ea_repository_object_i_dispatch, "Repository").ToIDispatch()

	oleutil.MustCallMethod(ea_repository_property, "OpenFile", connection_string).ToIDispatch()

	//ea_models := oleutil.MustGetProperty(ea_repository, "Models").ToIDispatch()
	test_object := oleutil.MustCallMethod(ea_repository_property, "GetElementByID", 7524).ToIDispatch()

	test_object_element_id := oleutil.MustGetProperty(test_object, "ElementID").Val
	fmt.Println(test_object_element_id)

}

func TestOleConnectivityExcel(t *testing.T) {
	ole.CoInitialize(0)
	unknown, _ := oleutil.CreateObject("Excel.Application")
	excel, _ := unknown.QueryInterface(ole.IID_IDispatch)
	oleutil.PutProperty(excel, "Visible", true)
	workbooks := oleutil.MustGetProperty(excel, "Workbooks").ToIDispatch()
	workbook := oleutil.MustCallMethod(workbooks, "Add", nil).ToIDispatch()
	worksheet := oleutil.MustGetProperty(workbook, "Worksheets", 1).ToIDispatch()
	cell := oleutil.MustGetProperty(worksheet, "Cells", 1, 1).ToIDispatch()
	oleutil.PutProperty(cell, "Value", 12345)

	time.Sleep(2000000000)

	oleutil.PutProperty(workbook, "Saved", true)
	oleutil.CallMethod(workbook, "Close", false)
	oleutil.CallMethod(excel, "Quit")
	excel.Release()

	ole.CoUninitialize()
}
