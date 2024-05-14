type Clan = {
    name: string,
    description: string,
    appearance: string,
    haven: string,
    background: string
    character: string,
    weakness: string,
    organization: string
    strongholds: [string]
}

enum CharType {
    attribute,
    talent,
    discipline,
    skill,
    background
}

type Characteristic = {
    name: string,
    type: CharType,
    description: string,
    descByVal: [string]
}

type Archetype = {
    name: string,
    description: string,
    sys: string
}

class Character {
    constructor(
        private id: string,
        public name: string,
        protected user: string,
        public nature: string,
        public demeanor: string,
        public concept:string
    ){}
}

class MasqueradeChar extends Character {
    constructor( 
        id: string,
        name: string,
        user: string, 
        nature: string,
        demeanor: string,
        concept:string,
        public clan: string,
        public generation: string,
        private Sire: string,
        public attributes: [[string, number]],
        public abilities: [[string, number]], 
        public advantages: [[string, number]],
        ){

        super(id,name,user,nature,demeanor, concept)
    }
}

