import { createTheme, type CSSVariablesResolver } from '@mantine/core'
import { themeToVars } from '@mantine/vanilla-extract'
import * as colors from './colors'

export const theme = createTheme({
    fontFamily: 'system-ui, Avenir, Helvetica, Arial, sans-serif',
    white: colors.White,
    cursorType: 'pointer',
    primaryColor: 'primary',
    colors: {
        primary: [
            colors.PurplePrimary,
            colors.PurplePrimary,
            colors.PurplePrimary,
            colors.PurplePrimary,
            colors.PurplePrimary,
            colors.PurplePrimary,
            colors.PurplePrimary,
            colors.PurplePrimary,
            colors.PurplePrimary,
            colors.PurplePrimary,
        ],
    },
    focusRing: 'auto',
    defaultRadius: 'lg',
    defaultGradient: {
        from: colors.PurplePrimary,
        to: colors.BlueSecondary,
        deg: 45,
    },
    breakpoints: {
        mobile: '600px',
        tablet: '768px',
        desktop: '1024px',
    },
})

export const variableResolver: CSSVariablesResolver = () => ({
    variables: {
        '--mantine-color-body': colors.BackgroundGray,
    },
    light: {
        '--mantine-color-body': colors.White,
    },
    dark: {
        '--mantine-color-body': colors.BackgroundGray,
    },
})

export const vars = themeToVars(theme)
