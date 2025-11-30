export interface User {
    ID: number;
    name: string;
    roleId: number;
    email: string;
    status: boolean;
    UpdatedAt: string;
    CreatedAt: string;
}

export interface UserForm {
    name: string,
    roleId:number
    email: string,
    password:string
}