// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: 0.4.0

import {
	PostCreateUserRequest as PostCreateUserRequest,
	PostCreateUserResponse as PostCreateUserResponse,

} from './classes//types';
export {
	PostCreateUserRequest as PostCreateUserRequest,
	PostCreateUserResponse as PostCreateUserResponse,

} from './classes//types';




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

	getRequestObject(param: any, routingPath: string[]): { [key: string]: any } {
		const obj = param.toObject();
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
		const resp = await fetch(
			`${this.baseURL}/${encodeURI(param.ID.toString())}`,
			{
				method: "POST",
				body: JSON.stringify(this.getRequestObject(param, excludeParams)),
				headers: {
					...this.headers,
					...headers,
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

    set response(value: string) {
        this._response = value;
    }

    set statusText(value: string) {
        this._statusText = value;
    }

    get statusCode(): number {
        return this._statusCode;
    }
}
