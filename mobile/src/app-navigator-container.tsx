import React, { memo, useRef } from 'react'
import {
  NavigationContainer,
  useNavigationContainerRef,
} from '@react-navigation/native'

// root stack navigator
import AppRootStackNavigator from '@/navigator/app-root-stack.navigator'

const AppNavigatorContainer = () => {
  const navigationRef = useNavigationContainerRef()
  const routeNameRef = useRef<string | undefined>()

  return (
    <NavigationContainer
      ref={navigationRef}
      onReady={() => {
        routeNameRef.current = navigationRef?.getCurrentRoute()?.name
      }}
    >
      <AppRootStackNavigator />
    </NavigationContainer>
  )
}

AppNavigatorContainer.displayName = 'AppNavigatorContainer'

export default memo(AppNavigatorContainer)
