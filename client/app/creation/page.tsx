import Link from "next/link"

export default function Creation() {
    return(
        <div>
            <button><Link href="creation/general">Create Character</Link></button>
            <button> <Link href="home">Generate Random Character</Link></button>
            <button><Link href="creation/attribute">Temp Button </Link></button>
        </div>
    )
}