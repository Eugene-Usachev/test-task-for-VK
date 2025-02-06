export interface Ping {
	containerId: number;
	pingTimeMS: number;
	wasSuccessful: boolean;
	date: Date;
}