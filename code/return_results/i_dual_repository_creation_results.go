package return_results

import "github.com/OntoLedgy/ea_interop_service/code/i_dual_objects"

type IDualRepositoryCreationResults struct { //class IDualRepositoryCreationResults:
	IDualRepository                   i_dual_objects.IRepository
	IDualRepositoryCreationResultType IDualRepositoryCreationResultTypes
}

//def __init__(
func (
	//self,
	iDualRepositoryCreationResults *IDualRepositoryCreationResults) Initialise(
	i_dual_repository i_dual_objects.IRepository, //IDualRepository: IRepository,
	i_dual_repository_creation_result_type IDualRepositoryCreationResultTypes) { //IDualRepositoryCreationResultType: IDualRepositoryCreationResultTypes):

	iDualRepositoryCreationResults.IDualRepository = //self.IDualRepository = \
		i_dual_repository //IDualRepository

	iDualRepositoryCreationResults.IDualRepositoryCreationResultType = //self.IDualRepositoryCreationResultType = \
		i_dual_repository_creation_result_type //IDualRepositoryCreationResultType

}

//----------------------

//def __enter__(
//self):
//return \
//self

//def __exit__(
//self,
//exception_type,
//exception_value,
//traceback):
//pass
