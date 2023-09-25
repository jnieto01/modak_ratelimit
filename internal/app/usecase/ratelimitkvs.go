package usecase

import (
	"net/http"
	"strconv"
	"time"

	"modak_ratelimit/internal/app/i18n"
	"modak_ratelimit/internal/app/utils/kvs"
	"modak_ratelimit/internal/app/utils/logger"
	"modak_ratelimit/config"
	"modak_ratelimit/internal/app/entity"
)



func RateLimitExe(rule entity.RuleByType, r *entity.Response) {

	logger.Info("Start rate limite service, rule: " + rule.Key)

	DB := entity.RedisDB{
		Addr:     config.App.Kvs.Addr,
		Password: config.App.Kvs.Password,
		Db:       config.App.Kvs.Db,
	}

	con, err := kvs.NewClient(DB)
	if err != nil {
		logger.Error("Error to connect to redis", err)
		internalError(r)
		return
	}
	defer con.Close()

	status, err := con.Exists(rule.Key)
	if err != nil {
		logger.Error("Error checking key on redis", err)
		internalError(r)
		return
	}

	if !status {
		err = con.SetWithTTL(rule.Key, time.Duration(rule.TimeInterval)*time.Minute, 1)
		logger.Info("Set new key on redis, key: " + rule.Key)

		r.Status = http.StatusOK
		r.Data.IsAllowed = true
		r.Data.Error = i18n.NotError
		return
	}

	value, err := con.Get(rule.Key)
	if err != nil {
		logger.Error("Error getting the count on redis", err)
		internalError(r)
		return
	}

	count, err := strconv.Atoi(value.(string))
	if err != nil {
		logger.Error("Error to conver value to int", err)
		internalError(r)
		return
	}

	if count >= rule.MaxRequests {

		totalSeg, err := con.GetExp(rule.Key)
		if err != nil {
			logger.Error("Error getting TTL", err)
			internalError(r)
			return
		}

		minLeft := int(totalSeg.Minutes())
		secLeft := int(totalSeg.Seconds()) % 60

		r.Status = http.StatusOK
		r.Data.IsAllowed = false
		r.Data.Error = i18n.SuspendedService
		r.Data.Error.Message = r.Data.Error.Message + strconv.Itoa(minLeft) + i18n.Minutes + strconv.Itoa(secLeft) + i18n.Seconds
		return
	}

	// Increment and update (keep duration of ttl)
	count++
	totalSeg, err := con.GetExp(rule.Key)
	if err != nil {
		logger.Error("Error getting TTL", err)
		internalError(r)
		return
	}
	err = con.SetWithTTL(rule.Key, 0, count)
	if err != nil {
		logger.Error("Error setting the increment of value TTL", err)
		internalError(r)
		return
	}
	err = con.SetExp(rule.Key, totalSeg)
	if err != nil {
		logger.Error("Error restoring TTL", err)
		internalError(r)
		return
	}

	r.Status = http.StatusOK
	r.Data.IsAllowed = true
	r.Data.Error = i18n.NotError

}

func internalError(r *entity.Response) {
	r.Status = http.StatusInternalServerError
	r.Data.IsAllowed = false
	r.Data.Error = i18n.InternalServerError
}
