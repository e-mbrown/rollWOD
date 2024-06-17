import React, { FormEvent, SetStateAction } from "react"

function submitConcept(e: FormEvent<HTMLFormElement>, setChar: React.Dispatch<SetStateAction<Partial<MasqueradeChar>>>) {
    e.preventDefault()
    const formData = new FormData(e.currentTarget)

    const obj:Partial<MasqueradeChar> = {
        clan: String(formData.get("clan")),
        demeanor: String(formData.get("demeanor")),
        concept: String(formData.get("concept")),
    }

    setChar(obj)
}




export{ submitConcept}