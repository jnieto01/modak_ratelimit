package kvs

import (
	/*
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"modak_ratelimit/config"
	"modak_ratelimit/internal/app/entity"
	*/
)


// Happy Path
// This unit test requires redis docker to be running
/*
func Skip_TestConnectToRedis(t *testing.T) {

	err := config.LoadConfig()
	assert.Nil(t, err)

	DB := entity.RedisDB{
		Addr:     config.App.Kvs.Addr,
		Password: config.App.Kvs.Password,
		Db:       config.App.Kvs.Db,
	}

	con, err := NewClient(DB)
	defer con.Close()

	assert.Nil(t, err)

}

func Skip_TestSetAndGetWithTTL(t *testing.T) {

	err := config.LoadConfig()
	assert.Nil(t, err)

	DB := entity.RedisDB{
		Addr:     config.App.Kvs.Addr,
		Password: config.App.Kvs.Password,
		Db:       config.App.Kvs.Db,
	}

	con, err := NewClient(DB)
	defer con.Close()

	assert.Nil(t, err)

	key := "my_key"
	value := "my_value"
	expiration := 1 * time.Minute

	err = con.SetWithTTL(key, expiration, value)
	assert.Nil(t, err)

	retrievedValue, err := con.Get(key)
	assert.Nil(t, err)
	assert.Equal(t, value, retrievedValue, "It is no the same value")
}

func Skip_TestDeleteKey(t *testing.T) {

	err := config.LoadConfig()
	assert.Nil(t, err)

	DB := entity.RedisDB{
		Addr:     config.App.Kvs.Addr,
		Password: config.App.Kvs.Password,
		Db:       config.App.Kvs.Db,
	}

	con, err := NewClient(DB)
	defer con.Close()

	key := "my_key"
	value := "my_value"
	expiration := 1 * time.Minute

	err = con.SetWithTTL(key, expiration, value)
	assert.Nil(t, err)

	err = con.Delete(key)
	assert.NoError(t, err, "Error to delete redis")

	_, err = con.Get(key)
	assert.Error(t, err, "Must get an error")
	assert.Equal(t, err.Error(), "redis: nil", "Must get an error")
}

func Skip_TestGetExp(t *testing.T) {

	err := config.LoadConfig()
	assert.Nil(t, err)

	DB := entity.RedisDB{
		Addr:     config.App.Kvs.Addr,
		Password: config.App.Kvs.Password,
		Db:       config.App.Kvs.Db,
	}

	con, err := NewClient(DB)
	defer con.Close()

	assert.Nil(t, err)

	key := "my_key"
	value := "my_value"
	expiration := 5 * time.Minute

	err = con.SetWithTTL(key, expiration, value)
	assert.Nil(t, err)

	totalSeg, err := con.GetExp(key)
	assert.Nil(t, err)

	minLeft := int(totalSeg.Minutes())
	assert.Equal(t, minLeft, 5 , "Must be the 5 Minutes")


}

func Skip_TestSetExp(t *testing.T) {
	err := config.LoadConfig()
	assert.Nil(t, err)

	DB := entity.RedisDB{
		Addr:     config.App.Kvs.Addr,
		Password: config.App.Kvs.Password,
		Db:       config.App.Kvs.Db,
	}

	con, err := NewClient(DB)
	defer con.Close()

	assert.Nil(t, err)

	key := "my_key"
	value := "my_value"	

	err = con.SetWithTTL(key, 0 , value)
	assert.Nil(t, err)


	expiration := 5 * time.Minute
	err = con.SetExp(key, expiration)
	assert.Nil(t, err)

	totalSeg, err := con.GetExp(key)
	assert.Nil(t, err)

	minLeft := int(totalSeg.Minutes())
	assert.Equal(t, minLeft, 5 , "Must be the 5 Minutes")


}

func Skip_TestExists(t *testing.T) {
	err := config.LoadConfig()
	assert.Nil(t, err)

	DB := entity.RedisDB{
		Addr:     config.App.Kvs.Addr,
		Password: config.App.Kvs.Password,
		Db:       config.App.Kvs.Db,
	}

	con, err := NewClient(DB)
	defer con.Close()

	assert.Nil(t, err)

	key := "my_key"
	value := "my_value"
	expiration := 1 * time.Minute

	err = con.SetWithTTL(key, expiration, value)
	assert.Nil(t, err)

	status, err := con.Exists(key)
	assert.Nil(t, err)
	assert.True(t, status)

}

func Skip_TestNotExists(t *testing.T) {
	err := config.LoadConfig()
	assert.Nil(t, err)

	DB := entity.RedisDB{
		Addr:     config.App.Kvs.Addr,
		Password: config.App.Kvs.Password,
		Db:       config.App.Kvs.Db,
	}

	con, err := NewClient(DB)
	assert.Nil(t, err)
	defer con.Close()

	key := "my_key_test"

	status, err := con.Exists(key)
	assert.Nil(t, err)
	assert.False(t, status)
}

func Skip_TestIntegration(t *testing.T) {

	err := config.LoadConfig()
	assert.Nil(t, err)

	DB := entity.RedisDB{
		Addr:     config.App.Kvs.Addr,
		Password: config.App.Kvs.Password,
		Db:       config.App.Kvs.Db,
	}

	con, err := NewClient(DB)
	defer con.Close()

	key := "my_key"
	value := "my_value"
	expiration := 10 * time.Second

	err = con.SetWithTTL(key, expiration, value)
	assert.Nil(t, err)

	retrievedValue, err := con.Get(key)
	assert.Nil(t, err)
	assert.Equal(t, value, retrievedValue, "It is no the same value")

	err = con.Delete(key)
	assert.NoError(t, err, "Error to delete redis")

	status, err := con.Exists(key)
	assert.Nil(t, err)
	assert.False(t, status, "the key still exists in redis")

}
*/