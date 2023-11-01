package main

import "fmt"




/*

day := 3

switch day {
    case 1:
        fmt.Println("Понедельник")
        fallthrough
    case 2:
        fmt.Println("Вторник")
        fallthrough
    case 3:
        fmt.Println("Среда")
        fallthrough
    case 4:
        fmt.Println("Четверг")
        fallthrough
    case 5:
        fmt.Println("Пятница")
    default:
        fmt.Println("Неправильный день")
}

 выведет
Среда
Четверг
Пятница



switch mounth:= "jan" {
  case mounth == "Dec":
    fmt.Print("Zima")
  case mounth == "Jun":
    fmt.Print("Leto")
}


/*switch x := 2; x {
    case 1:
        fmt.Println("Один")
    case 2:
        fmt.Println("Два")
    case 3:
        fmt.Println("Три")
    default:
        fmt.Println(x)
}

*/
/*
pass = 7055
if pass == 7055 { 
    fmt.Println("Пароль верный")
}
выведет: Пароль верный

pass = 7055
if pass == 3142 { 
    fmt.Println("Пароль неверный")
}
ничего не будет выведено

/* в Go этот вариант НЕ РАБОТАЕТ!
if условие 
{
    код будет выполнен если условие равно true
}

iq := 70

if iq > 120 {
    fmt.Println("Вы гений!")
} else {
    fmt.Println("Увы, но вы не гений")
}

if условие {
    код будет выполнен если условие равно true
} else if условие2 {
    код будет выполнен если условие2 равно true
} else {
    код будет выполнен если оба условия равны false
}

*/

/*func main() {
	var  a, b, c, d int
    fmt.Scan(&a, &b, &c, &d)
	fmt.Print(a, b, c, d)
} 

/*a := 7
b := a * 2
b /= 4
a += 3
fmt.Println(b)
const pi = 3.14
a := 31
b := 14
// cложение
var result int
result = a + b
fmt.Println(result)  // выведет 45
      
// вычитание
result = a - b
fmt.Println(result)  // выведет 17
      
// умножение
result = a * b
fmt.Println(result)  // выведет 434
      
// деление
result = a / b
fmt.Println(result)  // выведет 2
      
// остаток от деления
result = a % b
fmt.Println(result)  // выведет 3
-------------
var a int = 31
var b int = 14

a += b
fmt.Println(a)  // выведет 45

a -= b
fmt.Println(a)  // выведет 17

a *= b
fmt.Println(a)  // выведет 434

a /= b
fmt.Println(a)  // выведет 2

a %= b
fmt.Println(a)  // выведет 3
-------------
var a int = 31
var b int = 14
    
// равно
fmt.Println(a == b) // выведет false
    
// не равно
fmt.Println(a != b) // выведет true
    
// больше чем
fmt.Println(a > b)  // выведет true

// больше или равно
fmt.Println(a >= b) // выведет true
    
// меньше чем
fmt.Println(a < b)  // выведет false

// меньше или равно
fmt.Println(a <= b) // выведет false
-----------
var a int = 31
var b int = 14
    
// логическое И
fmt.Println(a != b && a >= b)  // выведет true
    
// логическое ИЛИ
fmt.Println(a == b || a > b)   // выведет true
    
// логическое НЕ
fmt.Println(!(a < b))          // выведет true

*/
