package config

import (
	"os"
	"unisun/api/auth-listener/src/constants"
)

func SetENV() {
	os.Setenv(constants.JWT_SECRET, "aSiAZgPRmmw7gN7p9WeQxQ==")
	os.Setenv(constants.CONTEXT_PATH, "/auth")
}
