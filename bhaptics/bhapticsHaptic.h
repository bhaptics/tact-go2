//
// Created by MJ on 2022-05-23.
//

#ifndef HAPTIC_LIBRARY_BHAPTICSHAPTIC_H
#define HAPTIC_LIBRARY_BHAPTICSHAPTIC_H

const std::string SdkVersion = "2.0.0";

const std::string SdkPlay = "SdkPlay";
const std::string SdkStopByRequestId = "SdkStopByRequestId";
const std::string SdkStopByEventId = "SdkStopByEventId";
const std::string SdkStopAll = "SdkStopAll";
const std::string SdkRequestAuth = "SdkRequestAuth";
const std::string SdkRequestAuthInit = "SdkRequestAuthInit";
const std::string SdkRequestReInit = "SdkRequestReInit";
const std::string ServerDevices = "ServerDevices";
const std::string ServerReady = "ServerReady";
const std::string ServerEventNameList = "ServerEventNameList";
const std::string ServerEventList = "ServerEventList";
const std::string ServerActiveEventNameList = "ServerActiveEventNameList";
const std::string ServerActiveRequestIdList = "ServerActiveRequestIdList";
const std::string SdkPing = "SdkPing";
const std::string SdkPingAll = "SdkPingAll";
const std::string SdkSwapPosition = "SdkSwapPosition";
const std::string SdkPlayDotMode = "SdkPlayDotMode";
const std::string SdkPlayPathMode = "SdkPlayPathMode";
const std::string SdkPlayLoop = "SdkPlayLoop";
const std::string SdkPlayWaveformMode = "SdkPlayWaveformMode";
const std::string SdkPlayWaveformMode8 = "SdkPlayWaveformMode8";
const std::string SdkDeviceSetVSM = "SdkDeviceSetVSM";
const std::string SdkPause = "SdkPause";
const std::string SdkResume = "SdkResume";
const std::string SdkPlayWithStartTime = "SdkPlayWithStartTime";

namespace bhaptics {

    struct BhapticsDevice {
        int Position;
        std::string DeviceName;
        std::string Address ;
        bool Connected;
        bool Paired ;
        int Battery;
        bool AudioJackIn;
    };

    void getResponseMessage(std::string message);

    std::vector<BhapticsDevice> getConnectedDevices();
    std::vector<std::string> getServerEventNameList();

    class Haptic {
    private:
    public:
        static bool sendAuthentication();

        static bool registryAndInit(const char* sdkAPIKey, const char* workspaceId, const char* initJsonMessage);
        static bool registryAndInitHost(const char* sdkAPIKey, const char* workspaceId, const char* initJsonMessage, const char* url);
        static bool reInitMessage(const char* sdkAPIKey, const char* workspaceId, const char* initJsonMessage);
        static int play(const char* key);
        static int playParam(const char* key, int requestId, float intensity, float duration, float angleX, float offsetY);
        static int playWithStartTime(const char* key, int requestId, int startMillis, float intensity, float duration, float angleX, float offsetY);
        static void playWithoutResult(const char* key, float intensity, float duration, float angleX, float offsetY);
        static int playDot(int requestId, int position, int duration, int* motorValues, int motorValueLen);
        static int playPath(int requestId, int position, int durationMillis, float* xValues, float* yValues, int* intensityValues, int Len);
        static int playWaveform(int requestId, int position, int* motorValues, int* playTimeValues, int* shapeValues, int repeatCount, int motorLen);
        static int playLoop(const char* eventId, int requestId, float intensity, float duration, float angleX, float offsetY, int interval, int maxCount);
        static int getEventTime(const char* eventId);

        static int pause(const char* eventId);
        static bool resume(const char* eventId);

        static bool stop(int requestKey);
        static bool stopByEventId(const char* eventId);
        static bool stopAll();


        static bool isbHapticsConnected(int position);
        static bool isPlaying();
        static bool isPlayingByRequestId(const int key);
        static bool isPlayingByEventId(const char* eventId);


        static const char* getDeviceInfoJson();
        static const char* getHapticMappingsJson();

        static bool ping(const char* address);
        static bool pingAll();
        static bool swapPosition(const char* address);
        static bool setDeviceVsm(const char* address, int vsm);
    };
}


#endif //HAPTIC_LIBRARY_BHAPTICSHAPTIC_H
