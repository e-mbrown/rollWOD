'use server'

import { getClient } from "@/app/apolloclient"
import { gql } from "@apollo/client"
import { BtnNum, GenButtons, AttrList} from "./utils"
import React from "react"
import { redirect } from "next/navigation"

const GET_ARCHETYPES = gql`
query getArchetypes {
   archetypes {
        name
        description
        sys
   }
}
`
const GET_CLANS = gql`
query getClans {
    clans {
        name
        description
    }
}
`
const GET_ATTRS = gql`
query getAttributes {
    charByType(type: attribute){
        name
        description
    }
}
`


const ArchProp = async () => {
    const client = getClient()
    const { data } = await client.query({
        query: GET_ARCHETYPES,
    })

    return (
        <div>
            {data.archetypes.map((arch: Archetype, idx: number) => (
                <div key={"arch" + idx}>
                    <h2>{arch.name}</h2>
                    <p>{arch.description}</p>
                    <p>{arch.sys}</p>
                    <input type="radio" name={"demeanor"} value={arch.name} /> 
                </div>
            ))}
        </div>
    )
}

const ClanProps = async () => {
    const client = getClient()
    const { data } =  await client.query({
        query: GET_CLANS,
        context:{}
})


    return (
        <div>
           {data.clans.map((clan: Clan, idx: number) => (
            <div key={"clan" + idx}>
                <h2>{clan.name}</h2>
                <p>{clan.description}</p> 
                <input type="radio" name={"clan"} value={clan.name}/> 
            </div>
           ))}
        </div>
    )
}


/// Irregular Server Actions

const AttrProps: React.FC<{data: {[k:string]: string}}> = async (props) => {
    let attributes: JSX.Element[] = []
    const client = getClient()
    const { data } = await client.query({
        query:GET_ATTRS,
        context:{},
})
    data.charByType.map((attr: Characteristic) =>{
        switch(attr.name) {
            case "charisma":
            case "manipulation":
            case "appearance":
                AttrList.social.push(attr)
                break
            case "perception":
            case "wits":
            case "intelligence":
                AttrList.mental.push(attr)
                break
            default:
                AttrList.physical.push(attr)
                break
        }
    })

    for (let attr in AttrList) {
        let ordinal = data[attr]
        AttrList[attr].forEach(attr => {
            let btns = GenButtons(attr.name)
            let element = <div className={ordinal} key={attr.name}>
                <h2>{attr.name}</h2>
                <p>{attr.description}</p>
                <label htmlFor={attr.name}>  {attr.name} value </label>
                { btns }
            </div>
            attributes.push(element)
    })};    
    return (
        <div>
            { ...attributes}
        </div>
    )
}


/// Redirects

const OrdinalRedirect = async (p: string, s: string, t:string) => {
    redirect(`/user/concept/${p}/${s}/${t}`)
}

const AttrRedirect = async () => {
    redirect(`user/`)
}

export {ArchProp, ClanProps, AttrProps, AttrRedirect, OrdinalRedirect}