import Config from 'react-native-config'

export const AppConfig = {
  APP_NAME: 'Example 2FA App',
  API_BASE_URL: Config.BASE_API_URL,
  PUSHER_API_KEY: Config.PUSHER_API_KEY,
  PUSHER_CLUSTER: Config.PUSHER_CLUSTER || 'ap1',
}
