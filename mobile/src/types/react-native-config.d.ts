declare module 'react-native-config' {
  export interface NativeConfig {
    BASE_API_URL: string
    PUSHER_API_KEY: string
    PUSHER_CLUSTER: string
  }

  export const Config: NativeConfig
  export default Config
}
