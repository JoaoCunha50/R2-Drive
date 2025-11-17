import axios from 'axios'
import { AuthTokenKey, LangKey } from '../storage/StorageKeys'
import { API } from './Environment'

export const client = axios.create({
    baseURL: API,
})

async function mergeCommonHeaders(headers: any) {
    const token = localStorage.getItem(AuthTokenKey)
    const language = localStorage.getItem(LangKey)

    return {
        Authorization: token ? `Bearer ${token}` : undefined,
        'Accept-language': language ? language : '',
        'Content-Type': 'application/json',
        ...headers,
    }
}

client.interceptors.request.use(async (request: any) => {
    const headers = await mergeCommonHeaders(request?.headers)
    request.headers = headers

    return request
})

client.interceptors.response.use(
    async (response: any) => response.data,
    (error: any) => {
        try {
            const exp = {
                ...(error && error.response && error.response.data
                    ? error.response.data
                    : ''),
                status:
                    error && error.response && error.response.status
                        ? error.response.status
                        : '',
            }
            return exp
        } catch (e) {
            return Promise.reject(e)
        }
    }
)
