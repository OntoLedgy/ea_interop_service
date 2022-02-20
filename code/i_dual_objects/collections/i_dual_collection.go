package collections

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

//from ea_interop_service_source.b_code.i_dual_objects.collections.i_collection import ICollection

//class IDualCollection(
//ICollection):
type IDualCollection struct {
	*ole.IDispatch
	count int64
}

//def __init__(
//self,
//ea_collection):
//ICollection.__init__(
//self)

func (IDualCollection) ICollection() {

}
func (iDualCollection IDualCollection) IDualCollection(eaCollection *ole.IDispatch) {

	//self.ea_collection = \
	//ea_collection
	iDualCollection.IDispatch = eaCollection

}

//def delete(
//self,
//index: int) \
//-> int:
//ea_attribute_parent_id = \
//self.ea_collection.Delete(
//index)
//
//return \
//ea_attribute_parent_id

//def add_new(
//self,
func (iDualCollection IDualCollection) AddNew(
	//ea_object_name: str,
	eaObjectName string,
	//ea_object_type: str):
	eaObjectType string) *ole.IDispatch {

	//ea_object = \
	eaObject :=
		oleutil.MustCallMethod(
			iDualCollection.IDispatch,
			//self.ea_collection.AddNew(
			"AddNew",
			//ea_object_name,
			eaObjectName,
			//ea_object_type)
			eaObjectType).ToIDispatch()

	//return \
	//ea_object
	return eaObject
}

//def __get_count(
//self) \
//-> int:
//ea_collection_count = \
//self.ea_collection.Count
//
//return \
//ea_collection_count

//count = \
//property(
//fget=__get_count)

//def refresh(
//self):
func (iDualCollection IDualCollection) Refresh() {

	oleutil.MustCallMethod(
		iDualCollection.IDispatch,
		//self.ea_collection.Refresh()
		"Refresh",
	)
}
