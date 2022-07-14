package main

type HomeBuilder struct {
	cKitchen  int
	cRoom     int
	cWc       int
	cFloor    int
	cGameRoom int
}

func NewHomeBuilder() HomeBuilder {
	return HomeBuilder{}
}

func (homeInstance *HomeBuilder) kitchen(count int) *HomeBuilder {
	homeInstance.cKitchen = count
	return homeInstance

}
func (homeInstance *HomeBuilder) Room(count int) *HomeBuilder {
	homeInstance.cRoom = count
	return homeInstance

}

func (homeInstance *HomeBuilder) Wc(count int) *HomeBuilder {
	homeInstance.cWc = count
	return homeInstance

}

func (homeInstance *HomeBuilder) Floor(count int) *HomeBuilder {
	homeInstance.cFloor = count
	return homeInstance
}

func (homeInstance *HomeBuilder) GameRoom(count int) *HomeBuilder {
	homeInstance.cGameRoom = count
	return homeInstance

}
