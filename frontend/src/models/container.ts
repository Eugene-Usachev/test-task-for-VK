export interface LastSuccessPing {
	pingTimeInMicroseconds: number;
	date: {
		seconds: number;
		nanos: number;
	};
}

export interface SuccessContainerFromBackend {
	id: number;
	ip_address: string;
	ping_time: number;
	date: {
		seconds: number;
		nanos: number;
	};
}

export interface InvalidContainerFromBackend {
	id: number;
	ip_address: string;
}

export function ParseContainerFromBackend(
	container: SuccessContainerFromBackend | InvalidContainerFromBackend
): Container {
	if ("ping_time" in container) {
		return {
			id: container.id,
			ipAddress: container.ip_address,
			lastPing: {
				pingTimeInMicroseconds: container.ping_time,
				date: {
					seconds: container.date.seconds,
					nanos: container.date.nanos,
				}
			}
		}
	} else {
		return {
			id: container.id,
			ipAddress: container.ip_address,
		}
	}
}

export interface Container {
	id: number;
	ipAddress: string;
	lastPing?: LastSuccessPing;
}