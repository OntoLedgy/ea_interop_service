package elements

import (
	"github.com/OntoLedgy/ea_interop_service/code/i_dual_objects/collections"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

//class IDualElement(
type IDualElement struct {
	//IElement):

	*ole.IDispatch

	//attributes = \
	attributes collections.IDualAttributeCollection

	//classifier_id = \
	classifier int64
	//fget=__get_classifier_id,//TODO
	//fset=__set_classifier_id)//TODO

	//connectors = \
	connectors string
	//fget=__get_connectors)//TODO

	element_guid string //TODO
	//element_guid = \
	//property(
	//fget=__get_element_guid)//TODO

	elmentId int64

	//element_id = \

	//fget=__get_element_id)//TODO

	//elements =
	elements string
	//fget=__get_elements) //TODO

	name string
	//name = \
	//property(
	//fget=__get_name,
	//fset=__set_name)

	notes string //TODO
	//notes = \

	//fget=__get_notes,
	//fset=__set_notes)

	packageId string //TODO
	//package_id = \

	//fget=__get_package_id,
	//fset=__set_package_id)

	stereotype string //TODO
	//stereotype = \

	//fget=__get_stereotype,
	//fset=__set_stereotype)

	stereotypeEx string //todo
	//stereotype_ex = \

	//fget=__get_stereotype_ex,
	//fset=__set_stereotype_ex)

}

//
//def __init__(
//self,
//element):
func (iDualElement IDualElement) Element() {

	//IElement.__init__(
	//self)
}

func (iDualElement *IDualElement) IDualElement(element *ole.IDispatch) {
	//self.element = \
	iDualElement.IDispatch =
		//element
		element
}

//
//def __get_attributes(
//self) \
func (iDualElement *IDualElement) Attributes() collections.IDualAttributeCollection { //TODO IDualAttributeCollection

	elementAttributes :=
		oleutil.MustGetProperty(
			iDualElement.IDispatch,
			//self.element.ElementID
			"Attributes").ToIDispatch()

	//attribute_collection = \
	attribute_collection :=
		//IDualAttributeCollection(
		collections.IDualAttributeCollection{
			//ea_collection=self.element.Attributes)
			elementAttributes}

	//return \
	//attribute_collection

	iDualElement.attributes = attribute_collection

	return attribute_collection
}

//def __get_classifier_id(
//self)
func (iDualElement *IDualElement) GetClassifierID() int64 { //-> int:

	//classifier_id = \
	//self.element.ClassifierID
	classifierId :=
		oleutil.MustGetProperty(
			iDualElement.IDispatch,
			//self.element.ElementID
			"ClassifierID").Val

	//return \
	//classifier_id
	return classifierId

}

//def __set_classifier_id(
//self,
//classifier_id: int):
//self.element.ClassifierID = \
//classifier_id
//
//self.element.Update()

//def __get_connectors(
//self) \
//-> IDualConnectorCollection:
//connector_collection = \
//IDualConnectorCollection(
//ea_collection=self.element.Connectors)
//
//return \
//connector_collection

//def __get_element_guid(
//self) \
//-> str:
//element_guid = \
//self.element.ElementGUID
//
//return \
//element_guid

//def __get_element_id(
//self) \
func (iDualElement IDualElement) ElementID() int64 { //-> int:

	//element_id = \
	elementId :=
		oleutil.MustGetProperty(
			iDualElement.IDispatch,
			//self.element.ElementID
			"ElementID").Val

	//return \
	//element_id
	return elementId

}

//def __get_elements(
//self):
//from ea_interop_service_source.b_code.i_dual_objects.collections.i_dual_element_collection import \
//IDualElementCollection
//
//element_collection = \
//IDualElementCollection(
//ea_collection=self.element.Elements)
//
//return \
//element_collection

//def __get_name(
//self) \
func (iDualElement IDualElement) ElementName() string { //-> str:

	//element_name = \

	elementName :=
		oleutil.MustGetProperty(
			iDualElement.IDispatch,
			//self.element.Name
			"Name").ToString()

	//return \
	//element_name
	return elementName
}

//def __set_name(
//self,
func (iDualElement IDualElement) SetElementName(
	//name: str):
	name string) { //-> str:

	//self.element.Name = \
	//name
	oleutil.MustPutProperty(
		iDualElement.IDispatch,
		//self.element.Name
		"Name",
		name)

	//
	//self.element.Update()
	oleutil.MustCallMethod(
		iDualElement.IDispatch,
		"Update")

}

//def __get_notes(
//self) \
//-> str:
//element_notes = \
//self.element.Notes

//return \
//element_notes

//def __set_notes(
func (iDualElement IDualElement) SetNotes(notes string) {

	//self,
	//notes: str):
	//self.element.Notes = \
	//notes

	oleutil.MustPutProperty(
		iDualElement.IDispatch,
		"Notes",
		notes)

}

//def __get_package_id(
//self) \
//-> int:
//package_id = \
//self.element.PackageID
//
//return \
//package_id

//def __set_package_id(
//self,
//package_id: int):
//self.element.PackageID = \
//package_id

//def __get_stereotype(
//self) \
//-> str:
//stereotype = \
//self.element.Stereotype
//
//return \
//stereotype

//def __set_stereotype(
//self,
//stereotype: str):
//self.element.Stereotype = \
//stereotype

func (iDualElement IDualElement) StereotypeEx(stereotypeEx string) {
	oleutil.MustPutProperty(
		iDualElement.IDispatch,
		"Stereotype",
		stereotypeEx)

}

//def __get_stereotype_ex(
//self) \
//-> str:
//stereotype_ex = \
//self.element.StereotypeEx
//
//return \
//stereotype_ex

//def __set_stereotype_ex(
//self,
//stereotype_ex: str):
//self.element.StereotypeEx = \
//stereotype_ex

func (iDualElement IDualElement) Update() {
	oleutil.MustCallMethod(
		iDualElement.IDispatch,
		"Update")
}

//def update(
//self):
//self.element.Update()

//def refresh(
//self):
//self.element.Refresh()

//attributes = \
//property(
//fget=__get_attributes)
//
//classifier_id = \
//property(
//fget=__get_classifier_id,
//fset=__set_classifier_id)
//
//connectors = \
//property(
//fget=__get_connectors)
//
//element_guid = \
//property(
//fget=__get_element_guid)
//
//element_id = \
//property(
//fget=__get_element_id)
//
//elements = \
//property(
//fget=__get_elements)
//
//name = \
//property(
//fget=__get_name,
//fset=__set_name)
//
//notes = \
//property(
//fget=__get_notes,
//fset=__set_notes)
//
//package_id = \
//property(
//fget=__get_package_id,
//fset=__set_package_id)
//
//stereotype = \
//property(
//fget=__get_stereotype,
//fset=__set_stereotype)
//
//stereotype_ex = \
//property(
//fget=__get_stereotype_ex,
//fset=__set_stereotype_ex)
