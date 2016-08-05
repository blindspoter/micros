package settings

const NAME = "micros"

var buildVersion = "dev"

func IsDevelop() bool {
	return "dev" == buildVersion
}
