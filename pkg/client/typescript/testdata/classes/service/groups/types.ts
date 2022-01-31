// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: devel

export type Company = {
	Metadata: Metadata;
}
export type Department = {
	Metadata: Metadata;
}
export type GetGroupsRequest = {}
export type GetGroupsResponse = {
	Companies: Company[] | null;
	Departments: Department[] | null;
}
export type Metadata = {
	CreatedAt: string;
	DeletedAt: string | null;
	ID: string;
	Name: string;
	UpdatedAt: string;
}
