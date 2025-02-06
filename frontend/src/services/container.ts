import axios, {AxiosResponse} from 'axios';
import {InvalidContainerFromBackend, SuccessContainerFromBackend} from "@/models/container";

const API_URL = 'http://localhost:4040';

interface GetContainerWithLatestPingResponse {
	/** These containers have latest successful ping */
	successful_containers: SuccessContainerFromBackend[];
	/** These containers __don't__ have latest successful ping */
	invalid_containers: InvalidContainerFromBackend[];
}

export const getContainers = async (): Promise<AxiosResponse<GetContainerWithLatestPingResponse>> => {
	return axios.get(`${API_URL}/container/with_latest_ping`);
}

export const registerContainer = async (ipAddress: string): Promise<AxiosResponse> => {
	return axios.post(`${API_URL}/container`, { ip_address: ipAddress });
}

export const unregisterContainer = async (containerID: number): Promise<AxiosResponse> => {
	return axios.delete(`${API_URL}/container/${containerID}`);
}