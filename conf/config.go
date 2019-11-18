package conf

import (
	_ "github.com/shampoo6/beemongo/hooks/app/start"
)

func init() {
	overrideRecoverFunc()
}
