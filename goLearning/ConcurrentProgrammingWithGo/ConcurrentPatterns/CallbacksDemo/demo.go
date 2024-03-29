package main

import "fmt"

func main() {
	po := new (PurchaseOrder)
	po.Value = 42.27

	ch := make(chan *PurchaseOrder)
	go SavePO(po,ch)
	newPo := <- ch
	fmt.Printf("PO: %v %T",newPo,newPo)
}

//PurchaseOrder struct
type PurchaseOrder struct{
	Number int
	Value float64
}
//SavePO to save PurchaseOrder
func SavePO(po *PurchaseOrder, callback chan *PurchaseOrder)  {
	po.Number = 1234
	callback <- po
}