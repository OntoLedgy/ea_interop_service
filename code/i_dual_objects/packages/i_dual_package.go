package packages

import (
	"github.com/OntoLedgy/ea_interop_service/code/i_dual_objects/collections"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

//
//from ea_interop_service_source.b_code.i_dual_objects.collections.i_dual_connector_collection import \
//IDualConnectorCollection
//from ea_interop_service_source.b_code.i_dual_objects.collections.i_dual_element_collection import IDualElementCollection
//from ea_interop_service_source.b_code.i_dual_objects.packages.i_package import IPackage

//class IDualPackage(
//IPackage):
type IDualPackage struct {
	*ole.IDispatch
}

//def __init__(
//self,
//package):
func (IDualPackage) IPackage() {
	//IPackage.__init__(
	//self)

}

func (iDualPackage IDualPackage) IDualPackage(iDualPackageDispatch *ole.IDispatch) {

	//self.package = \
	//package
	iDualPackage.IDispatch = iDualPackageDispatch

}

//def __get_elements(
//self) \
//-> IDualElementCollection:
func (iDualPackage IDualPackage) GetElements() collections.IDualElementCollection {

	//elements = \
	elements :=
		collections.IDualElementCollection{}

	elementsDispatch :=
		//IDualElementCollection(
		oleutil.MustGetProperty(
			iDualPackage.IDispatch,
			//ea_collection=self.package.Elements)
			"Elements").ToIDispatch()

	elements.IDualElementCollection(elementsDispatch)

	//return \
	//elements
	return elements

}

//def __get_name(
//self) \
//-> str:
func (iDualPackage IDualPackage) PackageName() string {

	//element_name = \
	//self.package.Name

	packageName := oleutil.MustGetProperty(
		iDualPackage.IDispatch,
		"Name").ToString()

	//return \
	//element_name
	return packageName
}

//def __set_name(
//self,
func (iDualPackage IDualPackage) SetPackageName(
	name string) {
	//name: str):

	//self.package.Name = \
	//name
	oleutil.MustPutProperty(
		iDualPackage.IDispatch,
		"Name",
		name)

}

//def __get_flags(
//self) \
//-> str:
//flags = \
//self.package.Flags
//
//return \
//flags

//def __set_flags(
//self,
//flags: str):
//self.package.Flags = \
//flags

//def __get_package_guid(
//self) \
//-> str:
//package_guid = \
//self.package.PackageGUID
//
//return \
//package_guid

//def __get_package_id(
//self) \
//-> int:
func (iDualPackage IDualPackage) PackageID() int64 {

	//package_id = \
	packageId :=
		//self.package.PackageID
		oleutil.MustGetProperty(
			iDualPackage.IDispatch,
			"PackageID").Val

	//return \
	//package_id

	return packageId
}

//def __get_packages(
//self):
func (iDualPackage IDualPackage) GetPackages() collections.IDualPackageCollection {

	//from ea_interop_service_source.b_code.i_dual_objects.collections.i_dual_package_collection import \
	//IDualPackageCollection

	//packages = \
	packages := collections.IDualPackageCollection{}

	//IDualPackageCollection(
	//ea_collection=self.package.Packages)
	packagesDispatch :=
		oleutil.MustGetProperty(
			iDualPackage.IDispatch,
			"Packages").ToIDispatch()

	packages.IDualPackageCollection(
		packagesDispatch)

	//return \
	//packages
	return packages

}

//def __get_connectors(
//self) \
//-> IDualConnectorCollection:
func (iDualPackage IDualPackage) GetConnectors() collections.IDualConnectorCollection {

	//connector_collection = \
	connectorCollection := collections.IDualConnectorCollection{}

	//IDualConnectorCollection(

	connectorCollectionDispatch :=
		oleutil.MustGetProperty(
			iDualPackage.IDispatch,
			//ea_collection=self.package.Connectors)
			"Connectors").ToIDispatch()

	connectorCollection.IDualConnectorCollection(
		connectorCollectionDispatch)

	//return \
	//connector_collection

	return connectorCollection
}

//def __get_stereotype_ex(
//self) \
//-> str:
//stereotype_ex = \
//self.package.StereotypeEx
//
//return \
//stereotype_ex

//def __set_stereotype_ex(
//self,
//stereotype_ex: str):
//self.package.StereotypeEx = \
//stereotype_ex

//def update(
//self):
//self.package.Update()

//def refresh(
//self):
//self.package.Refresh()

//elements = \
//property(
//fget=__get_elements)
//
//flags = \
//property(
//fget=__get_flags,
//fset=__set_flags)
//
//name = \
//property(
//fget=__get_name,
//fset=__set_name)
//
//package_guid = \
//property(
//fget=__get_package_guid)
//
//package_id = \
//property(
//fget=__get_package_id)
//
//packages = \
//property(
//fget=__get_packages)
//
//connectors = \
//property(
//fget=__get_connectors)
//
//stereotype_ex = \
//property(
//fget=__get_stereotype_ex,
//fset=__set_stereotype_ex)
