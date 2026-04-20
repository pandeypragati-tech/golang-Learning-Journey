# 📘 Module 5 - Strings 

---

## 🧠 Introduction

A string in Go is a sequence of bytes representing text (UTF-8 encoded).

✔ Immutable (cannot be changed directly)  
✔ Used for text processing  
✔ Very important for real-world applications  

---

## 🔤 String Basics

### 🔹 Declaration

```go
var str string = "Hello"

str := "Golang"

🔹 Access Characters
fmt.Println(str[0]) // byte value

⚠️ Returns ASCII/byte, not character

🔹 String Length
len(str)
🔁 Iterating Over String
🔹 Using Index
for i := 0; i < len(str); i++ {
    fmt.Println(string(str[i]))
}
🔹 Using Range (Recommended)
for _, ch := range str {
    fmt.Println(string(ch))
}

✔ Handles Unicode properly

🔄 String Immutability

❌ Cannot modify directly:

str[0] = 'H' // error

✔ Workaround:

Convert to slice
Modify
Convert back
🔧 Common String Operations
🔹 Concatenation
s1 := "Go"
s2 := "Lang"
result := s1 + s2
🔹 Comparison
a == b
🔹 Substring (Slicing)
str := "Golang"
sub := str[0:3]
📦 strings Package (Important)
🔹 To Upper / Lower
strings.ToUpper(str)
strings.ToLower(str)
🔹 Contains
strings.Contains(str, "go")
🔹 Replace
strings.ReplaceAll(str, " ", "")
🔹 Split
strings.Split(str, " ")
🔹 Fields (Word Split)
words := strings.Fields(str)