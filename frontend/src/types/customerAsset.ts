export interface CustomerAsset {
    ID: number;
    creditRequestId: number;
    customerId: number;
    assetId: number;
    description: string;
    marketValue: number;
    status: string;
    CreatedAt: string;
    UpdatedAt: string;
}

export interface CustomerAssetForm {
    creditRequestId: number;
    customerId: number;
    assetId: number;
    description: string;
    marketValue: number;
}