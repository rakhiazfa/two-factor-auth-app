import { Fragment, JSX } from 'react'

// react navigation
import { createNativeStackNavigator } from '@react-navigation/native-stack'

// interfaces
import { RootNavigatorParamList, ScreenType } from './navigator.type'

// screens
import HomeScreen from '@/screens/home.screen'

const rootScreen: Array<ScreenType> = [
  { name: 'home_screen', component: HomeScreen },
]

const RootStack = createNativeStackNavigator<RootNavigatorParamList>()

const RootStackNavigator = (): JSX.Element | null => {
  return (
    <Fragment>
      <RootStack.Navigator
        screenOptions={{ headerShown: false, gestureEnabled: true }}
      >
        {rootScreen.map(x => {
          return (
            <RootStack.Screen
              key={x.name}
              component={x.component}
              name={x.name}
              options={x.options}
            />
          )
        })}
      </RootStack.Navigator>
    </Fragment>
  )
}

export default RootStackNavigator
