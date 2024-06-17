'use client'
import React, { useState, FormEvent, ReactNode } from "react"
import { submitConcept } from "./functions"

 const CreationForm: React.FC<{slug:string, concept: ReactNode, attr: ReactNode}> = (props) => {
    let [char, setChar] = useState({} as Partial<MasqueradeChar>)
    let [state, setState] = useState(props.slug[0])
    console.log(props.slug)

    // function submitConcept(e: FormEvent<HTMLFormElement>) {
    //     e.preventDefault()
    //     const formData = new FormData(e.currentTarget)

    //     const obj:Partial<MasqueradeChar> = {
    //         clan: String(formData.get("clan")),
    //         demeanor: String(formData.get("demeanor")),
    //         concept: String(formData.get("concept")),
    //     }

    //     setChar(obj)
    //     // TODO: Can possibly lead to 
    //     AttrRedirect
    // }

    return(
        <div>
            {/* {
                (state =="general") && 

                <form onSubmit={(e) => submitConcept(e, setChar)}>
                    {props.concept}
                    <button type="submit">Submit</button>
                 </form> 
            } */}
             {/* {
                (state =="concept") && 
                 */}
                <form onSubmit={(e) => (e)}>
                    {props.attr}
                    {/* <button type="submit">Submit</button> */}
                 </form> 
            {/* } */}
            
        </div>
    )
}

export default CreationForm