import Link from "next/link"

export default function Creation() {
    return(
        <div>
            <button><Link href="creation/general/false">Create Character</Link></button>
            <button> <Link href="home">Generate Random Character</Link></button>
            <button><Link href="creation/user/attribute">Temp Button </Link></button>
        </div>
    )
}