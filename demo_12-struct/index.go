package main

import "fmt"

// --------------- Struct product ---------------
// Struct is Class Object in Java
type product struct {
	name  string
	price int
	stock int
}

// สร้าง Function ให้กับ Class product
func (p product) clear() product {
	p.price = 0
	p.stock = 0
	return p
}

// สร้าง Function ให้กับ Class product แบบส่งมาเป็น pointer (ทำให้ไม่ต้อง return ค่ากลับ)
func (p *product) addPrice(plusValue int) {
	p.price = p.price + plusValue
}

// ----------------------------------------------

func main() {
	var p1 product
	p1.name = "Arduino"
	p1.price = 100
	p1.stock = 20
	showProduct(p1)

	updateProduct(p1)
	showProduct(p1) // ค่าไม่เปลี่ยนแปลง เนื่องจากไม่ได้ส่งไปเป็น Pointer

	updateProductByPointer(&p1)
	showProduct(p1)

	// fmt.Printf("product address : %p\n", &p1)
	p1 = p1.clear() // assign ค่าที่ได้ ใส่ใน p1
	showProduct(p1)

	p1.addPrice(20)
	showProduct(p1)
	p1.addPrice(30)
	showProduct(p1)

}

func showProduct(p product) {
	fmt.Println("Product name :", p.name)
	fmt.Println("Product price :", p.price)
	fmt.Println("Product stock :", p.stock)
}

func updateProduct(p product) {
	p.name = "NodeMCU"
	p.price = 200
	p.stock = 10

	// ค่าจะเปลี่ยนแปลง แค่เฉพาะใน Function นี้เท่านั้น ออกนอก Function นี้ไปแล้ว จะเป็นค่าเดิม เนื่องจากเป็นคนละ Object กัน (Pointer ชี้ไปคนละ Address)
	fmt.Println("Value in function updateProduct :", p)
}

func updateProductByPointer(p *product) {
	p.name = "NodeMCU"
	p.price = 200
	p.stock = 10
}
