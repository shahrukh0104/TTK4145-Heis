package defines

const(
	//QUANTITY
	N_FLOORS			int = 4
	N_ELEVATORS			int = 1
	N_BUTTONS			int = 3


	//DIRECTIONS
	DIR_UP				int = 300
	DIR_DOWN			int = -300
	DIR_STOP			int = 0
	
	//LIGHTS
	LIGHT_ON			int = 1
	LIGHT_OFF			int = 0

	//LAMP CALL
	BUTTON_CALL_UP		int = 0
	BUTTON_CALL_DOWN	int = 1
	BUTTON_COMMAND		int = 2

	//STATES
	INIT				int = 0 
	IDLE				int = 1
	MOVING				int = 2
	DOORSOPEN			int = 3
	STOP 				int = 4
)

type MSG struct{
	State			int
	PrevFloor		int
	Dir				int
	OrderUp			[N_FLOORS] int
	OrderDown		[N_FLOORS] int
	OrderInside		[N_FLOORS] int
}

var Msg = MSG{}
