import React, { FormEvent } from "react"

 const Concept: React.FC<{setChar: Function, setState: Function, clan: React.ReactNode, arc: React.ReactNode}> = (props) => {
    function submitConcept(e: FormEvent<HTMLFormElement>) {
        e.preventDefault()
        const formData = new FormData(e.currentTarget)

        const obj:Partial<MasqueradeChar> = {
            clan: String(formData.get("clan")),
            demeanor: String(formData.get("demeanor")),
            concept: String(formData.get("concept")),
        }

        props.setChar(obj)
        props.setState("concept")
    }
    

    return (
        <div>
            <form onSubmit={submitConcept}>
                <div>
                    <h1>Concept</h1>
                    <p>
                    A character’s concept generally refers to who the character was
                    before becoming a vampire. Many Kindred cling desperately to any
                    salvageable aspects of their former selves — their self-image,
                    their occupation, how they lived, what was unique about them. In
                    their new nocturnal world, echoes of their mortal lives are all
                    that stand between many Kindred and madness.
                    </p> 
                    <figure>
                        <figcaption> Example concepts</figcaption>
                    </figure>
                    <ul>
                        <li>
                            Criminal — jailbird, Mafioso, drug dealer, pimp,
                            carjacker, thug, thief, fence
                        </li>
                        <li>
                            Drifter — bum, smuggler, prostitute, junkie, pilgrim,
                            biker, gambler
                        </li>
                        <li>
                            Entertainer — musician, film star, artist, club
                            kid,model, web celebrity
                        </li>
                        <li>
                            Intellectual — writer, student, scientist,philosopher,
                            social critic 
                        </li>
                        <li>
                            Investigator — detective, beat cop, governmentagent,
                            private eye, witch-hunter 
                        </li>
                        <li>
                            Politician — judge, public official, councilor,
                            aide,speechwriter
                        </li>
                        <li>
                            Professional — engineer, doctor, programmer,lawyer,
                            industrialist
                        </li>
                        <li>
                            Reporter — journalist, blogger, paparazzo, talkshow
                            host, media expert
                        </li>
                        <li>
                            Soldier — bodyguard, enforcer, soldier of fortune,
                            killer-for-hire</li>
                        <li>
                            Worker — trucker, farmer, wage earner, manservant,
                            construction laborer
                            </li>
                    </ul>
                    <input type="text" name="concept" size={60}></input>
                </div>
                <div>
                    <h1>Clan</h1>
                    <p> Clans determine the characters look, abilities and to some
                    extent their political affiliation.</p>
                    {props.clan}
                </div>
                <div>
                    <h1>Nature and Demeanor</h1>
                    <p> Your characters Nature and Demeanor generally describes your
                    characters behavioral traits. Their purpose is to assist in
                    understanding a characters personality. Demeanor is the way a
                    character presents herself to the outside world. It is the
                    “mask” she wears to protect her inner self. A character’s
                    Demeanor often differs from her Nature, though it might not.
                    Demeanor does not affect gameplay rules </p>
                    <p>Nature is the character’s “real” self, the person she truly
                    is. The Archetype a player chooses reflects that character’s
                    deep-rooted feelings about herself, others, and the world.
                    Nature should not be the only aspect of a character’s true
                    personality, merely the most dominant. Nature is also used to
                    determine a character’s ability to regain Willpower points</p>
                </div>
                {props.arc}
                <button onClick={(e) => submitConcept}>Submit</button>
            </form>
        </div> 
    )
}

export default Concept