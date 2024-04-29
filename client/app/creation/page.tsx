import Concept from "./concept"
import ArchProp from "./request/archetypesProps"
import ClanProps from "./request/clanprops"

// TODO: Move everything to seperate pages

export default function Creator(){
    const char: Partial<MasqueradeChar> = {}


    return(
         <div>
            <h1> Let's get this for character creation</h1>
            <Concept char={char} clan={ClanProps()} arch={ArchProp()} ></Concept>
        </div>
)}