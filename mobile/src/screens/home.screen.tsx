import { ImageAssets } from '@/assets'
import { pusher } from '@/services/pusher.service'
import { PusherEvent } from '@pusher/pusher-websocket-react-native'
import { useCallback, useEffect, useMemo, useRef } from 'react'
import { Alert, Button, Image, StatusBar, Text } from 'react-native'
import { View } from 'react-native'

import { AppConfig } from '@/config'
import BottomSheet, {
  BottomSheetModal,
  BottomSheetView,
} from '@gorhom/bottom-sheet'
import SampleBottomSheet from '@/components/SampleBottomSheet'

const HomeScreen = () => {
  // const bottomSheetRef = useRef<BottomSheetModal | null>(null)

  // ref
  const bottomSheetRef = useRef<BottomSheet | null>(null)
  const bottomSheetRef2 = useRef<BottomSheet | null>(null)

  // variables
  const snapPoints = useMemo(() => ['25%', '50%', '90%'], [])

  // callbacks
  const handleSheetChanges = useCallback((index: number) => {
    console.log('handleSheetChanges', index)
  }, [])

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
      <StatusBar translucent barStyle='dark-content' />
      <Text style={{ fontFamily: 'Jost', fontSize: 20, fontWeight: '500' }}>
        {AppConfig.APP_NAME}
      </Text>
      <Image
        source={ImageAssets.img1}
        style={{ height: 200, width: 200, resizeMode: 'contain' }}
      />

      <Button
        title='Open bottom sheet'
        onPress={() => bottomSheetRef.current?.snapToIndex(1)}
      />
      <SampleBottomSheet ref={bottomSheetRef} />
      <Button
        title='Open bottom sheet 2'
        onPress={() => bottomSheetRef2.current?.snapToIndex(1)}
      />

      <BottomSheet
        ref={bottomSheetRef2}
        snapPoints={snapPoints}
        onChange={handleSheetChanges}
      >
        <BottomSheetView style={{}}>
          <Text>Awesome 🎉</Text>
        </BottomSheetView>
      </BottomSheet>
    </View>
  )
}

export default HomeScreen
