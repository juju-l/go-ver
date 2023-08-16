package version

import (
	"os";"flag";"fmt"
)

			var ver = "v010206r"
			var vFlag = flag.Bool("v", false, "")
			///flag.Parse()

func Get() {
	if *vFlag == true {
	  /*;*/fmt.Println(ver);os.Exit(0)
	}
}

func init() {
	flag.Parse() //
}