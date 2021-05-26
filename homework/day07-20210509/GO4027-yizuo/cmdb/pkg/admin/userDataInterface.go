package admin

type AdminUserDataClient interface {
	ReadUsersData()
	WritesUsersData()
	// CheckUserDataFileExist()
	CopyUserDataFile()
	PersistenceOfLastNChanges()
}
