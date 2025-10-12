import { MantineProvider } from '@mantine/core'
import '@mantine/core/styles.css'
import AppRouter from './router/Router'
import { theme } from './styles/theme'

function App() {
    return (
        <MantineProvider theme={theme} defaultColorScheme="dark">
            <AppRouter />
        </MantineProvider>
    )
}

export default App
