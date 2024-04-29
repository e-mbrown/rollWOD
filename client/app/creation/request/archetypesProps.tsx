'use server'

import { getClient } from "@/app/apolloclient"
import { gql } from "@apollo/client"

const GET_ARCHETYPES = gql`
query getArchetypes {
   archetypes {
        name
        description
        sys
   }
}
`

const ArchProp = async () => {
    const client = getClient()
    const { data } = await client.query({
        query: GET_ARCHETYPES,
    })

    return (
        <div>
            {data.archetypes.map((arch: Archetype) => (
                <div key="arch">
                    <h2>{arch.name}</h2>
                    <p>{arch.description}</p>
                    <p>{arch.sys}</p>
                </div>
            ))}
        </div>
    )
}

export default ArchProp