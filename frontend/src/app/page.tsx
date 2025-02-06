'use client';

import {ContainerTable} from "@/app/_components/containerTable/containerTable";
import styles from './page.module.scss';

export default function HomePage() {
  return (
    <div className={styles.main}>
      <div className={styles.header}>
          <h1>Test task for VK</h1>
          <h2>List of containers</h2>
      </div>
      <ContainerTable/>
    </div>
  );
}