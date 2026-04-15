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
    <br>
    // code
    <br>
} else {
    <br>
    // code
    <br>
}
<br>
Key Points 
<br>
Condition must be boolean
<br>
Curly braces {} are mandatory
<br>
No parentheses required
<br>

🔁 Else-If Ladder
<br>
Syntax
<br>
if condition1 {
    <br>
} else if condition2 {
    <br>
} else {
    <br>
}
<br>
🔘 Switch Statement
<br>
Syntax
<br>
switch expression {
    <br>
case value1:
<br>
    // code
    <br>
case value2:
<br>
    // code
    <br>
default:
<br>
    // code
    <br>
}
<br>

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
    <br>
    fmt.Println(i)
    <br>
}

<br>
🔹 While-style Loop
i := 1
<br>
for i <= 5 {
    <br>
    fmt.Println(i)
    <br>
    i++
    <br>
}

<br>

🔹 Infinite Loop
for {
    <br>
    fmt.Println("Infinite Loop")
    <br>
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