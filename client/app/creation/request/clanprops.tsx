'use server'
import { gql, useQuery } from "@apollo/client"
import { getClient } from "@/app/apolloclient"

const GET_CLANS = gql`
query getClans {
    clans {
        name
        description
    }
}
`
const ClanProps = async () => {
    const client = getClient()
    const { data } =  await client.query({
        query: GET_CLANS,
        context:{}
})


    return (
        <div>
           {data.clans.map((clan: Clan) => (
            <div key="clan">
                <h2>{clan.name}</h2>
                <p>{clan.description}</p>    
            </div>
           ))}
        </div>
    )
}

export default ClanProps