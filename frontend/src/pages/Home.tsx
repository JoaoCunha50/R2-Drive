import { useContext } from 'react'
import { InfoContext } from '../context/Context'

export default function Home() {
    const { info } = useContext(InfoContext)

    return (
        <div>
            <h1>Home</h1>
        </div>
    )
}
