import axios from 'axios'
import { AuthTokenKey } from '../utils/storageKeys'
import { API } from './Environment'

export const api = axios.create({
    baseURL: API,
})

api.interceptors.request.use((config) => {
    const token = localStorage.getItem(AuthTokenKey)

    if (token) {
        config.headers.Authorization = `Bearer ${token}`
    }

    return config
})
