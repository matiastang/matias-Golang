/*
 * @Author: matiastang
 * @Date: 2022-04-24 15:28:43
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-24 15:43:46
 * @FilePath: /matias-Golang/learn/src/app/painkiller/pill.go
 * @Description:
 */
// pill.go
package painkiller

/*
stringer并不是 Go 自带的工具，需要手动安装。可以执行下面的命令安装：
go get golang.org/x/tools/cmd/stringer
*/
//go:generate stringer -type=Pill
type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)
