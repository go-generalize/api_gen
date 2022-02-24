// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: devel

export type GetRequest = {
	Enum: "A" | "B" | "C";
	Param: string;
	Time: string;
}
export type GetResponse = {
	Data: string;
}
export type PostCreateTableRequest = {
	File?: File | Blob;
	Files?: (File | Blob)[];
	Flag: number;
	ID: string;
	Text: string;
	map: {[key: number]: number};
}
export type PostCreateTableResponse = {
	ID: string;
	RequestTime: string;
}
export type PostCreateUserRequest = {
	Birthday: string;
	Gender: number;
	ID: string;
	Password: string;
	Roles: (Role | null)[] | null;
}
export type PostCreateUserResponse = {
	CreatedType: 0 | 1 | 2;
	Message: string;
	RequestedAt: string;
	Status: boolean;
}
export type Role = {
	ID: number;
	Name: string;
	RecursionRoles: Role[] | null;
}
