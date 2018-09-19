Data manipulation: counter, uniqueness, duplicates.


# Data structure

## List
Manipulation of the data list structure.

### Operation

#### Counter
Counter returns number of values. 

##### []float64
```
data := []float64{0.1, 0.5, 10, 100500.41, 0.1}
list := godata.NewFloat64List(data)
c := list.Counter()
fmt.Println(c)

// map[0.1:2 0.5:1 10:1 100500.41:1]
```

##### []string

```
data := []string{"Hi-ho", "Karlsson", "said", "Eric", "Where", "did", "you", "get", "to", "Why", "What", "do", "you", "mean", "asked", "Karlsson"}
list := godata.NewStringList(data)
c := list.Counter()
fmt.Println(c)

// map[said:1 to:1 mean:1 Karlsson:2 Eric:1 did:1 Why:1 What:1 Hi-ho:1 Where:1 you:2 get:1 do:1 asked:1]
```


#### Unique
Unique returns only unique values.

##### []float64

```
data := []float64{0.1, 0.5, 10, 100500.41, 0.1}
list := godata.NewFloat64List(data)
u := list.Unique()
fmt.Println(u)

// [0.1 0.5 10 100500.41]
```

##### []string

```
data := []string{"Hi-ho", "Karlsson", "said", "Eric", "Where", "did", "you", "get", "to", "Why", "What", "do", "you", "mean", "asked", "Karlsson"}
list := godata.NewStringList(data)
u := list.Unique()
fmt.Println(u)

// [Hi-ho Karlsson said Eric Where did you get to Why What do mean asked]
```

#### Duplicates
Duplicates returns a new Float64ListDuplicates reading from list. 


##### []float64

```
data := []float64{0.1, 0.5, 6.1, 10, 6.1, 100500.41, 0.1, 6.1, 6.2, 7.71}
list := godata.NewFloat64List(data)
ds := list.Duplicates()
fmt.Printf("%#v\n", ds)

// godata.Float64ListDuplicates{godata.Float64ListDuplicate{Value:0.1, Count:2}, godata.Float64ListDuplicate{Value:6.1, Count:3}}

```

Values 0.5, 10, 100500.41, 6.2, 7.71 aren't available, because aren't duplicates.

Duplicates can be sorted by Value and Count:

```

ds.SortByValue()
fmt.Printf("%#v\n", ds)

// godata.Float64ListDuplicates{godata.Float64ListDuplicate{Value:0.1, Count:2}, godata.Float64ListDuplicate{Value:6.1, Count:3}}

ds.SortByValueReverse()
fmt.Printf("%#v\n", ds)

// godata.Float64ListDuplicates{godata.Float64ListDuplicate{Value:6.1, Count:3}, godata.Float64ListDuplicate{Value:0.1, Count:2}}


ds.SortByCount()
fmt.Printf("%#v\n", ds)

// godata.Float64ListDuplicates{godata.Float64ListDuplicate{Value:0.1, Count:2}, godata.Float64ListDuplicate{Value:6.1, Count:3}}

ds.SortByCountReverse()
fmt.Printf("%#v\n", ds)

// godata.Float64ListDuplicates{godata.Float64ListDuplicate{Value:6.1, Count:3}, godata.Float64ListDuplicate{Value:0.1, Count:2}}

```

From duplicates, you can get a list of values and counts.

```

vs := ds.Values()
fmt.Println(vs)
// [0.1 6.1]

cs := ds.Counts()
fmt.Println(cs)
// [2 3]

```


# TODO

* Add more data type in List
* Add new data structure