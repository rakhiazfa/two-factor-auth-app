import React, { useCallback, useEffect, useState } from 'react'

import { Alert, Platform, StatusBar, TextInput } from 'react-native'

// gesture handler root view.
import { GestureHandlerRootView } from 'react-native-gesture-handler'

// safe area provider
import { SafeAreaProvider } from 'react-native-safe-area-context'

// bottom sheet modal provider.
import { BottomSheetModalProvider } from '@gorhom/bottom-sheet'

// react native screens.
import { enableScreens } from 'react-native-screens'

// messaging
import messaging from '@react-native-firebase/messaging'

// portal
import { PortalProvider } from '@gorhom/portal'

// navigator container
import AppNavigatorContainer from './app-navigator-container'

// utils
import { PermissionUtils } from '@/utilities/permission.util'

// apis
import { UserAPI } from '@/api/api'

// pusher
import initPusher from './services/pusher.service'

enableScreens()
initPusher()

const AppEntryPoint = () => {
  const [fcmToken, setFcmToken] = useState('')
  const initialRequestPermissions = async () => {
    Alert.alert(
      'This app needs to send you notification.',
      '',
      [
        {
          text: 'No',
          onPress: () => {
            Alert.alert('NOTIFICATION PERMISSION DENIED!')
            // storageUtils.save('NOTIFICATION_PERMISSION', 'denied')
          },
        },
        {
          text: 'Yes',
          onPress: async () => {
            // Alert.alert('OK Mantap!')
            if (Platform.OS === 'ios') {
              await PermissionUtils.requestUserPermission()
            }
            if (Platform.OS === 'android') {
              await PermissionUtils.requestNotificationPermissionAndroid()
            }
          },
        },
      ],
      { cancelable: false }
    )
  }

  const registerFCMToken = useCallback(async () => {
    try {
      // Register the device with FCM
      await messaging().registerDeviceForRemoteMessages()

      // Get the token
      const FCM_TOKEN = await messaging().getToken()
      setFcmToken(FCM_TOKEN)

      if (fcmToken) {
        try {
          await UserAPI.saveFcmToken({ username: Platform.OS, fcmToken })
          Alert.alert(`FCM TOKEN-> ${fcmToken}`)
          console.info('register token success')
        } catch (e) {
          console.info('failed to register token')
        }
      } else {
      }
    } catch (e) {}
  }, [])

  useEffect(() => {
    initialRequestPermissions()
  }, [])

  useEffect(() => {
    registerFCMToken()
  }, [])

  return (
    <GestureHandlerRootView style={{ flex: 1 }}>
      <PortalProvider>
        <BottomSheetModalProvider>
          <SafeAreaProvider>
            <StatusBar translucent backgroundColor='transparent' />
            <AppNavigatorContainer />
            <TextInput
              value={fcmToken}
              multiline={true}
              numberOfLines={4}
              style={{
                width: 280,
                alignSelf: 'center',
                borderWidth: 1,
                borderColor: '#ececec',
                borderRadius: 10,
              }}
            />
          </SafeAreaProvider>
        </BottomSheetModalProvider>
      </PortalProvider>
    </GestureHandlerRootView>
  )
}

export default AppEntryPoint
