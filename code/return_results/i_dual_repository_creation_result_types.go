package return_results

//class IDualRepositoryCreationResultTypes(
//Enum):

type IDualRepositoryCreationResultTypes int

const (

	//SUCCEEDED = \
	//auto()
	//
	SUCCEEDED IDualRepositoryCreationResultTypes = iota

	//FAILED_TO_OPEN_EA = \
	//auto()
	FAILED_TO_OPEN_EA

	//
	//FAILED_TO_OPEN_EA_PROJECT = \
	//auto()

	FAILED_TO_OPEN_EA_PROJECT
)
