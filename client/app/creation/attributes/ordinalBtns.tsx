'use client'
import React, {FormEvent,useState} from "react"
import { useRouter } from "next/navigation"


const OrdinalBtns: React.FC = (props) => {
    const router = useRouter()
    const [buttons, setButtons] = useState([] as JSX.Element[])
    const [display, setDisplay] = useState("primary")
    const [slug, setSlug] = useState("")
    let traits: {[k:string]: string} = {
        "physical": "",
        "mental": "",
        "social": "",
        "slug": ""
    }

    function restartAssign(e: FormEvent<HTMLButtonElement>) {
        e.preventDefault()
        traits = {"physical": "", "mental": "", "social": "", "slug": ""}
        setDisplay("primary")
        setButtons([])
    }


    function assignOrdinals(e: FormEvent<HTMLButtonElement>, ordinal: string){
        e.preventDefault()
        const trait = e.currentTarget.name
        traits[trait] = ordinal

        if (ordinal == "primary"){
            traits.slug = trait
            for (let k in traits) {
                if (traits[k] == ""){
                    let array = buttons
                    array.push(<button name={k} onClick={(e) =>assignOrdinals(e, "secondary")}>{k}</button>
                    )
                    setButtons(array) 
                } 
            }
            setDisplay("secondary")
        } else {
            for(let k in traits) {
                if (traits[k] == "") {
                    traits.slug = `${traits.slug}/${trait}/${k}`
                    traits[k] = "tertiary"
                }
            }
            setDisplay("confirm")
            // Reminder: using setState outside of useEffect misses the first
            // call of setSlug. Cant use traits.slug directly
            setSlug(traits.slug)
        }
    }

    return(
        <div>
                {
                    (display == "primary") && <div className="primary">
                        <h3>Primary Trait</h3>
                        <button name="physical" onClick={(e) =>assignOrdinals(e, "primary")}>Physical</button>
                        <button name="mental" onClick={(e) =>assignOrdinals(e, "primary")}>Mental</button>
                        <button name="social" onClick={(e) =>assignOrdinals(e, "primary")}>Social</button>
                    </div>
                }
                {
                    (display == "secondary") && <div>
                        <h3>Secondary Trait</h3>
                        {...buttons}
                    </div>
                }
                {
                    (display == "confirm") && <div>
                        <h3>Is this Correct?</h3>
                        <button onClick={(e) =>restartAssign(e)}>Nope</button>
                        <button onClick={() => router.push(`/creation/concept/true/${slug}`)}>Ye</button>

                    </div>
                }

        </div>
    )
}

export default OrdinalBtns