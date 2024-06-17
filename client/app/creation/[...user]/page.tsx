import CreationForm from "./form"
import Concept from "./concept"
import Attribute from "../attributes/page"

export default function General({params}: {params: {user:string}}){
    return(
        <div>
            Our form is below
            <CreationForm
            slug={params.user}
            concept={<Concept/>}
            attr ={<Attribute user={params.user}/>}
            />
            
            
        </div>
    )
}