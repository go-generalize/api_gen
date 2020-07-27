// THIS CODE WAS GENERATED AUTOMATICALLY
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS

import { GetArticleRequest as ServiceGetArticleRequest } from './classes/service/GetArticleRequest';
export { GetArticleRequest as ServiceGetArticleRequest } from './classes/service/GetArticleRequest';
import { GetArticleResponse as ServiceGetArticleResponse } from './classes/service/GetArticleResponse';
export { GetArticleResponse as ServiceGetArticleResponse } from './classes/service/GetArticleResponse';
import { GetUserJobGetRequest as ServiceUser2UserIDGetUserJobGetRequest } from './classes/service/user2/_userID/GetUserJobGetRequest';
export { GetUserJobGetRequest as ServiceUser2UserIDGetUserJobGetRequest } from './classes/service/user2/_userID/GetUserJobGetRequest';
import { GetUserJobGetResponse as ServiceUser2UserIDGetUserJobGetResponse } from './classes/service/user2/_userID/GetUserJobGetResponse';
export { GetUserJobGetResponse as ServiceUser2UserIDGetUserJobGetResponse } from './classes/service/user2/_userID/GetUserJobGetResponse';
import { GetUserRequest as ServiceUser2GetUserRequest } from './classes/service/user2/GetUserRequest';
export { GetUserRequest as ServiceUser2GetUserRequest } from './classes/service/user2/GetUserRequest';
import { GetUserResponse as ServiceUser2GetUserResponse } from './classes/service/user2/GetUserResponse';
export { GetUserResponse as ServiceUser2GetUserResponse } from './classes/service/user2/GetUserResponse';
import { PostCreateTableRequest as PostCreateTableRequest } from './classes//PostCreateTableRequest';
export { PostCreateTableRequest as PostCreateTableRequest } from './classes//PostCreateTableRequest';
import { PostCreateTableResponse as PostCreateTableResponse } from './classes//PostCreateTableResponse';
export { PostCreateTableResponse as PostCreateTableResponse } from './classes//PostCreateTableResponse';
import { PostCreateUserRequest as PostCreateUserRequest } from './classes//PostCreateUserRequest';
export { PostCreateUserRequest as PostCreateUserRequest } from './classes//PostCreateUserRequest';
import { PostCreateUserResponse as PostCreateUserResponse } from './classes//PostCreateUserResponse';
export { PostCreateUserResponse as PostCreateUserResponse } from './classes//PostCreateUserResponse';
import { PostUpdateUserNameRequest as ServiceUserPostUpdateUserNameRequest } from './classes/service/user/PostUpdateUserNameRequest';
export { PostUpdateUserNameRequest as ServiceUserPostUpdateUserNameRequest } from './classes/service/user/PostUpdateUserNameRequest';
import { PostUpdateUserNameRequest as ServiceUser2PostUpdateUserNameRequest } from './classes/service/user2/PostUpdateUserNameRequest';
export { PostUpdateUserNameRequest as ServiceUser2PostUpdateUserNameRequest } from './classes/service/user2/PostUpdateUserNameRequest';
import { PostUpdateUserNameResponse as ServiceUser2PostUpdateUserNameResponse } from './classes/service/user2/PostUpdateUserNameResponse';
export { PostUpdateUserNameResponse as ServiceUser2PostUpdateUserNameResponse } from './classes/service/user2/PostUpdateUserNameResponse';
import { PostUpdateUserNameResponse as ServiceUserPostUpdateUserNameResponse } from './classes/service/user/PostUpdateUserNameResponse';
export { PostUpdateUserNameResponse as ServiceUserPostUpdateUserNameResponse } from './classes/service/user/PostUpdateUserNameResponse';
import { PostUpdateUserPasswordRequest as ServiceUser2PostUpdateUserPasswordRequest } from './classes/service/user2/PostUpdateUserPasswordRequest';
export { PostUpdateUserPasswordRequest as ServiceUser2PostUpdateUserPasswordRequest } from './classes/service/user2/PostUpdateUserPasswordRequest';
import { PostUpdateUserPasswordRequest as ServiceUserPostUpdateUserPasswordRequest } from './classes/service/user/PostUpdateUserPasswordRequest';
export { PostUpdateUserPasswordRequest as ServiceUserPostUpdateUserPasswordRequest } from './classes/service/user/PostUpdateUserPasswordRequest';
import { PostUpdateUserPasswordResponse as ServiceUserPostUpdateUserPasswordResponse } from './classes/service/user/PostUpdateUserPasswordResponse';
export { PostUpdateUserPasswordResponse as ServiceUserPostUpdateUserPasswordResponse } from './classes/service/user/PostUpdateUserPasswordResponse';
import { PostUpdateUserPasswordResponse as ServiceUser2PostUpdateUserPasswordResponse } from './classes/service/user2/PostUpdateUserPasswordResponse';
export { PostUpdateUserPasswordResponse as ServiceUser2PostUpdateUserPasswordResponse } from './classes/service/user2/PostUpdateUserPasswordResponse';
import { PutJobRequest as ServiceUser2UserIDJobIDPutJobRequest } from './classes/service/user2/_userID/_JobID/PutJobRequest';
export { PutJobRequest as ServiceUser2UserIDJobIDPutJobRequest } from './classes/service/user2/_userID/_JobID/PutJobRequest';
import { PutJobResponse as ServiceUser2UserIDJobIDPutJobResponse } from './classes/service/user2/_userID/_JobID/PutJobResponse';
export { PutJobResponse as ServiceUser2UserIDJobIDPutJobResponse } from './classes/service/user2/_userID/_JobID/PutJobResponse';



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
			throw new Error(resp.statusText + ": " + await resp.text());
		}

		return new ServiceGetArticleResponse(await resp.json());
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
			throw new Error(resp.statusText + ": " + await resp.text());
		}

		return new ServiceStaticPageGetStaticPageResponse(await resp.json());
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
			throw new Error(resp.statusText + ": " + await resp.text());
		}

		return new ServiceUser2GetUserResponse(await resp.json());
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
			throw new Error(resp.statusText + ": " + await resp.text());
		}

		return new ServiceUser2PostUpdateUserNameResponse(await resp.json());
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
			throw new Error(resp.statusText + ": " + await resp.text());
		}

		return new ServiceUser2PostUpdateUserPasswordResponse(await resp.json());
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
			throw new Error(resp.statusText + ": " + await resp.text());
		}

		return new ServiceUser2UserIDGetUserJobGetResponse(await resp.json());
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
			throw new Error(resp.statusText + ": " + await resp.text());
		}

		return new ServiceUser2UserIDJobIDPutJobResponse(await resp.json());
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
			throw new Error(resp.statusText + ": " + await resp.text());
		}

		return new ServiceUserPostUpdateUserNameResponse(await resp.json());
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
			throw new Error(resp.statusText + ": " + await resp.text());
		}

		return new ServiceUserPostUpdateUserPasswordResponse(await resp.json());
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
			throw new Error(resp.statusText + ": " + await resp.text());
		}

		return new PostCreateUserResponse(await resp.json());
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
			throw new Error(resp.statusText + ": " + await resp.text());
		}

		return new PostCreateTableResponse(await resp.json());
	}

}
