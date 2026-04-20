# 📘 Module 4 - Arrays & Slices 

---

## 🧠 Introduction

Arrays and slices are used to store multiple values of the same type.

- Array → fixed size  
- Slice → dynamic size  

✔ Core data structures in Go  
✔ Used in almost every program  

---

## 📦 Arrays

### 🔹 Definition

An array is a fixed-size collection of elements of the same type.

### 🔹 Syntax

```go
var arr [5]int

🔹 Initialization
arr := [5]int{1, 2, 3, 4, 5}
🔹 Access Elements
fmt.Println(arr[0])
🔹 Traverse Array
for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}
🔹 Using Range
for index, value := range arr {
    fmt.Println(index, value)
}
⚠️ Array Limitations
Fixed size (cannot grow)
Memory is allocated at compile time
📌 Slices
🔹 Definition

A slice is a dynamic, flexible view of an array.

🔹 Syntax
s := []int{1, 2, 3}
🔹 Creating Slice from Array
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:4]
🔹 Slice Properties
Pointer → underlying array
Length → number of elements
Capacity → maximum elements
🔹 Length & Capacity
fmt.Println(len(s))
fmt.Println(cap(s))
➕ Append Function
s := []int{1, 2}
s = append(s, 3, 4)

✔ Automatically increases size

➖ Remove Element from Slice
s = append(s[:index], s[index+1:]...)
🔗 Merge Slices
s1 := []int{1, 2}
s2 := []int{3, 4}

merged := append(s1, s2...)