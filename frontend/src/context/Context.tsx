import i18n from 'i18next'
import LanguageDetector from 'i18next-browser-languagedetector'
import React, {
    createContext,
    useEffect,
    useState,
    type Dispatch,
    type ReactNode,
    type SetStateAction,
} from 'react'
import { initReactI18next } from 'react-i18next/initReactI18next'
import { useNavigate } from 'react-router'
import { api } from '../internal/config/api'
import LoadingPage from '../pages/LoadingPage/LoadingPage'
import type { User } from '../types/User'

const FallBackInfo = {
    languages: ['en', 'pt'],
    translations: {},
    defaultLanguage: 'en',
}

export interface UserContextType {
    user: User | null
    setUser: Dispatch<SetStateAction<User | null>>
}

export interface InfoContextType {
    info: any
    setInfo: Dispatch<SetStateAction<any>>
}

export const UserContext = createContext<UserContextType>({
    user: null,
    setUser: () => null,
})

export const InfoContext = createContext<InfoContextType>({
    info: FallBackInfo,
    setInfo: () => [],
})

function initTranslations(
    translations: any,
    defaultLanguage: string,
    langs: string[]
) {
    const resources: any = {}
    if (Array.isArray(translations)) {
        translations.map((item) => {
            const lang = item.lang
            const tag = item.tag
            const translationData = item.translation

            if (!resources[lang]) {
                resources[lang] = { translation: {} }
            }

            resources[lang].translation[tag] = translationData
        })
    }

    i18n.use(LanguageDetector)
        .use(initReactI18next)
        .init({
            returnNull: false,
            fallbackLng: defaultLanguage,
            interpolation: { escapeValue: false },
            resources,
            supportedLngs: langs,
        })
}

export function ContextProvider({ children }: { children: ReactNode }) {
    const [user, setUser] = useState<User | null>(null)
    const [info, setInfo] = useState<any>([])
    const [isLoading, setIsLoading] = useState(true)
    const navigate = useNavigate()

    useEffect(() => {
        const init = async () => {
            setIsLoading(true)
            try {
                const { data, status } = await api.get('/info')

                if (status === 200) {
                    setInfo(data)
                    initTranslations(
                        data?.translations,
                        data?.defaultLanguage,
                        data?.languagesSupported
                    )

                    if (data?.user) {
                        setUser(data.user)
                    } else {
                        localStorage.clear()
                        setUser(null)
                        navigate('/login')
                    }
                }
            } catch (err) {
                localStorage.clear()
                setUser(null)
                navigate('/login')
            } finally {
                setIsLoading(false)
            }
        }

        init()
    }, [])

    if (isLoading) return <LoadingPage />

    return (
        <InfoContext.Provider value={{ info, setInfo }}>
            <UserContext.Provider value={{ user, setUser }}>
                {children}
            </UserContext.Provider>
        </InfoContext.Provider>
    )
}

const Context = React.memo(ContextProvider)
export default Context
