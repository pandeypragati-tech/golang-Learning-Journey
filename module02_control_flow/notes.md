# 📘 Module 2 - Control Flow 

## 🧠 Introduction

Control Flow defines the order in which statements are executed in a program.

It helps in:
- Decision making  
- Repeating tasks  
- Building logic for real-world problems  

## 🔀 If-Else Statement

### Syntax

if condition {
    // code
} else {
    // code
}
Key Points 
<br>
Condition must be boolean
<br>
Curly braces {} are mandatory
<br>
No parentheses required

🔁 Else-If Ladder
<br>
Syntax
<br>
if condition1 {
} else if condition2 {
} else {
}
🔘 Switch Statement
<br>
Syntax
<br>
switch expression {
case value1:
    // code
case value2:
    // code
default:
    // code
}

Key Points
<br>
Cleaner than multiple if-else
<br>
No need of break (automatic)
<br>
Can match multiple values
<br>

🔄 Loops (for)
<br>

👉 Go has only one loop: for
<br>

🔹 Basic Loop
<br>
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
<br>
🔹 While-style Loop
i := 1
for i <= 5 {
    fmt.Println(i)
    i++
}
<br>

🔹 Infinite Loop
for {
    fmt.Println("Infinite Loop")
}
<br>
🎯 🔹 What is break?
<br>

👉 break immediately stops the loop and exits it
<br>
Simple Meaning
<br>
break = exit the loop completely
<br>
🎯 🔹 What is continue?
<br>

👉 continue skips the current iteration and moves to the next one
<br>

🧠 Simple Meaning
<br>
continue = skip this step, go to next iteration
<br>