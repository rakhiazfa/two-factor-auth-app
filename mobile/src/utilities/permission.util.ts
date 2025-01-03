import notifee, { AuthorizationStatus } from '@notifee/react-native'
import { PermissionsAndroid, Platform } from 'react-native'

export async function requestUserPermission() {
  const settings = await notifee.requestPermission()
  if (settings.authorizationStatus >= AuthorizationStatus.AUTHORIZED) {
    return true
  } else {
    return false
  }
}

export async function checkNotificationPermission(): Promise<boolean> {
  const settings = await notifee.requestPermission({
    sound: true,
    announcement: true,
  })
  if (settings.authorizationStatus) {
    // Alert.alert('User has notification permissions enabled')
    return true
  } else {
    return false
  }
}

const requestNotificationPermissionAndroid = async (): Promise<boolean> => {
  if (Platform.OS === 'android') {
    try {
      const result = await PermissionsAndroid.request(
        PermissionsAndroid.PERMISSIONS.POST_NOTIFICATIONS
      )
      if (result === 'granted') {
        return true
      } else if (result === 'denied') {
        return false
      }
    } catch (error) {
      return false
    }
  }
  return true
}

export const PermissionUtils = {
  requestUserPermission,
  checkNotificationPermission,
  requestNotificationPermissionAndroid,
}
