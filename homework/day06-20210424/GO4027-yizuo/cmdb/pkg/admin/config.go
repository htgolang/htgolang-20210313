package admin

import "cmdb/tools"

// 用户登录日志文件
var UserLoginLogFile string = "logs/admin/userLogin.log"

// 初始用户信息
var InitialUser = []string{"1", "yizuo", tools.Md5sum("yizuo"), "Wuhan", "66666666"}

// 用户数据
var (
	// 保留N个用户文件，包含自身文件
	KeepUserDataFileNum int = 4

	// 用户数据目录
	UserDataDir string = "data"

	// 用户数据目录
	UserDataDirCsv string = UserDataDir + "/csv/"

	// 用户数据文件名称
	UserDataCsvFileStr    string = "userData"
	UserDataCsvFileFormat string = ".csv"
	UserDataCsvFileName   string = UserDataCsvFileStr + UserDataCsvFileFormat

	// 用户数据文件Csv格式文件
	UserDataCsvFile string = UserDataDirCsv + UserDataCsvFileName
)
