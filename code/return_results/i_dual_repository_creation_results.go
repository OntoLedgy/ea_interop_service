package return_results

import "github.com/OntoLedgy/ea_interop_service/code/i_dual_objects"

type IDualRepositoryCreationResults struct { //class IDualRepositoryCreationResults:
	i_dual_repository                      i_dual_objects.IRepository
	i_dual_repository_creation_result_type IDualRepositoryCreationResultTypes
}

func ( //def __init__(
	iDualRepositoryCreationResults *IDualRepositoryCreationResults) Initialise( //self,

	i_dual_repository i_dual_objects.IRepository, //i_dual_repository: IRepository,
	i_dual_repository_creation_result_type IDualRepositoryCreationResultTypes) { //i_dual_repository_creation_result_type: IDualRepositoryCreationResultTypes):

	iDualRepositoryCreationResults.i_dual_repository = //self.i_dual_repository = \
		i_dual_repository //i_dual_repository

	iDualRepositoryCreationResults.i_dual_repository_creation_result_type = //self.i_dual_repository_creation_result_type = \
		i_dual_repository_creation_result_type //i_dual_repository_creation_result_type

}

//----------------------

//def __enter__(
//self):
//return \
//self
//
//def __exit__(
//self,
//exception_type,
//exception_value,
//traceback):
//pass
