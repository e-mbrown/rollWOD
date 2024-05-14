'use server'
import { getClient } from "@/app/apolloclient"
import { gql } from "@apollo/client"
const GET_ATTRS = gql`
query getAttributes {
    charByType(type: attribute){
        name
        description
    }
}
`

const AttrProps = async () => {
    const client = getClient()
    const { data } = await client.query({
        query:GET_ATTRS,
        context:{},
})

    const sortedChar: {[key:string]: JSX.Element[]} = {
        "physical":[],
        "mental": [],
        "social": [],
    }

    function genButtons(name: string): JSX.Element[]{
        let buttons = [] 
        for (let i= 1; i <= 5; i++) {
            if (i == 1) {
                buttons.push(
                    <input type="radio" name={name} value={i} id={name} checked/>
                )
            }
            buttons.push(
                <input type="radio" name={name} value={i} id={name}/>
            )
        }
        return buttons
    }

    // TODO: Model after the character sheet... or a prettier version. name and buttons on top and an onclick revealing any information needed by the user.
    // Not sure if request should broken up...or just overfetch and show if necessary? hm.
    data.charByType.forEach((attr: Characteristic, idx: number) => {
        var element: JSX.Element
        var btns: JSX.Element[]
        switch(attr.name) {
            case "charisma":
            case "manipulation":
            case "appearance":
                btns = genButtons(attr.name)
                element = <div className="social" key={attr.name}>
                    <h1>{attr.name}</h1>
                    <p>{attr.description}</p>
                    <label htmlFor={attr.name}>  {attr.name} value </label>
                    { btns }
                    </div>
                sortedChar.social.push(element)
                break
            case "perception":
            case "wits":
            case "intelligence":
                btns = genButtons(attr.name)
                element = <div className="mental" key={attr.name}>
                    <h1>{attr.name}</h1>
                    <p>{attr.description}</p>
                    <label htmlFor={attr.name}>  {attr.name} value </label>
                    { btns }          
                </div>
                
                sortedChar.mental.push(element)
                break
            default:
                btns = genButtons(attr.name)
                element = <div className="physical" key={attr.name}>
                    <h1>{attr.name}</h1>
                    <p>{attr.description}</p>
                    <label htmlFor={attr.name}>  {attr.name} value </label>
                    { btns }
                </div>
                sortedChar.physical.push(element)
                break
        }
    });
    return (
        <div>
            {sortedChar.physical}
            {sortedChar.mental}
            {sortedChar.social}
        </div>
    )
}

export default AttrProps