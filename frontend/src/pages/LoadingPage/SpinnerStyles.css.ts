import { globalStyle, keyframes, style } from '@vanilla-extract/css'

// Global style para o root
globalStyle('#root', {
    maxWidth: '1280px',
    margin: '0 auto',
    padding: '2rem',
    textAlign: 'center',
})

const spin = keyframes({
    from: {
        transform: 'rotate(0deg)',
    },
    to: {
        transform: 'rotate(360deg)',
    },
})

export const spinner = style({
    animation: `${spin} 1s linear infinite`,
    color: '#3498db',
})
