import React from "react"
import AttrBtns from "./attrBtns"
import { AttrProps } from "../[...user]/serverActions"

const AttrOrdinal = ({params}: {params: {user:string}}) => {
    console.log(params.user)
    const ordinal = {
        "primary": params.user[2],
        "secondary": params.user[3],
        "tertiary": params.user[4]
    }
    

    return (
       <div>
            pain
                {
                    (params.user[1] == "false") && 
                        <AttrBtns/>
                }
                {
                    (params.user[1] == "true") && 
                    // TODO: Work on AttrProp and final form
                    <AttrProps data = {ordinal}/>
                } 
                
       </div> 
    )
}

export default AttrOrdinal