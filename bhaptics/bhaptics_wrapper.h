// bhaptics_wrapper.h
#ifndef BHAPTICS_WRAPPER_H
#define BHAPTICS_WRAPPER_H

#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

bool registryAndInit(const char* sdkAPIKey, const char* workspaceId, const char* initJsonMessage);
bool wsIsConnected();
void wsClose();
int play(const char* key);
int playParam(const char* key, int requestId, float intensity, float duration, float angleX, float offsetY);
bool stopAll();
bool isbHapticsConnected(int position);
bool isPlaying();
bool isPlayingByRequestId(int requestId);
bool isPlayingByEventId(const char* eventId);
const char* getDeviceInfoJson();
bool isPlayerRunning();
bool isPlayerInstalled();
bool launchPlayer(bool tryLaunch);
bool registryAndInitHost(const char* sdkAPIKey, const char* workspaceId, const char* initJsonMessage, const char* url);
bool reInitMessage(const char* sdkAPIKey, const char* workspaceId, const char* initJsonMessage);
int playWithStartTime(const char* key, int requestId, int startMillis, float intensity, float duration, float angleX, float offsetY);
void playWithoutResult(const char* key, float intensity, float duration, float angleX, float offsetY);
int playDot(int requestId, int position, int duration, int* motorValues, int motorValueLen);
int playPath(int requestId, int position, int durationMillis, float* xValues, float* yValues, int* intensityValues, int Len);
int playWaveform(int requestId, int position, int* motorValues, int* playTimeValues, int* shapeValues, int repeatCount, int motorLen);
int playLoop(const char* eventId, int requestId, float intensity, float duration, float angleX, float offsetY, int interval, int maxCount);
int getEventTime(const char* eventId);
int pause(const char* eventId);
bool resume(const char* eventId);
bool stop(int requestKey);
bool stopByEventId(const char* eventId);
const char* getHapticMappingsJson();
bool ping(const char* address);
bool pingAll();
bool swapPosition(const char* address);
bool setDeviceVsm(const char* address, int vsm);
bool sendAuthentication();

#ifdef __cplusplus
}
#endif

#endif // BHAPTICS_WRAPPER_H