#include "steam.h"

static void *vtable_ISteamUtils006[] = {
    SteamAPI_ISteamUtils_GetSecondsSinceAppActive,
    SteamAPI_ISteamUtils_GetSecondsSinceComputerActive,
    SteamAPI_ISteamUtils_GetConnectedUniverse,
    SteamAPI_ISteamUtils_GetServerRealTime,
    SteamAPI_ISteamUtils_GetIPCountry,
    SteamAPI_ISteamUtils_GetImageSize,
    SteamAPI_ISteamUtils_GetImageRGBA,
    SteamAPI_ISteamUtils_GetCSERIPPort,
    SteamAPI_ISteamUtils_GetCurrentBatteryPower,
    SteamAPI_ISteamUtils_GetAppID,
    SteamAPI_ISteamUtils_SetOverlayNotificationPosition,
    SteamAPI_ISteamUtils_IsAPICallCompleted,
    SteamAPI_ISteamUtils_GetAPICallFailureReason,
    SteamAPI_ISteamUtils_GetAPICallResult,
    SteamAPI_ISteamUtils_RunFrame,
    SteamAPI_ISteamUtils_GetIPCCallCount,
    SteamAPI_ISteamUtils_SetWarningMessageHook,
    SteamAPI_ISteamUtils_IsOverlayEnabled,
    SteamAPI_ISteamUtils_BOverlayNeedsPresent,
    SteamAPI_ISteamUtils_CheckFileSignature,
    SteamAPI_ISteamUtils_ShowGamepadTextInput,
    SteamAPI_ISteamUtils_GetEnteredGamepadTextLength,
    SteamAPI_ISteamUtils_GetEnteredGamepadTextInput,
    SteamAPI_ISteamUtils_GetSteamUILanguage,
    SteamAPI_ISteamUtils_IsSteamRunningInVR,
};

static void *vtable_ISteamUtils007[] = {
    SteamAPI_ISteamUtils_GetSecondsSinceAppActive,
    SteamAPI_ISteamUtils_GetSecondsSinceComputerActive,
    SteamAPI_ISteamUtils_GetConnectedUniverse,
    SteamAPI_ISteamUtils_GetServerRealTime,
    SteamAPI_ISteamUtils_GetIPCountry,
    SteamAPI_ISteamUtils_GetImageSize,
    SteamAPI_ISteamUtils_GetImageRGBA,
    SteamAPI_ISteamUtils_GetCSERIPPort,
    SteamAPI_ISteamUtils_GetCurrentBatteryPower,
    SteamAPI_ISteamUtils_GetAppID,
    SteamAPI_ISteamUtils_SetOverlayNotificationPosition,
    SteamAPI_ISteamUtils_IsAPICallCompleted,
    SteamAPI_ISteamUtils_GetAPICallFailureReason,
    SteamAPI_ISteamUtils_GetAPICallResult,
    SteamAPI_ISteamUtils_RunFrame,
    SteamAPI_ISteamUtils_GetIPCCallCount,
    SteamAPI_ISteamUtils_SetWarningMessageHook,
    SteamAPI_ISteamUtils_IsOverlayEnabled,
    SteamAPI_ISteamUtils_BOverlayNeedsPresent,
    SteamAPI_ISteamUtils_CheckFileSignature,
    SteamAPI_ISteamUtils_ShowGamepadTextInput,
    SteamAPI_ISteamUtils_GetEnteredGamepadTextLength,
    SteamAPI_ISteamUtils_GetEnteredGamepadTextInput,
    SteamAPI_ISteamUtils_GetSteamUILanguage,
    SteamAPI_ISteamUtils_IsSteamRunningInVR,
    SteamAPI_ISteamUtils_SetOverlayNotificationInset,
};

static void *vtable_ISteamUtils008[] = {
    SteamAPI_ISteamUtils_GetSecondsSinceAppActive,
    SteamAPI_ISteamUtils_GetSecondsSinceComputerActive,
    SteamAPI_ISteamUtils_GetConnectedUniverse,
    SteamAPI_ISteamUtils_GetServerRealTime,
    SteamAPI_ISteamUtils_GetIPCountry,
    SteamAPI_ISteamUtils_GetImageSize,
    SteamAPI_ISteamUtils_GetImageRGBA,
    SteamAPI_ISteamUtils_GetCSERIPPort,
    SteamAPI_ISteamUtils_GetCurrentBatteryPower,
    SteamAPI_ISteamUtils_GetAppID,
    SteamAPI_ISteamUtils_SetOverlayNotificationPosition,
    SteamAPI_ISteamUtils_IsAPICallCompleted,
    SteamAPI_ISteamUtils_GetAPICallFailureReason,
    SteamAPI_ISteamUtils_GetAPICallResult,
    SteamAPI_ISteamUtils_RunFrame,
    SteamAPI_ISteamUtils_GetIPCCallCount,
    SteamAPI_ISteamUtils_SetWarningMessageHook,
    SteamAPI_ISteamUtils_IsOverlayEnabled,
    SteamAPI_ISteamUtils_BOverlayNeedsPresent,
    SteamAPI_ISteamUtils_CheckFileSignature,
    SteamAPI_ISteamUtils_ShowGamepadTextInput,
    SteamAPI_ISteamUtils_GetEnteredGamepadTextLength,
    SteamAPI_ISteamUtils_GetEnteredGamepadTextInput,
    SteamAPI_ISteamUtils_GetSteamUILanguage,
    SteamAPI_ISteamUtils_IsSteamRunningInVR,
    SteamAPI_ISteamUtils_SetOverlayNotificationInset,
    SteamAPI_ISteamUtils_IsSteamInBigPictureMode,
    SteamAPI_ISteamUtils_StartVRDashboard,
};

static void *vtable_ISteamUtils009[] = {
    SteamAPI_ISteamUtils_GetSecondsSinceAppActive,
    SteamAPI_ISteamUtils_GetSecondsSinceComputerActive,
    SteamAPI_ISteamUtils_GetConnectedUniverse,
    SteamAPI_ISteamUtils_GetServerRealTime,
    SteamAPI_ISteamUtils_GetIPCountry,
    SteamAPI_ISteamUtils_GetImageSize,
    SteamAPI_ISteamUtils_GetImageRGBA,
    SteamAPI_ISteamUtils_GetCSERIPPort, // private
    SteamAPI_ISteamUtils_GetCurrentBatteryPower,
    SteamAPI_ISteamUtils_GetAppID,
    SteamAPI_ISteamUtils_SetOverlayNotificationPosition,
    SteamAPI_ISteamUtils_IsAPICallCompleted,
    SteamAPI_ISteamUtils_GetAPICallFailureReason,
    SteamAPI_ISteamUtils_GetAPICallResult,
    SteamAPI_ISteamUtils_RunFrame, // private
    SteamAPI_ISteamUtils_GetIPCCallCount,
    SteamAPI_ISteamUtils_SetWarningMessageHook,
    SteamAPI_ISteamUtils_IsOverlayEnabled,
    SteamAPI_ISteamUtils_BOverlayNeedsPresent,
    SteamAPI_ISteamUtils_CheckFileSignature,
    SteamAPI_ISteamUtils_ShowGamepadTextInput,
    SteamAPI_ISteamUtils_GetEnteredGamepadTextLength,
    SteamAPI_ISteamUtils_GetEnteredGamepadTextInput,
    SteamAPI_ISteamUtils_GetSteamUILanguage,
    SteamAPI_ISteamUtils_IsSteamRunningInVR,
    SteamAPI_ISteamUtils_SetOverlayNotificationInset,
    SteamAPI_ISteamUtils_IsSteamInBigPictureMode,
    SteamAPI_ISteamUtils_StartVRDashboard,
    SteamAPI_ISteamUtils_IsVRHeadsetStreamingEnabled,
    SteamAPI_ISteamUtils_SetVRHeadsetStreamingEnabled,
    SteamAPI_ISteamUtils_IsSteamChinaLauncher,
    SteamAPI_ISteamUtils_InitFilterText,
    SteamAPI_ISteamUtils_FilterText,
    SteamAPI_ISteamUtils_GetIPv6ConnectivityState,
};

static void *vtable_ISteamUtils010[] = {
    SteamAPI_ISteamUtils_GetSecondsSinceAppActive,
    SteamAPI_ISteamUtils_GetSecondsSinceComputerActive,
    SteamAPI_ISteamUtils_GetConnectedUniverse,
    SteamAPI_ISteamUtils_GetServerRealTime,
    SteamAPI_ISteamUtils_GetIPCountry,
    SteamAPI_ISteamUtils_GetImageSize,
    SteamAPI_ISteamUtils_GetImageRGBA,
    SteamAPI_ISteamUtils_GetCSERIPPort, // private
    SteamAPI_ISteamUtils_GetCurrentBatteryPower,
    SteamAPI_ISteamUtils_GetAppID,
    SteamAPI_ISteamUtils_SetOverlayNotificationPosition,
    SteamAPI_ISteamUtils_IsAPICallCompleted,
    SteamAPI_ISteamUtils_GetAPICallFailureReason,
    SteamAPI_ISteamUtils_GetAPICallResult,
    SteamAPI_ISteamUtils_RunFrame, // private
    SteamAPI_ISteamUtils_GetIPCCallCount,
    SteamAPI_ISteamUtils_SetWarningMessageHook,
    SteamAPI_ISteamUtils_IsOverlayEnabled,
    SteamAPI_ISteamUtils_BOverlayNeedsPresent,
    SteamAPI_ISteamUtils_CheckFileSignature,
    SteamAPI_ISteamUtils_ShowGamepadTextInput,
    SteamAPI_ISteamUtils_GetEnteredGamepadTextLength,
    SteamAPI_ISteamUtils_GetEnteredGamepadTextInput,
    SteamAPI_ISteamUtils_GetSteamUILanguage,
    SteamAPI_ISteamUtils_IsSteamRunningInVR,
    SteamAPI_ISteamUtils_SetOverlayNotificationInset,
    SteamAPI_ISteamUtils_IsSteamInBigPictureMode,
    SteamAPI_ISteamUtils_StartVRDashboard,
    SteamAPI_ISteamUtils_IsVRHeadsetStreamingEnabled,
    SteamAPI_ISteamUtils_SetVRHeadsetStreamingEnabled,
    SteamAPI_ISteamUtils_IsSteamChinaLauncher,
    SteamAPI_ISteamUtils_InitFilterText,
    SteamAPI_ISteamUtils_FilterText,
    SteamAPI_ISteamUtils_GetIPv6ConnectivityState,
    SteamAPI_ISteamUtils_IsSteamRunningOnSteamDeck,
    SteamAPI_ISteamUtils_ShowFloatingGamepadTextInput,
    SteamAPI_ISteamUtils_SetGameLauncherMode,
    SteamAPI_ISteamUtils_DismissFloatingGamepadTextInput,
};

VtableEntry vtable_ISteamUtils[] = {
    {6, vtable_ISteamUtils006, sizeof(vtable_ISteamUtils006), 0},
    {7, vtable_ISteamUtils007, sizeof(vtable_ISteamUtils007), 0},
    {8, vtable_ISteamUtils008, sizeof(vtable_ISteamUtils008), 0},
    {9, vtable_ISteamUtils009, sizeof(vtable_ISteamUtils009), VTABLE_HAS_PRIVATE},
    {10, vtable_ISteamUtils010, sizeof(vtable_ISteamUtils010), VTABLE_HAS_PRIVATE},
    {0, NULL, 0, 0},
};
