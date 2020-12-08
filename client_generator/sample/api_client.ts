// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: devel

import {
	GetStaticPageRequest as ServiceStaticPageGetStaticPageRequest,
	GetStaticPageResponse as ServiceStaticPageGetStaticPageResponse,
} from './classes/service/static_page/types';

import {
	GetArticleRequest as ServiceGetArticleRequest,
	GetArticleResponse as ServiceGetArticleResponse,
} from './classes/service/types';

import {
	GetRequest as ServiceUserGetRequest,
	GetResponse as ServiceUserGetResponse,
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
	GetRequest as GetRequest,
	GetResponse as GetResponse,
	PostCreateTableRequest as PostCreateTableRequest,
	PostCreateTableResponse as PostCreateTableResponse,
	PostCreateUserRequest as PostCreateUserRequest,
	PostCreateUserResponse as PostCreateUserResponse,
} from './classes/types';

export type ApiClientBeforeMiddleware =
	(httpMethod: string, endpoint: string, request: unknown) => Promise<boolean>;
export type ApiClientAfterMiddleware =
	(httpMethod: string, endpoint: string, request: unknown, response: unknown) => Promise<boolean>;

export interface middlewareSet {
	beforeMiddleware?: ApiClientBeforeMiddleware[];
	afterMiddleware?: ApiClientAfterMiddleware[];
}


class ServiceClient {
	private beforeMiddleware: ApiClientBeforeMiddleware[] = [];
	private afterMiddleware: ApiClientAfterMiddleware[] = [];
	public static_page: ServiceStaticPageClient;
	public user: ServiceUserClient;
	public user2: ServiceUser2Client;
	constructor(
		private headers: {[key: string]: string},
		private options: {[key: string]: any},
		private baseURL: string,
		middleware: middlewareSet
	) {
		this.beforeMiddleware = middleware.beforeMiddleware!;
		this.afterMiddleware = middleware.afterMiddleware!;
		const childMiddlewareSet: middlewareSet = {
			beforeMiddleware: this.beforeMiddleware,
			afterMiddleware: this.afterMiddleware
		};
		this.static_page = new ServiceStaticPageClient(
			headers,
			options,
			baseURL,
			childMiddlewareSet
		);
		this.user = new ServiceUserClient(
			headers,
			options,
			baseURL,
			childMiddlewareSet
		);
		this.user2 = new ServiceUser2Client(
			headers,
			options,
			baseURL,
			childMiddlewareSet
		);
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
		for (const m of this.beforeMiddleware) {
			await m('GET', `${this.baseURL}/service/article`, param);
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
		const res = (await resp.json()) as ServiceGetArticleResponse;
		for (const m of this.afterMiddleware) {
			await m('GET', `${this.baseURL}/service/article`, param, res);
		}
		return res;
	}
}

class ServiceStaticPageClient {
	private beforeMiddleware: ApiClientBeforeMiddleware[] = [];
	private afterMiddleware: ApiClientAfterMiddleware[] = [];
	constructor(
		private headers: {[key: string]: string},
		private options: {[key: string]: any},
		private baseURL: string,
		middleware: middlewareSet
	) {
		this.beforeMiddleware = middleware.beforeMiddleware!;
		this.afterMiddleware = middleware.afterMiddleware!;
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
		for (const m of this.beforeMiddleware) {
			await m('GET', `${this.baseURL}/service/static_page/static_page`, param);
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
		const res = (await resp.json()) as ServiceStaticPageGetStaticPageResponse;
		for (const m of this.afterMiddleware) {
			await m('GET', `${this.baseURL}/service/static_page/static_page`, param, res);
		}
		return res;
	}
}

class ServiceUser2Client {
	private beforeMiddleware: ApiClientBeforeMiddleware[] = [];
	private afterMiddleware: ApiClientAfterMiddleware[] = [];
	public _userID: ServiceUser2UserIDClient;
	constructor(
		private headers: {[key: string]: string},
		private options: {[key: string]: any},
		private baseURL: string,
		middleware: middlewareSet
	) {
		this.beforeMiddleware = middleware.beforeMiddleware!;
		this.afterMiddleware = middleware.afterMiddleware!;
		const childMiddlewareSet: middlewareSet = {
			beforeMiddleware: this.beforeMiddleware,
			afterMiddleware: this.afterMiddleware
		};
		this._userID = new ServiceUser2UserIDClient(
			headers,
			options,
			baseURL,
			childMiddlewareSet
		);
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
		for (const m of this.beforeMiddleware) {
			await m('GET', `${this.baseURL}/service/user2/${encodeURI(param.id.toString())}`, param);
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
		const res = (await resp.json()) as ServiceUser2GetUserResponse;
		for (const m of this.afterMiddleware) {
			await m('GET', `${this.baseURL}/service/user2/${encodeURI(param.id.toString())}`, param, res);
		}
		return res;
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
		for (const m of this.beforeMiddleware) {
			await m('POST', `${this.baseURL}/service/user2/update_user_name`, param);
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
		const res = (await resp.json()) as ServiceUser2PostUpdateUserNameResponse;
		for (const m of this.afterMiddleware) {
			await m('POST', `${this.baseURL}/service/user2/update_user_name`, param, res);
		}
		return res;
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
		for (const m of this.beforeMiddleware) {
			await m('POST', `${this.baseURL}/service/user2/update_user_password`, param);
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
		const res = (await resp.json()) as ServiceUser2PostUpdateUserPasswordResponse;
		for (const m of this.afterMiddleware) {
			await m('POST', `${this.baseURL}/service/user2/update_user_password`, param, res);
		}
		return res;
	}
}

class ServiceUser2UserIDClient {
	private beforeMiddleware: ApiClientBeforeMiddleware[] = [];
	private afterMiddleware: ApiClientAfterMiddleware[] = [];
	public _JobID: ServiceUser2UserIDJobIDClient;
	constructor(
		private headers: {[key: string]: string},
		private options: {[key: string]: any},
		private baseURL: string,
		middleware: middlewareSet
	) {
		this.beforeMiddleware = middleware.beforeMiddleware!;
		this.afterMiddleware = middleware.afterMiddleware!;
		const childMiddlewareSet: middlewareSet = {
			beforeMiddleware: this.beforeMiddleware,
			afterMiddleware: this.afterMiddleware
		};
		this._JobID = new ServiceUser2UserIDJobIDClient(
			headers,
			options,
			baseURL,
			childMiddlewareSet
		);
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
		for (const m of this.beforeMiddleware) {
			await m('GET', `${this.baseURL}/service/user2/${encodeURI(param.UserID.toString())}/user_job_get`, param);
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
		const res = (await resp.json()) as ServiceUser2UserIDGetUserJobGetResponse;
		for (const m of this.afterMiddleware) {
			await m('GET', `${this.baseURL}/service/user2/${encodeURI(param.UserID.toString())}/user_job_get`, param, res);
		}
		return res;
	}
}

class ServiceUser2UserIDJobIDClient {
	private beforeMiddleware: ApiClientBeforeMiddleware[] = [];
	private afterMiddleware: ApiClientAfterMiddleware[] = [];
	constructor(
		private headers: {[key: string]: string},
		private options: {[key: string]: any},
		private baseURL: string,
		middleware: middlewareSet
	) {
		this.beforeMiddleware = middleware.beforeMiddleware!;
		this.afterMiddleware = middleware.afterMiddleware!;
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
		for (const m of this.beforeMiddleware) {
			await m('PUT', `${this.baseURL}/service/user2/${encodeURI(param.UserID.toString())}/${encodeURI(param.JobID.toString())}/job`, param);
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
		const res = (await resp.json()) as ServiceUser2UserIDJobIDPutJobResponse;
		for (const m of this.afterMiddleware) {
			await m('PUT', `${this.baseURL}/service/user2/${encodeURI(param.UserID.toString())}/${encodeURI(param.JobID.toString())}/job`, param, res);
		}
		return res;
	}
}

class ServiceUserClient {
	private beforeMiddleware: ApiClientBeforeMiddleware[] = [];
	private afterMiddleware: ApiClientAfterMiddleware[] = [];
	constructor(
		private headers: {[key: string]: string},
		private options: {[key: string]: any},
		private baseURL: string,
		middleware: middlewareSet
	) {
		this.beforeMiddleware = middleware.beforeMiddleware!;
		this.afterMiddleware = middleware.afterMiddleware!;
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

	async get(
		param: ServiceUserGetRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceUserGetResponse> {
	    const excludeParams: string[] = [];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}
		for (const m of this.beforeMiddleware) {
			await m('GET', `${this.baseURL}/service/user`, param);
		}
		const url = `${this.baseURL}/service/user?` + (new URLSearchParams(this.getRequestObject(param, excludeParams))).toString();
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
		await resp.text();
		const res = {} as ServiceUserGetResponse;
		for (const m of this.afterMiddleware) {
			await m('GET', `${this.baseURL}/service/user`, param, res);
		}
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
		for (const m of this.beforeMiddleware) {
			await m('POST', `${this.baseURL}/service/user/update_user_name`, param);
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
		const res = (await resp.json()) as ServiceUserPostUpdateUserNameResponse;
		for (const m of this.afterMiddleware) {
			await m('POST', `${this.baseURL}/service/user/update_user_name`, param, res);
		}
		return res;
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
		for (const m of this.beforeMiddleware) {
			await m('POST', `${this.baseURL}/service/user/update_user_password`, param);
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
		const res = (await resp.json()) as ServiceUserPostUpdateUserPasswordResponse;
		for (const m of this.afterMiddleware) {
			await m('POST', `${this.baseURL}/service/user/update_user_password`, param, res);
		}
		return res;
	}
}

export class APIClient {
	private headers: {[key: string]: string};
	private options: {[key: string]: any};
	private baseURL: string;

	private beforeMiddleware: ApiClientBeforeMiddleware[] = [];
	private afterMiddleware: ApiClientAfterMiddleware[] = [];

	public service: ServiceClient;

	constructor(
		token?: string,
		commonHeaders?: {[key: string]: string},
		baseURL?: string,
		commonOptions: {[key: string]: any} = {},
		middleware?: middlewareSet
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

		if (middleware) {
			this.beforeMiddleware = middleware.beforeMiddleware ?? [];
			this.afterMiddleware = middleware.afterMiddleware ?? [];
		}
		const childMiddlewareSet: middlewareSet = {
			beforeMiddleware: this.beforeMiddleware,
			afterMiddleware: this.afterMiddleware
		};

		this.service = new ServiceClient(
			headers,
			this.options,
			this.baseURL,
			childMiddlewareSet
		);
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

	async get(
		param: GetRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<GetResponse> {
	    const excludeParams: string[] = [];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}
		for (const m of this.beforeMiddleware) {
			await m('GET', `${this.baseURL}/`, param);
		}
		const url = `${this.baseURL}/?` + (new URLSearchParams(this.getRequestObject(param, excludeParams))).toString();
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
		const res = (await resp.json()) as GetResponse;
		for (const m of this.afterMiddleware) {
			await m('GET', `${this.baseURL}/`, param, res);
		}
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
		for (const m of this.beforeMiddleware) {
			await m('POST', `${this.baseURL}/create_table`, param);
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
		const res = (await resp.json()) as PostCreateTableResponse;
		for (const m of this.afterMiddleware) {
			await m('POST', `${this.baseURL}/create_table`, param, res);
		}
		return res;
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
		for (const m of this.beforeMiddleware) {
			await m('POST', `${this.baseURL}/create_user`, param);
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
		const res = (await resp.json()) as PostCreateUserResponse;
		for (const m of this.afterMiddleware) {
			await m('POST', `${this.baseURL}/create_user`, param, res);
		}
		return res;
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
