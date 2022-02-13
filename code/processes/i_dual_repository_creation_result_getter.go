package processes

import (
	"github.com/OntoLedgy/ea_interop_service/code/i_dual_objects"
	"github.com/OntoLedgy/ea_interop_service/code/return_results"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

//def get_i_dual_repository_creation_result(
func GetIDualRepositoryCreationResult(
	//ea_project_filename: str) \
	ea_project_filename string) *return_results.IDualRepositoryCreationResults {
	//i_dual_repository_creation_result = \
	i_dual_repository_creation_result :=
		//__get_i_dual_repository_creation_result(
		__get_i_dual_repository_creation_result(
			//ea_project_filename=ea_project_filename)
			ea_project_filename)

	//return \
	//i_dual_repository_creation_result
	return i_dual_repository_creation_result

}

//def __get_i_dual_repository_creation_result(
func __get_i_dual_repository_creation_result(
	//ea_project_filename: str) \
	ea_project_filename string) *return_results.IDualRepositoryCreationResults {

	//com_object_repository = \
	com_object_repository :=

		//__initialise_com_object_repository()
		__initialise_com_object_repository()

	//if com_object_repository is None:
	if com_object_repository == nil {

		//return \
		return nil

		//IDualRepositoryCreationResults(
		//i_dual_repository=INullRepository(),
		//i_dual_repository_creation_result_type=IDualRepositoryCreationResultTypes.FAILED_TO_OPEN_EA)

	}

	//com_object_repository = \
	com_object_repository =
		//__load_com_object_repository(
		__load_com_object_repository(
			//com_object_repository=com_object_repository,
			com_object_repository,
			//ea_project_filename=ea_project_filename)
			ea_project_filename)

	//if com_object_repository is None:
	if com_object_repository == nil {

		//return \
		//IDualRepositoryCreationResults(
		//i_dual_repository=INullRepository(),
		//i_dual_repository_creation_result_type=IDualRepositoryCreationResultTypes.FAILED_TO_OPEN_EA_PROJECT)
		//TODO

	}

	//i_dual_repository_creation_result = \
	i_dual_repository_creation_result :=
		//__get_successful_i_dual_repository_creation_result(
		__get_successful_i_dual_repository_creation_result(
			//com_object_repository=com_object_repository)
			com_object_repository)

	//return \
	//i_dual_repository_creation_result

	return i_dual_repository_creation_result

}

//def __initialise_com_object_repository():
func __initialise_com_object_repository() *ole.IDispatch {

	//try:
	//initial_com_object_repository = \
	//win32com.client.Dispatch(
	//'EA.Repository')

	ea_app_object, _ := oleutil.CreateObject("EA.App")

	ea_repository_object_i_dispatch, _ := ea_app_object.QueryInterface(ole.IID_IDispatch)

	initial_com_object_repository := oleutil.MustGetProperty(ea_repository_object_i_dispatch, "Repository").ToIDispatch()

	//except Exception:
	//initial_com_object_repository = \
	//None

	//return \
	//initial_com_object_repository
	return initial_com_object_repository

}

//def __load_com_object_repository(
func __load_com_object_repository(
	//com_object_repository,
	com_object_repository *ole.IDispatch,
	//ea_project_filename: str):
	ea_project_filename_or_connection_string string) *ole.IDispatch {

	//try:
	//com_object_repository.OpenFile2(
	//ea_project_filename,
	//1,
	//0)

	oleutil.MustCallMethod(
		com_object_repository,
		"OpenFile",
		ea_project_filename_or_connection_string).
		ToIDispatch()

	//
	//except Exception:
	//com_object_repository = \
	//None
	//
	//return \
	//com_object_repository
	//
	//TODO

	return com_object_repository
}

//def __get_successful_i_dual_repository_creation_result(
func __get_successful_i_dual_repository_creation_result(
	//com_object_repository) \
	com_object_repository *ole.IDispatch) *return_results.IDualRepositoryCreationResults { //-> IDualRepositoryCreationResults:

	i_repository :=
		&i_dual_objects.IRepository{}

	//i_dual_repository = \

	i_dual_repository :=
		//IDualRepository(
		&i_dual_objects.IDualRepository{}

	//i_dual_repository.
	i_dual_repository.Initialise(
		i_repository,
		//repository=com_object_repository)
		com_object_repository)

	//i_dual_repository_creation_result = \
	var i_dual_repository_creation_result =
	//IDualRepositoryCreationResults(
	&return_results.IDualRepositoryCreationResults{
		//i_dual_repository=i_dual_repository,
		*i_repository, //TODO - check if this needs to be passed with interface pattern
		//i_dual_repository_creation_result_type=IDualRepositoryCreationResultTypes.SUCCEEDED)
		return_results.SUCCEEDED}

	//return \
	//i_dual_repository_creation_result
	return i_dual_repository_creation_result

}
