package testing

import (
	"fmt"
	"github.com/OntoLedgy/ea_interop_service/code/i_dual_objects"
	"github.com/OntoLedgy/ea_interop_service/code/i_dual_objects/elements"
	"github.com/OntoLedgy/ea_interop_service/code/i_dual_objects/packages"
	"github.com/OntoLedgy/ea_interop_service/code/processes"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"reflect"
	"testing"
	"time"
)

func TestOleConnectivity(t *testing.T) {

	ole.CoInitialize(0)

	connection_string := "sparx_ea --- DBType=4;Connect=Provider=MSDASQL.1;Persist Security Info=False;Data Source=sparx_ea;LazyLoad=1;"

	ea_app_object,
		appOpenError :=
		oleutil.CreateObject(
			"EA.Repository")

	if appOpenError != nil {
		panic(appOpenError)
	}

	ea_repository_object_i_dispatch,
		iDispatchError :=
		ea_app_object.QueryInterface(
			ole.IID_IDispatch)

	if iDispatchError != nil {
		panic(iDispatchError)
	}

	oleutil.MustCallMethod(
		ea_repository_object_i_dispatch,
		"OpenFile2",
		connection_string, 1, 0)

	test_object := oleutil.MustCallMethod(
		ea_repository_object_i_dispatch,
		"GetElementByID",
		7523).ToIDispatch()

	test_object_element_id :=
		oleutil.MustGetProperty(
			test_object,
			"ElementID").Val

	test_object2 := oleutil.MustCallMethod(
		ea_repository_object_i_dispatch,
		"GetElementByGuid",
		"{B4298AF2-00A1-4e88-AFBB-07DB7B9F4CF5}").ToIDispatch()

	test_object2_element_id :=
		oleutil.MustGetProperty(
			test_object2,
			"ElementID").Val

	fmt.Println(
		test_object_element_id, test_object2_element_id)

	ea_app_object.Release()
	ea_repository_object_i_dispatch.Release()
	ole.CoUninitialize()

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

func TestEACreationResultsGetterConnectivity(t *testing.T) {

	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	connection_string :=
		"sparx_ea --- DBType=4;" +
			"Connect=Provider=MSDASQL.1;" +
			"Persist Security Info=False;" +
			"Data Source=sparx_ea;" +
			"LazyLoad=1;"

	var creationResults = processes.GetIDualRepositoryCreationResult(
		connection_string)

	fmt.Println(
		creationResults.IDualRepositoryCreationResultType)

	iDualRepository := creationResults.IDualRepository.(i_dual_objects.IDualRepository)

	element :=
		iDualRepository.GetElementByGuid("{B4298AF2-00A1-4e88-AFBB-07DB7B9F4CF5}")

	typeOfElements := reflect.TypeOf((*elements.IDualElement)(nil)).Elem()

	typeOfElement := reflect.TypeOf(element)

	if typeOfElement.String() == typeOfElements.String() {
		iDUalElement := element.(elements.IDualElement)

		elementID := iDUalElement.ElementID()
		elementName := iDUalElement.ElementName()
		iDUalElement.SetElementName("System Table Column Cells")

		fmt.Printf("element id : {%v} \n element name : {%s}\n",
			elementID,
			elementName)
	} else {
		fmt.Printf("type of element is: {%s}, should be : {%s}\n", typeOfElement, typeOfElements)
	}

	oleutil.MustCallMethod(
		iDualRepository.IDispatch,
		"CloseFile")

	iDualRepository.IDispatch.Release()

}

func TestEAPackages(t *testing.T) {

	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	connection_string :=
		"sparx_ea --- DBType=4;" +
			"Connect=Provider=MSDASQL.1;" +
			"Persist Security Info=False;" +
			"Data Source=sparx_ea;" +
			"LazyLoad=1;"

	var creationResults = processes.GetIDualRepositoryCreationResult(
		connection_string)

	fmt.Println(
		creationResults.IDualRepositoryCreationResultType)

	iDualRepository := creationResults.IDualRepository.(i_dual_objects.IDualRepository)

	eaPackage :=
		iDualRepository.GetPackageByID(38)

	typeOfElements := reflect.TypeOf((*packages.IDualPackage)(nil)).Elem()

	typeOfElement := reflect.TypeOf(eaPackage)

	if typeOfElement.String() == typeOfElements.String() {
		iDUalElement := eaPackage.(packages.IDualPackage)

		packageID := iDUalElement.PackageID() //TODO
		packageName := iDUalElement.PackageName()
		//iDUalElement.SetPackageName("System Table Column Cells")

		fmt.Printf("package id : {%v} \n package name : {%s}\n",
			packageID,
			packageName)
	} else {
		fmt.Printf("type of element is: {%s}, should be : {%s}\n", typeOfElement, typeOfElements)
	}

	oleutil.MustCallMethod(
		iDualRepository.IDispatch,
		"CloseFile")

	iDualRepository.IDispatch.Release()

}
