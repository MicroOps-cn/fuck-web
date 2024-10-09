package server

import (
	"os"

	logs "github.com/MicroOps-cn/fuck/log"
	"github.com/go-kit/log/level"
	"github.com/howeyc/gopass"
	"github.com/spf13/cobra"

	"github.com/MicroOps-cn/fuck-web/pkg/service"
	"github.com/MicroOps-cn/fuck-web/pkg/service/models"
)

var (
	username string
	password string
	email    string
	fullName string
	role     string
)

// migrateCmd represents the migrate command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User manager",
	Long:  `User manager tools.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var userAddCmd = &cobra.Command{
	Use:   "add",
	Short: "create user",
	Long:  `create user.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := logs.GetDefaultLogger()
		svc := service.New(cmd.Context())
		if password == "-" {
			p, err := gopass.GetPasswdPrompt("please input password: ", true, os.Stdin, os.Stderr)
			if err != nil {
				level.Error(logger).Log("msg", "failed to create user", "err", err)
				os.Exit(1)
			}
			password = string(p)
		}
		if len(username) == 0 {
			level.Error(logger).Log("msg", "username is null")
			os.Exit(1)
		}
		if len(password) == 0 {
			level.Error(logger).Log("msg", "password is null")
			os.Exit(1)
		}

		if len(fullName) == 0 {
			fullName = username
		}

		if err := svc.CreateUser(cmd.Context(), &models.User{
			Username: username,
			Password: []byte(password),
			Email:    email,
			FullName: fullName,
			Role:     role,
			Status:   models.UserMeta_normal,
		}); err != nil {
			level.Error(logger).Log("msg", "failed to create user", "err", err)
		}
	},
}

func init() {
	userCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "username (login name).")
	userCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "user password.")
	userCmd.PersistentFlags().StringVarP(&email, "email", "e", "", "user email.")
	userCmd.PersistentFlags().StringVarP(&fullName, "fullname", "f", "", "user full name.")
	userCmd.PersistentFlags().StringVarP(&role, "role", "r", "user", "user/admin")

	userCmd.AddCommand(userAddCmd)
}

func AddUserCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(userCmd)
	userCmd.PreRun = rootCmd.PreRun
}
