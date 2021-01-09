package main

import "fmt"

// 간단한 일과들..
func makeHotelReservation() {
	fmt.Println("호텔을 예약 한다.")
}
func bookFlightTickets() {
	fmt.Println("비행기를 예약 한다.")
}
func orderADress() {
	fmt.Println("드레스를 주문한다.")
}
func payCreditCardBills() {
	fmt.Println("신용카드 청구서를 지불 한다.")
}

// 아래 일과들은 여러개로 나누어진 일과 이다.

// 메일을 쓴다.
func writeAMail() {
	fmt.Println("메일의 1/3을 작성한다.")
	continueWritingMail1()
}
func continueWritingMail1() {
	fmt.Println("메일의 2/3까지 작성한다.")
	continueWritingMail2()
}
func continueWritingMail2() {
	fmt.Println("메일 작성을 완료 한다.")
}

// 오디오 북을 듣는다.
func listenToAudioBook() {
	fmt.Println("오디오북을 10분간 청취 한다.")
	continueListeningToAudioBook()
}
func continueListeningToAudioBook() {
	fmt.Println("오디오북 청취를 모두 완료 한다.")
}

// 하루 동안 처리해야 될 일과이다.
// 메일 쓰는것과 오디오 북을 듣는 것은 여러 작업으로 나누어져 있다.
var listOfTasks = []func(){
	makeHotelReservation, bookFlightTickets, orderADress,
	payCreditCardBills, writeAMail, listenToAudioBook,
}

func main() {
	for _, task := range listOfTasks {
		task()
	}
}
