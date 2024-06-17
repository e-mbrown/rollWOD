// Used in Attribute request
const AttrList: {[key:string]: Characteristic[]} = {
    "physical":[],
    "mental": [],
    "social": [],
}


function GenButtons(name: string): JSX.Element[]{
    let buttons = []
    for (let i= 1; i <= 5; i++) {
        if (i == 1) {
            buttons.push(
                <input type="radio" name={name} value={i} id={name} checked/>
            )
        } else {
            buttons.push(
                <input type="radio" name={name} value={i} id={name}/>
            )
        }
        
    }
    return buttons
}

///



export {GenButtons, AttrList}