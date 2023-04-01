package main

type address struct {
	location string
	ifSchool bool
}

func (a address) schoole() bool {
	return a.ifSchool
}
type room struct {
	int
	name string
	address address
	Price float64 `json:"price"`
}
func (r room) getRoom() *room {
	return &room{
		11,
		name: "111",
		address: address {
			location: "he fei",
			ifSchool: true,
		},
		Price: 1122.99,
	}
}
func main()  {

}
