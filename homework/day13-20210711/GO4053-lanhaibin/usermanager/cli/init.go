package cli

import (
	"usermanager/service"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"github.com/spf13/cobra"
)

var Cmd *cobra.Command

//  = new(cobra.Command)

var (
	username string
	password string
)

func init() {
	Cmd = &cobra.Command{
		Use:   "cmdb",
		Short: "资产管理系统",
		Long:  "资产管理系统",
	}

	Cmd.AddCommand(run())
	Cmd.AddCommand(syncDb())
	Cmd.AddCommand(createUser())
}

func run() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "runserver",
		Short: "运行web服务",
		Run: func(c *cobra.Command, args []string) {
			web.Run()
		},
	}
	return cmd
}

func syncDb() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "syncdb",
		Short: "同步数据库",
		Long:  "创建表结构及初始数据",
		Run: func(c *cobra.Command, args []string) {
			// dbName, _ := web.AppConfig.String("database::name")
			orm.RunSyncdb("default", false, false)
		},
	}
	return cmd
}

func createUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "创建用户",
		Run: func(c *cobra.Command, args []string) {
			service.Us.Create(username, nil, true, "", "", password)
		},
	}
	cmd.Flags().StringVarP(&username, "user", "u", "", "用户名")
	cmd.MarkFlagRequired("user")
	cmd.Flags().StringVarP(&password, "password", "p", "", "密码")
	cmd.MarkFlagRequired("password")
	return cmd
}
