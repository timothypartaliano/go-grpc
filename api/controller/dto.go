package controller

type ReqBodyCreateProduct struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Stock       int     `json:"stock"`
}

// message Product{
//     string id = 1;
//     string name = 2;
//     string description = 3;
//     float price = 4;
//     int32 stock = 5;
// }