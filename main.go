package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

var pl = fmt.Println

// ===================================== Input ==================================== //
func input() {
	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	}
}

// ===================================== Variable ==================================== //
func variable() {
	// var name type
	var vName string = "Petch"
	var v1, v2 = 1.2, 3.4
	var v3 = "hello"

	v4 := 2.4
	v4 = 5.4

	pl(vName, v1, v2, v3, v4)
}

// ===================================== Data Types ==================================== //
func dataTypes() {
	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.14))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("Hello"))
	pl(reflect.TypeOf("🥱"))
}

// ===================================== Casting ==================================== //
func casting() {
	cV1 := 1.5
	cV2 := int(cV1)
	pl(cV2)

	cV3 := "5000000"
	cV4, err := strconv.Atoi(cV3)
	pl(cV4, err, reflect.TypeOf(cV4))

	cV5 := 5000000
	cV6 := strconv.Itoa(cV5)
	pl(cV6, reflect.TypeOf(cV6))

	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		pl(cV8)
	}

	cV9 := fmt.Sprintf("%f", 3.14)
	pl(cV9)
}

// ===================================== IF Conditional ==================================== //
func ifConditional() {
	// conditional Operations : > < >= <= == !=
	// Logical Operations : && || !

	iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		pl("Important Birthday")
	} else if (iAge == 21) || (iAge == 50) {
		pl("Important Birthday")
	} else if iAge >= 65 {
		pl("Important Birthday")
	} else {
		pl("Not an Important Birthday")
	}

	pl("!true", !true)
}

// ===================================== String ==================================== //
func stringsFunc() {
	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")

	sV2 := replacer.Replace(sV1)
	pl(sV2)
	pl("Length :", len(sV2))
	pl("Contains Another :", strings.Contains(sV2, "Another"))
	pl("o index :", strings.Index(sV2, "o"))
	pl("Replace :", strings.Replace(sV2, "o", "0", -1))

	sV3 := "\nSome Words\n"
	sV3 = strings.TrimSpace(sV3)
	pl("Split :", strings.Split("a-b-c-d", "-"))
	pl("Lower :", strings.ToLower(sV2))
	pl("Upper :", strings.ToUpper(sV2))
	pl("Prefix :", strings.HasPrefix("tacocat", "taco"))
	pl("Suffix :", strings.HasSuffix("tacocat", "cat"))
}

// ===================================== Rune ==================================== //
func runes() {
	rStr := "abcdefg"
	pl("Rune Count :", utf8.RuneCountInString(rStr))

	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}
}

// ===================================== Time ==================================== //
func timeFunc() {
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())
}

// ===================================== Math ==================================== //
func mathFunc() {
	pl("5 + 1 =", 5+4)
	pl("5 - 1 =", 5-4)
	pl("5 * 1 =", 5*4)
	pl("5 / 1 =", 5/4)
	pl("5 % 1 =", 5%4)
	mInt := 1
	mInt = mInt + 1 // mInt += 1, mInt++

	pl("Float Precision", 0.11111111111111+0.11111111111111)

	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	pl("Random :", randNum)

	// There are many match functions
	pl("Abs(10) =", math.Abs(-10))
	pl("Pow(4,2) =", math.Pow(4, 2))
	pl("Sqrt(6) =", math.Sqrt(16))
	pl("Cbrt(8) =", math.Cbrt(8))
	pl("Ceil(4.4) =", math.Ceil(4.4))
	pl("Floor(4.4) =", math.Floor(4.4))
	pl("Round(4.4) =", math.Round(4.4))
	pl("Log2(8) =", math.Log2(8))
	pl("Log10(100) =", math.Log10(100))
	// Get the log of e to the power 0 2
	pl("Log(7.389) =", math.Log(math.Exp(2)))
	pl("Max(5, 4) =", math.Max(5, 4))
	pl("Min(5, 4) =", math.Min(5, 4))
	// Convert 90 degrees to radians
	r90 := 90 * math.Pi / 180
	d90 := r90 * (180 / math.Pi)
	fmt.Printf("%.2f radians = %.2f degrees\n", r90, d90)
	pl("Sin(90) =", math.Sin(r90))

	// %d : Integer
	// %c : Character
	// %f : Float
	// %t : Boolean
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses based on data type
	// %T : Type of supplied value

	fmt.Printf("%s %d %c %f %t %o %x\n", "Stuff", 1, 'A', 3.14, true, 1, 1)
	fmt.Printf("%9f\n", 3.14)
	fmt.Printf("%.2f\n", 3.141592)
	fmt.Printf("%9.f\n", 3.141592)

	sp1 := fmt.Sprintf("%9.f\n", 3.141592)
	pl(sp1)
}

// ===================================== For Loop ==================================== //
func forLoop() {
	for x := 1; x <= 5; x++ {
		pl(x)
	}
	pl("----------------------")
	for x := 5; x >= 1; x-- {
		pl(x)
	}
	pl("----------------------")
	fX := 0
	for fX < 5 {
		pl(fX)
		fX++
	}
}

// ===================================== While Loop ==================================== //
func whileLoop() {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	for true {
		fmt.Print("Guess a number between 0 and 50 :")
		pl("Random Number is :", randNum)

		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if iGuess > randNum {
			pl("Pick a Lower Value")
		} else if iGuess < randNum {
			pl("Pick a Higher Value")
		} else {
			pl("You Guess it")
			break
		}
	}
}

// ===================================== Range ==================================== //
func rangeFuc() {
	aNums := []int{1, 2, 3}
	for _, num := range aNums {
		pl(num)
	}
}

// ===================================== Array ==================================== //
func arrayFuc() {
	var arr1 [5]int
	arr1[0] = 1

	arr2 := [5]int{1, 2, 3, 4, 5}

	pl("Index 0 :", arr2[0])
	pl("Array Length", len(arr2))

	for i := 0; i < len(arr2); i++ {
		pl(arr2[i])
	}

	for i, v := range arr2 {
		fmt.Printf("%d : %d\n", i, v)
	}

	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			pl(arr3[i][j])
		}
	}

	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		fmt.Printf("Rune Array : %d\n", v)
	}
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	pl("I'm a string :", bStr)
}

// ===================================== Slice ==================================== //
func sliceFunc() {
	// var name []datatype
	sl1 := make([]string, 6)
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"
	pl("Slice Size :", len(sl1))

	for i := 0; i < len(sl1); i++ {
		pl(sl1[i])
	}
	for _, x := range sl1 {
		pl(x)
	}
	sArr := [5]int{1, 2, 3, 4, 5}
	sl3 := sArr[0:2]
	pl(sl3)
	pl("1st 3 :", sArr[:3])
	pl("Last 3 :", sArr[2:])
	sArr[0] = 10
	pl("sl3 :", sl3)
	sl3[0] = 1
	pl("sArr :", sArr)

	sl3 = append(sl3, 12)
	pl("sl3 :", sl3)
	pl("sArr :", sArr)

	sl4 := make([]string, 6)
	pl("sl4 :", sl4)
	pl("sl4[0] :", sl4[0])
}

// ===================================== Function ==================================== //
func sayHello() {
	pl("Hello")
}
func getSum(x, y int) int {
	return x + y
}
func getTwo(x int) (int, int) {
	return x + 1, x + 2
}
func getQuotient(x, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You can't divide by zero")
	} else {
		return x / y, nil
	}
}
func getSum2(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
func getArraySum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}
func changeVal(f3 int) int {
	f3 += 1
	return f3
}
func function() {
	// func funcName(parameters) returnType {BODY}
	sayHello()
	pl(getSum(5, 4))
	pl(getTwo(5))
	// return with error
	pl(getQuotient(5, 0))
	pl(getQuotient(5, 4))
	// varadic function
	pl(getSum2(1, 2, 3, 4, 5))
	// passing array
	vArr := []int{1, 2, 3, 4}
	pl("Array Sum :", getArraySum(vArr))

	f3 := 5
	pl("f3 before func :", f3)
	changeVal(f3)
	pl("f3 after func :", f3)
}

// ===================================== Pointer ==================================== //
func changeVal2(myPtr *int) {
	*myPtr = 12
}
func pointerFunc() {
	f3 := 10
	pl("f3 before func :", f3)
	changeVal2(&f3)
	pl("f3 after func :", f3)

	f4 := 10
	var f4Ptr *int = &f4
	pl("f4 Address :", f4Ptr)
	pl("f4 Value :", *f4Ptr)
	*f4Ptr = 11
	pl("f4 Value :", *f4Ptr)

	pl("f4 before func :", f4)
	changeVal2(&f4)
	pl("f4 after func :", f4)
}

// ===================================== Array Pointer ==================================== //

func dblArrVals(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}
func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))
	for _, val := range nums {
		sum += val
	}
	return (sum / numSize)
}
func arrayPointerFunc() {
	pArr := [4]int{1, 2, 3, 4}
	dblArrVals(&pArr)
	pl(pArr)

	iSlice := []float64{11, 13, 17}
	fmt.Printf("Average : %.3f\n", getAverage(iSlice...))
}

// ===================================== File IO ==================================== //
func fileIOFunc() {
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	iPrimeArr := []int{2, 3, 5, 7, 11}
	var sPrimeArr []string
	for _, i := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}
	for _, num := range sPrimeArr {
		_, err := f.WriteString(num + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scan1 := bufio.NewScanner(f)
	for scan1.Scan() {
		pl("Prime :", scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat("data.txt")
	if errors.Is(err, os.ErrNotExist) {
		pl("File Doesn't Exist")
	} else {
		f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if _, err := f.WriteString("13\n"); err != nil {
			log.Fatal(err)
		}
	}
}

// ===================================== Map ==================================== //
func mapFunc() {
	// var myMap map [keyType]valueType
	var heroes map[string]string
	heroes = make(map[string]string)

	villians := make(map[string]string)

	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["The Flash"] = "Barry Allen"
	villians["Lex Luther"] = "Lex Luther"

	superPets := map[int]string{1: "Krypto", 2: "Bat Hound"}

	fmt.Printf("Batman is %v\n", heroes["Batman"])
	pl("Chip :", superPets[3])
	_, ok := superPets[3]
	pl("Is there a 3rd pet :", ok)

	for k, v := range heroes {
		fmt.Printf("%s : %s\n", k, v)
	}
	delete(heroes, "The Flash")
	pl(heroes)
}

// ===================================== Generic And Constraints ==================================== //
type MyConstraint interface {
	int | float64
}

// pkg.go.dev/golang.org/x/exp/constraints
func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}
func genericAndConstraints() {
	pl("5 + 4 =", getSumGen(5, 4))
	pl("5.6 + 4.7 =", getSumGen(5.6, 4.7))
}

// ===================================== Struct ==================================== //
type customer struct {
	name    string
	address string
	bal     float64
}

func getCusInfo(c customer) {
	fmt.Printf("%s owes us %.2f\n", c.name, c.bal)
}
func newCusAdd(c *customer, address string) {
	c.address = address
}
func customerFunc() {
	var tS customer
	tS.name = "Tom Smith"
	tS.address = "5 main st"
	tS.bal = 234.56

	getCusInfo(tS)
	newCusAdd(&tS, "123 South st")
	pl("Address :", tS.address)

	sS := customer{"Sally Smith", "123 Main", 0.0}
	pl("Name :", sS.name)
}

type rectangle struct {
	length, height float64
}

func (r rectangle) Area() float64 {
	return r.length * r.height
}
func structFunc() {
	customerFunc()
	rect1 := rectangle{10.0, 15.0}
	pl("Rect Area :", rect1.Area())
}

// ===================================== Composition ==================================== //
type contact struct {
	fName string
	lName string
	phone string
}
type business struct {
	name    string
	address string
	contact
}

func (b business) info() {
	fmt.Printf("Contact at %s is %s %s\n", b.name, b.contact.fName, b.contact.lName)
}
func compositionFunc() {
	con1 := contact{
		"James",
		"Wang",
		"555-1212",
	}
	bus1 := business{
		"ABC Plumbing",
		"234 North St",
		con1,
	}
	bus1.info()
}

// ===================================== Defined Types ==================================== //
type Tsp float64
type TBs float64
type ML float64

func tspToML(tsp Tsp) ML {
	return ML(tsp * 4.92)
}

func TBToML(tsp TBs) ML {
	return ML(tsp * 14.79)
}

func (tsp Tsp) ToMls() ML {
	return ML(tsp * 4.92)
}

func (tbs TBs) ToMls() ML {
	return ML(tbs * 14.79)
}

func definedTypesFunc() {
	ml1 := ML(Tsp(3) * 4.92)
	fmt.Printf("3 tsps = %.2f ML\n", ml1)
	ml2 := ML(TBs(3) * 14.79)
	fmt.Printf("3 TBs = %.2f ML\n", ml2)
	pl("2 tsp + 4 tsp =", Tsp(2), Tsp(4))
	pl("2 tsp > 4 tsp =", Tsp(2) > Tsp(4))
	fmt.Printf("3 tsp = %.2f mL\n", tspToML(3))
	fmt.Printf("3 tsp = %.2f mL\n", TBToML(3))

	tsp1 := Tsp(3)
	fmt.Printf("%.2f tsp = %.2f mL\n", tsp1, tsp1.ToMls())
}

// ===================================== Interface ==================================== //
type Animal interface {
	AngrySound()
	HappySound()
}

type Cat string

func (c Cat) Attack() {
	pl("cat Attacks its Prey")
}
func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() {
	pl("Cat says Hissssss")
}
func (c Cat) HappySound() {
	pl("Cat says Purrrrrr")
}
func interfaceFunc() {
	var kitty Animal
	kitty = Cat("kitty")
	kitty.AngrySound()
	var kitty2 Cat = kitty.(Cat)
	kitty2.Attack()
	pl("Cats Name :", kitty2.Name())
}

// ===================================== Concurrency / Go Routines ==================================== //
func printTo15() {
	for i := 0; i < 15; i++ {
		pl("Func 1 :", i)
	}
}
func printTo10() {
	for i := 0; i < 10; i++ {
		pl("Func 2 :", i)
	}
}
func concurrencyFunc() {
	go printTo15()
	go printTo10()
	time.Sleep(2 * time.Second)
}

// ===================================== Channels ==================================== //
func nums1(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
}
func nums2(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
}
func channelFunc() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	go nums1(channel1)
	go nums2(channel2)
	pl(<-channel1)
	pl(<-channel1)
	pl(<-channel1)
	pl(<-channel2)
	pl(<-channel2)
	pl(<-channel2)
}

// ===================================== Mutex / Lock ==================================== //
type Account struct {
	balance int
	lock    sync.Mutex
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}
func (a *Account) Withdraw(v int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance {
		pl("Not enough money in account")
	} else {
		fmt.Printf("%d withdrawn : Balance : %d\n", v, a.balance)
		a.balance -= v
	}
}
func mutexLockFunc() {
	var acct Account
	acct.balance = 100
	pl("Balance :", acct.GetBalance())
	for i := 0; i < 12; i++ {
		go acct.Withdraw(10)
	}
	time.Sleep(2 * time.Second)
}

// ===================================== Closures & Passing Function ==================================== //
func useFunc(f func(int, int) int, x, y int) {
	pl("Answer :", (f(x, y)))
}

func sumVals(x, y int) int {
	return x + y
}

func closuresFuc() {
	intSum := func(x, y int) int {
		return x + y
	}
	pl("5 + 4 =", intSum(5, 4))
	samp1 := 1
	changeVar := func() {
		samp1 += 1
	}
	changeVar()
	pl("samp1", samp1)

	useFunc(sumVals, 5, 8)
}

// ===================================== Recursion ==================================== //
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
func recursionFunc() {
	pl("Factorial 4 =", factorial(4))
}

// ===================================== RegularExpression ==================================== //
func regexFunc() {
	reStr := "The ape was at the apex"
	match, _ := regexp.MatchString("(ape[^ ]?", reStr)
	pl(match)

	reStr2 := "Cat rat mat fat pat"
	r, _ := regexp.Compile("[crmfp]at")
	pl("MatchString :", r.MatchString(reStr2))
	pl("FindString :", r.FindString(reStr2))
	pl("FindStringIndex :", r.FindStringIndex(reStr2))
	pl("All String :", r.FindAllString(reStr2, -1))
	pl("1st 2 Strings :", r.FindAllString(reStr2, 2))
	pl("All Submatch Index:", r.FindAllStringSubmatchIndex(reStr2, -1))
	pl(r.ReplaceAllString(reStr2, "Dog"))
}

// ===================================== Automated Testing ==================================== //

func main() {
	// input()
	// variable()
	// dataTypes()
	// casting()
	// ifConditional()
	// stringsFunc()
	// runes()
	// timeFunc()
	// mathFunc()
	// forLoop()
	// whileLoop()
	// rangeFuc()
	// arrayFuc()
	// sliceFunc()
	// function()
	// pointerFunc()
	// arrayPointerFunc()
	// fileIOFunc()
	// mapFunc()
	// genericAndConstraints()
	// structFunc()
	// compositionFunc()
	// definedTypesFunc()
	// interfaceFunc()
	// concurrencyFunc()
	// channelFunc()
	// mutexLockFunc()
	// closuresFuc()
	// recursionFunc()
	regexFunc()
}
