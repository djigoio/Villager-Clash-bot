package villagers

type Villager struct {
	name  string
	photo string
}

// maxwin = each time poll ends, if villager win == maxwin maxwin += 1
func GetRandomVillager() Villager {
	var villager Villager
	//get random villager from json
	//where lost == 0 && win < maxWin
	villager.name = ""
	villager.photo = ""
	return villager
}

func VillagerKillsVillager() {
	//winner := GetRandomVillager()
	//loser := GetRandomVillager()
	// sendTweet winner.name weapon.random loser.name and makes loser.name leave the island :(
	//loser.alive = false
}

func VillagerKillsItself() {
	//loser := GetRandomVillager()
	//loser.suicide
	// sendTweet loser suicidePhrase and leaves the island in an ambulance
	//loser.alive = false
}

func MakeVillagerWin() {
	// sendTweet Villager.name has won the Island clash!
}
