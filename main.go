package main

import "fmt"

Диапазон (range)
Теперь, когда мы знаем, как создавать массивы и срезы, давайте научимся выполнять перебор их элементов с помощью цикла!

Форма range (диапазон) цикла for позволяет выполнять итерации по срезу:

package main 

import "fmt"

func main() {
    a := make([]int, 5)
    a[1] = 2
    a[2] = 3

    for i, v := range a {
        fmt.Print(i,v)
    }
}

Во время каждой итерации цикла он возвращает два значения: индекс элемента и его значение.

Первым указывается переменнная для индекса, вторым для значений среза/массива

Если вам нужны только значения, вы можете пропустить индекс, используя знак подчеркивания:

for _, v := range a {
    fmt.Println(v)
  }


  Вы можете использовать диапазоны как для срезов, так и для массивов.


  Есть массив nums. Вам нужно написать цикл для перебора его значений с использованием range.
    В цикле нужно посчитать сумму всех элементов массива. Для этого объявлена переменная sum.

    sum := 0

    for _, v := range nums {
        sum += v
    }

    res := 0
    nums := [5]int{1, 2, 3, 4, 5}
    for i, v := range nums {
        if i % 2 != 0 {
            res += v
        }
    }
    fmt.Println(res)
Выведет 6, так как условие проходится по индексам, а именно если индекс элемента массива не является четным то прибавляется к res, 
следовательно индекс 1 имеет цифра 2, и индекс 3 имеет число 4
4+2 =6


Срезы (slices)
Массив имеет фиксированный размер, то есть после его определения вы не можете изменить количество содержащихся в нем элементов.

Для преодоления этой проблемы Go предоставляет срез, который представляет собой динамически изменяемое представление элементов массива.

Срез основан на массиве и определяется путем указания двух индексов, нижней и верхней границы, разделенных двоеточием:

a := [5]int{0, 2, 4, 6, 8}

var s []int = a[1:3]
Этот код выбирает из массива a элементы с индексами от 1 до 3, включая первый заданный индекс, но исключая последний.

Таким образом, срез s теперь будет включать значения [2, 4]:

fmt.Println(s) // выведет [2 4]
Вы можете опустить нижнюю или верхнюю границу.
Если опустить нижнюю границу, то будет взято значение 0, а если опустить верхнюю границу, то будет взята длина массива.
Например: a[:3] возьмет первые 3 элемента массива.
Вы можете получить доступ к значениям среза, используя тот же синтаксис, что и при работе с массивами:

a := [5]int{0, 2, 4, 6, 8}
var s []int = a[1:3]

fmt.Println(s[1]) // выведет 4
В срезе не хранятся никакие данные; он просто описывает секцию исходного массива. Изменение элементов среза изменяет соответствующие элементы исходного массива.

package main

import "fmt"

func main() {
  a := [5]int{0, 2, 4, 6, 8}
  var s []int = a[1:3]

  s[0] = 8
  fmt.Println(a) // выведет [0 8 4 6 8]
}
Вы можете иметь несколько срезов одного и того же массива. Изменение в любом из них будет видно во всех срезах, так как оно повлияет на основной массив.
x := [6]int{4, 3, 7, 2, 1, 9}
var y []int = x[2:5]
fmt.Println(y[2])
выведет 1, так как изначально объявлен массив x длинной в 6 целых чисел, далее объявляется слайс y который берет отсчет с третьего элемента массива (число 7) до четвертого элемента
так как правое число не учитывается (а именно в этом случае 5),  далее мы вызываем 2 элемент по слайсу у, соотственно 0 индекс = 7, 1 индекс = 2 и 2 индекс равен 1

Go предоставляет функцию make() для создания срезов.
Так создаются массивы динамического размера.

Например:

a := make([]int, 5)
Функция make создает массив заданного типа и размера и возвращает срез, который ссылается на этот массив.

После создания среза мы можем добавлять в него новые элементы с помощью функции append():

a := make([]int, 5)
a = append(a, 8)
fmt.Println(a) // выведет [0 0 0 0 0 8]
Функция append() принимает в качестве первого аргумента срез, а в качестве следующего аргумента - элементы, которые должны быть добавлены в конец среза.
Затем она возвращает новый срез, содержащий старый срез плюс новые добавленные элементы.

Вы можете добавлять несколько значений одновременно, просто разделяя их запятыми в качестве аргументов, например: append(s, 1, 2, 3)

x := make([]int, 2)
x = append(x, 4, 2, 1)
fmt.Println(x[3])
создается слайс x, состоящий из двух целых чисел [0,0], далее добавляем в конец числа, выходит [0, 0, 4, 2, 1]
cоответственно элементов с индексом 3 является число 2

y := make([]int, 4)
y = append(y, 4)
fmt.Println(y[3])

создается слайс y, состоящий из четырех целых чисел [0, 0, 0, 0], далее в конец добавляем 4 через, выходит [0, 0, 0, 0, 4]
выведет 0

----------------------------------------------------------------    
Массив - это последовательность элементов одного типа.

Массив задается с помощью квадратных скобок, число в которых определяет количество элементов в массиве. После скобок указывает тип значений, которые будут хранится в массиве.

Например:

var a [5]int
Теперь a - это массив из 5 целых чисел.

Вы также можете определить и сразу инициализировать значения массива, используя следующий синтаксис:

a := [5]int{0, 2, 4, 6, 8} 
Как вы видите, при объявлении массива нам необходимо указывать его размер. Это означает, что массивы имеют фиксированный размер.
Создайте массив countries, который может хранить три строковые значения
var countries [3]string
После объявления массива мы можем получить доступ к его элементам, используя квадратные скобки и их индексы.

Например:

var a [5]int

a[0] = 8
a[1] = 42

fmt.Println(a[1]) // выведет 42

---------------------------------------------------------------------------------------------------------------------------------------
Структура — это тип данных, который содержит именованные поля. Это позволяет группировать различные данные.
var countries [3]string


Например, давайте создадим структуру для хранения данных Contacts:

type Contact struct {
    name string
    age  int
} 
Теперь мы можем создать новый экземпляр нашей структуры Contact, используя следующий синтаксис:

x := Contact{"Андрей", 33}
x теперь является структурой, которая инициализируется с данными, которые указаны в фигурных скобках.

Мы также можем указать имена полей при создании нового экземпляра структуры.
Например:
x := Contact{name: "Андрей", age: 33}.
Это облегчает чтение кода.

Вам нужно создать структуру Car, которая должна иметь три поля: color и brand строчного типа, а также year целочисленного.

type Car struct {
    color string 
    brand string
    year int
}

В результате выполнения следующего фрагмента кода будет ответ 2
type Coord struct {
    x int
    y int
}

func main() {
    a := Coord{7, 5}
    fmt.Println(a.x - a.y)
}
Так как a становится частью структуры со значениями 7 по x и 5 по y, далее пишем разность 7-5

fmt.Println(x.age)  // выведет 27
fmt.Println(x.name) // выведет Андрей


Мы можем получить доступ к полям структуры, используя её имя и точку:
x :=Contact{"Андрей, 33"}

x.age = 27

Указатели на структуры
Подобно простым указателям, мы также можем создавать указатели на структуры с помощью оператора &:

x := Contact{"Андрей", 33}
p := &x
Указатели на структуры автоматически разыменовываются, что означает, что мы можем получить доступ к значениям полей, просто используя точку:

x := Contact{"Андрей", 33}
p := &x
fmt.Println(p.age)
Мы могли бы использовать (*p).age для доступа к полю age в структуре, но это выглядит сложно и трудночитаемо. 
Go позволяет сократить этот синтаксис и просто использовать вместо него p.age

Заполните пропуски, чтобы сделать указатель на структуру Car и вывести ее поле name, используя указатель:
a := Car{"Мерседес", "красный", 2022}
p := &a
fmt.Println(p.name)

Мы также можем использовать указатели при создании новой структуры:

p := &Contact{"Андрей", 22}
Теперь p - это указатель на только что созданный экземпляр структуры Contact.

Указатели на структуры полезны, так как позволяют передавать их в функции в качестве аргументов.

Заполните пропуски, чтобы определить структуру Dog, создать указатель на только что созданный экземпляр структуры Dog и вывести значение поля age этого экземпляра.

type Dog struct {
    name string
    age int
}

func main () {
    x:=&Dog {"Шарик",7}
    fmt.Println(x.age)
}
Методы
Мы можем добавить функциональность нашим структурам с помощью методов!
Методы - это просто функции со специальным аргументом-получателем.
Давайте рассмотрим пример:

func (x Contact) welcome() {
    fmt.Println(x.name)
    fmt.Println(x.age)
} 
Получатель указывается между ключевым словом func и именем метода.
В приведенном выше примере приемником является структура Contact.

Обратите внимание, что в методе мы можем получить доступ к полям структуры получателя.

После определения метода мы можем вызвать его на нашей структуре, используя синтаксис точки:

package main

import "fmt"

type Contact struct {
    name string
    age int
}

func (x Contact) welcome() {
    fmt.Println(x.name)
    fmt.Println(x.age)
}

func main() {
    x := Contact{"Андрей", 33}
    x.welcome()
}
Поскольку методы - это просто функции с аргументом-получателем, мы можем достичь той же функциональности, используя обычную функцию, которая принимает структуру в качестве аргумента:

func welcome(x Contact) {
    fmt.Println(x.name)
    fmt.Println(x.age)
}
func main() {
    x := Contact{"Андрей", 33}
    welcome(x)
}
В приведенном выше коде welcome() - это функция, принимающая в качестве аргумента структуру.

В результате данного кода type Test struct {
    a int
    b int
}

func (x Test) do() int {
    return x.a - x.b
}

func main() {
    t := Test{11, 2}
    fmt.Println(t.do())
}
Будет выведено 9, так как структура Test была объявлена методов фукнции Do, 
следовательно эта функция возвращала разницу между первым и вторым числом, переданным в нее как в структуру
Далее через мейн мы объявили новый экземпляр структуры t и присвоили ему значения  11 к а и 2 к b, далее вызвали вывод функции do, следовательно было возвращено 11-2


Если нам нужно изменить данные структуры в методе, мы можем использовать указатели в качестве получателей метода.
Например:

func (ptr *Contact) increase(val int) {
    ptr.age += val
}
Теперь наш метод increase() использует указатель в качестве получателя и может изменять поле age в структуре, на которую он вызван:

x := Contact{"Андрей", 33}
x.increase(5)
fmt.Println(x.age) // выведет 38
Go автоматически разыменовывает указатели, поэтому мы просто вызываем метод с именем структуры, как и в случае с простым получателем.

Поскольку методы часто нуждаются в модификации своего получателя, получатели указателей встречаются чаще, чем получатели значений.

Заполните пробелы, чтобы создать метод под названием drive, который принимает указатель-получатель c на структуру Car, а также целочисленный аргумент speed.

func (c *Car) drive (speed int) {
    return c.Speed += speed // это доп
}

Классическая задача на указатели. 

Вам надо написать функцию changeStrings, которая должна принимать указатели на две строки и затем менять местами их значения.

Выводить или возвращать ничего не нужно.

func changeStrings(a,b *string) {
    *a,*b = *b,*a
}
чтобы объявить метод под названием deposit для структуры Account, принимающий один целочисленный аргумент (значение в структуре в итоге должно измениться).

func (acc *Account) deposit(x int){
    acc.balance += x
}
Что выведет следующий фрагмент кода:

type T struct {
    val int
}
func (p *T) a() {
    p.val += 1
}
func (p T) b() {
    p.val += 2
}
func main() {
    x := T{5}
    x.a()
    x.b()
    fmt.Println(x.val)
}
Выведет 6, так как через x мы передали в структуру значение 5, в методе b нет указателя на струкутуру *, следовательно к экземпляру структуры val прибавляется значение 1 как в функции a
так как она имеет указатель на структуру 

--------------------------------------------------------------------------------------------------------------------------------
func double(a int) {
    a := 10
    p := &a
    a -= 4
    *p = *p + 3
    fmt.Println(a)
} 
Вам нужно написать функцию doublePointer, которая должна принимать в качестве аргумента указатель на целочисленую переменную x, а затем увеличить значение этой переменной в два раза. 
Значение должно измениться не только в переменной внутри функции, но и за ее пределами. Именно для этого вам надо использовать указатель.
func doublePointer(x *int) {
    *x *= 2   
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