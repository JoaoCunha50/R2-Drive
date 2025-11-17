import { LoaderCircle as Spinner } from 'lucide-react'
import { spinner } from './SpinnerStyles.css'

export default function LoadingPage() {
    return (
        <div>
            <Spinner className={spinner} />
            <h1>Loading...</h1>
        </div>
    )
}
