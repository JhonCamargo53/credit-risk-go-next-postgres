export interface Asset{
    ID: number;
    name: string;
    description: string;
    status:boolean
}

export interface AssetForm{
    name: string;
    description: string;
}