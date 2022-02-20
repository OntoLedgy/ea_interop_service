package collections

//from ea_interop_service_source.b_code.i_dual_objects.attributes.i_dual_attribute import IDualAttribute
//from ea_interop_service_source.b_code.i_dual_objects.collections.i_dual_collection import IDualCollection

import (
	"github.com/OntoLedgy/ea_interop_service/code/i_dual_objects/attributes"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

//class IDualAttributeCollection(

type IDualAttributeCollection struct {
	*ole.IDispatch
}

//IDualCollection):
//def __init__(
//self,
func (IDualAttributeCollection) IDualCollection() {
	//ea_collection):
	//IDualCollection.__init__(
	//self,
	//ea_collection=ea_collection)
}

func (IDualAttributeCollection) ICollection() {

}

func (IDualAttributeCollection) AddNew() {}

func (iDualAttributeCollection *IDualAttributeCollection) SetElementDispatch(iDualAttributeCollectionDispatch *ole.IDispatch) {

	//self.element = \
	iDualAttributeCollection.IDispatch =
		//element
		iDualAttributeCollectionDispatch

}

//def get_at(
//self,
//index: int) \
func (iDualAttributeCollection IDualAttributeCollection) GetAt(index int64) *attributes.IDualAttribute { //-> IDualAttribute:

	//collection_item = \

	//index)
	collection_item :=
		oleutil.MustCallMethod(
			iDualAttributeCollection.IDispatch,
			//self.ea_collection.GetAt(
			"GetAt",
			index).ToIDispatch()
	//
	//attribute = \
	//IDualAttribute(
	attribute :=
		&attributes.IDualAttribute{
			//attribute=collection_item)
			collection_item}

	//return \
	//attribute
	return attribute

}
