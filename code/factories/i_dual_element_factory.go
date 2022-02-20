package factories

//from ea_interop_service_source.b_code.i_dual_objects.elements.i_dual_element import IDualElement

import (
	"github.com/OntoLedgy/ea_interop_service/code/i_dual_objects/collections"
	"github.com/OntoLedgy/ea_interop_service/code/i_dual_objects/elements"
)

//def create_i_dual_element(
func CreatIDualElement(
	//container,
	container collections.IDualElementCollection,
	//element_name: str,
	elementName string,
	//element_type: str,
	elementType string,
	//element_notes=str(),
	elementNotes string,
	//stereotype_ex=str()) \
	stereotypeEx string) *elements.IDualElement { //-> IDualElement:

	//i_dual_element = \
	iDualElement :=
		&elements.IDualElement{}
	//IDualElement(

	//container.elements.add_new(
	container.AddNew(
		//ea_object_name=element_name,
		elementName,
		//ea_object_type=element_type))
		elementType)

	//if element_notes:
	if elementNotes != "" {
		//i_dual_element.notes = \
		//element_notes
		iDualElement.SetNotes(elementNotes)
	}

	//if stereotype_ex:
	if stereotypeEx != "" {

		//i_dual_element.stereotype_ex = \
		//stereotype_ex

		iDualElement.StereotypeEx(
			stereotypeEx)
	}

	//i_dual_element.update()
	iDualElement.Update()

	//container.elements.refresh()
	container.Refresh()

	//return \
	//i_dual_element

	return iDualElement

}
