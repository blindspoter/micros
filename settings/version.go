package settings

const NAME = "ymir"

var buildVersion = "dev"

func IsDevelop() bool {
	return "dev" == buildVersion
}
