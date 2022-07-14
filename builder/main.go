package main

func main() {
	homeInstance := NewHome()
	builderInstance := NewHomeBuilder()
	builderInstance.Wc(1).Room(2)
	homeInstance.build(builderInstance)

}
