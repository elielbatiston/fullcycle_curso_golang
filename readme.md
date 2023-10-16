
mutation createOrder {
	createOrder(input:{id:"ccc",Price:100.0, Tax:2.0}){
    id
    Price
    Tax
    FinalPrice
  }
}


