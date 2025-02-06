import {useCallback, useEffect, useState} from "react";
import {getContainers} from "@/services/container";
import {Table} from "antd";
import styles from './containerTable.module.scss';
import {Container, ParseContainerFromBackend} from "@/models/container";

interface RowData {
	id: number;
	ipAddress: string;
	pingTime: string | "None";
	lastSuccessDate: string | "Never";
}

function RowDataFromContainer(container: Container): RowData {
	if (container.lastPing) {
		const secondsInMs = container.lastPing!.date.seconds * 1000;
		const nanosInMs = container.lastPing!.date.nanos * 1e-6;
		const date = new Date(secondsInMs + nanosInMs);
		const dateHours = date.getHours() < 10 ? "0" + date.getHours() : date.getHours();
		const dateMinutes = date.getMinutes() < 10 ? "0" + date.getMinutes() : date.getMinutes();
		const dateSeconds = date.getSeconds() < 10 ? "0" + date.getSeconds() : date.getSeconds();
		const dateString = `${dateHours}:${dateMinutes}:${dateSeconds} ${date.getDate()}/${date.getMonth() + 1}/${date.getFullYear()}`

		const micros = container.lastPing!.pingTimeInMicroseconds;
		const time = micros > 1e5 ? (micros / 1e3) + "ms" : micros + "μs";

		return {
			id: container.id,
			ipAddress: container.ipAddress,
			pingTime: time,
			lastSuccessDate: dateString,
		}
	} else {
		return {
			id: container.id,
			ipAddress: container.ipAddress,
			pingTime: "None",
			lastSuccessDate: "Never",
		}
	}
}

async function getRows(): Promise<RowData[] | null> {
	let res = await getContainers();

	if (res.status === 200) {
		const { successfulContainers, invalidContainers } = res.data;
		let containers: Container[] = [];

		for (const container of successfulContainers) {
			containers.push(ParseContainerFromBackend(container));
		}

		for (const container of invalidContainers) {
			containers.push(ParseContainerFromBackend(container));
		}

		containers.sort((f, s) => {
			if (f.lastPing && s.lastPing) {
				return 0
			}

			if (f.lastPing) {
				return -1;
			}

			if (s.lastPing) {
				return 1;
			}

			return 0;
		});

		let rows: RowData[] = [];

		for (const container of containers) {
			rows.push(RowDataFromContainer(container));
		}

		return rows;
	} else {
		const msg = "Failed to get containers, reason: " + res.data;

		alert(msg);
		console.log(msg);

		return null;
	}
}

export const ContainerTable = () => {
	const [containers, setContainers] = useState<RowData[]>([]);
	const getAndSetRows = useCallback(() => {
		getRows().then((rows) => {
			if (rows) {
				setContainers(rows);
			}
		})
	}, []);

	useEffect(() => {
		const interval = setInterval(() => {
			getAndSetRows();
		}, 5000);

		getAndSetRows();

		return () => clearInterval(interval);
	}, []);

	return (
		<div>
			<Table className={styles.containerTable} dataSource={containers} rowKey="id" pagination={false}>
				<Table.Column title="IP Address" dataIndex="ipAddress" />
				<Table.Column title="Ping Time" dataIndex="pingTime" />
				<Table.Column title="Last Successful Attempt" dataIndex="lastSuccessDate" />
			</Table>
		</div>
	);
};