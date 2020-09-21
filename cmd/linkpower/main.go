package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ghanto/linkpower/pkg/dataloader"
	"github.com/ghanto/linkpower/pkg/linkpower"
)

var fileFlag = flag.String("f", "./resources/test_data.json", "input json path")

func main() {
	flag.Parse()
	f := *fileFlag

	dl := dataloader.NewDataLoader(f)
	testCases, err := dl.LoadTestCases()
	if err != nil {
		log.Fatal(err)
	}

	for _, tc := range *testCases {
		world := linkpower.NewWorld()
		world.AddDevice(linkpower.NewDevice(tc.Device.X, tc.Device.Y))
		for _, s := range tc.Stations {
			err := world.AddStation(linkpower.NewStation(s.X, s.Y, s.R))
			if err != nil {
				log.Fatalf(err.Error())
			}
		}

		bestLink, power, err := world.FindBestLink()
		if err != nil {
			fmt.Printf("No link station within reach for point (%d,%d) \n", tc.Device.X, tc.Device.Y)
			continue
		}

		fmt.Printf(
			"Best link station for point (%d,%d) is (%d,%d) with power %g \n",
			tc.Device.X, tc.Device.Y,
			bestLink.PosX, bestLink.PosY, power,
		)
	}
}
