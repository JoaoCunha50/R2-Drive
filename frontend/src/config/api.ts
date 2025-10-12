import axios from 'axios'
import { API } from './Environment'

export const api = axios.create({
    baseURL: API,
})
