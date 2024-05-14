'use client'
import AttrProps from "./attrProp"
import dynamic from "next/dynamic"

const AttributeNoSSR = dynamic(() => import('./attributes'))

export default function Attributes(){
    function placeholder(){

    }

    return (
        <div>
           <AttributeNoSSR 
            setChar={placeholder} 
           setState={placeholder} 
           attr={<AttrProps />}></AttributeNoSSR>
            {/* Todo section attrs by categories. Figure out how to distribute points
                Maybe choose primary category => on click fetch the attributes distribute points. Choose secondary then do x and then third submit reuse same flow for step 3*/}
        </div>
    )
}