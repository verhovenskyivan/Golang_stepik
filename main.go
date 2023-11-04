package main

import "fmt"

func double(a int) {
    fmt.Println(a*2) 
} 

func main() {
    for i := 4; i > 0; i-- {
        defer double(i)
    }
}

Вам нужно написать для этой программы функцию marsAge(), которая будет принимать на вход целое число, возраст в земных годах, а возвращать возраст на Марсе. 
Возвращать функция должна целое число. Округлять ничего не надо. Просто отбрасываем дробную часть.
func marsAge(a int) int {
    return a * 365 / 687
}


Напишите функцию squareSumm(), которая должна принимать на вход два целых числа a и b и возвращать сумму квадратов чисел от a до b (включительно).
func squareSumm(a,b int) int {
    sum := 0
    for i:=a; i <= b; i++ {        
    sum+= i*i
    }
    return sum
}


Напишите функцию isEven(), которая принимает в качестве аргумента одно целое число и возвращает true если оно четное и false в ином случае.

func isEven(a int) bool {
    return a % 2 == 0
}

Напишите функцию calc() которая должна принимать на вход одно целое число, а возвращать два значения - число умноженное на два и число возведенное в квадрат.

func calc(a int)( int, int) {
    return a*2, a*a   
   }
   
------------
Напишите функцию max() которая должна принимать в качестве параметров два целых числа и возвращать большее из них.
func max(a,b int)  int {
    var x int
    switch {
    case a > b:
    return a
    case b > a:
    return b
  }
    return x
}


var x = 13

func one() {
    x += 1
}

func two(x int) {
    x += 3
}

func main() {
    one()
    two(x)
    fmt.Println(x)
}
/*

func main() {
    helloStepik()
}

func helloStepik() {
    fmt.Println("Hello stepik")
}
func plusminus(a, b int) (int, int){
    return a + b , a - b
}
func concat(a, b string) string{
    return a + b   
   }

Defer
Оператор defer гарантирует, что функция будет вызвана только после того, как вернется окружающая функция.

Например:

package main

import "fmt"

func welcome() {
    fmt.Println("Добро пожаловать")
}

func main() {
    defer welcome()
    fmt.Println("Привет!")
}
Приведенный выше код сначала выведет Привет! и только после этого выведет результат функции welcome().
Это происходит потому, что вызов функции welcome() отложен, то есть она ждет, пока main() закончит выполнение, и только после этого происходит её вызов
/*
package main

import "fmt"

func square(num int) int {
    return num * num
}

func main() {
    res := square(4)
    fmt.Println(res) // выведет: 16
}
func argFunc(a int, b int, c int){
    какой-то код
}

или так, если все они одного типа:
func argFunc(a, b, c int){
    как
func main(){
    var res int
    fmt.Scan(&res)
    sum := 1
    for  i := 1; i <= res; i++ {
    sum *= i
    }
}

/*
    fmt.Print(sum)
    sum := 0
    for i := 0; i <= 30; i += 5 {
        sum += i
    } 
    fmt.Println(sum)

{
    var money int
    fmt.Scan(&money)
    switch  {
        case money > 1000:
        fmt.Print("Apple")
        case 500 <= money && money <= 1000:
        fmt.Print("Samsung")
        case money < 500:
        fmt.Print("Nokia с фонариком")
    }
for [инициализация счетчика]; [условие]; [изменение счетчика]{
    // это так называемое тело цикла, 
    // где находится код, который выполняется во время работы цикла
}

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