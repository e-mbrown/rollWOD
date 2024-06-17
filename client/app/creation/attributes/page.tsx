import React from "react"
import AttrOrdinal from "./attrOrdinal"

const Attribute: React.FC<{user: string}> = (params)  => {
        console.log(params)
    return (
        <div key="attribute">
            <h1>Character Attributes</h1>
                    <p>
                    Attributes are the natural abilities and raw capabilities a
                    character is made of. How strong is a character? How attractive? How
                    quick? How smart? Attributes take all these questions and more into
                    account. All Vampire characters have nine Attributes, which are
                    divided into three categories: Physical (Strength, Dexterity,
                    Stamina), Social (Charisma, Manipulation, Appearance), and Mental
                    (Perception, Intelligence, Wits).
                    </p>
                    <p>
                    First, the player must select which group of Attributes is his
                    character’s strong suit (primary). The player then selects the group in which the character is average (secondary). Finally, the remaining group is designated as the character’s weak point
                    (tertiary).
                    </p>
                    <AttrOrdinal params={params}/>
                    
        </div>
    )
}

export default Attribute