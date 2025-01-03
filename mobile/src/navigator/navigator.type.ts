import React from 'react'
import {
  NativeStackNavigationOptions,
  NativeStackNavigationProp,
  NativeStackScreenProps,
} from '@react-navigation/native-stack'

export type RootNavigatorParamList = {
  home_screen: undefined
}

export type ScreenType = {
  label?: string
  name: keyof Partial<RootNavigatorParamList>
  component: React.ComponentType<object> | (() => JSX.Element)
  options?: NativeStackNavigationOptions
}

export type NavigationProps = NativeStackNavigationProp<
  Partial<RootNavigatorParamList>
>

export type AppStackScreenProps<T extends keyof RootNavigatorParamList> =
  NativeStackScreenProps<RootNavigatorParamList, T>
