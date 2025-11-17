import type { AnyObject, ApiResponse } from '../../types/GeneralTypes'
import type { User } from '../../types/User'
import { client } from '../axios/AxiosConfig'

export const Login = async (
    data: AnyObject
): Promise<ApiResponse<{ user: User; token: string }>> =>
    client.post('/users/login', data)
