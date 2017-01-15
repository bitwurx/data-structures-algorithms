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
}

func (arr *Array) Delete(val interface{}) {
    for i, item := range arr.Items {
        if arr.Items[i] == nil {
            break
        } else if item.CheckValue(val) {
            for j := i; j < arr.Count; j++ {
                arr.Items[j] = arr.Items[j+1]
            }
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