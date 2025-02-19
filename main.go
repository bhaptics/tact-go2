// main.go
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"bhaptics_go/bhaptics"
)

func printHelp() {
	fmt.Println("\nAvailable commands:")
	fmt.Println("o - Initialize SDK with test keys")
	fmt.Println("quit - Close connection and exit")
	fmt.Println("re - ReInit with test keys")
	fmt.Println("s - Play 'head' event")
	fmt.Println("z - Play custom event")
	fmt.Println("x - Pause event")
	fmt.Println("c - Resume event")
	fmt.Println("r - Play hit-slash with params")
	fmt.Println("e - Play hit-web with params")
	fmt.Println("d - Play 'all' event")
	fmt.Println("f - Stop by request ID")
	fmt.Println("a - Stop all events")
	fmt.Println("i - Check if playing by request ID")
	fmt.Println("l - Check if any event is playing")
	fmt.Println("long - Play long pattern")
	fmt.Println("web - Play hit-web")
	fmt.Println("loop/loops/loopx/loopy/loopxy - Play loop patterns")
	fmt.Println("w - Check if hit-web is playing")
	fmt.Println("q - Stop hit-web event")
	fmt.Println("m - Get device info JSON")
	fmt.Println("n - Get haptic mappings JSON")
	fmt.Println("u - Ping specific device")
	fmt.Println("j - Ping all devices")
	fmt.Println("h - Swap position")
	fmt.Println("vsm - Set device VSM")
	fmt.Println("k - Play dot pattern")
	fmt.Println("event - Get event time")
	fmt.Println("g - Play waveform pattern")
	fmt.Println("pp - Play path pattern")
}

func main() {
	fmt.Println("Hello bHaptics SDK")

	appId := ""
	apiKey := ""
	event := "hit-slash"
	mac := ""

	// Initialize the SDK
	res := bhaptics.RegistryAndInit(apiKey, appId, "")
	if !res {
		fmt.Println("Connection Failed")
		return
	}

	connected := bhaptics.WsIsConnected()
	fmt.Printf("Connected: %v\n", connected)

	scanner := bufio.NewScanner(os.Stdin)
	var requestId int

	printHelp()

	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "o":
			bhaptics.RegistryAndInit(apiKey, appId, "")
		case "quit":
			bhaptics.WsClose()
			return
		case "re":
			result := bhaptics.ReInitMessage(apiKey, appId, "")
			fmt.Println(result)
		case "s":
			requestId = bhaptics.Play("head")
		case "z":
			requestId = bhaptics.Play(event)
		case "x":
			bhaptics.Pause(event)
		case "c":
			bhaptics.Resume(event)
		case ",":
			requestId = bhaptics.PlayWithStartTime(event, rand.Int(), 500, 1.0, 1.0, 0.0, 0.0)
		case "qwer":
			requestId = bhaptics.PlayParam(event, rand.Int(), 1.0, 1.0, 180.0, 0.0)
		case "r":
			requestId = bhaptics.PlayParam("hit-slash", 0, 10.0, 3.0, 0.0, 0.0)
		case "e":
			requestId = bhaptics.PlayParam("hit-web", 0, 1.0, 10.0, 0.0, 0.0)
		case "d":
			requestId = bhaptics.Play("all")
		case "f":
			bhaptics.Stop(requestId)
		case "a":
			bhaptics.StopAll()
		case "i":
			fmt.Println(bhaptics.IsPlayingByRequestId(requestId))
		case "l":
			fmt.Println(bhaptics.IsPlaying())
		case "long":
			requestId = bhaptics.Play("long_pattern")
		case "web":
			requestId = bhaptics.Play("hit-web")
		case "loop":
			requestId = bhaptics.PlayLoop(event, 0, 5.0, 1.0, 0.0, 0.0, 3500, 10)
		case "loops":
			requestId = bhaptics.PlayLoop(event, 0, 5.0, 1.0, 0.0, 0.0, 3500, 10)
		case "loopx":
			requestId = bhaptics.PlayLoop(event, 0, 5.0, 1.0, 90.0, 0.0, 3500, 10)
		case "loopy":
			requestId = bhaptics.PlayLoop(event, 0, 5.0, 1.0, 0.0, 0.25, 3500, 10)
		case "loopxy":
			requestId = bhaptics.PlayLoop(event, 0, 5.0, 1.0, 90.0, 0.25, 3500, 10)
		case "w":
			fmt.Println(bhaptics.IsPlayingByEventId(event))
		case "q":
			fmt.Println(bhaptics.StopByEventId(event))
		case "m":
			fmt.Println(bhaptics.GetDeviceInfoJson())
		case "n":
			fmt.Println(bhaptics.GetHapticMappingsJson())
		case "u":
			fmt.Println(bhaptics.Ping(mac))
		case "j":
			fmt.Println(bhaptics.PingAll())
		case "h":
			fmt.Println(bhaptics.SwapPosition(mac))
		case "vsm":
			fmt.Println(bhaptics.SetDeviceVsm(mac, 40))
		case "k":
			motorValues := make([]int, 32)
			for i := range motorValues {
				motorValues[i] = 40
			}
			fmt.Println(bhaptics.PlayDot(0, 0, 2000, motorValues, len(motorValues)))
		case "event":
			fmt.Println(bhaptics.GetEventTime("scan4"))
		case "g":
			motorValues := make([]int, 8)
			playTime := make([]int, 8)
			shape := make([]int, 8)

			for i := range motorValues {
				motorValues[i] = 100
				playTime[i] = 8
				shape[i] = 2
			}

			repeatCount := 100

			for i := 0; i < 5; i++ {
				fmt.Println(bhaptics.PlayWaveform(0, 8, motorValues, playTime, shape, repeatCount, len(motorValues)))
				time.Sleep(40 * time.Millisecond)
			}

			time.Sleep(500 * time.Millisecond)

			for i := 0; i < 5; i++ {
				fmt.Println(bhaptics.PlayWaveform(0, 9, motorValues, playTime, shape, repeatCount, len(motorValues)))
				time.Sleep(40 * time.Millisecond)
			}
		case "pp":
			xValues := []float32{0.738, 0.723, 0.709, 0.696, 0.682, 0.667, 0.653, 0.639, 0.624, 0.611, 0.597, 0.584, 0.57, 0.557, 0.542, 0.528, 0.515, 0.501, 0.487, 0.474, 0.46, 0.447, 0.432, 0.419, 0.406, 0.393, 0.378, 0.365, 0.352, 0.338, 0.324, 0.311, 0.297}
			yValues := []float32{0.68, 0.715, 0.749, 0.782, 0.816, 0.852, 0.885, 0.921, 0.956, 0.952, 0.918, 0.885, 0.848, 0.816, 0.779, 0.743, 0.71, 0.673, 0.639, 0.606, 0.571, 0.537, 0.5, 0.467, 0.434, 0.4, 0.363, 0.329, 0.296, 0.261, 0.226, 0.192, 0.157}

			intensity := make([]int, len(xValues))
			for i := range intensity {
				intensity[i] = 100
			}

			fmt.Printf("pp: %d\n", len(xValues))
			fmt.Println(bhaptics.PlayPath(0, 0, 1000, xValues, yValues, intensity, len(xValues)))
		case "help":
			printHelp()
		default:
			fmt.Println("Unknown command. Type 'help' for available commands.")
		}

		time.Sleep(100 * time.Millisecond)
	}
}
