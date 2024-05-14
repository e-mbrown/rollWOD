import React, { FormEvent } from "react"

const AttrForm: React.FC<{setChar:Function, setState: Function, attr: React.ReactNode}> = (props) => {
    function submitAttr(e:FormEvent<HTMLFormElement>) {
        e.preventDefault()
        const formData = new FormData(e.currentTarget)

        const obj:Partial<MasqueradeChar> = {

        }

        // props.setChar(obj)
        // props.setState("attributes")
    }

    return (
        <div>
            <form onSubmit={submitAttr}>
                <div>
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
                    character’s strong suit (primary). The player then selects the group
                    in which the character is average (secondary). Finally, the
                    remaining group is designated as the character’s    weak point
                    (tertiary).
                    </p>
                    {props.attr}
                    {/* Todo section attrs by categories. Figure out how to distribute points
                        Maybe choose primary category => on click fetch the attributes distribute points. Choose secondary then do x and then third submit reuse same flow for step 3*/}
                </div>  
            </form>
        </div>
    )
}

export default AttrForm

