// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: 2.0

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

export interface MiddlewareContext {
	httpMethod: string;
	endpoint: string;
	request: unknown;
	response?: unknown;
	baseURL: string;
	headers: {[key: string]: string};
	options: {[key: string]: any};
}

export enum MiddlewareResult {
	CONTINUE = 1,
	MIDDLEWARE_STOP = 2,
	STOP = 4
}

export type ApiClientMiddlewareFunc = (context: MiddlewareContext) => Promise<MiddlewareResult|boolean>;

export interface middlewareSet {
	beforeMiddleware?: ApiClientMiddlewareFunc[];
	afterMiddleware?: ApiClientMiddlewareFunc[];
}

class ServiceClient {
	private beforeMiddleware: ApiClientMiddlewareFunc[] = [];
	private afterMiddleware: ApiClientMiddlewareFunc[] = [];
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

	async callMiddleware(
		middlewares: ApiClientMiddlewareFunc[],
		context: MiddlewareContext
	) {
		for (const m of middlewares) {
			const func: ApiClientMiddlewareFunc = m;
			const mr = await func(context);
			if (typeof mr === 'boolean') {
				if (!mr) {
					break;
				}
			} else {
				if (mr === MiddlewareResult.CONTINUE) {
					continue;
				} else if (mr === MiddlewareResult.MIDDLEWARE_STOP) {
					break;
				} else if (mr === MiddlewareResult.STOP) {
					throw new ApiMiddlewareStop();
				}
			}
		}
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

		const reqHeader = {
			...this.headers,
			...headers,
			...mockHeaders,
		};
		const reqOption = {
			...this.options,
			...options,
		};
		const context: MiddlewareContext = {
			httpMethod: 'GET',
			endpoint: `${this.baseURL}/service/article`,
			request: param,
			baseURL: this.baseURL,
			headers: reqHeader,
			options: reqOption,
		};
		await this.callMiddleware(this.beforeMiddleware, context);
		const url = `${this.baseURL}/service/article?` + (new URLSearchParams(this.getRequestObject(param, excludeParams))).toString();
		const resp = await fetch(
			url,
			{
				method: "GET",
				headers: reqHeader,
				...reqOption,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			const responseText = await resp.text();
			throw new ApiError(resp, responseText);
		}
		const res = (await resp.json()) as ServiceGetArticleResponse;
		context.response = res;
		await this.callMiddleware(this.afterMiddleware, context);
		return res;
	}
}

class ServiceStaticPageClient {
	private beforeMiddleware: ApiClientMiddlewareFunc[] = [];
	private afterMiddleware: ApiClientMiddlewareFunc[] = [];
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

	async callMiddleware(
		middlewares: ApiClientMiddlewareFunc[],
		context: MiddlewareContext
	) {
		for (const m of middlewares) {
			const func: ApiClientMiddlewareFunc = m;
			const mr = await func(context);
			if (typeof mr === 'boolean') {
				if (!mr) {
					break;
				}
			} else {
				if (mr === MiddlewareResult.CONTINUE) {
					continue;
				} else if (mr === MiddlewareResult.MIDDLEWARE_STOP) {
					break;
				} else if (mr === MiddlewareResult.STOP) {
					throw new ApiMiddlewareStop();
				}
			}
		}
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

		const reqHeader = {
			...this.headers,
			...headers,
			...mockHeaders,
		};
		const reqOption = {
			...this.options,
			...options,
		};
		const context: MiddlewareContext = {
			httpMethod: 'GET',
			endpoint: `${this.baseURL}/service/static_page/static_page`,
			request: param,
			baseURL: this.baseURL,
			headers: reqHeader,
			options: reqOption,
		};
		await this.callMiddleware(this.beforeMiddleware, context);
		const url = `${this.baseURL}/service/static_page/static_page?` + (new URLSearchParams(this.getRequestObject(param, excludeParams))).toString();
		const resp = await fetch(
			url,
			{
				method: "GET",
				headers: reqHeader,
				...reqOption,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			const responseText = await resp.text();
			throw new ApiError(resp, responseText);
		}
		const res = (await resp.json()) as ServiceStaticPageGetStaticPageResponse;
		context.response = res;
		await this.callMiddleware(this.afterMiddleware, context);
		return res;
	}
}

class ServiceUser2Client {
	private beforeMiddleware: ApiClientMiddlewareFunc[] = [];
	private afterMiddleware: ApiClientMiddlewareFunc[] = [];
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

	async callMiddleware(
		middlewares: ApiClientMiddlewareFunc[],
		context: MiddlewareContext
	) {
		for (const m of middlewares) {
			const func: ApiClientMiddlewareFunc = m;
			const mr = await func(context);
			if (typeof mr === 'boolean') {
				if (!mr) {
					break;
				}
			} else {
				if (mr === MiddlewareResult.CONTINUE) {
					continue;
				} else if (mr === MiddlewareResult.MIDDLEWARE_STOP) {
					break;
				} else if (mr === MiddlewareResult.STOP) {
					throw new ApiMiddlewareStop();
				}
			}
		}
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

		const reqHeader = {
			...this.headers,
			...headers,
			...mockHeaders,
		};
		const reqOption = {
			...this.options,
			...options,
		};
		const context: MiddlewareContext = {
			httpMethod: 'GET',
			endpoint: `${this.baseURL}/service/user2/${encodeURI(param.id.toString())}`,
			request: param,
			baseURL: this.baseURL,
			headers: reqHeader,
			options: reqOption,
		};
		await this.callMiddleware(this.beforeMiddleware, context);
		const url = `${this.baseURL}/service/user2/${encodeURI(param.id.toString())}?` + (new URLSearchParams(this.getRequestObject(param, excludeParams))).toString();
		const resp = await fetch(
			url,
			{
				method: "GET",
				headers: reqHeader,
				...reqOption,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			const responseText = await resp.text();
			throw new ApiError(resp, responseText);
		}
		const res = (await resp.json()) as ServiceUser2GetUserResponse;
		context.response = res;
		await this.callMiddleware(this.afterMiddleware, context);
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

		const reqHeader = {
			...this.headers,
			...headers,
			...mockHeaders,
		};
		const reqOption = {
			...this.options,
			...options,
		};
		const context: MiddlewareContext = {
			httpMethod: 'POST',
			endpoint: `${this.baseURL}/service/user2/update_user_name`,
			request: param,
			baseURL: this.baseURL,
			headers: reqHeader,
			options: reqOption,
		};
		await this.callMiddleware(this.beforeMiddleware, context);
		const url = `${this.baseURL}/service/user2/update_user_name`;
		const resp = await fetch(
			url,
			{
				method: "POST",
				body: JSON.stringify(this.getRequestObject(param, excludeParams)),
				headers: reqHeader,
				...reqOption,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			const responseText = await resp.text();
			throw new ApiError(resp, responseText);
		}
		const res = (await resp.json()) as ServiceUser2PostUpdateUserNameResponse;
		context.response = res;
		await this.callMiddleware(this.afterMiddleware, context);
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

		const reqHeader = {
			...this.headers,
			...headers,
			...mockHeaders,
		};
		const reqOption = {
			...this.options,
			...options,
		};
		const context: MiddlewareContext = {
			httpMethod: 'POST',
			endpoint: `${this.baseURL}/service/user2/update_user_password`,
			request: param,
			baseURL: this.baseURL,
			headers: reqHeader,
			options: reqOption,
		};
		await this.callMiddleware(this.beforeMiddleware, context);
		const url = `${this.baseURL}/service/user2/update_user_password`;
		const resp = await fetch(
			url,
			{
				method: "POST",
				body: JSON.stringify(this.getRequestObject(param, excludeParams)),
				headers: reqHeader,
				...reqOption,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			const responseText = await resp.text();
			throw new ApiError(resp, responseText);
		}
		const res = (await resp.json()) as ServiceUser2PostUpdateUserPasswordResponse;
		context.response = res;
		await this.callMiddleware(this.afterMiddleware, context);
		return res;
	}
}

class ServiceUser2UserIDClient {
	private beforeMiddleware: ApiClientMiddlewareFunc[] = [];
	private afterMiddleware: ApiClientMiddlewareFunc[] = [];
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

	async callMiddleware(
		middlewares: ApiClientMiddlewareFunc[],
		context: MiddlewareContext
	) {
		for (const m of middlewares) {
			const func: ApiClientMiddlewareFunc = m;
			const mr = await func(context);
			if (typeof mr === 'boolean') {
				if (!mr) {
					break;
				}
			} else {
				if (mr === MiddlewareResult.CONTINUE) {
					continue;
				} else if (mr === MiddlewareResult.MIDDLEWARE_STOP) {
					break;
				} else if (mr === MiddlewareResult.STOP) {
					throw new ApiMiddlewareStop();
				}
			}
		}
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

		const reqHeader = {
			...this.headers,
			...headers,
			...mockHeaders,
		};
		const reqOption = {
			...this.options,
			...options,
		};
		const context: MiddlewareContext = {
			httpMethod: 'GET',
			endpoint: `${this.baseURL}/service/user2/${encodeURI(param.UserID.toString())}/user_job_get`,
			request: param,
			baseURL: this.baseURL,
			headers: reqHeader,
			options: reqOption,
		};
		await this.callMiddleware(this.beforeMiddleware, context);
		const url = `${this.baseURL}/service/user2/${encodeURI(param.UserID.toString())}/user_job_get?` + (new URLSearchParams(this.getRequestObject(param, excludeParams))).toString();
		const resp = await fetch(
			url,
			{
				method: "GET",
				headers: reqHeader,
				...reqOption,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			const responseText = await resp.text();
			throw new ApiError(resp, responseText);
		}
		const res = (await resp.json()) as ServiceUser2UserIDGetUserJobGetResponse;
		context.response = res;
		await this.callMiddleware(this.afterMiddleware, context);
		return res;
	}
}

class ServiceUser2UserIDJobIDClient {
	private beforeMiddleware: ApiClientMiddlewareFunc[] = [];
	private afterMiddleware: ApiClientMiddlewareFunc[] = [];
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

	async callMiddleware(
		middlewares: ApiClientMiddlewareFunc[],
		context: MiddlewareContext
	) {
		for (const m of middlewares) {
			const func: ApiClientMiddlewareFunc = m;
			const mr = await func(context);
			if (typeof mr === 'boolean') {
				if (!mr) {
					break;
				}
			} else {
				if (mr === MiddlewareResult.CONTINUE) {
					continue;
				} else if (mr === MiddlewareResult.MIDDLEWARE_STOP) {
					break;
				} else if (mr === MiddlewareResult.STOP) {
					throw new ApiMiddlewareStop();
				}
			}
		}
	}

	async putJob(
		param: ServiceUser2UserIDJobIDPutJobRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceUser2UserIDJobIDPutJobResponse> {
	    const excludeParams: string[] = ['JobID', 'UserID', ];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}

		const reqHeader = {
			...this.headers,
			...headers,
			...mockHeaders,
		};
		const reqOption = {
			...this.options,
			...options,
		};
		const context: MiddlewareContext = {
			httpMethod: 'PUT',
			endpoint: `${this.baseURL}/service/user2/${encodeURI(param.UserID.toString())}/${encodeURI(param.JobID.toString())}/job`,
			request: param,
			baseURL: this.baseURL,
			headers: reqHeader,
			options: reqOption,
		};
		await this.callMiddleware(this.beforeMiddleware, context);
		const url = `${this.baseURL}/service/user2/${encodeURI(param.UserID.toString())}/${encodeURI(param.JobID.toString())}/job`;
		const resp = await fetch(
			url,
			{
				method: "PUT",
				body: JSON.stringify(this.getRequestObject(param, excludeParams)),
				headers: reqHeader,
				...reqOption,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			const responseText = await resp.text();
			throw new ApiError(resp, responseText);
		}
		const res = (await resp.json()) as ServiceUser2UserIDJobIDPutJobResponse;
		context.response = res;
		await this.callMiddleware(this.afterMiddleware, context);
		return res;
	}
}

class ServiceUserClient {
	private beforeMiddleware: ApiClientMiddlewareFunc[] = [];
	private afterMiddleware: ApiClientMiddlewareFunc[] = [];
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

	async callMiddleware(
		middlewares: ApiClientMiddlewareFunc[],
		context: MiddlewareContext
	) {
		for (const m of middlewares) {
			const func: ApiClientMiddlewareFunc = m;
			const mr = await func(context);
			if (typeof mr === 'boolean') {
				if (!mr) {
					break;
				}
			} else {
				if (mr === MiddlewareResult.CONTINUE) {
					continue;
				} else if (mr === MiddlewareResult.MIDDLEWARE_STOP) {
					break;
				} else if (mr === MiddlewareResult.STOP) {
					throw new ApiMiddlewareStop();
				}
			}
		}
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

		const reqHeader = {
			...this.headers,
			...headers,
			...mockHeaders,
		};
		const reqOption = {
			...this.options,
			...options,
		};
		const context: MiddlewareContext = {
			httpMethod: 'GET',
			endpoint: `${this.baseURL}/service/user/`,
			request: param,
			baseURL: this.baseURL,
			headers: reqHeader,
			options: reqOption,
		};
		await this.callMiddleware(this.beforeMiddleware, context);
		const url = `${this.baseURL}/service/user/?` + (new URLSearchParams(this.getRequestObject(param, excludeParams))).toString();
		const resp = await fetch(
			url,
			{
				method: "GET",
				headers: reqHeader,
				...reqOption,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			const responseText = await resp.text();
			throw new ApiError(resp, responseText);
		}
		await resp.text();
		const res = {} as ServiceUserGetResponse;
		context.response = res;
		await this.callMiddleware(this.afterMiddleware, context);
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

		const reqHeader = {
			...this.headers,
			...headers,
			...mockHeaders,
		};
		const reqOption = {
			...this.options,
			...options,
		};
		const context: MiddlewareContext = {
			httpMethod: 'POST',
			endpoint: `${this.baseURL}/service/user/update_user_name`,
			request: param,
			baseURL: this.baseURL,
			headers: reqHeader,
			options: reqOption,
		};
		await this.callMiddleware(this.beforeMiddleware, context);
		const url = `${this.baseURL}/service/user/update_user_name`;
		const resp = await fetch(
			url,
			{
				method: "POST",
				body: JSON.stringify(this.getRequestObject(param, excludeParams)),
				headers: reqHeader,
				...reqOption,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			const responseText = await resp.text();
			throw new ApiError(resp, responseText);
		}
		const res = (await resp.json()) as ServiceUserPostUpdateUserNameResponse;
		context.response = res;
		await this.callMiddleware(this.afterMiddleware, context);
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

		const reqHeader = {
			...this.headers,
			...headers,
			...mockHeaders,
		};
		const reqOption = {
			...this.options,
			...options,
		};
		const context: MiddlewareContext = {
			httpMethod: 'POST',
			endpoint: `${this.baseURL}/service/user/update_user_password`,
			request: param,
			baseURL: this.baseURL,
			headers: reqHeader,
			options: reqOption,
		};
		await this.callMiddleware(this.beforeMiddleware, context);
		const url = `${this.baseURL}/service/user/update_user_password`;
		const resp = await fetch(
			url,
			{
				method: "POST",
				body: JSON.stringify(this.getRequestObject(param, excludeParams)),
				headers: reqHeader,
				...reqOption,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			const responseText = await resp.text();
			throw new ApiError(resp, responseText);
		}
		const res = (await resp.json()) as ServiceUserPostUpdateUserPasswordResponse;
		context.response = res;
		await this.callMiddleware(this.afterMiddleware, context);
		return res;
	}
}

export class APIClient {
	private headers: {[key: string]: string};
	private options: {[key: string]: any};
	private baseURL: string;

	private beforeMiddleware: ApiClientMiddlewareFunc[] = [];
	private afterMiddleware: ApiClientMiddlewareFunc[] = [];

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

	async callMiddleware(
		middlewares: ApiClientMiddlewareFunc[],
		context: MiddlewareContext
	) {
		for (const m of middlewares) {
			const func: ApiClientMiddlewareFunc = m;
			const mr = await func(context);
			if (typeof mr === 'boolean') {
				if (!mr) {
					break;
				}
			} else {
				if (mr === MiddlewareResult.CONTINUE) {
					continue;
				} else if (mr === MiddlewareResult.MIDDLEWARE_STOP) {
					break;
				} else if (mr === MiddlewareResult.STOP) {
					throw new ApiMiddlewareStop();
				}
			}
		}
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

		const reqHeader = {
			...this.headers,
			...headers,
			...mockHeaders,
		};
		const reqOption = {
			...this.options,
			...options,
		};
		const context: MiddlewareContext = {
			httpMethod: 'GET',
			endpoint: `${this.baseURL}/`,
			request: param,
			baseURL: this.baseURL,
			headers: reqHeader,
			options: reqOption,
		};
		await this.callMiddleware(this.beforeMiddleware, context);
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
		context.response = res;
		await this.callMiddleware(this.afterMiddleware, context);
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

		const reqHeader = {
			...this.headers,
			...headers,
			...mockHeaders,
		};
		const reqOption = {
			...this.options,
			...options,
		};
		const context: MiddlewareContext = {
			httpMethod: 'POST',
			endpoint: `${this.baseURL}/create_table`,
			request: param,
			baseURL: this.baseURL,
			headers: reqHeader,
			options: reqOption,
		};
		await this.callMiddleware(this.beforeMiddleware, context);
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
		context.response = res;
		await this.callMiddleware(this.afterMiddleware, context);
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

		const reqHeader = {
			...this.headers,
			...headers,
			...mockHeaders,
		};
		const reqOption = {
			...this.options,
			...options,
		};
		const context: MiddlewareContext = {
			httpMethod: 'POST',
			endpoint: `${this.baseURL}/create_user`,
			request: param,
			baseURL: this.baseURL,
			headers: reqHeader,
			options: reqOption,
		};
		await this.callMiddleware(this.beforeMiddleware, context);
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
		context.response = res;
		await this.callMiddleware(this.afterMiddleware, context);
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

export class ApiMiddlewareStop extends Error {}

export interface MockOption {
	wait_ms: number;
	target_file: string;
}
