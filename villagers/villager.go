package villagers

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
)

type Villagers struct {
	Villagers []Villager `json:"villagers"`
}

type Villager struct {
	name  string `json:"name"`
	photo string `json:"photo"`
	alive bool   `json:"alive"`
}

func openVillagersJSON(villagers Villagers) {
	villagersFile, err := os.Open("villagers.json")

	if err != nil {
		println("Could not open or find Json File")
	}

	byteValue, _ := ioutil.ReadAll((villagersFile))
	json.Unmarshal(byteValue, &villagers)
}
func randomVillagerNumber(villagers Villagers) int {
	for {
		randomVillagerNumber := rand.Int() % len(villagers.Villagers)
		if villagers.Villagers[randomVillagerNumber].alive {
			return randomVillagerNumber
		}
	}
}
func GetRandomVillager() Villager {
	var villagers Villagers
	var selectedVillager Villager

	openVillagersJSON(villagers)
	villagerNumber := randomVillagerNumber(villagers)

	selectedVillager.name = villagers.Villagers[villagerNumber].name
	selectedVillager.photo = villagers.Villagers[villagerNumber].photo
	return selectedVillager
}
func makeVillagersFight(villager Villager, villager2 Villager) {

	firstContestant := GetRandomVillager()
	secondContestant := GetRandomVillager()

	println(firstContestant.name + " getWeapon + getPhrase" + secondContestant.name)

	winnerNumber := rand.Intn(2)

	switch winnerNumber {
	case 1:
		villager.Win()
		villager2.Lose()
		break
	case 2:
		villager.Lose()
		villager2.Win()
	default:
		villagerKillsItself()
	}

	// sendTweet winner.name weapon.random loser.name and makes loser.name leave the island :(
	//loser.alive = false
}
func (v *Villager) Win() {
	// sendTweet Villager.name has won the Island clash!
}
func (v *Villager) Lose() {
	//write json, look for villager,then change alive property
	v.alive = false
}
func villagerKillsItself() {
	//loser.suicide
	//loser.alive = false
}
