package queue

import(
	. "../defines"		
	"fmt"
	"math"
	. "../network"
)


func costLocal(e *ElevatorState, floor int, button int) int {
	

	const distCost = 4
	const numOrderCost = 6


	distanceCost := int(math.Abs(float64((e.Floor - floor)))) * distCost

	numOrders := 0
	for i := 0; i < N_FLOORS; i++ {
		if e.OrderUp[i] == 1{
			numOrders++
		}
		if e.OrderDown[i] == 1{
		 	numOrders++
		}
		if e.OrderInside[i] == 1{
		 	numOrders++
		}
	}



	
	return  numOrders*numOrderCost + distanceCost
}

func CostGlobal(states map[string]*ElevatorState, peerList []string, localIP string, floor int, button int) string {
	//costs := map[string]int

	fmt.Println("Cost fn enter")
	fmt.Println("states: ", states)
	fmt.Println("peerList: ", peerList)
	fmt.Println("localIP: ", localIP)
	
	minstCost := 1000
	minstCostIP := localIP
	for IP, state := range states {
		for i := 0; i < len(peerList); i++ {

			if IP == peerList[i] {
				cost := costLocal(state, floor, button)
				fmt.Println("Cost for ", IP, ": ", cost)
			
				if cost < minstCost {
					minstCost = cost
					minstCostIP = IP 
				}
			}
		}
	}
	return minstCostIP
}
