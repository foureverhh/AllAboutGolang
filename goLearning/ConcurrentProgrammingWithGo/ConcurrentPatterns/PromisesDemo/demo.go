package main

import(
	"fmt"
	"errors"
	"time"
)
func main() {
	po := new(PurchaseOrder)
	po.Value = 42.37
	SavePO(po,true).Then(func(obj interface{}) error {
		po := obj.(*PurchaseOrder)
		fmt.Printf("Purchase Order saved with ID: %d\n", po.Number)
		return nil
	},func(err error){
		fmt.Print("Failed to save Purchase Order: " + err.Error() + "\n")
	}).Then(func (obj interface{}) error {
		fmt.Println("Second promise success")
		return nil
	},func (err error)  {
		fmt.Println("Second promise failed: "+err.Error())
	})
	fmt.Printf("PO is: %v and type is %T",po,po)
	//fmt.Scanln()
}
//PurchaseOrder structure
type PurchaseOrder struct{
	Number int
	Value float64
}
//SavePO with promise
func SavePO(po *PurchaseOrder,isFail bool) *Promise{
	result := new (Promise)
	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error,1)
	go func ()  {
		time.Sleep(2 * time.Second)
		if isFail {
			result.failureChannel <- errors.New("Failed to save purchase")
		}else {
			po.Number = 1234
			result.successChannel <- po
		}	
	}()
	return result
}

//Promise type
type Promise struct{
	successChannel chan interface{}
	failureChannel chan error
}

//Then structure
func (promise *Promise) Then(success func(interface{}) error, failure func (error)) *Promise {
	result := new(Promise)
	result.successChannel = make(chan interface{},1)
	result.failureChannel = make(chan error,1)
	timeout := time.After(1 * time.Second)
	go func ()  {
		select {
		case obj:=<-promise.successChannel:
			newErr := success(obj)
			if newErr==nil {
				result.successChannel<-obj
			}else {
				result.failureChannel<-newErr
			}
		case err:=<-promise.failureChannel:
				failure(err)
				result.failureChannel <- err
		case <- timeout:
				err := errors.New("Promise timed out")
				failure(err)
				result.failureChannel <- err
		}
	}()

	return result
}