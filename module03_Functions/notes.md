# 📘 Module 3 - Functions (Golang Notes)

---

## 🧠 Introduction

Functions are reusable blocks of code that perform a specific task.

✔ Improve code reusability  
✔ Reduce repetition  
✔ Make code modular and readable  

---

## 🔧 Function Declaration

### Syntax

```go
func functionName(parameters) returnType {
    // code
}

📞 Function Call

greet()

👉 Calls the function and executes its code

📥 Parameters

👉 Parameters are inputs to a function

func add(a int, b int) {
    fmt.Println(a + b)
}
📤 Return Values

👉 Functions can return values

func add(a int, b int) int {
    return a + b
}
🔁 Multiple Return Values

👉 Go supports returning multiple values

func calc(a, b int) (int, int) {
    return a + b, a * b
}
Usage

sum, product := calc(3, 4)

🏷 Named Return Values

func calc(a, b int) (sum int, product int) {
    sum = a + b
    product = a * b
    return
}
🔄 Recursion

👉 A function calling itself

Example

func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

Key Points

Must have base case

Avoid infinite recursion