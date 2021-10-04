package main

import (
	database "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-transform/database"
	myLog "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-transform/myLog"
	subscriber "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-transform/subscriber"
)

func main() {
	myLog.MyLog(myLog.Get_level_INFO(), "main(start)")
	client := subscriber.Connect("tcp://localhost:1883", "my-client-id")
	client.Publish("a/b/#", 0, false, "my great message")
	database.Insert()
	myLog.MyLog(myLog.Get_level_INFO(), "main(end)")
}
