import React from "react";
import AttrBtns from "./attrBtns";

const AttrDivs: React.FC<{ordinal: string, data: Characteristic[]}> = (props) => {
    let divs: JSX.Element[] = []
    // Have Attrbtns return 3 buttons that are controlled within the component,(total and what not) but can be placed independently within each div below

    props.data.forEach(attr => {
        divs.push(
        <div className={props.ordinal} key={attr.name}>
            <h2>{attr.name}</h2>
            <p>{attr.description}</p>
        </div>
        )
    });
    return (
        <>
            {...divs}
            <AttrBtns ordinal={props.ordinal} data={props.data}/>
        </>
    )
}


export default AttrDivs