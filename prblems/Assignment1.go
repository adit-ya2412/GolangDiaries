package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type resutrant struct {
	price    int
	dishname string
}

var menu = map[string]resutrant{
	"C": resutrant{price: 40, dishname: "Cofee"},
	"D": resutrant{price: 80, dishname: "dosa"},
	"T": resutrant{price: 20, dishname: "tomato soup"},
	"I": resutrant{price: 20, dishname: "idli"},
	"V": resutrant{price: 25, dishname: "vada"},
	"B": resutrant{price: 30, dishname: "bhature&chhole"},
	"P": resutrant{price: 30, dishname: "paneer pakoda"},
	"M": resutrant{price: 90, dishname: "manchurian"},
	"H": resutrant{price: 70, dishname: "hakka noodle"},
	"F": resutrant{price: 30, dishname: "french fries"},
	"J": resutrant{price: 30, dishname: "jalebi"},
	"L": resutrant{price: 15, dishname: "lemonade"},
	"S": resutrant{price: 20, dishname: "spring roll"},
}

func main() {
	var input string

	var total int
	fmt.Println("Plz place the order :")
	reader := bufio.NewReader(os.Stdin)
	input, er := reader.ReadString('\n')
	if er != nil {
		fmt.Println("An error occured while reading input. Please try again", er)
		return
	}
	input = strings.TrimSpace(input)
	runeSample := []rune(input)
	input1 := (string(runeSample[0]))
	input2 := (string(runeSample[len(runeSample)-1]))
	for input != "END" {

		items, e := strconv.Atoi(input1)
		data, ifexist := menu[input2]

		if ifexist {
			if e == nil {
				total = total + items*data.price
				fmt.Println("***************************")
				fmt.Printf("Bill(%d*price of %s)=", items, data.dishname)
				fmt.Println(data.price * items)
				fmt.Println("***************************")
			} else {
				fmt.Println("there is error", e)

			}
		} else {
			fmt.Println("Wrong input")
		}
		fmt.Println("Plz place the order :")
		input, er := reader.ReadString('\n')
		if er != nil {
			fmt.Println("An error occured while reading input. Please try again", er)
			return
		}
		input = strings.TrimSpace(input)
		if input == "END" {
			break
		}
		runeSample := []rune(input)
		input1 = (string(runeSample[0]))
		input2 = (string(runeSample[len(runeSample)-1]))

	}
	fmt.Println("The total income for the day is ", total)

}
