package godata

import (
	"sort"
)

type Float64List struct {
	data []float64
}

// Counter returns number of values. 
func (list *Float64List) Counter() map[float64]int {
	c := make(map[float64]int)
	for _, v := range list.data {
		c[v]++
	}
	return c
}

// Unique returns only unique values.
func (list *Float64List) Unique() []float64 {
	u := make([]float64, 0, len(list.data))
	isExist := make(map[float64]bool)
	for _, value := range list.data {
		if !isExist[value] {
			u = append(u, value)
			isExist[value] = true
		}
	}
	return u
}

// Duplicates returns a new Float64ListDuplicates reading from list.
func (list *Float64List) Duplicates() Float64ListDuplicates {
	c := list.Counter()
	ds := make(Float64ListDuplicates, 0, len(c))
	for _, v := range list.Unique() {
		if c[v] > 1 {
			ds = append(ds, Float64ListDuplicate{v, c[v]})
		}
	}
	return ds
}

func NewFloat64List(data []float64) *Float64List {
	return &Float64List{
		data: data,
	}
}

type Float64ListDuplicate struct {
	Value float64
	Count int
}

type Float64ListDuplicates []Float64ListDuplicate

func (ds Float64ListDuplicates) SortByValue() {
	sort.Slice(ds, func(i, j int) bool { 
		return ds[i].Value < ds[j].Value 
	})
}

func (ds Float64ListDuplicates) SortByValueReverse() {
	sort.Slice(ds, func(i, j int) bool { 
		return ds[i].Value > ds[j].Value 
	})	
}

func (ds Float64ListDuplicates) SortByCount() {
	sort.Slice(ds, func(i, j int) bool { 
		return ds[i].Count < ds[j].Count 
	})	
}

func (ds Float64ListDuplicates) SortByCountReverse() {
	sort.Slice(ds, func(i, j int) bool { 
		return ds[i].Count > ds[j].Count 
	})		
}

func (ds Float64ListDuplicates) Values() []float64 {
	vs := make([]float64, len(ds))
	for i := 0; i < len(ds); i++ {
		vs[i] = ds[i].Value
	}
	return vs
}

func (ds Float64ListDuplicates) Counts() []int {
	cs := make([]int, len(ds))
	for i := 0; i < len(ds); i++ {
		cs[i] = ds[i].Count
	}
	return cs
}

type StringList struct {
	data []string
}

// Counter returns number of values. 
func (list *StringList) Counter() map[string]int {
	c := make(map[string]int)
	for _, v := range list.data {
		c[v]++
	}
	return c
}

// Unique returns only unique values.
func (list *StringList) Unique() []string {
	u := make([]string, 0, len(list.data))
	isExist := make(map[string]bool)
	for _, value := range list.data {
		if !isExist[value] {
			u = append(u, value)
			isExist[value] = true
		}
	}
	return u
}

func NewStringList(data []string) *StringList {
	return &StringList{
		data: data,
	}
}