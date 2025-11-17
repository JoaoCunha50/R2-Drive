export interface ApiResponse<T> {
    success: boolean
    data: T
    title?: string
    message?: string
}

export interface AnyObject {
    [key: string]: any
}
