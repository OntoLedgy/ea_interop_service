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
	eaProjectFilename string) *return_results.IDualRepositoryCreationResults {

	//i_dual_repository_creation_result = \
	iDualRepositoryCreationResult :=
		//getIDualRepositoryCreationResult(
		getIDualRepositoryCreationResult(
			//ea_project_filename=ea_project_filename)
			eaProjectFilename)

	//return \
	//i_dual_repository_creation_result
	return iDualRepositoryCreationResult

}

//def __get_i_dual_repository_creation_result(
func getIDualRepositoryCreationResult(
	//ea_project_filename: str) \
	eaProjectFilename string) *return_results.IDualRepositoryCreationResults {

	//com_object_repository = \
	comObjectRepository :=
		//__initialise_com_object_repository()
		initialiseComObjectRepository()

	//if com_object_repository is None:
	if comObjectRepository == nil {
		//return \
		//IDualRepositoryCreationResults(
		nullRepository := i_dual_objects.INullRepository{}

		iDualRepositoryCreationResultsFailedToOpenEA := &return_results.IDualRepositoryCreationResults{
			//i_dual_repository=INullRepository(),
			nullRepository,
			//i_dual_repository_creation_result_type=IDualRepositoryCreationResultTypes.FAILED_TO_OPEN_EA)
			return_results.FAILED_TO_OPEN_EA,
		}
		return iDualRepositoryCreationResultsFailedToOpenEA
	}

	//com_object_repository = \
	comObjectRepository =
		//__load_com_object_repository(
		loadComObjectRepository(
			//com_object_repository=com_object_repository,
			comObjectRepository,
			//ea_project_filename=ea_project_filename)
			eaProjectFilename)

	//if com_object_repository is None:
	if comObjectRepository == nil {
		//return \
		//IDualRepositoryCreationResults(
		nullRepository := i_dual_objects.INullRepository{}

		iDualRepositoryCreationResultsFailedToOpenEAProject := &return_results.IDualRepositoryCreationResults{
			//i_dual_repository=INullRepository(),
			nullRepository,
			//i_dual_repository_creation_result_type=IDualRepositoryCreationResultTypes.FAILED_TO_OPEN_EA_PROJECT)
			return_results.FAILED_TO_OPEN_EA_PROJECT,
		}

		return iDualRepositoryCreationResultsFailedToOpenEAProject
	}

	//i_dual_repository_creation_result = \
	i_dual_repository_creation_result :=
		//__get_successful_i_dual_repository_creation_result(
		getSuccessfulIDualRepositoryCreationResult(
			//com_object_repository=com_object_repository)
			comObjectRepository)

	//return \
	//i_dual_repository_creation_result
	return i_dual_repository_creation_result
}

//def __initialise_com_object_repository():
func initialiseComObjectRepository() *ole.IDispatch {

	//try:
	ea_app_object,
		appOpenError :=
		//win32com.client.Dispatch(
		oleutil.CreateObject(
			//'EA.Repository')
			"EA.Repository")

	if appOpenError != nil {
		panic(appOpenError)
	}

	//initial_com_object_repository = \
	initial_com_object_repository,
		iDispatchError :=
		ea_app_object.QueryInterface(
			ole.IID_IDispatch)

	if iDispatchError != nil {

		//except Exception:
		//initial_com_object_repository = \
		//None

		panic(iDispatchError)
	}

	//return \
	//initial_com_object_repository
	return initial_com_object_repository

}

//def __load_com_object_repository(
func loadComObjectRepository(
	//com_object_repository,
	comObjectRepository *ole.IDispatch,
	//ea_project_filename: str):
	eaProjectFilenameOrConnectionString string) *ole.IDispatch {

	//try:

	openFileResult := oleutil.MustCallMethod(
		comObjectRepository,
		//com_object_repository.OpenFile2(
		"OpenFile2",
		//ea_project_filename,
		eaProjectFilenameOrConnectionString,
		//1,
		1,
		//0
		0)

	//TODO
	if openFileResult == nil {

		//except Exception:
		//com_object_repository = \
		//None
	}

	//return \
	//com_object_repository
	return comObjectRepository
}

//def __get_successful_i_dual_repository_creation_result(
func getSuccessfulIDualRepositoryCreationResult(
	//com_object_repository) \
	comObjectRepository *ole.IDispatch) *return_results.IDualRepositoryCreationResults { //-> IDualRepositoryCreationResults:

	//i_dual_repository =
	iDualRepository :=
		//IDualRepository(
		i_dual_objects.IDualRepository{}

	//i_dual_repository.
	iDualRepository.SetIDualRepositoryDispatch(
		//repository=com_object_repository)
		comObjectRepository)

	//i_dual_repository_creation_result = \
	var i_dual_repository_creation_result =
	//IDualRepositoryCreationResults(
	&return_results.IDualRepositoryCreationResults{
		//i_dual_repository=i_dual_repository,
		iDualRepository, //TODO - check if this needs to be passed with interface pattern
		//i_dual_repository_creation_result_type=IDualRepositoryCreationResultTypes.SUCCEEDED)
		return_results.SUCCEEDED}

	//return \
	//i_dual_repository_creation_result
	return i_dual_repository_creation_result

}
