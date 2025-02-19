// bhaptics.go
package bhaptics

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo LDFLAGS: -L${SRCDIR} -lbhaptics_library

#include "bhaptics_wrapper.h"
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

func RegistryAndInit(sdkAPIKey, workspaceId, initJsonMessage string) bool {
	cSDKAPIKey := C.CString(sdkAPIKey)
	cWorkspaceId := C.CString(workspaceId)
	cInitJsonMessage := C.CString(initJsonMessage)
	defer func() {
		C.free(unsafe.Pointer(cSDKAPIKey))
		C.free(unsafe.Pointer(cWorkspaceId))
		C.free(unsafe.Pointer(cInitJsonMessage))
	}()

	return bool(C.registryAndInit(cSDKAPIKey, cWorkspaceId, cInitJsonMessage))
}

func WsIsConnected() bool {
	return bool(C.wsIsConnected())
}

func WsClose() {
	C.wsClose()
}

func Play(key string) int {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	return int(C.play(cKey))
}

func PlayParam(key string, requestId int, intensity, duration, angleX, offsetY float32) int {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	return int(C.playParam(cKey, C.int(requestId), C.float(intensity), C.float(duration), C.float(angleX), C.float(offsetY)))
}

func StopAll() bool {
	return bool(C.stopAll())
}

func IsBHapticsConnected(position int) bool {
	return bool(C.isbHapticsConnected(C.int(position)))
}

func IsPlaying() bool {
	return bool(C.isPlaying())
}

func IsPlayingByRequestId(requestId int) bool {
	return bool(C.isPlayingByRequestId(C.int(requestId)))
}

func IsPlayingByEventId(eventId string) bool {
	cEventId := C.CString(eventId)
	defer C.free(unsafe.Pointer(cEventId))
	return bool(C.isPlayingByEventId(cEventId))
}

func GetDeviceInfoJson() string {
	return C.GoString(C.getDeviceInfoJson())
}

func IsPlayerRunning() bool {
	return bool(C.isPlayerRunning())
}

func IsPlayerInstalled() bool {
	return bool(C.isPlayerInstalled())
}

func LaunchPlayer(tryLaunch bool) bool {
	return bool(C.launchPlayer(C.bool(tryLaunch)))
}

func RegistryAndInitHost(sdkAPIKey, workspaceId, initJsonMessage, url string) bool {
	cSDKAPIKey := C.CString(sdkAPIKey)
	cWorkspaceId := C.CString(workspaceId)
	cInitJsonMessage := C.CString(initJsonMessage)
	cUrl := C.CString(url)
	defer func() {
		C.free(unsafe.Pointer(cSDKAPIKey))
		C.free(unsafe.Pointer(cWorkspaceId))
		C.free(unsafe.Pointer(cInitJsonMessage))
		C.free(unsafe.Pointer(cUrl))
	}()

	return bool(C.registryAndInitHost(cSDKAPIKey, cWorkspaceId, cInitJsonMessage, cUrl))
}

func ReInitMessage(sdkAPIKey, workspaceId, initJsonMessage string) bool {
	cSDKAPIKey := C.CString(sdkAPIKey)
	cWorkspaceId := C.CString(workspaceId)
	cInitJsonMessage := C.CString(initJsonMessage)
	defer func() {
		C.free(unsafe.Pointer(cSDKAPIKey))
		C.free(unsafe.Pointer(cWorkspaceId))
		C.free(unsafe.Pointer(cInitJsonMessage))
	}()

	return bool(C.reInitMessage(cSDKAPIKey, cWorkspaceId, cInitJsonMessage))
}

func PlayWithStartTime(key string, requestId, startMillis int, intensity, duration, angleX, offsetY float32) int {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	return int(C.playWithStartTime(cKey, C.int(requestId), C.int(startMillis), C.float(intensity), C.float(duration), C.float(angleX), C.float(offsetY)))
}

func PlayWithoutResult(key string, intensity, duration, angleX, offsetY float32) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	C.playWithoutResult(cKey, C.float(intensity), C.float(duration), C.float(angleX), C.float(offsetY))
}

func PlayDot(requestId, position, duration int, motorValues []int, motorValueLen int) int {
	cMotorValues := make([]C.int, len(motorValues))
	for i, v := range motorValues {
		cMotorValues[i] = C.int(v)
	}
	return int(C.playDot(C.int(requestId), C.int(position), C.int(duration), &cMotorValues[0], C.int(motorValueLen)))
}

func PlayPath(requestId, position, durationMillis int, xValues, yValues []float32, intensityValues []int, length int) int {
	cXValues := make([]C.float, len(xValues))
	cYValues := make([]C.float, len(yValues))
	cIntensityValues := make([]C.int, len(intensityValues))

	for i, v := range xValues {
		cXValues[i] = C.float(v)
	}
	for i, v := range yValues {
		cYValues[i] = C.float(v)
	}
	for i, v := range intensityValues {
		cIntensityValues[i] = C.int(v)
	}

	return int(C.playPath(C.int(requestId), C.int(position), C.int(durationMillis), &cXValues[0], &cYValues[0], &cIntensityValues[0], C.int(length)))
}

func PlayWaveform(requestId, position int, motorValues, playTimeValues, shapeValues []int, repeatCount, motorLen int) int {
	cMotorValues := make([]C.int, len(motorValues))
	cPlayTimeValues := make([]C.int, len(playTimeValues))
	cShapeValues := make([]C.int, len(shapeValues))

	for i, v := range motorValues {
		cMotorValues[i] = C.int(v)
	}
	for i, v := range playTimeValues {
		cPlayTimeValues[i] = C.int(v)
	}
	for i, v := range shapeValues {
		cShapeValues[i] = C.int(v)
	}

	return int(C.playWaveform(C.int(requestId), C.int(position), &cMotorValues[0], &cPlayTimeValues[0], &cShapeValues[0], C.int(repeatCount), C.int(motorLen)))
}

func PlayLoop(eventId string, requestId int, intensity, duration, angleX, offsetY float32, interval, maxCount int) int {
	cEventId := C.CString(eventId)
	defer C.free(unsafe.Pointer(cEventId))
	return int(C.playLoop(cEventId, C.int(requestId), C.float(intensity), C.float(duration), C.float(angleX), C.float(offsetY), C.int(interval), C.int(maxCount)))
}

func GetEventTime(eventId string) int {
	cEventId := C.CString(eventId)
	defer C.free(unsafe.Pointer(cEventId))
	return int(C.getEventTime(cEventId))
}

func Pause(eventId string) int {
	cEventId := C.CString(eventId)
	defer C.free(unsafe.Pointer(cEventId))
	return int(C.pause(cEventId))
}

func Resume(eventId string) bool {
	cEventId := C.CString(eventId)
	defer C.free(unsafe.Pointer(cEventId))
	return bool(C.resume(cEventId))
}

func Stop(requestKey int) bool {
	return bool(C.stop(C.int(requestKey)))
}

func StopByEventId(eventId string) bool {
	cEventId := C.CString(eventId)
	defer C.free(unsafe.Pointer(cEventId))
	return bool(C.stopByEventId(cEventId))
}

func GetHapticMappingsJson() string {
	return C.GoString(C.getHapticMappingsJson())
}

func Ping(address string) bool {
	cAddress := C.CString(address)
	defer C.free(unsafe.Pointer(cAddress))
	return bool(C.ping(cAddress))
}

func PingAll() bool {
	return bool(C.pingAll())
}

func SwapPosition(address string) bool {
	cAddress := C.CString(address)
	defer C.free(unsafe.Pointer(cAddress))
	return bool(C.swapPosition(cAddress))
}

func SetDeviceVsm(address string, vsm int) bool {
	cAddress := C.CString(address)
	defer C.free(unsafe.Pointer(cAddress))
	return bool(C.setDeviceVsm(cAddress, C.int(vsm)))
}
