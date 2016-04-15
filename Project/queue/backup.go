package queue

import(
		"encoding/json"
		"os"
		."../network"
		"fmt"
)


func BackupSavetoFile(e *ElevatorState) {
	file, _ := os.Create("ordersInside.json")
	buf, err := json.Marshal(e.OrderInside)
	fmt.Println(err)
	file.WriteAt(buf, 0)
}


func BackupLoadFromFile(e *ElevatorState) {
	file, _ := os.Open("ordersInside.json")
	buf := make([]byte, 1024)
	n, _ := file.ReadAt(buf, 0)
	fmt.Println(string(buf))
	err := json.Unmarshal(buf[:n], &e.OrderInside)
	fmt.Println(err)
}
