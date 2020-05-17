// Скомпилируется ли код ниже
// 1 - да
// 2 - нет, ошибка на 18
// 3 - нет, ошибка на 19

package structs

import "testing"

type A struct {
	name string
}
type B struct {
	slice []interface{}
}

func TestA(t *testing.T) {
	a := make(map[A]struct{}) // 18
	b := make(map[B]struct{}) // 19
	t.Log(a)
	t.Log(b)
}
