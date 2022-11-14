package configloader

import "gcreate/conf"

type Config interface {
	Load() *conf.Configuration
}
