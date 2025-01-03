import Axios from '@/api/axios'
import { AxiosResponse } from 'axios'

export interface IUser {
  username: string
  fcmToken: string
}

export const UserAPI = {
  saveFcmToken: async (body: IUser): Promise<AxiosResponse<IUser>> => {
    const response = await Axios.post('/user', body)
    return response
  },
}
