package conf

import (
	_ "beemongo/hooks/app/start"
)

func init() {
	overrideRecoverFunc()
}
