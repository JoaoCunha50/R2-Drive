import {
    createContext,
    useState,
    type Dispatch,
    type ReactNode,
    type SetStateAction,
} from 'react'
import type { Translation } from '../types/Translations'
import type { User } from '../types/User'

export interface UserContextType {
    user: User | null
    setUser: Dispatch<SetStateAction<User | null>>
}

export interface translationsContextType {
    translations: Translation[]
    setTranslations: Dispatch<SetStateAction<Translation[]>>
}

export const UserContext = createContext<UserContextType>({
    user: null,
    setUser: () => null,
})

export const TranslationsContext = createContext<translationsContextType>({
    translations: [],
    setTranslations: () => [],
})

export function ContextProvider({ children }: { children: ReactNode }) {
    const [user, setUser] = useState<User | null>(null)
    const [translations, setTranslations] = useState<Translation[]>([])

    return (
        <UserContext.Provider value={{ user, setUser }}>
            <TranslationsContext.Provider
                value={{ translations, setTranslations }}
            >
                {children}
            </TranslationsContext.Provider>
        </UserContext.Provider>
    )
}
