// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: 0.4.0

import {
	PostCreateTableRequest as PostCreateTableRequest,
	PostCreateTableResponse as PostCreateTableResponse,
	PostCreateUserRequest as PostCreateUserRequest,
	PostCreateUserResponse as PostCreateUserResponse,
} from './classes//types';
export {
	PostCreateTableRequest as PostCreateTableRequest,
	PostCreateTableResponse as PostCreateTableResponse,
	PostCreateUserRequest as PostCreateUserRequest,
	PostCreateUserResponse as PostCreateUserResponse,
} from './classes//types';
import {
	GetArticleRequest as ServiceGetArticleRequest,
	GetArticleResponse as ServiceGetArticleResponse,
} from './classes/service/types';
export {
	GetArticleRequest as ServiceGetArticleRequest,
	GetArticleResponse as ServiceGetArticleResponse,
} from './classes/service/types';
import {
	PostUpdateUserNameRequest as ServiceUserPostUpdateUserNameRequest,
	PostUpdateUserNameResponse as ServiceUserPostUpdateUserNameResponse,
	PostUpdateUserPasswordRequest as ServiceUserPostUpdateUserPasswordRequest,
	PostUpdateUserPasswordResponse as ServiceUserPostUpdateUserPasswordResponse,
} from './classes/service/user/types';
export {
	PostUpdateUserNameRequest as ServiceUserPostUpdateUserNameRequest,
	PostUpdateUserNameResponse as ServiceUserPostUpdateUserNameResponse,
	PostUpdateUserPasswordRequest as ServiceUserPostUpdateUserPasswordRequest,
	PostUpdateUserPasswordResponse as ServiceUserPostUpdateUserPasswordResponse,
} from './classes/service/user/types';
import {
	PutJobRequest as ServiceUser2UserIDJobIDPutJobRequest,
	PutJobResponse as ServiceUser2UserIDJobIDPutJobResponse,
} from './classes/service/user2/_userID/_JobID/types';
export {
	PutJobRequest as ServiceUser2UserIDJobIDPutJobRequest,
	PutJobResponse as ServiceUser2UserIDJobIDPutJobResponse,
} from './classes/service/user2/_userID/_JobID/types';
import {
	GetUserJobGetRequest as ServiceUser2UserIDGetUserJobGetRequest,
	GetUserJobGetResponse as ServiceUser2UserIDGetUserJobGetResponse,
} from './classes/service/user2/_userID/types';
export {
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
export {
	GetUserRequest as ServiceUser2GetUserRequest,
	GetUserResponse as ServiceUser2GetUserResponse,
	PostUpdateUserNameRequest as ServiceUser2PostUpdateUserNameRequest,
	PostUpdateUserNameResponse as ServiceUser2PostUpdateUserNameResponse,
	PostUpdateUserPasswordRequest as ServiceUser2PostUpdateUserPasswordRequest,
	PostUpdateUserPasswordResponse as ServiceUser2PostUpdateUserPasswordResponse,
} from './classes/service/user2/types';


class PropsClient {

	constructor(private headers: {[key: string]: string}, private options: {[key: string]: any}, private baseURL: string) {

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

}

class ServiceClient {

	public static_page: ServiceStaticPageClient;
	public table: ServiceTableClient;
	public user: ServiceUserClient;
	public user2: ServiceUser2Client;
	constructor(private headers: {[key: string]: string}, private options: {[key: string]: any}, private baseURL: string) {

		this.static_page = new ServiceStaticPageClient(headers, options, baseURL);
		this.table = new ServiceTableClient(headers, options, baseURL);
		this.user = new ServiceUserClient(headers, options, baseURL);
		this.user2 = new ServiceUser2Client(headers, options, baseURL);
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

	async getArticle(
		param: ServiceGetArticleRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceGetArticleResponse> {
	    const excludeParams: string[] = [];
		const resp = await fetch(
			`${this.baseURL}/service/article?` + (new URLSearchParams(this.getRequestObject(param, excludeParams))).toString(),
			{
				method: "GET",
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
		return (await resp.json()) as ServiceGetArticleResponse;
	}
}

class ServiceStaticPageClient {

	constructor(private headers: {[key: string]: string}, private options: {[key: string]: any}, private baseURL: string) {

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

	async getStaticPage(
		param: ServiceStaticPageGetStaticPageRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceStaticPageGetStaticPageResponse> {
	    const excludeParams: string[] = [];
		const resp = await fetch(
			`${this.baseURL}/service/static_page/static_page?` + (new URLSearchParams(this.getRequestObject(param, excludeParams))).toString(),
			{
				method: "GET",
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
		return (await resp.json()) as ServiceStaticPageGetStaticPageResponse;
	}
}

class ServiceTableClient {

	constructor(private headers: {[key: string]: string}, private options: {[key: string]: any}, private baseURL: string) {

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

}

class ServiceUser2Client {

	public _userID: ServiceUser2UserIDClient;
	constructor(private headers: {[key: string]: string}, private options: {[key: string]: any}, private baseURL: string) {

		this._userID = new ServiceUser2UserIDClient(headers, options, baseURL);
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

	async getUser(
		param: ServiceUser2GetUserRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceUser2GetUserResponse> {
	    const excludeParams: string[] = ['id', ];
		const resp = await fetch(
			`${this.baseURL}/service/user2/${encodeURI(param.id.toString())}?` + (new URLSearchParams(this.getRequestObject(param, excludeParams))).toString(),
			{
				method: "GET",
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
		return (await resp.json()) as ServiceUser2GetUserResponse;
	}
	async postUpdateUserName(
		param: ServiceUser2PostUpdateUserNameRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceUser2PostUpdateUserNameResponse> {
	    const excludeParams: string[] = [];
		const resp = await fetch(
			`${this.baseURL}/service/user2/update_user_name`,
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
		return (await resp.json()) as ServiceUser2PostUpdateUserNameResponse;
	}
	async postUpdateUserPassword(
		param: ServiceUser2PostUpdateUserPasswordRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceUser2PostUpdateUserPasswordResponse> {
	    const excludeParams: string[] = [];
		const resp = await fetch(
			`${this.baseURL}/service/user2/update_user_password`,
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
		return (await resp.json()) as ServiceUser2PostUpdateUserPasswordResponse;
	}
}

class ServiceUser2UserIDClient {

	public _JobID: ServiceUser2UserIDJobIDClient;
	constructor(private headers: {[key: string]: string}, private options: {[key: string]: any}, private baseURL: string) {

		this._JobID = new ServiceUser2UserIDJobIDClient(headers, options, baseURL);
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

	async getUserJobGet(
		param: ServiceUser2UserIDGetUserJobGetRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceUser2UserIDGetUserJobGetResponse> {
	    const excludeParams: string[] = ['UserID', ];
		const resp = await fetch(
			`${this.baseURL}/service/user2/${encodeURI(param.UserID.toString())}/user_job_get?` + (new URLSearchParams(this.getRequestObject(param, excludeParams))).toString(),
			{
				method: "GET",
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
		return (await resp.json()) as ServiceUser2UserIDGetUserJobGetResponse;
	}
}

class ServiceUser2UserIDJobIDClient {

	constructor(private headers: {[key: string]: string}, private options: {[key: string]: any}, private baseURL: string) {

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

	async putJob(
		param: ServiceUser2UserIDJobIDPutJobRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceUser2UserIDJobIDPutJobResponse> {
	    const excludeParams: string[] = ['UserID', 'JobID', ];
		const resp = await fetch(
			`${this.baseURL}/service/user2/${encodeURI(param.UserID.toString())}/${encodeURI(param.JobID.toString())}/job`,
			{
				method: "PUT",
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
		return (await resp.json()) as ServiceUser2UserIDJobIDPutJobResponse;
	}
}

class ServiceUserClient {

	constructor(private headers: {[key: string]: string}, private options: {[key: string]: any}, private baseURL: string) {

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

	async postUpdateUserName(
		param: ServiceUserPostUpdateUserNameRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceUserPostUpdateUserNameResponse> {
	    const excludeParams: string[] = [];
		const resp = await fetch(
			`${this.baseURL}/service/user/update_user_name`,
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
		return (await resp.json()) as ServiceUserPostUpdateUserNameResponse;
	}
	async postUpdateUserPassword(
		param: ServiceUserPostUpdateUserPasswordRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceUserPostUpdateUserPasswordResponse> {
	    const excludeParams: string[] = [];
		const resp = await fetch(
			`${this.baseURL}/service/user/update_user_password`,
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
		return (await resp.json()) as ServiceUserPostUpdateUserPasswordResponse;
	}
}


export class APIClient {
	private headers: {[key: string]: string};
	private options: {[key: string]: any};
	private baseURL: string;

	public props: PropsClient;
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


		this.props = new PropsClient(headers, this.options, this.baseURL);
		this.service = new ServiceClient(headers, this.options, this.baseURL);
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

	async postCreateTable(
		param: PostCreateTableRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<PostCreateTableResponse> {
	    const excludeParams: string[] = [];
		const resp = await fetch(
			`${this.baseURL}/create_table`,
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
		return (await resp.json()) as PostCreateTableResponse;
	}

	async postCreateUser(
		param: PostCreateUserRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<PostCreateUserResponse> {
	    const excludeParams: string[] = [];
		const resp = await fetch(
			`${this.baseURL}/create_user`,
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
