package settings

const NAME = "gin"

var buildVersion = "dev"

func IsDevelop() bool {
	return "dev" == buildVersion
}
