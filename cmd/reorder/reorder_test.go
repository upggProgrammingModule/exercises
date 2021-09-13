package main

import (
	"reflect"
	"testing"
)

//func commaListToInts(commaList string) []int {
//func oneBasedToZeroBased(inputNumbers []int) []int {
//func reorderColumns(line string, delimiter string, columnOrder []int) string {

var commaListToIntsTests = []struct {
	commaList string
	expected  []int
}{
	{"1,2,3", []int{1, 2, 3}},
	{"3,2,1", []int{3, 2, 1}},
	{"1", []int{1}},
	{"2,7,1,3", []int{2, 7, 1, 3}},
}

var oneToZeroTests = []struct {
	input    []int
	expected []int
}{
	{[]int{1, 2, 3}, []int{0, 1, 2}},
	{[]int{5, 2, 8}, []int{4, 1, 7}},
}

var reorderTests = []struct {
	line     string
	delim    string
	order    []int
	expected string
}{
	{"hello:how:are:you", ":", []int{0, 2}, "hello:are"},
	{"duke_university_blue_devils", "_", []int{2, 3, 0}, "blue_devils_duke"},
}

func TestCommaList(t *testing.T) {
	var actual []int
	for _, curr := range commaListToIntsTests {
		actual = commaListToInts(curr.commaList)
		if !reflect.DeepEqual(actual, curr.expected) {
			t.Errorf("Error when parsing a comma list. list:%s expected:%v actual:%v", curr.commaList, curr.expected, actual)
		}
	}
}

func TestOneToZero(t *testing.T) {
	var actual []int
	for _, curr := range oneToZeroTests {
		actual = oneBasedToZeroBased(curr.input)
		if !reflect.DeepEqual(actual, curr.expected) {
			t.Errorf("Error when converting to zero based. input:%v expected:%v actual:%v", curr.input, curr.expected, actual)
		}
	}
}

func TestReorderColumns(t *testing.T) {
	var actual string
	for _, curr := range reorderTests {
		actual = reorderColumns(curr.line, curr.delim, curr.order)
		if !reflect.DeepEqual(actual, curr.expected) {
			t.Errorf("error when reordering columns. line:%v, delimiter:%v, order:%v, expected:%v, actual:%v", curr.line, curr.delim, curr.order, curr.expected, actual)}
	}
}
