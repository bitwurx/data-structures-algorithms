package structs

type ArrayItem interface {
    CheckValue(val interface{}) bool
}

type IntItem struct {
    Value int
}

func (item IntItem) CheckValue(val interface{}) bool {
    return val.(int) == item.Value
}

type Array struct {
    Items []ArrayItem
    Count int
    Size int
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
    arr := &Array{make([]ArrayItem, size), 0, size}
    return arr
}