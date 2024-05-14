'use server'
import { gql } from "@apollo/client"
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
           {data.clans.map((clan: Clan, idx: number) => (
            <div key={"clan" + idx}>
                <h2>{clan.name}</h2>
                <p>{clan.description}</p> 
                <input type="radio" name={"clan"} value={clan.name}/> 
            </div>
           ))}
        </div>
    )
}

export default ClanProps