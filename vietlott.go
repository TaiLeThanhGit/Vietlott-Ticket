package main

import (
    "fmt"
    "io/ioutil"
	"strconv"
	"math/rand"
	"time"
	"sort"
	"flag"
	"strings"
	"os"
)
const (
	MIN_MEGA 		= 1
	MAX_MEGA 		= 45
	
	MIN_POWER 		= 1
	MAX_POWER 		= 55
	
	MEGA_TYPE 		= 1
	POWER_TYPE 		= 2
	
	NUMBER_TICKETS 	= 1
	
	QUANTITY_NUMBERS	= 6
	
	MEGA_NAME		= "MEGA"
	POWER_NAME		= "POWER"
)

type ticket struct {
	nums 		[]int
}

type vietLott struct {
	vietlottType			int
	min 					int
	max 					int
	numberTickets 			int
	tickets 				[]ticket
	unexpectedNumbersFile	string
	unexpectedNumbers 		[]int
}

func main() {
	vl := NewVietLott()
	
	vl.Init()
	
	vl.getTickets()
	vl.PrintTickets()
}

func Usage() {
	defer os.Exit(0)
	
	fmt.Println(os.Args[0] , " [-t type] [-n quantity] [-f file] [-h]")
	fmt.Println("\t type: 1 is Mega, 2 is Power")
	fmt.Println("\t quantity: > 0")
	fmt.Println("\t file: file name contains the list of unexpected numbers")
	fmt.Println("\t -h: show this help")
	
}

func (this *vietLott) GetFlags() {

	var (
		strVietLottType, strNumberTickets string
		help bool
		unexpectedNumbersFile string
	)
	
	flag.StringVar(&strVietLottType, "t", strconv.Itoa(MEGA_TYPE), "VietLott type")
	flag.StringVar(&strNumberTickets, "n", strconv.Itoa(NUMBER_TICKETS), "Number of tickets")
	flag.StringVar(&unexpectedNumbersFile, "f", "", "file name")
	flag.BoolVar(&help, "h", false, "Usage")
	flag.Parse()
	
	if (help) {
		Usage()
	}
	vietLottType, err := strconv.Atoi(strVietLottType);
	   
	if  err != nil {
		fmt.Println("VietLott Type param is invalid: -t ", strVietLottType)
		Usage()
		os.Exit(0)
	}
	
	numberTickets, err := strconv.Atoi(strNumberTickets);
	   
	if  err != nil {
		fmt.Println("Number Tickets is invalid: -n ", strNumberTickets)
		Usage()
		os.Exit(0)
	}

	if vietLottType != MEGA_TYPE && vietLottType != POWER_TYPE {
		vietLottType = MEGA_TYPE
	}
	if numberTickets < 0 {
		numberTickets = NUMBER_TICKETS
	}
	
	this.SetType(vietLottType)
	this.SetNumberTickets(numberTickets)
	this.unexpectedNumbersFile = unexpectedNumbersFile
}

func (this *vietLott) Init() {
	// get params from commandline
	this.GetFlags()
	// get unexpected numbers from file
	this.GetUnpextedNumbers()
}

func (this *vietLott) GetUnpextedNumbers() {
	if this.unexpectedNumbersFile == "" {
		return
	}
	content, err := ioutil.ReadFile(this.unexpectedNumbersFile)
	if err != nil {
		fmt.Println("cannot read file " + this.unexpectedNumbersFile)
		panic(err)
	}
	
	//replace all new line chars with one space
	strContent := strings.Replace(string(content), "\r\n", " ", -1)
	
	//convert string to string slice
	strNumbers := strings.Split(strContent, " ")
	
	for _, n := range strNumbers {
		//convert string to number
		temp, err := strconv.Atoi(n)
		if (err != nil) {
			panic(err)
		}
		// check duplicate number
		if !inArray(this.unexpectedNumbers, temp) {
			this.unexpectedNumbers = append(this.unexpectedNumbers, temp)
		}
	}
	
}

func (this *vietLott) SetType(vietLottType int) {
	this.vietlottType = vietLottType
	
	if vietLottType == MEGA_TYPE {
		this.min = MIN_MEGA
		this.max = MAX_MEGA
	} else if vietLottType == POWER_TYPE {
		this.min = MIN_POWER
		this.max = MAX_POWER	
	}
}

func (this *vietLott) SetNumberTickets(n int) {
	this.numberTickets = n
}

func (this *vietLott) GetVietlottName() string {
	name := ""
	if (this.vietlottType == MEGA_TYPE) {
		name = MEGA_NAME
	}
	if (this.vietlottType == POWER_TYPE) {
		name = POWER_NAME
	}
	return name
}
func (this vietLott) PrintTickets() {
	
	fmt.Printf("Unexpected Numbers: %v\n", this.unexpectedNumbers)
	fmt.Println("Type: ", this.GetVietlottName())
	fmt.Println("Number of Tickets: ", this.numberTickets)
	
	for i, ticket := range this.tickets {
		fmt.Print(i+1, ": ")
		for _, v := range ticket.nums {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}

func (this *vietLott) getTickets() {
	count := 0
	for count < this.numberTickets{
		this.tickets = append(this.tickets, this.GenerateTicket())
		count++
	}
	
}

func (this vietLott) GenerateTicket() ticket {
	
	var randomTicket []int
	for {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(this.max - this.min) + this.min
		
		if !inArray(randomTicket, randNum) {
			randomTicket = append(randomTicket, randNum)
		}

		if len(randomTicket) == QUANTITY_NUMBERS {
			break
		}
	}
	//fmt.Println(randomTicket)
	sort.Ints(randomTicket)
	return ticket{nums: randomTicket}
}

func inArray(a []int, element int) bool {
	for _, v := range a {
		if (v == element) {
			return true
		}
	}
	return false
}

func NewVietLott() vietLott {
	return vietLott{
		vietlottType: MEGA_TYPE, 
		min: MIN_MEGA, 
		max: MAX_MEGA, 
		numberTickets: NUMBER_TICKETS, 
		unexpectedNumbersFile: "",
	}
}
