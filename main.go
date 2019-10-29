package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var attackerNumShip int
	var defenderNumShip int

	defenderWinCount := 0
	attackerWinCount := 0

	for i:= 0; i < 1000; i++ {
		attackerNumShip = 3
		defenderNumShip = 3
		combatSimulation(&attackerNumShip, &defenderNumShip)
		//fmt.Printf("Attacker Ships Left: %d, Defender Ships Left: %d\n", attackerNumShip, defenderNumShip)

		if attackerNumShip < defenderNumShip {
			defenderWinCount++
		} else {
			attackerWinCount++
		}
	}

	fmt.Printf("Attacker wins: %d, Defender wins: %d\n", attackerWinCount, defenderWinCount)
}

// retreat returns if retreat is successful
func retreat() bool {
	max := 6
	min := 1

	thresholdForSuccess := 5
	return rand.Intn(max - min) + min >= thresholdForSuccess
}

// attack returns if attacker successfully hits defender
func attack() (defender int, success bool) {
	attackerMax := 6
	attackerMin := 1
	attacker := rand.Intn(attackerMax - attackerMin) + attackerMin

	defenderMax := 6
	defenderMin := 1
	defender = rand.Intn(defenderMax - defenderMin) + defenderMin

	success = attacker >= defender

	return defender, success
}

// maneuver returns if attacker successfully outmaneuvers the defender
func maneuver(defenderPrevRole int) bool {
	maneuverMax := 6
	maneuverMin := 1
	maneuverRole := rand.Intn(maneuverMax - maneuverMin) + maneuverMin

	if maneuverRole > defenderPrevRole {
		return true
	}

	return false
}

// combatSimulation will simulate if a full combat with
func combatSimulation(attackerNumShip, defenderNumShip *int) {
	if *defenderNumShip == 0 || *attackerNumShip == 0 {
		return
	}

	// Only try to retreat if defender ships <= attacker ships
	if *defenderNumShip <= *attackerNumShip {
		success := retreat()
		if success {
			return
		}
	}

Attack:
	defenderRole, success := attack()
	if success {
		*defenderNumShip -= 1
	}

	if *defenderNumShip == 0 {
		return
	}

	success = maneuver(defenderRole)
	if success {
		goto Attack
	} else {
		// Defender Counter attacks
		combatSimulation(defenderNumShip, attackerNumShip)
		return
	}
}