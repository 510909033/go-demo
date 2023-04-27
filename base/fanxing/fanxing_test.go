package main

import (
	"log"
	"testing"
)

func Test_sum(t *testing.T) {
	sum1 := sum("todo")
	sum2 := sum(33)

	t.Log(sum1, sum2)
}

type User2 struct {
}

func TestDemoCompare(t *testing.T) {
	DemoCompare(3, 4)
	DemoCompare("a", 8)

	var a1 *int
	DemoCompare(8, a1)

	var a2 *int
	DemoCompare(a1, a2)

	// var a3 interface{}
	// DemoCompare[interface{}, *int](a3, a2) //interface{} does not implement comparable

	// var a4 map[string]int
	// DemoCompare(a4, a2)//map[string]int does not implement comparable

}

type Vector[T comparable] struct {
	data_ [4]T
}

func (v *Vector[T]) Contains(e T) bool {
	for _, x := range v.data_ {
		if x == e {
			return true
		}
	}
	return false
}

func TestVector(t *testing.T) {
	_ = Vector[Vector[int]]{}
	v2 := Vector[int]{data_: [4]int{1, 2, 3, 4}}
	_ = v2
	log.Printf("%#v\n", v2)
}

type Infos interface {
	int | string | []bool
}

func getInfos[a Infos](info a) {
	log.Printf("%#v\n", info)
}

func TestGetInfos(t *testing.T) {
	getInfos(22)
	getInfos("str")
	getInfos([]bool{true, false})
	// getInfos(int16(1))//int16 does not implement Infos
	// getInfos[string](22)//cannot use 22 (untyped int constant) as string value in argument to getInfos[string]
}

type Detail[info Infos, twoInfo Infos] struct {
	Info info
	List []info
	Two  twoInfo
}

func DemoDetail(detail *Detail[string, int], list []Detail[int, []bool]) {

	log.Printf("detail, info=%s, List=%+v, two=%d\n", detail.Info, detail.List, detail.Two)

}

func TestDemoDetail(t *testing.T) {
	DemoDetail(&Detail[string, int]{
		Info: "22",
		List: []string{},
		Two:  1,
	}, []Detail[int, []bool]{})

}

type Nickname struct {
	Name string
}

func Max[T comparable, t int | int64 | Nickname](com1 T, com2 T, a, b t) t {
	if com1 == com2 {
		return a
	}
	return b
}

func TestMax(t *testing.T) {
	log.Println(Max(1, 2, 1, 2))
	log.Println(Max(1, 2, int64(1), 2))

}

type TypeDetail interface {
	Detail[int, []bool] | int
}

func Somes[T TypeDetail | int | int64 | int32](a T) T {

	log.Printf("a=%#v, info=%+v\n", a, a)

	return a
}

func TestSomes(t *testing.T) {
	Somes(Detail[int, []bool]{
		Info: 1000,
		List: []int{1, 2},
		Two:  []bool{true, false, true},
	})
	Somes[int](2)
	Somes(int32(22))
}

func SomeV2[T Detail[int, []bool]](de T) {

}

type AgeT interface {
	int8 | int16
}

type NameE interface {
	string
}

type User[T AgeT, E NameE] struct {
	age  T
	name E
}

// 获取age
func (u *User[T, E]) GetAge() T {
	return u.age
}

// 获取name
func (u *User[T, E]) GetName() E {
	return u.name
}

func TestAge(t *testing.T) {
	user := User[int16, string]{age: int16(11), name: "hh"}

	user.GetAge()
	// user.age

}

func Get[T any]() T {
	var t T

	var t2 interface{} = t
	switch t2.(type) {
	case int:
		t2 = 18
	}

	return t
}

func GetV2[T any]() T {
	var t T

	var ti interface{} = &t
	switch v := ti.(type) {
	case *int:
		*v = 18
	}

	return t
}
