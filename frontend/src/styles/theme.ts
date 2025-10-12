import { createTheme } from '@mantine/core'
import { themeToVars } from '@mantine/vanilla-extract'

export const theme = createTheme({
    fontFamily: 'system-ui, Avenir, Helvetica, Arial, sans-serif',
    colors: {
        // Define custom colors
        brand: [
            '#e8e9ff',
            '#c4c6ff',
            '#9fa3ff',
            '#7a80ff',
            '#535bf2',
            '#646cff', // primary
            '#4850d9',
            '#3640bf',
            '#2430a6',
            '#12208c',
        ],
    },

    primaryColor: 'brand',

    defaultRadius: 'md',

    components: {
        Button: {
            defaultProps: {
                radius: 'md',
            },
            styles: (theme: any) => ({
                root: {
                    padding: '0.6em 1.2em',
                    fontSize: '1em',
                    fontWeight: 500,
                    backgroundColor:
                        theme.colorScheme === 'dark' ? '#251764ff' : '#f9f9f9',
                    border: '1px solid transparent',
                    transition: 'border-color 0.25s',

                    '&:hover': {
                        borderColor: '#646cff',
                    },

                    '&:focus, &:focus-visible': {
                        outline: '4px auto -webkit-focus-ring-color',
                    },
                },
            }),
        },

        Anchor: {
            styles: {
                root: {
                    fontWeight: 500,
                    color: '#646cff',
                    textDecoration: 'inherit',

                    '&:hover': {
                        color: '#535bf2',
                    },
                },
            },
        },
    },
})

export const vars = themeToVars(theme)
