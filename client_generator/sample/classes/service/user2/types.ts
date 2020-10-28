// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: unknown

export type GetUserRequest = {
	id: string;
	search_request: string;
}
export type GetUserResponse = {
	ID: string;
	RequestTime: string;
	SearchRequest: string;
}
export type PostUpdateUserNameRequest = {
	Name: string;
}
export type PostUpdateUserNameResponse = {
	Message: string;
	RequestTime: string;
	Status: boolean;
}
export type PostUpdateUserPasswordRequest = {
	Password: string;
	PasswordConfirm: string;
}
export type PostUpdateUserPasswordResponse = {
	Message: string;
	RequestTime: string;
	Status: boolean;
}
