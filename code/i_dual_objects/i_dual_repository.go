package i_dual_objects

import (
	"github.com/OntoLedgy/ea_interop_service/code/i_dual_objects/elements"
	"github.com/OntoLedgy/ea_interop_service/code/i_dual_objects/packages"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

type IDualRepository struct {
	*ole.IDispatch
}

//def __init__(
//self,
func (iDualRepository IDualRepository) InitialiseRepository() {

}

func (iDualRepository *IDualRepository) SetIDualRepositoryDispatch(
	//repository):
	dispatch *ole.IDispatch) {

	iDualRepository.IDispatch =
		dispatch
}

//def custom_command(
//self,
//*args):
//self.repository.CustomCommand(
//*args)

//def get_attribute_by_guid(
//self,
//attribute_guid: str) \
//-> IAttribute:
//attribute = \
//self.repository.GetAttributeByGuid(
//attribute_guid)
//
//if not attribute:
//return \
//INullAttribute()
//
//i_dual_attribute = \
//IDualAttribute(
//attribute = attribute)
//
//return \
//i_dual_attribute

//def get_connector_by_guid(
//self,
//connector_ea_guid: str) \
//-> IConnector:
//connector = \
//self.repository.GetConnectorByGuid(
//connector_ea_guid)
//
//if not connector:
//return \
//INullConnector()
//
//i_dual_connector = \
//IDualConnector(
//connector = connector)
//
//return \
//i_dual_connector

//def get_connector_by_id(
//self,
//connector_id: int) \
//-> IConnector:
//connector = \
//self.repository.GetConnectorByID(
//connector_id)
//
//if not connector:
//return \
//INullConnector()
//
//i_dual_connector = \
//IDualConnector(
//connector = connector)
//
//return \
//i_dual_connector

//def get_element_by_guid(
//self,
func (iDualRepository IDualRepository) GetElementByGuid(
	//element_ea_guid: str) \
	elementEaGuid string) elements.IElement { //-> IElement:

	//element = \
	element :=
		//self.repository.GetElementByGuid(
		oleutil.MustCallMethod(
			iDualRepository.IDispatch,
			"GetElementByGuid",
			elementEaGuid).
			ToIDispatch()
	//element_ea_guid)

	//
	//if not element:
	if element == nil {
		iNullElement := elements.INullElement{}
		iNullElement.Element()
		//return \
		//INullElement()
		return iNullElement
	}

	//
	//i_dual_element = \
	iDualElement := elements.IDualElement{}

	//IDualElement(
	iDualElement.IDispatch = element
	//element = element)

	//return \
	//i_dual_element
	return iDualElement

}

//def get_element_by_id(
//self,
func (iDualRepository IDualRepository) GetElementById(
	//element_id: int) \
	elementId int) elements.IElement { //-> IElement:

	//element = \
	element :=
		//self.repository.GetElementByID(
		//element_id)
		oleutil.MustCallMethod(
			iDualRepository.IDispatch,
			"GetElementByID",
			elementId).
			ToIDispatch()

	//if not element:
	if element == nil {
		iNullElement := elements.INullElement{}
		iNullElement.Element()
		//return \
		//INullElement()
		return iNullElement

	}

	//i_dual_element = \
	iDualElement := elements.IDualElement{}
	//IDualElement(
	iDualElement.IDispatch = element
	//element = element)

	//return \
	//i_dual_element
	return iDualElement
}

//def get_diagram_by_guid(
//self,
//diagram_ea_guid: str) \
//-> IDiagram:
//diagram = \
//self.repository.GetDiagramByGuid(
//diagram_ea_guid)
//
//if not diagram:
//return \
//INullDiagram()
//
//i_dual_diagram = \
//IDualDiagram(
//diagram = diagram)
//
//return \
//i_dual_diagram
//
//def get_diagram_by_id(
//self,
//diagram_id: int) \
//-> IDiagram:
//diagram = \
//self.repository.GetDiagramByID(
//diagram_id)
//
//if not diagram:
//return \
//INullDiagram()
//
//i_dual_diagram = \
//IDualDiagram(
//diagram = diagram)
//
//return \
//i_dual_diagram

//def get_package_by_guid(
//self,
//package_ea_guid: str) \
//-> IPackage:
//package = \
//self.repository.GetPackageByGuid(
//package_ea_guid)
//
//if not package:
//return \
//INullPackage()
//
//i_dual_package = \
//IDualPackage(
//package = package )
//
//return \
//i_dual_package

//def get_package_by_id(
//self,
func (iDualRepository IDualRepository) GetPackageByID(
	//package_id: int) \
	packageId int) packages.IPackage { //-> IPackage:

	//package = \

	dualPackage :=
		oleutil.MustCallMethod(
			iDualRepository.IDispatch,
			//self.repository.GetPackageByID(
			"GetPackageByID",
			//package_id)
			packageId).
			ToIDispatch()

	//if not package:
	if dualPackage == nil {
		iNullPackage := packages.INullPackage{}
		iNullPackage.IPackage()
		//return \
		//INullPackage()
		return iNullPackage

	}

	//i_dual_package = \
	//IDualPackage(
	//package = package )
	iDualPackage := packages.IDualPackage{}

	//IDualElement(
	iDualPackage.IDispatch = dualPackage
	//element = element)

	//return \
	//i_dual_package
	return iDualPackage

}

//def exit(
//self):
//self.repository.Exit()
//
//def get_project_interface(
//self) \
//-> IProject:
//project = \
//self.repository.GetProjectInterface()
//
//if not project:
//return \
//INullProject()
//
//i_dual_project = \
//IDualProject(
//project = project)
//
//return \
//i_dual_project
//
//def refresh_model_view(
//self,
//package_id: int):
//self.repository.RefreshModelView(
//package_id)
//
//def sql_query(
//self,
//sql: str) \
//-> str:
//xml_string = \
//self.repository.SQLQuery(
//sql)
//
//return \
//xml_string
//
//def __get_batch_append(
//self) \
//-> bool:
//batch_append = \
//self.repository.BatchAppend
//
//return \
//batch_append
//
//def __set_batch_append(
//self,
//batch_append: bool):
//self.repository.BatchAppend = \
//batch_append
//
//def __get_enable_ui_updates(
//self) \
//-> bool:
//enable_ui_updates = \
//self.repository.EnableUIUpdates
//
//return \
//enable_ui_updates
//
//def __set_enable_ui_updates(
//self,
//enable_ui_updates: bool):
//self.repository.EnableUIUpdates = \
//enable_ui_updates
//
//def __get_instance_guid(
//self) \
//-> str:
//instance_guid = \
//self.repository.InstanceGUID
//
//return \
//instance_guid
//
//def __get_connection_string(
//self) \
//-> str:
//connection_string = \
//self.repository.ConnectionString
//
//return \
//connection_string
//
//def __get_library_version(
//self) \
//-> int:
//library_version = \
//self.repository.LibraryVersion
//
//return \
//library_version
//
//def __get_models(
//self) \
//-> IDualPackageCollection:
//models = \
//IDualPackageCollection(
//ea_collection = self.repository.Models)
//
//return \
//models
//
//def __get_stereotypes(
//self) \
//-> IDualStereotypeCollection:
//stereotypes = \
//IDualStereotypeCollection(
//ea_collection = self.repository.Stereotypes)
//
//return \
//stereotypes
//
//batch_append = \
//property(
//fget = __get_batch_append,
//fset = __set_batch_append)
//
//connection_string = \
//property(
//fget = __get_connection_string)
//
//enable_ui_updates = \
//property(
//fget = __get_enable_ui_updates,
//fset = __set_enable_ui_updates)
//
//instance_guid = \
//property(
//fget = __get_instance_guid)
//
//library_version = \
//property(
//fget = __get_library_version)
//
//models = \
//property(
//fget = __get_models)
//
//stereotypes = \
//property(
//fget = __get_stereotypes)
