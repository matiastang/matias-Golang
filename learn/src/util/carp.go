// /*
// * @Author: tangdaoyong
// * @Date: 2021-02-01 15:21:09
// * @LastEditors: tangdaoyong
// * @LastEditTime: 2021-02-01 15:25:12
// * @Description: CARP 算法
// */
package carp

import (
	"crypto/md5"
	"errors"
)

var (
	ErrHaveNoRedis = errors.New("have no redis")
)

// 根据 carp 算法, 从 redisArr 的数组中选择合适的 redis, 返回值为下标
func GetTargetIndex(key string, redisArr []string) (idx int, err error) {
	if len(redisArr) < 1 {
		return -1, ErrHaveNoTarget
	} else if len(redisArr) == 1 {
		return 0, nil
	}

	hashArr := make([]string, len(redisArr))
	for k, v := range redisArr {
		hashArr[k] = hashString(v + key)
	}
	idx = minIdx(hashArr)
	return
}

// 返回 arr 数组中最小值的下标
func minIdx(arr []string) (idx int) {
	if len(arr) < 1 {
		return -1
	}

	idx, min := 0, arr[0]

	for k, v := range arr {
		if v < min {
			idx = k
			min = v
		}
	}
	return
}

func hashString(str string) (hash string) {
	md5sum := md5.Sum([]byte(str))
	return string(md5sum[:])
}