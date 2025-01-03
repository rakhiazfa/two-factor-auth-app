/**
 * @format
 */

import { AppRegistry } from 'react-native'
import { name as appName } from './app.json'
import {
  messagingBackgroundMessageHandler,
  messagingForegroundMessageHandler,
  messagingGetInitialNotification,
  messagingOnNotificationOpenedApp,
  notifeeOnBackgroundEvent,
} from '@/services/notification.service'
import App from './src/app'

// Android background handler
messagingBackgroundMessageHandler()

// On notification opened app
messagingOnNotificationOpenedApp()

// Get initial notification
messagingGetInitialNotification()

// Notification foreground handler
messagingForegroundMessageHandler()

// notifee background event
notifeeOnBackgroundEvent()

const HeadlessCheck = ({ isHeadless }) => {
  if (isHeadless) {
    // App has been launched in the background by iOS, ignore
    return null
  }
  return <App />
}

AppRegistry.registerComponent(appName, () => HeadlessCheck)
