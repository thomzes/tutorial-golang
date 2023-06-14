package main

import "fmt";

func main() {
	type NoKTP string;
	type Married bool;

	var noKtpThomas NoKTP = "23902930239012381";
	fmt.Println(noKtpThomas);

	var marriedStatus Married = true;
	fmt.Println(marriedStatus);

}