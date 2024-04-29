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

const attrProps = async () => {
    const client = getClient()
    const { data } = await client.query({
        query:GET_ATTRS,
        context:{},
})
    return (
        <div>
            {data.attributes.map((attr: Characteristic, idx: number) =>(
                <div>
                    <h1 key={idx}>{attr.name}</h1>
                    <p>{attr.description}</p>
                </div>
                
            ))}
        </div>
    )
}