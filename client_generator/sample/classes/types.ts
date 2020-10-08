// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: unknown

export type PostCreateTableRequest = {
	ID: string;
	Text: string;
}
export type PostCreateTableResponse = {
	ID: string;
	Payload: Table;
}
export type PostCreateUserRequest = {
	Gender: number;
	ID: string;
	Password: string;
}
export type PostCreateUserResponse = {
	CreatedType: 0 | 1 | 2;
	Message: string;
	Status: boolean;
}
export type Pos = {
	X: number;
	Y: number;
}
export type Table = {
	Pos: Pos;
}
