'use client'

import { useEffect } from "react"

// TODO: Improve ErrorBoundaries

export default function ErrorBoundary({
    error,
    reset,
}: {
    error: Error & {digest?: string}
    reset: () => void
}) {
    useEffect(() => {
        console.error(error)
    }, [error])

    return (
        <div>
            <h2> We have run into an error!</h2>
            <p> {error?.message}</p>
            <button onClick={() => reset()}>Attempt Recovery</button>
        </div>
    )
}