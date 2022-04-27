/*
 * @Author: matiastang
 * @Date: 2022-04-27 16:10:30
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-27 16:13:43
 * @FilePath: /matias-Golang/learn/src/test/example/example_test.go
 * @Description: example_go
 */
package example

import (
	"testing"
)

func TestAdd(t *testing.T) {
	r := add(2, 4)
	if r != 6 {
		t.Fatalf("add(2, 4) error, expect:%d, actual:%d", 6, r)
	}
	t.Logf("test add success")
}

func TestSub(t *testing.T) {
	r := sub(4, 2)
	if r != 2 {
		t.Fatalf("sub(4, 2) error, expect:%d, actual:%d", 2, r)
	}
	t.Logf("test sub success")
}
