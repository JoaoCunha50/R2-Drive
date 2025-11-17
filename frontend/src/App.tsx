import { MantineProvider } from '@mantine/core'
import '@mantine/core/styles.css'
import Context from './context/Context'
import AppRouter from './router/Router'
import { theme, variableResolver } from './styles/Theme'

function App() {
    return (
        <MantineProvider
            theme={theme}
            cssVariablesResolver={variableResolver}
            defaultColorScheme="auto"
        >
            <Context>
                <AppRouter />
            </Context>
        </MantineProvider>
    )
}

export default App
