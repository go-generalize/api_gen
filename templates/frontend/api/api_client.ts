// THIS CODE WAS GENERATED AUTOMATICALLY
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS

import { PostCreateUserRequest as PostCreateUserRequest } from './classes//PostCreateUserRequest';
export { PostCreateUserRequest as PostCreateUserRequest } from './classes//PostCreateUserRequest';
import { PostCreateUserResponse as PostCreateUserResponse } from './classes//PostCreateUserResponse';
export { PostCreateUserResponse as PostCreateUserResponse } from './classes//PostCreateUserResponse';




export class APIClient {
	private headers: {[key: string]: string};
	private options: {[key: string]: any};
	private baseURL: string;


	constructor(
		token?: string,
		commonHeaders?: {[key: string]: string},
		baseURL?: string,
		commonOptions: {[key: string]: any} = {}
	) {
		const headers: {[key: string]: string} =  {
			'Content-Type': 'application/json',
			...commonHeaders,
		};

		if (token !== undefined) {
			headers['Authorization'] = 'Bearer ' + token;
		}

		this.baseURL =  (baseURL === undefined) ? "" : baseURL;
		this.options = commonOptions;
		this.headers = headers;


	}

	async postCreateUser(
		param: PostCreateUserRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<PostCreateUserResponse> {
		const resp = await fetch(
			this.baseURL + "/create_user",
			{
				method: "POST",
				body: JSON.stringify(param),
				headers: {
					...this.headers,
					...headers,
				},
				...this.options,
				...options,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			throw new Error(resp.statusText + ": " + await resp.text());
		}

		return new PostCreateUserResponse(await resp.json());
	}
}
