import Image from "next/image";
import styles from "./page.module.css";
import Link from "next/link";
import ClanProps from "./creation/request/clanprops";


export default function Home() {
  return (
    <div>
      <h1> Our Homepage. Making progress! Woo</h1>
      <Link href="creation">Creation Page</Link>
      

    </div>
  
  )
}
