'use client'
import { useState } from "react"
import ArchProp from "./archetypesProps"
import ClanProps from "./clanprops"
import dynamic from "next/dynamic"


const ConceptNoSSR = dynamic(() => import('./concept'), {ssr: false})

export default function General(){
    // TODO: Make function that sets char and state and pass to components instead of the two functions
    let [char, setChar] = useState({} as Partial<MasqueradeChar>)
    let [state, setState] = useState("general")


    return(
         <div>
           {
            state == "general" && 
            <div>
                <h1> Let's get this for character creation</h1>
                <ConceptNoSSR 
                    setChar={setChar}
                    setState = {setState}
                    clan={<ClanProps />}
                    arc={<ArchProp />}></ConceptNoSSR>
            </div>
           }
   

           {
            state == "concept" && 
            <div>
                <h1> We did it</h1>
                <p>{char.clan} {char.concept}</p>
            </div>
           } 

        </div>
)}