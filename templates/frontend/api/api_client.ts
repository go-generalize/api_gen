// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: unknown

import {
	PostCreateUserRequest as PostCreateUserRequest,
	PostCreateUserResponse as PostCreateUserResponse,
} from './classes/types';


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

	getRequestObject(obj: any, routingPath: string[]): { [key: string]: any } {
		let res: { [key: string]: any } = {};
		Object.keys(obj).forEach((key) => {
			if (routingPath.indexOf(key) === -1) {
				res[key] = obj[key];
			}
		});
		return res;
	}

	async postCreateUser(
		param: PostCreateUserRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<PostCreateUserResponse> {
	    const excludeParams: string[] = ['ID', ];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}
		const url = `${this.baseURL}/${encodeURI(param.ID.toString())}`;

		const resp = await fetch(
			url,
			{
				method: "POST",
				body: JSON.stringify(this.getRequestObject(param, excludeParams)),
				headers: {
					...this.headers,
					...headers,
					...mockHeaders,
				},
				...this.options,
				...options,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			const responseText = await resp.text();
			throw new ApiError(resp, responseText);
		}
		return (await resp.json()) as PostCreateUserResponse;
	}
}

export class ApiError extends Error {
    private _statusCode: number;
    private _statusText: string;
    private _response: string;

    constructor(response: Response, responseText: string) {
        super();
        this._statusCode = response.status;
        this._statusText = response.statusText;
        this._response = responseText
    }

    get statusCode(): number {
        return this._statusCode;
    }

	get statusText(): string {
        return this._statusText;
    }

	get response(): string {
        return this._response;
    }
}

export interface MockOption {
	wait_ms: number;
	target_file: string;
}
