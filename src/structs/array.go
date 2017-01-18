package structs

const (
    ResizeAmount = 10
)

type ArrayItem interface {
    CheckValue(interface{}) bool
    Compare(interface{}) int
    GreaterThan(interface{}) bool
}

type IntItem struct {
    Value int
}

func (item IntItem) CheckValue(val interface{}) bool {
    return val.(int) == item.Value
}

func (item IntItem) Compare(val interface{}) int {
    if val.(int) == item.Value {
        return 0
    } else if val.(int) > item.Value {
        return 1
    } else {
        return -1
    }
}

func (item IntItem) GreaterThan(comp interface{}) bool {
    return item.Value >= comp.(IntItem).Value
}

type Array struct {
    Items []ArrayItem
    Count int
}

func (arr *Array) Delete(val interface{}) {
    for i := 0; i < arr.Count; i++ {
        item := arr.Items[i]
        if item.CheckValue(val) == true {
            for j := i; j < arr.Count-1; j++ {
                arr.Items[j] = arr.Items[j+1]
            }
            arr.Count--
        }
    }
}

func (arr *Array) Find(val interface{}) int {
    for i, item := range arr.Items {
        if arr.Items[i] == nil {
            break
        } else if item.CheckValue(val) {
            return i
        }
    }
    return -1
}

func (arr *Array) Insert(item ArrayItem) {
    arr.Items[arr.Count] = item
    arr.Count++
}

func NewArray(size int) *Array {
    arr := &Array{make([]ArrayItem, size), 0}
    return arr
}

type OrderedArray struct {
    Items []ArrayItem
    Count int
    Size int
}

func (arr *OrderedArray) Delete(val interface{}) {
    for i := 0; i < arr.Count; i++ {
        for {
            item := arr.Items[i]

            if item.CheckValue(val) == true {
                for j := i; j < arr.Count-1; j++ {
                    if arr.Items[j] == nil {
                        break
                    }
                    arr.Items[j] = arr.Items[j+1]
                }
                arr.Count--
            } else {
                break
            }
        }
    }
}

func (arr *OrderedArray) Find(val interface{}) int {
    lower := 0
    upper := arr.Count - 1

    for {
        i := (lower + upper) / 2
        v := arr.Items[i].Compare(val)
        if lower > upper {
            return -1
        }
        if v == 0 {
            return i
        }
        if v == 1 {
            lower = i + 1
        }
        if v == -1 {
            upper = i - 1
        }
    }
}

func (arr *OrderedArray) GetItems() []ArrayItem {
    items := make([]ArrayItem, arr.Count)

    for i, item := range arr.Items {
        if i == arr.Count {
            break
        }
        items[i] = item
    }

    return items
}

func (arr *OrderedArray) Insert(item ArrayItem) {
    if (arr.Count + 1 > arr.Size) {
        arr.Items = append(arr.Items, make([]ArrayItem, ResizeAmount)...)
        arr.Size += ResizeAmount
    }

    for i := 0; i <= arr.Count; i++ {
        if arr.Items[i] == nil {
            arr.Items[i] = item
        }

        if arr.Items[i].GreaterThan(item) {
            for j := arr.Count; j > i; j-- {
                arr.Items[j] = arr.Items[j-1]
            }

            arr.Items[i] = item
            break
        }
    }

    arr.Count++
}

func NewOrderedArray(size int) *OrderedArray {
    arr := &OrderedArray{make([]ArrayItem, size), 0, size}
    return arr
}