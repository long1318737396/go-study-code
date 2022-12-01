package main
import "fmt"
//银行初始储蓄10000
var totalBalance float64=10000
//客户结构体
type Customer struct {
	Name string
	Balance float64
}
//存钱
func (customerInstance *Customer) deposit(money float64){
	customerInstance.Balance+=money
	totalBalance+=money
	fmt.Println(customerInstance.Name,"当前余额：",customerInstance.Balance)
	fmt.Println("银行总储蓄：",totalBalance)
}
//取钱
func (customerInstance *Customer) withdraw(money float64){
	if customerInstance.Balance<money{
		fmt.Println("账户余额不足，无法支取")
	}else if totalBalance<money{
		fmt.Println("银行储蓄不足，无法支取")
	}else{
		customerInstance.Balance-=money
		totalBalance-=money
		fmt.Println(customerInstance.Name,"当前余额：",customerInstance.Balance)
		fmt.Println("银行总储蓄：",totalBalance)
	}
}
func main(){
	//声明customerA，即储户A
	var customerA Customer
	customerA.Name="储户A"
	customerA.Balance=2000
	//声明customerB，即储户B
	var customerB Customer
	customerB.Name="储户B"
	customerB.Balance=3000
	//声明customerC，即储户C
	var customerC Customer
	customerC.Name="储户C"
	customerC.Balance=5000
	//储户A存入5000
	customerA.deposit(5000)
	//储户C支取3000
	customerC.withdraw(3000)
	//储户B支取4000
	customerB.withdraw(4000)
	//储户A支取6000
	customerA.deposit(6000)
	//储户B存入2000
	customerB.deposit(2000)
	//储户C存入10000
	customerC.deposit(10000)
	//储户B支取500
	customerB.withdraw(500)
}