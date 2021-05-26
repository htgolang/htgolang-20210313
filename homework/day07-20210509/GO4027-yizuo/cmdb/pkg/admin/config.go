package admin

import "cmdb/tools"

// 用户登录日志文件,注意创建日志的目录必须存在
var UserLoginLogFile string = "logs/admin/userLogin.log"

// 初始用户信息
var InitialUser = []string{"1", "yizuo", tools.Md5sum("yizuo"), "Wuhan", "66666666"}

// 用户数据
var (
	// 用户数据存储类型
	UsersDataStorageType string = "json"

	// 保留N个用户文件，包含自身文件
	KeepUserDataFileNum int = 4

	// 用户数据目录
	UserDataDir string = "data"
	// 用户数据文件名字前缀
	UserDataFileStr string = "userData"

	// Json用户数据目录路径
	UserDataDirJson string = UserDataDir + "/json/"
	// Json用户数据文件后缀
	UserDataJsonFileFormat string = ".json"
	// Json用户数据文件名称
	UserDataJsonFileName string = UserDataFileStr + UserDataJsonFileFormat
	// Json用户数据文件Json格式文件绝对路径
	UserDataJsonFile string = UserDataDirJson + UserDataJsonFileName

	// CSV用户数据目录路径
	UserDataDirCsv string = UserDataDir + "/csv/"
	// CSV用户数据文件后缀
	UserDataCsvFileFormat string = ".csv"
	// CSV用户数据文件名称
	UserDataCsvFileName string = UserDataFileStr + UserDataCsvFileFormat
	// CSV用户数据文件Csv格式文件绝对路径
	UserDataCsvFile string = UserDataDirCsv + UserDataCsvFileName
)
