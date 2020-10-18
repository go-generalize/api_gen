// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: unknown

import {
	GetArticleRequest as ServiceGetArticleRequest,
	GetArticleResponse as ServiceGetArticleResponse,
} from './classes/service/types';

import {
	PostUpdateUserNameRequest as ServiceUserPostUpdateUserNameRequest,
	PostUpdateUserNameResponse as ServiceUserPostUpdateUserNameResponse,
	PostUpdateUserPasswordRequest as ServiceUserPostUpdateUserPasswordRequest,
	PostUpdateUserPasswordResponse as ServiceUserPostUpdateUserPasswordResponse,
} from './classes/service/user/types';

import {
	PutJobRequest as ServiceUser2UserIDJobIDPutJobRequest,
	PutJobResponse as ServiceUser2UserIDJobIDPutJobResponse,
} from './classes/service/user2/_userID/_JobID/types';

import {
	GetUserJobGetRequest as ServiceUser2UserIDGetUserJobGetRequest,
	GetUserJobGetResponse as ServiceUser2UserIDGetUserJobGetResponse,
} from './classes/service/user2/_userID/types';

import {
	GetUserRequest as ServiceUser2GetUserRequest,
	GetUserResponse as ServiceUser2GetUserResponse,
	PostUpdateUserNameRequest as ServiceUser2PostUpdateUserNameRequest,
	PostUpdateUserNameResponse as ServiceUser2PostUpdateUserNameResponse,
	PostUpdateUserPasswordRequest as ServiceUser2PostUpdateUserPasswordRequest,
	PostUpdateUserPasswordResponse as ServiceUser2PostUpdateUserPasswordResponse,
} from './classes/service/user2/types';

import {
	PostCreateTableRequest as PostCreateTableRequest,
	PostCreateTableResponse as PostCreateTableResponse,
	PostCreateUserRequest as PostCreateUserRequest,
	PostCreateUserResponse as PostCreateUserResponse,
} from './classes/types';


class ServiceClient {
	public static_page: ServiceStaticPageClient;
	public user: ServiceUserClient;
	public user2: ServiceUser2Client;
	constructor(private headers: {[key: string]: string}, private options: {[key: string]: any}, private baseURL: string) {
		this.static_page = new ServiceStaticPageClient(headers, options, baseURL);
		this.user = new ServiceUserClient(headers, options, baseURL);
		this.user2 = new ServiceUser2Client(headers, options, baseURL);
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

	async getArticle(
		param: ServiceGetArticleRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceGetArticleResponse> {
	    const excludeParams: string[] = [];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}
		const url = `${this.baseURL}/service/article?` + (new URLSearchParams(this.getRequestObject(param, excludeParams))).toString();
		const resp = await fetch(
			url,
			{
				method: "GET",
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
		return (await resp.json()) as ServiceGetArticleResponse;
	}
}

class ServiceStaticPageClient {
	constructor(private headers: {[key: string]: string}, private options: {[key: string]: any}, private baseURL: string) {
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

	async getStaticPage(
		param: ServiceStaticPageGetStaticPageRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceStaticPageGetStaticPageResponse> {
	    const excludeParams: string[] = [];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}
		const url = `${this.baseURL}/service/static_page/static_page?` + (new URLSearchParams(this.getRequestObject(param, excludeParams))).toString();
		const resp = await fetch(
			url,
			{
				method: "GET",
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
		return (await resp.json()) as ServiceStaticPageGetStaticPageResponse;
	}
}

class ServiceUser2Client {
	public _userID: ServiceUser2UserIDClient;
	constructor(private headers: {[key: string]: string}, private options: {[key: string]: any}, private baseURL: string) {
		this._userID = new ServiceUser2UserIDClient(headers, options, baseURL);
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

	async getUser(
		param: ServiceUser2GetUserRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceUser2GetUserResponse> {
	    const excludeParams: string[] = ['id', ];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}
		const url = `${this.baseURL}/service/user2/${encodeURI(param.id.toString())}?` + (new URLSearchParams(this.getRequestObject(param, excludeParams))).toString();
		const resp = await fetch(
			url,
			{
				method: "GET",
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
		return (await resp.json()) as ServiceUser2GetUserResponse;
	}

	async postUpdateUserName(
		param: ServiceUser2PostUpdateUserNameRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceUser2PostUpdateUserNameResponse> {
	    const excludeParams: string[] = [];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}
		const url = `${this.baseURL}/service/user2/update_user_name`;
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
		return (await resp.json()) as ServiceUser2PostUpdateUserNameResponse;
	}

	async postUpdateUserPassword(
		param: ServiceUser2PostUpdateUserPasswordRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceUser2PostUpdateUserPasswordResponse> {
	    const excludeParams: string[] = [];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}
		const url = `${this.baseURL}/service/user2/update_user_password`;
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
		return (await resp.json()) as ServiceUser2PostUpdateUserPasswordResponse;
	}
}

class ServiceUser2UserIDClient {
	public _JobID: ServiceUser2UserIDJobIDClient;
	constructor(private headers: {[key: string]: string}, private options: {[key: string]: any}, private baseURL: string) {
		this._JobID = new ServiceUser2UserIDJobIDClient(headers, options, baseURL);
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

	async getUserJobGet(
		param: ServiceUser2UserIDGetUserJobGetRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceUser2UserIDGetUserJobGetResponse> {
	    const excludeParams: string[] = ['UserID', ];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}
		const url = `${this.baseURL}/service/user2/${encodeURI(param.UserID.toString())}/user_job_get?` + (new URLSearchParams(this.getRequestObject(param, excludeParams))).toString();
		const resp = await fetch(
			url,
			{
				method: "GET",
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
		return (await resp.json()) as ServiceUser2UserIDGetUserJobGetResponse;
	}
}

class ServiceUser2UserIDJobIDClient {
	constructor(private headers: {[key: string]: string}, private options: {[key: string]: any}, private baseURL: string) {
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

	async putJob(
		param: ServiceUser2UserIDJobIDPutJobRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceUser2UserIDJobIDPutJobResponse> {
	    const excludeParams: string[] = ['UserID', 'JobID', ];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}
		const url = `${this.baseURL}/service/user2/${encodeURI(param.UserID.toString())}/${encodeURI(param.JobID.toString())}/job`;
		const resp = await fetch(
			url,
			{
				method: "PUT",
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
		return (await resp.json()) as ServiceUser2UserIDJobIDPutJobResponse;
	}
}

class ServiceUserClient {
	constructor(private headers: {[key: string]: string}, private options: {[key: string]: any}, private baseURL: string) {
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

	async postUpdateUserName(
		param: ServiceUserPostUpdateUserNameRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceUserPostUpdateUserNameResponse> {
	    const excludeParams: string[] = [];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}
		const url = `${this.baseURL}/service/user/update_user_name`;
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
		return (await resp.json()) as ServiceUserPostUpdateUserNameResponse;
	}

	async postUpdateUserPassword(
		param: ServiceUserPostUpdateUserPasswordRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceUserPostUpdateUserPasswordResponse> {
	    const excludeParams: string[] = [];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}
		const url = `${this.baseURL}/service/user/update_user_password`;
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
		return (await resp.json()) as ServiceUserPostUpdateUserPasswordResponse;
	}
}

export class APIClient {
	private headers: {[key: string]: string};
	private options: {[key: string]: any};
	private baseURL: string;

	public service: ServiceClient;

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

		this.service = new ServiceClient(headers, this.options, this.baseURL);
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

	async postCreateTable(
		param: PostCreateTableRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<PostCreateTableResponse> {
	    const excludeParams: string[] = [];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}
		const url = `${this.baseURL}/create_table`;

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
		return (await resp.json()) as PostCreateTableResponse;
	}

	async postCreateUser(
		param: PostCreateUserRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<PostCreateUserResponse> {
	    const excludeParams: string[] = [];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}
		const url = `${this.baseURL}/create_user`;

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
	wait_ms: string;
	target_file: string;
}
