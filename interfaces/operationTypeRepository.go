package interfaces

type DbOperationTypeRepo DbRepo

func NewDbOperationTypeRepo(dbHandlers map[string]DbHandler) *DbOperationTypeRepo {
	dbOperationTypeRepo := new(DbOperationTypeRepo)
	dbOperationTypeRepo.dbHandlers = dbHandlers
	dbOperationTypeRepo.dbHandler = dbHandlers["DbOperationTypeRepo"]
	return dbOperationTypeRepo
}
