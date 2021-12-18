package utils

import (
	"crypto/md5"
	"encoding/hex"
)

type Pagination struct {
	Limit uint `form:"limit" binding:"required"`
	Page  uint `form:"page" binding:"required"`
}

func IsNotEqualMd5(stringPassword string, stringMD5 string) bool {
	hasher := md5.New()
	hasher.Write([]byte(stringPassword))
	return hex.EncodeToString(hasher.Sum(nil)) != stringMD5
}

func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
