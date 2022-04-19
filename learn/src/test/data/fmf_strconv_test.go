/*
 * @Author: matiastang
 * @Date: 2022-04-19 18:52:13
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-19 19:05:19
 * @FilePath: /matias-Golang/learn/src/test/data/fmf_strconv_test.go
 * @Description: fmf转换
 */
package data

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

/*
* go test -v -bench=. fmf_strconv_test.go
goos: darwin
goarch: arm64
Benchmark_fmf_sprint_test
5577006791947779410
4011359550169803385
5274308539948715278
7127531329315789732
3616463329458536717
Benchmark_fmf_sprint_test-8             11428642               100.7 ns/op
Benchmark_strconv_Itoa_test
6503432924651047465
6125781018688046522
3175377706831416265
6299558379657647971
7728232659349344884
Benchmark_strconv_Itoa_test-8           24395553                46.62 ns/op
PASS
ok      command-line-arguments  2.916s
* 基本数据类型与字符串之间的转换，优先使用 strconv 而不是 fmt，因为前者性能更佳。
* fmt 实现上利用反射来达到范型的效果，在运行时进行类型的动态判断，所以带来了一定的性能损耗。
*/
func Benchmark_fmf_sprint_test(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = fmt.Sprint(rand.Int())
	}
	fmt.Println(s)
}

func Benchmark_strconv_itoa_test(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = strconv.Itoa(rand.Int())
	}
	fmt.Println(s)
}
