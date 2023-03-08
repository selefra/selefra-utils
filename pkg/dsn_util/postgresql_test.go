package dsn_util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfigByDSN(t *testing.T) {

	// ok case
	dsn := "postgres://xx_username:veryhardpasswd@xxx.xxx.xxx.xxx.xxx.com:5432/xxx_database_name"
	exceptedMosaicPasswdDSN := "postgres://xx_username:*******@xxx.xxx.xxx.xxx.xxx.com:5432/xxx_database_name"
	exceptedOriginalPasswdDSN := dsn
	c, err := NewConfigByDSN(dsn)
	assert.Nil(t, err)
	assert.Equal(t, exceptedMosaicPasswdDSN, c.ToDSN(true))
	assert.Equal(t, exceptedOriginalPasswdDSN, c.ToDSN(false))

	// bad case
	_, err = NewConfigByDSN("asdasdasd")
	assert.NotNil(t, err)

	dsn = "host=127.0.0.1 user=postgres_user password=postgres_user port=25432 dbname=test sslmode=disable"
	c, err = NewConfigByDSN(dsn)
	assert.Nil(t, err)
	t.Log(c)

}
