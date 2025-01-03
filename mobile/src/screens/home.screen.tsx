import { ImageAssets } from '@/assets'
import { pusher } from '@/services/pusher.service'
import { PusherEvent } from '@pusher/pusher-websocket-react-native'
import { useEffect } from 'react'
import { Alert, Image, Text } from 'react-native'
import { View } from 'react-native'

import { AppConfig } from '@/config'

const HomeScreen = () => {
  useEffect(() => {
    pusher.subscribe({
      channelName: 'auth',
      onEvent: (event: PusherEvent) => {
        Alert.alert(`Event received: ${event}`)
        // const data = JSON.parse(event.data)
      },
    })
  }, [])

  return (
    <View style={{ flex: 1, alignItems: 'center', justifyContent: 'center' }}>
      <Text style={{ fontFamily: 'Jost', fontSize: 20, fontWeight: '500' }}>
        {AppConfig.APP_NAME}
      </Text>
      <Image
        source={ImageAssets.img1}
        style={{ height: 200, width: 200, resizeMode: 'contain' }}
      />
    </View>
  )
}

export default HomeScreen
